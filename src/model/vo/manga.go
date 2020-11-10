package vo

type MangaPage struct {
	Bid           uint64               // 漫画编号
	Bname         string               // 漫画标题
	Bpic          string               // 漫画封面
	Url           string               // 漫画链接
	PublishYear   string               // 出品年代
	Zone          string               // 漫画地区
	AlphabetIndex string               // 字母索引
	Type          string               // 漫画剧情
	AuthorName    string               // 漫画作者
	Alias         string               // 漫画别名
	Finished      bool                 // 是否完结
	NewestChapter string               // 最新一话
	NewestDate    string               // 更新时间
	Introduction  string               // 漫画介绍
	Rank          string               // 漫画排名
	Chapters      []*MangaChapterGroup // 章节链接
}

type MangaChapterGroup struct {
	Title string              // 分组标题
	Links []*MangaChapterLink // 章节集合
}

type MangaChapterLink struct {
	Cid   uint64 // 章节编号
	Cname string // 章节标题
	Url   string // 章节链接
	Pages int32  // 章节页数
	New   bool   // 最近发布
}

type MangaChapter struct {
	Bid      uint64   `json:"bid"`      // 漫画编号
	Bname    string   `json:"bname"`    // 漫画标题
	Bpic     string   `json:"bpic"`     // 漫画封面
	Cid      uint64   `json:"cid"`      // 章节编号
	Cname    string   `json:"cname"`    // 章节标题
	Url      string   `json:"url"`      // 章节链接
	Files    []string `json:"files"`    // 每页链接
	Finished bool     `json:"finished"` // 是否完结
	Len      int32    `json:"len"`      // 总共页数
	Path     string   `json:"path"`     // 链接路径
	NextId   int32    `json:"nextId"`   // 下一章节
	PrevId   int32    `json:"prevId"`   // 上一章节
	Sl       *struct {
		E int64  `json:"e"` // e
		M string `json:"m"` // m
	} `json:"sl"` // 查询加密
}
