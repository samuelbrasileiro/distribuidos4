package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func handleConnection(conn net.Conn) {
	// Lógica para lidar com a conexão do cliente
	// Implemente aqui o processamento necessário para atender a solicitação do cliente

	// Lógica para lidar com a conexão do cliente
	buffer := make([]byte, 1024)
	for {
		// Receba a solicitação do cliente
		_, err := conn.Read(buffer)
		if err != nil {
			if err == io.EOF {
				// Cliente encerrou a conexão
				fmt.Println("Cliente encerrou a conexão.")
				return
			}
			fmt.Println("Erro ao receber solicitação")
			return
		}

		// Processar a solicitação do cliente
		// request := string(buffer[:n])
		// print("Servidor recebeu: " + request)

		// Envie a resposta para o cliente
		response := "Resposta\n"
		_, err = conn.Write([]byte(response))
		if err != nil {
			fmt.Println("Erro ao enviar resposta")
			return
		}
	}
	
	// Fecha Conexão
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			fmt.Println("Erro ao fechar conexão")
			os.Exit(0)
		}
	}(conn)

}

func main() {
	// Define o endpoint do servidor TCP
	r, err := net.ResolveTCPAddr("tcp", ":8080")
	if err != nil {
		fmt.Println("Erro ao iniciar o servidor")
		return
	}

	// Cria umlistener TCP
	listener, errListener := net.ListenTCP("tcp", r)
	if errListener != nil {
		fmt.Println("Erro ao criar um listener TCP")
		return
	}

	fmt.Println("Servidor TCP aguardando conexões...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Erro ao aceitar conexão")
			continue
		}

		go handleConnection(conn)
	}
}
