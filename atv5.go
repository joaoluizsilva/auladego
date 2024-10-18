//Grie um array de 6 números inteiros e escreva um programa que calcule o produto de todos os elementos
package main

import "fmt"

// Função para calcular o produto de todos os elementos do array
func calcularProduto(numeros [6]int) int {
    produto := 1
    for _, valor := range numeros {
        produto *= valor
    }
    return produto
}

func main() {
    var numeros [6]int

    // Solicitando que o usuário insira 6 números inteiros
    fmt.Println("Insira 6 números inteiros:")

    for i := 0; i < 6; i++ {
        fmt.Printf("Número %d: ", i+1)
        fmt.Scan(&numeros[i])
    }

    // Calculando o produto de todos os elementos
    produto := calcularProduto(numeros)

    // Exibindo o resultado
    fmt.Printf("O produto de todos os elementos é: %d\n", produto)
}
