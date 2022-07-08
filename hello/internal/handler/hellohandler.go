package handler

import (
	"crypto/sha256"
	"math/big"
	"net/http"
	"runtime"
	"strconv"
	"time"

	"github.com/dennissetiawan/go-web-framework-benchmark/hello/internal/logic"
	"github.com/dennissetiawan/go-web-framework-benchmark/hello/internal/svc"
	"github.com/dennissetiawan/go-web-framework-benchmark/hello/shared"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func HelloHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		if shared.SharedHandlerParam.CpuBound {
			pow(shared.SharedHandlerParam.Target)
		} else {
			if shared.SharedHandlerParam.SleepTime > 0 {
				time.Sleep(shared.SharedHandlerParam.SleepTimeDuration)
			} else {
				runtime.Gosched()
			}
		}

		l := logic.NewHelloLogic(r.Context(), svcCtx)
		resp, err := l.Hello()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}

}

func pow(targetBits int) {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetBits))

	var hashInt big.Int
	var hash [32]byte
	nonce := 0

	for {
		data := "hello world " + strconv.Itoa(nonce)
		hash = sha256.Sum256([]byte(data))
		hashInt.SetBytes(hash[:])

		if hashInt.Cmp(target) == -1 {
			break
		} else {
			nonce++
		}

		if nonce%100 == 0 {
			runtime.Gosched()
		}
	}
}
