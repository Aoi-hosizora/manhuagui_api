package vo

// 漫画页的完整信息
type MangaPage struct {
	Mid               uint64               // 漫画编号
	Title             string               // 漫画标题
	Cover             string               // 漫画封面
	Url               string               // 漫画链接
	PublishYear       string               // 出品年代
	MangaZone         string               // 漫画地区
	AlphabetIndex     string               // 字母索引
	Genres            []*Category          // 漫画剧情
	Authors           []*TinyAuthor        // 漫画作者
	Alias             string               // 漫画别名
	Finished          bool                 // 是否完结
	NewestChapter     string               // 最新一话
	NewestDate        string               // 更新时间
	BriefIntroduction string               // 简介介绍
	Introduction      string               // 漫画介绍
	MangaRank         string               // 漫画排名
	AverageScore      float32              // 平均给分
	ScoreCount        int32                // 给分人数
	PerScores         [6]float32           // 具体给分
	ChapterGroups     []*MangaChapterGroup // 章节链接
}

// 漫画章节的完整信息
type MangaChapter struct {
	Cid        uint64   `json:"cid"`      // 章节编号
	Title      string   `json:"cname"`    // 章节标题
	Mid        uint64   `json:"bid"`      // 漫画编号
	MangaTitle string   `json:"bname"`    // 漫画标题
	Cover      string   `json:"bpic"`     // 漫画封面
	Url        string   `json:"url"`      // 章节链接
	Pages      []string `json:"files"`    // 每页链接
	Finished   bool     `json:"finished"` // 是否完结
	PageCount  int32    `json:"len"`      // 总共页数
	Path       string   `json:"path"`     // 链接路径
	NextId     int32    `json:"nextId"`   // 下一章节
	PrevId     int32    `json:"prevId"`   // 上一章节
	Sl         *struct {
		E int64  `json:"e"` // e
		M string `json:"m"` // m
	} `json:"sl"` // 查询加密
}

// 漫画页的部分信息 (Tiny)
type TinyMangaPage struct {
	Mid           uint64 // 漫画编号
	Title         string // 漫画标题
	Cover         string // 漫画封面
	Url           string // 漫画链接
	Finished      bool   // 是否完结
	NewestChapter string // 最新一话
	NewestDate    string // 更新时间
}

// 漫画章节的部分信息 (Tiny)
type TinyMangaChapter struct {
	Cid       uint64 // 章节编号
	Title     string // 章节标题
	Mid       uint64 // 漫画编号
	Url       string // 章节链接
	PageCount int32  // 章节页数
	IsNew     bool   // 最近发布
}

// 漫画页分组 (Group)
type MangaPageGroup struct {
	Title  string           // 分组标题
	Mangas []*TinyMangaPage // 章节集合
}

// 漫画章节分组 (Group)
type MangaChapterGroup struct {
	Title    string              // 分组标题
	Chapters []*TinyMangaChapter // 章节集合
}

// 主页的漫画列表
type MangaPageGroupList struct {
	Title       string            // 列表标题
	TopGroup    *MangaPageGroup   // 置顶分组
	Groups      []*MangaPageGroup // 类别分组
	OtherGroups []*MangaPageGroup // 其他分组
}
