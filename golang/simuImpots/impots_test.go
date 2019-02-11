package main

import (
	"reflect"
	"testing"
)

func TestCalculeTranche(t *testing.T) {
	TranchesMock := TrancheBuilderMock{}.Build()
	dependencies := Dependencies{
		barême: MakeBarême(TrancheBuilderMock{}),
	}
	calculateurTranche := NewCalculateurTranche(&dependencies)

	type args struct {
		revenu float32
	}
	tests := []struct {
		name                string
		args                args
		wantTranches        []float32
		wantRevenuImposable float32
	}{
		// TODO: Add test cases.
		{
			name:         "Non Imposable",
			args:         args{revenu: TranchesMock[0].threshold / 2},
			wantTranches: []float32{TranchesMock[0].threshold / 2},
		},
		{
			name:         "Dans 1er tranche",
			args:         args{revenu: TranchesMock[0].threshold + (TranchesMock[1].threshold-TranchesMock[0].threshold)/2},
			wantTranches: []float32{TranchesMock[0].threshold, (TranchesMock[1].threshold - TranchesMock[0].threshold) / 2},
		},
		{
			name:         "Dans 2ieme tranche",
			args:         args{revenu: TranchesMock[1].threshold + (TranchesMock[2].threshold-TranchesMock[1].threshold)/2},
			wantTranches: []float32{TranchesMock[0].threshold, TranchesMock[1].threshold - TranchesMock[0].threshold, (TranchesMock[2].threshold - TranchesMock[1].threshold) / 2},
		},
		{
			name:         "Au delà de la dernière tranche",
			args:         args{revenu: TranchesMock[2].threshold + 1000},
			wantTranches: []float32{TranchesMock[0].threshold, TranchesMock[1].threshold - TranchesMock[0].threshold, TranchesMock[2].threshold - TranchesMock[1].threshold, 1000},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotTranches := calculateurTranche.Execute(tt.args.revenu); !reflect.DeepEqual(gotTranches, tt.wantTranches) {
				t.Errorf("Revenu:%f - CalculeTranche() = %v, want %v", tt.args.revenu, gotTranches, tt.wantTranches)
			}
		})
	}
}

func TestResizeSlice(t *testing.T) {
	type args struct {
		inputSlice []float32
		newSize    int
	}
	tests := []struct {
		name            string
		args            args
		wantOutputSlice []float32
	}{
		// TODO: Add test cases.
		{
			name: "Add 1 element to empty slice",
			args: args{
				inputSlice: nil,
				newSize:    1,
			},
			wantOutputSlice: []float32{0},
		},
		{
			name: "Add 2 elements to empty slice",
			args: args{
				inputSlice: nil,
				newSize:    2,
			},
			wantOutputSlice: []float32{0, 0},
		},
		{
			name: "Add 10 elements to empty slice",
			args: args{
				inputSlice: nil,
				newSize:    10,
			},
			wantOutputSlice: []float32{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			name: "Remove 1 element in non-empty",
			args: args{
				inputSlice: make([]float32, 4),
				newSize:    3,
			},
			wantOutputSlice: []float32{0, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOutputSlice := ResizeSlice(tt.args.inputSlice, tt.args.newSize); !reflect.DeepEqual(gotOutputSlice, tt.wantOutputSlice) {
				t.Errorf("ResizeSlice() = %v, want %v", gotOutputSlice, tt.wantOutputSlice)
			}
		})
	}
}

func TestBarême_Size(t *testing.T) {
	type fields struct {
		outBoundTranche []OutBoundTranche
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
		{
			name: "Taille Barême de 3 tranches imposables",
			fields: fields{
				outBoundTranche: MakeBarême(TrancheBuilderMock{}).outBoundTranche,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Barême{
				outBoundTranche: tt.fields.outBoundTranche,
			}
			if got := b.Size(); got != tt.want {
				t.Errorf("Barême.Size() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBarême_Get(t *testing.T) {
	TranchesMock := TrancheBuilderMock{}.Build()
	type fields struct {
		tranches []float32
	}
	type args struct {
		index int
	}
	tests := []struct {
		name string
		args args
		want *Tranche
	}{
		// TODO: Add test cases.
		{
			name: "Première tranche",
			args: args{
				index: 0,
			},
			want: &Tranche{0, TranchesMock[0].threshold, false, TranchesMock[0].coefficient},
		},
		// TODO: Add test cases.
		{
			name: "Deuxieme tranche",
			args: args{
				index: 1,
			},
			want: &Tranche{TranchesMock[0].threshold, TranchesMock[1].threshold, false, TranchesMock[1].coefficient},
		},
		// TODO: Add test cases.
		{
			name: "Derniere tranche",
			args: args{
				index: 2,
			},
			want: &Tranche{TranchesMock[1].threshold, TranchesMock[2].threshold, true, TranchesMock[2].coefficient},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := MakeBarême(TrancheBuilderMock{})
			if got := b.Get(tt.args.index); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Barême.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
