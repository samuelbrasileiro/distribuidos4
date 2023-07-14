package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	// Defina o número de processos
	numProcesses := 1

	// Iniciar vários processos
	for i := 0; i < numProcesses; i++ {
		cmd := exec.Command("./myprocess", fmt.Sprintf("%d", i))

		// Redirecionar a saída padrão do processo para a saída padrão do programa
		cmd.Stdout = os.Stdout

		// Executar o processo
		err := cmd.Start()
		if err != nil {
			fmt.Printf("Erro ao iniciar o processo %d: %v\n", i, err)
			continue
		}

		// // Aguardar o término do processo
		// err = cmd.Wait()
		// if err != nil {
		// 	fmt.Printf("Erro ao aguardar o término do processo %d: %v\n", i, err)
		// }
	}
}