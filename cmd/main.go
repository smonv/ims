package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/solher/arangolite"
	"github.com/tthanh/ims/arango"
	"github.com/tthanh/ims/config"
	"github.com/tthanh/ims/server"
)

var (
	s *server.Server
	c *config.Config
)

func main() {
	c = config.LoadConfig()
	db := arangolite.New().LoggerOptions(false, false, false).
		Connect(fmt.Sprintf("http://%s:%s", c.Arango.Host, c.Arango.Port), "_system", "", "")

	arango.InitDatabase(db, c.Arango)

	imageStore := arango.NewImageStore(db)
	tagStore := arango.NewTagStore(db)
	imageTagStore := arango.NewImageTagStore(db)
	s = server.NewServer(imageStore, tagStore, imageTagStore)

	r := mux.NewRouter()
	r.HandleFunc("/", s.Home).Methods("GET")
	r.HandleFunc("/api/tags", s.CreateTag).Methods("POST")
	r.HandleFunc("/api/tags", s.GetTags).Methods("GET")
	r.HandleFunc("/api/tags/{key}", s.GetTag).Methods("GET")
	r.HandleFunc("/api/images/{key}", s.GetImage).Methods("GET")
	r.HandleFunc("/api/image_tag", s.CreateImageTag).Methods("POST")
	r.HandleFunc("/api/image_tag/images/{key}", s.GetImagesByTag).Methods("GET")
	r.HandleFunc("/api/image_tag/tags/{key}", s.GetTagsByImage).Methods("GET")

	http.ListenAndServe(":8080", r)
}
