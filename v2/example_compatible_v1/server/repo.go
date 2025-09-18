package main

import (
	"context"

	gootel "github.com/hafizarrDev/go-opentelemetry-dev/v2"
)

func repoGetFoo(ctx context.Context) {
	ctx, span := gootel.NewSpan(ctx, "repoGetFoo", "") //nolint:staticcheck,ineffassign
	defer span.End()
}
