package provider

type Provider interface {
	Register()
	Boot()
}

// UnimplementedProvider is a default implementation of the Provider interface.
type UnimplementedProvider struct {
}

// Register is a default implementation of the Provider interface.
func (u *UnimplementedProvider) Register() {
}

// Boot is a default implementation of the Provider interface.
func (u *UnimplementedProvider) Boot() {
}
