package easypanel

import "context"

// DomainsService handles domain-related API operations (newer Easypanel API).
type DomainsService struct {
	client *httpClient
}

// Create creates a new domain.
func (s *DomainsService) Create(ctx context.Context, params CreateDomainParams) (RestResponse[Domain], error) {
	var resp RestResponse[Domain]
	err := s.client.post(ctx, routeCreateDomain, params, &resp)
	return resp, err
}

// Update updates an existing domain.
func (s *DomainsService) Update(ctx context.Context, params UpdateDomainParams) error {
	return s.client.post(ctx, routeUpdateDomain, params, nil)
}

// Delete deletes a domain by ID.
func (s *DomainsService) Delete(ctx context.Context, params DeleteDomainParams) error {
	return s.client.post(ctx, routeDeleteDomain, params, nil)
}

// List returns all domains for a project/service.
func (s *DomainsService) List(ctx context.Context, params ListDomainsParams) (RestResponse[[]Domain], error) {
	var resp RestResponse[[]Domain]
	err := s.client.get(ctx, routeListDomains, params, &resp)
	return resp, err
}
