package rest

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"
	"todo/pkg/common"
	"todo/pkg/types"

	"github.com/gorilla/mux"
)

type Server struct {
	server *http.Server
}

type processor interface {
	GetListDetails(r common.ApiRequest) (resp common.ApiResponse)
}

func NewServer(url string, p processor) (*Server, error) {
	router := mux.NewRouter()

	router.Handle("/v1/{listUUID}/list", restJSONHandler{
		GET: restHandler{
			req:     &types.GetListDetailsReq{},
			handler: p.GetListDetails,
			label:   "GetListDetails",
		},
		POST: restHandler{
			req:     &types.GetListDetailsReq{}, // Impliment UpdateListDetails struct
			handler: p.GetListDetails,
			label:   "UpdateListDetails",
		},
		DELETE: restHandler{
			req:     &types.GetListDetailsReq{}, // Impliment DeleteListDetails struct
			handler: p.GetListDetails,
			label:   "DeleteListDetails",
		},
	})

	router.Handle("/v1/list", restJSONHandler{
		GET: restHandler{
			req:     &types.GetCompleteListReq{},
			handler: p.GetListDetails,
			label:   "GetAllListItem",
		},
	})

	router.Handle("/v1/add", restJSONHandler{
		POST: restHandler{
			req:     &types.GetListDetailsReq{}, // Impliment AddListItem struct
			handler: p.GetListDetails,
			label:   "AddListItem",
		},
	})

	srv := &http.Server{
		Addr:              url,
		Handler:           router,
		ReadTimeout:       1 * time.Minute, // Will come from config
		ReadHeaderTimeout: 1 * time.Minute,
		WriteTimeout:      1 * time.Minute,
		IdleTimeout:       1 * time.Minute,
	}

	return &Server{srv}, nil
}

func (s *Server) Run() error {
	fmt.Printf("Starting HTTPS server at [%s]", s.server.Addr)

	go func() {
		err := s.server.ListenAndServeTLS("", "")
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("HTTP Server Run Error:[%s]. Exiting.", err.Error())
			os.Exit(1)
		}
	}()

	return nil
}

// Stop gracefully stops http server.
func (s *Server) Stop() error {
	return s.server.Shutdown(context.Background())
}
