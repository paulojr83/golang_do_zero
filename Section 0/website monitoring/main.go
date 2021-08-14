package main

import "fmt"
import "os"
import "net/http"

func main() {

	exibirIntruducao()
	exibeMenu()
	comando := lerComando()

	switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibir Logs")

		case 0:
			fmt.Println("Sair do Programa")
			os.Exit(0)
		default:
			fmt.Printf("Não reconheço esse comando %d", comando)
			os.Exit(-1)
	}

}

func iniciarMonitoramento(){
	fmt.Println("Iniciar Monitoramento...")
	site := "https://www.alura.com.br"
	resp, _ := http.Get(site)
	if resp.StatusCode == 200 {
		fmt.Println(resp.Status)
	}else {
		fmt.Printf("O site %s está offline", site)
	}

}
func exibirIntruducao() {
	var nome string
	fmt.Println("Olá! qual é o seu nome?")
	fmt.Scan(&nome)
	versao := 1.1
	fmt.Println("Olá sr.", nome, "seja bem vindo!")
	fmt.Println("Este programa está na versão", versao)
}

func lerComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)

	fmt.Println("O comando escolhido foi", comandoLido)
	return comandoLido
}

func exibeMenu(){
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
}