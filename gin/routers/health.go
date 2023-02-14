package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jjh930301/needsss_common/database"
)

// @Tags health
// @Summary health check
// @Accept  json
// @Produce  json
// @Router /health/check [get]
func HealthCheck(c *gin.Context) {
	db, err := database.DB.DB()
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"result": false,
		})
		return
	}
	stats := db.Stats()
	c.JSON(http.StatusOK, gin.H{
		"result": true,
		"idle":   stats.Idle,
	})
}
