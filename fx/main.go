package main

import (
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

type A struct {
	logger *zap.Logger
	b      *B
}

func NewA(logger *zap.Logger, b *B) *A {
	return &A{
		logger: logger,
		b:      b,
	}
}

type B struct {
	V int
}

func NewB() *B {
	return &B{
		V: 1,
	}
}

func main() {
	fx.New(
		fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: log}
		}),
		fx.Provide(
			NewA,
			NewB,
			zap.NewExample,
		),
		fx.Invoke(func(b *B, logger *zap.Logger) {

		}),
	).Run()
}
