package content

//Content contains metadata on entity to allow them be treated as content types
type Content struct {
	icon        string
	visibility  string
	label       string
	description string
	path        string
	pagination  string
}

//Icon to represent the content type in backoffice
func (c Content) Icon() string {
	return c.icon
}

//Visibility rules for the content type
func (c Content) Visibility() string {
	return c.visibility
}

//Label to use when the content type is displayed in the backoffice
func (c Content) Label() string {
	return c.label
}

//Description of the content type
func (c Content) Description() string {
	return c.description
}

//Path to full view for the content type
func (c Content) Path() string {
	return c.path
}

//Pagination rules for listing of the content type
func (c Content) Pagination() string {
	return c.pagination
}

//New is a helper to initialize a content instance
func New(icon, visibility, label, description, path, pagination string) Content {
	return Content{
		icon:        icon,
		visibility:  visibility,
		label:       label,
		description: description,
		path:        path,
		pagination:  pagination,
	}
}
