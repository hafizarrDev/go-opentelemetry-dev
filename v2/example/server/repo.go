package main

import (
	"context"

	gootel "github.com/hafizarrDev/go-opentelemetry-dev/v2"
)

func repoGetFoo(ctx context.Context) {
	ctx, span := gootel.RecordSpan(ctx) //nolint:staticcheck,ineffassign
	defer span.End()
}
