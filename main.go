package main

import (
	//"flag"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Joeljm1/todo-go/todo"
)

func createJsonFileIfNotExist(filename string) error {
	_, err := os.Stat(filename)
	if err != nil {
		if os.IsNotExist(err) {
			err := os.WriteFile(filename, []byte("[]"), 0644)
			if err != nil {
				return err
			}
			return nil
		}
		return err
	}
	return nil
}

func help() {
	fmt.Println("available commands are add,delete,update,help,complete")
}

func main() {
	data := todo.TodoSlice{}
	filename := "E:\\code\\go\\todoCLI\\data.json"
	err := createJsonFileIfNotExist(filename)
	if err != nil {
		log.Fatal(err)
	}
	err = todo.Load(filename, &data)
	if err != nil {
		log.Fatal(err)
	}
	if len(os.Args) < 2 { // no commandline arg
		data.List()
		return
	}
	switch os.Args[1] {
	case "":
		help()
	case "help":
		help()
	case "add":
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Enter title")
		// fmt.Scanln(&title)
		title, _ := reader.ReadString('\n')
		title = strings.TrimSpace(title)
		fmt.Println("Enter descripttion")
		// fmt.Scanln(&desc)
		desc, _ := reader.ReadString('\n')
		desc = strings.TrimSpace(desc)
		data.Add(title, desc)
		err = data.Save(filename)
		if err != nil {
			log.Fatal(err)
		}
	case "delete":
		var index int
		fmt.Println("Enter index to delete")
		fmt.Scan(&index)
		index--
		err := data.Delete(index)
		if err != nil {
			log.Fatal(err)
		}
		err = data.Save(filename)
		if err != nil {
			log.Fatal(err)
		}
	case "list":
		data.List()
	case "complete":
		var index int
		fmt.Println("Enter index to complete")
		fmt.Scan(&index)
		index--
		err := data.Complete(index)
		if err != nil {
			log.Fatal(err)
		}
		err = data.Save(filename)
		if err != nil {
			log.Fatal(err)
		}
	default:
		fmt.Println("Wrong comand")
		help()

	}
}
