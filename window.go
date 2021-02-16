package gwl

type Window interface {
	GetTitle() string
	SetTitle(title string) error

	GetBounds() Bounds
	SetBounds(bounds Bounds) error

	WaitForEvent() (Event, error)

	Close() error
}

func NewWindow(title string, bounds Bounds) (Window, error) {
	return createPlatformWindow(title, bounds)
}
