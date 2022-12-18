package service

import (
	"fmt"
	"github.com/Aoi-hosizora/ahlib/xdi"
	"github.com/Aoi-hosizora/ahlib/xnumber"
	"github.com/Aoi-hosizora/manhuagui-api/src/model/vo"
	"github.com/Aoi-hosizora/manhuagui-api/src/provide/sn"
	"github.com/Aoi-hosizora/manhuagui-api/src/static"
	"github.com/PuerkitoBio/goquery"
)

type RankService struct {
	httpService   *HttpService
	authorService *AuthorService
}

func NewRankService() *RankService {
	return &RankService{
		httpService:   xdi.GetByNameForce(sn.SHttpService).(*HttpService),
		authorService: xdi.GetByNameForce(sn.SAuthorService).(*AuthorService),
	}
}

func (r *RankService) getRankingList(time string, typ string) ([]*vo.MangaRank, error) {
	url := ""
	part := ""
	if typ == "" || typ == "all" {
		if time == "day" {
			part = "" // /rank/
		} else {
			part = time + ".html" // /rank/week.html
		}
	} else {
		if time == "day" {
			part = typ + ".html" // /rank/japan.html
		} else {
			part = fmt.Sprintf("%s_%s.html", typ, time) // /rank/japan_week.html
		}
	}
	url = fmt.Sprintf(static.MANGA_RANK_URL, part)
	_, doc, err := r.httpService.HttpGetDocument(url, nil)
	if err != nil {
		return nil, err
	} else if doc == nil {
		return nil, nil
	}
	return r.getRankingListFromDoc(doc), nil
}

func (r *RankService) getRankingListFromDoc(doc *goquery.Document) []*vo.MangaRank {
	out := make([]*vo.MangaRank, 0)
	trs := doc.Find("div.top-cont tr:not(.rank-split-first):not(.rank-split):not(:first-child)")
	trs.Each(func(idx int, tr *goquery.Selection) {
		order, _ := xnumber.Atoi8(tr.Find("td.rank-no").Text())
		title := tr.Find("td.rank-title a").Text()
		url := tr.Find("td.rank-title a").AttrOr("href", "")
		status := tr.Find("td.rank-title span").Text()
		authorA := tr.Find("div.rank-author a")
		authors := make([]*vo.TinyAuthor, 0)
		authorA.Each(func(idx int, sel *goquery.Selection) {
			authors = append(authors, r.authorService.GetAuthorFromA(sel))
		})
		newestChapter := tr.Find("div.rank-update a").Text()
		newestDate := tr.Find("td.rank-time").Text()
		score, _ := xnumber.Atof64(tr.Find("td.rank-score").Text())
		trend := uint8(0)
		if tr.Find("td.rank-trend span.trend-down").Length() > 0 {
			trend = uint8(2)
		} else if tr.Find("td.rank-trend span.trend-up").Length() > 0 {
			trend = uint8(1)
		}
		id := static.ParseMid(url)
		rank := &vo.MangaRank{
			Mid:           id,
			Title:         title,
			Cover:         fmt.Sprintf(static.MANGA_COVER_URL, id),
			Url:           static.HOMEPAGE_URL + url,
			Finished:      status == "完结",
			Authors:       authors,
			NewestChapter: newestChapter,
			NewestDate:    newestDate,
			Order:         order,
			Score:         score,
			Trend:         trend,
		}
		out = append(out, rank)
	})

	return out
}

func (r *RankService) GetDayRanking(typ string) ([]*vo.MangaRank, error) {
	return r.getRankingList("day", typ)
}

func (r *RankService) GetWeekRanking(typ string) ([]*vo.MangaRank, error) {
	return r.getRankingList("week", typ)
}

func (r *RankService) GetMonthRanking(typ string) ([]*vo.MangaRank, error) {
	return r.getRankingList("month", typ)
}

func (r *RankService) GetTotalRanking(typ string) ([]*vo.MangaRank, error) {
	return r.getRankingList("total", typ)
}
