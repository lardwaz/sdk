package auth

// Endpoints is a list of common provider endpoints
type Endpoints struct {
	Login          string `json:"login"`
	Logout         string `json:"logout"`
	Profile        string `json:"profile"`
	ForgetPassword string `json:"forget_password"`
	Register       string `json:"register"`
}
