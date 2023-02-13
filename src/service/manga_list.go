package service

import (
	"fmt"
	"github.com/Aoi-hosizora/ahlib/xdi"
	"github.com/Aoi-hosizora/ahlib/xnumber"
	"github.com/Aoi-hosizora/manhuagui-api/src/model/param"
	"github.com/Aoi-hosizora/manhuagui-api/src/model/vo"
	"github.com/Aoi-hosizora/manhuagui-api/src/provide/sn"
	"github.com/Aoi-hosizora/manhuagui-api/src/static"
	"github.com/PuerkitoBio/goquery"
	"strings"
	"sync"
)

type MangaListService struct {
	httpService     *HttpService
	categoryService *CategoryService
	rankService     *RankService
}

func NewMangaListService() *MangaListService {
	return &MangaListService{
		httpService:     xdi.GetByNameForce(sn.SHttpService).(*HttpService),
		categoryService: xdi.GetByNameForce(sn.SCategoryService).(*CategoryService),
		rankService:     xdi.GetByNameForce(sn.SRankService).(*RankService),
	}
}

func (m *MangaListService) getMangas(doc *goquery.Document, tagIndex int, tagName string) (*vo.MangaGroup, []*vo.MangaGroup, []*vo.MangaGroup) {
	// get top mangas
	topMangas := make([]*vo.TinyManga, 0)
	topMangaUl := doc.Find("div.cmt-cont ul:nth-child(" + xnumber.Itoa(tagIndex) + ")") // <<<
	topMangaUl.Find("li").Each(func(idx int, li *goquery.Selection) {
		topMangas = append(topMangas, m.getTinyMangaPageFromLi(li, true))
	})
	topGroup := &vo.MangaGroup{
		Title:  "",
		Mangas: m.tinyMangasToTinyBlockMangas(topMangas),
	}

	// get group mangas
	groups := make([]*vo.MangaGroup, 0)
	otherMangaUl := doc.Find("div#" + tagName + "Cont ul") // <<<
	otherMangaUl.Each(func(idx int, sel *goquery.Selection) {
		groupTitle := doc.Find("div#" + tagName + "Bar li:nth-child(" + xnumber.Itoa(idx+1) + ")").Text() // <<<
		groupMangas := make([]*vo.TinyManga, 0)
		sel.Find("li").Each(func(idx int, li *goquery.Selection) {
			groupMangas = append(groupMangas, m.getTinyMangaPageFromLi(li, true))
		})
		groups = append(groups, &vo.MangaGroup{
			Title:  groupTitle,
			Mangas: m.tinyMangasToTinyBlockMangas(groupMangas),
		})
	})

	// get other group mangas
	otherGroups := make([]*vo.MangaGroup, 0)
	if tagIndex == 1 || tagIndex == 2 {
		scContDiv := doc.Find("div.idx-sc-cont")
		scContDiv.Each(func(idx int, sel *goquery.Selection) {
			groupMangas := make([]*vo.TinyManga, 0)
			groupTitle := sel.Find("h4").Text()
			otherMangaUl := sel.Find("div.idx-sc-list ul:nth-child(" + xnumber.Itoa(tagIndex) + ")") // <<<
			otherMangaUl.Children().Each(func(idx int, li *goquery.Selection) {
				manga := m.getTinyMangaPageFromLi(li, false)
				manga.Finished = tagIndex == 2
				groupMangas = append(groupMangas, manga)
			})
			otherGroups = append(otherGroups, &vo.MangaGroup{
				Title:  groupTitle,
				Mangas: m.tinyMangasToTinyBlockMangas(groupMangas),
			})
		})
	} else if tagIndex == 3 {
		entity := func(mid uint64, finished bool, title string, chapter string) *vo.TinyBlockManga {
			return &vo.TinyBlockManga{
				Mid:           mid,
				Title:         title,
				Cover:         fmt.Sprintf(static.MANGA_COVER_URL, mid),
				Url:           fmt.Sprintf(static.MANGA_PAGE_URL, mid),
				Finished:      finished,
				NewestChapter: chapter,
			}
		}
		otherGroups = append(otherGroups, &vo.MangaGroup{
			Title: "推理/恐怖/悬疑",
			Mangas: []*vo.TinyBlockManga{
				// 推理
				entity(43991, true, "S-终极警官", "第10卷"),
				entity(43865, true, "异法人", "第03卷"),
				entity(43816, false, "小林少年和狂妄怪人", "第02卷"),
				entity(43815, false, "侦探少女有纱事件簿 来自沟口的爱", "第02卷"),
				entity(43550, false, "毒之樱", "第01卷"),
				// 恐怖
				entity(46675, true, "大叔狩猎", "短篇"),
				entity(46635, true, "被支配的行尸", "短篇"),
				entity(46420, true, "细菌", "短篇"),
				entity(46378, true, "万圣节", "短篇"),
				entity(46296, false, "死者的葬列", "第02话"),
				// 悬疑
				entity(46362, true, "怪医黑杰克：特别篇", "第01卷"),
				entity(46086, false, "我与骚扰狂", "s07话"),
				entity(45989, false, "裘格斯的二人", "第01话"),
				entity(45670, true, "电子保姆的纯情", "短篇"),
				entity(45651, true, "自杀挑战", "短篇"),
			},
		})
		otherGroups = append(otherGroups, &vo.MangaGroup{
			Title: "百合/后宫/治愈",
			Mangas: []*vo.TinyBlockManga{
				// 百合
				entity(46723, false, "邀你一起在户外共进美餐", "第01话"),
				entity(46712, true, "synergy", "短篇"),
				entity(46710, true, "day/life REmain", "短篇"),
				entity(46707, true, "oh! Chabashira", "短篇"),
				entity(46704, true, "二重身", "短篇"),
				// 后宫
				entity(46611, false, "虽然是召唤勇者，因为被认定为最下级，所以自己制造女仆后宫", "第01卷"),
				entity(45312, false, "星天的塔鲁克-帝国后宫秘史", "第01话"),
				entity(43990, false, "养女儿开后宫", "公告"),
				entity(39301, false, "靠着qs技能在异世界开无双", "第02话"),
				entity(39275, false, "麻烦不断的女仆们", "第06话"),
				// 治愈
				entity(46552, true, "HIGH SCHOOL RUNWAY", "短篇"),
				entity(46388, true, "人鱼海格", "短篇"),
				entity(46362, true, "怪医黑杰克：特别篇", "第01卷"),
				entity(46361, true, "怪医黑杰克画集", "第01卷"),
				entity(46294, false, "下北泽购物游记", "第02话"),
			},
		})
		otherGroups = append(otherGroups, &vo.MangaGroup{
			Title: "社会/历史/战争",
			Mangas: []*vo.TinyBlockManga{
				// 社会
				entity(46622, false, "元祖大四叠半大物语", "第02话"),
				entity(45339, true, "某一天系列", "黑心商人的靈機一動"),
				entity(45172, true, "我的夜晚比你的白天更美", "第02卷"),
				entity(45088, true, "链偶", "短篇"),
				entity(44866, true, "会长岛耕作", "第13卷"),
				// 历史
				entity(46053, false, "帝后轶闻", "第02话"),
				entity(45316, false, "超抢跑赵云", "第36话"),
				entity(45075, false, "古董屋优子", "第02卷"),
				entity(44857, false, "足下定江山", "第04话"),
				entity(44657, true, "剑豪与萩饼", "短篇"),
				// 战争
				entity(46203, false, "黎明军团", "第01卷"),
				entity(45620, false, "勇敢的正义公主琪琪", "第02话"),
				entity(45056, false, "过眼云烟的爱", "家政服务机器HAL02"),
				entity(44828, false, "星球大战：原力释放", "第01卷"),
				entity(44818, false, "使命召唤：幽灵", "第03卷"),
			},
		})
		otherGroups = append(otherGroups, &vo.MangaGroup{
			Title: "校园/励志/冒险",
			Mangas: []*vo.TinyBlockManga{
				// 校园
				entity(46725, false, "你其实是喜欢我的对吧？", "第01话"),
				entity(46720, true, "魔法事故", "短篇"),
				entity(46714, false, "莫斯科的早晨", "第02话"),
				entity(46697, false, "白根同学的告白", "短篇03"),
				entity(46683, true, "泥之龟", "短篇"),
				// 励志
				entity(45748, true, "Build a chair", "短篇"),
				entity(45204, true, "贵石", "短篇"),
				entity(45140, false, "Blue Period", "第57话"),
				entity(44971, true, "地下的天使", "短篇"),
				entity(44880, true, "○さわ@@的竹林组短漫", "短篇"),
				// 冒险
				entity(46717, false, "木叶新传：汤烟忍法帖", "第02话"),
				entity(46693, false, "勇者斗恶龙 达伊的大冒险 勇者阿邦和狱炎的魔王", "序"),
				entity(46656, false, "蓝甲虫：毕业日", "第01卷"),
				entity(46629, false, "弑神魔王转生成为最弱种族成就史上最强", "第04话"),
				entity(46618, false, "以破损技能开始的现代迷宫攻略", "第01话"),
			},
		})
	}

	return topGroup, groups, otherGroups
}

func (m *MangaListService) getTinyMangaPageFromLi(li *goquery.Selection, hasCover bool) *vo.TinyManga {
	if hasCover {
		url := li.Find("a").AttrOr("href", "")
		title := li.Find("a").AttrOr("title", "")
		cover := li.Find("a img").AttrOr("src", "")
		if cover == "" {
			cover = li.Find("a img").AttrOr("data-src", "")
		}
		tt := li.Find("span.tt").Text()
		newestChapter := strings.TrimSuffix(strings.TrimPrefix(strings.TrimPrefix(tt, "更新至"), "共"), "[全]")
		newestDate := li.Find("span.dt").Text()
		return &vo.TinyManga{
			Mid:           static.ParseMid(url),
			Title:         title,
			Cover:         static.ParseCoverUrl(cover),
			Url:           static.HOMEPAGE_URL + url,
			Finished:      strings.HasPrefix(tt, "共"),
			NewestChapter: newestChapter,
			NewestDate:    newestDate,
		}
	} else {
		title := li.Find("h6 a").AttrOr("title", "")
		url := li.Find("h6 a").AttrOr("href", "")
		newestChapter := li.Find("h6 span a").AttrOr("title", "")
		id := static.ParseMid(url)
		return &vo.TinyManga{
			Mid:           id,
			Title:         title,
			Cover:         fmt.Sprintf(static.MANGA_COVER_URL, id),
			Url:           static.HOMEPAGE_URL + url,
			Finished:      true, // true
			NewestChapter: newestChapter,
			NewestDate:    "",
		}
	}
}

func (m *MangaListService) tinyMangasToTinyBlockMangas(mangas []*vo.TinyManga) []*vo.TinyBlockManga {
	out := make([]*vo.TinyBlockManga, len(mangas))
	for idx, manga := range mangas {
		out[idx] = &vo.TinyBlockManga{
			Mid:           manga.Mid,
			Title:         manga.Title,
			Cover:         manga.Cover,
			Url:           manga.Url,
			Finished:      manga.Finished,
			NewestChapter: manga.NewestChapter,
		}
	}
	return out
}

func (m *MangaListService) GetHotSerialMangas() (*vo.MangaGroupList, error) {
	_, doc, err := m.httpService.HttpGetDocument(static.HOMEPAGE_URL, nil)
	if err != nil {
		return nil, err
	}

	topGroup, groups, otherGroups := m.getMangas(doc, 1, "serial")
	return &vo.MangaGroupList{
		Title:       "热门连载",
		TopGroup:    topGroup,
		Groups:      groups,
		OtherGroups: otherGroups,
	}, nil
}

func (m *MangaListService) GetFinishedMangas() (*vo.MangaGroupList, error) {
	_, doc, err := m.httpService.HttpGetDocument(static.HOMEPAGE_URL, nil)
	if err != nil {
		return nil, err
	}

	topGroup, groups, otherGroups := m.getMangas(doc, 2, "finish")
	return &vo.MangaGroupList{
		Title:       "经典完结",
		TopGroup:    topGroup,
		Groups:      groups,
		OtherGroups: otherGroups,
	}, nil
}

func (m *MangaListService) GetLatestMangas() (*vo.MangaGroupList, error) {
	_, doc, err := m.httpService.HttpGetDocument(static.HOMEPAGE_URL, nil)
	if err != nil {
		return nil, err
	}

	topGroup, groups, otherGroups := m.getMangas(doc, 3, "latest")
	return &vo.MangaGroupList{
		Title:       "最新上架",
		TopGroup:    topGroup,
		Groups:      groups,
		OtherGroups: otherGroups, // X
	}, nil
}

func (m *MangaListService) GetHomepageMangaGroupList() (*vo.HomepageMangaGroupList, error) {
	wg := &sync.WaitGroup{}
	wg.Add(2)

	var doc1 *goquery.Document
	var doc2 *goquery.Document
	var err1 error
	var err2 error
	go func() {
		defer func() { wg.Done() }()
		_, doc1, err1 = m.httpService.HttpGetDocument(static.HOMEPAGE_URL, nil)
	}()
	go func() {
		defer func() { wg.Done() }()
		_, doc2, err2 = m.httpService.HttpGetDocument(fmt.Sprintf(static.MANGA_RANK_URL, ""), nil)
	}()
	wg.Wait()
	if err1 != nil {
		return nil, err1
	}
	if err2 != nil {
		return nil, err2
	}

	sTopGroup, sGroups, sOtherGroups := m.getMangas(doc1, 1, "serial")
	fTopGroup, fGroups, fOtherGroups := m.getMangas(doc1, 2, "finish")
	lTopGroup, lGroups, lOtherGroups := m.getMangas(doc1, 3, "latest")
	daily := m.rankService.getRankingListFromDoc(doc2)
	categories := m.categoryService.getAllCategories(doc2)
	genres := categories.Genres
	zones := categories.Zones
	ages := categories.Ages

	return &vo.HomepageMangaGroupList{
		Serial: &vo.MangaGroupList{Title: "热门连载", TopGroup: sTopGroup, Groups: sGroups, OtherGroups: sOtherGroups},
		Finish: &vo.MangaGroupList{Title: "经典完结", TopGroup: fTopGroup, Groups: fGroups, OtherGroups: fOtherGroups},
		Latest: &vo.MangaGroupList{Title: "最新上架", TopGroup: lTopGroup, Groups: lGroups, OtherGroups: lOtherGroups},
		Daily:  daily,
		Genres: genres,
		Zones:  zones,
		Ages:   ages,
	}, nil
}

func (m *MangaListService) GetRecentUpdatedMangas(pa *param.PageParam) ([]*vo.TinyManga, int32, error) {
	_, doc, err := m.httpService.HttpGetDocument(static.MANGA_UPDATE_URL, nil)
	if err != nil {
		return nil, 0, err
	}

	latestLis := doc.Find("div.latest-list li")
	allMangas := make([]*vo.TinyManga, latestLis.Length())
	latestLis.Each(func(idx int, li *goquery.Selection) {
		allMangas[idx] = m.getTinyMangaPageFromLi(li, true)
	})
	totalLength := int32(len(allMangas))

	out := make([]*vo.TinyManga, 0)
	start := pa.Limit * (pa.Page - 1)
	end := start + pa.Limit
	for i := start; i < end && i < totalLength; i++ {
		out = append(out, allMangas[i])
	}
	return out, totalLength, nil
}
