package albums

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/Arash-81/go-app/src/metrics"
)

// Album represents the structure of an album
type Album struct {
    ID     string  `json:"id"`
    Title  string  `json:"title"`
    Artist string  `json:"artist"`
}

// albums slice to seed record album data.
var albums = []Album{
    {ID: "1", Title: "Album One", Artist: "Artist One"},
    {ID: "2", Title: "Album Two", Artist: "Artist Two"},
}

// GetAlbums retrieves the list of albums
func GetAlbums(c *gin.Context) {
    metrics.RequestCounter.WithLabelValues("GET", "/albums").Inc()
    c.IndentedJSON(http.StatusOK, albums)
}

// PostAlbums adds a new album to the albums slice
func PostAlbums(c *gin.Context) {
    var newAlbum Album

    // Call BindJSON to bind the received JSON to newAlbum.
    if err := c.BindJSON(&newAlbum); err != nil {
        metrics.RequestCounter.WithLabelValues("POST", "/albums").Inc()
        c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
        return
    }

    albums = append(albums, newAlbum)
    metrics.RequestCounter.WithLabelValues("POST", "/albums").Inc()
    c.JSON(http.StatusCreated, newAlbum)
}