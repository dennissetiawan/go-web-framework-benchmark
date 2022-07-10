package handler

import (
	"net/http"

	"github.com/dennissetiawan/go-web-framework-benchmark/hello/internal/logic"
	"github.com/dennissetiawan/go-web-framework-benchmark/hello/internal/svc"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func HelloHandler(svcCtx *svc.ServiceContext, goZeroExtraLogic func()) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		goZeroExtraLogic()

		l := logic.NewHelloLogic(r.Context(), svcCtx)
		resp, err := l.Hello()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}

}
