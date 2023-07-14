package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

func main() {
	// Definir o número de clientes
	numClients := 80

	// Esperar por todos os clientes concluírem
	var wg sync.WaitGroup
	wg.Add(numClients)

	for i := 0; i < numClients; i++ {
		// Para passar o valor de 'i' dentro de cada rotina, precisei passar ele como parâmetro
		go func(i int) {
			// Conectar ao servidor TCP na porta 8080
			conn, err := net.Dial("tcp", "localhost:8080")
			if err != nil {
				fmt.Println("Erro ao conectar ao servidor:", err)
				wg.Done()
				return
			}
			defer conn.Close()

			// Enviar solicitações para o servidor e medir o tempo de resposta
			numRequests := 10000
			if i == 0 {
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
						wg.Done()
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
						wg.Done()
						return
					}

				}

			}

			wg.Done()
		}(i)
	}

	// Aguardar todas as goroutines terminarem
	wg.Wait()
}
