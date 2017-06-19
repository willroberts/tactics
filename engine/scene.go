package engine

type Scene interface {
	Setup() error
	Main() error
	Teardown() error
}
