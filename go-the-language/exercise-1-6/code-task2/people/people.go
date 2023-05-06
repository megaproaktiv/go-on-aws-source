package people

type Mood int

const (
    Neutral Mood = iota
    Excited
    Happy
    Curious
    Grateful
	Relaxed
)

type Person struct {
    Name string
    Mood Mood
}

func (m Mood) String() string {
	switch m {
	case Neutral:
		return "neutral"
	case Excited:
		return "excited"
	case Happy:
		return "happy"
	case Curious:
		return "curious"
	case Grateful:
		return "grateful"
	case Relaxed:
		return "relaxed"
	default:
		return "unknown"
	}
}
