// Comentário de uma linha

/*
	Comentário em múltiplas linhas

	Código GO é por padrão UTF-8
	Package é o comando inicial para determinar em qual pacote está o código
*/
package main

// Import serve para importar os pacotes nativos e criador por você mesmo
import "fmt"

// main() é um função reservada, ela é o início da aplicação
func main() {
	// Usando a função Println que está dentro do pacote fmt
	fmt.Println("Hello world")
}
