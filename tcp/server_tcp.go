package main

import (
	"fmt"
	"io"
	"net"
)

func handleConnection(conn net.Conn) {
	// Lógica para lidar com a conexão do cliente
	// Implemente aqui o processamento necessário para atender a solicitação do cliente

	// Lógica para lidar com a conexão do cliente
	buffer := make([]byte, 1024)
	for {
		// Receba a solicitação do cliente
		n, err := conn.Read(buffer)
		if err != nil {
			if err == io.EOF {
				// Cliente encerrou a conexão
				fmt.Println("Cliente encerrou a conexão.")
				return
			}
			fmt.Println("Erro ao receber solicitação:", err)
			return
		}

		// Processar a solicitação do cliente
		request := string(buffer[:n])
		print("Servidor recebeu: " + request)

		// Envie a resposta para o cliente
		response := "Resposta\n"
		_, err = conn.Write([]byte(response))
		if err != nil {
			fmt.Println("Erro ao enviar resposta:", err)
			return
		}
	}
	conn.Close()

}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Erro ao iniciar o servidor:", err)
		return
	}

	defer listener.Close()

	fmt.Println("Servidor TCP aguardando conexões...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Erro ao aceitar conexão:", err)
			continue
		}

		go handleConnection(conn)
	}
}
