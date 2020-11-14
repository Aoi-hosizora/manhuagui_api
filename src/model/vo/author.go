package vo

// 作者的完整信息
type Author struct {
	Aid  uint64 `json:"aid"`
	Name string `json:"name"`
	Url  string `json:"url"`
}

// 作者的部分信息 (Tiny)
type TinyAuthor struct {
	Aid  uint64 `json:"aid"`
	Name string `json:"name"`
	Url  string `json:"url"`
}
