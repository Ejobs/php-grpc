package grpc

import (
	"github.com/opentracing/opentracing-go"
	"github.com/spf13/cobra"
	rrpc "github.com/spiral/php-grpc"
	rr "github.com/spiral/roadrunner/cmd/rr/cmd"
	"go.elastic.co/apm/module/apmot"
)

func init() {
	cobra.OnInitialize(func() {
		svc, _ := rr.Container.Get(rrpc.ID)
		if _, ok := svc.(*rrpc.Service); ok {
			opentracing.SetGlobalTracer(apmot.New())
		}
	})
}
