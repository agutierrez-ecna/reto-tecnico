Desafío Técnico - Interseguro
Solución de arquitectura de microservicios para el procesamiento, rotación y análisis estadístico de matrices.  

Descripción del Proyecto
El sistema está compuesto por dos microservicios y un frontend con autenticación JWT:  

- API 1 (Go + Fiber): Recibe una matriz rectangular, aplica un algoritmo de rotación de 90 grados, valida la seguridad con JWT y reenvía los datos al segundo microservicio.  

- API 2 (Node.js + Express): Recibe la matriz procesada y calcula métricas estadísticas (valor máximo, valor mínimo, promedio, suma total y verificación de matriz diagonal).  

- Frontend: Interfaz web interactiva construida en HTML y JavaScript para gestionar autenticación e ingresar datos a la API.  

Tecnologías Utilizadas
Lenguajes: Go, JavaScript (Node.js), HTML5  
Frameworks: Fiber (Go), Express.js (Node.js)  
Seguridad: JSON Web Tokens (JWT)  
Documentación: Swagger / OpenAPI  
Contenerización: Docker, Docker Compose  
Testing: Go Testing Package, Jest, Supertest  

Documentación de la API (Swagger)
La API construida en Go incluye especificación Swagger integrada:  
Swagger UI: http://localhost:8080/swagger/index.html

Puntos de Enlace (Endpoints)
Servicio    Método  Ruta                    Descripción
Go API      POST    /login                  Autentica credenciales y genera el token JWT  
Go API      POST    /api/process-matrix     Recibe matriz, la rota y consume a Node.js (Requiere JWT)  Node.js API POST    /analyze-matrix         Recibe la matriz y calcula estadísticas  

Despliegue Local con Docker Compose
Para construir y levantar todos los servicios en contenedores aislados, ejecuta el siguiente comando en la raíz del repositorio:  
- docker-compose up --build
URLs de Acceso
API Go: http://localhost:8080
API Node.js: http://localhost:3000

Ejecución de Pruebas Automatizadas
Pruebas en Go (API 1)
cd go-api
go test -v ./...
Pruebas en Node.js (API 2)
cd node-api
npm test
