package mapkha

type TextRange struct {
	s        int
	e        int
	EdgeType Etype
}

// Improved as Roger Peppe suggested in his tweet
// https://twitter.com/rogpeppe/status/574911374645682176
func (w *Wordcut) pathToRanges(path []Edge) []TextRange {
	w.MakeOrReuseRangesSlice(len(path))
	j := len(w.ranges) - 1
	for e := len(path) - 1; e > 0; {
		s := path[e].S
		w.ranges[j] = TextRange{s, e, path[e].EdgeType}
		j--
		e = s
	}
	return w.ranges[j+1:]
}
