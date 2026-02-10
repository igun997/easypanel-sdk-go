package easypanel

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	client := New(Config{
		Endpoint: "https://panel.example.com",
		Token:    "my-token",
	})

	require.NotNil(t, client)
	assert.NotNil(t, client.Projects)
	assert.NotNil(t, client.Services)
	assert.NotNil(t, client.Monitor)
	assert.NotNil(t, client.Settings)
	assert.NotNil(t, client.Domains)
}

func TestGetUser(t *testing.T) {
	client := setupTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/api/trpc/auth.getUser", r.URL.Path)
		assert.Equal(t, "test-token", r.Header.Get("Authorization"))

		writeJSON(t, w, newRestResponse(User{
			ID:    "1",
			Email: "admin@test.com",
			Admin: true,
		}))
	})

	resp, err := client.GetUser(context.Background())
	require.NoError(t, err)
	assert.Equal(t, "1", resp.Result.Data.JSON.ID)
	assert.Equal(t, "admin@test.com", resp.Result.Data.JSON.Email)
	assert.True(t, resp.Result.Data.JSON.Admin)
}

func TestGetLicensePayload(t *testing.T) {
	client := setupTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/api/trpc/lemonLicense.getLicensePayload", r.URL.Path)
		assert.Equal(t, "test-token", r.Header.Get("Authorization"))

		w.WriteHeader(http.StatusOK)
	})

	err := client.GetLicensePayload(context.Background(), LicenseTypeLemon)
	require.NoError(t, err)
}

func TestActivateLicense(t *testing.T) {
	client := setupTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/api/trpc/portalLicense.activate", r.URL.Path)
		assert.Equal(t, "test-token", r.Header.Get("Authorization"))

		w.WriteHeader(http.StatusOK)
	})

	err := client.ActivateLicense(context.Background(), LicenseTypePortal)
	require.NoError(t, err)
}

func TestErrorResponse(t *testing.T) {
	client := setupTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(Error{
			OK:           false,
			ErrorMessage: "access denied",
		})
	})

	_, err := client.GetUser(context.Background())
	require.Error(t, err)

	var apiErr *Error
	require.ErrorAs(t, err, &apiErr)
	assert.Equal(t, http.StatusForbidden, apiErr.StatusCode)
	assert.Equal(t, "access denied", apiErr.ErrorMessage)
	assert.Equal(t, "access denied", apiErr.Error())
}
