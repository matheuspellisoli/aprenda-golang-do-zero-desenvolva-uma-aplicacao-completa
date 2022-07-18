package address

import "strings"

func TypeAddress(address string) string {
	addressValid := []string{"rua", "avenida", "estrada", "rodovia"}

	first := strings.Split(strings.ToLower(address), " ")[0]
	valid := false

	for _, typeAddress := range addressValid {
		if typeAddress == first {
			valid = true
		}
	}

	if valid {
		return first
	}

	return "inválido"
}
