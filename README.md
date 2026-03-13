# Backend

## Prerequisites
- Go 1.22 or higher

## Running the Server

```bash
cd backend
go run main.go
```

The server will start on `http://localhost:8080`

## Testing

```bash
# Health check
curl http://localhost:8080/health

# Products endpoint (to be implemented)
curl http://localhost:8080/products
```

## Data Files

- `data/metadata.json` - Product metadata (id, name, base_price, image_url)
- `data/details.json` - Product details (id, discount_percent, bestseller, colors, stock)

