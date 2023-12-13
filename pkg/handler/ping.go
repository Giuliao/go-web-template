package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"my.net/module/api/types"
)

// Ping godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {object} types.Pong
// @Router /api/v1/ping [get]
func Ping(c *gin.Context) {

	c.JSON(http.StatusOK, types.Pong{
		Message: "pong",
	})
}

// PPing godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {object} types.Pong
// @Router /api/v2/ping [get]
func PPing(c *gin.Context) {

	c.JSON(http.StatusOK, types.Pong{
		Message: "pong",
	})
}
