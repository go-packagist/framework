package provider

type Provider interface {
	Register()
	Boot()
}
