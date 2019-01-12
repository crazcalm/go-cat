package main

import(
	"fmt"
	"io/ioutil"
	"bytes"
)

func cat(filename string) [][]byte {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	return bytes.Split(b, []byte("\n"))
}

func print_to_screen( lines [][]byte){
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

func suppress_blank_lines(lines [][]byte) [][]byte {
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

func add_show_ends(lines [][]byte) [][]byte {
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

func add_line_numbers(lines [][]byte) [][]byte {
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
	lines := cat("notes")
	lines = suppress_blank_lines(lines)
	lines = add_show_ends(lines)
	print_to_screen(lines)
}
