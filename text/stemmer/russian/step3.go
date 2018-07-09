package russian

import (
	snowballword "github.com/belfinor/Helium/text/stemmer/word"
)

// Step 3 is the removal of the derivational suffix.
//
func step3(word *snowballword.Word) bool {

	// Search for a DERIVATIONAL ending in R2 (i.e. the entire
	// ending must lie in R2), and if one is found, remove it.

	suffix, _ := word.RemoveFirstSuffixIn(word.R2start, "ост", "ость")
	if suffix != "" {
		return true
	}
	return false
}
