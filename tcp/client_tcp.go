package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	serverAddr := "localhost:8080"

	conn, err := net.Dial("tcp", serverAddr)
	if err != nil {
		fmt.Println("Erro ao conectar ao servidor:", err)
		return
	}

	defer conn.Close()

	// Medir o tempo de resposta (RTT) para N solicitações
	numRequests := 10000
	totalDuration := time.Duration(0)

	for i := 0; i < numRequests; i++ {
		start := time.Now()

		// Envie a solicitação para o servidor
		request := "solicitação\n"
		// print("Cliente mandou: " + request)
		_, err := conn.Write([]byte(request))
		if err != nil {
			fmt.Println("Erro ao enviar solicitação: ", err)
			continue
		}

		// Receba a resposta do servidor
		buffer := make([]byte, 1024)
		_, err = conn.Read(buffer)
		if err != nil {
			fmt.Println("Erro ao receber resposta: ", err)
			continue
		}

		duration := time.Since(start)
		totalDuration += duration
	}

	averageRTT := totalDuration / time.Duration(numRequests)
	fmt.Println("Tempo médio de resposta (TCP):", averageRTT)
}
