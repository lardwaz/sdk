package content

//Contentable indicates that a content-type is being represented
//It provides meta data on the content type
type Contentable interface {
	Icon() string
	Visibility() string
	Label() string
	LabelField() string
	Description() string
	Type() string
	Pagination() int
}

//Content contains metadata on entity to allow them be treated as content types
type content struct {
	icon           string
	visibility     string
	label          string
	labelField     string
	description    string
	entityType     string
	pagination     int
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

//Type of the content type
func (c content) Type() string {
	return c.entityType
}

//Pagination rules for listing of the content type
func (c content) Pagination() int {
	return c.pagination
}

//New is a helper to initialize a content instance
func New(icon, visibility, label, labelField, description, entityType string, pagination int) Contentable {
	return content{
		icon:           icon,
		visibility:     visibility,
		label:          label,
		labelField:     labelField,
		description:    description,
		entityType:     entityType,
		pagination:     pagination,
	}
}
