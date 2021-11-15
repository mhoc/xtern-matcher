package loader

import (
	"fmt"

	"github.com/mhoc/xtern-matcher/model"
)

type Loader interface {
	Companies() (model.Companies, error)
	Students() (model.Students, error)
}

func Get(id string, args map[string]string) Loader {
	switch id {
	case "csv":
		return NewCSV(args)
	default:
		panic(fmt.Sprintf("no data loader found with id %v", id))
	}
}
