package types

import (
	"net/http"
	"time"
	"todo/pkg/common"

	"github.com/gorilla/mux"
)

type GetListDetailsReq struct {
	ListUUID string
}

func (req *GetListDetailsReq) Parse(r *http.Request) (err error) {
	req.ListUUID = mux.Vars(r)["listUUID"]
	return
}
func (req *GetListDetailsReq) GetNewObj() common.ApiRequest {
	return &GetListDetailsReq{}
}
func (req *GetListDetailsReq) Validate() (err error) {
	return
}

type GetListDetailsResp struct {
	Title            string    `json:"title"` // Mandatory
	Desciption       string    `json:"desciption,omitempty"`
	TS               time.Time `json:"timestamp,omitempty"`
	IsMarkedComplete bool      `json:"completed"` // // Mandatory
}

type GetCompleteListReq struct {
}

func (req *GetCompleteListReq) Parse(r *http.Request) (err error) {
	return
}
func (req *GetCompleteListReq) GetNewObj() common.ApiRequest {
	return &GetCompleteListReq{}
}
func (req *GetCompleteListReq) Validate() (err error) {
	return
}

type GetCompleteListResp struct {
	List []GetCompleteListResp `json:"list"`
}
