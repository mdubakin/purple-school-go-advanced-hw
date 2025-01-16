package random

import (
	"net/http"
	"strconv"
)

const maxRandomInt = 6

func RandomInt() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(strconv.Itoa(getRandomInt(maxRandomInt))))
	}
}
