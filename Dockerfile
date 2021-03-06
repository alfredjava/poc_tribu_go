# Generar imagen de API Rest para incidencias en dos pasos
FROM golang:latest AS builder

# Disponibilizar dependencias
RUN go get github.com/gorilla/mux

# Compilar API y generar ejecutable
WORKDIR /app
COPY ./*.go ./
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -a -o ./main .

FROM alpine:latest
RUN apk --no-cache add ca-certificates

# Copiar artefacto generado por el builder en el paso anterior
COPY --from=builder /app/main ./
RUN chmod +x ./main

# Configurar el puerto de escucha para la aplciación.
EXPOSE 8083

# Ejecutar API
CMD ["./main"]