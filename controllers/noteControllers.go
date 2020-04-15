package controllers

import (
	"encoding/json"
	"net/http"
	models "personal_note/persistenseStore/postgresql"
	u "personal_note/utils"

	"github.com/gorilla/mux"
)

// CreateNote will create note using note id
var CreateNote = func(w http.ResponseWriter, r *http.Request) {

	user := r.Context().Value("user").(string) //Grab the id of the user that send the request

	note := &models.Note{}
	
	err := json.NewDecoder(r.Body).Decode(note)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	note.UserID = user
	resp := note.Create()
	u.Respond(w, resp)
}

// UpdateNote will update note using note id
var UpdateNote = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	note := &models.Note{}

	err := json.NewDecoder(r.Body).Decode(note)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	resp := note.Update(id)
	u.Respond(w, resp)
}

// DeleteNote will delete note using note id
var DeleteNote = func(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id := params["id"]
	
	resp := models.Delete(id)
	u.Respond(w, resp)
}

// ArchiveNote will archive note using note id
var ArchiveNote = func(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id := params["id"]
	
	resp := models.Archive(id, true)
	u.Respond(w, resp)
}

// UnArchiveNote will unarchive note using note id
var UnArchiveNote = func(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id := params["id"]
	
	resp := models.Archive(id, false)
	u.Respond(w, resp)
}

// GetArtchivedList will list of archived notes using user id
var GetArtchivedList = func(w http.ResponseWriter, r *http.Request) {

	userID := r.Context().Value("user").(string)
	data := models.GetArtchivedList(userID, true)
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

// GetUnArtchivedList will list of unarchived notes using user id
var GetUnArtchivedList = func(w http.ResponseWriter, r *http.Request) {

	userID := r.Context().Value("user").(string)
	data := models.GetArtchivedList(userID, false)
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

