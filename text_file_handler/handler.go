package TextFileHandler

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
)

type TextFileHandler interface {
	// GetLines Return a string slice that contains all the lines in the file.
	GetLines() ([]string, error)

	//AddLine Add a new line in the file.
	AddLine(line string) error

	//RemoveLine remove a line from the file by the line position.
	RemoveLine(position int) error

	//UpdateLine update a line on the file by the line position.
	UpdateLine(position int, newLine string) error
}

type textFileHandler struct {
	filename string
}

func NewTextFileHandler(filename string) (TextFileHandler, error) {
	return &textFileHandler{filename: filename}, nil
}

func (f textFileHandler) GetLines() ([]string, error) {
	file, err := os.Open(f.filename)
	if err != nil {
		log.Println("[GetLines] Error Open", err)
		return nil, err
	}

	lines, err := f.convertFileLinesIntoSlice(file)
	if err != nil {
		log.Println("[GetLines] Error convertFileLinesIntoSlice", err)
		return nil, err
	}

	err = file.Close()
	if err != nil {
		log.Println("[GetLines] Error Close", err)
		return nil, err
	}

	return lines, err
}

func (f textFileHandler) AddLine(line string) error {
	file, err := os.OpenFile(f.filename, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Println("[AddLine] Error AddLine", err)
		return err
	}

	lines, err := f.convertFileLinesIntoSlice(file)
	if err != nil {
		log.Println("[AddLine] Error convertFileLinesIntoSlice", err)
		return err
	}
	lines = append(lines, line)

	err = os.WriteFile(f.filename, []byte(strings.Join(lines, "\n")), 0666)
	if err != nil {
		log.Println("[AddLine] Error WriteString", err)
		return err
	}

	err = file.Close()
	if err != nil {
		log.Println("[AddLine] Error Close", err)
		return err
	}

	return nil
}

func (f textFileHandler) RemoveLine(position int) error {
	file, err := os.OpenFile(f.filename, os.O_RDWR, 0666)
	if err != nil {
		log.Println("[RemoveLine] Error AddLine", err)
		return err
	}

	lines, err := f.convertFileLinesIntoSlice(file)
	if err != nil {
		log.Println("[RemoveLine] Error convertFileLinesIntoSlice", err)
		return err
	}

	lines = append(lines[:position], lines[position+1:]...)

	err = os.WriteFile(f.filename, []byte(strings.Join(lines, "\n")), 0666)
	if err != nil {
		log.Println("[RemoveLine] Error WriteString", err)
		return err
	}

	err = file.Close()
	if err != nil {
		log.Println("[RemoveLine] Error Close", err)
		return err
	}

	return nil
}

func (f textFileHandler) UpdateLine(position int, newLine string) error {
	file, err := os.OpenFile(f.filename, os.O_RDWR, 0666)
	if err != nil {
		log.Println("[UpdateLine] Error AddLine", err)
		return err
	}

	lines, err := f.convertFileLinesIntoSlice(file)
	if err != nil {
		log.Println("[UpdateLine] Error convertFileLinesIntoSlice", err)
		return err
	}

	lines[position] = newLine + "\n"

	err = os.WriteFile(f.filename, []byte(strings.Join(lines, "\n")), 0666)
	if err != nil {
		log.Println("[UpdateLine] Error WriteString", err)
		return err
	}

	err = file.Close()
	if err != nil {
		log.Println("[UpdateLine] Error Close", err)
		return err
	}

	return nil
}

// convertFileLinesIntoSlice convert the file lines into a string slice.
func (f textFileHandler) convertFileLinesIntoSlice(file io.Reader) ([]string, error) {
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
		if err := scanner.Err(); err != nil {
			return nil, err
		}
	}

	return lines, nil
}
