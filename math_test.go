package main

import "testing"

func TestSoma(t *testing.T){
	total := Soma(15, 15)
	if total != 30{
		t.Errorf("Resultado errado, deu %d, era pra ter dado %d", total, 30)
	}
}