package robot

import (
	"fmt"
	"math"
	"time"

	"github.com/ev3go/ev3"
	"github.com/ev3go/ev3dev"
	"github.com/paulheg/kit_plm_cad/robot/internal/parser"
)

type Robot interface {
	parser.Strategy
	Dispose()
}

var _ Robot = &robot{}

type robot struct {
	leftDrive, rightDrive *ev3dev.TachoMotor

	heightDrive, pitchDrive *ev3dev.TachoMotor

	speaker *ev3dev.Speaker

	maxLarge float32

	// is set to true if the robot cam to a stop (default: true)
	stopped bool

	autoTilt bool
}

// TiltDown implements Robot.
func (r *robot) TiltDown() {
	r.pitchDrive.SetSpeedSetpoint(int(10 * r.maxLarge)).Command(string(RunForever))
	time.Sleep(100 * time.Millisecond)
	r.pitchDrive.Command(string(Stop))
}

// TiltUp implements Robot.
func (r *robot) TiltUp() {
	r.pitchDrive.SetSpeedSetpoint(-int(10 * r.maxLarge)).Command(string(RunForever))
	time.Sleep(100 * time.Millisecond)
	r.pitchDrive.Command(string(Stop))
}

// MoveDown implements Robot.
func (r *robot) MoveDown() {
	r.heightDrive.SetSpeedSetpoint(int(10 * r.maxLarge)).Command(string(RunForever))
	time.Sleep(100 * time.Millisecond)
	r.heightDrive.Command(string(Stop))
}

// MoveUp implements Robot.
func (r *robot) MoveUp() {
	r.heightDrive.SetSpeedSetpoint(-int(10 * r.maxLarge)).Command(string(RunForever))
	time.Sleep(100 * time.Millisecond)
	r.heightDrive.Command(string(Stop))
}

func New() (Robot, error) {

	left, err := getTachoMotor(D, Large)
	if err != nil {
		return nil, err
	}

	right, err := getTachoMotor(A, Large)
	if err != nil {
		return nil, err
	}

	height, err := getTachoMotor(C, Large)
	if err != nil {
		return nil, err
	}

	pitch, err := getTachoMotor(B, Medium)
	if err != nil {
		return nil, err
	}

	// SoundPath is the path to the ev3 sound events.
	const SoundPath = "/dev/input/by-path/platform-sound-event"
	speaker := ev3dev.NewSpeaker(SoundPath)
	err = speaker.Init()
	if err != nil {
		speaker.Close()
		return nil, err
	}

	err = ev3.LCD.Init(true)
	if err != nil {
		ev3.LCD.Close()
		return nil, err
	}

	return &robot{
		leftDrive:   left,
		rightDrive:  right,
		heightDrive: height,
		pitchDrive:  pitch,
		speaker:     speaker,
		stopped:     true,
		maxLarge:    float32(left.MaxSpeed() / 10),
	}, nil

}

func getTachoMotor(out Output, driver MotorSize) (*ev3dev.TachoMotor, error) {
	motor, err := ev3dev.TachoMotorFor(string(out), string(driver))
	if err != nil {
		err = fmt.Errorf("failed to find %s motor on %s: %v", driver, out, err)
		return nil, err
	}

	err = motor.SetStopAction("brake").Err()
	if err != nil {
		err = fmt.Errorf("failed to set the stop action: %v", err)
		return nil, err
	}

	return motor, nil
}

func getServoMotor(out Output, driver MotorSize) (*ev3dev.ServoMotor, error) {
	motor, err := ev3dev.ServoMotorFor(string(out), string(driver))
	if err != nil {
		err = fmt.Errorf("failed to find %s motor on %s: %v", driver, out, err)
		return nil, err
	}

	return motor, nil
}

func (r *robot) Dispose() {
	if r.speaker != nil {
		r.speaker.Close()
	}
}

// SetAutoTilt implements parser.Strategy.
func (r *robot) SetAutoTilt(tilt bool) {
	r.autoTilt = tilt

	if tilt {
		r.adjustTilt()
	}
}

func (r *robot) playStartSound() {
	r.speaker.Tone(420)
	time.Sleep(200 * time.Millisecond)
	r.speaker.Tone(0)
}

func (r *robot) playStopSound() {
	r.speaker.Tone(400)
	time.Sleep(200 * time.Millisecond)
	r.speaker.Tone(0)
}

func (r *robot) adjustTilt() {
	// r.pitchDrive.SetSpeedSetpoint(int(0.2 * r.maxLarge

}

// OnMove implements parser.Strategy.
func (r *robot) Move(x float32, y float32) {

	if x == 0 && y == 0 {
		r.stopped = true
		r.playStopSound()

		r.leftDrive.Command(string(Stop))
		r.rightDrive.Command(string(Stop))
	} else {

		if r.stopped {
			r.stopped = false

			// play sound
			r.playStartSound()
		}

		r.drive(x, y)
	}
}

func (r *robot) drive(x, y float32) {

	left, right := JoyconToLeftRight(x, y)

	r.leftDrive.SetSpeedSetpoint(int(left * r.maxLarge)).Command(string(RunForever))
	r.rightDrive.SetSpeedSetpoint(int(right * r.maxLarge)).Command(string(RunForever))

}

func JoyconToLeftRight(x, y float32) (left, right float32) {

	const sqh = float32(math.Sqrt2 / 2)

	left = x*sqh + y*sqh
	right = -x*sqh + y*sqh

	return
}
