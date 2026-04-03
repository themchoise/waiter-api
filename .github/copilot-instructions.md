## Swagger / swaggo

Este proyecto usa **swaggo/swag** para documentar la API automáticamente.

### Regla obligatoria

Cada vez que se cree o modifique un handler HTTP en `interfaces/http/`, **SIEMPRE** agregar o actualizar las annotations de Swagger encima de la función. Formato:

```go
// NombreFuncion godoc
// @Summary      Descripción corta
// @Description  Descripción larga (opcional)
// @Tags         nombre-del-recurso
// @Accept       json
// @Produce      json
// @Param        nombre  path/body/query  tipo  requerido  "descripción"
// @Success      200  {object}  TipoRespuesta
// @Failure      400  {object}  map[string]string
// @Failure      422  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /api/v1/ruta [método]
```

### Reglas de annotations

- `@Tags` debe coincidir con el recurso: `requests`, `feedback`, `restaurants`, `tables`
- `@Param` debe incluir todos los parámetros: path params, query params y body
- Para body usar: `@Param request body usecase.NombreInput true "descripción"`
- Para path params: `@Param id path string true "ID del recurso"`
- `@Success` y `@Failure` deben reflejar los códigos HTTP reales del handler
- `@Router` debe coincidir exactamente con la ruta definida en `router.go`

### Ejemplo completo

```go
// Create godoc
// @Summary      Crear solicitud
// @Description  Crea una nueva solicitud de mesa (llamar mozo, pedir cuenta, ayuda)
// @Tags         requests
// @Accept       json
// @Produce      json
// @Param        request  body      usecase.CreateRequestInput  true  "Datos de la solicitud"
// @Success      201      {object}  entity.Request
// @Failure      400      {object}  map[string]string
// @Failure      422      {object}  map[string]string
// @Router       /api/v1/requests [post]
func (h *RequestHandler) Create(c *gin.Context) {
```

### Generación

Después de modificar annotations, ejecutar:

```bash
swag init -g cmd/server/main.go -o docs
```

### Anotación del main

El archivo `cmd/server/main.go` debe tener las annotations globales de la API:

```go
// @title          Waiter API
// @version        1.0
// @description    Sistema de llamado de mozos para restaurantes
// @host           localhost:8080
// @BasePath       /api/v1
```
