package config

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func ConfigurationCors() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "http://localhost:3000", "http://3.219.79.220"},
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},  
        AllowHeaders:     []string{"Content-Type", "Authorization"},            
        AllowCredentials: true,
        MaxAge:           12 * time.Hour,
	})
}