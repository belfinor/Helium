package text

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.003
// @date    2018-08-08

import (
	"io"
	"strings"
	"unicode"
)

const (
	WSO_ENDS int = 0
)

type wsOpts struct {
	ends bool
}

func isEndStatement(run rune) bool {
	return run == '.' || run == '!' || run == '?' || run == '…'
}

func WordStream(rdr io.RuneReader, opts ...int) <-chan string {

	var opt wsOpts

	for _, o := range opts {
		switch o {
		case WSO_ENDS:
			opt.ends = true
		}
	}

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
						if opt.ends && isEndStatement(run) {
							output <- "."
						}
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

					if prev == '.' && opt.ends {
						output <- "."
					}

					state = 0
				}

			case 3:

				if !(unicode.IsLetter(run) || unicode.IsDigit(run)) {
					if run != '#' {
						if opt.ends && isEndStatement(run) {
							output <- "."
						}
						state = 0
					}
				}

			}
		}

		if state == 1 || state == 2 {
			output <- builder.String()
		}

		if opt.ends {
			output <- "."
		}

		close(output)

	}()

	return output
}
