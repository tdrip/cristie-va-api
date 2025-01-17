package client

import "testing"

func TestClientError(t *testing.T) {
	bytes := []byte{}
	_, err := GetError(bytes)
	// should have an error as body is empty
	if err == nil {
		t.Fatal("Empty data should have produced an error")
	}
}
