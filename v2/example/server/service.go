package main

import (
	"context"

	gootel "github.com/hafizarrDev/go-opentelemetry-dev/v2"
)

func serviceFoo(ctx context.Context) {
	ctx, span := gootel.RecordSpan(ctx)
	defer span.End()

	repoGetFoo(ctx)
}
