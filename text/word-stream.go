package text

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.001
// @date    2018-08-06

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

			switch state {
			case 0:

				if unicode.IsLetter(run) || unicode.IsDigit(run) {
					builder.WriteRune(run)
					state = 1
				} else if run == '#' {
					state = 3
				}

			case 1:

				if run == '-' || run == '.' {
					prev = run
					state = 2
				} else if unicode.IsLetter(run) || unicode.IsDigit(run) {
					builder.WriteRune(run)
				} else {
					output <- builder.String()
					builder.Reset()
					if run == '#' {
						state = 3
					} else {
						state = 0
					}
				}

			case 2:

				if unicode.IsLetter(run) || unicode.IsDigit(run) {
					builder.WriteRune(prev)
					builder.WriteRune(run)
					state = 1
				} else {
					output <- builder.String()
					builder.Reset()
					state = 0
				}

			case 3:

				if !(unicode.IsLetter(run) || unicode.IsDigit(run)) {
					if run != '#' {
						state = 0
					}
				}

			}
		}

		if state == 1 || state == 2 {
			output <- builder.String()
		}

		close(output)

	}()

	return output
}
