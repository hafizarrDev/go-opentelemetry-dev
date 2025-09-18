package main

import (
	"context"
	"log/slog"

	gootel "github.com/hafizarrDev/go-opentelemetry-dev/v2"
	"github.com/hafizarrDev/go-opentelemetry-dev/v2/example/pbfoo"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func extapiGRPCFoo() {
	ctx, span := gootel.RecordSpan(context.Background())
	defer span.End()

	conn, err := grpc.NewClient("localhost:4001", grpc.WithStatsHandler(otelgrpc.NewClientHandler()), grpc.WithTransportCredentials(insecure.NewCredentials()))
	fatalIfErr(err)
	client := pbfoo.NewExampleClient(conn)
	res, err := client.Foo(ctx, &pbfoo.ReqFoo{})
	fatalIfErr(err)
	slog.Info("success", "res", res)
}
