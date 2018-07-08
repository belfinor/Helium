package text

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.000
// @date    2018-07-08

import (
	"io"
	"strings"
	"unicode"
)

func WordStream(rdr io.RuneReader) <-chan string {

	output := make(chan string, 10000)

	go func() {

		state := 0
		prev := '0'

		builder := strings.Builder{}

		for {

			run, _, err := rdr.ReadRune()
			if err != nil {
				break
			}

			run = unicode.ToLower(run)
			if run == 'ё' {
				run = 'е'
			}

			if unicode.IsLetter(run) || unicode.IsDigit(run) {
				if state == 2 {
					builder.WriteRune(prev)
				}
				builder.WriteRune(run)
				state = 1
				continue
			}

			if state == 0 {
				continue
			}

			if state > 1 {
				output <- builder.String()
				builder.Reset()
				state = 0
				continue
			}

			if run == '-' || run == '.' {
				prev = run
				state = 2
				continue
			}

			output <- builder.String()
			builder.Reset()
			state = 0

		}

		if state > 0 {
			output <- builder.String()
		}

		close(output)

	}()

	return output
}
