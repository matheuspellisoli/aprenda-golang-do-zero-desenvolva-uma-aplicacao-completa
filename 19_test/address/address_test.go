package address

import "testing"

func TestTypeAddress(t *testing.T) {

	addressValue := "Rua dos bobos"

	typeAddress := "rua"

	if TypeAddress(addressValue) != typeAddress {
		t.Error("tipo não é o esperado")
	}
}
