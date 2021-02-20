package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
 * Faça um algoritmo em Go que preencha dois vetores A e B (TAMANHO: 5) com valores inteiros.
 * Depois de preencher os vetores A e B, exiba um vetor C, sendo que C = A + B. Usar o RAND.
 */

const tamanho = 10

func main() {

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	var vetorA [tamanho]int
	var vetorB [tamanho]int
	var vetorC [tamanho]int

	for i := 0; i < 5; i++ {

		vetorA[i] = r1.Intn(100)
		vetorB[i] = r1.Intn(100)
		vetorC[i] = vetorA[i] + vetorB[i]

	}

	for i := 0; i < 5; i++ {
		fmt.Printf("Valor de vetorA[%d]: %d\n", i, vetorA[i])
		fmt.Printf("Valor de vetorB[%d]: %d\n", i, vetorB[i])
		fmt.Printf("Valor do vetorC[%d]: %d\n", i, vetorC[i])
		fmt.Printf("---------------------------------\n")

	}

}
