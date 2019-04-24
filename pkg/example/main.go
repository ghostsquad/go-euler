package main

//go:generate go build -buildmode=plugin -o main.so main.go

import (
	"context"
	"github.com/ghostsquad/goeuler/pkg"
)

type solution struct {}

func (s solution) Solve(ctx context.Context) {
	pkg.SolveWith(ctx, "example", func() int {
		answer := 0
		return answer
	})
}

var Solution solution