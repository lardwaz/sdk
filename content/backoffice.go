package sdk

//Content represents a content type
type Content struct {
	Icon        string `json:"icon"`
	Visibility  string `json:"visibility"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Path        string `json:"path"`
	Pagination  string `json:"pagination"`
}

//Field represents a field in a content type
type Field struct {
	Label string `json:"label"`
	Name  string `json:"name"`
	Edit  Widget `json:"edit"`
	List  Widget `json:"list"`
}

//Widget represents which widget
type Widget struct {
	Type       string `json:"type"`
	Visibility string `json:"visibility"`
}
