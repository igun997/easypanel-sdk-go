package easypanel

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestProjectsCanCreate(t *testing.T) {
	expected := newRestResponse(true)

	client := setupTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/api/trpc/projects.canCreateProject", r.URL.Path)
		writeJSON(t, w, expected)
	})

	resp, err := client.Projects.CanCreate(context.Background())
	require.NoError(t, err)
	assert.Equal(t, true, resp.Result.Data.JSON)
}

func TestProjectsCreate(t *testing.T) {
	expectedProject := ProjectInfo{
		Name:      "test-project",
		CreatedAt: "2025-01-01T00:00:00.000Z",
	}
	expected := newRestResponse(expectedProject)

	client := setupTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/api/trpc/projects.createProject", r.URL.Path)

		var params ProjectName
		decodeTRPCBody(t, r, &params)
		assert.Equal(t, "test-project", params.Name)

		writeJSON(t, w, expected)
	})

	resp, err := client.Projects.Create(context.Background(), ProjectName{Name: "test-project"})
	require.NoError(t, err)
	assert.Equal(t, "test-project", resp.Result.Data.JSON.Name)
	assert.Equal(t, "2025-01-01T00:00:00.000Z", resp.Result.Data.JSON.CreatedAt)
}

func TestProjectsDestroy(t *testing.T) {
	client := setupTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/api/trpc/projects.destroyProject", r.URL.Path)

		var params ProjectName
		decodeTRPCBody(t, r, &params)
		assert.Equal(t, "test-project", params.Name)

		w.WriteHeader(http.StatusOK)
	})

	err := client.Projects.Destroy(context.Background(), ProjectName{Name: "test-project"})
	require.NoError(t, err)
}

func TestProjectsInspect(t *testing.T) {
	expectedInspect := ProjectInspect{
		Project: ProjectInfo{
			Name:      "test-project",
			CreatedAt: "2025-01-01T00:00:00.000Z",
		},
		Services: []Service{
			{
				SelectService: SelectService{
					ProjectName: "test-project",
					ServiceName: "web",
				},
				Type:    ServiceTypeApp,
				Enabled: true,
			},
		},
	}
	expected := newRestResponse(expectedInspect)

	client := setupTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/api/trpc/projects.inspectProject", r.URL.Path)

		var query ProjectQuery
		decodeTRPCQuery(t, r, &query)
		assert.Equal(t, "test-project", query.ProjectName)

		writeJSON(t, w, expected)
	})

	resp, err := client.Projects.Inspect(context.Background(), ProjectQuery{ProjectName: "test-project"})
	require.NoError(t, err)
	assert.Equal(t, "test-project", resp.Result.Data.JSON.Project.Name)
	assert.Len(t, resp.Result.Data.JSON.Services, 1)
	assert.Equal(t, "web", resp.Result.Data.JSON.Services[0].ServiceName)
	assert.Equal(t, ServiceTypeApp, resp.Result.Data.JSON.Services[0].Type)
}

func TestProjectsList(t *testing.T) {
	projects := []ProjectInfo{
		{Name: "project-one", CreatedAt: "2025-01-01T00:00:00.000Z"},
		{Name: "project-two", CreatedAt: "2025-02-01T00:00:00.000Z"},
	}
	expected := newRestResponse(projects)

	client := setupTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/api/trpc/projects.listProjects", r.URL.Path)
		writeJSON(t, w, expected)
	})

	resp, err := client.Projects.List(context.Background())
	require.NoError(t, err)
	require.Len(t, resp.Result.Data.JSON, 2)
	assert.Equal(t, "project-one", resp.Result.Data.JSON[0].Name)
	assert.Equal(t, "project-two", resp.Result.Data.JSON[1].Name)
}

func TestProjectsListWithServices(t *testing.T) {
	data := ProjectsWithServices{
		Projects: []ProjectInfo{
			{Name: "project-one", CreatedAt: "2025-01-01T00:00:00.000Z"},
		},
		Services: []Service{
			{
				SelectService: SelectService{
					ProjectName: "project-one",
					ServiceName: "api",
				},
				Type:    ServiceTypeApp,
				Enabled: true,
			},
			{
				SelectService: SelectService{
					ProjectName: "project-one",
					ServiceName: "db",
				},
				Type:    ServiceTypePostgres,
				Enabled: true,
			},
		},
	}
	expected := newRestResponse(data)

	client := setupTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/api/trpc/projects.listProjectsAndServices", r.URL.Path)
		writeJSON(t, w, expected)
	})

	resp, err := client.Projects.ListWithServices(context.Background())
	require.NoError(t, err)
	require.Len(t, resp.Result.Data.JSON.Projects, 1)
	assert.Equal(t, "project-one", resp.Result.Data.JSON.Projects[0].Name)
	require.Len(t, resp.Result.Data.JSON.Services, 2)
	assert.Equal(t, "api", resp.Result.Data.JSON.Services[0].ServiceName)
	assert.Equal(t, ServiceTypeApp, resp.Result.Data.JSON.Services[0].Type)
	assert.Equal(t, "db", resp.Result.Data.JSON.Services[1].ServiceName)
	assert.Equal(t, ServiceTypePostgres, resp.Result.Data.JSON.Services[1].Type)
}
