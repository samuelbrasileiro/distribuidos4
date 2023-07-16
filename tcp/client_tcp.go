// Para criar o executável do client usa o comando:
//  Linux -> go build -o client_tcp client_tcp.go
// Windows -> go build -o client_tcp.exe client_tcp.go

package main

import (
	"fmt"
	"net"
	"time"
	"os"
	"math"
)

func main() {
	args := os.Args[1:2] // Ignorar o primeiro argumento, que é o nome do executável
	
	// Retorna o endereço do endpoint TCP
	r, err := net.ResolveTCPAddr("tcp", ":8080")
	if err != nil {
		fmt.Println("Erro ao buscar endereço do endpoint TCP")
	}
	
	// Conectar ao servidor TCP na porta 8080
	conn, err := net.DialTCP("tcp", nil, r)
	if err != nil {
		fmt.Println("Erro ao conectar ao servidor")
		return
	}

	// Fecha Conexão
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			fmt.Println("Erro ao fechar conexão")
			os.Exit(0)
		}
	}(conn)

	// Enviar solicitações para o servidor e medir o tempo de resposta
	numRequests := 10000
	RTTs := make([]time.Duration, numRequests)
	if args[0] == "0" {
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
				fmt.Println("Erro ao receber resposta do servidor")
				return
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
			request := "Solicitação do cliente"
			conn.Write([]byte(request))

			// Receber resposta do servidor
			buffer := make([]byte, 1024)
			_, err := conn.Read(buffer)
			if err != nil {
				fmt.Println("Erro ao receber resposta do servidor")
				return
			}

		}

	}
}
