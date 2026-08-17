package httpapi

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"task045-lognorm/internal/lognorm"
)

func TestProbe_EmptyLogsReturnsArray(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/logs", nil)
	rec := httptest.NewRecorder()
	New(lognorm.NewService()).Handler().ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("status=%d want 200; body=%s", rec.Code, rec.Body.String())
	}
	var body map[string]json.RawMessage
	if err := json.Unmarshal(rec.Body.Bytes(), &body); err != nil {
		t.Fatalf("decode response: %v", err)
	}
	var logs []json.RawMessage
	if err := json.Unmarshal(body["logs"], &logs); err != nil {
		t.Fatalf("decode logs: %v", err)
	}
	if logs == nil {
		t.Fatal("logs must be an empty JSON array, not null")
	}
	if len(logs) != 0 {
		t.Fatalf("logs=%d want 0", len(logs))
	}
}
