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
)

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
			// Enviar solicitação para o servidor
			req := []byte("Solicitação do cliente")
			_, err := conn.Write(req)
			if err != nil {
				fmt.Println("Erro ao tentar enviar request:", err)
				os.Exit(0)
			}

			// Receber resposta do servidor
			rep := make([]byte, 1024)
			_, _, errResp := conn.ReadFromUDP(rep)
			if errResp != nil {
				fmt.Println("Erro ao receber resposta do servidor:", errResp)
				os.Exit(0)
			}

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
			// Enviar solicitação para o servidor
			req := []byte("Solicitação do cliente")
			_, err := conn.Write(req)
			if err != nil {
				fmt.Println("Erro ao tentar enviar request:", err)
				os.Exit(0)
			}

			// Receber resposta do servidor
			rep := make([]byte, 1024)
			_, _, errResp := conn.ReadFromUDP(rep)
			if errResp != nil {
				fmt.Println("Erro ao receber resposta do servidor:", errResp)
				os.Exit(0)
			}

		}

	}
}