package handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
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
		ImageUri: "https://example.com/coding.png",
		Date:     "Mar 02 2021",
		ReadTime: "8 mins read",
		Href:     "/blogs/2",
	},
	{
		Id:       "3",
		Topic:    "Web Development",
		Title:    "Building Web Apps with Go",
		Intro:    "Create web applications using the Go programming language.",
		ImageUri: "https://example.com/webapp.png",
		Date:     "Jan 11 2021",
	},
	{
		Id:       "4",
		Topic:    "DevOps",
		Title:    "Container Orchestration with Go",
		Intro:    "Manage containers efficiently with Go.",
		ImageUri: "https://example.com/devops.png",
		Date:     "Jan 11 2021",
	},
	{
		Id:       "5",
		Topic:    "Data Science",
		Title:    "Data Analysis in Go",
		Intro:    "Use Go for data analysis and visualization.",
		ImageUri: "https://example.com/datascience.png",
		Date:     "Jan 11 2021",
	},
	{
		Id:       "6",
		Topic:    "Security",
		Title:    "Go Security Best Practices",
		Intro:    "Secure your Go applications against threats.",
		ImageUri: "https://example.com/security.png",
		Date:     "Jan 11 2021",
	},
}

func ServeBlog(w http.ResponseWriter, r *http.Request, logger *logger.Logger) {

	logger.Debug("Accessed route: '/blog'")

	tmpl, err := template.ParseFS(content, "templates/*.html", "templates/common/*.html")
	if err != nil {
		logger.Error(fmt.Sprintf("Error rendering HTML template: %s", err))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	data := BlogPageData{
		Title:     "Sys.D Solutions - Blogs",
		Blogs:     blogs[0:3],
		BlogCount: len(blogs),
	}

	w.WriteHeader(http.StatusOK)

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
	data := BlogPageData{
		Title:     fmt.Sprintf("Sys.D Solutions - %s", blogs[blogIndex].Title),
		Blogs:     blogs[0:3],
		BlogCount: len(blogs),
	}

	w.WriteHeader(http.StatusOK)

	err = tmpl.ExecuteTemplate(w, fmt.Sprintf("blog-%s.html", blogID), data)
	if err != nil {
		logger.Error(fmt.Sprintf("Error rendering HTML template: %s", err))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

}
