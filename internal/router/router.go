package router

import (
	"net/http"
	"os"

	"github.com/dioncodes/go-surreal-starter/handlers/userHandler"
	"github.com/dioncodes/go-surreal-starter/pkg/env"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Setup(env *env.Env, r *gin.Engine) {
	r.SetTrustedProxies([]string{"127.0.0.1"})

	// CORS middleware
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowHeaders = []string{"Content-Type", "Origin", "Authorization", "Debug"}

	if os.Getenv("ENV") == "dev" {
		corsConfig.AllowAllOrigins = true
	} else {
		corsConfig.AllowOrigins = []string{"http://localhost:5173", "http://localhost:8080", "https://example.com"}
	}

	r.Use(cors.New(corsConfig))

	// r.LoadHTMLGlob(os.Getenv("BASE_DIR") + "/templates/email/*.html")

	userHandler.Register(env, r)

	r.GET("", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"details": "system up and running",
		})
	})
}
