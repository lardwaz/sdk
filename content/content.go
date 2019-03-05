package content

//Content represents a content type
type Content interface {
	Icon() string
	Visibility() string
	Label() string
	Description() string
	Path() string
	Pagination() string
}

type content struct {
	icon        string
	visibility  string
	label       string
	description string
	path        string
	pagination  string
}

func (c content) Icon() string {
	return c.icon
}

func (c content) Visibility() string {
	return c.visibility
}

func (c content) Label() string {
	return c.label
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
func NewContent(icon, visibility, label, description, path, pagination string) Content {
	return content{
		icon:        icon,
		visibility:  visibility,
		label:       label,
		description: description,
		path:        path,
		pagination:  pagination,
	}
}
