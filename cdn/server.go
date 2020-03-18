package cdn

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

const (
	defaultStaticDirName = "assets"
	defaultPort          = "8080"
)

// Server represents a cdn http server
type Server struct {
	ContentRoot string
	Port        string
	AuthToken   string
}

// NewHTTPServer creates a new cdn http server
func NewHTTPServer(contentRoot string, port string, authToken string) *Server {
	if contentRoot == "" {
		contentRoot = defaultStaticDirName
	}
	contentRoot = "./" + contentRoot + "/"
	if port == "" {
		port = defaultPort
	}
	return &Server{ContentRoot: contentRoot, Port: port, AuthToken: authToken}
}

// Serve will start the cdn service
func (s *Server) Serve() {
	router := s.newRouter()
	err := http.ListenAndServe(":"+s.Port, applyHandlers(router))
	if err != nil {
		log.Fatal("ListenAndServe Error: ", err)
	}
}

func (s *Server) newRouter() *mux.Router {
	router := mux.NewRouter()

	if _, err := os.Stat(s.ContentRoot); os.IsNotExist(err) {
		os.Mkdir(s.ContentRoot, os.ModeDir)
	}

	router.NewRoute().Handler(http.StripPrefix("/", http.FileServer(http.Dir(s.ContentRoot)))).Methods("GET")
	router.HandleFunc("/assets", s.applyAuthentication(s.create)).Methods("POST", "PUT")
	router.HandleFunc("/assets/{name}", s.applyAuthentication(s.delete)).Methods("DELETE")
	router.HandleFunc("/assets/{name}", s.applyAuthentication(s.replace)).Methods("PATCH")

	return router
}

func (s *Server) applyAuthentication(fun func(w http.ResponseWriter, r *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if s.AuthToken != "" {
			auth := r.Header.Get("Authorization")
			if auth == "" {
				unauthorized("missing authentication header", w)
				return
			} else if auth != s.AuthToken {
				forbidden("invalid authentication", w)
				return
			}
		}
		fun(w, r)
	}
}

func (s *Server) create(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(32 << 20)
	file, header, err := r.FormFile("file")
	if err != nil {
		if err == http.ErrMissingFile {
			badRequest("missing file to upload", w)
		} else {
			internalServerError("could not read file", w)
		}
		return
	}
	name := header.Filename
	b, e := strconv.ParseBool(r.FormValue("generate_name"))
	if e == nil || b == true {
		name = uuid.New().String() + filepath.Ext(header.Filename)
	}
	defer file.Close()
	_, err = os.Stat(s.ContentRoot + name)
	if err == nil {
		badRequest("file with this name already exists", w)
		return
	}
	f, err := os.OpenFile(s.ContentRoot+name, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		internalServerError("could not create file", w)
		return
	}
	defer f.Close()
	io.Copy(f, file)
	if b == true {
		m, _ := json.Marshal(NewUploadUUIDResponse(name))
		okWithContent(m, w)
	} else {
		noContent(w)
	}
}

func (s *Server) replace(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(32 << 20)
	vars := mux.Vars(r)
	name := vars["name"]
	file, _, err := r.FormFile("file")
	if err != nil {
		if err == http.ErrMissingFile {
			badRequest("missing file to replace", w)
		} else {
			internalServerError("could not read file", w)
		}
		return
	}
	defer file.Close()
	_, err = os.Stat(s.ContentRoot + name)
	if os.IsNotExist(err) {
		badRequest("file with this name doesn't exists", w)
		return
	}
	f, err := os.OpenFile(s.ContentRoot+name, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		internalServerError("could not open file", w)
		return
	}
	defer f.Close()
	io.Copy(f, file)
	noContent(w)
}

func (s *Server) delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	path := s.ContentRoot + vars["name"]
	_, err := os.Stat(path)
	if err != nil {
		badRequest("file with this name doesn't exist", w)
		return
	}
	err = os.Remove(path)
	if err != nil {
		internalServerError("could not remove file", w)
		return
	}
	noContent(w)
}
