# Sistema de Llamado de Mozos — Arquitectura DDD (Go + Next.js)

## 🧠 Contexto del dominio

Sistema para restaurantes donde:

- Cada mesa tiene un QR
- El cliente accede a una web app
- Puede llamar al mozo, pedir la cuenta o hacer consultas
- El restaurante recibe solicitudes en tiempo real
- Se registran métricas y feedback

---

## 🧱 Enfoque arquitectónico

Estilo: **DDD (Domain-Driven Design)** + **Clean Architecture**

```
/domain        → reglas de negocio puras
/application   → casos de uso
/infrastructure→ DB, WebSockets, APIs
/interfaces    → HTTP handlers / controllers
```

---

## 🧩 Dominio

### Restaurant

```go
type Restaurant struct {
    ID   string
    Name string
    Plan string
}
```

### Table

```go
type Table struct {
    ID           string
    Number       int
    RestaurantID string
    QRCode       string
}
```

### Request

```go
type RequestType string

const (
    CallWaiter RequestType = "CALL_WAITER"
    AskBill    RequestType = "ASK_BILL"
    AskHelp    RequestType = "ASK_HELP"
)

type RequestStatus string

const (
    Pending   RequestStatus = "PENDING"
    InProcess RequestStatus = "IN_PROCESS"
    Done      RequestStatus = "DONE"
)

type Request struct {
    ID        string
    TableID   string
    Type      RequestType
    Status    RequestStatus
    CreatedAt time.Time
}
```

### Feedback

```go
type Feedback struct {
    ID        string
    TableID   string
    Score     int
    CreatedAt time.Time
}
```

---

## ⚙️ Application Layer

### Casos de uso

- Crear solicitud
- Marcar como atendida
- Obtener solicitudes activas
- Registrar feedback

---

## 🏗 Infrastructure

- Base de datos: MySQL / PostgreSQL
- WebSockets (recomendado)
- Generación de QR dinámicos

---

## 🌐 Interfaces

### Cliente

- POST /requests
- GET /table/{id}/status
- POST /feedback

### Restaurante

- GET /requests/active
- PATCH /requests/{id}/complete

---

## 🖥 Frontend (Next.js)

### Cliente

- Botonera simple
- Estado en tiempo real

### Dashboard

- Lista de mesas
- Acciones rápidas

---

## 🔄 Flujo

Cliente → Backend → Evento → Restaurante → Atención

---

## 🚀 Stack

Backend:

- Go (Gin / Fiber)
- WebSockets

Frontend:

- Next.js
- Tailwind

Infra:

- Docker
- Nginx

---

## 🧠 Resumen

Sistema basado en:

- Eventos simples
- Estado controlado
- Tiempo real
- Dominio claro
