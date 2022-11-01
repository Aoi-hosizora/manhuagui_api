package service

import (
	"encoding/json"
	"fmt"
	"github.com/Aoi-hosizora/ahlib/xdi"
	"github.com/Aoi-hosizora/manhuagui-api/src/model/vo"
	"github.com/Aoi-hosizora/manhuagui-api/src/provide/sn"
	"github.com/Aoi-hosizora/manhuagui-api/src/static"
	"strings"
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
	bs, _, err := c.httpService.HttpGet(url, nil)
	if err != nil {
		return nil, 0, err
	}
	commentsObj := &vo.Comments{}
	err = json.Unmarshal(bs, commentsObj)
	if err != nil {
		return nil, 0, err
	}

	objArr := commentsObj.CommentIds
	chains := make([][]string, len(objArr))
	for idx, idsStr := range objArr {
		chain := strings.Split(idsStr, ",")
		if len(chain) > 0 {
			chains[idx] = chain
		}
	}

	out := make([]*vo.Comment, 0, len(chains))
	for _, chain := range chains {
		cmt, ok := commentsObj.Comments[chain[0]]
		if !ok {
			continue
		}
		if cmt.Username == "" {
			cmt.Username = "-"
		}
		if cmt.Avatar == "" {
			cmt.Avatar = "https://cf.hamreus.com/images/default.png"
		}

		timeline := make([]*vo.RepliedComment, 0, len(chain)-1)
		if len(chain) > 1 {
			for idx := len(chain) - 1; idx >= 1; idx-- {
				repliedId := chain[idx]
				if reply, ok := commentsObj.Comments[repliedId]; ok {
					cmt := vo.NewRepliedComment(reply)
					if cmt.Username == "" {
						cmt.Username = "-"
					}
					if cmt.Avatar == "" {
						cmt.Avatar = "https://cf.hamreus.com/images/default.png"
					}
					timeline = append(timeline, cmt)
				}
			}
		}
		cmt.ReplyTimeline = timeline
		out = append(out, cmt)
	}

	return out, commentsObj.Total, nil
}
