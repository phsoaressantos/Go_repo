package main

// começando testes em go

import "fmt"

func main() {
	nome := "Paulo"
	versao := 1.1
	fmt.Println("Olá, sr.", nome)
	fmt.Println("Este programa esta na versao", versao)

	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")

	var comando int
	fmt.Scan(&comando)
	fmt.Println("O endereço da minha variavel comando é", &comando)
	fmt.Println("O comando escolhido foi", comando)

<<<<<<< HEAD
	// if comando == 1 {
	// 	fmt.Println("Monitorando...")
	// } else if comando == 2 {
	// 	fmt.Println("Exibindo Logs...")
	// } else if comando == 0 {
	// 	fmt.Println("Saindo..")
	// } else {
	// 	fmt.Println("Não conheço este comando")
	// }

	// nao precisa break no switch
	switch comando {
	case 1:
		fmt.Println("Monitorando...")
	case 2:
		fmt.Println("Exibindo Logs...")
	case 3:
		fmt.Println("Saindo..")
	default:
		fmt.Println("Não conheço este comando")
	}
=======
>>>>>>> 55ba25867fe19dbae8b10f5b7d9fa429e6088391
}
