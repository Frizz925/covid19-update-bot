package country

type Country string

const (
	ID Country = "id"
	JP Country = "jp"
)

var countryNameMap = map[Country]string{
	ID: "Indonesia",
	JP: "Japan",
}

func (c Country) ID() string {
	return string(c)
}

func (c Country) Name() string {
	if v, ok := countryNameMap[c]; ok {
		return v
	}
	return c.String()
}

func (c Country) String() string {
	return string(c)
}
