package easypanel

import "context"

// MonitorService handles monitoring-related API operations.
type MonitorService struct {
	client *httpClient
}

// GetAdvancedStats returns advanced monitoring statistics (CPU, disk, memory, network over time).
func (s *MonitorService) GetAdvancedStats(ctx context.Context) (RestResponse[AdvancedStats], error) {
	var resp RestResponse[AdvancedStats]
	err := s.client.get(ctx, routeGetAdvancedStats, nil, &resp)
	return resp, err
}

// GetDockerTaskStats returns Docker task status for all services.
func (s *MonitorService) GetDockerTaskStats(ctx context.Context) (RestResponse[DockerTaskStats], error) {
	var resp RestResponse[DockerTaskStats]
	err := s.client.get(ctx, routeGetDockerTaskStats, nil, &resp)
	return resp, err
}

// GetMonitorTableData returns container-level statistics for all running services.
func (s *MonitorService) GetMonitorTableData(ctx context.Context) (RestResponse[[]ContainerStats], error) {
	var resp RestResponse[[]ContainerStats]
	err := s.client.get(ctx, routeGetMonitorTableData, nil, &resp)
	return resp, err
}

// GetSystemStats returns system-wide statistics (uptime, CPU, memory, disk, network).
func (s *MonitorService) GetSystemStats(ctx context.Context) (RestResponse[SystemStats], error) {
	var resp RestResponse[SystemStats]
	err := s.client.get(ctx, routeGetSystemStats, nil, &resp)
	return resp, err
}
