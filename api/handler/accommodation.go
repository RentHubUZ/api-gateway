package handler

import (
	"fmt"
	"net/http"
	"strconv"

	pbaccom "api_gateway/genproto/accommodation"
	"api_gateway/internal/models"

	"github.com/gin-gonic/gin"
)

// CreateHouse godoc
// @Summary      Create a new house
// @Description  Create a new house with specified properties
// @Tags         houses
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        body body models.CreateProperties true "Create House Request"
// @Success      202 {object} accommodation.CreateHouseRes
// @Failure      400 {object} string
// @Failure      500 {object} string
// @Router       /api/properties/propertiescreate [post]
func (h *Handler) CreateHouse(c *gin.Context) {
	userId, exists := c.Get("user_id")
	if !exists {
		h.Log.Error("User ID not found in context")
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID not found in context"})
		return
	}
	id := userId.(string)

	var properties models.CreateProperties
	if err := c.ShouldBindJSON(&properties); err != nil {
		h.Log.Error("Error binding JSON: ", "error", err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	req := pbaccom.CreateHouseReq{
		OwnerId:       id,
		Address:       properties.Address,
		Price:         float32(properties.Price),
		PropertyType:  properties.Property_type,
		Bedrooms:      properties.Bedrooms,
		Bathrooms:     properties.Bathrooms,
		SquareFootage: float32(properties.Square_footage),
		ListingStatus: properties.Listing_status,
		Description:   properties.Description,
		RoommateCount: properties.Roommate_count,
		LeaseTerms:    properties.Lease_terms,
		LeaseDuration: properties.Lease_duration,
		TopStatus:     properties.Top_status,
		ImageUrl:      properties.Image_url,
		Latitude:      float32(properties.Latitude),
		Longitude:     float32(properties.Longitude),
	}

	resp, err := h.AccommodationService.CreateHouse(c, &req)
	if err != nil {
		h.Log.Error(fmt.Sprintf("could not find the information about the house %v", err))
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, resp)
}

// UpdateHouse godoc
// @Summary      Update an existing house
// @Description  Update an existing house with new details
// @Tags         houses
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        body body models.UpdateProperties true "Update House Request"
// @Success      202 {object} accommodation.UpdateHouseRes
// @Failure      400 {object} string
// @Failure      500 {object} string
// @Router       /api/properties/propertiesupdate [put]
func (h *Handler) UpdateHouse(c *gin.Context) {
	userId, exists := c.Get("user_id")
	if !exists {
		h.Log.Error("User ID not found in context")
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID not found in context"})
		return
	}
	id := userId.(string)

	var updateproperties models.UpdateProperties
	if err := c.ShouldBindJSON(&updateproperties); err != nil {
		h.Log.Error("Error binding JSON: ", "error", err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	req := pbaccom.UpdateHouseReq{
		Id:            updateproperties.Id,
		OwnerId:       id,
		Address:       updateproperties.Address,
		Price:         float32(updateproperties.Price),
		PropertyType:  updateproperties.Property_type,
		Bedrooms:      updateproperties.Bedrooms,
		Bathrooms:     updateproperties.Bathrooms,
		SquareFootage: float32(updateproperties.Square_footage),
		ListingStatus: updateproperties.Listing_status,
		Description:   updateproperties.Description,
		RoommateCount: updateproperties.Roommate_count,
		LeaseTerms:    updateproperties.Lease_terms,
		LeaseDuration: updateproperties.Lease_duration,
		TopStatus:     updateproperties.Top_status,
		ImageUrl:      updateproperties.Image_url,
		Latitude:      float32(updateproperties.Latitude),
		Longitude:     float32(updateproperties.Longitude),
	}

	resp, err := h.AccommodationService.UpdateHouse(c, &req)
	if err != nil {
		h.Log.Error(fmt.Sprintf("error updating home data %v", err))
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, resp)
}

// GetAllHouse godoc
// @Summary      Get all houses
// @Description  Retrieve a list of all houses with pagination
// @Tags         houses
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        limit path int true "Limit"
// @Param        page  path int true "Page"
// @Success      202 {object} accommodation.GetAllHouseRes
// @Failure      400 {object} string
// @Failure      500 {object} string
// @Router       /api/properties/propertiesgetall/{limit}/{page} [get]
func (h *Handler) GetAllHouse(c *gin.Context) {
	limitStr := c.Param("limit")
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit parameter"})
		return
	}

	pageStr := c.Param("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page parameter"})
		return
	}

	limitInt32 := int32(limit)
	pageInt32 := int32(page)

	req := pbaccom.GetallHouseReq{
		Limit: limitInt32,
		Page:  pageInt32,
	}

	resp, err := h.AccommodationService.GetAllHouse(c, &req)
	if err != nil {
		h.Log.Error(fmt.Sprintf("error retrieving all house information %v", err))
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, resp)
}

// GetByIdHouse godoc
// @Summary      Get house by ID
// @Description  Retrieve house information by ID
// @Tags         houses
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        properties_id path string true "Property ID"
// @Success      202 {object} accommodation.GetByIdHouseRes
// @Failure      500 {object} string
// @Router       /api/properties/propertiesgetbyid/{properties_id} [get]
func (h *Handler) GetByIdHouse(c *gin.Context) {
	id := c.Param("properties_id")

	req := pbaccom.GetByIdHouseReq{
		Id: id,
	}

	resp, err := h.AccommodationService.GetByIdHouse(c, &req)
	if err != nil {
		h.Log.Error(fmt.Sprintf("error in getting information about house id %v", err))
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, resp)
}

// DeleteHouse godoc
// @Summary      Delete a house
// @Description  Delete a house by ID
// @Tags         houses
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        properties_id path string true "Property ID"
// @Success      202 {object} accommodation.DeleteHouseRes
// @Failure      500 {object} string
// @Router       /api/properties/propertiesdelete/{properties_id} [delete]
func (h *Handler) DeleteHouse(c *gin.Context) {
	id := c.Param("properties_id")

	req := pbaccom.DeleteHouseReq{
		Id: id,
	}

	resp, err := h.AccommodationService.DeleteHouse(c, &req)
	if err != nil {
		h.Log.Error(fmt.Sprintf("error when deleting house id information %v", err))
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, resp)
}
