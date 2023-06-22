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

//begin mood
func (m Mood) String() string {
	switch m {
	case Neutral:
		return "neutral"
	//end mood
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
	//begin mood
	default:
		return "unknown"
	}
}
//end mood
