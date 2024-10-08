package helper

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"golang.org/x/exp/rand"
)

const (
	// Alphabet ...
	Alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	// Numerals ...
	Numerals = "0123456789"

	// Ascii ...
	Ascii = Alphabet + Numerals + "~!@#$%^&*()-_+={}[]\\|<,>.?/\"';:`"
)

// GeneratorExprRanges ...
type GeneratorExprRanges [][]byte

func seedAndReturnRandom(n int) int {
	return rand.Intn(n)
}

func alphabetSlice(from, to byte) (string, error) {
	leftPos := strings.Index(Ascii, string(from))
	rightPos := strings.LastIndex(Ascii, string(to))
	if leftPos > rightPos {
		return "", fmt.Errorf("invalid range specified: %s-%s", string(from), string(to))
	}
	return Ascii[leftPos:rightPos], nil
}

func replaceWithGenerated(s *string, expresion string, ranges [][]byte, length int) error {
	var alphabet string
	for _, r := range ranges {
		switch string(r[0]) + string(r[1]) {
		case `\w`:
			alphabet += Ascii
		case `\d`:
			alphabet += Numerals
		default:
			if slice, err := alphabetSlice(r[0], r[1]); err != nil {
				return err
			} else {
				alphabet += slice
			}
		}
	}
	if len(alphabet) == 0 {
		return fmt.Errorf("empty range in expresion: %s", expresion)
	}
	result := make([]byte, length)
	for i := 0; i <= length-1; i++ {
		result[i] = alphabet[seedAndReturnRandom(len(alphabet))]
	}
	*s = strings.Replace(*s, expresion, string(result), 1)
	return nil
}

func findExpresionPos(s string) GeneratorExprRanges {
	rangeExp, _ := regexp.Compile(`([\\]?[a-zA-Z0-9]\-?[a-zA-Z0-9]?)`)
	matches := rangeExp.FindAllStringIndex(s, -1)
	result := make(GeneratorExprRanges, len(matches))
	for i, r := range matches {
		result[i] = []byte{s[r[0]], s[r[1]-1]}
	}
	return result
}

func rangesAndLength(s string) (string, int, error) {
	expr := s[0:strings.LastIndex(s, "{")]
	length, err := parseLength(s)
	return expr, length, err
}

func parseLength(s string) (int, error) {
	lengthStr := string(s[strings.LastIndex(s, "{")+1 : len(s)-1])
	if l, err := strconv.Atoi(lengthStr); err != nil {
		return 0, fmt.Errorf("unable to parse length from %v", s)
	} else {
		return l, nil
	}
}

// Generate random string
func Generate(template string) (string, error) {
	result := template
	generatorsExp, _ := regexp.Compile(`\[([a-zA-Z0-9\-\\]+)\](\{([0-9]+)\})`)
	matches := generatorsExp.FindAllStringIndex(template, -1)
	for _, r := range matches {
		ranges, length, err := rangesAndLength(template[r[0]:r[1]])
		if err != nil {
			return "", err
		}
		positions := findExpresionPos(ranges)
		if err := replaceWithGenerated(&result, template[r[0]:r[1]], positions, length); err != nil {
			return "", err
		}
	}
	return result, nil
}
