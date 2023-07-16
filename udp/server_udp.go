package main

import (
	"fmt"
	"net"
)

func handleUDPConnection(conn *net.UDPConn, addr *net.UDPAddr, req []byte, n int) {
	// Processar a solicitação do cliente
	// request := string(req[:n])
	// print("Servidor recebeu: " + request + "\n")

	// Envie a resposta para o cliente
	response := "resposta"
	_, err := conn.WriteToUDP([]byte(response), addr)
	if err != nil {
		fmt.Println("Erro ao enviar resposta:", err)
		return
	}

}

func main() {
	// Define o endpont do servidor UDP
	addr, err := net.ResolveUDPAddr("udp", ":8080")
	if err != nil {
		fmt.Println("Erro ao resolver endereço")
		return
	}

	// Prepara o endpoint UDP para receber requests
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println("Erro ao iniciar o servidor")
		return
	}

	// Fecha conn
	defer conn.Close()

	fmt.Println("Servidor UDP aguardando conexões...")

	for {
		req := make([]byte, 1024)

		// Recebe request
		n, clientAddr, err := conn.ReadFromUDP(req)
		if err != nil {
			fmt.Println("Erro ao ler dados UDP")
			continue
		}

		go handleUDPConnection(conn, clientAddr, req, n)
	}
}
