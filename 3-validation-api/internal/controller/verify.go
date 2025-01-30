package controller

import (
	"log"
	"net/http"
	"validation/internal/model"
	"validation/internal/usecase/hash"
	"validation/pkg/request"
)

type VerifyHandlerDeps struct {
	hash.HashService
}

type VerifyHandler struct {
	VerifyHandlerDeps
}

func NewVerifyHandler(mux *http.ServeMux, deps VerifyHandlerDeps) {
	h := VerifyHandler{VerifyHandlerDeps: deps}
	mux.HandleFunc("GET /verify/{hash}", h.Verify())
	mux.HandleFunc("POST /send", h.Send())
}

func (h VerifyHandler) Verify() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		hash := r.PathValue("hash")
		email, err := h.HashService.VerifyEmail(hash)
		if err != nil {
			log.Println(err)
			w.Write([]byte("false"))
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.Write([]byte("hello " + email + ": true"))
	}
}

func (h VerifyHandler) Send() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req, err := request.HandleBody[model.SendRequest](w, r)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if err := h.HashService.SaveEmailHash(req.Email); err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write([]byte("verify your email"))
	}
}
