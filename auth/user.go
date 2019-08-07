package auth

// User contains the information common amongst most providers.
// All of the "raw" datafrom the provider can be found in the `RawData` field.
type User struct {
	RawData   map[string]interface{} `json:"raw_data"`
	Provider  string                 `json:"provider"` // Provider name
	Token     string                 `json:"-"`        // :sensitive:
	UserID    string                 `json:"user_id"`
	Username  string                 `json:"username"`
	Email     string                 `json:"email"`
	AvatarURL string                 `json:"avatar_url"`
	FirstName string                 `json:"first_name"`
	LastName  string                 `json:"last_name"`
	Location  string                 `json:"location"`
	Types     []string               `json:"types"`   // `Roles` for some Providers
	Actions   []string               `json:"actions"` // `Permissions` for some Providers
}
