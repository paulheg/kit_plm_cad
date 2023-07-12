package parser

import (
	"strconv"
	"strings"
)

type Strategy interface {
	SetAutoTilt(tilt bool)
	MoveUp()
	MoveDown()
	TiltUp()
	TiltDown()
	Move(x, y float32)
}

type Command string

const (
	AutoTilt Command = "auto_tilt"
	TiltUp   Command = "tilt_up"
	TiltDown Command = "tilt_down"
	MoveUp   Command = "move_up"
	MoveDown Command = "move_down"
	Move     Command = "move"
)

const Delimeter = "|"

func Parse(input string, strategy Strategy) {

	args := strings.Split(input, Delimeter)
	if len(args) == 0 {
		return
	}

	command := Command(args[0])
	args = args[1:]

	switch command {
	case MoveUp:
		if len(args) == 0 {
			strategy.MoveUp()
		}
	case MoveDown:
		if len(args) == 0 {
			strategy.MoveDown()
		}
	case AutoTilt:
		if len(args) == 1 {
			autoTilt, err := strconv.ParseBool(args[0])
			if err != nil {
				return
			}

			strategy.SetAutoTilt(autoTilt)
		}
	case TiltDown:
		if len(args) == 0 {
			strategy.TiltDown()
		}

	case TiltUp:
		if len(args) == 0 {
			strategy.TiltUp()
		}

	case Move:
		if len(args) == 2 {
			x, err := strconv.ParseFloat(args[0], 32)
			if err != nil {
				return
			}

			y, err := strconv.ParseFloat(args[1], 32)
			if err != nil {
				return
			}

			strategy.Move(float32(x), float32(y))
		}
	}

}
