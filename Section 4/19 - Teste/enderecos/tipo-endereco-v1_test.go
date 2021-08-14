package enderecos

import "testing"

func TestTipoDeEndereco(t *testing.T) {
	t.Parallel()
	enderecoParaTeste := "Rua das bananeiras"
	tipoEnderecoEsperado := "Rua"

	tipoEnderecoRecibido := TipoDeEndereco(enderecoParaTeste)

	if tipoEnderecoRecibido != tipoEnderecoEsperado {
		t.Errorf("O tipo recebido é dirente do esperado!\nEsperava %s e recebou %s",
			tipoEnderecoEsperado,
			tipoEnderecoRecibido)
	}

}