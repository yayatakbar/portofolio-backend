package models

// Struct untuk Data Diri
type PersonalInfo struct {
	Name       string `json:"name" bson:"name"`
	BirthPlace string `json:"birth_place" bson:"birth_place"`
	BirthDate  string `json:"birth_date" bson:"birth_date"`
	Gender     string `json:"gender" bson:"gender"`
	Address    string `json:"address" bson:"address"`
	Hp         string `json:"hp" bson:"hp"`
	Email      string `json:"email" bson:"email"`
	Github     string `json:"github" bson:"github"`
	Facebook   string `json:"facebook" bson:"facebook"`
	Instagram  string `json:"instagram" bson:"instagram"`
	Twitter    string `json:"twitter" bson:"twitter"`
}
