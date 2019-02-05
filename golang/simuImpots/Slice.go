package main

func ResizeSlice(inputSlice []float32, newSize int) (outputSlice []float32) {
	deltaSize := newSize - len(inputSlice)
	outputSlice = make([]float32, len(inputSlice)+deltaSize, cap(inputSlice)+deltaSize)
	copy(outputSlice, inputSlice)
	return outputSlice
}
