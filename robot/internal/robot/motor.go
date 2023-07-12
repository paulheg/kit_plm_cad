package robot

type MotorSize string

const (
	Medium MotorSize = "lego-ev3-m-motor"
	Large  MotorSize = "lego-ev3-l-motor"
)

type MotorCommand string

const (
	Stop        MotorCommand = "stop"
	RunForever  MotorCommand = "run-forever"
	RunToAbsPos MotorCommand = "run-to-abs-pos"
	RunToRelPos MotorCommand = "run-to-rel-pos"
)
