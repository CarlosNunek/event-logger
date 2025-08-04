# -------- Etapa 1: Build --------
FROM golang:1.21 AS builder

# Crear directorio de trabajo dentro del contenedor
WORKDIR /app

# Copiar archivos de dependencias y descargar módulos
COPY go.mod go.sum ./
RUN go mod download

# Copiar el resto del código
COPY . .

# Compilar la aplicación
RUN go build -o eventlogger main.go

# -------- Etapa 2: Imagen final --------
FROM debian:bookworm-slim

WORKDIR /app

# Copiar solo el ejecutable
COPY --from=builder /app/eventlogger .

# Copiar archivo .env si lo necesitas dentro del contenedor (opcional)
COPY .env .

# (Opcional) Exponer puerto si luego agregas servidor HTTP
# EXPOSE 8080

# Comando de inicio
CMD ["./eventlogger"]

