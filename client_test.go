package spinwriter_test

import (
	"api/spinwriter"
	"testing"
)

func TestNew(t *testing.T) {
	client := spinwriter.New("email", "api_key")

	if client == nil {
		t.Error("Expected client to be instance of spinwriter.Client")
	}
}
