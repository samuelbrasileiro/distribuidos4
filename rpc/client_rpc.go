package main

import (
	"fmt"
	"math/rand"
	"net/rpc"
	"time"
	"os"
	"math"
)

type Pokemon struct {
	Name  string
	Type  string
	Level int
}

func main() {
	args := os.Args[1:2] // Ignorar o primeiro argumento, que é o nome do executável
	client, err := rpc.Dial("tcp", ":8080")
	if err != nil {
		fmt.Println("Erro ao conectar ao servidor:", err)
		return
	}

	defer func(client *rpc.Client) {
		err := client.Close()
		if err != nil {
			fmt.Println(err)
			return
		}
	}(client)


	numRequests := 10000
	RTTs := make([]time.Duration, numRequests)

	if args[0] == "0" {
		var totalRTT time.Duration

		for j := 0; j < numRequests; j++ {
			start := time.Now()
			// Gerador de números aleatórios com base na hora atual
			rand.Seed(time.Now().UnixNano())
			// Gerar um número aleatório entre 0 e 149
			randomNumber := rand.Intn(150)
			var pokemon Pokemon
			errResp := client.Call("Pokedex.GetPokemon", randomNumber, &pokemon)
			if errResp != nil {
				fmt.Println("Erro ao chamar o método:", errResp)
				continue
			}

			//fmt.Printf("Name: %s\n", pokemon.Name)
			//fmt.Printf("Type: %s\n", pokemon.Type)
			//fmt.Printf("Level: %d\n", pokemon.Level)
			//fmt.Println("------------------------")

			// Calcular o tempo de resposta
			RTT := time.Since(start)
			RTTs[j] = RTT
			totalRTT += RTT
		}
		// Calcular a média e o desvio padrão dos tempos de resposta
		// time.Duration(numRequests) é para converter para o tipo ficar igual
		avgRTT := totalRTT / time.Duration(numRequests)
		fmt.Println("Tempo médio de resposta:", avgRTT)

		// Calcula o desvio padrão dos RTTs
		var sumSquareDiffs time.Duration
		for _, RTT := range RTTs {
			diff := RTT - avgRTT
			sumSquareDiffs += (diff * diff)
		}
		variance := sumSquareDiffs / time.Duration(numRequests)
		stdDev := time.Duration(math.Sqrt(float64(variance)))
		fmt.Println("Desvio padrão dos RTTs:", stdDev)

	} else {
		for j := 0; j < numRequests; j++ {
			// Gerador de números aleatórios com base na hora atual
			rand.Seed(time.Now().UnixNano())
			// Gerar um número aleatório entre 0 e 149
			randomNumber := rand.Intn(150)
			var pokemon Pokemon
			errResp := client.Call("Pokedex.GetPokemon", randomNumber, &pokemon)
			if errResp != nil {
				fmt.Println("Erro ao chamar o método:", errResp)
				continue
			}
		}
	}
}
