package auth

import (
	"net/http"

	"github.com/ardianilyas/go-auth-domain/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Authhandler struct {
	service AuthService
}

func NewAuthHandler(service AuthService) *Authhandler {
	return &Authhandler{service}
}

func (h *Authhandler) Register(c *gin.Context) {
	var req RegisterRequest
	if !utils.BindAndValidate(c, &req) {
		return
	}

	user, err := h.service.Register(req.Username, req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp := RegisterResponse{
		ID:        	user.ID.String(),
		Username:  	user.Username,
		Email:     	user.Email,
		Role: 		user.Role,
	}

	utils.RespondSuccess(c, http.StatusCreated, "user registered successfully", resp)
}

func (h *Authhandler) Login(c *gin.Context) {
	var req LoginRequest
	if !utils.BindAndValidate(c, &req) {
		return
	}

	user, accessToken, refreshToken, err := h.service.Login(req.Email, req.Password)
	if err != nil {
		utils.Unauthorized(c, err.Error())
		return
	}

	c.SetCookie("access_token", accessToken, 3600, "/", "", false, true)
	c.SetCookie("refresh_token", refreshToken, 3600, "/", "", false, true)

	resp := LoginResponse{
		ID: user.ID.String(),
		Username: user.Username,
		Email: user.Email,
		Role: user.Role,
	}

	utils.RespondSuccess(c, http.StatusOK, "user logged in successfully", resp)
}

func (h *Authhandler) Refresh(c *gin.Context) {
	refreshToken, err := c.Cookie("refresh_token")
	if err != nil {
		utils.Unauthorized(c, "refresh token missing")
		return
	}

	newAccess, newRefresh, err := h.service.RefreshToken(refreshToken)
	if err != nil {
		utils.Unauthorized(c, err.Error())
		return
	}

	c.SetCookie("access_token", newAccess, 3600, "/", "", false, true)
	c.SetCookie("refresh_token", newRefresh, 3600, "/", "", false, true)

	utils.RespondSuccess(c, http.StatusOK, "token refreshed successfully", nil)
}

func (h *Authhandler) Logout(c *gin.Context) {
	refreshToken, err := c.Cookie("refresh_token")
	if err != nil {
		utils.Unauthorized(c, "refresh token missing")
		return
	}

	if err := h.service.Logout(refreshToken); err != nil {
		utils.InternalError(c, err.Error())
		return
	}

	c.SetCookie("access_token", "", -1, "/", "", false, true)
	c.SetCookie("refresh_token", "", -1, "/", "", false, true)

	utils.RespondSuccess(c, http.StatusOK, "user logged out successfully", nil)
}

func (h *Authhandler) Me(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		utils.Unauthorized(c, "unauthorized")
		return
	}

	uid, err := uuid.Parse(userID.(string))
	if err != nil {
		utils.Unauthorized(c, "invalid user id")
		return
	}

	user, err := h.service.GetUserByID(uid)
	if err != nil {
		utils.BadRequest(c, "user not found")
		return
	}

	resp := MeResponse{
		ID: user.ID.String(),
		Email: user.Email,
		Role: user.Role,
	}

	utils.RespondSuccess(c, http.StatusOK, "user details", resp)
}