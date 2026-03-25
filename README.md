# Weather Monorepo

API REST que consume dos protocolos distintos — REST y SOAP — desde una librería interna compartida, construida en Go con arquitectura de monorepo.

## Estructura
```
weather-monorepo/
├── apiclient/          # Librería interna compartida
│   ├── client_rest.go  # Cliente REST (Open-Meteo)
│   ├── client_soap.go  # Cliente SOAP (CDYNE CountryInfo)
│   ├── envelope.go     # Structs XML para SOAP
│   ├── errors.go       # Errores personalizados
│   └── types.go        # Modelos compartidos
├── weather-app/        # Servicio principal
│   ├── main.go         # Entrypoint
│   ├── server.go       # Configuración del servidor
│   ├── routes.go       # Definición de rutas
│   ├── weather_handler.go     # HTTP handlers
│   ├── weather_service.go     # Lógica de negocio
│   ├── weather_service_test.go
│   └── weather_handler_test.go
├── Dockerfile
├── go.work
└── go.work.sum
```

## Tecnologías

- **Go 1.23**
- **Go Workspaces** — monorepo con módulos independientes
- **Open-Meteo API** — clima actual gratuito sin API key
- **CDYNE CountryInfo SOAP** — obtención de capital por código de país
- **testify** — assertions en tests unitarios
- **Docker** — imagen multistage

## Endpoints

### `POST /api/weather/rest`
Obtiene el clima actual de una ciudad directamente vía REST.

**Request:**
```json
{
  "city": "Guatemala"
}
```

**Response:**
```json
{
  "city": "Guatemala",
  "country": "GT",
  "temperature": 22.5,
  "feels_like": 23.1,
  "description": "parcialmente nublado",
  "humidity": 75,
  "source": "REST"
}
```

---

### `POST /api/weather/soap`
Obtiene la capital del país vía SOAP y luego consulta el clima de esa capital vía REST.

**Request:**
```json
{
  "country": "GT"
}
```

**Response:**
```json
{
  "city": "GT",
  "country": "GT",
  "temperature": 19.8,
  "feels_like": 20.1,
  "description": "cielo despejado en Guatemala City",
  "humidity": 68,
  "source": "SOAP + REST"
}
```

**Códigos de país soportados:** `GT`, `US`, `MX`, `JP`, `ES`, `DE`, `FR`, `BR`, `AR`, `CO` y cualquier código ISO 3166-1 alpha-2.

## Flujo interno
```
REST:  ciudad → Geocoding API → lat/lon → clima actual
SOAP:  código país → capital (SOAP) → lat/lon → clima actual (REST)
```

## Correr localmente
```bash
# Clonar el repo
git clone https://github.com/Xavez05/weather-monorepo.git
cd weather-monorepo

# Correr la app
cd weather-app
go run .

# El servidor queda en http://localhost:8080
```

## Correr con Docker
```bash
docker build -t weather-app .
docker run -p 8080:8080 weather-app
```

## Tests
```bash
cd weather-app
go test ./... -v
```

Cobertura incluye:
- Casos exitosos REST y SOAP
- Errores de servicio (ciudad no encontrada, país inválido)
- Validación de request body vacío
- HTTP status codes correctos (200, 400, 502)

## Principios aplicados

- **SOLID** — interfaces `WeatherFetcher` para inversión de dependencias, responsabilidad única por archivo
- **Clean Code** — handlers, servicios y rutas separados
- **Go Workspaces** — `apiclient` es reutilizable como librería independiente en otros proyectos del monorepo
- **Multistage Docker** — imagen final liviana basada en `alpine`