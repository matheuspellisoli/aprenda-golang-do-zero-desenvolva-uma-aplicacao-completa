package address

import "testing"

type testScene struct {
	received string
	expected string
}

func TestTypeAddress2(t *testing.T) {

	testScene := []testScene{
		{"Rua dos bobos", "rua"},
		{"Avenida Paulista", "avenida"},
		{"estrada 01", "estrada"},
		{"rodovia 03", "rodovia"},
		{"praca nova", "inválido"},
		{"", "inválido"},
	}

	for _, scene := range testScene {
		if TypeAddress(scene.received) != scene.expected {
			t.Errorf("tipo não é o esperado received: %s, expected: %s ", scene.received, scene.expected)
		}
	}

}
