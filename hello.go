package main

// começando testes em go
// bibliotecas padroes https://pkg.go.dev/std

import (
	"fmt"
	"net/http"
	"os"
	"reflect"
)

func main() {

	exibeIntroducao()
	exibeNomes()
	// nao existe while true em go. for se condicao é equivalente

	for {
		exibeMenu()
		comando := leComando()

		// nao precisa break no switch
		switch comando {
		case 1:
			iniciarMonitoramento()

		case 2:
			fmt.Println("Exibindo Logs...")
		case 0:
			fmt.Println("Até Logo! ;)")
			os.Exit(0)
		// utilizamos -1 no os.exit p mostrar que saimos de forma inesperada.
		default:
			fmt.Println("Não conheço este comando")
			os.Exit(-1)
		}

	}

	// if comando == 1 {
	// 	fmt.Println("Monitorando...")
	// } else if comando == 2 {
	// 	fmt.Println("Exibindo Logs...")
	// } else if comando == 0 {
	// 	fmt.Println("Saindo..")
	// } else {
	// 	fmt.Println("Não conheço este comando")
	// }

}

func exibeIntroducao() {
	nome := "Paulo"
	versao := 1.1
	fmt.Println("Olá, sr.", nome)
	fmt.Println("Este programa esta na versao", versao)
}

func exibeMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O comando escolhido foi", comandoLido)

	return comandoLido
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	// array tem tamanho fixo, nao é dinamico. trabalhamos com slice
	var sites [4]string
	sites[0] = "https://www.alura.com.br"
	sites[1] = "https://random-status-code.herokuapp.com/"
	sites[2] = "https://globo.com"

	fmt.Println(sites)

	site := "https://www.alura.com.br"
	// _ serve para ignorar a variavel q nao quer neste caso o err
	resp, _ := http.Get(site)
	fmt.Println(resp)

	if resp.StatusCode == 200 {
		fmt.Println("O Site:", site, "foi carregado corretamente")
	} else {
		fmt.Println("O Site:", site, "esta com problema")
	}
	// da p trazer conteudo?
	// https://higordiego.com.br/posts/golang-jira/
}

func exibeNomes() {
	nomes := []string{"Douglas", "Daniel", "Julio"}
	fmt.Println(nomes)
	fmt.Println(reflect.TypeOf(nomes))
	fmt.Println("O meu slice tem a capacidade para", cap(nomes), "itens. Além disso tem", len(nomes))

	nomes = append()
}
