package gwl

import (
	"fmt"
)

func createPlatformWindow(title string, bounds Bounds) (Window, error) {
	compositor, err := getCompositor()
	if err != nil {
		return nil, err
	}
	switch compositor {
	case "x11":
		w, err := createX11Window(title, bounds)
		return w, err
	default:
		return nil, fmt.Errorf("compositor (%s) not implemented", compositor)
	}
}
