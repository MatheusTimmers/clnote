package main

import (
	"fmt"
	"os"

	"github.com/MatheusTimmers/clnote/db"
	"github.com/MatheusTimmers/clnote/notes"

	_ "github.com/lib/pq"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Uso: note <comando> [argumentos]")
		os.Exit(1)
	}

  db.InitDB()
  
	command := os.Args[2]

	fmt.Printf("Numero de argumentos %d \n", len(os.Args))
	switch command {
	case "add":
		// add new entry
		if len(os.Args) != 4 {
			fmt.Println("Uso: note add <titulo>")
			os.Exit(1)
		}

		notes.AddNote(os.Args[3])

  case "list":
    // list all entrys
    if (len(os.Args) != 3) {
      fmt.Println("Uso: note list")
      os.Exit(1)
    }
    
    notes.ListNotes()

  case "get":
    // get a entry in db
    if len(os.Args) != 4 {
			fmt.Println("Uso: note get <titulo>")
			os.Exit(1)
		}

		notes.GetNote(os.Args[3])
    
	}
}
