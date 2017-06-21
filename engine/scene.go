package engine

// Scene is a generic interface for types which interact with SDLEngine. They
// may resize the game window, make draw calls, interact with the event loop,
// and more.
type Scene interface {
	Setup() error
	Main() error
	Teardown() error
}
