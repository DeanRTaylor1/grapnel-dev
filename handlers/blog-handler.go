package handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"sort"
	"strconv"

	"github.com/DeanRTaylor1/deans-site/logger"
)

type Blog struct {
	Id       string
	Topic    string
	Title    string
	Intro    string
	ImageUri string
	Date     string
	ReadTime string
	Href     string
}

var blogs = []Blog{
	{
		Id:       "1",
		Topic:    "Marketing",
		Title:    "Unlocking Business Success: The Power of SEO in the Digital Age",
		Intro:    "In today's digital age, where the online marketplace is more competitive than ever, businesses are constantly seeking ways to stand out and thrive. One of the most effective tools in their arsenal is Search Engine Optimization (SEO). ",
		ImageUri: "/images/seo-image.jpg",
		Date:     "Jan 11 2021",
		ReadTime: "6 mins read",
		Href:     "/blogs/1",
	},
	{
		Id:       "2",
		Topic:    "Software",
		Title:    "Tailored to Thrive: The Advantages of Custom Software Solutions",
		Intro:    "In today's fast-paced business landscape, staying competitive often means embracing technology to streamline operations, enhance productivity, and deliver exceptional user experiences...",
		ImageUri: "/images/software.jpg",
		Date:     "Mar 02 2021",
		ReadTime: "8 mins read",
		Href:     "/blogs/2",
	},
	{
		Id:       "3",
		Topic:    "Ecommerce",
		Title:    "Elevate Your Online Store: The Magic of Ecommerce Integrations",
		Intro:    "            In the fast-paced world of online retail, staying competitive and delivering exceptional user experiences is crucial...",
		ImageUri: "/images/ecommerce.jpg",
		Date:     "Oct 08 2021",
		ReadTime: "7 mins read",
		Href:     "/blogs/3",
	},
	{
		Id:       "4",
		Topic:    "Business",
		Title:    " The ROI of Custom Web Development: Why Invest in a Bespoke Website?",
		Intro:    "Creating a custom website, tailored specifically to a brand's unique needs and identity, can significantly enhance a business's online presence and profitability...",
		ImageUri: "/images/bespoke.jpg",
		Date:     "Jan 14 2022",
		ReadTime: "8 mins read",
		Href:     "/blogs/4",
	},
	{
		Id:       "5",
		Topic:    "SEO",
		Title:    " Cracking the Code: How SEO and Custom Web Development Go Hand in hand",
		Intro:    " In the digital world, the duo of SEO (Search Engine Optimization) and custom web development is like bread and butter...",
		ImageUri: "/images/seo-image.jpg",
		Date:     "Feb 17 2022",
		ReadTime: "6 mins read",
		Href:     "/blogs/5",
	},
	{
		Id:       "6",
		Topic:    "Case Study",
		Title:    "Case Study: Elevating Nhimsallyfilm.com with Customized Software and SEO Expertise",
		Intro:    " In the digital era, tailored software and strategic SEO have become vital for businesses to...",
		ImageUri: "/images/nhimsallyfilm.jpg",
		Date:     "Mar 21 2023",
		ReadTime: "8 mins read",
		Href:     "/blogs/6",
	},
}

func ServeBlog(w http.ResponseWriter, r *http.Request, logger *logger.Logger) {

	logger.Debug("Accessed route: '/blog'")

	tmpl, err := template.ParseFS(content, "templates/*.html", "templates/common/*.html")
	if err != nil {
		logger.Error(fmt.Sprintf("Error rendering HTML template: %s", err))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	sort.Slice(blogs, func(i, j int) bool {
		return blogs[i].Id > blogs[j].Id
	})

	data := BlogPageData{
		Title:     "Grapnel - Blogs",
		Blogs:     blogs[0:3],
		BlogCount: len(blogs),
	}

	w.Header().Set("Content-Type", ContentTypeHTML)

	err = tmpl.ExecuteTemplate(w, "blog.html", data)
	if err != nil {
		logger.Error(fmt.Sprintf("Error rendering HTML template: %s", err))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func GetBlogs(w http.ResponseWriter, r *http.Request, logger *logger.Logger) {
	pageParam := r.URL.Query().Get("page")

	page := 1
	if pageParam != "" {
		page, _ = strconv.Atoi(pageParam)
	}

	postsPerPage := 3

	startIndex := (page - 1) * postsPerPage
	if startIndex < 0 || startIndex >= len(blogs) {
		http.Error(w, "Page number out of range", http.StatusBadRequest)
		return
	}

	endIndex := startIndex + postsPerPage
	if endIndex > len(blogs) {
		endIndex = len(blogs)
	}

	pageBlogs := blogs[startIndex:endIndex]

	jsonData, err := json.Marshal(pageBlogs)
	if err != nil {
		logger.Error(fmt.Sprintf("Error marshaling JSON: %s", err))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func GetBlogByID(w http.ResponseWriter, r *http.Request, logger *logger.Logger, blogID string) {
	logger.Debug("Accessed route: '/blog/id'")

	tmpl, err := template.ParseFS(content, "templates/*.html", "templates/common/*.html")
	if err != nil {
		logger.Error(fmt.Sprintf("Error rendering HTML template: %s", err))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	blogIndex, err := strconv.Atoi(blogID)
	if err != nil {
		http.Error(w, "Something went wrong.", http.StatusBadRequest)
	}
	data := PageData{
		Title: fmt.Sprintf("Grapnel - %s", blogs[blogIndex-1].Title),
	}

	w.Header().Set("Content-Type", ContentTypeHTML)

	err = tmpl.ExecuteTemplate(w, fmt.Sprintf("blog-%s.html", blogID), data)
	if err != nil {
		logger.Error(fmt.Sprintf("Error rendering HTML template: %s", err))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

}
