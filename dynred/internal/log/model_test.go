package log

import (
	"testing"
)

func TestGetDomain_port(t *testing.T) {
	expected := "localhost"
	actual, err := getDomain("localhost:8080")
	if err != nil {
		t.Fatal(err)
	}
	if actual != expected {
		t.Errorf(actual, expected)
	}
}

func TestGetDomain_subdomain(t *testing.T) {
	expected := "example.com"
	actual, err := getDomain("t.example.com")
	if err != nil {
		t.Fatal(err)
	}
	if actual != expected {
		t.Errorf(actual, expected)
	}
}
