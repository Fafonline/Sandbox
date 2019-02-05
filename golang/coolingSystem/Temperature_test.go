package main

import "testing"

func TestTemperatureSensor_Get(t *testing.T) {
	type fields struct {
		Estimator   Estimator
		temperature int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		name: "When Temperature sensor is created",
		want: 0,
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sensor := &TemperatureSensor{
				Estimator:   tt.fields.Estimator,
				temperature: tt.fields.temperature,
			}
			if got := sensor.Get(); got != tt.want {
				t.Errorf("TemperatureSensor.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
