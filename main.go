package main

import (
    "log"

    "github.com/gin-gonic/gin"
    "github.com/prometheus/client_golang/prometheus/promhttp"
    "github.com/Arash-81/go-app/src/albums"
    "github.com/Arash-81/go-app/src/metrics"
)

func main() {
    // initialise Gin router:
    router := gin.Default()

    // Middleware to track status codes
    router.Use(func(c *gin.Context) {
        c.Next()
        metrics.RequestCounter.WithLabelValues(c.Request.Method, c.FullPath()).Inc()
    })

    router.GET("/albums", albums.GetAlbums)
    router.POST("/albums", albums.PostAlbums)

    // Prometheus metrics endpoint
    router.GET("/metrics", gin.WrapH(promhttp.Handler()))

    // run the server on port 3000:
    err := router.Run(":3000")
    if err != nil {
        log.Fatalf("[Error] failed to start Gin server due to: %v", err)
    }
}