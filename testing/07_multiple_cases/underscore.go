package color

import "errors"

func ColorNew(name string) (string , error) {
	switch name {
	case "blue":
		return "#0000FF", nil
	case "white":
		return "#FFFFFF", nil
	case "black":
		return "#000000", nil
	case "grey":
		return "#888888", nil
	default:
		return "", errors.New("error")
	}
}