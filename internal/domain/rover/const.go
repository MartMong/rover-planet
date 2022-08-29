package rover

type direction float64

const (
	N direction = iota * 0.5
	NE
	E
	SE
	S
	SW
	W
	NW
)

func getDirection(d direction) string {
	switch d {
	case N:
		return "N"

	case NE:
		return "NE"

	case E:
		return "E"

	case SE:
		return "SE"

	case S:
		return "S"

	case SW:
		return "SW"

	case W:
		return "W"

	case NW:
		return "NW"

	default:
		return ""
	}
}
