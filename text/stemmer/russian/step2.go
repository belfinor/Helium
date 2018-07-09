package russian

import (
	snowballword "github.com/belfinor/Helium/text/stemmer/word"
)

// Step 2 is the removal of the "и" suffix.
//
func step2(word *snowballword.Word) bool {
	suffix, _ := word.RemoveFirstSuffixIn(word.RVstart, "и")
	if suffix != "" {
		return true
	}
	return false
}
