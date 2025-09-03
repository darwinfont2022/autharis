package main

import (
	"fmt"
	"log"
	"os"

	"github.com/darwinfont2022/autharis/internal/database/db"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Cargar variables de entorno
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️  No se encontró el archivo .env, usando variables del sistema")
	}

	// Leer puerto desde .env o usar por defecto 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	db.InitDB()

	// Inicializar router de Gin
	router := gin.Default()

	// Endpoint de prueba
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Levantar servidor
	addr := fmt.Sprintf(":%s", port)
	log.Printf("🚀 Autharis corriendo en http://localhost%s\n", addr)
	if err := router.Run(addr); err != nil {
		log.Fatalf("❌ Error iniciando servidor: %v", err)
	}
}
