package server

import (
	"github.com/datatogether/api/apiutil"
	"github.com/ipfs/go-datastore"
	"net/http"
)

func (s *Server) HandleIPFSPath(w http.ResponseWriter, r *http.Request) {
	data, err := s.store.Get(datastore.NewKey(r.URL.Path))
	if err != nil {
		apiutil.WriteErrResponse(w, http.StatusInternalServerError, err)
		return
	}

	w.Write(data)
}

// WebappHandler renders the home page
func WebappHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "webapp")
}