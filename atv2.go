package main

import "fmt"

// Função para encontrar o maior e o menor valor em um array
func encontrarMaiorMenor(numeros [10]int) (int, int) {
    maior := numeros[0]
    menor := numeros[0]

    for _, valor := range numeros {
        if valor > maior {
            maior = valor
        }
        if valor < menor {
            menor = valor
        }
    }
    return maior, menor
}

func main() {
    var numeros [10]int

    // Solicitando que o usuário insira 10 números inteiros
    fmt.Println("Insira 10 números inteiros:")

    for i := 0; i < 10; i++ {
        fmt.Printf("Número %d: ", i+1)
        fmt.Scan(&numeros[i])
    }

    // Encontrar o maior e o menor valor
    maior, menor := encontrarMaiorMenor(numeros)

    // Exibir o maior e o menor valor
    fmt.Printf("O maior valor é: %d\n", maior)
    fmt.Printf("O menor valor é: %d\n", menor)
}
