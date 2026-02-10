package easypanel

import "context"

// SettingsService handles settings-related API operations.
type SettingsService struct {
	client *httpClient
}

// ChangeCredentials changes the user's email and password.
func (s *SettingsService) ChangeCredentials(ctx context.Context, params ChangeCredentialsParams) error {
	return s.client.post(ctx, routeChangeCredentials, params, nil)
}

// GetGithubToken returns the configured GitHub token.
func (s *SettingsService) GetGithubToken(ctx context.Context) (RestResponse[string], error) {
	var resp RestResponse[string]
	err := s.client.get(ctx, routeGetGithubToken, nil, &resp)
	return resp, err
}

// GetLetsEncryptEmail returns the configured Let's Encrypt email.
func (s *SettingsService) GetLetsEncryptEmail(ctx context.Context) (RestResponse[string], error) {
	var resp RestResponse[string]
	err := s.client.get(ctx, routeGetLetsEncryptEmail, nil, &resp)
	return resp, err
}

// GetPanelDomain returns the panel domain configuration.
func (s *SettingsService) GetPanelDomain(ctx context.Context) (RestResponse[PanelDomain], error) {
	var resp RestResponse[PanelDomain]
	err := s.client.get(ctx, routeGetPanelDomain, nil, &resp)
	return resp, err
}

// GetServerIp returns the server IP address.
func (s *SettingsService) GetServerIp(ctx context.Context) (RestResponse[string], error) {
	var resp RestResponse[string]
	err := s.client.get(ctx, routeGetServerIp, nil, &resp)
	return resp, err
}

// GetTraefikCustomConfig returns the custom Traefik configuration.
func (s *SettingsService) GetTraefikCustomConfig(ctx context.Context) (RestResponse[string], error) {
	var resp RestResponse[string]
	err := s.client.get(ctx, routeGetTraefikCustomConfig, nil, &resp)
	return resp, err
}

// PruneDockerBuilder triggers a Docker builder cache prune.
func (s *SettingsService) PruneDockerBuilder(ctx context.Context) (RestResponse[string], error) {
	var resp RestResponse[string]
	err := s.client.post(ctx, routePruneDockerBuilder, nil, &resp)
	return resp, err
}

// PruneDockerImages triggers a Docker image prune.
func (s *SettingsService) PruneDockerImages(ctx context.Context) (RestResponse[string], error) {
	var resp RestResponse[string]
	err := s.client.post(ctx, routePruneDockerImages, nil, &resp)
	return resp, err
}

// RefreshServerIp refreshes the server IP address.
func (s *SettingsService) RefreshServerIp(ctx context.Context) error {
	return s.client.post(ctx, routeRefreshServerIp, nil, nil)
}

// RestartEasypanel restarts the Easypanel service.
func (s *SettingsService) RestartEasypanel(ctx context.Context) error {
	return s.client.post(ctx, routeRestartEasypanel, nil, nil)
}

// RestartTraefik restarts the Traefik reverse proxy.
func (s *SettingsService) RestartTraefik(ctx context.Context) error {
	return s.client.post(ctx, routeRestartTraefik, nil, nil)
}

// SetDockerPruneDaily enables or disables daily Docker pruning.
func (s *SettingsService) SetDockerPruneDaily(ctx context.Context, params PruneDockerDailyParams) (RestResponse[bool], error) {
	var resp RestResponse[bool]
	err := s.client.post(ctx, routeSetPruneDockerDaily, params, &resp)
	return resp, err
}

// SetGithubToken sets the GitHub token.
func (s *SettingsService) SetGithubToken(ctx context.Context, params GithubTokenParams) (RestResponse[string], error) {
	var resp RestResponse[string]
	err := s.client.post(ctx, routeSetGithubToken, params, &resp)
	return resp, err
}

// SetLetsEncryptEmail sets the Let's Encrypt email.
func (s *SettingsService) SetLetsEncryptEmail(ctx context.Context, params LetsEncryptParams) (RestResponse[string], error) {
	var resp RestResponse[string]
	err := s.client.post(ctx, routeSetLetsEncryptEmail, params, &resp)
	return resp, err
}

// SetPanelDomain sets the panel domain configuration.
func (s *SettingsService) SetPanelDomain(ctx context.Context, params PanelDomainParams) error {
	return s.client.post(ctx, routeSetPanelDomain, params, nil)
}

// UpdateTraefikCustomConfig updates the custom Traefik configuration.
func (s *SettingsService) UpdateTraefikCustomConfig(ctx context.Context, params TraefikConfParams) error {
	return s.client.post(ctx, routeUpdateTraefikCustomConfig, params, nil)
}
