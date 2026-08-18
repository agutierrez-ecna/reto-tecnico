## Desafío Técnico - Arquitectura de Microservicios

Solución de arquitectura de microservicios distribuida para el procesamiento, rotación de 90 grados y análisis estadístico de matrices rectangulares.

## Demo y Despliegue en la Nube (GCP Cloud Run)

Los servicios se encuentran 100% operativos y desplegados en Google Cloud Run:
  - Aplicación Web (Frontend + Go API): https://reto-tecnico-go-1021109336611.europe-west1.run.app
  - Microservicio de Análisis (Node.js API): https://reto-tecnico-nodejs-1021109336611.europe-west1.run.app

## Descripción del Proyecto

El sistema está compuesto por dos microservicios y un cliente web con autenticación JWT:

1. API 1 (Go + Fiber):
  Sirve la interfaz web, autentica credenciales para generar tokens JWT, recibe la matriz, ejecuta el algoritmo de rotación de 90 grados y reenvía el resultado al microservicio de análisis.
2. API 2 (Node.js + Express):
  Recibe la matriz rotada y calcula métricas (valor máximo, mínimo, promedio, suma total y validación de matriz diagonal).
3. Frontend: 
  Interfaz web interactiva construida en HTML y Tailwind CSS que gestiona la autenticación JWT y el procesamiento de matrices en tiempo real.
## Tecnologías Utilizadas

- Lenguajes: Go, JavaScript (Node.js), HTML5, CSS3 (Tailwind CSS)
- Frameworks: Fiber (Go), Express.js (Node.js)
- Seguridad: JSON Web Tokens (JWT)
- Documentación: Swagger / OpenAPI
- Contenerización y Despliegue: Docker, Docker Compose(Local), Google Cloud Run, Cloud Build
- Testing: Go Testing Package, Jest, Supertest

## Credenciales de Prueba

Para probar el flujo en Swagger, Postman o la app web:

- Usuario: `admin`
- Contraseña: `123456`

## Puntos de Enlace (Endpoints)

| Servicio | Método | Ruta | Descripción |
| Go API | `POST` | `/login` | Autentica credenciales y devuelve el token JWT. |
| Go API | `POST` | `/api/process-matrix` | Recibe la matriz, la rota y consume a Node.js -(Requiere Header `Authorization: Bearer <token>`)-. |
| Node.js API | `POST` | `/analyze-matrix` | Recibe la matriz procesada y calcula métricas estadísticas. |

## Despliegue Local con Docker Compose

Si deseas ejecutar el proyecto localmente en contenedores aislados:

docker-compose up --build
