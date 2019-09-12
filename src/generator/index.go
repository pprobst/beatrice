package generator

type IndexPage struct {
	BlogTitle string
	Author    string
	Posts     []*Post
    About     *Post
	Theme     string
	Descr     string
}
