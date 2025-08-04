# 📦 event-logger

Microservicio desarrollado en Go que escucha eventos desde Redis (por canales como `usuarios` y `login`) y almacena los datos en InfluxDB como registros de serie temporal.

---

## 📌 Arquitectura utilizada

- **Event-Driven Architecture (EDA)** basada en eventos.
- **Monolito distribuido simple**.
- Comunicación asíncrona a través de Redis Pub/Sub.
- Persistencia en **InfluxDB** (serie temporal).

---

## 🎯 Patrón de diseño aplicado

- **Publisher/Subscriber**: Redis es el intermediario de eventos.
- **Separation of concerns**: cada paquete (`config/`, `redis/`, `influx/`) tiene una responsabilidad clara

---

### 🛠 Instalación local
go run main.go
