package robot

import (
	"log"

	"github.com/paulheg/kit_plm_cad/robot/internal/parser"
)

var _ parser.Strategy = &LoggerStrategy{}

type LoggerStrategy struct {
}

// TiltDown implements parser.Strategy.
func (*LoggerStrategy) TiltDown() {
	log.Print("Tilt Down")
}

// TiltUp implements parser.Strategy.
func (*LoggerStrategy) TiltUp() {
	log.Print("Tilt Up")
}

// MoveDown implements parser.Strategy.
func (*LoggerStrategy) MoveDown() {
	log.Print("Move Down")
}

// MoveUp implements parser.Strategy.
func (*LoggerStrategy) MoveUp() {
	log.Print("Move Up")
}

// Move implements Strategy.
func (*LoggerStrategy) Move(x float32, y float32) {
	left, right := JoystickFixFlutterData(JoystickRotationStrategy)(x, y)
	log.Printf("X: %f, Y: %f", x, y)
	log.Printf("Left: %f, Right: %f", left, right)
}

// SetAutoTilt implements Strategy.
func (*LoggerStrategy) SetAutoTilt(tilt bool) {
	log.Printf("Auto Tilt: %t", tilt)
}

// SetHeight implements Strategy.
func (*LoggerStrategy) SetHeight(height float32) {
	log.Printf("Height: %f", height)
}

// SetTilt implements Strategy.
func (*LoggerStrategy) SetTilt(tilt float32) {
	log.Printf("Tilt: %f", tilt)
}
