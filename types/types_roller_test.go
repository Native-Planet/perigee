package types

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestShipListRPCs(t *testing.T) {
	for _, raw := range []string{
		`[{"name":"~zod","address":"owner"}]`,
		`["zod"]`,
		`[0]`,
	} {
		t.Run(raw, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.Header.Get("X-Test") != "header-value" {
					t.Error("custom header missing")
				}
				json.NewEncoder(w).Encode(map[string]any{"jsonrpc": "2.0", "result": json.RawMessage(raw), "id": "1"})
			}))
			defer server.Close()
			client := &Client{Endpoint: server.URL, HttpClient: server.Client(), Headers: http.Header{"X-Test": []string{"header-value"}}}
			spawned, err := client.GetSpawned(context.Background(), "~zod")
			if err != nil {
				t.Fatal(err)
			}
			unspawned, err := client.GetUnspawned(context.Background(), "~zod")
			if err != nil {
				t.Fatal(err)
			}
			for _, ships := range [][]ShipInfo{spawned, unspawned} {
				if len(ships) != 1 || ships[0].Name != "~zod" {
					t.Fatalf("ships = %#v", ships)
				}
			}
		})
	}
}

func TestWhenNextBatchTimestamp(t *testing.T) {
	target := time.Now().Add(time.Hour).Truncate(time.Millisecond)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]any{"jsonrpc": "2.0", "result": target.UnixMilli(), "id": "1"})
	}))
	defer server.Close()
	client := &Client{Endpoint: server.URL, HttpClient: server.Client()}
	before := time.Now()
	got, err := client.WhenNextBatch(context.Background())
	after := time.Now()
	if err != nil {
		t.Fatal(err)
	}
	if got < target.Sub(after) || got > target.Sub(before) {
		t.Fatalf("duration %v outside request interval", got)
	}
}
