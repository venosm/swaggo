package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "test_enum_app/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type StavZpravy string

const (
	StavAktivni     StavZpravy = "AKTIVNI"
	StavArchivovana StavZpravy = "ARCHIVOVANA"
	StavSmazana     StavZpravy = "SMAZANA"
)

type Zprava struct {
	ID          int        `json:"id" example:"1"`
	StavZpravy  StavZpravy `json:"stav_zpravy,omitempty" validate:"omitempty,oneof=AKTIVNI ARCHIVOVANA SMAZANA AKTIVNI ARCHIVOVANA SMAZANA" example:"AKTIVNI" description:"Stav zprávy (defaultní hodnota je aktivní)."`
}

// @title Test API pro enum duplicity
// @version 1.0
// @host localhost:8080
// @BasePath /api/v1
func main() {
	r := gin.Default()
	r.GET("/test", GetZprava)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run(":8080")
}

// GetZprava test endpoint
// @Summary Test endpoint
// @Description Test endpoint for enum duplication
// @Tags test
// @Produce json
// @Success 200 {object} Zprava
// @Router /test [get]
func GetZprava(c *gin.Context) {
	zprava := Zprava{
		ID: 1,
		StavZpravy: StavAktivni,
	}
	c.JSON(http.StatusOK, zprava)
}