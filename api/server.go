package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/guidogimeno/evm/types"
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
	router.GET("/ping", s.handlePing)
	router.POST("/project-evaluation", s.handleProjectEvaluation)
}

func (s *Server) handlePing(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func (s *Server) handleProjectEvaluation(c *gin.Context) {
	project := types.Project{}

	if err := c.BindJSON(&project); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"cv":  project.CostVariance(),
		"cpi": project.CostPerformanceIndex(),
		"sv":  project.ScopeVariance(),
		"spi": project.ScopePerformanceIndex(),
	})
}
