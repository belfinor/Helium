package english

import (
	snowballword "github.com/belfinor/Helium/text/stemmer/word"
)

// Applies various transformations necessary for the
// other, subsequent stemming steps.  Most important
// of which is defining the two regions R1 & R2.
//
func preprocess(word *snowballword.Word) {

	// Clean up apostrophes
	normalizeApostrophes(word)
	trimLeftApostrophes(word)

	// Capitalize Y's that are not behaving
	// as vowels.
	capitalizeYs(word)

	// Find the two regions, R1 & R2
	r1start, r2start := r1r2(word)
	word.R1start = r1start
	word.R2start = r2start
}
