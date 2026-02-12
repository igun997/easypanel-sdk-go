package easypanel

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/websocket"
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

func TestServicesStop(t *testing.T) {
	params := SelectService{
		ProjectName: "proj",
		ServiceName: "svc",
	}

	client := setupTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/api/trpc/services.app.stopService", r.URL.Path)
		assert.Equal(t, "test-token", r.Header.Get("Authorization"))

		var body SelectService
		decodeTRPCBody(t, r, &body)
		assert.Equal(t, params, body)

		w.WriteHeader(http.StatusOK)
	})

	err := client.Services.Stop(context.Background(), ServiceTypeApp, params)
	require.NoError(t, err)
}

func TestServicesRestart(t *testing.T) {
	params := SelectService{
		ProjectName: "proj",
		ServiceName: "svc",
	}

	client := setupTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/api/trpc/services.app.restartService", r.URL.Path)
		assert.Equal(t, "test-token", r.Header.Get("Authorization"))

		var body SelectService
		decodeTRPCBody(t, r, &body)
		assert.Equal(t, params, body)

		w.WriteHeader(http.StatusOK)
	})

	err := client.Services.Restart(context.Background(), ServiceTypeApp, params)
	require.NoError(t, err)
}

func TestServicesUpdateEnv(t *testing.T) {
	params := UpdateEnv{
		SelectService: SelectService{
			ProjectName: "proj",
			ServiceName: "svc",
		},
		Env:          "KEY=value",
		CreateDotEnv: true,
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

func TestServicesUpdateSourceDockerfile(t *testing.T) {
	params := UpdateDockerfile{
		SelectService: SelectService{
			ProjectName: "proj",
			ServiceName: "svc",
		},
		Dockerfile: "FROM node:18\nRUN npm install",
	}

	client := setupTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/api/trpc/services.app.updateSourceDockerfile", r.URL.Path)
		assert.Equal(t, "test-token", r.Header.Get("Authorization"))

		var body UpdateDockerfile
		decodeTRPCBody(t, r, &body)
		assert.Equal(t, params, body)

		w.WriteHeader(http.StatusOK)
	})

	err := client.Services.UpdateSourceDockerfile(context.Background(), ServiceTypeApp, params)
	require.NoError(t, err)
}

func TestServicesUpdateRedirects(t *testing.T) {
	params := UpdateRedirects{
		SelectService: SelectService{
			ProjectName: "proj",
			ServiceName: "svc",
		},
		Redirects: []RedirectParams{
			{
				Name:        "Old Domain to New Domain",
				Enabled:     true,
				Regex:       "^https?://old-domain.com/(.*)",
				Replacement: "https://new-domain.com/${1}",
				Permanent:   false,
			},
		},
	}

	client := setupTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/api/trpc/services.app.updateRedirects", r.URL.Path)
		assert.Equal(t, "test-token", r.Header.Get("Authorization"))

		var body UpdateRedirects
		decodeTRPCBody(t, r, &body)
		assert.Equal(t, params, body)

		w.WriteHeader(http.StatusOK)
	})

	err := client.Services.UpdateRedirects(context.Background(), ServiceTypeApp, params)
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

func TestServicesStreamLogs(t *testing.T) {
	upgrader := websocket.Upgrader{}

	messages := []LogMessage{
		{Output: "2024-01-01T00:00:00Z service started\r\n"},
		{Output: "2024-01-01T00:00:01Z listening on :8080\r\n"},
		{Output: "2024-01-01T00:00:02Z request received\r\n"},
	}

	var capturedToken, capturedService, capturedCompose string

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/ws/serviceLogs", r.URL.Path)

		capturedToken = r.URL.Query().Get("token")
		capturedService = r.URL.Query().Get("service")
		capturedCompose = r.URL.Query().Get("compose")

		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			t.Fatalf("upgrade: %v", err)
			return
		}
		defer conn.Close()

		for _, msg := range messages {
			if err := conn.WriteJSON(msg); err != nil {
				return
			}
		}
		conn.WriteMessage(websocket.CloseMessage,
			websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	}))
	defer server.Close()

	client := New(Config{
		Endpoint: server.URL,
		Token:    "test-token",
	})

	ctx := context.Background()
	ch, err := client.Services.StreamLogs(ctx, StreamLogsParams{
		ProjectName: "myproj",
		ServiceName: "web",
		Token:       "deploy-token-abc",
		Compose:     false,
	})
	require.NoError(t, err)

	var received []LogMessage
	for msg := range ch {
		received = append(received, msg)
	}

	assert.Equal(t, messages, received)
	assert.Equal(t, "deploy-token-abc", capturedToken)
	assert.Equal(t, "myproj_web", capturedService)
	assert.Equal(t, "false", capturedCompose)
}

func TestServicesStreamLogs_ContextCancel(t *testing.T) {
	upgrader := websocket.Upgrader{}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			return
		}
		defer conn.Close()

		// Keep sending messages until the connection is closed.
		for {
			err := conn.WriteJSON(LogMessage{Output: "log line\r\n"})
			if err != nil {
				return
			}
		}
	}))
	defer server.Close()

	client := New(Config{
		Endpoint: server.URL,
		Token:    "test-token",
	})

	ctx, cancel := context.WithCancel(context.Background())
	ch, err := client.Services.StreamLogs(ctx, StreamLogsParams{
		ProjectName: "proj",
		ServiceName: "svc",
		Token:       "tok",
	})
	require.NoError(t, err)

	// Read one message to confirm stream works.
	msg := <-ch
	assert.Equal(t, "log line\r\n", msg.Output)

	// Cancel context — channel should close.
	cancel()

	// Drain remaining messages; channel must eventually close.
	for range ch {
	}
}
