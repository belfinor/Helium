package text

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.008
// @date    2018-11-30

import (
	"io"
	"strings"
	"unicode"
)

const (
	WSO_ENDS         int = 0 // ADD dot on ent of statement
	WSO_NO_URLS      int = 1 // SKIP ulrs
	WSO_NO_NUMS      int = 2 // SKIP 123, 12.10, 10:10, 12.08.2018, 10-10 etc
	WSO_HASHTAG      int = 3 // use hashtags
	WSO_NO_XXX_COLON int = 4 // skip xxx: and XXX:
)

const (
	ST_NULL  int = 0
	ST_NUM   int = 1
	ST_ALPHA int = 2
)

type wsOpts struct {
	ends       bool
	noUrls     bool
	noNums     bool
	hashTags   bool
	NoXxxColon bool
}

func isEndStatement(run rune) bool {
	return run == '.' || run == '!' || run == '?' || run == ';' || run == '…'
}

func WordStream(rdr io.RuneReader, opts ...int) <-chan string {

	var opt wsOpts

	for _, o := range opts {
		switch o {
		case WSO_ENDS:
			opt.ends = true
		case WSO_NO_URLS:
			opt.noUrls = true
		case WSO_NO_NUMS:
			opt.noNums = true
		case WSO_HASHTAG:
			opt.hashTags = true
		case WSO_NO_XXX_COLON:
			opt.NoXxxColon = true
		}
	}

	output := make(chan string, 10000)

	go func() {

		state := 0
		prev := '0'
		st := ST_NULL

		builder := strings.Builder{}

		waitDot := false

		for {

			if state == 0 {
				st = ST_NULL
			}

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

				if unicode.IsLetter(run) {
					builder.WriteRune(run)
					state = 1
					st = st | ST_ALPHA
					waitDot = true
				} else if unicode.IsDigit(run) {
					builder.WriteRune(run)
					state = 1
					st = st | ST_NUM
					waitDot = true
				} else if run == '#' {
					state = 3
				} else if opt.ends && isEndStatement(run) {
					if waitDot {
						waitDot = false
						output <- "."
					}
				}

			case 1:

				if run == '-' || run == '.' {
					prev = run
					state = 2
				} else if unicode.IsLetter(run) {
					builder.WriteRune(run)
					st = st | ST_ALPHA
				} else if unicode.IsDigit(run) {
					builder.WriteRune(run)
					st = st | ST_NUM
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
						} else if opt.NoXxxColon && (str == "xxx" || str == "ххх") {
							builder.Reset()
							state = 0
						} else {
							if !opt.noNums || st&ST_ALPHA != 0 {
								output <- str
							}
							builder.Reset()
							state = 0
						}

					} else {

						if !opt.noNums || st&ST_ALPHA != 0 {
							output <- builder.String()
						}
						builder.Reset()

						if run == '#' {
							state = 3
						} else {
							if opt.ends && isEndStatement(run) {
								output <- "."
								waitDot = false
							}
							state = 0
						}
					}
				}

			case 2:

				if unicode.IsLetter(run) {
					builder.WriteRune(prev)
					builder.WriteRune(run)
					st = st | ST_ALPHA
					state = 1
				} else if unicode.IsDigit(run) {
					builder.WriteRune(prev)
					builder.WriteRune(run)
					st = st | ST_NUM
					state = 1
				} else {
					if !opt.noNums || st&ST_ALPHA != 0 {
						output <- builder.String()
					}
					builder.Reset()

					if prev == '.' && opt.ends {
						output <- "."
						waitDot = false
					}

					state = 0
				}

			case 3:

				if !(unicode.IsLetter(run) || unicode.IsDigit(run)) {

					str := builder.String()
					builder.Reset()

					if str != "" && opt.hashTags {
						if !opt.noNums || st&ST_ALPHA != 0 {
							waitDot = true
							output <- str
						}
					}

					if run != '#' {

						if opt.ends && isEndStatement(run) {
							output <- "."
							waitDot = false
						}
						state = 0
					} else {
						st = ST_NULL
					}
				} else {

					if unicode.IsDigit(run) {
						st = st | ST_NUM
					} else {
						st = st | ST_ALPHA
					}

					builder.WriteRune(run)
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

			if state == 4 || !opt.noNums || st&ST_ALPHA != 0 {
				output <- str
			}
		} else if state == 3 {
			str := builder.String()
			builder.Reset()

			if str != "" && opt.hashTags {
				if !opt.noNums || st&ST_ALPHA != 0 {
					output <- str
					waitDot = true
				}
			}
		}

		if opt.ends && waitDot {
			output <- "."
		}

		close(output)

	}()

	return output
}
