package handler

import (
	pbpay "api_gateway/genproto/payment"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreatePayment godoc
// @Summary      Create a new payment
// @Description  Create a new payment with specified details
// @Tags         payments
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        body body payment.CreatePaymentReq true "Create Payment Request"
// @Success      200 {object} payment.CreatePaymentRes
// @Failure      400 {object} string
// @Router       /api/payment/createpayment [post]
func (h *Handler) CreatePayment(c *gin.Context) {
	var paymentCreate pbpay.CreatePaymentReq
	if err := c.ShouldBindJSON(&paymentCreate); err != nil {
		h.Log.Error("Error binding JSON: ", "error", err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.PaymentService.Create(c, &paymentCreate)
	if err != nil {
		h.Log.Error(fmt.Sprintf("payment creating error: %v", err.Error()))
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, resp)
}

// GetPayment godoc
// @Summary      Get payment by ID
// @Description  Retrieve payment details by payment ID
// @Tags         payments
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        payment_id path string true "Payment ID"
// @Success      200 {object} payment.GetPaymentRes
// @Failure      400 {object} string
// @Router       /api/payment/getbyidpayment/{payment_id} [get]
func (h *Handler) GetPayment(c *gin.Context) {
	id := c.Param("payment_id")

	req := pbpay.GetPaymentReq{
		Id: id,
	}

	resp, err := h.PaymentService.Get(c, &req)
	if err != nil {
		h.Log.Error(fmt.Sprintf("getpayment error: %v", err.Error()))
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, resp)
}

// GetAllPayment godoc
// @Summary      Get all payments
// @Description  Retrieve a list of all payments with pagination
// @Tags         payments
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        limit path int true "Limit"
// @Param        page  path int true "Page"
// @Success      200 {object} payment.GetAllPaymentRes
// @Failure      400 {object} string
// @Router       /api/payment/getallpayment/{limit}/{page} [get]
func (h *Handler) GetAllPaymet(c *gin.Context) {
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

	req := pbpay.GetAllPaymentReq{
		Limit: limitInt32,
		Page:  pageInt32,
	}

	resp, err := h.PaymentService.GetAll(c, &req)
	if err != nil {
		h.Log.Error(fmt.Sprintf("GetAllPayment error: %v", err.Error()))
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, resp)
}

// DeletePayment godoc
// @Summary      Delete payment
// @Description  Delete a payment by payment ID
// @Tags         payments
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        payment_id path string true "Payment ID"
// @Success      200 {object} payment.DeletePaymentRes
// @Failure      400 {object} string
// @Router       /api/payment/deletepayment/{payment_id} [delete]
func (h *Handler) DeletePayment(c *gin.Context) {
	id := c.Param("payment_id")

	req := pbpay.DeletePaymentReq{
		Id: id,
	}

	resp, err := h.PaymentService.Delete(c, &req)
	if err != nil {
		h.Log.Error(fmt.Sprintf("Payment deleting error: %v", err.Error()))
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, resp)
}
