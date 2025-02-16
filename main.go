package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/Arash-81/go-app/src/albums"
)

var (
	httpStatusCodes = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests by status code",
		},
		[]string{"code"},
	)
)

func init() {
	prometheus.MustRegister(httpStatusCodes)
}

func main() {
	// initialise Gin router:
	router := gin.Default()

	// Middleware to track status codes
	router.Use(func(c *gin.Context) {
		c.Next()
		httpStatusCodes.WithLabelValues(http.StatusText(c.Writer.Status())).Inc()
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