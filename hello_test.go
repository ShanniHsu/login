package main

import (
	"github.com/gin-gonic/gin"
	"login/controll"
	"testing"
)

func TestHello(t *testing.T) {
	t.Log("hello world")
}

func TestWorld(t *testing.T) {
	t.Error("world of bug")
}

func add(a, b int) int {
	return a + b
}

func Test_add_1_2(t *testing.T) {
	if add(1, 2) != 3 {
		t.Error("wrong result")
	}
}

func TestRegister(t *testing.T) {
	var ctx *gin.Context
	controll.Register(ctx)
}
