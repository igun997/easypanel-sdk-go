package easypanel

import "context"

// ServicesService handles service-related API operations.
type ServicesService struct {
	client *httpClient
}

// Create creates a new service of the given type.
func (s *ServicesService) Create(ctx context.Context, st ServiceType, params CreateServiceParams) (RestResponse[Service], error) {
	var resp RestResponse[Service]
	err := s.client.post(ctx, serviceRoute(routeCreateService, st), params, &resp)
	return resp, err
}

// Inspect returns detailed information about a service.
func (s *ServicesService) Inspect(ctx context.Context, st ServiceType, params SelectService) (RestResponse[Service], error) {
	var resp RestResponse[Service]
	err := s.client.get(ctx, serviceRoute(routeInspectService, st), params, &resp)
	return resp, err
}

// Destroy deletes a service.
func (s *ServicesService) Destroy(ctx context.Context, st ServiceType, params SelectService) error {
	return s.client.post(ctx, serviceRoute(routeDestroyService, st), params, nil)
}

// Deploy triggers a deployment for a service.
func (s *ServicesService) Deploy(ctx context.Context, st ServiceType, params SelectService) error {
	return s.client.post(ctx, serviceRoute(routeDeployService, st), params, nil)
}

// Disable disables a service.
func (s *ServicesService) Disable(ctx context.Context, st ServiceType, params SelectService) error {
	return s.client.post(ctx, serviceRoute(routeDisableService, st), params, nil)
}

// Enable enables a service.
func (s *ServicesService) Enable(ctx context.Context, st ServiceType, params SelectService) error {
	return s.client.post(ctx, serviceRoute(routeEnableService, st), params, nil)
}

// ExposeService exposes a service port externally.
func (s *ServicesService) ExposeService(ctx context.Context, st ServiceType, params ExposeServiceParams) error {
	return s.client.post(ctx, serviceRoute(routeExposeService, st), params, nil)
}

// RefreshDeployToken refreshes the deploy token for a service.
func (s *ServicesService) RefreshDeployToken(ctx context.Context, st ServiceType, params SelectService) error {
	return s.client.post(ctx, serviceRoute(routeRefreshDeployToken, st), params, nil)
}

// UpdateSourceGithub updates the GitHub source configuration.
func (s *ServicesService) UpdateSourceGithub(ctx context.Context, st ServiceType, params UpdateGithub) error {
	return s.client.post(ctx, serviceRoute(routeUpdateSourceGithub, st), params, nil)
}

// UpdateSourceGit updates the Git source configuration.
func (s *ServicesService) UpdateSourceGit(ctx context.Context, st ServiceType, params UpdateGit) error {
	return s.client.post(ctx, serviceRoute(routeUpdateSourceGit, st), params, nil)
}

// UpdateSourceImage updates the Docker image source configuration.
func (s *ServicesService) UpdateSourceImage(ctx context.Context, st ServiceType, params UpdateImage) error {
	return s.client.post(ctx, serviceRoute(routeUpdateSourceImage, st), params, nil)
}

// UpdateBuild updates the build configuration for a service.
func (s *ServicesService) UpdateBuild(ctx context.Context, st ServiceType, params UpdateBuildParams) error {
	return s.client.post(ctx, serviceRoute(routeUpdateBuild, st), params, nil)
}

// UpdateEnv updates the environment variables for a service.
func (s *ServicesService) UpdateEnv(ctx context.Context, st ServiceType, params UpdateEnv) error {
	return s.client.post(ctx, serviceRoute(routeUpdateEnv, st), params, nil)
}

// UpdateDomains updates the domain configuration for a service.
func (s *ServicesService) UpdateDomains(ctx context.Context, st ServiceType, params CreateServiceParams) error {
	return s.client.post(ctx, serviceRoute(routeUpdateDomains, st), params, nil)
}

// UpdateRedirects updates the redirect rules for a service.
func (s *ServicesService) UpdateRedirects(ctx context.Context, st ServiceType, params UpdateRedirects) error {
	return s.client.post(ctx, serviceRoute(routeUpdateRedirects, st), params, nil)
}

// UpdateBasicAuth updates the basic auth configuration for a service.
func (s *ServicesService) UpdateBasicAuth(ctx context.Context, st ServiceType, params UpdateBasicAuth) error {
	return s.client.post(ctx, serviceRoute(routeUpdateBasicAuth, st), params, nil)
}

// UpdateMounts updates the mount configuration for a service.
func (s *ServicesService) UpdateMounts(ctx context.Context, st ServiceType, params MountParams) error {
	return s.client.post(ctx, serviceRoute(routeUpdateMounts, st), params, nil)
}

// UpdatePorts updates the port mappings for a service.
func (s *ServicesService) UpdatePorts(ctx context.Context, st ServiceType, params UpdatePorts) error {
	return s.client.post(ctx, serviceRoute(routeUpdatePorts, st), params, nil)
}

// UpdateResources updates the resource limits for a service.
func (s *ServicesService) UpdateResources(ctx context.Context, st ServiceType, params UpdateResources) error {
	return s.client.post(ctx, serviceRoute(routeUpdateResources, st), params, nil)
}

// UpdateDeploy updates the deployment configuration for a service.
func (s *ServicesService) UpdateDeploy(ctx context.Context, st ServiceType, params DeployParams) error {
	return s.client.post(ctx, serviceRoute(routeUpdateDeploy, st), params, nil)
}

// UpdateBackup updates the backup configuration for a service.
func (s *ServicesService) UpdateBackup(ctx context.Context, st ServiceType, params UpdateBackupParams) error {
	return s.client.post(ctx, serviceRoute(routeUpdateBackup, st), params, nil)
}

// UpdateAdvanced updates the advanced settings for a service.
func (s *ServicesService) UpdateAdvanced(ctx context.Context, st ServiceType, params UpdateAdvancedParams) error {
	return s.client.post(ctx, serviceRoute(routeUpdateAdvanced, st), params, nil)
}

// UpdateSourceInline updates a compose service with inline docker-compose content.
func (s *ServicesService) UpdateSourceInline(ctx context.Context, st ServiceType, params UpdateSourceInline) error {
	return s.client.post(ctx, serviceRoute(routeUpdateSourceInline, st), params, nil)
}

// UpdateSourceGitCompose updates a compose service with a Git source.
func (s *ServicesService) UpdateSourceGitCompose(ctx context.Context, st ServiceType, params UpdateSourceGitCompose) error {
	return s.client.post(ctx, serviceRoute(routeUpdateSourceGit, st), params, nil)
}

// GetServiceLogs retrieves logs for a service.
func (s *ServicesService) GetServiceLogs(ctx context.Context, params SelectService) (RestResponse[string], error) {
	var resp RestResponse[string]
	err := s.client.get(ctx, routeGetServiceLogs, params, &resp)
	return resp, err
}
