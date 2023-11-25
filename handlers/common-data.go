package handlers

type PageData struct {
	Title string
}

type BlogPageData struct {
	Title     string
	Blogs     []Blog
	BlogCount int
}
