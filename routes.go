package easypanel

import "strings"

// API route constants for Easypanel tRPC endpoints.
const (
	// Auth routes
	routeGetUser = "/api/trpc/auth.getUser"
	routeLogout  = "/api/trpc/auth.logout"
	routeLogin   = "/api/trpc/auth.login"

	// Project routes
	routeListProjects            = "/api/trpc/projects.listProjects"
	routeListProjectsAndServices = "/api/trpc/projects.listProjectsAndServices"
	routeCanCreateProject        = "/api/trpc/projects.canCreateProject"
	routeInspectProject          = "/api/trpc/projects.inspectProject"
	routeCreateProject           = "/api/trpc/projects.createProject"
	routeDestroyProject          = "/api/trpc/projects.destroyProject"

	// Monitor routes
	routeGetAdvancedStats    = "/api/trpc/monitor.getAdvancedStats"
	routeGetSystemStats      = "/api/trpc/monitor.getSystemStats"
	routeGetDockerTaskStats  = "/api/trpc/monitor.getDockerTaskStats"
	routeGetMonitorTableData = "/api/trpc/monitor.getMonitorTableData"
	routeGetServiceStats     = "/api/trpc/monitor.getServiceStats"

	// Settings routes
	routeRestartEasypanel         = "/api/trpc/settings.restartEasypanel"
	routeGetServerIp              = "/api/trpc/settings.getServerIp"
	routeRefreshServerIp          = "/api/trpc/settings.refreshServerIp"
	routeGetGithubToken           = "/api/trpc/settings.getGithubToken"
	routeSetGithubToken           = "/api/trpc/settings.setGithubToken"
	routeGetPanelDomain           = "/api/trpc/settings.getPanelDomain"
	routeSetPanelDomain           = "/api/trpc/settings.setPanelDomain"
	routeGetLetsEncryptEmail      = "/api/trpc/settings.getLetsEncryptEmail"
	routeSetLetsEncryptEmail      = "/api/trpc/settings.setLetsEncryptEmail"
	routeGetTraefikCustomConfig   = "/api/trpc/settings.getTraefikCustomConfig"
	routeUpdateTraefikCustomConfig = "/api/trpc/settings.updateTraefikCustomConfig"
	routeRestartTraefik           = "/api/trpc/settings.restartTraefik"
	routePruneDockerImages        = "/api/trpc/settings.pruneDockerImages"
	routePruneDockerBuilder       = "/api/trpc/settings.pruneDockerBuilder"
	routeSetPruneDockerDaily      = "/api/trpc/settings.setPruneDockerDaily"
	routeChangeCredentials        = "/api/trpc/settings.changeCredentials"

	// Domain routes
	routeCreateDomain = "/api/trpc/domains.createDomain"
	routeUpdateDomain = "/api/trpc/domains.updateDomain"
	routeDeleteDomain = "/api/trpc/domains.deleteDomain"
	routeListDomains  = "/api/trpc/domains.listDomains"

	// Log routes
	routeGetServiceLogs = "/api/trpc/logs.getServiceLogs"

	// Action routes
	routeListActions = "/api/trpc/actions.listActions"
	routeGetAction   = "/api/trpc/actions.getAction"
)

// Service route templates (use serviceRoute to interpolate type).
const (
	routeCreateService       = "/api/trpc/services.{type}.createService"
	routeInspectService      = "/api/trpc/services.{type}.inspectService"
	routeDestroyService      = "/api/trpc/services.{type}.destroyService"
	routeDeployService       = "/api/trpc/services.{type}.deployService"
	routeStopService         = "/api/trpc/services.{type}.stopService"
	routeRestartService      = "/api/trpc/services.{type}.restartService"
	routeDisableService      = "/api/trpc/services.{type}.disableService"
	routeEnableService       = "/api/trpc/services.{type}.enableService"
	routeExposeService       = "/api/trpc/services.{type}.exposeService"
	routeRefreshDeployToken  = "/api/trpc/services.{type}.refreshDeployToken"
	routeUpdateSourceGithub  = "/api/trpc/services.{type}.updateSourceGithub"
	routeUpdateSourceGit     = "/api/trpc/services.{type}.updateSourceGit"
	routeUpdateSourceImage   = "/api/trpc/services.{type}.updateSourceImage"
	routeUpdateBuild         = "/api/trpc/services.{type}.updateBuild"
	routeUpdateEnv           = "/api/trpc/services.{type}.updateEnv"
	routeUpdateDomains       = "/api/trpc/services.{type}.updateDomains"
	routeUpdateRedirects     = "/api/trpc/services.{type}.updateRedirects"
	routeUpdateBasicAuth     = "/api/trpc/services.{type}.updateBasicAuth"
	routeUpdateMounts        = "/api/trpc/services.{type}.updateMounts"
	routeUpdatePorts         = "/api/trpc/services.{type}.updatePorts"
	routeUpdateResources     = "/api/trpc/services.{type}.updateResources"
	routeUpdateDeploy        = "/api/trpc/services.{type}.updateDeploy"
	routeUpdateBackup        = "/api/trpc/services.{type}.updateBackup"
	routeUpdateAdvanced      = "/api/trpc/services.{type}.updateAdvanced"
	routeUpdateSourceInline  = "/api/trpc/services.{type}.updateSourceInline"
)

// License route templates.
const (
	routeGetLicensePayload = "/api/trpc/{type}License.getLicensePayload"
	routeActivateLicense   = "/api/trpc/{type}License.activate"
)

// serviceRoute replaces {type} in a route template with the given service type.
func serviceRoute(template string, st ServiceType) string {
	return strings.Replace(template, "{type}", string(st), 1)
}

// licenseRoute replaces {type} in a license route template with the given license type.
func licenseRoute(template string, lt LicenseType) string {
	return strings.Replace(template, "{type}", string(lt), 1)
}
