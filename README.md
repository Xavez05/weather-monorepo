# Weather Monorepo

![CI](https://github.com/Xavez05/weather-monorepo/actions/workflows/ci.yml/badge.svg)

API that consumes two different protocols — REST and SOAP — from a shared internal library, built in Go with a monorepo architecture.

## Structure
```
weather-monorepo/
├── apiclient/               # Shared internal library
│   ├── client_rest.go       # REST client (Open-Meteo)
│   ├── client_soap.go       # SOAP client (CDYNE CountryInfo)
│   ├── envelope.go          # XML structs for SOAP
│   ├── errors.go            # Custom error types
│   └── types.go             # Shared models
├── weather-app/             # Main service
│   ├── main.go              # Entrypoint
│   ├── server.go            # Server setup
│   ├── routes.go            # Route definitions
│   ├── weather_handler.go   # HTTP handlers
│   ├── weather_service.go   # Business logic
│   ├── weather_service_test.go
│   └── weather_handler_test.go
├── .github/
│   └── workflows/
│       └── ci.yml           # GitHub Actions CI
├── Dockerfile
├── go.work
└── go.work.sum
```

## Tech Stack

- **Go 1.23**
- **Go Workspaces** — monorepo with independent modules
- **Open-Meteo API** — free weather data, no API key required
- **CDYNE CountryInfo SOAP** — capital city lookup by country code
- **testify** — unit test assertions
- **Docker** — multistage image
- **GitHub Actions** — automated CI on every push

## CI/CD

The pipeline runs automatically on every push to `main` with three sequential jobs:
```
Build → Test → Docker Build
```

It can also be triggered manually from the **Actions** tab in GitHub using the **Run workflow** button.

## Endpoints

### `POST /api/weather/rest`
Returns current weather for a city via REST.

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
  "description": "partly cloudy",
  "humidity": 75,
  "source": "REST"
}
```

---

### `POST /api/weather/soap`
Fetches the country capital via SOAP, then retrieves weather for that capital via REST.

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
  "description": "clear sky in Guatemala City",
  "humidity": 68,
  "source": "SOAP + REST"
}
```

**Supported country codes:** `GT`, `US`, `MX`, `JP`, `ES`, `DE`, `FR`, `BR`, `AR`, `CO` and any ISO 3166-1 alpha-2 code.

## Internal Flow
```
REST:  city → Geocoding (Open-Meteo) → lat/lon → current weather
SOAP:  country code → capital city (CDYNE SOAP) → lat/lon → current weather (Open-Meteo)
```

## Run Locally
```bash
git clone https://github.com/Xavez05/weather-monorepo.git
cd weather-monorepo/weather-app
go run .
```

Server available at `http://localhost:8080`.

## Run with Docker
```bash
docker build -t weather-app .
docker run -p 8080:8080 weather-app
```

## Tests
```bash
cd weather-app
go test ./... -v
```

Coverage includes:
- Successful REST and SOAP responses
- Service errors (city not found, invalid country)
- Empty request body validation
- Correct HTTP status codes (200, 400, 502)

## Principles Applied

- **SOLID** — `WeatherFetcher` interface for dependency inversion, single responsibility per file
- **Clean Code** — handlers, services and routes separated into independent files
- **Go Workspaces** — `apiclient` is reusable as an independent library across monorepo projects
- **Multistage Docker** — lightweight final image based on `alpine`
- **GitHub Actions** — CI pipeline with automated build, tests and docker build
