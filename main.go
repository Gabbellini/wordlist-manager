package main

import (
	"flag"
	"fmt"
	"log"
	TextFileHandler "wordlist-manager/text_file_handler"
)

func main() {
	filename := flag.String("path", "wordlist.txt", "")
	flag.Parse()

	for {
		fmt.Println("[0] SHOW FILE WORDS\n[1] ADD WORD\n[2] UPDATE WORD\n[3] REMOVE WORD\n[4] EXIT")
		var option int
		_, err := fmt.Scan(&option)
		if err != nil {
			log.Println("[main] Error Scan:", err)
			continue
		}

		handler, err := TextFileHandler.NewTextFileHandler(*filename)
		if err != nil {
			log.Println("[main] Error NewTextFileHandler:", err)
			continue
		}

		switch option {
		case 0:
			showWordsOnFile(handler)
		case 1:
			addWordToFile(handler)
		case 2:
			updateWordOnFile(handler)
		case 3:
			removeWordFromFile(handler)
		case 4:
			return
		default:
			fmt.Println("Invalid option.")
		}
	}
}

func showWordsOnFile(handler TextFileHandler.TextFileHandler) {
	words, err := handler.GetLines()
	if err != nil {
		log.Println("[showWordsOnFile] Error GetWords", err)
		return
	}

	fmt.Println("Words:", words)
}

func addWordToFile(handler TextFileHandler.TextFileHandler) {
	fmt.Print("New Word: ")
	var word string
	_, err := fmt.Scan(&word)
	if err != nil {
		log.Println("[addWordToFile] Error Scan:", err)
		return
	}

	err = handler.AddLine(word)
	if err != nil {
		log.Println("[addWordToFile] Error AddWord:", err)
		return
	}
}

func removeWordFromFile(handler TextFileHandler.TextFileHandler) {
	fmt.Println("Word position: ")
	var position int
	_, err := fmt.Scan(&position)
	if err != nil {
		log.Println("[removeWordFromFile] Error Scan:", err)
		return
	}

	err = handler.RemoveLine(position)
	if err != nil {
		log.Println("[removeWordFromFile] Error RemoveWord:", err)
		return
	}
}

func updateWordOnFile(handler TextFileHandler.TextFileHandler) {
	fmt.Print("Word position: ")
	var position int
	_, err := fmt.Scan(&position)
	if err != nil {
		log.Println("[updateWordOnFile] Error Scan:", err)
		return
	}

	fmt.Print("New Word: ")
	var newWord string
	_, err = fmt.Scan(&newWord)
	if err != nil {
		log.Println("[updateWordOnFile] Error Scan:", err)
		return
	}

	err = handler.UpdateLine(position, newWord)
	if err != nil {
		log.Println("[updateWordOnFile] Error UpdateWord:", err)
		return
	}
}
