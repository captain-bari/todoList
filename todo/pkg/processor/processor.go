package processor

import (
	"net/http"
	"time"
	"todo/pkg/common"
	"todo/pkg/types"
)

func NewProcessor(db database) *processor {
	return &processor{
		db: db,
	}
}

type database interface {
	GetListDetails(uuid string)
}

type processor struct {
	db database
}

func (p *processor) GetListDetails(r common.ApiRequest) (resp common.ApiResponse) {
	req := r.(*types.GetListDetailsReq)
	listUUID := req.ListUUID

	// Get from DB
	p.db.GetListDetails(listUUID)

	resp = common.ApiResponse{
		Resp: &types.GetListDetailsResp{
			Desciption: "description from db",
			TS:         time.Now(), // from DB
		},
		Code: http.StatusOK,
	}
	return
}
