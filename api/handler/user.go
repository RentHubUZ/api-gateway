package handler

import (
	pb "api_gateway/genproto/user"
	"api_gateway/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetProfile godoc
// @Summary Gets profile
// @Description Retrieves user profile
// @Tags user
// @Security ApiKeyAuth
// @Success 200 {object} user.Profile
// @Failure 401 {object} string
// @Failure 500 {object} string
// @Router       /user/profile [get]
func (h *Handler) GetProfile(c *gin.Context) {
	h.Log.Info("GetProfile handler is invoked")

	ctx := c.Request.Context() // Request context'ni alohida o'zgaruvchiga olib chiqish
	id, err := getUserID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error1": err.Error()})
		return
	}

	// gRPC so'rovi uchun alohida o'zgaruvchi
	profileRequest := &pb.ID{Id: id}

	resp, err := h.UserService.GetProfile(ctx, profileRequest)
	if err != nil {
		h.Log.Error("GetProfile error", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error2": err.Error()})
		return
	}

	h.Log.Info("GetProfile handler completed", "response", resp)
	c.JSON(http.StatusOK, resp)
}

// UpdateProfile godoc
// @Summary Updates profile
// @Description Updates user profile
// @Tags user
// @Security ApiKeyAuth
// @Param data body models.UserUpdate true "New user data"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Failure 401 {object} string
// @Failure 500 {object} string
// @Router       /user/profile/update [put]
func (h *Handler) UpdateProfile(c *gin.Context) {
	h.Log.Info("UpdateProfile handler is invoked")

	ctx := c.Request.Context() // Request context
	id, err := getUserID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error1": err.Error()})
		return
	}

	// So'rov body'sini alohida o'zgaruvchiga olish
	var req models.UserUpdate
	if err := c.ShouldBind(&req); err != nil {
		h.Log.Error("UpdateProfile binding error", "error2", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error2": err.Error()})
		return
	}

	// gRPC uchun yangi profil ma'lumotlari
	updateRequest := &pb.NewData{
		Id:          id,
		FullName:    req.FullName,
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
	}

	_, err = h.UserService.UpdateProfile(ctx, updateRequest)
	if err != nil {
		h.Log.Error("UpdateProfile error", "error3", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error3": err.Error()})
		return
	}

	h.Log.Info("UpdateProfile handler completed")
	c.JSON(http.StatusOK, "User updated successfully")
}

// DeleteProfile godoc
// @Summary Deletes profile
// @Description Deletes user profile
// @Tags user
// @Security ApiKeyAuth
// @Success 200 {object} string
// @Failure 401 {object} string
// @Failure 500 {object} string
// @Router       /user/profile/delete [delete]
func (h *Handler) DeleteProfile(c *gin.Context) {
	h.Log.Info("DeleteProfile handler is invoked")

	ctx := c.Request.Context() // Request context
	id, err := getUserID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error1": err.Error()})
		return
	}

	// gRPC so'rovi uchun ID o'zgaruvchisi
	deleteRequest := &pb.ID{Id: id}

	_, err = h.UserService.DeleteProfile(ctx, deleteRequest)
	if err != nil {
		h.Log.Error("DeleteProfile error", "error2", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error2": err.Error()})
		return
	}

	h.Log.Info("DeleteProfile handler completed")
	c.JSON(http.StatusOK, "User deleted successfully")
}

// ChangePassword godoc
// @Summary Changes password
// @Description Changes user password
// @Tags user
// @Security ApiKeyAuth
// @Param data body models.ChangePassword true "Passwords"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Failure 401 {object} string
// @Failure 500 {object} string
// @Router       /user/password [put]
func (h *Handler) ChangePassword(c *gin.Context) {
	h.Log.Info("ChangePassword handler is invoked")

	ctx := c.Request.Context() // Request context
	id, err := getUserID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error1": err.Error()})
		return
	}

	// Parollar so'rovi
	var req models.ChangePassword
	if err := c.ShouldBind(&req); err != nil {
		h.Log.Error("ChangePassword binding error", "error2", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error2": err.Error()})
		return
	}

	// gRPC so'rovi uchun yangi parol ma'lumotlari
	changePasswordRequest := &pb.NewPass{
		Id:          id,
		OldPassword: req.OldPassword,
		NewPassword: req.NewPassword,
	}

	_, err = h.UserService.ChangePassword(ctx, changePasswordRequest)
	if err != nil {
		h.Log.Error("ChangePassword error", "error3", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	h.Log.Info("ChangePassword handler completed")
	c.JSON(http.StatusOK, "Password changed successfully")
}
