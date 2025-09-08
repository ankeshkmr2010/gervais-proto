package controllers

import (
	gin "github.com/gin-gonic/gin"
)

type BaseController struct {
}

func NewBaseController() *BaseController {
	return &BaseController{}
}

func (b *BaseController) HealthCheck(c *gin.Context) {
	c.JSON(200, "ok")
}
