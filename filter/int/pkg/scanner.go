package pkg

import (
	"bufio"
	"os"
)

func Scanner(file *os.File) ([]string, error) {

	var txtLines []string
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		txtLines = append(txtLines, scanner.Text())
	}

	file.Close()

	return txtLines, nil

}
