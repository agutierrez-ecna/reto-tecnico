package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	_ "go-api/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

// @title API en Go - Interseguro Challenge
// @version 1.0
// @description API que procesa y rota matrices rectangulares protegida con JWT.
// @host localhost:8080
// @BasePath /

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Escribe 'Bearer ' seguido de tu token JWT generador en /login (Ejemplo: Bearer eyJhbGci...).

var jwtSecret = []byte("secreto-interseguro")

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type MatrixRequest struct {
	Matrix [][]float64 `json:"matrix"`
}

type MatrixResponse struct {
	RotatedMatrix [][]float64 `json:"rotated_matrix"`
}

func rotateMatrix(matrix [][]float64) [][]float64 {
	rows := len(matrix)
	if rows == 0 {
		return [][]float64{}
	}
	cols := len(matrix[0])

	rotated := make([][]float64, cols)
	for i := range rotated {
		rotated[i] = make([]float64, rows)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			rotated[j][rows-1-i] = matrix[i][j]
		}
	}
	return rotated
}

// @Summary Iniciar sesión para obtener Token JWT
// @Description Autentica al usuario y devuelve un token JWT válido
// @Accept json
// @Produce json
// @Param request body Credentials true "Credenciales de acceso"
// @Success 200 {object} map[string]string
// @Router /login [post]
func Login(c *fiber.Ctx) error {
	var creds Credentials
	if err := c.BodyParser(&creds); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Datos inválidos"})
	}

	if creds.Username != "admin" || creds.Password != "123456" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Credenciales incorrectas"})
	}

	claims := jwt.MapClaims{
		"username": creds.Username,
		"exp":      time.Now().Add(time.Hour * 2).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString(jwtSecret)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "No se pudo generar el token"})
	}

	return c.JSON(fiber.Map{"token": t})
}

// @Summary Recibe y rota una matriz (Protegido con JWT)
// @Description Recibe una matriz de números, la rota 90 grados y la reenvía a Node.js
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body MatrixRequest true "Matriz de entrada"
// @Success 200 {object} map[string]interface{}
// @Router /api/process-matrix [post]
func ProcessMatrixHandler(c *fiber.Ctx) error {
	var req MatrixRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "El cuerpo de la solicitud (Body) tiene un formato JSON inválido.",
		})
	}

	if len(req.Matrix) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "La matriz no puede estar vacía.",
		})
	}

	rotated := rotateMatrix(req.Matrix)
	payload := MatrixResponse{RotatedMatrix: rotated}
	jsonPayload, _ := json.Marshal(payload)

	nodeServiceURL := os.Getenv("NODE_SERVICE_URL")
	if nodeServiceURL == "" {
		nodeServiceURL = "http://node-api:3000/analyze-matrix"
	}
	resp, err := http.Post(nodeServiceURL, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"error": "Error de comunicación con el microservicio de análisis en Node.js.",
		})
	}
	defer resp.Body.Close()

	var nodeResult map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&nodeResult)

	return c.JSON(fiber.Map{
		"message":        "Procesado exitosamente",
		"rotated_matrix":  rotated,
		"node_statistics": nodeResult,
	})
}

func main() {
	app := fiber.New()

	app.Use(cors.New())

    app.Static("/", "../frontend")

	app.Get("/swagger/*", fiberSwagger.WrapHandler)

	app.Post("/login", Login)

	api := app.Group("/api")
	api.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtSecret,
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "No autorizado: Token JWT faltante o inválido",
			})
		},
	}))
	api.Post("/process-matrix", ProcessMatrixHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("API en Go ejecutándose en el puerto %s con seguridad JWT y CORS", port)
	log.Fatal(app.Listen(":" + port))
}