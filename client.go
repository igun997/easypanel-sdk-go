package easypanel

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

type httpClient struct {
	baseURL    string
	token      string
	httpClient *http.Client
}

func newHTTPClient(baseURL, token string) *httpClient {
	return &httpClient{
		baseURL: baseURL,
		token:   token,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// trpcInput wraps a value in the tRPC envelope: {"json": value}
type trpcInput struct {
	JSON any `json:"json"`
	Meta any `json:"meta,omitempty"`
}

// get performs a GET request. If input is non-nil, it is JSON-encoded as {"json": input}
// and sent as the ?input= query parameter (tRPC convention). The response is decoded into result.
func (c *httpClient) get(ctx context.Context, route string, input any, result any) error {
	u, err := url.Parse(c.baseURL + route)
	if err != nil {
		return fmt.Errorf("easypanel: invalid url: %w", err)
	}

	// Always send input envelope, even if input is nil (as null)
	envelope := trpcInput{JSON: input}
	inputJSON, err := json.Marshal(envelope)
	if err != nil {
		return fmt.Errorf("easypanel: marshal input: %w", err)
	}
	q := u.Query()
	q.Set("input", string(inputJSON))
	u.RawQuery = q.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return fmt.Errorf("easypanel: create request: %w", err)
	}

	return c.do(req, result)
}

// post performs a POST request with a JSON body wrapped in tRPC envelope {"json": body}.
// The response is decoded into result. If result is nil, the response body is discarded.
func (c *httpClient) post(ctx context.Context, route string, body any, result any) error {
	var buf []byte
	if body != nil {
		envelope := trpcInput{JSON: body}
		var err error
		buf, err = json.Marshal(envelope)
		if err != nil {
			return fmt.Errorf("easypanel: marshal body: %w", err)
		}
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.baseURL+route, bytes.NewReader(buf))
	if err != nil {
		return fmt.Errorf("easypanel: create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	return c.do(req, result)
}

// do executes the request with retry logic (1 retry on 5xx, 1s delay).
func (c *httpClient) do(req *http.Request, result any) error {
	req.Header.Set("Authorization", c.token)

	// Read body into memory for potential retry
	var bodyBytes []byte
	if req.Body != nil {
		var err error
		bodyBytes, err = io.ReadAll(req.Body)
		if err != nil {
			return fmt.Errorf("easypanel: read request body: %w", err)
		}
		req.Body = io.NopCloser(bytes.NewReader(bodyBytes))
	}

	var lastErr error
	for attempt := 0; attempt < 2; attempt++ {
		if attempt > 0 {
			time.Sleep(1 * time.Second)
			// Reset body for retry
			if bodyBytes != nil {
				req.Body = io.NopCloser(bytes.NewReader(bodyBytes))
			}
		}

		resp, err := c.httpClient.Do(req)
		if err != nil {
			lastErr = fmt.Errorf("easypanel: request failed: %w", err)
			continue
		}

		respBody, err := io.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			lastErr = fmt.Errorf("easypanel: read response: %w", err)
			continue
		}

		// Retry on 5xx
		if resp.StatusCode >= 500 && attempt < 1 {
			lastErr = fmt.Errorf("easypanel: server error: %d", resp.StatusCode)
			continue
		}

		// Non-2xx error
		if resp.StatusCode < 200 || resp.StatusCode >= 300 {
			apiErr := &Error{StatusCode: resp.StatusCode}
			if json.Unmarshal(respBody, apiErr) != nil || apiErr.ErrorMessage == "" {
				apiErr.ErrorMessage = string(respBody)
			}
			return apiErr
		}

		// Decode response
		if result != nil {
			if err := json.Unmarshal(respBody, result); err != nil {
				return fmt.Errorf("easypanel: decode response: %w", err)
			}
		}

		return nil
	}

	return lastErr
}
