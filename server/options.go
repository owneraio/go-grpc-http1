package server

type options struct {
	preferGRPCWeb  bool
	genericMethods bool
}

// Option is an object that controls the behavior of the downgrading gRPC server.
type Option interface {
	apply(o *options)
}

type optionFunc func(o *options)

func (f optionFunc) apply(o *options) {
	f(o)
}

// PreferGRPCWeb instructs the server to always send a downgraded gRPC-Web response
// if the client indicates it accepts gRPC-Web responses.
func PreferGRPCWeb(prefer bool) Option {
	return optionFunc(func(o *options) {
		o.preferGRPCWeb = prefer
	})
}

// WithGenericMethods enables the downgradable mode for every method
func WithGenericMethods() Option {
	return optionFunc(func(o *options) {
		o.genericMethods = true
	})
}
