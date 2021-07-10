package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func pickElements(args []string) (string, error) {
	var (
		source        string
		baseIndices   string
		containsComma bool
		indices       []string
		elements      []string
		result        []string
	)

	switch len(args) {
	case 1:
		fmt.Print("source: \n> ")
		fmt.Scanln(&source)
		fmt.Print("indices: \n> ")
		fmt.Scanln(&baseIndices)
	case 3:
		source = args[1]
		baseIndices = args[2]
	default:
		return "", errors.New("Wrong number of arguments.\n\n% elgram <source> <indices>\n")
	}

	containsComma = strings.Contains(baseIndices, ",")

	elements = strings.Split(source, "")
	if containsComma {
		indices = strings.Split(baseIndices, ",")
	} else {
		indices = strings.Split(baseIndices, "")
	}

	for _, indexStr := range indices {
		index, err := strconv.Atoi(indexStr)
		if err != nil {
			return "", errors.New("Illegal index contained.\n2nd argument must contain of only number.")
		}

		if index > len(elements) {
			return "", errors.New("Illegal index `" + indexStr + "` contained.")
		}

		result = append(result, elements[index-1])
	}

	return strings.Join(result, ""), nil
}

func main() {
	result, err := pickElements(os.Args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(result)

	os.Exit(0)
}
