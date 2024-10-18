//Implimente um programa que armazene 20 números inteiros em um array e, em sequida exiba os valores em ordem inversa.
package main

import "fmt"

func main() {
    var numeros [20]int

    // Solicitando que o usuário insira 20 números inteiros
    fmt.Println("Insira 20 números inteiros:")

    for i := 0; i < 20; i++ {
        fmt.Printf("Número %d: ", i+1)
        fmt.Scan(&numeros[i])
    }

    // Exibindo os números na ordem inversa
    fmt.Println("\nNúmeros na ordem inversa:")
    for i := 19; i >= 0; i-- {
        fmt.Printf("%d ", numeros[i])
    }
    fmt.Println() // Para pular uma linha ao final
}
