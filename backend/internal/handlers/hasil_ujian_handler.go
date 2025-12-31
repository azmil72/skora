package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"backend/internal/models"
)

type HasilUjianHandler struct {
	DB *gorm.DB
}

func NewHasilUjianHandler(db *gorm.DB) *HasilUjianHandler {
	return &HasilUjianHandler{DB: db}
}

func (h *HasilUjianHandler) CreateHasilUjian(c *gin.Context) {
	var hasil models.HasilUjian
	if err := c.ShouldBindJSON(&hasil); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.DB.Create(&hasil).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, hasil)
}

func (h *HasilUjianHandler) GetHasilUjians(c *gin.Context) {
	var hasils []models.HasilUjian
	if err := h.DB.Preload("SesiUjian").Find(&hasils).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, hasils)
}

func (h *HasilUjianHandler) GetHasilUjian(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var hasil models.HasilUjian
	if err := h.DB.Preload("SesiUjian").First(&hasil, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Hasil ujian not found"})
		return
	}

	c.JSON(http.StatusOK, hasil)
}

func (h *HasilUjianHandler) UpdateHasilUjian(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var hasil models.HasilUjian
	if err := h.DB.First(&hasil, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Hasil ujian not found"})
		return
	}

	if err := c.ShouldBindJSON(&hasil); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.DB.Save(&hasil).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, hasil)
}

func (h *HasilUjianHandler) DeleteHasilUjian(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := h.DB.Delete(&models.HasilUjian{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Hasil ujian deleted successfully"})
}