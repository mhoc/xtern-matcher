package matcher

import (
	"time"

	"github.com/mhoc/xtern-matcher/model"
)

const (
	EvolveEjectionCount    = 0.1
	EvolveFitnessThreshold = 0.6
	EvolveRunFor           = 30 * time.Second
)

/**
 * FindMatchesEvolve is an evolution matcher.
 *
 * We define the "fitness" of a match to be the distance at which a company ranked a student; if they're
 * assigned a student ranked #0, the fitness is 1. If they're assigned a student ranked 2, the fitness
 * would be 0 (assuming there are 3 global maximum ranks across all companies).
 *
 * In addition to a per-match fitness, we consider the Global Average of all match fitnesses to be the
 * goal the evolution algorithm is trying to optimize for.
 *
 * The algorithm selects a set of matches essentially at random, and calculates the per-match
 * fitness and global fitness. It then ejects all of the scores below a given fitness threshold
 * and re-calculates. It repeats this process until it arrives at a global fitness that is higher
 * than the last generation.
 *
 * The next generation repeats this process, ideally with more matches above the fitness threshold.
 *
 * Additionally, this algorithm implements a fitness reaper; every N generations represents a Reap
 * Generation in which 1 match above the fitness threshold is ejected at random along with the unfit
 * matches. The same process applies here; if it can find a higher global fitness then the generation
 * is saved, otherwise it is reset back and a normal generation continues.
 *
 * This process can run indefinitely.
 */
func Evolve(students model.Students, companies model.Companies) model.Matches {
	var matches model.Matches
	return matches
}
