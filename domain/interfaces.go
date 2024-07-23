package domain

import "time"

type Sensors interface {
	GenerateData(randomizer bool, timestamp time.Time) Sensors
	String() string
}
