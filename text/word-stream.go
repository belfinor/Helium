package text

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.004
// @date    2018-08-10

import (
	"io"
	"strings"
	"unicode"
)

const (
	WSO_ENDS    int = 0
	WSO_NO_URLS int = 1
)

type wsOpts struct {
	ends   bool
	noUrls bool
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
		case WSO_NO_URLS:
			opt.noUrls = true
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

					if run == ':' {

						str := builder.String()
						if str == "http" || str == "https" {

							run, _, err = rdr.ReadRune()
							if err != nil || run != '/' {
								output <- str
								builder.Reset()
								state = 0
							} else {
								builder.WriteRune(':')
								builder.WriteRune('/')
								state = 4
							}
						} else {
							output <- str
							builder.Reset()
							state = 0
						}

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

			case 4:
				if run == ' ' || run == '\t' || run == '\n' || run == '\r' {

					str := builder.String()
					builder.Reset()

					if !opt.noUrls {

						str = strings.TrimRight(str, ".,!?\"'")

						output <- str
					}

					state = 0
				} else {
					builder.WriteRune(run)
				}
			}
		}

		if state == 1 || state == 2 || (state == 4 && !opt.noUrls) {

			str := builder.String()
			if state == 4 {
				str = strings.TrimRight(str, ".,!?\"'")
			}
			output <- str
		}

		if opt.ends {
			output <- "."
		}

		close(output)

	}()

	return output
}
