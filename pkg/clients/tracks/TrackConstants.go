package tracks

type MusicType int

const (
	Party MusicType = iota
	Pop
	Rock
	Classical
)

func (t MusicType) String() string {
	switch t {
	case Party:
		return "party"
	case Pop:
		return "pop"
	case Rock:
		return "rock"
	case Classical:
		return "classical"
	default:
		return "invalid"
	}
}
