package handlers

import (
	"context"
	"github.com/gabrielsscti/ifood-backend-test/pkg/clients/authorization"
	"github.com/gabrielsscti/ifood-backend-test/pkg/clients/tracks"
	"github.com/gabrielsscti/ifood-backend-test/pkg/clients/weather"
	"github.com/gabrielsscti/ifood-backend-test/pkg/server/models"
	track_manager "github.com/gabrielsscti/ifood-backend-test/pkg/track-manager"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

type TracksHandler struct {
	ctx context.Context
}

func NewTracksHandler(ctx context.Context) *TracksHandler {
	return &TracksHandler{ctx: ctx}
}

func (t *TracksHandler) createTrackClientByService(service string) tracks.TrackClient {
	switch service {
	case "spotify":
		return tracks.NewSpotifyTrackClient(authorization.NewClientCredentials(
			os.Getenv("SPOTIFY_CLIENT_ID"),
			os.Getenv("SPOTIFY_CLIENT_SECRET"),
			tracks.SpotifyTokenURL))
	default:
		return tracks.NewSpotifyTrackClient(authorization.NewClientCredentials(
			os.Getenv("SPOTIFY_CLIENT_ID"),
			os.Getenv("SPOTIFY_CLIENT_SECRET"),
			tracks.SpotifyTokenURL))
	}
}

func (t *TracksHandler) TracksByCityName(c *gin.Context) {
	var cityTracks models.CityTracksRequest
	if err := c.ShouldBind(&cityTracks); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	trackClient := t.createTrackClientByService(cityTracks.Service)
	weatherClient := weather.NewOpenWeatherClient(authorization.ApiKey{ApiKey: os.Getenv("OPEN_WEATHER_API_KEY")})

	trackManager := track_manager.CreateTrackManager(trackClient, weatherClient)

	resp, err := trackManager.GetPlaylist(&cityTracks.Location)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, resp)
}

func (t *TracksHandler) TracksByCoordinate(c *gin.Context) {
	var cityTracks models.CoordinateTracksRequest
	if err := c.ShouldBind(cityTracks); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	trackClient := t.createTrackClientByService(cityTracks.Service)
	weatherClient := weather.NewOpenWeatherClient(authorization.ApiKey{ApiKey: os.Getenv("OPEN_WEATHER_API_KEY")})

	trackManager := track_manager.CreateTrackManager(trackClient, weatherClient)

	resp, err := trackManager.GetPlaylist(&cityTracks.Coordinates)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, resp)
}
