package day12

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/umbe77/aoc-2019/utils"
)

var planetNames = []string{
	"Io",
	"Europa",
	"Ganymede",
	"Callisto",
}

func Execute() {
	planets := make([]*Planet, 0)
	nameIndex := 0
	utils.ReadFile(fmt.Sprintf("./day12/input.txt"), func(line string) {
		planets = append(planets, ParsePlanet(line, planetNames[nameIndex]))
		nameIndex++
	})

	fmt.Printf("Part 1: %d\n", TotalEnergy(planets, 1000))
}

func ParsePlanet(data, name string) *Planet {
	posPair := strings.Split(strings.TrimRight(strings.TrimLeft(data, "<"), ">"), ", ")
	var x, y, z int
	for i, pair := range posPair {
		v, _ := strconv.Atoi(strings.Split(pair, "=")[1])
		switch i {
		case 0:
			x = v
		case 1:
			y = v
		case 2:
			z = v
		}
	}
	return &Planet{
		Name: name,
		Pos: Vec{
			X: x,
			Y: y,
			Z: z,
		},
		Vel: Vec{0, 0, 0},
	}
}

type Vec struct {
	X int
	Y int
	Z int
}

type Planet struct {
	Name string
	Pos  Vec
	Vel  Vec
}

func Add(v1, v2 Vec) Vec {
	return Vec{
		X: v1.X + v2.X,
		Y: v1.Y + v2.Y,
		Z: v1.Z + v2.Z,
	}
}

func printPlanets(planets []*Planet) {
	for _, p := range planets {
		fmt.Printf("%+v\n", p)
	}
	fmt.Println("================")
}

func ApplyGravity(v1, v2 int) int {
	if v1 < v2 {
		return 1
	} else if v1 > v2 {
		return -1
	}
	return 0
}

func TotalEnergy(planets []*Planet, totalStep int) int {
	for i := 0; i < totalStep; i++ {
		resultantVel := make(map[string]Vec)
		for _, planet := range planets {
			resultantVel[planet.Name] = planet.Vel
			for _, other := range planets {
				if planet.Name != other.Name {
					vel := Vec{
						X: ApplyGravity(planet.Pos.X, other.Pos.X),
						Y: ApplyGravity(planet.Pos.Y, other.Pos.Y),
						Z: ApplyGravity(planet.Pos.Z, other.Pos.Z),
					}
					resultantVel[planet.Name] = Add(resultantVel[planet.Name], vel)
				}
			}
		}
		for _, planet := range planets {
			planet.Vel = resultantVel[planet.Name]
			planet.Pos = Add(planet.Pos, planet.Vel)
		}
	}

	totalE := 0
	for _, planet := range planets {
		potential := utils.Abs(planet.Pos.X) + utils.Abs(planet.Pos.Y) + utils.Abs(planet.Pos.Z)
		kinetic := utils.Abs(planet.Vel.X) + utils.Abs(planet.Vel.Y) + utils.Abs(planet.Vel.Z)
		totalE += potential * kinetic
	}
	return totalE
}
