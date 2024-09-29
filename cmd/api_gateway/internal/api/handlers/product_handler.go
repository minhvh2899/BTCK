package handlers

import (
	"my-project/cmd/api_gateway/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
    productService *service.ProductService
}

func NewProductHandler(productService *service.ProductService) *ProductHandler {
    return &ProductHandler{productService: productService}
}

func (h *ProductHandler) CreateProduct(c *gin.Context) {
    var input struct {
        Name        string  `json:"name" binding:"required"`
        Description string  `json:"description" binding:"required"`
        Price       float64 `json:"price" binding:"required"`
    }

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    response, err := h.productService.CreateProduct(c.Request.Context(), input.Name, input.Description, input.Price)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, response)
}

func (h *ProductHandler) GetProduct(c *gin.Context) {
    id := c.Param("id")

    response, err := h.productService.GetProduct(c.Request.Context(), id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, response)
}

func (h *ProductHandler) ListProducts(c *gin.Context) {
    page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
    limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

    response, err := h.productService.ListProducts(c.Request.Context(), int32(page), int32(limit))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, response)
}

func (h *ProductHandler) UpdateProduct(c *gin.Context) {
    id := c.Param("id")

    var input struct {
        Name        string  `json:"name" binding:"required"`
        Description string  `json:"description" binding:"required"`
        Price       float64 `json:"price" binding:"required"`
    }

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    response, err := h.productService.UpdateProduct(c.Request.Context(), id, input.Name, input.Description, input.Price)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, response)
}

func (h *ProductHandler) DeleteProduct(c *gin.Context) {
    id := c.Param("id")

    response, err := h.productService.DeleteProduct(c.Request.Context(), id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

	c.JSON(http.StatusOK, response)
}
