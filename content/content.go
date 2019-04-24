package content

//Contentable indicates that a content-type is being represented
//It provides meta data on the content type
type Contentable interface {
	Icon() string
	Visibility() string
	Label() string
	Description() string
	Path() string
	Pagination() string
}

//Content contains metadata on entity to allow them be treated as content types
type content struct {
	icon        string
	visibility  string
	label       string
	description string
	path        string
	pagination  string
}

//Icon to represent the content type in backoffice
func (c content) Icon() string {
	return c.icon
}

//Visibility rules for the content type
func (c content) Visibility() string {
	return c.visibility
}

//Label to use when the content type is displayed in the backoffice
func (c content) Label() string {
	return c.label
}

//Description of the content type
func (c content) Description() string {
	return c.description
}

//Path to full view for the content type
func (c content) Path() string {
	return c.path
}

//Pagination rules for listing of the content type
func (c content) Pagination() string {
	return c.pagination
}

//New is a helper to initialize a content instance
func New(icon, visibility, label, description, path, pagination string) Contentable {
	return content{
		icon:        icon,
		visibility:  visibility,
		label:       label,
		description: description,
		path:        path,
		pagination:  pagination,
	}
}
