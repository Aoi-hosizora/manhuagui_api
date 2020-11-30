package service

import (
	"encoding/json"
	"fmt"
	"github.com/Aoi-hosizora/ahlib/xdi"
	"github.com/Aoi-hosizora/manhuagui-backend/src/model/vo"
	"github.com/Aoi-hosizora/manhuagui-backend/src/provide/sn"
	"github.com/Aoi-hosizora/manhuagui-backend/src/static"
)

type CommentService struct {
	httpService *HttpService
}

func NewCommentService() *CommentService {
	return &CommentService{
		httpService: xdi.GetByNameForce(sn.SHttpService).(*HttpService),
	}
}

func (c *CommentService) GetComments(mid uint64, page int32) ([]*vo.Comment, int32, error) {
	url := fmt.Sprintf(static.MANGA_COMMENT_URL, mid, page)
	bs, err := c.httpService.HttpGet(url)
	if err != nil {
		return nil, 0, err
	}
	commentsObj := &vo.Comments{}
	err = json.Unmarshal(bs, commentsObj)
	if err != nil {
		return nil, 0, err
	}

	return []*vo.Comment{}, 0, nil
}
