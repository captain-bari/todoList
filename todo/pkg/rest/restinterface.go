package rest

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand/v2"
	"net/http"
	"todo/pkg/common"
)

type restJSONHandler struct {
	GET    http.Handler
	POST   http.Handler
	PUT    http.Handler
	PATCH  http.Handler
	DELETE http.Handler
}

// Will be called from mux
func (h restJSONHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defaultFunc := func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("[ %s %s %s ] RespCode: %d", r.RequestURI, r.Method, r.RemoteAddr, http.StatusMethodNotAllowed)
		http.Error(w, r.Method+" method not supported", http.StatusMethodNotAllowed)
	}

	switch r.Method {
	case http.MethodGet:
		if h.GET != nil {
			defaultFunc = h.GET.ServeHTTP
		}
	case http.MethodPost:
		if h.POST != nil {
			defaultFunc = h.POST.ServeHTTP
		}
	case http.MethodPut:
		if h.PUT != nil {
			defaultFunc = h.PUT.ServeHTTP
		}
	case http.MethodPatch:
		if h.PATCH != nil {
			defaultFunc = h.PATCH.ServeHTTP
		}
	case http.MethodDelete:
		if h.DELETE != nil {
			defaultFunc = h.DELETE.ServeHTTP
		}
	}
	defaultFunc(w, r)
}

// restHandler will server as a FILTER layer
type restHandler struct {
	handler func(common.ApiRequest) common.ApiResponse
	req     common.ApiRequest
	label   string
}

func (h restHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	eteId := rand.Uint32()
	txnId := rand.Uint32()
	debug := fmt.Sprintf("[ %s %s %s : %d %d ] %s", r.RequestURI, r.Method, r.RemoteAddr, eteId, txnId, h.label)
	var req common.ApiRequest

	if h.req != nil {
		req = h.req.GetNewObj()
		// pre processing
		body, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			fmt.Printf("%s [ RespCode: %d ] [ %s ]", debug, http.StatusBadRequest, err.Error())
			http.Error(w, r.Method+" bad request", http.StatusBadRequest)
			return
		}

		if len(body) != 0 {
			err = json.Unmarshal(body, req)
			if err != nil {
				fmt.Printf("%s [ RespCode: %d ] [ %s ]", debug, http.StatusBadRequest, err.Error())
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
		}

		err = req.Parse(r)
		if err != nil {
			fmt.Printf("%s [ RespCode: %d ] [ %s ]", debug, http.StatusBadRequest, err.Error())
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = req.Validate()
		if err != nil {
			fmt.Printf("%s [ RespCode: %d ] [ %s ]", debug, http.StatusBadRequest, err.Error())
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// process request
		resp := h.handler(req)

		// post processing, response
		if resp.Err != nil {
			fmt.Printf("%s [ RespCode: %d ] [ %s ]", debug, http.StatusInternalServerError, resp.Err.Error())
			http.Error(w, resp.Err.Error(), http.StatusInternalServerError)
			return
		}

		writeSuccessResponse(resp.Code, resp.Resp, w)
	}
}

func writeSuccessResponse(respCode int, response common.Response, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(respCode)
	json.NewEncoder(w).Encode(response)
}
