package request

import (
	"net/http"
	"validation/pkg/response"
)

func HandleBody[T any](w http.ResponseWriter, r *http.Request) (*T, error) {
	body, err := decode[T](r.Body)
	if err != nil {
		response.EncodeJSONWithHeaders(w, err.Error(), http.StatusBadRequest, nil)
		return nil, err
	}

	if err := validate(body); err != nil {
		response.EncodeJSONWithHeaders(w, err.Error(), http.StatusBadRequest, nil)
		return nil, err
	}
	return &body, nil
}
