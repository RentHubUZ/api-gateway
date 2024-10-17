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
// @Router /api/user/profile [get]
func (h *Handler) GetProfile(c *gin.Context) {
	h.Log.Info("GetProfile handler is invoked")

	id, err := getUserID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.UserService.GetProfile(c.Request.Context(), &pb.ID{Id: id})
	if err != nil {
		h.Log.Error("GetProfile error", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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
// @Router /api/user/profile/update [put]
func (h *Handler) UpdateProfile(c *gin.Context) {
	h.Log.Info("UpdateProfile handler is invoked")

	id, err := getUserID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	var req models.UserUpdate
	if err := c.ShouldBind(&req); err != nil {
		h.Log.Error("UpdateProfile binding error", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err = h.UserService.UpdateProfile(c.Request.Context(), &pb.NewData{
		Id:          id,
		FullName:    req.FullName,
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
	})
	if err != nil {
		h.Log.Error("UpdateProfile error", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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
// @Router /api/user/profile/delete [delete]
func (h *Handler) DeleteProfile(c *gin.Context) {
	h.Log.Info("DeleteProfile handler is invoked")

	id, err := getUserID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	_, err = h.UserService.DeleteProfile(c.Request.Context(), &pb.ID{Id: id})
	if err != nil {
		h.Log.Error("DeleteProfile error", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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
// @Router /api/user/password [put]
func (h *Handler) ChangePassword(c *gin.Context) {
	h.Log.Info("ChangePassword handler is invoked")

	id, err := getUserID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	var req models.ChangePassword
	if err := c.ShouldBind(&req); err != nil {
		h.Log.Error("ChangePassword binding error", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err = h.UserService.ChangePassword(c.Request.Context(), &pb.NewPass{
		Id:          id,
		OldPassword: req.OldPassword,
		NewPassword: req.NewPassword,
	})
	if err != nil {
		h.Log.Error("ChangePassword error", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	h.Log.Info("ChangePassword handler completed")
	c.JSON(http.StatusOK, "Password changed successfully")
}
