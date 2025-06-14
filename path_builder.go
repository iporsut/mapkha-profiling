package mapkha

var globalContext = &EdgeBuildingContext{}

func (w *Wordcut) buildPath(textRunes []rune, edgeBuilders []EdgeBuilder) []Edge {
	w.MakeOrReusePathSlice(len(textRunes) + 1)
	w.path[0] = Edge{S: 0, EdgeType: INIT, WordCount: 0, UnkCount: 0}
	leftBoundary := 0
	for i, ch := range textRunes {
		var bestEdge Edge
		for _, edgeBuilder := range edgeBuilders {
			globalContext.runes = textRunes
			globalContext.Path = w.path
			globalContext.I = i
			globalContext.Ch = ch
			globalContext.LeftBoundary = leftBoundary
			globalContext.BestEdge = bestEdge

			edge, found := edgeBuilder.Build(globalContext)

			if found && ((bestEdge == Edge{}) || edge.IsBetterThan(&bestEdge)) {
				bestEdge = edge
			}
		}

		if bestEdge.EdgeType == 0 {
			panic("bestEdge must not be nil")
		}

		if bestEdge.EdgeType != UNK {
			leftBoundary = i + 1
		}

		w.path[i+1] = bestEdge
	}
	return w.path
}
