package main

import (
	"fmt"
	"net"
	"time"
	"os"
)

func main() {
	args := os.Args[1:2] // Ignorar o primeiro argumento, que é o nome do executável
	// Conectar ao servidor TCP na porta 8082
	conn, err := net.Dial("udp", "localhost:8082")
	if err != nil {
		fmt.Println("Erro ao conectar ao servidor:", err)
		return
	}
	defer conn.Close()

	// Enviar solicitações para o servidor e medir o tempo de resposta
	numRequests := 10000
	if args[0] == "0" {
		fmt.Println("Entrei meu peixe")
		var totalRTT time.Duration

		for j := 0; j < numRequests; j++ {
			start := time.Now()

			// Enviar solicitação para o servidor
			request := "Solicitação do cliente"
			conn.Write([]byte(request))

			// Receber resposta do servidor
			buffer := make([]byte, 1024)
			_, err := conn.Read(buffer)
			if err != nil {
				fmt.Println("Erro ao receber resposta do servidor:", err)
				return
			}

			// Calcular o tempo de resposta
			rtt := time.Since(start)
			totalRTT += rtt
		}
		// Calcular a média e o desvio padrão dos tempos de resposta
		avgRTT := totalRTT / time.Duration(numRequests)
		fmt.Println("Tempo médio de resposta:", avgRTT)

	} else {
		for j := 0; j < numRequests; j++ {

			// Enviar solicitação para o servidor
			request := "Solicitação do cliente"
			conn.Write([]byte(request))

			// Receber resposta do servidor
			buffer := make([]byte, 1024)
			_, err := conn.Read(buffer)
			if err != nil {
				fmt.Println("Erro ao receber resposta do servidor:", err)
				return
			}

		}

	}
}
