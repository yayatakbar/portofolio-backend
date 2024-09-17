package models

// Struct untuk Data Diri
type Sidebar struct {
	Name    string `json:"name" bson:"name"`
	Link    string `json:"link" bson:"link"`
	IsAdmin int    `json:"is_admin" bson:"is_admin"`
}
