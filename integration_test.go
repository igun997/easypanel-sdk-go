package easypanel

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func newIntegrationClient() *Client {
	endpoint := os.Getenv("EASYPANEL_ENDPOINT")
	token := os.Getenv("EASYPANEL_TOKEN")
	if endpoint == "" || token == "" {
		panic("EASYPANEL_ENDPOINT and EASYPANEL_TOKEN must be set")
	}
	return New(Config{
		Endpoint: endpoint,
		Token:    token,
	})
}

func TestIntegration_GetUser(t *testing.T) {
	client := newIntegrationClient()
	resp, err := client.GetUser(context.Background())
	require.NoError(t, err)
	assert.NotEmpty(t, resp.Result.Data.JSON.ID)
	assert.NotEmpty(t, resp.Result.Data.JSON.Email)
	t.Logf("User: id=%s email=%s admin=%v", resp.Result.Data.JSON.ID, resp.Result.Data.JSON.Email, resp.Result.Data.JSON.Admin)
}

func TestIntegration_ProjectsList(t *testing.T) {
	client := newIntegrationClient()
	resp, err := client.Projects.List(context.Background())
	require.NoError(t, err)
	t.Logf("Found %d projects", len(resp.Result.Data.JSON))
	for _, p := range resp.Result.Data.JSON {
		t.Logf("  Project: %s (created: %s)", p.Name, p.CreatedAt)
	}
}

func TestIntegration_ProjectsListWithServices(t *testing.T) {
	client := newIntegrationClient()
	resp, err := client.Projects.ListWithServices(context.Background())
	require.NoError(t, err)
	t.Logf("Found %d projects, %d services", len(resp.Result.Data.JSON.Projects), len(resp.Result.Data.JSON.Services))
	for _, s := range resp.Result.Data.JSON.Services {
		t.Logf("  Service: %s/%s type=%s enabled=%v", s.ProjectName, s.ServiceName, s.Type, s.Enabled)
	}
}

func TestIntegration_MonitorGetSystemStats(t *testing.T) {
	client := newIntegrationClient()
	resp, err := client.Monitor.GetSystemStats(context.Background())
	require.NoError(t, err)
	stats := resp.Result.Data.JSON
	t.Logf("System Stats:")
	t.Logf("  Uptime: %.0f seconds", stats.Uptime)
	t.Logf("  CPU: %.1f%% (%d cores)", stats.CPUInfo.UsedPercentage, stats.CPUInfo.Count)
	t.Logf("  Memory: %.0f/%.0f MB (%.1f%%)", stats.MemInfo.UsedMemMb, stats.MemInfo.TotalMemMb, stats.MemInfo.UsedMemPercentage)
	t.Logf("  Disk: %s/%s GB (%s%%)", stats.DiskInfo.UsedGb, stats.DiskInfo.TotalGb, stats.DiskInfo.UsedPercentage)
	assert.Greater(t, stats.Uptime, 0.0)
	assert.Greater(t, stats.CPUInfo.Count, 0)
	assert.Greater(t, stats.MemInfo.TotalMemMb, 0.0)
}

func TestIntegration_MonitorGetDockerTaskStats(t *testing.T) {
	client := newIntegrationClient()
	resp, err := client.Monitor.GetDockerTaskStats(context.Background())
	require.NoError(t, err)
	t.Logf("Docker tasks: %d entries", len(resp.Result.Data.JSON))
	for name, status := range resp.Result.Data.JSON {
		t.Logf("  %s: actual=%d desired=%d", name, status.Actual, status.Desired)
	}
}

func TestIntegration_SettingsGetServerIp(t *testing.T) {
	client := newIntegrationClient()
	resp, err := client.Settings.GetServerIp(context.Background())
	require.NoError(t, err)
	assert.NotEmpty(t, resp.Result.Data.JSON)
	t.Logf("Server IP: %s", resp.Result.Data.JSON)
}

func TestIntegration_SettingsGetPanelDomain(t *testing.T) {
	client := newIntegrationClient()
	resp, err := client.Settings.GetPanelDomain(context.Background())
	require.NoError(t, err)
	t.Logf("Panel Domain: %s (default: %s, serveOnIp: %v)",
		resp.Result.Data.JSON.PanelDomain,
		resp.Result.Data.JSON.DefaultPanelDomain,
		resp.Result.Data.JSON.ServeOnIP)
}

func TestIntegration_ProjectsCanCreate(t *testing.T) {
	client := newIntegrationClient()
	resp, err := client.Projects.CanCreate(context.Background())
	require.NoError(t, err)
	t.Logf("Can create projects: %v", resp.Result.Data.JSON)
}

// TestIntegration_FullServiceLifecycle is an end-to-end test that creates a project,
// deploys zedtux/docker-coming-soon as an app service, exercises all service operations,
// and cleans up afterwards.
func TestIntegration_FullServiceLifecycle(t *testing.T) {
	client := newIntegrationClient()
	ctx := context.Background()

	const projectName = "sdk-test"
	const serviceName = "coming-soon"

	// --- Cleanup helper (runs even if test fails) ---
	t.Cleanup(func() {
		t.Log("=== CLEANUP ===")
		_ = client.Services.Destroy(ctx, ServiceTypeApp, SelectService{
			ProjectName: projectName,
			ServiceName: serviceName,
		})
		_ = client.Projects.Destroy(ctx, ProjectName{Name: projectName})
		t.Log("Cleanup complete")
	})

	// --- Step 1: Create project ---
	t.Log("=== Step 1: Create project ===")
	projResp, err := client.Projects.Create(ctx, ProjectName{Name: projectName})
	require.NoError(t, err, "create project")
	assert.Equal(t, projectName, projResp.Result.Data.JSON.Name)
	t.Logf("Created project: %s", projResp.Result.Data.JSON.Name)

	// --- Step 2: Verify project in list ---
	t.Log("=== Step 2: Verify project exists ===")
	listResp, err := client.Projects.List(ctx)
	require.NoError(t, err, "list projects")
	found := false
	for _, p := range listResp.Result.Data.JSON {
		if p.Name == projectName {
			found = true
			break
		}
	}
	assert.True(t, found, "project %q should appear in project list", projectName)

	// --- Step 3: Create app service ---
	t.Log("=== Step 3: Create app service ===")
	svcResp, err := client.Services.Create(ctx, ServiceTypeApp, CreateServiceParams{
		SelectService: SelectService{
			ProjectName: projectName,
			ServiceName: serviceName,
		},
	})
	require.NoError(t, err, "create service")
	// API may return "name" instead of "serviceName" in create response
	svcName := svcResp.Result.Data.JSON.ServiceName
	if svcName == "" {
		svcName = svcResp.Result.Data.JSON.Name
	}
	t.Logf("Created service: %s/%s type=%s", svcResp.Result.Data.JSON.ProjectName, svcName, svcResp.Result.Data.JSON.Type)

	// --- Step 4: Update source to Docker image ---
	t.Log("=== Step 4: Set Docker image source ===")
	err = client.Services.UpdateSourceImage(ctx, ServiceTypeApp, UpdateImage{
		ProjectName: projectName,
		ServiceName: serviceName,
		Image:       "zedtux/docker-coming-soon",
	})
	require.NoError(t, err, "update source image")
	t.Log("Source set to zedtux/docker-coming-soon")

	// --- Step 5: Update environment variables ---
	t.Log("=== Step 5: Set environment variables ===")
	envVars := "TITLE=SDK Test App\nPRODUCT_NAME=Easypanel Go SDK\nCATCHY_PHRASE=Deployed via Go SDK integration test\nEMAIL=test@example.com"
	err = client.Services.UpdateEnv(ctx, ServiceTypeApp, UpdateEnv{
		SelectService: SelectService{
			ProjectName: projectName,
			ServiceName: serviceName,
		},
		Env: envVars,
	})
	require.NoError(t, err, "update env")
	t.Logf("Environment set:\n%s", envVars)

	// --- Step 6: Deploy the service ---
	t.Log("=== Step 6: Deploy service ===")
	err = client.Services.Deploy(ctx, ServiceTypeApp, SelectService{
		ProjectName: projectName,
		ServiceName: serviceName,
	})
	require.NoError(t, err, "deploy service")
	t.Log("Deployment triggered")

	// --- Step 7: Inspect the service ---
	t.Log("=== Step 7: Inspect service ===")
	inspectResp, err := client.Services.Inspect(ctx, ServiceTypeApp, SelectService{
		ProjectName: projectName,
		ServiceName: serviceName,
	})
	require.NoError(t, err, "inspect service")
	svc := inspectResp.Result.Data.JSON
	// API may return "name" instead of "serviceName"
	inspectedName := svc.ServiceName
	if inspectedName == "" {
		inspectedName = svc.Name
	}
	t.Logf("Service details:")
	t.Logf("  Project: %s", svc.ProjectName)
	t.Logf("  Service: %s (name=%s, serviceName=%s)", inspectedName, svc.Name, svc.ServiceName)
	t.Logf("  Type: %s", svc.Type)
	t.Logf("  Enabled: %v", svc.Enabled)
	assert.Equal(t, projectName, svc.ProjectName)
	assert.Equal(t, serviceName, inspectedName)

	// --- Step 8: Inspect project (should include the service) ---
	t.Log("=== Step 8: Inspect project with services ===")
	projInspect, err := client.Projects.Inspect(ctx, ProjectQuery{ProjectName: projectName})
	require.NoError(t, err, "inspect project")
	t.Logf("Project %s has %d services", projInspect.Result.Data.JSON.Project.Name, len(projInspect.Result.Data.JSON.Services))
	assert.GreaterOrEqual(t, len(projInspect.Result.Data.JSON.Services), 1)

	// --- Step 9: List actions for the service ---
	t.Log("=== Step 9: List actions ===")
	actionsResp, err := client.Actions.List(ctx, ListActionsParams{
		ProjectName: projectName,
		ServiceName: serviceName,
	})
	require.NoError(t, err, "list actions")
	t.Logf("Found %d actions", len(actionsResp.Result.Data.JSON))
	for _, a := range actionsResp.Result.Data.JSON {
		t.Logf("  Action: id=%s type=%s status=%s", a.ID, a.Type, a.Status)
	}

	// --- Step 10: Destroy service ---
	t.Log("=== Step 10: Destroy service ===")
	err = client.Services.Destroy(ctx, ServiceTypeApp, SelectService{
		ProjectName: projectName,
		ServiceName: serviceName,
	})
	require.NoError(t, err, "destroy service")
	t.Log("Service destroyed")

	// --- Step 11: Destroy project ---
	t.Log("=== Step 11: Destroy project ===")
	err = client.Projects.Destroy(ctx, ProjectName{Name: projectName})
	require.NoError(t, err, "destroy project")
	t.Log("Project destroyed")

	// --- Step 12: Verify project is gone ---
	t.Log("=== Step 12: Verify project removed ===")
	listResp2, err := client.Projects.List(ctx)
	require.NoError(t, err, "list projects after cleanup")
	for _, p := range listResp2.Result.Data.JSON {
		assert.NotEqual(t, projectName, p.Name, "project should be removed")
	}
	t.Log("Verified: project no longer exists")

	t.Log("=== FULL LIFECYCLE TEST PASSED ===")
}
