package http

import (
	"net/http"
	"strconv"

	"github.com/gene-qxsi/Blog-user/internal/domain"
	"github.com/gene-qxsi/Blog-user/internal/presentation/dto"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	srv domain.UserService
}

func NewUserHandler(userService domain.UserService) *UserHandler {
	return &UserHandler{srv: userService}
}

// TODO: заменить работу с ошибками
func (h *UserHandler) CreateUser(c *gin.Context) {
	const op = "user-service>internal>presentation>handlers>http>user_create.go>CreateUser()"

	var userCreateReuqest dto.UserRequest
	if err := c.ShouldBindJSON(&userCreateReuqest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid json: %s" + err.Error() + ". op: " + op})
		return
	}

	id, err := h.srv.CreateUser(c.Request.Context(), userCreateReuqest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "validation error: %s" + err.Error() + ". op: " + op})
		return
	}

	c.JSON(http.StatusCreated, dto.UserCreateResponse{
		ID: id,
	})
}

func (h *UserHandler) GetUser(c *gin.Context) {
	const op = "user-service>internal>presentation>handlers>http>user_create.go>GetUser()"

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid param: %s" + err.Error() + ". op: " + op})
		return
	}

	user, err := h.srv.GetUser(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid response from GetUser. op: " + op})
		return
	}

	c.JSON(http.StatusOK, dto.UserResponse{
		ID:       user.ID(),
		Email:    user.Email(),
		Password: user.Password(),
	})
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	const op = "user-service>internal>presentation>handlers>http>user_create.go>DeleteUser()"

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid param: %s" + err.Error() + ". op: " + op})
		return
	}

	if err := h.srv.DeleteUser(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error deleted user: %s" + err.Error() + ". op: " + op})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	const op = "user-service>internal>presentation>handlers>http>user_create.go>UpdateUser()"

	id := c.GetHeader("X-User-ID")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is empty. op: " + op})
		return
	}

	userID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id header: %s" + err.Error() + ". op: " + op})
		return
	}

	var updateUserRequest dto.UserRequest
	if err := c.ShouldBindJSON(&updateUserRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid json: %s" + err.Error() + ". op: " + op})
		return
	}

	user, err := h.srv.UpdateUser(c.Request.Context(), userID, updateUserRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error updated user: %s" + err.Error() + ". op: " + op})
		return
	}

	c.JSON(http.StatusOK, dto.UserResponse{
		ID:       user.ID(),
		Email:    user.Email(),
		Password: user.Password(),
	})
}

func (h *UserHandler) RegisterUserRoutes(r *gin.RouterGroup) {
	userGroup := r.Group("/users")
	{
		userGroup.GET("/:id", h.GetUser)
		userGroup.POST("/", h.CreateUser)
		userGroup.DELETE("/:id", h.DeleteUser)
		userGroup.PUT("/", h.UpdateUser)
	}
}
