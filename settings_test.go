package easypanel

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSettingsChangeCredentials(t *testing.T) {
	client := setupTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/api/trpc/settings.changeCredentials", r.URL.Path)
		assert.Equal(t, "test-token", r.Header.Get("Authorization"))

		var params ChangeCredentialsParams
		decodeTRPCBody(t, r, &params)
		assert.Equal(t, "a@b.com", params.Email)
		assert.Equal(t, "old", params.OldPassword)
		assert.Equal(t, "new", params.NewPassword)

		w.WriteHeader(http.StatusOK)
	})

	err := client.Settings.ChangeCredentials(context.Background(), ChangeCredentialsParams{
		Email:       "a@b.com",
		OldPassword: "old",
		NewPassword: "new",
	})
	require.NoError(t, err)
}

func TestSettingsGetGithubToken(t *testing.T) {
	client := setupTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/api/trpc/settings.getGithubToken", r.URL.Path)
		assert.Equal(t, "test-token", r.Header.Get("Authorization"))

		writeJSON(t, w, newRestResponse("ghp_test123"))
	})

	resp, err := client.Settings.GetGithubToken(context.Background())
	require.NoError(t, err)
	assert.Equal(t, "ghp_test123", resp.Result.Data.JSON)
}

func TestSettingsGetLetsEncryptEmail(t *testing.T) {
	client := setupTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/api/trpc/settings.getLetsEncryptEmail", r.URL.Path)
		assert.Equal(t, "test-token", r.Header.Get("Authorization"))

		writeJSON(t, w, newRestResponse("admin@test.com"))
	})

	resp, err := client.Settings.GetLetsEncryptEmail(context.Background())
	require.NoError(t, err)
	assert.Equal(t, "admin@test.com", resp.Result.Data.JSON)
}

func TestSettingsGetPanelDomain(t *testing.T) {
	client := setupTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/api/trpc/settings.getPanelDomain", r.URL.Path)
		assert.Equal(t, "test-token", r.Header.Get("Authorization"))

		writeJSON(t, w, newRestResponse(PanelDomain{
			ServeOnIP:          true,
			PanelDomain:        "panel.example.com",
			DefaultPanelDomain: "panel.example.com",
		}))
	})

	resp, err := client.Settings.GetPanelDomain(context.Background())
	require.NoError(t, err)
	assert.True(t, resp.Result.Data.JSON.ServeOnIP)
	assert.Equal(t, "panel.example.com", resp.Result.Data.JSON.PanelDomain)
	assert.Equal(t, "panel.example.com", resp.Result.Data.JSON.DefaultPanelDomain)
}

func TestSettingsGetServerIp(t *testing.T) {
	client := setupTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/api/trpc/settings.getServerIp", r.URL.Path)
		assert.Equal(t, "test-token", r.Header.Get("Authorization"))

		writeJSON(t, w, newRestResponse("1.2.3.4"))
	})

	resp, err := client.Settings.GetServerIp(context.Background())
	require.NoError(t, err)
	assert.Equal(t, "1.2.3.4", resp.Result.Data.JSON)
}

func TestSettingsRestartEasypanel(t *testing.T) {
	client := setupTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/api/trpc/settings.restartEasypanel", r.URL.Path)
		assert.Equal(t, "test-token", r.Header.Get("Authorization"))

		w.WriteHeader(http.StatusOK)
	})

	err := client.Settings.RestartEasypanel(context.Background())
	require.NoError(t, err)
}

func TestSettingsSetGithubToken(t *testing.T) {
	client := setupTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/api/trpc/settings.setGithubToken", r.URL.Path)
		assert.Equal(t, "test-token", r.Header.Get("Authorization"))

		var params GithubTokenParams
		decodeTRPCBody(t, r, &params)
		assert.Equal(t, "ghp_new", params.GithubToken)

		writeJSON(t, w, newRestResponse("ghp_new"))
	})

	resp, err := client.Settings.SetGithubToken(context.Background(), GithubTokenParams{
		GithubToken: "ghp_new",
	})
	require.NoError(t, err)
	assert.Equal(t, "ghp_new", resp.Result.Data.JSON)
}

func TestSettingsSetPanelDomain(t *testing.T) {
	client := setupTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/api/trpc/settings.setPanelDomain", r.URL.Path)
		assert.Equal(t, "test-token", r.Header.Get("Authorization"))

		var params PanelDomainParams
		decodeTRPCBody(t, r, &params)
		assert.True(t, params.ServeOnIP)
		assert.Equal(t, "panel.example.com", params.PanelDomain)
		assert.Equal(t, "panel.example.com", params.DefaultPanelDomain)

		w.WriteHeader(http.StatusOK)
	})

	err := client.Settings.SetPanelDomain(context.Background(), PanelDomainParams{
		ServeOnIP:          true,
		PanelDomain:        "panel.example.com",
		DefaultPanelDomain: "panel.example.com",
	})
	require.NoError(t, err)
}

func TestSettingsPruneDockerImages(t *testing.T) {
	client := setupTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/api/trpc/settings.pruneDockerImages", r.URL.Path)
		assert.Equal(t, "test-token", r.Header.Get("Authorization"))

		writeJSON(t, w, newRestResponse("pruned 3 images"))
	})

	resp, err := client.Settings.PruneDockerImages(context.Background())
	require.NoError(t, err)
	assert.Equal(t, "pruned 3 images", resp.Result.Data.JSON)
}

func TestSettingsSetDockerPruneDaily(t *testing.T) {
	client := setupTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/api/trpc/settings.setPruneDockerDaily", r.URL.Path)
		assert.Equal(t, "test-token", r.Header.Get("Authorization"))

		var params PruneDockerDailyParams
		decodeTRPCBody(t, r, &params)
		assert.True(t, params.PruneDockerDaily)

		writeJSON(t, w, newRestResponse(true))
	})

	resp, err := client.Settings.SetDockerPruneDaily(context.Background(), PruneDockerDailyParams{
		PruneDockerDaily: true,
	})
	require.NoError(t, err)
	assert.True(t, resp.Result.Data.JSON)
}
