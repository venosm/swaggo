# Example App - OpenAPI 3.0 Demo

Tato ukázková aplikace demonstruje použití upraveného Swag nástroje pro generování OpenAPI 3.0 dokumentace místo původní Swagger 2.0.

## Vlastnosti

- **REST API** pro správu uživatelů
- **OpenAPI 3.0** dokumentace generovaná pomocí upraveného Swag nástroje
- **České pojmenování** struktur a endpointů
- **Gin framework** pro webový server
- **Swagger UI** pro interaktivní prohlížení API

## Struktura API

### Endpointy

- `GET /api/v1/uzivatele` - Získej všechny uživatele
- `GET /api/v1/uzivatele/{id}` - Získej konkrétního uživatele podle ID
- `POST /api/v1/uzivatele` - Vytvoř nového uživatele

### Dokumentace

- `GET /swagger/index.html` - Swagger UI interface
- `GET /swagger/doc.json` - OpenAPI 3.0 JSON dokumentace

## Spuštění

```bash
# Nainstaluj závislosti
go mod tidy

# Spusť aplikaci
go run .
```

Aplikace bude dostupná na `http://localhost:8080`

## Testování API

```bash
# Získej všechny uživatele
curl http://localhost:8080/api/v1/uzivatele

# Získej konkrétního uživatele
curl http://localhost:8080/api/v1/uzivatele/1

# Vytvoř nového uživatele
curl -X POST http://localhost:8080/api/v1/uzivatele \
  -H "Content-Type: application/json" \
  -d '{"jmeno": "Nový Uživatel", "email": "novy@example.com", "vek": 25}'
```

## Swagger UI

Interaktivní dokumentaci najdeš na: http://localhost:8080/swagger/index.html

## Generování dokumentace

Dokumentace se generuje pomocí našeho upraveného swag nástroje:

```bash
# Ze složky swag-master
./swag init -g exampleapp/main.go --output exampleapp/docs
```

## Hlavní rozdíly oproti Swagger 2.0

- **OpenAPI verze**: "3.0.0" místo "2.0"
- **Aktualizovaná dokumentace**: všechny reference na Swagger 2.0 změněny na OpenAPI 3.0
- **Zachovaná kompatibilita**: všechna existující funkcionalita zůstává nezměněna