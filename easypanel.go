package easypanel

import "context"

// Config holds the configuration for an Easypanel client.
type Config struct {
	Endpoint string // Base URL, e.g. "https://panel.example.com"
	Token    string // Authorization token
}

// Client is the main entry point for the Easypanel SDK.
type Client struct {
	Projects *ProjectsService
	Services *ServicesService
	Monitor  *MonitorService
	Settings *SettingsService
	Domains  *DomainsService
	Actions  *ActionsService

	client *httpClient
}

// New creates a new Easypanel client with the given configuration.
func New(cfg Config) *Client {
	c := newHTTPClient(cfg.Endpoint, cfg.Token)
	return &Client{
		Projects: &ProjectsService{client: c},
		Services: &ServicesService{client: c},
		Monitor:  &MonitorService{client: c},
		Settings: &SettingsService{client: c},
		Domains:  &DomainsService{client: c},
		Actions:  &ActionsService{client: c},
		client:   c,
	}
}

// GetUser returns the authenticated user's information.
func (c *Client) GetUser(ctx context.Context) (RestResponse[User], error) {
	var resp RestResponse[User]
	err := c.client.get(ctx, routeGetUser, nil, &resp)
	return resp, err
}

// GetLicensePayload retrieves the license payload for the given license type.
func (c *Client) GetLicensePayload(ctx context.Context, lt LicenseType) error {
	return c.client.get(ctx, licenseRoute(routeGetLicensePayload, lt), nil, nil)
}

// ActivateLicense activates the license for the given license type.
func (c *Client) ActivateLicense(ctx context.Context, lt LicenseType) error {
	return c.client.post(ctx, licenseRoute(routeActivateLicense, lt), nil, nil)
}
