package vo

type MangaPage struct {
}

type MangaChapter struct {
	Bid      uint64   `json:"bid"`
	Bname    string   `json:"bname"`
	Bpic     string   `json:"bpic"`
	Cid      uint64   `json:"cid"`
	Cname    string   `json:"cname"`
	Files    []string `json:"files"`
	Finished bool     `json:"finished"`
	Len      int32    `json:"len"`
	Path     string   `json:"path"`
	NextId   int32    `json:"nextId"`
	PrevId   int32    `json:"prevId"`
	Sl       *struct {
		E int64  `json:"e"`
		M string `json:"m"`
	} `json:"sl"`
}
