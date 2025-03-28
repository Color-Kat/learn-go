package link

import (
	"demo/http/configs"
	"demo/http/pkg/event"
	"demo/http/pkg/middleware"
	"demo/http/pkg/request"
	"demo/http/pkg/response"
	"fmt"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type LinkHandlerDeps struct {
	LinkRepository *LinkRepository
	EventBus       *event.EventBus
	Config         *configs.Config
}

type LinkHandler struct {
	LinkRepository *LinkRepository
	EventBus       *event.EventBus
}

func NewLinkHandler(router *http.ServeMux, deps LinkHandlerDeps) {
	handler := LinkHandler{
		LinkRepository: deps.LinkRepository,
		EventBus:       deps.EventBus,
	}
	router.Handle("POST /link", middleware.IsAuthed(handler.Create(), deps.Config))
	router.Handle("PATCH /link/{id}", middleware.IsAuthed(handler.Update(), deps.Config))
	router.Handle("DELETE /link/{id}", middleware.IsAuthed(handler.Delete(), deps.Config))
	router.HandleFunc("GET /{hash}", handler.GoTo())
	router.HandleFunc("GET /link", handler.GetAll())
}

func (handler *LinkHandler) GoTo() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		hash := req.PathValue("hash")
		link, err := handler.LinkRepository.GetByHash(hash)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		//handler.StatRepository.AddClick(link.ID)
		go handler.EventBus.Publish(event.Event{
			Type: event.EventLinkVisited,
			Data: link.ID,
		})

		http.Redirect(w, req, link.Url, http.StatusTemporaryRedirect)
	}
}

func (handler *LinkHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		body, err := request.HandleBody[LinkCreateRequest](&w, req)
		if err != nil {
			return
		}

		link := NewLink(body.Url)

		for {
			existedLink, _ := handler.LinkRepository.GetByHash(link.Hash)
			if existedLink == nil {
				break
			} else {
				link.GenerateHash()
			}
		}

		createdLink, err := handler.LinkRepository.Create(link)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		response.Json(w, createdLink, http.StatusCreated)
	}
}

func (handler *LinkHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		email, ok := req.Context().Value(middleware.ContextEmailKey).(string)
		if ok {
			fmt.Println(email)
		}

		body, err := request.HandleBody[LinkUpdateRequest](&w, req)
		if err != nil {
			return
		}
		id, err := strconv.Atoi(req.PathValue("id"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		link, err := handler.LinkRepository.Update(&Link{
			Model: gorm.Model{ID: uint(id)},
			Url:   body.Url,
			Hash:  body.Hash,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		response.Json(w, link, http.StatusCreated)
	}
}

func (handler *LinkHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		id, err := strconv.Atoi(req.PathValue("id"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		_, err = handler.LinkRepository.GetById(uint(id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		err = handler.LinkRepository.Delete(uint(id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		response.Json(w, nil, 200)
	}
}

func (handler *LinkHandler) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		limit, err := strconv.Atoi(req.URL.Query().Get("limit"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		offset, err := strconv.Atoi(req.URL.Query().Get("offset"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		links := handler.LinkRepository.GetAll(limit, offset)
		count := handler.LinkRepository.Count()
		response.Json(w, &GetAllLinksResponse{
			Links: links,
			Count: count,
		}, 200)
	}
}
