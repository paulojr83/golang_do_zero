package main

import "fmt"

func diaDaSemana(numero uint) string  {
	switch numero{
	case 1:
		return "domingo"
	case 2:
		return "Segunda"
	case 3:
		return "Terça-feira"
	case 4:
		return "Quarta-feira"
	case 5:
		return "Quinta-feira"
	case 6:
		return "Sexta-feira"
	case 7:
		return "Sábado"
	default:
		return "Não é um valor valido"
	}
}

func diaDaSemana2(numero uint) string  {
	switch {
	case numero == 1:
		return "domingo"
	case numero == 2:
		return "Segunda"
	case numero == 3:
		return "Terça-feira"
	case numero == 4:
		return "Quarta-feira"
	case numero == 5:
		return "Quinta-feira"
	case numero == 6:
		return "Sexta-feira"
	case numero == 7:
		return "Sábado"
	default:
		return "Não é um valor valido"
	}
}

func diaDaSemana3(numero uint) string  {
	var diaSemana string
	switch {
	case numero == 1:
		diaSemana = "domingo"
		fallthrough // Com isso ela vai executar a próximo case
	case numero == 2:
		diaSemana = "Segunda"
	case numero == 3:
		diaSemana = "Terça-feira"
	case numero == 4:
		diaSemana = "Quarta-feira"
	case numero == 5:
		diaSemana = "Quinta-feira"
	case numero == 6:
		diaSemana = "Sexta-feira"
	case numero == 7:
		diaSemana = "Sábado"
	default:
		diaSemana = "Não é um valor valido"
	}

	return diaSemana
}

func main() {
	diaSemana1 := diaDaSemana(1)
	fmt.Println(diaSemana1)

	diaSemana2 := diaDaSemana2(1)
	fmt.Println(diaSemana2)

	diaSemana3 := diaDaSemana3(1)
	fmt.Println(diaSemana3)
}
