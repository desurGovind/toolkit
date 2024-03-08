package toolkit

import "crypto/rand"

const randomStringSource = "abcdfgABC"

// Tools is the type used to instantiate this module.
type Tools struct{}

// Uppercase will export the fn to other modules
// RandomString returns a string of random chars of len
// as src of the string

func (t *Tools) RandomString(n int) string {

	s, r := make([]rune, n), []rune(randomStringSource)

	for i := range s {
		p, _ := rand.Prime(rand.Reader, len(r))
		x, y := p.Uint64(), uint64(len(r))
		s[i] = r[x%y]

	}

	return string(s)
}
