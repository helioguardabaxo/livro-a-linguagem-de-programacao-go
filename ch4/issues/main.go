package main

import (
	"fmt"
	"log"
	"os"

	"/home/helio/projetos/go/src/github.com/helioguardabaxo/livro-a-linguagem-de-programacao-go/ch4/github"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.0s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
}
