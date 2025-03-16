package main

import (
	"fmt"
	"os"
	"os/exec"

	_ "github.com/lib/pq"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Uso: note <comando> [argumentos]")
		os.Exit(1)
	}

	command := os.Args[2]

	fmt.Printf("Numero de argumentos %d \n", len(os.Args))
	switch command {
	case "add":
		// add new entry
		if len(os.Args) != 4 {
			fmt.Println("Uso: note add <titulo>")
			os.Exit(1)
		}

		addNote(os.Args[3])
	}

}

func addNote(title string) {
	editor := os.Getenv("CLNOTE_EDITOR")
	if editor == "" {
		editor = "nvim" // Fallback
	}

	tmpfile, err := os.CreateTemp("", title + "-*.md")
	if err != nil {
		fmt.Println("Erro ao criar arquivo temporário:", err)
		return
	}
	defer os.Remove(tmpfile.Name())

	cmd := exec.Command(editor, tmpfile.Name())
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		fmt.Println("Erro ao abrir o editor:", err)
		return
	}

	content, err := os.ReadFile(tmpfile.Name())
	if err != nil {
		fmt.Println("Erro ao ler a nota:", err)
		return
	}

	if len(content) == 0 {
		fmt.Println("Nota vazia. Cancelando operação.")
		return
	}

	fmt.Println("Nota salva com sucesso!")
	fmt.Println("Conteúdo da nota:\n", string(content))
}
