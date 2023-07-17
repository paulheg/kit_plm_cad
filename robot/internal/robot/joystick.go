package robot

import "math"

type JoystickStrategy func(x, y float32) (left, right float32)

// Fix the inputs of the flutter joystick
// Input
// (-0.7,-0.7)	(0,-1) 	(0.7,-0.7)
// (-1,0)		(0,0) 	(1,0)
// (-0.7,0.7) 	(0,1) 	(0.7,0.7)
// Output
// (-1, 1)	(0, 1)	(1, 1)
// (-1, 0)	(0, 0)	(1, 0)
// (-1, -1)	(0, -1)	(1, -1)
func JoystickFixFlutterData(baseStrategy JoystickStrategy) JoystickStrategy {
	return func(x, y float32) (left float32, right float32) {
		return baseStrategy(x, -y)
	}
}

func JoystickRotationStrategy(x, y float32) (left, right float32) {

	const sqh = float32(math.Sqrt2 / 2)

	left = x*sqh + y*sqh
	right = -x*sqh + y*sqh

	return
}

func JoystickPolarStrategy(turn_damping float64) JoystickStrategy {

	return func(x, y float32) (left float32, right float32) {
		theta := math.Atan2(float64(y), float64(x))
		r := math.Sqrt(float64(x*x + y*y))

		var max_r float64 = 0

		if math.Abs(float64(x)) > math.Abs(float64(y)) {
			max_r = math.Abs(r / float64(x))
		} else {
			max_r = math.Abs(r / float64(y))
		}

		magnitude := r / max_r

		left = float32(magnitude * (math.Sin(theta) + math.Cos(theta)/turn_damping))
		right = float32(magnitude * (math.Sin(theta) - math.Cos(theta)/turn_damping))
		return
	}

}
