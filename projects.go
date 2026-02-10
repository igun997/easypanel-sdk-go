package easypanel

import "context"

// ProjectsService handles project-related API operations.
type ProjectsService struct {
	client *httpClient
}

// CanCreate checks if a new project can be created.
func (s *ProjectsService) CanCreate(ctx context.Context) (RestResponse[bool], error) {
	var resp RestResponse[bool]
	err := s.client.get(ctx, routeCanCreateProject, nil, &resp)
	return resp, err
}

// Create creates a new project.
func (s *ProjectsService) Create(ctx context.Context, params ProjectName) (RestResponse[ProjectInfo], error) {
	var resp RestResponse[ProjectInfo]
	err := s.client.post(ctx, routeCreateProject, params, &resp)
	return resp, err
}

// Destroy deletes a project.
func (s *ProjectsService) Destroy(ctx context.Context, params ProjectName) error {
	return s.client.post(ctx, routeDestroyProject, params, nil)
}

// Inspect returns detailed information about a project including its services.
func (s *ProjectsService) Inspect(ctx context.Context, params ProjectQuery) (RestResponse[ProjectInspect], error) {
	var resp RestResponse[ProjectInspect]
	err := s.client.get(ctx, routeInspectProject, params, &resp)
	return resp, err
}

// List returns all projects.
func (s *ProjectsService) List(ctx context.Context) (RestResponse[[]ProjectInfo], error) {
	var resp RestResponse[[]ProjectInfo]
	err := s.client.get(ctx, routeListProjects, nil, &resp)
	return resp, err
}

// ListWithServices returns all projects along with their services.
func (s *ProjectsService) ListWithServices(ctx context.Context) (RestResponse[ProjectsWithServices], error) {
	var resp RestResponse[ProjectsWithServices]
	err := s.client.get(ctx, routeListProjectsAndServices, nil, &resp)
	return resp, err
}
