package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"my-project/cmd/product/internal/models"
	"my-project/cmd/product/internal/service"
)

type ProductHandler struct {
    service service.ProductService
}

func NewProductHandler(service service.ProductService) *ProductHandler {
    return &ProductHandler{service: service}
}

func (h *ProductHandler) CreateProduct(c *gin.Context) {
    var input models.Product
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    product, err := h.service.CreateProduct(input.Title, input.Desc)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, product)
}

func (h *ProductHandler) GetProduct(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    product, err := h.service.GetProductByID(uint(id))
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
        return
    }
    c.JSON(http.StatusOK, product)
}

func (h *ProductHandler) UpdateProduct(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    var input struct {
        Title string `json:"title" binding:"required"`
        Desc  string `json:"desc"`
    }
    if err := c.BindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    product, err := h.service.UpdateProduct(uint(id), input.Title, input.Desc)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, product)
}

func (h *ProductHandler) DeleteProduct(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    if err := h.service.DeleteProduct(uint(id)); err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
        return
    }
    c.JSON(http.StatusNoContent, nil)
}

func (h *ProductHandler) ListProducts(c *gin.Context) {
    products, err := h.service.ListProducts()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, products)
}