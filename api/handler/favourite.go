package handler

import (
	"net/http"
	"github.com/gin-gonic/gin"
	favorites "api_gateway/genproto/favorites"
)

// CreateFavorites godoc
// @Summary      Create a new favorite
// @Description  Add a new property to user's favorites
// @Tags         favorites
// @Accept       json
// @Produce      json
// @Param        body body favorites.CreateFavoritesReq true "Create Favorite Request"
// @Success      200 {object} favorites.CreateFavoritesRes
// @Failure      400 {object} string
// @Router       /favorites/create [post]
func (h *Handler) CreateFavorites(c *gin.Context) {
	var req favorites.CreateFavoritesReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := h.FavouriteService.CreateFavorites(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// DeleteFavorites godoc
// @Summary      Delete a favorite
// @Description  Remove a property from user's favorites
// @Tags         favorites
// @Accept       json
// @Produce      json
// @Param        id path string true "Favorite ID"
// @Success      200 {object} favorites.DeleteFavoritesRes
// @Failure      400 {object} string
// @Router       /favorites/delete/{id} [delete]
func (h *Handler) DeleteFavorites(c *gin.Context) {
	var req favorites.DeleteFavoritesReq
	req.Id = c.Param("id")
	if req.Id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	res, err := h.FavouriteService.DeleteFavorites(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// GetAllFavorites godoc
// @Summary      Get all favorites
// @Description  Retrieve a list of all user's favorite properties
// @Tags         favorites
// @Accept       json
// @Produce      json
// @Param        limit query int true "Limit"
// @Param        page query int true "Page"
// @Success      200 {object} favorites.GetAllFavoritesRes
// @Failure      400 {object} string
// @Router       /favorites/getall [get]
func (h *Handler) GetAllFavorites(c *gin.Context) {
	var req favorites.GetAllFavoritesReq
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := h.FavouriteService.GetAllFavorites(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// GetByIdFavorites godoc
// @Summary      Get a favorite by ID
// @Description  Retrieve a specific favorite by its ID
// @Tags         favorites
// @Accept       json
// @Produce      json
// @Param        id path string true "Favorite ID"
// @Success      200 {object} favorites.GetByIdFavoritesRes
// @Failure      400 {object} string
// @Router       /favorites/getbyid/{id} [get]
func (h *Handler) GetByIdFavorites(c *gin.Context) {
	var req favorites.GetByIdFavoritesReq
	req.Id = c.Param("id")
	if req.Id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	res, err := h.FavouriteService.GetByIdFavorites(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}
