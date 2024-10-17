package handler

import (
	report "api_gateway/genproto/report"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateReport godoc
// @Summary      Create a new report
// @Description  Add a new user report
// @Tags         report
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        body body report.CreateReportRequest true "Create Report Request"
// @Success      200 {object} report.CreateReportResponse
// @Failure      400 {object} string
// @Router       /report/create [post]
func (h *Handler) CreateReport(c *gin.Context) {
	var req report.CreateReportRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if req.UserId == "" || req.UserId == "string" {
		UserId, err := getUserID(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		req.UserId = UserId
	}

	res, err := h.ReportService.Create(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// GetReportById godoc
// @Summary      Get a report by ID
// @Description  Retrieve a specific report by its ID
// @Tags         report
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        id path string true "Report ID"
// @Success      200 {object} report.GetReportResponse
// @Failure      400 {object} string
// @Router       /report/get/{id} [get]
func (h *Handler) GetReportById(c *gin.Context) {
	var req report.GetReportRequest
	req.Id = c.Param("id")
	if req.Id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	res, err := h.ReportService.Get(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// DeleteReport godoc
// @Summary      Delete a report
// @Description  Remove a report by its ID
// @Tags         report
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        id path string true "Report ID"
// @Success      200 {object} report.Void
// @Failure      400 {object} string
// @Router       /report/delete/{id} [delete]
func (h *Handler) DeleteReport(c *gin.Context) {
	var req report.DeleteReportRequest
	req.Id = c.Param("id")
	if req.Id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	_, err := h.ReportService.Delete(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Report deleted successfully"})
}
