package content

//Content represents a content type
type Content interface {
	Icon() string
	Visibility() string
	Name() string
	Description() string
	Path() string
	Pagination() string
}

type content struct {
	icon        string `json:"icon"`
	visibility  string `json:"visibility"`
	name        string `json:"name"`
	description string `json:"description"`
	path        string `json:"path"`
	pagination  string `json:"pagination"`
}

func (c content) Icon() string {
	return c.icon
}

func (c content) Visibility() string {
	return c.visibility
}

func (c content) Name() string {
	return c.name
}

func (c content) Description() string {
	return c.description
}

func (c content) Path() string {
	return c.path
}

func (c content) Pagination() string {
	return c.pagination
}

//NewContent is a helper to initialize a content instance
func NewContent(Icon, Visibility, Name, Description, Path, Pagination string) Content {
	return content{
		icon:        Icon,
		visibility:  Visibility,
		name:        Name,
		description: Description,
		path:        Path,
		pagination:  Pagination,
	}
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
