package handler

import (
	"net/http"

	"github.com/dennissetiawan/go-web-framework-benchmark/hello/internal/logic"
	"github.com/dennissetiawan/go-web-framework-benchmark/hello/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func HelloHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewHelloLogic(r.Context(), svcCtx)
		resp, err := l.Hello()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
