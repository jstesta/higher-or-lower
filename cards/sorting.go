package cards

type ByFaceValue Deck

func (t ByFaceValue) Len() int           { return len(t) }
func (t ByFaceValue) Swap(i, j int)      { t[i], t[j] = t[j], t[i] }
func (t ByFaceValue) Less(i, j int) bool { return t[i].CompareTo(t[j]) == -1 }
