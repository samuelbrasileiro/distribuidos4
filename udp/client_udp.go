// Para criar o executável do client usa o comando:
// Linux -> go build -o client_udp client_udp.go
// Windows -> go build -o client_udp.exe client_udp.go
package main

import (
	"fmt"
	"net"
	"os"
	"time"
	"math"
	"math/rand"
	"encoding/json"
)


type Pokemon struct {
	Name  string
	Type  string
	Level int
}

func main() {
	args := os.Args[1:2] // Ignorar o primeiro argumento, que é o nome do executável
	// Retorna o endereço do endpoint UDP 8080
	addr, err := net.ResolveUDPAddr("udp", "localhost:8080")
	if err != nil {
		fmt.Println("Erro ao retornar endereço do endpoint UDP")
		return
	}

	// Conectar ao servidor UDP na porta 8080
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		fmt.Println("Erro ao conectar ao servidor")
		return
	}
	defer conn.Close()

	// Enviar solicitações para o servidor e medir o tempo de resposta
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

			// Enviar solicitação para o servidor para retornar um pokemon
			// Convertendo o número inteiro em um slice de bytes
			req := make([]byte, 4)
			req[0] = byte(randomNumber >> 24)
			req[1] = byte(randomNumber >> 16)
			req[2] = byte(randomNumber >> 8)
			req[3] = byte(randomNumber)
			_, err := conn.Write(req)
			if err != nil {
				fmt.Println("Erro ao tentar enviar request:", err)
				os.Exit(0)
			}

			// Receber resposta do servidor que é um objeto com informações de um pokemon
			rep := make([]byte, 1024)
			n, _, errResp := conn.ReadFromUDP(rep)
			if errResp != nil {
				fmt.Println("Erro ao receber resposta do servidor:", errResp)
				os.Exit(0)
			}

			// Desserializando a resposta para criar um novo objeto Pokemon
			var pokemon Pokemon
			errDess := json.Unmarshal(rep[:n], &pokemon)
			if errDess != nil {
				fmt.Println("Erro ao desserializar a resposta do servidor:", errDess)
				os.Exit(1)
			}

			// fmt.Printf("Name: %s\n", pokemon.Name)
			// fmt.Printf("Type: %s\n", pokemon.Type)
			// fmt.Printf("Level: %d\n", pokemon.Level)
			// fmt.Println("------------------------")

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

			// Enviar solicitação para o servidor para retornar um pokemo
			// Convertendo o número inteiro em um slice de bytes
			req := make([]byte, 4)
			req[0] = byte(randomNumber >> 24)
			req[1] = byte(randomNumber >> 16)
			req[2] = byte(randomNumber >> 8)
			req[3] = byte(randomNumber)
			_, err := conn.Write(req)
			if err != nil {
				fmt.Println("Erro ao tentar enviar request:", err)
				os.Exit(0)
			}

			// Receber resposta do servidor que é um objeto com informações de um pokemon
			rep := make([]byte, 1024)
			n, _, errResp := conn.ReadFromUDP(rep)
			if errResp != nil {
				fmt.Println("Erro ao receber resposta do servidor:", errResp)
				os.Exit(0)
			}

			// Desserializando a resposta para criar um novo objeto Pokemon
			var pokemon Pokemon
			errDess := json.Unmarshal(rep[:n], &pokemon)
			if errDess != nil {
				fmt.Println("Erro ao desserializar a resposta do servidor:", errDess)
				os.Exit(1)
			}

			// fmt.Printf("Name: %s\n", pokemon.Name)
			// fmt.Printf("Type: %s\n", pokemon.Type)
			// fmt.Printf("Level: %d\n", pokemon.Level)
			// fmt.Println("------------------------")

		}
	}
}