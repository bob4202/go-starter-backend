package user

import (
	"go-starter-backend/pkg/appcommon"
	"go-starter-backend/pkg/apperror"
	"go-starter-backend/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
}

type Messages string

const (
	UserRegisterSuccess   Messages = "user register successfully"
	UserDeletedSuccess    Messages = "user deleted successfully"
	UserNotFound          Messages = "user not found"
	UserUpdateSuccess     Messages = "user update successfully"
	PasswordUpdateSuccess Messages = "password update successfully"
)

func NewHandler(service *Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Register(c *gin.Context) {
	var req RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, apperror.ErrBadRequest.Error())
	}

	user, err := h.service.Register(req.Name, req.Email, req.Password)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
	}

	response.Created(c, string(UserRegisterSuccess), user)
}

func (h *Handler) DeleteUser(c *gin.Context) {
	var req DeleteUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, apperror.ErrBadRequest.Error())
	}

	err := h.service.DeleteUser(req.ID)

	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
	}

	response.Success(c, http.StatusOK, string(UserDeletedSuccess), nil)
}

func (h *Handler) ChangePassword(c *gin.Context) {
	var req ChangePasswordRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, apperror.ErrBadRequest.Error())
	}

	userID := c.GetString(appcommon.UserID)

	user, err := h.service.Me(userID)

	if err != nil {
		response.Error(c, http.StatusNotFound, string(UserNotFound))
	}
	err = h.service.ChangePassword(user.ID, req.Password)

	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
	}
	response.Success(c, http.StatusOK, string(PasswordUpdateSuccess), nil)

}

func (h *Handler) UpdateUser(c *gin.Context) {
	var req UpdateRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, apperror.ErrBadRequest.Error())
	}

	user, err := h.service.UpdateUser(req.ID, req.Name, req.Email)

	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
	}

	response.Success(c, http.StatusOK, string(UserUpdateSuccess), user)

}

func (h *Handler) GetUserById(c *gin.Context) {

	var req GetUserByIDRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, apperror.ErrBadRequest.Error())
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

	response.Success(c, http.StatusOK, string(appcommon.Success), res)
}

func (h *Handler) GetUserByEmail(c *gin.Context) {

	var req GetUserByEmailRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, apperror.ErrBadRequest.Error())
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

	response.Success(c, http.StatusOK, string(appcommon.Success), res)
}
func (h *Handler) Me(c *gin.Context) {
	userID := c.GetString(appcommon.UserID)

	user, err := h.service.Me(userID)

	if err != nil {
		response.Error(c, http.StatusUnauthorized, string(UserNotFound))
	}
	res := PrivateUserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}

	response.Success(c, http.StatusOK, string(appcommon.Success), res)

}
func (h *Handler) Login(c *gin.Context) {

	var req LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, apperror.ErrBadRequest.Error())
	}

	token, err := h.service.Login(req.Email, req.Password)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
	}

	res := LoginResponse{
		Token: token,
	}

	response.Success(c, http.StatusOK, string(appcommon.Success), res)
}
