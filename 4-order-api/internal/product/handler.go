package product

import (
	"fmt"
	"net/http"
	"orderapi/pkg/request"
	"orderapi/pkg/response"
	"strconv"
)

type ProductHandler struct {
	*ProductService
}

func NewProductHandler(router *http.ServeMux, productService *ProductService) {
	p := ProductHandler{ProductService: productService}
	router.HandleFunc("GET /product/{id}", p.GetByID)
	router.HandleFunc("POST /product", p.Create)
	router.HandleFunc("PUT /product/{id}", p.Update)
	router.HandleFunc("DELETE /product/{id}", p.Delete)
}

func (h *ProductHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(r.PathValue("id"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	product, err := h.ProductService.GetByID(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	response.EncodeJSON(w, product)
}

type GetProductResponse struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Category string `json:"category"`
	Price    int    `json:"price"`
}

func (h *ProductHandler) Create(w http.ResponseWriter, r *http.Request) {
	b, err := request.HandleBody[CreateProductRequest](w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	product, err := h.ProductService.Create(b.Name, b.Category, b.Price)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response.EncodeJSON(w, product)
}

type CreateProductRequest struct {
	Name     string `json:"name" validate:"required,min=2"`
	Category string `json:"category" validate:"required,min=2"`
	Price    int    `json:"price" validate:"gt=0"`
}

func (h *ProductHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(r.PathValue("id"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	b, err := request.HandleBody[UpdateProductRequest](w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	product, err := h.ProductService.Update(uint(id), b.Name, b.Category, b.Price)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	response.EncodeJSON(w, product)
}

type UpdateProductRequest struct {
	Name     string `json:"name,omitempty" validate:"omitempty"`
	Category string `json:"category,omitempty" validate:"omitempty"`
	Price    int    `json:"price,omitempty" validate:"omitempty,gt=0"`
}

func (h *ProductHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(r.PathValue("id"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.ProductService.Delete(uint(id)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	response.EncodeJSON(w, fmt.Sprintf("Продукт с ID %v удален", id))
}
