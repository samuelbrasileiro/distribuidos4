package main

import (
	"fmt"
	"math/rand"
	"time"
	"math"
	"os"
	"strconv"
	"encoding/json"
	"github.com/streadway/amqp"
	"github.com/google/uuid"
)

type Pokemon struct {
	Name  string
	Type  string
	Level int
}

func errorCheck(msg string, err error) {
	if err != nil {
		fmt.Println("%s: %s", msg, err)
	}
}

func makeSingleRequest(ch *amqp.Channel, q amqp.Queue, msgs <-chan amqp.Delivery) {
	// Gerador de números aleatórios com base na hora atual
	rand.Seed(time.Now().UnixNano())

	// Gerar um número aleatório entre 0 e 149
	randomNumber := rand.Intn(150)

	correlationID := uuid.New().String()

	var err error

	// Enviar o número do Pokémon para a fila
	err = ch.Publish(
		"",
		"pokedex_requests", // Routing key (nome da fila)
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(strconv.Itoa(randomNumber)),
			ReplyTo: q.Name,
			CorrelationId: correlationID,
		},
	)

	if err != nil {
		fmt.Println("Erro ao publicar mensagem: %v", err)
		return
	}

	response := <-msgs
	var pokemon Pokemon
	err = json.Unmarshal(response.Body, &pokemon)
	if err != nil {
		fmt.Println("Erro ao decodificar resposta JSON:", err)
	} else {
		fmt.Printf("Name: %s\n", pokemon.Name)
		fmt.Printf("Type: %s\n", pokemon.Type)
		fmt.Printf("Level: %d\n", pokemon.Level)
		fmt.Println("------------------------")
	}
}

func main() {
	args := os.Args[1:2] // Ignorar o primeiro argumento, que é o nome do executável

	// Conectar ao servidor RabbitMQ
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	errorCheck("Erro ao se conectar com servidor de mensageria", err)
	defer conn.Close()

	ch, err := conn.Channel()
	errorCheck("Erro ao estabelecer um canal de comunicação com o servidor de mensageria", err)
	defer ch.Close()

	// declara a fila para as respostas
	q, err := ch.QueueDeclare(
		"pokedex_response",
		false,
		false,
		true,
		false,
		nil,
	)

	// cria servidor da fila de response
	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil)
	errorCheck("Falha ao registrar o servidor no broker", err)

	numRequests := 10000
	RTTs := make([]time.Duration, numRequests)
	var totalRTT time.Duration

	if args[0] == "0" {
		for j := 0; j < numRequests; j++ {
			start := time.Now()

			makeSingleRequest(ch, q, msgs)

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
			makeSingleRequest(ch, q, msgs)
		}
	}
}
