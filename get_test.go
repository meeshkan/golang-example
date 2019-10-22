package main

import "testing"

func TestGetCall(t *testing.T) {
	_, err := makeGetCall()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}
