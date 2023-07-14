package main

import (
	"fmt"
	"net"
)

func handleUDPConnection(conn *net.UDPConn) {
	defer conn.Close()

	buffer := make([]byte, 1024)
	for {
		// Receba a solicitação do cliente
		n, clientAddr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("Erro ao receber solicitação:", err)
			return
		}

		// Processar a solicitação do cliente
		request := string(buffer[:n])
		print("Servidor recebeu: " + request)

		// Envie a resposta para o cliente
		response := "resposta"
		_, err = conn.WriteToUDP([]byte(response), clientAddr)
		if err != nil {
			fmt.Println("Erro ao enviar resposta:", err)
			return
		}
	}
}

func main() {
	addr, err := net.ResolveUDPAddr("udp", ":8082")
	if err != nil {
		fmt.Println("Erro ao resolver endereço:", err)
		return
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println("Erro ao iniciar o servidor:", err)
		return
	}

	defer conn.Close()

	fmt.Println("Servidor UDP aguardando conexões...")

	for {
		buffer := make([]byte, 1024)
		_, _, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("Erro ao ler dados UDP:", err)
			continue
		}

		go handleUDPConnection(conn)
	}
}
