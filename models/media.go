package models

type Media struct {
	Model
	FileName string
	FileType string
}

func (m *Media) GetURL() string {
	return "/media/" + m.ID + "." + m.FileType
}
