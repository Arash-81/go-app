package albums

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "github.com/Arash-81/go-app/src/metrics"
)

// Album represents the structure of an album
type Album struct {
    ID     string  `json:"id"`
    Title  string  `json:"title"`
    Artist string  `json:"artist"`
    Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []Album{
    {ID: "1", Title: "Album One", Artist: "Artist One", Price: 10.00},
    {ID: "2", Title: "Album Two", Artist: "Artist Two", Price: 20.00},
}

// GetAlbums retrieves the list of albums
func GetAlbums(c *gin.Context) {
    statusCode := strconv.Itoa(http.StatusOK)
    metrics.RequestCounter.WithLabelValues("GET", "/albums", statusCode).Inc()
    c.IndentedJSON(http.StatusOK, albums)
}

// PostAlbums adds a new album to the albums slice
func PostAlbums(c *gin.Context) {
    var newAlbum Album

    // Call BindJSON to bind the received JSON to newAlbum.
    if err := c.BindJSON(&newAlbum); err != nil {
        statusCode := strconv.Itoa(http.StatusBadRequest)
        metrics.RequestCounter.WithLabelValues("POST", "/albums", statusCode).Inc()
        c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
        return
    }

    albums = append(albums, newAlbum)
    statusCode := strconv.Itoa(http.StatusCreated)
    metrics.RequestCounter.WithLabelValues("POST", "/albums", statusCode).Inc()
    c.JSON(http.StatusCreated, newAlbum)
}