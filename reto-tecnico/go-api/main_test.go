package main

import (
	"reflect"
	"testing"
)

func TestRotateMatrix(t *testing.T) {
	input := [][]float64{
		{1, 2},
		{3, 4},
		{5, 6},
	}

	expected := [][]float64{
		{5, 3, 1},
		{6, 4, 2},
	}

	result := rotateMatrix(input)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Resultado incorrecto. Esperado %v, obtenido %v", expected, result)
	}
}