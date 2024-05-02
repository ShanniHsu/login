package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	t.Log("hello world")
}

func TestWorld(t *testing.T) {
	t.Error("world of bug")
}
