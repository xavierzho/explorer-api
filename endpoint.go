package explorer

type Endpoint interface {
	// Name returns the module name
	Name() string
}
