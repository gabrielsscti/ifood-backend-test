package tracks

import (
	"fmt"
	"github.com/gabrielsscti/ifood-backend-test/pkg/clients"
	"github.com/gabrielsscti/ifood-backend-test/pkg/clients/authorization"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var spotifyTrackClient TrackClient

func TestMain(t *testing.M) {
	clients.TryLoadEnvironmentFile()
	spotifyTrackClient = NewSpotifyTrackClient(authorization.NewClientCredentials(os.Getenv("SPOTIFY_CLIENT_ID"), os.Getenv("SPOTIFY_CLIENT_SECRET"), SpotifyTokenURL))
	os.Exit(t.Run())
}

func TestFetchTracks(t *testing.T) {
	val, err := spotifyTrackClient.FetchTracks(Rock)

	if err != nil {
		t.Errorf(err.Error())
	}
	fmt.Println(val)
	assert.NotNil(t, val)
}
