package rest_api

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/MartMong/rover-planet/internal/utils"
)

type RoverRestAPI struct {
	server *http.Server
}

func NewRoverRestAPI(port int) *RoverRestAPI {
	mux := http.NewServeMux()
	mux.HandleFunc("/rover", handlerRover)

	return &RoverRestAPI{
		server: &http.Server{
			Addr:    fmt.Sprintf(":%d", port),
			Handler: mux,
		},
	}
}

func (r *RoverRestAPI) Start() error {
	return r.server.ListenAndServe()
}

func handlerRover(w http.ResponseWriter, r *http.Request) {
	scanner := bufio.NewScanner(r.Body)
	results := make([]string, 0, 10)

	err := utils.ExecuteRoverWithScanner(scanner, func(outputLog string) {
		results = append(results, outputLog)
	})

	if err != nil {
		results = append(results, err.Error())
		w.WriteHeader(http.StatusBadRequest)
	}

	io.WriteString(w, strings.Join(results, "\n"))
}
