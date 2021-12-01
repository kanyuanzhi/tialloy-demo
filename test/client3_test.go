package test

import "testing"

func TestClient3(t *testing.T) {
	client := NewClient(2, 1, "cc")
	client.Start()
}
