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
			args:         args{revenu: TranchesMock[0] / 2},
			wantTranches: []float32{TranchesMock[0] / 2},
		},
		{
			name:         "Dans 1er tranche",
			args:         args{revenu: TranchesMock[0] + (TranchesMock[1]-TranchesMock[0])/2},
			wantTranches: []float32{TranchesMock[0], (TranchesMock[1] - TranchesMock[0]) / 2},
		},
		{
			name:         "Dans 2ieme tranche",
			args:         args{revenu: TranchesMock[1] + (TranchesMock[2]-TranchesMock[1])/2},
			wantTranches: []float32{TranchesMock[0], TranchesMock[1] - TranchesMock[0], (TranchesMock[2] - TranchesMock[1]) / 2},
		},
		{
			name:         "Au delà de la dernière tranche",
			args:         args{revenu: TranchesMock[2] + 1000},
			wantTranches: []float32{TranchesMock[0], TranchesMock[1] - TranchesMock[0], TranchesMock[2] - TranchesMock[1], 1000},
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
		tranches []float32
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
				tranches: MakeBarême(TrancheBuilderMock{}).tranches,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Barême{
				tranches: tt.fields.tranches,
			}
			if got := b.Size(); got != tt.want {
				t.Errorf("Barême.Size() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMakeBarêmeError(t *testing.T) {
	type args struct {
		errorString string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "Make an error",
			args: args{
				"Test Error",
			},
			want: "Test Error",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MakeBarêmeError(tt.args.errorString); !reflect.DeepEqual(got.Error(), tt.want) {
				t.Errorf("MakeBarêmeError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBarême_Get(t *testing.T) {
	type fields struct {
		tranches []float32
	}
	type args struct {
		index int
	}
	tests := []struct {
		name string
		args args
		want float32
	}{
		// TODO: Add test cases.
		{
			name: "Première tranche",
			args: args{
				index: 0,
			},
			want: TRANCHE_01,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := MakeBarême(TrancheBuilderMock{})
			if got := b.Get(tt.args.index); got != tt.want {
				t.Errorf("Barême.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBarême_TrancheComplete(t *testing.T) {
	type fields struct {
		tranches []float32
	}
	type args struct {
		indiceTranche int
	}
	tests := []struct {
		name        string
		args        args
		wantTranche float32
		wantErr     bool
	}{
		// TODO: Add test cases.
		{
			name: "Première tranche",
			args: args{
				indiceTranche: 0,
			},
			wantTranche: TrancheBuilderMock{}.Build()[1] - TrancheBuilderMock{}.Build()[0],
			wantErr:     false,
		},
		{
			name: "Dernière tranche",
			args: args{
				indiceTranche: MakeBarême(TrancheBuilderMock{}).Size() - 1,
			},
			wantTranche: 0,
			wantErr:     true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := MakeBarême(TrancheBuilderMock{})
			gotTranche, err := b.TrancheComplete(tt.args.indiceTranche)
			if (err != nil) != tt.wantErr {
				t.Errorf("Barême.TrancheComplete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotTranche != tt.wantTranche {
				t.Errorf("Barême.TrancheComplete() = %v, want %v", gotTranche, tt.wantTranche)
			}
		})
	}
}
