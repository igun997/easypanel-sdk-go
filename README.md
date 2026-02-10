# easypanel-sdk-go

Go SDK for the [Easypanel](https://easypanel.io) API. Provides a fully typed, idiomatic Go client for managing projects, services, domains, monitoring, settings, and deployment actions.

## Features

- Full coverage of the Easypanel tRPC API
- Zero runtime dependencies (only `net/http` from stdlib)
- Generic `RestResponse[T]` for type-safe responses
- Support for all service types: `app`, `mysql`, `mariadb`, `postgres`, `mongo`, `redis`, `compose`
- Domain management (create, update, delete, list)
- Deployment action tracking
- Automatic retry on 5xx errors

## Installation

```bash
go get github.com/igun997/easypanel-sdk-go
```

Requires Go 1.23+.

## Quick Start

```go
package main

import (
    "context"
    "fmt"
    "log"

    easypanel "github.com/igun997/easypanel-sdk-go"
)

func main() {
    client := easypanel.New(easypanel.Config{
        Endpoint: "https://your-panel.example.com",
        Token:    "your-api-token",
    })
    ctx := context.Background()

    // List all projects
    projects, err := client.Projects.List(ctx)
    if err != nil {
        log.Fatal(err)
    }
    for _, p := range projects.Result.Data.JSON {
        fmt.Printf("Project: %s\n", p.Name)
    }
}
```

## Usage Examples

### Create a Project

```go
resp, err := client.Projects.Create(ctx, easypanel.ProjectName{
    Name: "my-project",
})
```

### Create and Deploy an App Service

```go
// Create service
_, err := client.Services.Create(ctx, easypanel.ServiceTypeApp, easypanel.CreateServiceParams{
    SelectService: easypanel.SelectService{
        ProjectName: "my-project",
        ServiceName: "web",
    },
})

// Set Docker image
err = client.Services.UpdateSourceImage(ctx, easypanel.ServiceTypeApp, easypanel.UpdateImage{
    ProjectName: "my-project",
    ServiceName: "web",
    Image:       "nginx:latest",
})

// Set environment variables
err = client.Services.UpdateEnv(ctx, easypanel.ServiceTypeApp, easypanel.UpdateEnv{
    SelectService: easypanel.SelectService{
        ProjectName: "my-project",
        ServiceName: "web",
    },
    Env: "PORT=8080\nNODE_ENV=production",
})

// Deploy
err = client.Services.Deploy(ctx, easypanel.ServiceTypeApp, easypanel.SelectService{
    ProjectName: "my-project",
    ServiceName: "web",
})
```

### Manage Domains

```go
// Create a domain
_, err := client.Domains.Create(ctx, easypanel.Domain{
    ID:                  "unique-domain-id",
    HTTPS:               true,
    Host:                "app.example.com",
    Path:                "/",
    Middlewares:         []string{},
    CertificateResolver: "letsencrypt",
    DestinationType:     "service",
    ServiceDestination: &easypanel.ServiceDestination{
        Protocol:    "http",
        Port:        80,
        Path:        "/",
        ProjectName: "my-project",
        ServiceName: "web",
    },
})

// List domains for a service
domains, err := client.Domains.List(ctx, easypanel.ListDomainsParams{
    ProjectName: "my-project",
    ServiceName: "web",
})
```

### Deploy a Compose Service

```go
// Create a compose service
_, err := client.Services.Create(ctx, easypanel.ServiceTypeCompose, easypanel.CreateServiceParams{
    SelectService: easypanel.SelectService{
        ProjectName: "my-project",
        ServiceName: "stack",
    },
})

// Set inline docker-compose content
err = client.Services.UpdateSourceInline(ctx, easypanel.ServiceTypeCompose, easypanel.UpdateSourceInline{
    ProjectName:    "my-project",
    ServiceName:    "stack",
    ComposeFile:    "docker-compose.yml",
    ComposeContent: "version: '3'\nservices:\n  web:\n    image: nginx:latest\n    ports:\n      - '80:80'\n",
})

// Or use a Git source
err = client.Services.UpdateSourceGitCompose(ctx, easypanel.ServiceTypeCompose, easypanel.UpdateSourceGitCompose{
    ProjectName: "my-project",
    ServiceName: "stack",
    Repo:        "https://github.com/user/repo",
    Ref:         "main",
    ComposeFile: "docker-compose.yml",
    AutoDeploy:  true,
})
```

### Track Deployment Actions

```go
actions, err := client.Actions.List(ctx, easypanel.ListActionsParams{
    ProjectName: "my-project",
    ServiceName: "web",
})
for _, a := range actions.Result.Data.JSON {
    fmt.Printf("Action: %s status=%s\n", a.Type, a.Status)
}

// Get details for a specific action
detail, err := client.Actions.Get(ctx, easypanel.GetActionParams{
    ActionID: "action-id",
})
fmt.Printf("Log: %s\n", detail.Result.Data.JSON.Log)
```

### Monitoring

```go
// System stats
stats, err := client.Monitor.GetSystemStats(ctx)
fmt.Printf("CPU: %.1f%%, Memory: %.0f MB\n",
    stats.Result.Data.JSON.CPUInfo.UsedPercentage,
    stats.Result.Data.JSON.MemInfo.UsedMemMb)

// Docker task status
tasks, err := client.Monitor.GetDockerTaskStats(ctx)

// Container-level metrics
containers, err := client.Monitor.GetMonitorTableData(ctx)
```

### Settings

```go
// Get server IP
ip, err := client.Settings.GetServerIp(ctx)

// Set Let's Encrypt email
_, err = client.Settings.SetLetsEncryptEmail(ctx, easypanel.LetsEncryptParams{
    Email: "admin@example.com",
})

// Get panel domain
domain, err := client.Settings.GetPanelDomain(ctx)
```

## API Reference

### Client

| Method | Description |
|--------|-------------|
| `New(Config)` | Create a new client |
| `GetUser(ctx)` | Get current user info |
| `GetLicensePayload(ctx)` | Get license information |
| `ActivateLicense(ctx, params)` | Activate a license |

### Projects

| Method | Description |
|--------|-------------|
| `Projects.CanCreate(ctx)` | Check if project creation is allowed |
| `Projects.Create(ctx, params)` | Create a new project |
| `Projects.Destroy(ctx, params)` | Delete a project |
| `Projects.Inspect(ctx, params)` | Get project details with services |
| `Projects.List(ctx)` | List all projects |
| `Projects.ListWithServices(ctx)` | List projects with their services |

### Services

| Method | Description |
|--------|-------------|
| `Services.Create(ctx, type, params)` | Create a service |
| `Services.Inspect(ctx, type, params)` | Inspect a service |
| `Services.Destroy(ctx, type, params)` | Delete a service |
| `Services.Deploy(ctx, type, params)` | Trigger deployment |
| `Services.Disable(ctx, type, params)` | Disable a service |
| `Services.Enable(ctx, type, params)` | Enable a service |
| `Services.ExposeService(ctx, type, params)` | Expose a port |
| `Services.UpdateSourceImage(ctx, type, params)` | Set Docker image source |
| `Services.UpdateSourceGithub(ctx, type, params)` | Set GitHub source |
| `Services.UpdateSourceGit(ctx, type, params)` | Set Git source |
| `Services.UpdateSourceInline(ctx, type, params)` | Set inline compose content |
| `Services.UpdateSourceGitCompose(ctx, type, params)` | Set Git source for compose |
| `Services.UpdateEnv(ctx, type, params)` | Update environment variables |
| `Services.UpdateBuild(ctx, type, params)` | Update build config |
| `Services.UpdateDomains(ctx, type, params)` | Update domains (legacy) |
| `Services.UpdateRedirects(ctx, type, params)` | Update redirects |
| `Services.UpdateBasicAuth(ctx, type, params)` | Update basic auth |
| `Services.UpdateMounts(ctx, type, params)` | Update mount points |
| `Services.UpdatePorts(ctx, type, params)` | Update port mappings |
| `Services.UpdateResources(ctx, type, params)` | Update resource limits |
| `Services.UpdateDeploy(ctx, type, params)` | Update deploy config |
| `Services.UpdateBackup(ctx, type, params)` | Update backup config |
| `Services.UpdateAdvanced(ctx, type, params)` | Update advanced settings |
| `Services.GetServiceLogs(ctx, params)` | Get service logs |

### Domains

| Method | Description |
|--------|-------------|
| `Domains.Create(ctx, params)` | Create a domain |
| `Domains.Update(ctx, params)` | Update a domain |
| `Domains.Delete(ctx, params)` | Delete a domain |
| `Domains.List(ctx, params)` | List domains for a service |

### Actions

| Method | Description |
|--------|-------------|
| `Actions.List(ctx, params)` | List deployment actions |
| `Actions.Get(ctx, params)` | Get action details with logs |

### Monitor

| Method | Description |
|--------|-------------|
| `Monitor.GetAdvancedStats(ctx)` | CPU, disk, memory, network over time |
| `Monitor.GetDockerTaskStats(ctx)` | Docker task status per service |
| `Monitor.GetMonitorTableData(ctx)` | Container-level statistics |
| `Monitor.GetSystemStats(ctx)` | System-wide stats |

### Settings

| Method | Description |
|--------|-------------|
| `Settings.ChangeCredentials(ctx, params)` | Change email/password |
| `Settings.GetGithubToken(ctx)` | Get GitHub token |
| `Settings.SetGithubToken(ctx, params)` | Set GitHub token |
| `Settings.GetLetsEncryptEmail(ctx)` | Get LE email |
| `Settings.SetLetsEncryptEmail(ctx, params)` | Set LE email |
| `Settings.GetPanelDomain(ctx)` | Get panel domain config |
| `Settings.SetPanelDomain(ctx, params)` | Set panel domain |
| `Settings.GetServerIp(ctx)` | Get server IP |
| `Settings.RefreshServerIp(ctx)` | Refresh server IP |
| `Settings.GetTraefikCustomConfig(ctx)` | Get Traefik config |
| `Settings.UpdateTraefikCustomConfig(ctx, params)` | Update Traefik config |
| `Settings.PruneDockerBuilder(ctx)` | Prune builder cache |
| `Settings.PruneDockerImages(ctx)` | Prune unused images |
| `Settings.SetDockerPruneDaily(ctx, params)` | Toggle daily prune |
| `Settings.RestartEasypanel(ctx)` | Restart Easypanel |
| `Settings.RestartTraefik(ctx)` | Restart Traefik |

## Service Types

| Constant | Value |
|----------|-------|
| `ServiceTypeApp` | `app` |
| `ServiceTypeMySQL` | `mysql` |
| `ServiceTypeMariaDB` | `mariadb` |
| `ServiceTypePostgres` | `postgres` |
| `ServiceTypeMongo` | `mongo` |
| `ServiceTypeRedis` | `redis` |
| `ServiceTypeCompose` | `compose` |

## Testing

```bash
# Run unit tests
go test -v -run 'Test[^I]' ./...

# Run integration tests (requires EASYPANEL_ENDPOINT and EASYPANEL_TOKEN)
go test -v -run 'TestIntegration' ./...
```

## License

MIT
