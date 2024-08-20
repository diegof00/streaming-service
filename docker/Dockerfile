# Usar una imagen base ligera de Go 1.23
FROM golang:1.23-alpine

# Crear directorio de trabajo dentro del contenedor
WORKDIR /app

# Copiar go.mod y go.sum desde el contexto (.. en docker-compose.yml)
COPY go.mod go.sum ./

# Descargar dependencias del proyecto
RUN go mod download

# Copiar el resto del código desde el contexto
COPY . .

# Construir la aplicación
RUN go build -o main ./cmd/server/main.go

# Verificar que el binario `main` se haya generado correctamente
RUN ls -la /app

# Exponer el puerto 8080
EXPOSE 8080

# Ejecutar la aplicación
CMD ["./main"]
