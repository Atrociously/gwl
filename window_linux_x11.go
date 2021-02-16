package gwl

import (
	"fmt"

	"github.com/jezek/xgb"
	"github.com/jezek/xgb/xproto"
)

type x11window struct {
	x   *xgb.Conn
	wid xproto.Window
	ctx xproto.Gcontext
}

func createX11Window(title string, bounds Bounds) (*x11window, error) {
	x, err := xgb.NewConn()
	if err != nil {
		return nil, err
	}
	screen := xproto.Setup(x).DefaultScreen(x)

	wid, err := xproto.NewWindowId(x)
	if err != nil {
		return nil, err
	}
	create_cookie := xproto.CreateWindowChecked(x, screen.RootDepth, wid, screen.Root,
		int16(bounds.X), int16(bounds.Y), uint16(bounds.Width), uint16(bounds.Height), 0,
		xproto.WindowClassInputOutput, screen.RootVisual,
		xproto.CwBackPixel|xproto.CwEventMask,
		[]uint32{ // values must be in the order defined by the protocol
			screen.WhitePixel,
			xproto.EventMaskStructureNotify |
			xproto.EventMaskKeyPress |
			xproto.EventMaskKeyRelease |
			xproto.EventMaskSubstructureNotify |
			xproto.EventMaskExposure |
			xproto.EventMaskButtonPress})

	if err = create_cookie.Check(); err != nil {
		fmt.Println("window create error")
		return nil, err
	}
	map_cookie := xproto.MapWindowChecked(x, wid)
	if err := map_cookie.Check(); err != nil {
		fmt.Println("window map error")
		return nil, err
	}

	ctx, err := xproto.NewGcontextId(x)
	if err != nil {
		return nil, err
	}
	ctx_cookie := xproto.CreateGCChecked(x, ctx, xproto.Drawable(wid),
		xproto.GcForeground|xproto.GcGraphicsExposures,
		[]uint32{screen.BlackPixel, 0})
	if err = ctx_cookie.Check(); err != nil {
		fmt.Println("graphics context create error")
		return nil, err
	}

	w := &x11window{
		x:   x,
		wid: wid,
		ctx: ctx,
	}
	err = w.SetTitle(title)

	return w, err
}

func (w *x11window) GetTitle() string {
	cookie := xproto.GetProperty(w.x, false, w.wid, xproto.AtomWmName, xproto.AtomString, 0, 0)
	reply, err := cookie.Reply()
	if err != nil {
		return err.Error()
	}
	return string(reply.Value)
}

func (w *x11window) SetTitle(title string) error {
	cookie := xproto.ChangePropertyChecked(w.x, xproto.PropModeReplace, w.wid, xproto.AtomWmName, xproto.AtomString, 8, uint32(len(title)), []byte(title))
	err := cookie.Check()
	return err
}

func (w *x11window) GetBounds() Bounds {
	cookie := xproto.GetGeometry(w.x, xproto.Drawable(w.wid))
	reply, err := cookie.Reply()
	if err != nil {
		return Bounds{}
	}
	return B16(reply.X, reply.Y, reply.Width, reply.Height)
}

func (w *x11window) SetBounds(bounds Bounds) error {
	cookie := xproto.ConfigureWindowChecked(w.x, w.wid,
		xproto.ConfigWindowX|xproto.ConfigWindowY|xproto.ConfigWindowWidth|xproto.ConfigWindowHeight,
		[]uint32{uint32(bounds.X), uint32(bounds.Y), uint32(bounds.Width), uint32(bounds.Height)})
	if err := cookie.Check(); err != nil {
		return err
	}
	_, err := cookie.Reply()
	return err
}

func (w *x11window) WaitForEvent() (Event, error) {
	ev, err := w.x.WaitForEvent()
	if err != nil {
		return nil, err
	}
	if ev == nil {
		return nil, nil
	}
	event := eventToEvent(ev)
	if event == nil {
		return nil, fmt.Errorf("%s", ev.String())
	}
	return event, nil
}

func (w *x11window) Close() error {
	cookie := xproto.DestroyWindowChecked(w.x, w.wid)
	err := cookie.Check()
	return err
}

// TODO implement this
func eventToEvent(ev xgb.Event) Event {
	switch event := ev.(type) {
	case xproto.ButtonPressEvent:
		return EventButton{
			Key: Key(event.Detail),
		}
	case xproto.KeyPressEvent:
		return EventButton{
			Key: Key(event.Detail),
		}
	default:
		return nil
	}
}
