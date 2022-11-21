package day6

import (
	"fmt"
	"strings"

	"github.com/umbe77/aoc-2019/utils"
)

func Execute() {
	orbs := make([]string, 0)
	utils.ReadFile(fmt.Sprintf("./day6/input.txt"), func(line string) {
		orbs = append(orbs, line)
	})
	com := InitializeOrbits(orbs)
	count := CountOrbits(com)
	fmt.Printf("Part 1: %d\n", count)

	g := InitializeOrbitsGraph(orbs)
	ot := FindShortest(g, "YOU", "SAN", []string{})
	fmt.Printf("Part 2:\n%q\n-- %d\n", ot, len(ot)-3)
}

type Orbit struct {
	Planet     string
	SubPlanets []*Orbit
}

func CountOrbits(orbs map[string]string) int {
	count := 0

	for _, v := range orbs {
		parent := v
		for {
			count++
			if p, ok := orbs[parent]; ok {
				parent = p
			} else {
				break
			}
		}
	}

	return count
}

func isInPath(path []string, val string) bool {
	for _, v := range path {
		if v == val {
			return true
		}
	}
	return false
}

func FindShortest(orbs map[string][]string, start, end string, path []string) []string {
	if _, ok := orbs[start]; !ok {
		return path
	}
	path = append(path, start)
	if start == end {
		return path
	}

	shortest := make([]string, 0)

	for _, node := range orbs[start] {
		if !isInPath(path, node) {
			newPath := FindShortest(orbs, node, end, path)
			if len(newPath) > 0 && (len(shortest) == 0 || len(newPath) < len(shortest)) {
				shortest = newPath
			}
		}
	}

	return shortest
}

type Planet struct {
	Name    string
	Parents []string
}

func InitializeOrbits(orbs []string) map[string]string {
	c := make(map[string]string)
	for _, v := range orbs {
		planets := strings.Split(v, ")")
		c[planets[1]] = planets[0]
	}
	return c
}

func InitializeOrbitsGraph(orbs []string) map[string][]string {
	g := make(map[string][]string)

	var addPlanet = func(p1, p2 string) {
		if _, ok := g[p1]; ok {
			g[p1] = append(g[p1], p2)
		} else {
			g[p1] = []string{p2}
		}

		if _, ok := g[p2]; ok {
			g[p2] = append(g[p2], p1)
		} else {
			g[p2] = []string{p1}
		}
	}

	for _, v := range orbs {
		planets := strings.Split(v, ")")
		addPlanet(planets[0], planets[1])
	}

	return g
}

func InitializeOrbitsTree(orbs []string) *Orbit {
	c := make(map[string]*Orbit)
	for _, v := range orbs {
		planets := strings.Split(v, ")")
		var planet *Orbit
		var subplanet *Orbit
		if o, ok := c[planets[0]]; ok {
			planet = o
		} else {
			planet = &Orbit{
				Planet: planets[0],
			}
			c[planets[0]] = planet
		}
		if o, ok := c[planets[1]]; ok {
			subplanet = o
		} else {
			subplanet = &Orbit{
				Planet: planets[1],
			}
			c[planets[1]] = subplanet
		}
		planet.SubPlanets = append(planet.SubPlanets, subplanet)
	}
	return c["COM"]
}
