package main

import "fmt"

// Função que recebe um array de inteiros e retorna a média
func calcularMedia(numeros [20]int) float64 {
    soma := 0
    for _, valor := range numeros {
        soma += valor
    }
    return float64(soma) / float64(len(numeros))
}

func main() {
    // Criando um array com 20 números inteiros
    numeros := [20]int{10, 15, 20, 25, 30, 35, 40, 45, 50, 55, 60, 65, 70, 75, 80, 85, 90, 95, 100, 105}

    // Chamando a função para calcular a média
    media := calcularMedia(numeros)

    // Exibindo a média
    fmt.Printf("A média dos valores é: %.2f\n", media)
}
