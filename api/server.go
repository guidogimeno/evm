package api

import (
	"fmt"
	"net/http"

	"github.com/ggimeno96/project/types"
	"github.com/gin-gonic/gin"
)

type Server struct {
	ListenAddr string
	router     *gin.Engine
}

func (s *Server) Start() error {
	router := gin.Default()
	s.mapRoutes(router)
	return router.Run()
}

func (s *Server) mapRoutes(router *gin.Engine) {
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.POST("/evm", func(c *gin.Context) {
		var evm types.EarnedValueManagement

		if err := c.BindJSON(&evm); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		fmt.Println((evm))

		c.JSON(http.StatusOK, gin.H{
			"cv":  evm.CostVariance(),
			"cpi": evm.CostPerformanceIndex(),
			"sv":  evm.ScopeVariance(),
			"spi": evm.ScopePerformanceIndex(),
		})
	})
}
