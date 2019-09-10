package generator

type IndexPage struct {
	BlogTitle string
	Author    string
	Posts     []*Post
	Theme     string
	Descr     string
}
