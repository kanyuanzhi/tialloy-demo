package test

import "testing"

func TestClient2(t *testing.T) {
	client := NewClient(2, 1, "bb")
	client.Start()
}
