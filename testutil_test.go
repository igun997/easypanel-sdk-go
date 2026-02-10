package easypanel

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

// setupTestClient creates a test HTTP server and returns a Client configured to use it.
// The handler receives all requests made by the client.
func setupTestClient(t *testing.T, handler http.HandlerFunc) *Client {
	t.Helper()
	server := httptest.NewServer(handler)
	t.Cleanup(server.Close)
	return New(Config{
		Endpoint: server.URL,
		Token:    "test-token",
	})
}

// mustMarshal JSON-encodes v and panics on error. For use in test fixtures only.
func mustMarshal(t *testing.T, v any) []byte {
	t.Helper()
	b, err := json.Marshal(v)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	return b
}

// newRestResponse is a test helper to build a RestResponse[T] value.
func newRestResponse[T any](data T) RestResponse[T] {
	var resp RestResponse[T]
	resp.Result.Data.JSON = data
	return resp
}

// decodeTRPCBody decodes a tRPC-enveloped POST body {"json": ...} into target.
func decodeTRPCBody(t *testing.T, r *http.Request, target any) {
	t.Helper()
	var envelope struct {
		JSON json.RawMessage `json:"json"`
	}
	err := json.NewDecoder(r.Body).Decode(&envelope)
	require.NoError(t, err, "decode tRPC envelope")
	err = json.Unmarshal(envelope.JSON, target)
	require.NoError(t, err, "decode tRPC body")
}

// decodeTRPCQuery decodes a tRPC GET ?input={"json": ...} query param into target.
func decodeTRPCQuery(t *testing.T, r *http.Request, target any) {
	t.Helper()
	inputRaw := r.URL.Query().Get("input")
	require.NotEmpty(t, inputRaw, "input query param should be set")
	var envelope struct {
		JSON json.RawMessage `json:"json"`
	}
	err := json.Unmarshal([]byte(inputRaw), &envelope)
	require.NoError(t, err, "decode tRPC query envelope")
	err = json.Unmarshal(envelope.JSON, target)
	require.NoError(t, err, "decode tRPC query body")
}

// writeJSON writes a JSON response to the http.ResponseWriter.
func writeJSON(t *testing.T, w http.ResponseWriter, v any) {
	t.Helper()
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(v); err != nil {
		t.Fatalf("encode response: %v", err)
	}
}
