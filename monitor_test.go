package easypanel

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMonitorGetAdvancedStats(t *testing.T) {
	want := AdvancedStats{
		CPU: []TimeValue{
			{Value: "12.5", Time: "2025-01-01T00:00:00Z"},
			{Value: "15.3", Time: "2025-01-01T00:01:00Z"},
		},
		Disk: []TimeValue{
			{Value: "45.0", Time: "2025-01-01T00:00:00Z"},
		},
		Memory: []TimeValue{
			{Value: "62.1", Time: "2025-01-01T00:00:00Z"},
			{Value: "63.8", Time: "2025-01-01T00:01:00Z"},
		},
		Network: []NetworkTimeValue{
			{
				Value: NetworkValue{Input: 1024, Output: 2048},
				Time:  "2025-01-01T00:00:00Z",
			},
			{
				Value: NetworkValue{Input: 1500, Output: 3000},
				Time:  "2025-01-01T00:01:00Z",
			},
		},
	}

	client := setupTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/api/trpc/monitor.getAdvancedStats", r.URL.Path)
		assert.Equal(t, "test-token", r.Header.Get("Authorization"))

		writeJSON(t, w, newRestResponse(want))
	})

	resp, err := client.Monitor.GetAdvancedStats(context.Background())
	require.NoError(t, err)

	got := resp.Result.Data.JSON
	assert.Len(t, got.CPU, 2)
	assert.Equal(t, "12.5", got.CPU[0].Value)
	assert.Equal(t, "2025-01-01T00:00:00Z", got.CPU[0].Time)
	assert.Equal(t, "15.3", got.CPU[1].Value)

	assert.Len(t, got.Disk, 1)
	assert.Equal(t, "45.0", got.Disk[0].Value)

	assert.Len(t, got.Memory, 2)
	assert.Equal(t, "62.1", got.Memory[0].Value)
	assert.Equal(t, "63.8", got.Memory[1].Value)

	assert.Len(t, got.Network, 2)
	assert.Equal(t, 1024, got.Network[0].Value.Input)
	assert.Equal(t, 2048, got.Network[0].Value.Output)
	assert.Equal(t, "2025-01-01T00:00:00Z", got.Network[0].Time)
	assert.Equal(t, 1500, got.Network[1].Value.Input)
	assert.Equal(t, 3000, got.Network[1].Value.Output)
}

func TestMonitorGetDockerTaskStats(t *testing.T) {
	want := DockerTaskStats{
		"project_web": TaskStatus{Actual: 3, Desired: 3},
		"project_db":  TaskStatus{Actual: 1, Desired: 1},
		"project_api": TaskStatus{Actual: 2, Desired: 5},
	}

	client := setupTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/api/trpc/monitor.getDockerTaskStats", r.URL.Path)
		assert.Equal(t, "test-token", r.Header.Get("Authorization"))

		writeJSON(t, w, newRestResponse(want))
	})

	resp, err := client.Monitor.GetDockerTaskStats(context.Background())
	require.NoError(t, err)

	got := resp.Result.Data.JSON
	require.Len(t, got, 3)

	assert.Equal(t, 3, got["project_web"].Actual)
	assert.Equal(t, 3, got["project_web"].Desired)

	assert.Equal(t, 1, got["project_db"].Actual)
	assert.Equal(t, 1, got["project_db"].Desired)

	assert.Equal(t, 2, got["project_api"].Actual)
	assert.Equal(t, 5, got["project_api"].Desired)
}

func TestMonitorGetMonitorTableData(t *testing.T) {
	want := []ContainerStats{
		{
			ID: "abc123",
			Stats: ContainerStat{
				CPU: struct {
					Percent float64 `json:"percent"`
				}{Percent: 25.5},
				Memory: struct {
					Usage   int     `json:"usage"`
					Percent float64 `json:"percent"`
				}{Usage: 512000000, Percent: 48.2},
				Network: struct {
					In  int `json:"in"`
					Out int `json:"out"`
				}{In: 10240, Out: 20480},
			},
			ProjectName:   "myproject",
			ServiceName:   "web",
			ContainerName: "myproject_web.1",
		},
		{
			ID: "def456",
			Stats: ContainerStat{
				CPU: struct {
					Percent float64 `json:"percent"`
				}{Percent: 5.1},
				Memory: struct {
					Usage   int     `json:"usage"`
					Percent float64 `json:"percent"`
				}{Usage: 128000000, Percent: 12.0},
				Network: struct {
					In  int `json:"in"`
					Out int `json:"out"`
				}{In: 500, Out: 1500},
			},
			ProjectName:   "myproject",
			ServiceName:   "db",
			ContainerName: "myproject_db.1",
		},
	}

	client := setupTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/api/trpc/monitor.getMonitorTableData", r.URL.Path)
		assert.Equal(t, "test-token", r.Header.Get("Authorization"))

		writeJSON(t, w, newRestResponse(want))
	})

	resp, err := client.Monitor.GetMonitorTableData(context.Background())
	require.NoError(t, err)

	got := resp.Result.Data.JSON
	require.Len(t, got, 2)

	assert.Equal(t, "abc123", got[0].ID)
	assert.Equal(t, "myproject", got[0].ProjectName)
	assert.Equal(t, "web", got[0].ServiceName)
	assert.Equal(t, "myproject_web.1", got[0].ContainerName)
	assert.Equal(t, 25.5, got[0].Stats.CPU.Percent)
	assert.Equal(t, 512000000, got[0].Stats.Memory.Usage)
	assert.Equal(t, 48.2, got[0].Stats.Memory.Percent)
	assert.Equal(t, 10240, got[0].Stats.Network.In)
	assert.Equal(t, 20480, got[0].Stats.Network.Out)

	assert.Equal(t, "def456", got[1].ID)
	assert.Equal(t, "myproject", got[1].ProjectName)
	assert.Equal(t, "db", got[1].ServiceName)
	assert.Equal(t, "myproject_db.1", got[1].ContainerName)
	assert.Equal(t, 5.1, got[1].Stats.CPU.Percent)
	assert.Equal(t, 128000000, got[1].Stats.Memory.Usage)
}

func TestMonitorGetSystemStats(t *testing.T) {
	want := SystemStats{
		Uptime: 86400.0,
		MemInfo: MemInfo{
			TotalMemMb:        16384.0,
			UsedMemMb:         8192.0,
			FreeMemMb:         8192.0,
			UsedMemPercentage: 50.0,
			FreeMemPercentage: 50.0,
		},
		DiskInfo: DiskInfo{
			TotalGb:        "500",
			UsedGb:         "200",
			FreeGb:         "300",
			UsedPercentage: "40",
			FreePercentage: "60",
		},
		CPUInfo: CPUInfo{
			UsedPercentage: 35.5,
			Count:          8,
			Loadavg:        []float64{1.5, 2.0, 1.8},
		},
		Network: NetworkInfo{
			InputMb:  1024.5,
			OutputMb: 2048.3,
		},
	}

	client := setupTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/api/trpc/monitor.getSystemStats", r.URL.Path)
		assert.Equal(t, "test-token", r.Header.Get("Authorization"))

		// Ensure input parameter is present and equals {"json":null}
		inputRaw := r.URL.Query().Get("input")
		require.NotEmpty(t, inputRaw, "input query param should be set")
		assert.JSONEq(t, `{"json":null}`, inputRaw)

		writeJSON(t, w, newRestResponse(want))
	})

	resp, err := client.Monitor.GetSystemStats(context.Background())
	require.NoError(t, err)

	got := resp.Result.Data.JSON
	assert.Equal(t, 86400.0, got.Uptime)

	assert.Equal(t, 16384.0, got.MemInfo.TotalMemMb)
	assert.Equal(t, 8192.0, got.MemInfo.UsedMemMb)
	assert.Equal(t, 8192.0, got.MemInfo.FreeMemMb)
	assert.Equal(t, 50.0, got.MemInfo.UsedMemPercentage)
	assert.Equal(t, 50.0, got.MemInfo.FreeMemPercentage)

	assert.Equal(t, "500", got.DiskInfo.TotalGb)
	assert.Equal(t, "200", got.DiskInfo.UsedGb)
	assert.Equal(t, "300", got.DiskInfo.FreeGb)
	assert.Equal(t, "40", got.DiskInfo.UsedPercentage)
	assert.Equal(t, "60", got.DiskInfo.FreePercentage)

	assert.Equal(t, 35.5, got.CPUInfo.UsedPercentage)
	assert.Equal(t, 8, got.CPUInfo.Count)
	assert.Equal(t, []float64{1.5, 2.0, 1.8}, got.CPUInfo.Loadavg)

	assert.Equal(t, 1024.5, got.Network.InputMb)
	assert.Equal(t, 2048.3, got.Network.OutputMb)
}
