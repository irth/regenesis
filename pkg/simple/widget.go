package simple

type Widget interface {
	Render() (string, error)
	Update(stdout Output) ([]BoundEventHandler, error)
}
