package main

import (
	"testing"
)

func TestHelloMessage(t *testing.T) {
	if HelloMessage() != "Hello world" {
		t.Fail()
	}
}
