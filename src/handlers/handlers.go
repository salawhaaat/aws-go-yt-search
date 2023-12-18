package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/salawhaaat/aws-go-yt-search/src/search"
)

func SearchHandler(ctx *gin.Context) {
	switch ctx.Request.Method {
	case http.MethodGet:
		handleGetRequest(ctx)
	case http.MethodPost:
		handlePostRequest(ctx)
	default:
		ctx.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Method not allowed"})
	}
}

func handleGetRequest(ctx *gin.Context) {
	query := ctx.Query("q")
	if query == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Missing query parameter"})
		return
	}
	link, err := search.Search(query)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"video": link})
}

func handlePostRequest(ctx *gin.Context) {
	var links []string
	if err := ctx.ShouldBindJSON(&links); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res := search.ProcessLinks(links)
	ctx.JSON(http.StatusOK, gin.H{"videos": res})
}
