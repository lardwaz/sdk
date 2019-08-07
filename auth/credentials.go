package auth

// Credentials represents connection to a service having a protocol, identifier, password, host and port
type Credentials interface {
	GetProtocol() string
	GetIdentifier() string
	GetPassword() string
	GetHost() string
	GetPort() string
}
