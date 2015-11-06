package parser

// Parser parses stats from uWSGI server.
type Parser interface {
	// Parse stats from the given url.
	Parse(url string) (map[string]interface{}, error)
}

// StatsParser is a Parser that parses uWSGI stats server.
type StatsParser struct {
}

// NewStatsParser constructs a StatsParser.
func NewStatsParser() Parser {
	return &StatsParser{}
}

// Parse uWSGI stats from given url.
func (p *StatsParser) Parse(url string) (map[string]interface{}, error) {
	return nil, nil
}
