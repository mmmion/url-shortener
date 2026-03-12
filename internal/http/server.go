package http

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"text/template"
	"url_shortner/internal/shortener"
)

type URL shortener.URL

type Server struct {
	Address string
	Port    int
}

func NewServer(addr string, port int) (*Server, error) {
	if port <= 0 || port >= 65535 {
		return nil, fmt.Errorf("Invalid port.")
	}

	return &Server{
		Address: addr,
		Port:    port,
	}, nil
}

func (server *Server) RedirectHandler(w http.ResponseWriter, r *http.Request) {
	shortUrl := r.URL.Path[1:]

	fmt.Println("SHORT URL: " + shortUrl)

	url, exists := shortener.Find(shortUrl)

	fmt.Println(exists)
	fmt.Println(url)

	if !exists {
		type PageData struct {
			ErrorMsg string
			Code     string
		}
		data := PageData{
			ErrorMsg: "URL not found!",
			Code:     shortUrl,
		}

		tmpl, err := template.ParseFiles("internal/http/pages/index.html")
		if err != nil {
			http.Error(w, "Template missing", 500)
			return
		}

		tmpl.Execute(w, data)
		return
	}

	http.Redirect(w, r, url.Original, http.StatusSeeOther)
}

func (server *Server) ShortenURLHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var requestData struct {
		LongURL string `json:"long_url"`
	}

	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil || requestData.LongURL == "" {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	urlObj, err := shortener.NewURL(requestData.LongURL)

	if err != nil {
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(urlObj)
}

func (server *Server) Start() {
	fullAddr := fmt.Sprintf("%s:%d", server.Address, server.Port)

	fmt.Printf("Server starting on %s...\n", fullAddr)

	http.HandleFunc("/", server.RedirectHandler)
	http.HandleFunc("/api/shorten", server.ShortenURLHandler)

	log.Fatal(http.ListenAndServe(fullAddr, nil))
}
