package main

type TemperatureSensor struct {
	Estimator
	temperature int
}

func MakeTemperatureSensor() *TemperatureSensor {
	return &TemperatureSensor{
		temperature: 0,
	}
}

func (sensor *TemperatureSensor) Get() int {
	return sensor.temperature
}
