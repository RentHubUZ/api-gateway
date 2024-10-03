package handler

import (
	pbtarif "api_gateway/genproto/tariff"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateTarif godoc
// @Summary      Create a new tariff
// @Description  Create a new tariff with specified details
// @Tags         tariffs
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        body body tariff.CreateTariffReq true "Create Tariff Request"
// @Success      200 {object} tariff.CreateTariffRes
// @Failure      400 {object} string
// @Router       /api/tarif/createtarif [post]
func (h *Handler) CreateTarif(c *gin.Context) {
	var tarif pbtarif.CreateTariffReq
	if err := c.ShouldBindJSON(&tarif); err != nil {
		h.Log.Error("Error binding JSON: ", "error", err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.TarifService.Create(c, &tarif)
	if err != nil {
		h.Log.Error(fmt.Sprintf("Create tarif error: %v", err.Error()))
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, resp)
}

// UpdateTarif godoc
// @Summary      Update a tariff
// @Description  Update an existing tariff with specified details
// @Tags         tariffs
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        body body tariff.UpdateTariffReq true "Update Tariff Request"
// @Success      200 {object} tariff.UpdateTariffRes
// @Failure      400 {object} string
// @Router       /api/tarif/updatetarif [put]
func (h *Handler) UpdateTarif(c *gin.Context) {
	var tarifupdate pbtarif.UpdateTariffReq
	if err := c.ShouldBindJSON(&tarifupdate); err != nil {
		h.Log.Error("Error binding JSON: ", "error", err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.TarifService.Update(c, &tarifupdate)
	if err != nil {
		h.Log.Error(fmt.Sprintf("Updating tarif error: %v", err.Error()))
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, resp)
}

// GetByIdTarif godoc
// @Summary      Get tariff by ID
// @Description  Retrieve tariff details by tariff ID
// @Tags         tariffs
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        tarif_id path string true "Tariff ID"
// @Success      200 {object} tariff.GetTariffRes
// @Failure      400 {object} string
// @Router       /api/tarif/getbyidtarif/{tarif_id} [get]
func (h *Handler) GetByIdTarif(c *gin.Context) {
	id := c.Param("tarif_id")

	req := pbtarif.GetTariffReq{
		Id: id,
	}

	resp, err := h.TarifService.Get(c, &req)
	if err != nil {
		h.Log.Error(fmt.Sprintf("GetById tarif error: %v", err.Error()))
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, resp)
}

// GetAllTarif godoc
// @Summary      Get all tariffs
// @Description  Retrieve a list of all tariffs with pagination
// @Tags         tariffs
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        limit path int true "Limit"
// @Param        page  path int true "Page"
// @Success      200 {object} tariff.GetAllTariffRes
// @Failure      400 {object} string
// @Router       /api/tarif/getalltarif/{limit}/{page} [get]
func (h *Handler) GetAllTarif(c *gin.Context) {
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

	req := pbtarif.GetAllTariffReq{
		Limit: limitInt32,
		Page: pageInt32,
	}

	resp,err := h.TarifService.GetAll(c,&req)
	if err != nil {
		h.Log.Error(fmt.Sprintf("GetallTarif error: %v",err.Error()))
		c.JSON(400,gin.H{"error":err.Error()})
		return
	}

	c.JSON(200,resp)
}

// DeleteTarif godoc
// @Summary      Delete tariff
// @Description  Delete a tariff by tariff ID
// @Tags         tariffs
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        tarif_id path string true "Tariff ID"
// @Success      200 {object} tariff.DeleteTariffRes
// @Failure      400 {object} string
// @Router       /api/tarif/deletetarif/{tarif_id} [delete]
func (h *Handler) DeleteTarif(c *gin.Context) {
	id := c.Param("tarif_id")

	req := pbtarif.DeleteTariffReq{
		Id: id,
	}

	resp,err := h.TarifService.Delete(c,&req)
	if err != nil {
		h.Log.Error(fmt.Sprintf("Delete tarif error: %v",err.Error()))
		return
	}

	c.JSON(200,resp)
}