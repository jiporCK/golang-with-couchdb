package entity

type Course struct {
	ID			string `json:"_id,omitempty"`
	Rev			string `json:"_rev,omitempty"`
	Name		string `json:"name"` 
	TeacherID	string `json:"teacher_id"`
}