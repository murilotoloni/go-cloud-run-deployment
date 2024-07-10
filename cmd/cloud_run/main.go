package main

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/{$}", helloWorld())
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic("Something went wrong when running http server...")
	}
}

func helloWorld() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write(MarshalResponse(Response{Message: "Hello World!"}))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(MarshalResponse(ErrorResponse{Message: "Something went wrong..."}))
		}
	}
}

func MarshalResponse(res interface{}) []byte {
	marshalledResponse, err := json.Marshal(res)
	if err != nil {
		return []byte{}
	}
	return marshalledResponse
}
