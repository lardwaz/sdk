package content

//Contentable indicates that a content-type is being represented
//It provides meta data on the content type
type Contentable interface {
	Icon() string
	Visibility() string
	Label() string
	LabelField() string
	Description() string
	Path() string
	Pagination() string
	ContentBuilder() bool
}

//Content contains metadata on entity to allow them be treated as content types
type content struct {
	icon           string
	visibility     string
	label          string
	labelField     string
	description    string
	path           string
	pagination     string
	contentBuilder bool
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

//LabelField to use when the content type is displayed in the backoffice
func (c content) LabelField() string {
	return c.labelField
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

//ContentBuilder rules for listing of the content type
func (c content) ContentBuilder() bool {
	return c.contentBuilder
}

//New is a helper to initialize a content instance
func New(icon, visibility, label, labelField, description, path, pagination string, contentBuilder bool) Contentable {
	return content{
		icon:           icon,
		visibility:     visibility,
		label:          label,
		labelField:     labelField,
		description:    description,
		path:           path,
		pagination:     pagination,
		contentBuilder: contentBuilder,
	}
}
