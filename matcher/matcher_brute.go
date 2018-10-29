package matcher

import (
	"time"

	"github.com/mhoc/xtern-matcher/model"
)

const (
	BruteRunFor = 30 * time.Second
)

func Brute(students model.Students, companies model.Companies) model.Matches {
	var matches model.Matches
	return matches
}
