package authorization

type SpotifyAuthorization struct {
	Authorizer
}

func (s SpotifyAuthorization) GetAuthorization() (Authorization, error) {
	authorization, err := s.Authorizer.SetAuthorization()
	if err != nil {
		return nil, err
	}

	return authorization, nil
}
