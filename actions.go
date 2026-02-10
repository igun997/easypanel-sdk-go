package easypanel

import "context"

// ActionsService handles action/deployment tracking API operations.
type ActionsService struct {
	client *httpClient
}

// List returns all actions for a project/service.
func (s *ActionsService) List(ctx context.Context, params ListActionsParams) (RestResponse[[]Action], error) {
	var resp RestResponse[[]Action]
	err := s.client.get(ctx, routeListActions, params, &resp)
	return resp, err
}

// Get returns detailed information about a specific action.
func (s *ActionsService) Get(ctx context.Context, params GetActionParams) (RestResponse[ActionDetail], error) {
	var resp RestResponse[ActionDetail]
	err := s.client.get(ctx, routeGetAction, params, &resp)
	return resp, err
}
