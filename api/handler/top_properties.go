package handler

import (
	pbtop "api_gateway/genproto/top_properties"
	"api_gateway/internal/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateTopProperties godoc
// @Summary      Create top properties
// @Description  Create a new top property with specified details
// @Tags         top_properties
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        body body models.Top_Properties true "Top Properties"
// @Success      200 {object} top_properties.CreateTopPropertyRes
// @Failure      400 {object} string
// @Router       /topproperties/createtopproperties [post]
func (h *Handler) CreateTopProperties(c *gin.Context) {
	userId, exists := c.Get("user_id")
	if !exists {
		h.Log.Error("User ID not found in context")
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID not found in context"})
		return
	}
	id := userId.(string)

	var topproperties models.Top_Properties
	if err := c.ShouldBindJSON(&topproperties); err != nil {
		h.Log.Error("Error binding JSON: ", "error", err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	req := pbtop.CreateTopPropertyReq{
		UserId:     id,
		PropertyId: topproperties.Property_id,
		TariffName: topproperties.Tariff_name,
	}

	resp, err := h.TopPropertiesService.Create(c, &req)
	if err != nil {
		h.Log.Error(fmt.Sprintf("Create topperoperties error: %v", err.Error()))
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, resp)
}

// UpdateTopProperties godoc
// @Summary      Update top properties
// @Description  Update an existing top property by its ID
// @Tags         top_properties
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        body body models.Top_UpdateProperties true "Update Top Properties"
// @Success      200 {object} top_properties.UpdateTopPropertyRes
// @Failure      400 {object} string
// @Router       /topproperties/updatetopproperties [put]
func (h *Handler) UpdateTopProperties(c *gin.Context) {
	userId, exists := c.Get("user_id")
	if !exists {
		h.Log.Error("User ID not found in context")
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID not found in context"})
		return
	}
	id := userId.(string)

	var updatetopproperties models.Top_UpdateProperties
	if err := c.ShouldBindJSON(&updatetopproperties); err != nil {
		h.Log.Error("Error binding JSON: ", "error", err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	req := pbtop.UpdateTopPropertyReq{
		Id:         updatetopproperties.Id,
		PropertyId: updatetopproperties.Property_id,
		UserId:     id,
		StartDate:  updatetopproperties.Start_date,
		EndDate:    updatetopproperties.End_date,
		TariffName: updatetopproperties.Tariff_name,
	}

	resp, err := h.TopPropertiesService.Update(c, &req)
	if err != nil {
		h.Log.Error(fmt.Sprintf("Update top properties error: %v", err.Error()))
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, resp)
}

// GetByIdTopProperties godoc
// @Summary      Get top properties by ID
// @Description  Retrieve top properties details by its ID
// @Tags         top_properties
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        top_properties_id path string true "Top Properties ID"
// @Success      200 {object} top_properties.GetTopPropertyRes
// @Failure      400 {object} string
// @Router       /topproperties/getbyidtopproperties/{top_properties_id} [get]
func (h *Handler) GetByIdTopProperties(c *gin.Context) {
	id := c.Param("top_properties_id")

	req := pbtop.GetTopPropertyReq{
		Id: id,
	}

	resp, err := h.TopPropertiesService.Get(c, &req)
	if err != nil {
		h.Log.Error(fmt.Sprintf("GetById top properties error: %v", err.Error()))
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, resp)
}

// GetAllTopProperties godoc
// @Summary      Get all top properties
// @Description  Retrieve a list of all top properties with pagination
// @Tags         top_properties
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        limit path int true "Limit"
// @Param        page  path int true "Page"
// @Success      200 {object} top_properties.GetAllTopPropertyRes
// @Failure      400 {object} string
// @Router       /topproperties/getalltopproperties/{limit}/{page} [get]
func (h *Handler) GetAllTopProperties(c *gin.Context) {
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

	req := pbtop.GetAllTopPropertyReq{
		Limit: limitInt32,
		Page:  pageInt32,
	}

	resp, err := h.TopPropertiesService.GetAll(c, &req)
	if err != nil {
		h.Log.Error(fmt.Sprintf("Getall properties error: %v", err.Error()))
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, resp)
}

// DeleteTopProperties godoc
// @Summary      Delete top properties
// @Description  Delete a top property by its ID
// @Tags         top_properties
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        top_properties_id path string true "Top Properties ID"
// @Success      200 {object} top_properties.DeleteTopPropertyRes
// @Failure      400 {object} string
// @Router       /topproperties/deletetopproperties/{top_properties_id} [delete]
func (h *Handler) DeleteTopProperties(c *gin.Context) {
	id := c.Param("top_properties_id")

	req := pbtop.DeleteTopPropertyReq{
		Id: id,
	}

	resp, err := h.TopPropertiesService.Delete(c, &req)
	if err != nil {
		h.Log.Error(fmt.Sprintf("Delete top properties error: %v", err.Error()))
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, resp)
}
