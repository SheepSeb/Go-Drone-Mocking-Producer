package domain

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

type IMUSensor struct {
	Wx         float64   `json:"wx"`
	Wy         float64   `json:"wy"`
	Wz         float64   `json:"wz"`
	Timestamp  time.Time `json:"timestamp"`
	TypeSensor string    `json:"typeSensor"`
}

type DOFSensor struct {
	MagX       float64   `json:"magX"`
	MagY       float64   `json:"magY"`
	MagZ       float64   `json:"magZ"`
	GyroX      float64   `json:"gyroX"`
	GyroY      float64   `json:"gyroY"`
	GyroZ      float64   `json:"gyroZ"`
	Timestamp  time.Time `json:"timestamp"`
	TypeSensor string    `json:"typeSensor"`
}

type Temperature struct {
	Grades     float64   `json:"grades"`
	Timestamp  time.Time `json:"timestamp"`
	TypeSensor string    `json:"typeSensor"`
}

type Motors struct {
	Voltage    float64   `json:"voltage"`
	VoltageRef float64   `json:"voltageRef"`
	Timestamp  time.Time `json:"timestamp"`
}

type Drone struct {
	DroneId   int       `json:"droneId"`
	DroneName string    `json:"droneName"`
	Motors    []Motors  `json:"motors"`
	Sensor    []Sensors `json:"sensors"`
}

func (imu IMUSensor) GenerateData(randomizer bool, timestamp time.Time) Sensors {
	var minVal, maxVal float64
	if randomizer == true {
		minVal = -900.0
		maxVal = 900.0
	} else {
		minVal = 10.0
		maxVal = 25.0
	}
	println(time.Now().String())
	return IMUSensor{
		Wx:         minVal + rand.Float64()*(maxVal-minVal),
		Wy:         minVal + rand.Float64()*(maxVal-minVal),
		Wz:         minVal + rand.Float64()*(maxVal-minVal),
		Timestamp:  timestamp,
		TypeSensor: "IMU",
	}
}

func (dof DOFSensor) GenerateData(randomizer bool, timestamp time.Time) Sensors {
	var minVal, maxVal float64
	if randomizer == true {
		minVal = -900.0
		maxVal = 900.0
	} else {
		minVal = 10.0
		maxVal = 25.0
	}
	return DOFSensor{
		GyroX:      minVal + rand.Float64()*(maxVal-minVal),
		GyroY:      minVal + rand.Float64()*(maxVal-minVal),
		GyroZ:      minVal + rand.Float64()*(maxVal-minVal),
		MagX:       minVal + rand.Float64()*(maxVal-minVal),
		MagY:       minVal + rand.Float64()*(maxVal-minVal),
		MagZ:       minVal + rand.Float64()*(maxVal-minVal),
		Timestamp:  timestamp,
		TypeSensor: "DOF",
	}
}

func (temp Temperature) GenerateData(randomizer bool, timestamp time.Time) Sensors {
	var minVal, maxVal float64
	if randomizer == true {
		minVal = -900.0
		maxVal = 900.0
	} else {
		minVal = 10.0
		maxVal = 25.0
	}
	return Temperature{
		Grades:     minVal + rand.Float64()*(maxVal-minVal),
		Timestamp:  timestamp,
		TypeSensor: "Temperature",
	}
}

func (motor Motors) GenerateData(randomizer bool, timestamp time.Time) Sensors {
	var minVal, maxVal float64
	if randomizer == true {
		minVal = -900.0
		maxVal = 900.0
	} else {
		minVal = 10.0
		maxVal = 25.0
	}
	constantVoltage := minVal + rand.Float64()*(maxVal-minVal)
	return Motors{
		Voltage:    constantVoltage,
		VoltageRef: constantVoltage,
		Timestamp:  timestamp,
	}
}

func (imu IMUSensor) String() string {
	return fmt.Sprintf("IMUSensor{Wx: %v, Wy: %v, Wz: %v, Timestamp: %v}", imu.Wx, imu.Wy, imu.Wz, imu.Timestamp)
}

func (dof DOFSensor) String() string {
	return fmt.Sprintf("DOFSensor{GyroX: %v, GyroY: %v, GyroZ: %v, MagX: %v, MagY: %v, MagZ: %v, Timestamp: %v}", dof.GyroX, dof.GyroY, dof.GyroZ, dof.MagX, dof.MagY, dof.MagZ, dof.Timestamp)
}

func (temp Temperature) String() string {
	return fmt.Sprintf("Temperature{Grades: %v, Timestamp: %v}", temp.Grades, temp.Timestamp)
}

func (motor Motors) String() string {
	return fmt.Sprintf("Motors{Voltage: %v, VoltageRef: %v, Timestamp: %v}", motor.Voltage, motor.VoltageRef, motor.Timestamp)
}

func NewDrone(droneId int, droneName string, randomize bool) Drone {
	currentTime := time.Now()
	motors := make([]Motors, 0)
	var imu, temperature, dof Sensors

	for i := 0; i < 5; i++ {
		motorVar := Motors{}.GenerateData(randomize, currentTime).(Motors)
		motors = append(motors, motorVar)
	}

	imu = IMUSensor{}.GenerateData(randomize, currentTime).(IMUSensor)
	temperature = Temperature{}.GenerateData(randomize, currentTime).(Temperature)
	dof = DOFSensor{}.GenerateData(randomize, currentTime).(DOFSensor)

	sensors := []Sensors{imu, temperature, dof}

	drone := Drone{
		DroneId:   droneId,
		DroneName: droneName,
		Motors:    motors,
		Sensor:    sensors,
	}
	return drone
}

func NewDroneSensorsIncluded(droneId int, droneName string, randomize bool, sensors []Sensors) Drone {
	currentTime := time.Now()
	motors := make([]Motors, 0)

	for i := 0; i < 5; i++ {
		motorVar := Motors{}.GenerateData(randomize, currentTime).(Motors)
		motors = append(motors, motorVar)
	}

	drone := Drone{
		DroneId:   droneId,
		DroneName: droneName,
		Motors:    motors,
		Sensor:    sensors,
	}
	return drone
}

func (d Drone) String() string {
	motorsStr := ""
	for _, motor := range d.Motors {
		motorsStr += motor.String() + "\n"
	}

	sensorsStr := ""
	for _, sensor := range d.Sensor {
		sensorsStr += sensor.String() + "\n"
	}

	return fmt.Sprintf("Drone{DroneId: %d, DroneName: %s, Motors: [\n%s], Sensor: [\n%s]}", d.DroneId, d.DroneName, motorsStr, sensorsStr)
}

func (d Drone) ToJson() []byte {
	jsonPayload, err := json.Marshal(d)
	if err != nil {
		println(err)
	}
	return jsonPayload
}
