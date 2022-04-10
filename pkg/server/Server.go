package server

import (
	"context"
	"fmt"
	"github.com/gabrielsscti/ifood-backend-test/pkg/server/handlers"
	"github.com/gin-gonic/gin"
)

func RunServer(port string) {
	ctx := context.Background()
	tracksHandler := handlers.NewTracksHandler(ctx)

	router := gin.Default()

	router.POST("/city", tracksHandler.TracksByCityName)
	router.POST("/coords", tracksHandler.TracksByCoordinate)

	if err := router.Run(port); err != nil {
		fmt.Println(err.Error())
	}
}
