package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	serverAddr, err := net.ResolveUDPAddr("udp", "localhost:8080")
	if err != nil {
		fmt.Println("Erro ao resolver endereço:", err)
		return
	}

	conn, err := net.DialUDP("udp", nil, serverAddr)
	if err != nil {
		fmt.Println("Erro ao conectar ao servidor:", err)
		return
	}

	defer conn.Close()

	// Medir o tempo de resposta (RTT) para N solicitações
	numRequests := 2
	totalDuration := time.Duration(0)

	for i := 0; i < numRequests; i++ {
		start := time.Now()

		// Envie a solicitação para o servidor
		_, err := conn.Write([]byte("request"))
		if err != nil {
			fmt.Println("Erro ao enviar solicitação:", err)
			continue
		}

		// Receba a resposta do servidor
		buffer := make([]byte, 1024)
		_, _, err = conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("Erro ao receber resposta:", err)
			continue
		}

		duration := time.Since(start)
		totalDuration += duration
	}

	averageRTT := totalDuration / time.Duration(numRequests)
	fmt.Println("Tempo médio de resposta (UDP):", averageRTT)
}
