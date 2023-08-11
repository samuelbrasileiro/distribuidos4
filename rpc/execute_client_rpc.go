package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	// Defina o número de processos
	numProcesses := 80

	// Iniciar vários processos
	for i := 0; i < numProcesses; i++ {
		// No windows é necessário passar o caminho todo, no linux um ./ resolve linux >>>
		cmd := exec.Command("C:/Users/washi/Documents/Go/distribuidos4/rcp/client_rcp.exe", fmt.Sprintf("%d", i))

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