// Package easypanel provides a Go SDK for the Easypanel API.
//
// Easypanel is a modern server control panel that uses Docker under the hood.
// This SDK wraps the Easypanel tRPC API and provides a typed, idiomatic Go client
// for managing projects, services, domains, monitoring, and settings.
//
// # Getting Started
//
// Create a client with your Easypanel endpoint and API token:
//
//	client := easypanel.New(easypanel.Config{
//	    Endpoint: "https://panel.example.com",
//	    Token:    "your-api-token",
//	})
//
// # Projects
//
// Projects are the top-level grouping for services:
//
//	// Create a project
//	resp, err := client.Projects.Create(ctx, easypanel.ProjectName{Name: "my-app"})
//
//	// List all projects
//	list, err := client.Projects.List(ctx)
//
//	// Inspect a project with its services
//	info, err := client.Projects.Inspect(ctx, easypanel.ProjectQuery{ProjectName: "my-app"})
//
// # Services
//
// Services run inside projects. Each service has a type that determines its behavior.
// Supported types: app, mysql, mariadb, postgres, mongo, redis, compose.
//
//	// Create an app service
//	_, err := client.Services.Create(ctx, easypanel.ServiceTypeApp, easypanel.CreateServiceParams{
//	    SelectService: easypanel.SelectService{
//	        ProjectName: "my-app",
//	        ServiceName: "api",
//	    },
//	})
//
//	// Set Docker image and deploy
//	err = client.Services.UpdateSourceImage(ctx, easypanel.ServiceTypeApp, easypanel.UpdateImage{
//	    ProjectName: "my-app",
//	    ServiceName: "api",
//	    Image:       "node:20-alpine",
//	})
//	err = client.Services.Deploy(ctx, easypanel.ServiceTypeApp, easypanel.SelectService{
//	    ProjectName: "my-app",
//	    ServiceName: "api",
//	})
//
// # Compose Services
//
// Compose services allow deploying multi-container stacks using docker-compose:
//
//	_, err := client.Services.Create(ctx, easypanel.ServiceTypeCompose, easypanel.CreateServiceParams{
//	    SelectService: easypanel.SelectService{
//	        ProjectName: "my-app",
//	        ServiceName: "stack",
//	    },
//	})
//	err = client.Services.UpdateSourceInline(ctx, easypanel.ServiceTypeCompose, easypanel.UpdateSourceInline{
//	    ProjectName:    "my-app",
//	    ServiceName:    "stack",
//	    ComposeFile:    "docker-compose.yml",
//	    ComposeContent: composeYAML,
//	})
//
// # Domains
//
// Domains are managed separately from services:
//
//	_, err := client.Domains.Create(ctx, easypanel.Domain{
//	    ID:                  "unique-id",
//	    HTTPS:               true,
//	    Host:                "api.example.com",
//	    Path:                "/",
//	    CertificateResolver: "letsencrypt",
//	    DestinationType:     "service",
//	    ServiceDestination: &easypanel.ServiceDestination{
//	        Protocol:    "http",
//	        Port:        80,
//	        Path:        "/",
//	        ProjectName: "my-app",
//	        ServiceName: "api",
//	    },
//	})
//
// # Actions
//
// Track deployment progress through actions:
//
//	actions, err := client.Actions.List(ctx, easypanel.ListActionsParams{
//	    ProjectName: "my-app",
//	    ServiceName: "api",
//	})
//
// # Monitoring
//
// Get system and container statistics:
//
//	stats, err := client.Monitor.GetSystemStats(ctx)
//	containers, err := client.Monitor.GetMonitorTableData(ctx)
//
// # Response Format
//
// All responses are wrapped in [RestResponse] which follows the tRPC envelope format.
// Access the data through the nested Result.Data.JSON field:
//
//	resp, err := client.Projects.List(ctx)
//	for _, p := range resp.Result.Data.JSON {
//	    fmt.Println(p.Name)
//	}
package easypanel
