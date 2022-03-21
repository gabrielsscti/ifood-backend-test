package tracks

import (
	"encoding/json"
	"errors"
	"github.com/gabrielsscti/ifood-backend-test/pkg/clients/authorization"
	"io/ioutil"
	"net/http"
	"strings"
)

type SpotifyTrackClient struct {
	authorization authorization.Authorizable
}

type PlaylistsByCategoryRequest struct {
	Playlists struct {
		Items []struct {
			Id string `json:"id"`
		} `json:"items"`
	} `json:"playlists"`
}

type PlaylistRequest struct {
	Tracks struct {
		Items []struct {
			Track struct {
				Name string `json:"name"`
			} `json:"track"`
		} `json:"items"`
	} `json:"tracks"`
}

const spotifyBrowseCategoryURL = "https://api.spotify.com/v1/browse/categories/@CategoryID/playlists"
const categoryWildcard = "@CategoryID"

const spotifyPlaylistURL = "https://api.spotify.com/v1/playlists/@PlaylistID"
const playlistWildcard = "@PlaylistID"

func NewSpotifyTrackClient(authorizer authorization.Authorizer) SpotifyTrackClient {
	return SpotifyTrackClient{
		authorization.SpotifyAuthorization{Authorizer: authorizer},
	}
}

func (s SpotifyTrackClient) getPlaylistID(auth authorization.Authorization, musicType MusicType) (string, error) {
	urlWithCategory := strings.Replace(spotifyBrowseCategoryURL, categoryWildcard, musicType.String(), -1)
	req, err := http.NewRequest("GET", urlWithCategory, nil)
	if err != nil {
		return "", err
	}
	auth.Authorize(req)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	if resp.StatusCode != http.StatusOK {
		return "", errors.New(resp.Status)
	}
	defer resp.Body.Close()

	byteValue, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var playlists PlaylistsByCategoryRequest
	err = json.Unmarshal(byteValue, &playlists)
	if err != nil {
		return "", err
	}

	return playlists.Playlists.Items[0].Id, nil
}

func (s SpotifyTrackClient) getTracks(auth authorization.Authorization, playlistID string) (Tracks, error) {
	urlWithPlaylistID := strings.Replace(spotifyPlaylistURL, playlistWildcard, playlistID, -1)
	req, err := http.NewRequest("GET", urlWithPlaylistID, nil)
	if err != nil {
		return nil, err
	}

	auth.Authorize(req)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}
	defer resp.Body.Close()

	byteValue, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var playlist PlaylistRequest
	err = json.Unmarshal(byteValue, &playlist)
	if err != nil {
		return nil, err
	}

	var tracks Tracks
	for _, item := range playlist.Tracks.Items {
		tracks = append(tracks, item.Track.Name)
	}
	return tracks, nil
}

func (s SpotifyTrackClient) FetchTracks(musicType MusicType) (Tracks, error) {
	auth, err := s.authorization.GetAuthorization()
	if err != nil {
		return nil, err
	}

	playlistID, err := s.getPlaylistID(auth, musicType)
	if err != nil {
		return nil, err
	}

	tracks, err := s.getTracks(auth, playlistID)
	if err != nil {
		return nil, err
	}

	return tracks, nil
}
