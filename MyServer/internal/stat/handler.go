package stat

import (
	"demo/http/configs"
	"demo/http/pkg/middleware"
	"demo/http/pkg/response"
	"net/http"
	"time"
)

const (
	GroupByDay   = "day"
	GroupByMonth = "month"
)

type StatHandlerDeps struct {
	StatRepository *StatRepository
	Config         *configs.Config
}

type StatHandler struct {
	StatRepository *StatRepository
}

func NewStatHandler(router *http.ServeMux, deps StatHandlerDeps) {
	handler := StatHandler{
		StatRepository: deps.StatRepository,
	}

	router.Handle("GET /stat", middleware.IsAuthed(handler.GetAll(), deps.Config))
}

func (handler *StatHandler) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		from, err1 := time.Parse("2006-01-02", req.URL.Query().Get("from"))
		to, err2 := time.Parse("2006-01-02", req.URL.Query().Get("to"))
		if err1 != nil || err2 != nil {
			http.Error(w, "Wrong date", http.StatusBadRequest)
			return
		}

		by := req.URL.Query().Get("by")
		if by != GroupByDay && by != GroupByMonth {
			http.Error(w, "Wrong by", http.StatusBadRequest)
			return
		}

		stats := handler.StatRepository.GetStats(by, from, to)

		response.Json(w, &stats, 200)
	}
}
