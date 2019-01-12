package main

import(
	"fmt"
	"io/ioutil"
	"bytes"
	"flag"
	"strings"
)

//Flags
var numbered_lines bool
var fileNames string

func init(){
	flag.BoolVar(&numbered_lines, "b", false, "number all output lines")
	flag.BoolVar(&numbered_lines, "number", false,  "number all output lines")
	flag.StringVar(&fileNames, "f", "", "Path to files seperated by spaces")
	flag.StringVar(&fileNames, "files", "", "Path to files seperated by spaces")
}

func parseFileNames(files string) (results []string) {
	tempt := strings.Split(files, " ")
	for _, item := range tempt {
		results = append(results, strings.TrimSpace(item))
	}
	return
}

func openFile(filename string) ([][]byte, error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return [][]byte{[]byte("")}, err
	}

	return bytes.Split(b, []byte("\n")), nil
}

func printToScreen( lines [][]byte){
	num_of_lines := len(lines) - 1

	for line_number, line := range(lines){
		if line_number != num_of_lines {
			fmt.Println(string(line))
		} else {
			if len(line) == 0 {
				// suppressing the last line if empty
				continue
			}
			fmt.Println(string(line))
		}
	}
}

func suppressBlankLines(lines [][]byte) [][]byte {
	var lines2 [][]byte
	for index, line := range lines {
		if index == 0 {
			lines2 = append(lines2, line)
			continue
		}
			
		if len(line) == 0 && len(lines[index-1]) ==0 {
			// Skipping this line
			continue
		}
	
		lines2 = append(lines2, line)
	}
	return lines2
}

func addShowEnds(lines [][]byte) [][]byte {
	buf := bytes.NewBuffer([]byte(""))

	num_of_lines := len(lines) - 1

	var result [][]byte
	for index, line := range lines {
		if index == num_of_lines {
			if len(line) == 0 {
				continue
			}
		}
		buf.Write(line)
		buf.Write([]byte("$"))

		result = append(result, []byte(buf.String()))
		buf.Reset()
	}
	return result
}

func addLineNumbers(lines [][]byte) [][]byte {
	buf := bytes.NewBuffer([]byte(""))

	num_of_lines := len(lines) - 1

	var result [][]byte
	for index, line := range lines {
		if index == num_of_lines {
			if len(line) == 0 {
				continue
			}
		}
	
		prepend_text := fmt.Sprintf("    %d  ", index + 1)
		buf.Write([]byte(prepend_text))
		buf.Write(line)
	
		result = append(result, []byte(buf.String()))
		buf.Reset()  // Clearing the buffer
	}

	return result
}

func main(){
	flag.Parse()

	files := parseFileNames(fileNames)

	for _, file := range files {
		lines, err := openFile(file)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		printToScreen(lines)
	}
}
