package client

import "testing"

func TestClientError(t *testing.T) {
	bytes := []byte{}
	_, err := GetError(bytes)
	if err != nil {
		t.Fatalf("%v", err.Error())
	}
}