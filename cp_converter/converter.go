package cp_converter

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

func convertScssVariableDeclaration(files []os.FileInfo) {
	for _, file := range files {
		println(file.Name())
	}
}

func ConvertScssVarToCustomProps(files []os.FileInfo) {
	convertScssVariableDeclaration(files)
}

var sassVarDeclarationReg = regexp.MustCompile(`\$([\w-]+\d*):`)
var sassVarAssignationReg = regexp.MustCompile(`\$([\w-]+\d*)`)

func isDeclaration(line string) bool {
	return sassVarDeclarationReg.MatchString(line)
}

func isAnAssignation(line string) bool {
	return sassVarAssignationReg.MatchString(line)
}

func ConvertFile(path string) {
	file, err := os.OpenFile(path, os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	outputFile, err := os.Create("output.tmp")
	if err != nil {
		fmt.Println("Error creating temp file:", err)
		return
	}
	defer outputFile.Close()

	scanner := bufio.NewScanner(file)
	writer := bufio.NewWriter(outputFile)

	for scanner.Scan() {
		line := scanner.Text()
		println(line)
		//var newLine string
		if isDeclaration(line) {
			// convert declaration to custom property
			line = sassVarDeclarationReg.ReplaceAllString(line, `--$1:`)
		}
		if isAnAssignation(line) {
			// convert assignation to custom property
			line = sassVarAssignationReg.ReplaceAllString(line, `var(--$1)`)
		}
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			fmt.Println("Error writing to temp file:", err)
			return
		}

	}

	err = writer.Flush()
	if err != nil {
		fmt.Println("Error flushing writer:", err)
		return
	}

	file.Close()
	outputFile.Close()

	err = os.Rename("output.tmp", path)
	if err != nil {
		fmt.Println("Error renaming temp file:", err)
		return
	}

	fmt.Println("File updated successfully.")

}
