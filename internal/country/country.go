package country

type Country string

const (
	ID Country = "id"
	JP Country = "jp"
)

type countryExt struct {
	Name  string
	Emoji string
}

var countryExtMap = map[Country]*countryExt{
	ID: {
		Name:  "Indonesia",
		Emoji: "🇮🇩",
	},
	JP: {
		Name:  "Japan",
		Emoji: "🇯🇵",
	},
}

func (c Country) ID() string {
	return string(c)
}

func (c Country) Name() string {
	if ext, ok := c.getExtended(); ok {
		return ext.Name
	}
	return c.String()
}

func (c Country) Emoji() string {
	if ext, ok := c.getExtended(); ok {
		return ext.Emoji
	}
	return c.String()
}

func (c Country) String() string {
	return string(c)
}

func (c Country) getExtended() (*countryExt, bool) {
	v, ok := countryExtMap[c]
	return v, ok
}
