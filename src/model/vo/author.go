package vo

// 作者的部分信息 (Small)
type SmallAuthor struct {
	Aid        uint64 // 作者编号
	Name       string // 作者名称
	Zone       string // 作者地区
	Url        string // 作者链接
	MangaCount int32  // 漫画数量
	NewestDate string // 更新时间
}

// 作者的部分信息 (Tiny)
type TinyAuthor struct {
	Aid  uint64 // 作者编号
	Name string // 作者名称
	Url  string // 作者链接
}
