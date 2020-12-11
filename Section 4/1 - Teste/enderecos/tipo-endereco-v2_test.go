package enderecos

import "testing"


type cenariosDeTeste struct {
	enderecoParaTeste 		string
	tipoEnderecoEsperado 	string
}

func TestTipoDeEnderecoV2(t *testing.T) {
	t.Parallel()

	cenariosDeTeste := []cenariosDeTeste{
		{"Rua das bananeiras","Rua", },
		{"Avenida das bananeiras","Avenida", },
		{"Estrada das bananeiras","Estrada", },
		{"Rodovia das bananeiras","Rodovia", },
		{"Das bananeiras","Não é um tipo valido", },
		{"","Não é um tipo valido", },
	}

	for _, cenario := range cenariosDeTeste {
		tipoEnderecoRecibido := TipoDeEndereco(cenario.enderecoParaTeste)

		if tipoEnderecoRecibido != cenario.tipoEnderecoEsperado {
			t.Errorf("O tipo recebido é dirente do esperado!\nEsperava %s e recebou %s",
				cenario.tipoEnderecoEsperado,
				tipoEnderecoRecibido)
		}
	}


}