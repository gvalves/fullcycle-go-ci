package main

import "testing"

func TestSum(t *testing.T) {
	sum := Sum(15, 15)
	expected := 30

	if sum != expected {
		t.Errorf("Falha na soma. Esperado: %d. Recebido: %d\n", expected, sum)
	}
}
