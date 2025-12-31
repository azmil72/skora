package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"backend/internal/models"
)

type PertanyaanHandler struct {
	DB *gorm.DB
}

func NewPertanyaanHandler(db *gorm.DB) *PertanyaanHandler {
	return &PertanyaanHandler{DB: db}
}

func (h *PertanyaanHandler) CreatePertanyaan(c *gin.Context) {
	var pertanyaan models.Pertanyaan
	if err := c.ShouldBindJSON(&pertanyaan); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.DB.Create(&pertanyaan).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, pertanyaan)
}

func (h *PertanyaanHandler) GetPertanyaans(c *gin.Context) {
	var pertanyaans []models.Pertanyaan
	if err := h.DB.Preload("Room").Preload("QuestionOptions").Find(&pertanyaans).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, pertanyaans)
}

func (h *PertanyaanHandler) GetPertanyaan(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var pertanyaan models.Pertanyaan
	if err := h.DB.Preload("Room").Preload("QuestionOptions").First(&pertanyaan, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pertanyaan not found"})
		return
	}

	c.JSON(http.StatusOK, pertanyaan)
}

func (h *PertanyaanHandler) UpdatePertanyaan(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var pertanyaan models.Pertanyaan
	if err := h.DB.First(&pertanyaan, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pertanyaan not found"})
		return
	}

	if err := c.ShouldBindJSON(&pertanyaan); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.DB.Save(&pertanyaan).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, pertanyaan)
}

func (h *PertanyaanHandler) DeletePertanyaan(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := h.DB.Delete(&models.Pertanyaan{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Pertanyaan deleted successfully"})
}