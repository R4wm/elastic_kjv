package elastic_kjv

const lastOldTestamentVerseNum = 23145

// ElasticVerse contents for bible verse.
type ElasticVerse struct {
	LinearOrderedVerse int    `json:"linearOrderedVerse"`
	Testament          string `json:"testament"`
	Chapter            int    `json:"chapter"`
	Book               string `json:"book"`
	Verse              int    `json:"verse"`
	Text               string `json:"text"`
}

// GetTestament Updates Testament value to be old or new.
func (e *ElasticVerse) GetTestament() {

	// Determine if verse is old or new testament.
	if e.Verse > lastOldTestamentVerseNum {
		e.Testament = "new"
		return
	}

	e.Testament = "old"
}

type elasticBook struct {
	Book string
}
