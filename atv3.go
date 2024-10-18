//Escreva um programa que crie um array com 20 números inteiros e conte quantas são paris e quantes são impares.
package main

import "fmt"

// Função para contar números pares e ímpares
func contarParesImpares(numeros [20]int) (int, int) {
    pares := 0
    impares := 0

    for _, valor := range numeros {
        if valor%2 == 0 {
            pares++
        } else {
            impares++
        }
    }

    return pares, impares
}

func main() {
    var numeros [20]int

    // Solicitando que o usuário insira 20 números inteiros
    fmt.Println("Insira 20 números inteiros:")

    for i := 0; i < 20; i++ {
        fmt.Printf("Número %d: ", i+1)
        fmt.Scan(&numeros[i])
    }

    // Contar números pares e ímpares
    pares, impares := contarParesImpares(numeros)

    // Exibir a quantidade de pares e ímpares
    fmt.Printf("Quantidade de números pares: %d\n", pares)
    fmt.Printf("Quantidade de números ímpares: %d\n", impares)
}

