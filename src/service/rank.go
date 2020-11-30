package service

import (
	"github.com/Aoi-hosizora/ahlib/xdi"
	"github.com/Aoi-hosizora/manhuagui-backend/src/provide/sn"
)

type RankService struct {
	httpService *HttpService
}

func NewRankService() *RankService {
	return &RankService{
		httpService: xdi.GetByNameForce(sn.SHttpService).(*HttpService),
	}
}
