package notes

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/MatheusTimmers/clnote/db"
)

func AddNote(title string) {
	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = "nvim"
	}

	tmpfile, err := os.CreateTemp("", "note-*.md")
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

	_, err = db.DB.Exec("INSERT INTO notes (title, content) VALUES ($1, $2)", title, string(content))
	if err != nil {
		fmt.Println("Erro ao salvar nota no banco:", err)
		return
	}

	fmt.Println("Nota salva com sucesso!")
}

func ListNotes() {
	rows, err := db.DB.Query("SELECT id, title, created_at FROM notes ORDER BY created_at DESC")
	if err != nil {
		fmt.Println("Erro ao buscar notas:", err)
		return
	}
	defer rows.Close()

	fmt.Println("Notas salvas:")
	for rows.Next() {
		var id int
		var title string
		var createdAt string

		err := rows.Scan(&id, &title, &createdAt)
		if err != nil {
			fmt.Println("Erro ao ler nota:", err)
			continue
		}

		fmt.Printf("[%d] %s - %s\n", id, title, createdAt)
	}
}

func GetNote(title string) {
	rows, err := db.DB.Query("SELECT id, title, content, created_at FROM notes WHERE title = $1", title)
	if err != nil {
		fmt.Println("Erro ao buscar nota:", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var title string
		var createdAt string
		var content string

		err := rows.Scan(&id, &title, &content, &createdAt)
		if err != nil {
			fmt.Println("Erro ao ler nota:", err)
			continue
		}

		fmt.Printf("[%d] %s \n %s\n \n %s\n", id, title, content, createdAt)
	}

}
