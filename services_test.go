package easypanel

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestServicesCreate(t *testing.T) {
	params := CreateServiceParams{
		SelectService: SelectService{
			ProjectName: "proj",
			ServiceName: "svc",
		},
	}

	want := Service{
		SelectService: SelectService{
			ProjectName: "proj",
			ServiceName: "svc",
		},
		Type:    ServiceTypeApp,
		Enabled: true,
		Token:   "deploy-token-123",
	}

	client := setupTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/api/trpc/services.app.createService", r.URL.Path)
		assert.Equal(t, "test-token", r.Header.Get("Authorization"))
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"))

		var body CreateServiceParams
		decodeTRPCBody(t, r, &body)
		assert.Equal(t, params, body)

		writeJSON(t, w, newRestResponse(want))
	})

	resp, err := client.Services.Create(context.Background(), ServiceTypeApp, params)
	require.NoError(t, err)
	assert.Equal(t, "proj", resp.Result.Data.JSON.ProjectName)
	assert.Equal(t, "svc", resp.Result.Data.JSON.ServiceName)
	assert.Equal(t, ServiceTypeApp, resp.Result.Data.JSON.Type)
	assert.True(t, resp.Result.Data.JSON.Enabled)
	assert.Equal(t, "deploy-token-123", resp.Result.Data.JSON.Token)
}

func TestServicesInspect(t *testing.T) {
	params := SelectService{
		ProjectName: "proj",
		ServiceName: "svc",
	}

	want := Service{
		SelectService: SelectService{
			ProjectName: "proj",
			ServiceName: "svc",
		},
		Type:    ServiceTypeMySQL,
		Enabled: true,
		Token:   "token-mysql",
	}

	client := setupTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/api/trpc/services.mysql.inspectService", r.URL.Path)
		assert.Equal(t, "test-token", r.Header.Get("Authorization"))

		var input SelectService
		decodeTRPCQuery(t, r, &input)
		assert.Equal(t, params, input)

		writeJSON(t, w, newRestResponse(want))
	})

	resp, err := client.Services.Inspect(context.Background(), ServiceTypeMySQL, params)
	require.NoError(t, err)
	assert.Equal(t, "proj", resp.Result.Data.JSON.ProjectName)
	assert.Equal(t, "svc", resp.Result.Data.JSON.ServiceName)
	assert.Equal(t, ServiceTypeMySQL, resp.Result.Data.JSON.Type)
	assert.True(t, resp.Result.Data.JSON.Enabled)
	assert.Equal(t, "token-mysql", resp.Result.Data.JSON.Token)
}

func TestServicesDestroy(t *testing.T) {
	params := SelectService{
		ProjectName: "proj",
		ServiceName: "svc",
	}

	client := setupTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/api/trpc/services.app.destroyService", r.URL.Path)
		assert.Equal(t, "test-token", r.Header.Get("Authorization"))

		var body SelectService
		decodeTRPCBody(t, r, &body)
		assert.Equal(t, params, body)

		w.WriteHeader(http.StatusOK)
	})

	err := client.Services.Destroy(context.Background(), ServiceTypeApp, params)
	require.NoError(t, err)
}

func TestServicesDeploy(t *testing.T) {
	params := SelectService{
		ProjectName: "proj",
		ServiceName: "svc",
	}

	client := setupTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/api/trpc/services.app.deployService", r.URL.Path)
		assert.Equal(t, "test-token", r.Header.Get("Authorization"))

		var body SelectService
		decodeTRPCBody(t, r, &body)
		assert.Equal(t, params, body)

		w.WriteHeader(http.StatusOK)
	})

	err := client.Services.Deploy(context.Background(), ServiceTypeApp, params)
	require.NoError(t, err)
}

func TestServicesDisable(t *testing.T) {
	params := SelectService{
		ProjectName: "proj",
		ServiceName: "svc",
	}

	client := setupTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/api/trpc/services.redis.disableService", r.URL.Path)
		assert.Equal(t, "test-token", r.Header.Get("Authorization"))

		var body SelectService
		decodeTRPCBody(t, r, &body)
		assert.Equal(t, params, body)

		w.WriteHeader(http.StatusOK)
	})

	err := client.Services.Disable(context.Background(), ServiceTypeRedis, params)
	require.NoError(t, err)
}

func TestServicesEnable(t *testing.T) {
	params := SelectService{
		ProjectName: "proj",
		ServiceName: "svc",
	}

	client := setupTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/api/trpc/services.redis.enableService", r.URL.Path)
		assert.Equal(t, "test-token", r.Header.Get("Authorization"))

		var body SelectService
		decodeTRPCBody(t, r, &body)
		assert.Equal(t, params, body)

		w.WriteHeader(http.StatusOK)
	})

	err := client.Services.Enable(context.Background(), ServiceTypeRedis, params)
	require.NoError(t, err)
}

func TestServicesUpdateEnv(t *testing.T) {
	params := UpdateEnv{
		SelectService: SelectService{
			ProjectName: "proj",
			ServiceName: "svc",
		},
		Env: "KEY=value",
	}

	client := setupTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/api/trpc/services.app.updateEnv", r.URL.Path)
		assert.Equal(t, "test-token", r.Header.Get("Authorization"))

		var body UpdateEnv
		decodeTRPCBody(t, r, &body)
		assert.Equal(t, params, body)

		w.WriteHeader(http.StatusOK)
	})

	err := client.Services.UpdateEnv(context.Background(), ServiceTypeApp, params)
	require.NoError(t, err)
}

func TestServicesUpdateDomains(t *testing.T) {
	params := CreateServiceParams{
		SelectService: SelectService{
			ProjectName: "proj",
			ServiceName: "svc",
		},
		Domains: []DomainParams{
			{Host: "example.com", HTTPS: true, Port: 443},
			{Host: "api.example.com", Port: 8080, Path: "/v1"},
		},
	}

	client := setupTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/api/trpc/services.app.updateDomains", r.URL.Path)
		assert.Equal(t, "test-token", r.Header.Get("Authorization"))

		var body CreateServiceParams
		decodeTRPCBody(t, r, &body)
		assert.Equal(t, params, body)

		w.WriteHeader(http.StatusOK)
	})

	err := client.Services.UpdateDomains(context.Background(), ServiceTypeApp, params)
	require.NoError(t, err)
}

func TestServicesUpdateResources(t *testing.T) {
	params := UpdateResources{
		SelectService: SelectService{
			ProjectName: "proj",
			ServiceName: "svc",
		},
		Resources: Resources{
			CPULimit:          2.0,
			CPUReservation:    0.5,
			MemoryLimit:       1024,
			MemoryReservation: 512,
		},
	}

	client := setupTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/api/trpc/services.app.updateResources", r.URL.Path)
		assert.Equal(t, "test-token", r.Header.Get("Authorization"))

		var body UpdateResources
		decodeTRPCBody(t, r, &body)
		assert.Equal(t, params, body)

		w.WriteHeader(http.StatusOK)
	})

	err := client.Services.UpdateResources(context.Background(), ServiceTypeApp, params)
	require.NoError(t, err)
}

func TestServicesGetServiceLogs(t *testing.T) {
	params := SelectService{
		ProjectName: "proj",
		ServiceName: "svc",
	}

	wantLogs := "2024-01-01T00:00:00Z service started\n2024-01-01T00:00:01Z listening on :8080"

	client := setupTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/api/trpc/logs.getServiceLogs", r.URL.Path)
		assert.Equal(t, "test-token", r.Header.Get("Authorization"))

		var input SelectService
		decodeTRPCQuery(t, r, &input)
		assert.Equal(t, params, input)

		writeJSON(t, w, newRestResponse(wantLogs))
	})

	resp, err := client.Services.GetServiceLogs(context.Background(), params)
	require.NoError(t, err)
	assert.Equal(t, wantLogs, resp.Result.Data.JSON)
}
