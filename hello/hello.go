package hello

import (
	"flag"
	"fmt"

	"github.com/dennissetiawan/go-web-framework-benchmark/hello/internal/config"
	"github.com/dennissetiawan/go-web-framework-benchmark/hello/internal/handler"
	"github.com/dennissetiawan/go-web-framework-benchmark/hello/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "hello/etc/hello-api.yaml", "the config file")

func StartServer(port int, goZeroExtraLogic func()) {

	// shared.SharedHandlerParam = handlerParam

	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx, goZeroExtraLogic)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, port)
	server.Start()
}
