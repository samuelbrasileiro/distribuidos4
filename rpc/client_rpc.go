package main

import (
	"fmt"
	"math/rand"
	"net/rpc"
	"time"
)

type Pokemon struct {
	Name  string
	Type  string
	Level int
}

func main() {
	client, err := rpc.Dial("tcp", ":8080")
	if err != nil {
		fmt.Println("Erro ao conectar ao servidor:", err)
		return
	}

	numRequests := 10000
	RTTs := make([]time.Duration, numRequests)
	var totalRTT time.Duration

	for j := 0; j < numRequests; j++ {
		start := time.Now()

		randomNumber := rand.Intn(150)
		var pokemon Pokemon
		err := client.Call("Pokedex.GetPokemon", randomNumber, &pokemon)
		if err != nil {
			fmt.Println("Erro ao chamar o método:", err)
			continue
		}

		RTT := time.Since(start)
		RTTs[j] = RTT
		totalRTT += RTT
	}

	avgRTT := totalRTT / time.Duration(numRequests)
	fmt.Println("Tempo médio de resposta:", avgRTT)

}
