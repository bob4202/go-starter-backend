package user

import (
	"go-starter-backend/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{
		service: service,
	}
}

type PublicUserResponse struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
type RegisterRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required, min=6"`
}

type UpdateRequest struct {
	ID    string `json:"id" binding:"required"`
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
}

type ChangePasswordRequest struct {
	ID       string `json:"id" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type DeleteUserRequest struct {
	ID string `json:"id" binding:"required"`
}

type GetUserByIDRequest struct {
	ID string `json:"id" binding:"required"`
}

type GetUserByEmailRequest struct {
	Email string `json:"email" binding:"required"`
}

func (h *Handler) Register(c *gin.Context) {
	var req RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "invalid request body")
	}

	user, err := h.service.Register(req.Name, req.Email, req.Password)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
	}

	response.Created(c, "user register successfully", user)
}

func (h *Handler) DeleteUser(c *gin.Context) {
	var req DeleteUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "invalid request body")
	}

	err := h.service.DeleteUser(req.ID)

	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
	}

	response.Success(c, http.StatusAccepted, "user deleted successfully", nil)
}

func (h *Handler) ChangePassword(c *gin.Context) {
	var req ChangePasswordRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "invalid request body")
	}

	err := h.service.ChangePassword(req.ID, req.Password)

	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
	}
	response.Success(c, http.StatusAccepted, "update password successfully", nil)

}

func (h *Handler) UpdateUser(c *gin.Context) {
	var req UpdateRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "invalid request body")
	}

	user, err := h.service.UpdateUser(req.ID, req.Name, req.Email)

	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
	}

	response.Success(c, http.StatusAccepted, "user update successfully", user)

}

func (h *Handler) GetUserById(c *gin.Context) {

	var req GetUserByIDRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "invalid request body")
	}

	user, err := h.service.GetUserByID(req.ID)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
	}

	res := PublicUserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}

	response.Success(c, http.StatusAccepted, "success", res)
}

func (h *Handler) GetUserByEmail(c *gin.Context) {

	var req GetUserByEmailRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "invalid request body")
	}

	user, err := h.service.GetUserByID(req.Email)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
	}

	res := PublicUserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}

	response.Success(c, http.StatusAccepted, "success", res)
}
