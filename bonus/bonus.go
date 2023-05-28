package bonus

import "errors"

// Você recebe uma lista de caminhos, onde `caminhos[i] = [cidadeAi, cidadeBi]` significa que existe um caminho direto que
//vai de cidadeAi para cidadeBi. Retorne a cidade de destino, ou seja, a cidade sem nenhum caminho que saia dela.

func Destino(caminhos [][2]string) (string, error) {
	if len(caminhos) == 0 {
		return "", errors.New("Not implemented yet")
	}
	mapinha := make(map[string]int)
	if len(caminhos) == 1 {
		return caminhos[0][1], nil
	}

	for i := 0; i < len(caminhos); i++ {
		mapinha[string(caminhos[i][0])]++
		for j := 0; j < 2; j++ {
			mapinha[string(caminhos[i][j])]++
		}
	}
	for chave, valoramt := range mapinha {
		if valoramt == 1 {
			return chave, nil
		}
	}
	return "", errors.New("Ó O GAS")
}
