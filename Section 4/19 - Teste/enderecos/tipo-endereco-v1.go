package enderecos

import "strings"

//TipoDeEndereco verifica se o endereço tem um tipo valido
func TipoDeEndereco(endereco string) string  {
	tiposValidos := []string{"rua", "avenida", "estrada", "rodovia"}

	enderecoEmLetraMinuscula := strings.ToLower(endereco)
	primeiraOalavraDoEndereco := strings.Split(enderecoEmLetraMinuscula, " ")[0]

	enderecoTemUmTipoValido := false
	for _, tipo := range tiposValidos {
		if( tipo == primeiraOalavraDoEndereco ) {
			enderecoTemUmTipoValido = true
		}
	}
	if enderecoTemUmTipoValido {
		return strings.Title(primeiraOalavraDoEndereco)
	}

	return "Não é um tipo valido"
}

