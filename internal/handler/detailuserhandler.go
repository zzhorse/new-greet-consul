package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"new-greet/internal/logic"
	"new-greet/internal/svc"
	"new-greet/internal/types"
)

func DetailUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		//if err := httpx.Parse(r, &req); err != nil {
		//	httpx.ErrorCtx(r.Context(), w, err)
		//	return
		//}

		l := logic.NewDetailUserLogic(r.Context(), svcCtx)
		resp, err := l.DetailUser(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
