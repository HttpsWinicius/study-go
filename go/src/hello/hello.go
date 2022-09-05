package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	exibeIntroducao()

	for {
		exibeMenu()

		comando := leComando()

		if comando == 1 {
			iniciarMonitoramento()
		} else if comando == 2 {
			fmt.Println("Exibindo Logs...")
		} else if comando == 0 {
			fmt.Println("Saindo do programa...")
			os.Exit(0)
		} else {
			fmt.Println("Comando não reconhecido")
			leComando()
			os.Exit(-1)
		}
	}

}

func exibeIntroducao() {

	var nome string
	fmt.Println("Por favor, digite seu nome: ")
	fmt.Scan(&nome)
	fmt.Println("Olá, ", nome)
}

func exibeMenu() {

	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do programa")

}

func leComando() int {

	var comandoLido int
	fmt.Println("Por favor, escolha um comando do menu: ")
	//o " & " siginfica que estamos esperando o valor da variavel
	fmt.Scan(&comandoLido)

	fmt.Println("O comando escolhido foi", comandoLido)

	return comandoLido
}

func iniciarMonitoramento() {
	fmt.Println("Por favor, digite o site que deseja monitorar:")
	var site string
	fmt.Scan(&site)

	fmt.Println("Iniciando monitoramento do site:", site)

	resposta, erro := http.Get(site)

	if erro != nil {
		fmt.Println("Site", site, "está fora do ar")
	} else if resposta.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site:", site, "esta com problemas. Status Code: ", resposta.StatusCode)
	}

}
