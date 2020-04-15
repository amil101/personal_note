package models

import (
	u "personal_note/utils"

	"github.com/jinzhu/gorm"
)

// basic structure for personal note
type Note struct {
	gorm.Model
	id      string // note Id
	Content string `json:"content"`
	Archive bool   `json:"archive"` // note Archive Status
	UserID  string `json:"user_id"` // The user that this Note belongs to
}

/*
 This struct function validate the required parameters sent through the http request body

returns message and true if the requirement is met
*/
func (note *Note) Validate() (map[string]interface{}, bool) {

	if note.Content == "" {
		return u.Message(false, "Note should be on the payload"), false
	}

	//All the required parameters are present
	return u.Message(true, "success"), true
}

// Create a note 
func (note *Note) Create() map[string]interface{} {

	if resp, ok := note.Validate(); !ok {
		return resp
	}

	GetDB().Create(note)

	resp := u.Message(true, "success")
	resp["note"] = note
	return resp
}

// Update a note using note id
func (note *Note) Update(id string) map[string]interface{} {

	if resp, ok := note.Validate(); !ok {
		return resp
	}

	err := GetDB().Model(&note).Where("ID = ?", id).UpdateColumn("content", note.Content).Error
	if err != nil {
		return nil
	}

	resp := u.Message(true, "success")
	return resp
}

// Archive a note using note id
func Archive(id string, flag bool) map[string]interface{} {

	note := &Note{}

	err := GetDB().Model(&note).Where("ID = ?", id).UpdateColumn("archive", flag).Error
	if err != nil {
		return nil
	}

	resp := u.Message(true, "success")
	return resp
}

// Delete a note using note id
func Delete(id string) map[string]interface{} {

	note := &Note{}

	err := GetDB().Where("ID = ?", id).Delete(note).Error
	if err != nil {
		return nil
	}

	resp := u.Message(true, "success")
	return resp
}

// get archived / unarchived list of notes accoring to the user
func GetArtchivedList(user string, flag bool) []*Note {

	notes := make([]*Note, 0)

	err := GetDB().Table("notes").Where("user_id = ? AND archive = ?", user, flag).Find(&notes).Error
	if err != nil {
		return nil
	}
	return notes
}
