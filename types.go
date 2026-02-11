package easypanel

// RestResponse is the generic API response wrapper matching Easypanel's tRPC format.
type RestResponse[T any] struct {
	Result struct {
		Data struct {
			JSON T `json:"json"`
		} `json:"data"`
	} `json:"result"`
}

// Error represents an API error response.
type Error struct {
	OK           bool   `json:"ok"`
	ErrorMessage string `json:"errorMessage"`
	StatusCode   int    `json:"status,omitempty"`
}

func (e *Error) Error() string {
	if e.ErrorMessage != "" {
		return e.ErrorMessage
	}
	return "easypanel: unknown error"
}

// ServiceType represents the type of service in Easypanel.
type ServiceType string

const (
	ServiceTypeApp      ServiceType = "app"
	ServiceTypeMySQL    ServiceType = "mysql"
	ServiceTypeMariaDB  ServiceType = "mariadb"
	ServiceTypePostgres ServiceType = "postgres"
	ServiceTypeMongo    ServiceType = "mongo"
	ServiceTypeRedis    ServiceType = "redis"
	ServiceTypeCompose  ServiceType = "compose"
)

// LicenseType represents the license provider type.
type LicenseType string

const (
	LicenseTypeLemon  LicenseType = "lemon"
	LicenseTypePortal LicenseType = "portal"
)

// --- Auth Types ---

// User represents an Easypanel user.
type User struct {
	ID        string `json:"id"`
	CreatedAt string `json:"createdAt"`
	Email     string `json:"email"`
	Admin     bool   `json:"admin"`
}

// --- Project Types ---

// ProjectName is the parameter for creating/destroying a project.
type ProjectName struct {
	Name string `json:"name"`
}

// ProjectQuery is the parameter for inspecting a project.
type ProjectQuery struct {
	ProjectName string `json:"projectName"`
}

// ProjectInfo represents basic project information.
type ProjectInfo struct {
	Name      string `json:"name"`
	CreatedAt string `json:"createdAt"`
}

// ProjectInspect represents a project with its services.
type ProjectInspect struct {
	Project  ProjectInfo `json:"project"`
	Services []Service   `json:"services"`
}

// ProjectsWithServices represents the response for listing projects with services.
type ProjectsWithServices struct {
	Projects []ProjectInfo `json:"projects"`
	Services []Service     `json:"services"`
}

// --- Service Types ---

// SelectService identifies a specific service within a project.
type SelectService struct {
	ProjectName  string `json:"projectName"`
	ServiceName  string `json:"serviceName"`
	Password     string `json:"password,omitempty"`
	RootPassword string `json:"rootPassword,omitempty"`
	Image        string `json:"image,omitempty"`
}

// CreateServiceParams contains parameters for creating a service.
type CreateServiceParams struct {
	SelectService
	Domains []DomainParams `json:"domains,omitempty"`
}

// DomainParams represents domain configuration for a service.
type DomainParams struct {
	Host  string `json:"host"`
	HTTPS bool   `json:"https,omitempty"`
	Port  int    `json:"port,omitempty"`
	Path  string `json:"path,omitempty"`
}

// RedirectParams represents a redirect rule.
type RedirectParams struct {
	Regex       string `json:"regex"`
	Replacement string `json:"replacement"`
	Permanent   bool   `json:"permanent"`
}

// PortParams represents a port mapping.
type PortParams struct {
	Protocol  string `json:"protocol"` // "tcp" or "udp"
	Published int    `json:"published"`
	Target    int    `json:"target"`
}

// MountEntry represents a single mount configuration.
type MountEntry struct {
	Type      string `json:"type"` // "bind", "volume", "file"
	HostPath  string `json:"hostPath,omitempty"`
	Name      string `json:"name,omitempty"`
	Content   string `json:"content,omitempty"`
	MountPath string `json:"mountPath"`
}

// MountParams contains parameters for updating service mounts.
type MountParams struct {
	SelectService
	Mounts []MountEntry `json:"mounts"`
}

// UserParams represents basic auth credentials.
type UserParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// DeployParams contains parameters for deployment configuration.
type DeployParams struct {
	SelectService
	Replicas     int      `json:"replicas"`
	Command      []string `json:"command"`
	ZeroDowntime bool     `json:"zeroDowntime"`
	CapAdd       []string `json:"capAdd"`
	CapDrop      []string `json:"capDrop"`
	Sysctls      []string `json:"sysctls"`
}

// Resources represents resource limits and reservations.
type Resources struct {
	CPULimit          float64 `json:"cpuLimit"`
	CPUReservation    float64 `json:"cpuReservation"`
	MemoryLimit       float64 `json:"memoryLimit"`
	MemoryReservation float64 `json:"memoryReservation"`
}

// UpdateResources contains parameters for updating service resources.
type UpdateResources struct {
	SelectService
	Resources Resources `json:"resources"`
}

// UpdateBuildParams contains parameters for updating build configuration.
type UpdateBuildParams struct {
	SelectService
	BuildType string `json:"type,omitempty"` // "nixpacks", "herokuBuildpacks", "dockerfile", "none"
}

// GitParams represents Git source configuration.
type GitParams struct {
	Repo   string `json:"repo"`
	Branch string `json:"branch"`
	Path   string `json:"path,omitempty"`
}

// GithubParams represents GitHub source configuration.
type GithubParams struct {
	Owner  string `json:"owner"`
	Repo   string `json:"repo"`
	Branch string `json:"branch"`
	Path   string `json:"path,omitempty"`
}

// DockerImageParams represents Docker image source configuration.
type DockerImageParams struct {
	Image    string `json:"image"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

// ServiceSource represents the source configuration of a service.
type ServiceSource struct {
	AutoDeploy bool `json:"autoDeploy"`
	GitParams
	GithubParams
	DockerImageParams
}

// UpdateGithub contains parameters for updating GitHub source.
type UpdateGithub struct {
	SelectService
	GithubParams
	AutoDeploy bool `json:"autoDeploy"`
}

// UpdateGit contains parameters for updating Git source.
type UpdateGit struct {
	SelectService
	GitParams
	AutoDeploy bool `json:"autoDeploy"`
}

// UpdateImage contains parameters for updating Docker image source.
type UpdateImage struct {
	ProjectName string `json:"projectName"`
	ServiceName string `json:"serviceName"`
	Image       string `json:"image"`
	Username    string `json:"username,omitempty"`
	Password    string `json:"password,omitempty"`
}

// UpdateEnv contains parameters for updating environment variables.
type UpdateEnv struct {
	SelectService
	Env string `json:"env"`
}

// UpdateRedirects contains parameters for updating redirect rules.
type UpdateRedirects struct {
	SelectService
	Redirects []RedirectParams `json:"redirects"`
}

// UpdateBasicAuth contains parameters for updating basic auth.
type UpdateBasicAuth struct {
	SelectService
	BasicAuth []UserParams `json:"basicAuth"`
}

// UpdatePorts contains parameters for updating port mappings.
type UpdatePorts struct {
	SelectService
	Ports []PortParams `json:"ports"`
}

// ExposeServiceParams contains parameters for exposing a service port.
type ExposeServiceParams struct {
	SelectService
	ExposedPort int `json:"exposedPort"`
}

// UpdateAdvancedParams contains parameters for advanced service settings.
type UpdateAdvancedParams struct {
	SelectService
	Hostname string `json:"hostname,omitempty"`
}

// UpdateBackupParams contains parameters for backup configuration.
type UpdateBackupParams struct {
	SelectService
	// Backup-specific fields can be extended as needed
}

// Service represents a full service configuration.
type Service struct {
	SelectService
	Name          string           `json:"name,omitempty"` // Some endpoints return "name" instead of "serviceName"
	Type          ServiceType      `json:"type"`
	Enabled       bool             `json:"enabled"`
	Token         string           `json:"token"`
	Env           string           `json:"env,omitempty"`
	Command       string           `json:"command,omitempty"`
	Deploy        *DeployParams    `json:"deploy,omitempty"`
	Domains       []DomainParams   `json:"domains,omitempty"`
	Mounts        []MountEntry     `json:"mounts,omitempty"`
	Ports         []PortParams     `json:"ports,omitempty"`
	Redirects     []RedirectParams `json:"redirects,omitempty"`
	BasicAuth     []UserParams     `json:"basicAuth,omitempty"`
	ExposedPort   int              `json:"exposedPort,omitempty"`
	DeploymentURL string           `json:"deploymentUrl,omitempty"`
	Source        *ServiceSource   `json:"source,omitempty"`
	Resources     Resources        `json:"resources"`
}

// --- Monitor Types ---

// TimeValue represents a time-series data point.
type TimeValue struct {
	Value string `json:"value"`
	Time  string `json:"time"`
}

// NetworkValue represents network I/O at a point in time.
type NetworkValue struct {
	Input  int `json:"input"`
	Output int `json:"output"`
}

// NetworkTimeValue represents a network data point with timestamp.
type NetworkTimeValue struct {
	Value NetworkValue `json:"value"`
	Time  string       `json:"time"`
}

// AdvancedStats represents advanced monitoring statistics.
type AdvancedStats struct {
	CPU     []TimeValue        `json:"cpu"`
	Disk    []TimeValue        `json:"disk"`
	Memory  []TimeValue        `json:"memory"`
	Network []NetworkTimeValue `json:"network"`
}

// MemInfo represents memory information.
type MemInfo struct {
	TotalMemMb        float64 `json:"totalMemMb"`
	UsedMemMb         float64 `json:"usedMemMb"`
	FreeMemMb         float64 `json:"freeMemMb"`
	UsedMemPercentage float64 `json:"usedMemPercentage"`
	FreeMemPercentage float64 `json:"freeMemPercentage"`
}

// DiskInfo represents disk usage information.
type DiskInfo struct {
	TotalGb        string `json:"totalGb"`
	UsedGb         string `json:"usedGb"`
	FreeGb         string `json:"freeGb"`
	UsedPercentage string `json:"usedPercentage"`
	FreePercentage string `json:"freePercentage"`
}

// CPUInfo represents CPU information.
type CPUInfo struct {
	UsedPercentage float64   `json:"usedPercentage"`
	Count          int       `json:"count"`
	Loadavg        []float64 `json:"loadavg"`
}

// NetworkInfo represents network traffic information.
type NetworkInfo struct {
	InputMb  float64 `json:"inputMb"`
	OutputMb float64 `json:"outputMb"`
}

// SystemStats represents system-wide statistics.
type SystemStats struct {
	Uptime   float64     `json:"uptime"`
	MemInfo  MemInfo     `json:"memInfo"`
	DiskInfo DiskInfo    `json:"diskInfo"`
	CPUInfo  CPUInfo     `json:"cpuInfo"`
	Network  NetworkInfo `json:"network"`
}

// TaskStatus represents the status of a Docker task.
type TaskStatus struct {
	Actual  int `json:"actual"`
	Desired int `json:"desired"`
}

// DockerTaskStats maps service names to their task statuses.
type DockerTaskStats map[string]TaskStatus

// ContainerStat represents the stats for a single container.
type ContainerStat struct {
	CPU struct {
		Percent float64 `json:"percent"`
	} `json:"cpu"`
	Memory struct {
		Usage   int     `json:"usage"`
		Percent float64 `json:"percent"`
	} `json:"memory"`
	Network struct {
		In  int `json:"in"`
		Out int `json:"out"`
	} `json:"network"`
}

// ContainerStats represents statistics for a container.
type ContainerStats struct {
	ID            string        `json:"id"`
	Stats         ContainerStat `json:"stats"`
	ProjectName   string        `json:"projectName"`
	ServiceName   string        `json:"serviceName"`
	ContainerName string        `json:"containerName"`
}

// --- Settings Types ---

// ChangeCredentialsParams contains parameters for changing user credentials.
type ChangeCredentialsParams struct {
	Email       string `json:"email"`
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
}

// PruneDockerDailyParams contains parameters for Docker prune daily setting.
type PruneDockerDailyParams struct {
	PruneDockerDaily bool `json:"pruneDockerDaily"`
}

// GithubTokenParams contains parameters for setting GitHub token.
type GithubTokenParams struct {
	GithubToken string `json:"githubToken"`
}

// PanelDomainParams contains parameters for setting panel domain.
type PanelDomainParams struct {
	ServeOnIP          bool   `json:"serveOnIp"`
	DefaultPanelDomain string `json:"defaultPanelDomain"`
	PanelDomain        string `json:"panelDomain"`
}

// PanelDomain represents the panel domain configuration response.
type PanelDomain struct {
	ServeOnIP          bool   `json:"serveOnIp"`
	PanelDomain        string `json:"panelDomain"`
	DefaultPanelDomain string `json:"defaultPanelDomain"`
}

// TraefikConfParams contains parameters for Traefik configuration.
type TraefikConfParams struct {
	Config string `json:"config"`
}

// LetsEncryptParams contains parameters for Let's Encrypt email.
type LetsEncryptParams struct {
	LetsEncryptEmail string `json:"letsEncryptEmail"`
}

// --- Domain Types (newer Easypanel API) ---

// ServiceDestination represents the target service for a domain.
type ServiceDestination struct {
	Protocol       string `json:"protocol"`
	Port           int    `json:"port"`
	Path           string `json:"path"`
	ProjectName    string `json:"projectName"`
	ServiceName    string `json:"serviceName"`
	ComposeService string `json:"composeService,omitempty"`
}

// Domain represents a domain configuration in the newer Easypanel API.
type Domain struct {
	ID                  string              `json:"id"`
	HTTPS               bool                `json:"https"`
	Host                string              `json:"host"`
	Path                string              `json:"path"`
	Middlewares         []string            `json:"middlewares"`
	CertificateResolver string             `json:"certificateResolver"`
	Wildcard            bool                `json:"wildcard"`
	DestinationType     string              `json:"destinationType"`
	ServiceDestination  *ServiceDestination `json:"serviceDestination,omitempty"`
}

// CreateDomainParams contains parameters for creating a domain.
type CreateDomainParams = Domain

// UpdateDomainParams contains parameters for updating a domain.
type UpdateDomainParams = Domain

// DeleteDomainParams contains parameters for deleting a domain.
type DeleteDomainParams struct {
	ID string `json:"id"`
}

// ListDomainsParams contains parameters for listing domains.
type ListDomainsParams struct {
	ProjectName string `json:"projectName"`
	ServiceName string `json:"serviceName"`
}

// --- Compose Service Types ---

// UpdateSourceInline contains parameters for updating a compose service with inline content.
type UpdateSourceInline struct {
	ProjectName    string `json:"projectName"`
	ServiceName    string `json:"serviceName"`
	ComposeFile    string `json:"composeFile"`
	ComposeContent string `json:"composeContent"`
}

// UpdateSourceGitCompose contains parameters for updating a compose service with a Git source.
type UpdateSourceGitCompose struct {
	ProjectName string `json:"projectName"`
	ServiceName string `json:"serviceName"`
	Repo        string `json:"repo"`
	Ref         string `json:"ref"`
	RootPath    string `json:"rootPath,omitempty"`
	ComposeFile string `json:"composeFile,omitempty"`
	AutoDeploy  bool   `json:"autoDeploy"`
}

// --- Action Types ---

// Action represents a deployment action.
type Action struct {
	ID          string `json:"id"`
	Type        string `json:"type"`
	Status      string `json:"status"` // "success", "error", "running"
	ProjectName string `json:"projectName"`
	ServiceName string `json:"serviceName"`
	ServiceType string `json:"serviceType"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

// ActionDetail represents detailed information about an action including logs.
type ActionDetail struct {
	Action
	Log string `json:"log"`
}

// ListActionsParams contains parameters for listing actions.
type ListActionsParams struct {
	ProjectName string `json:"projectName"`
	ServiceName string `json:"serviceName"`
}

// GetActionParams contains parameters for getting a single action.
type GetActionParams struct {
	ActionID string `json:"actionId"`
}
