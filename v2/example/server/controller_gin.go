package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	gootel "github.com/hafizarrDev/go-opentelemetry-dev/v2"
)

func controllerGinFoo(c *gin.Context) {
	ctx, span := gootel.RecordSpan(c)
	defer span.End()

	serviceFoo(ctx)

	c.JSON(http.StatusOK, gin.H{"trace_id": span.SpanContext().TraceID().String()})
}
