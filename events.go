package gwl

const (
	eventButtonType = 0x00
)

type Event interface {
	Type() uint8
}

type EventButton struct {
	Key Key
	Pressed bool
	Released bool
}

func (eb EventButton) Type() uint8 { return eventButtonType }
