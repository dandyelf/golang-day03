package jWriter

import (
	"encoding/json"
	"log"
	"myHttp/types"
	"net/http"
	"strconv"
)

func JWriter(w http.ResponseWriter, r *http.Request, GetPlaces func(limit int, offset int) ([]types.Place, int, error)) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	page := 1
	if queryPage := r.URL.Query().Get("page"); len(queryPage) != 0 {
		if pageNum, err := strconv.Atoi(queryPage); err == nil {
			page = pageNum
		} else {
			returnError(w, "404 Page should be a number!", http.StatusBadRequest)
			return
		}
	}

	perPage := 10
	if queryPerPage := r.URL.Query().Get("per-page"); len(queryPerPage) != 0 {
		if perPageNum, err := strconv.Atoi(queryPerPage); err == nil {
			perPage = perPageNum
		} else {
			returnError(w, "404 Page should be a number!", http.StatusBadRequest)
			return
		}
	}
	list, total, err := GetPlaces(perPage, page)
	if err != nil {
		returnError(w, "400 Invalid page value: "+strconv.Itoa(page), http.StatusBadRequest)
		return
	}
	if err := createJson(types.PageData{
		Total: total,
		Prev:  page - 1,
		Next:  page + 1,
		List:  list,
	}, w); err != nil {
		returnError(w, "400 Server Response Error", http.StatusBadRequest)
		return
	}
}

func createJson(data types.PageData, w http.ResponseWriter) error {
	result := types.Foodcorts{
		Name:     "Places",
		Total:    data.Total,
		Places:   data.List,
		PrevPage: data.Prev,
		NextPage: data.Next,
		LastPage: data.Total / 10,
	}
	return json.NewEncoder(w).Encode(&result)
}

func returnError(w http.ResponseWriter, errorMessage string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	errorResponse := map[string]string{"error": errorMessage}
	err := json.NewEncoder(w).Encode(errorResponse)
	if err != nil {
		log.Println("Error create error.")
	}
}
