package makename

import (
	"embed"
	"encoding/json"
	"fmt"
	"math/rand"
)

//go:embed names.json
var f embed.FS

type names struct {
	FirstNames struct {
		Male   []string `json:"male"`
		Female []string `json:"female"`
	} `json:"first_names"`
	MiddleNames struct {
		Male   []string `json:"male"`
		Female []string `json:"female"`
	} `json:"middle_names"`
	LastNames []string `json:"last_names"`
	Initials  []string `json:"initials"`
}

type NameParam struct {
	Include, Initial bool
}

type NameGenerator struct {
	names names
}

func NewNameGenerator() *NameGenerator {
	f, err := f.Open("names.json")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var n names
	if err := json.NewDecoder(f).Decode(&n); err != nil {
		panic(err)
	}

	ng := NameGenerator{
		names: n,
	}

	return &ng
}

func (n *NameGenerator) Rand() string {
	exclude := NameParam{}
	include := NameParam{Include: true}
	includeInitial := NameParam{Include: true, Initial: true}

	chance := rand.Intn(10)
	male := rand.Intn(2) == 0

	if chance < 5 {
		return n.Generate(include, include, include, male)
	} else if chance < 7 {
		return n.Generate(include, exclude, include, male)
	} else if chance < 8 {
		return n.Generate(include, includeInitial, exclude, male)
	} else {
		return n.Generate(include, exclude, includeInitial, male)
	}
}

func (n *NameGenerator) Generate(first, middle, last NameParam, male bool) string {
	if male {
		return n.randMale(first, middle, last)
	}

	return n.randFemale(first, middle, last)
}

func (n *NameGenerator) randMale(first, middle, last NameParam) string {
	name := ""

	if first.Include {
		if first.Initial {
			name += n.names.Initials[rand.Intn(len(n.names.Initials))]
		} else {
			name += n.names.FirstNames.Male[rand.Intn(len(n.names.FirstNames.Male))]
		}
	}

	if middle.Include {
		if middle.Initial {
			name += fmt.Sprintf(" %s", n.names.Initials[rand.Intn(len(n.names.Initials))])
		} else {
			name += fmt.Sprintf(" %s", n.names.MiddleNames.Male[rand.Intn(len(n.names.MiddleNames.Male))])
		}
	}

	if last.Include {
		if last.Initial {
			name += fmt.Sprintf(" %s", n.names.Initials[rand.Intn(len(n.names.Initials))])
		} else {
			name += fmt.Sprintf(" %s", n.names.LastNames[rand.Intn(len(n.names.LastNames))])
		}
	}

	return name
}

func (n *NameGenerator) randFemale(first, middle, last NameParam) string {
	name := ""

	if first.Include {
		if first.Initial {
			name += n.names.Initials[rand.Intn(len(n.names.Initials))]
		} else {
			name += n.names.FirstNames.Female[rand.Intn(len(n.names.FirstNames.Female))]
		}
	}

	if middle.Include {
		if middle.Initial {
			name += fmt.Sprintf(" %s", n.names.Initials[rand.Intn(len(n.names.Initials))])
		} else {
			name += fmt.Sprintf(" %s", n.names.MiddleNames.Female[rand.Intn(len(n.names.MiddleNames.Female))])
		}
	}

	if last.Include {
		if last.Initial {
			name += fmt.Sprintf(" %s", n.names.Initials[rand.Intn(len(n.names.Initials))])
		} else {
			name += fmt.Sprintf(" %s", n.names.LastNames[rand.Intn(len(n.names.LastNames))])
		}
	}

	return name
}
