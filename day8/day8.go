package day8

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/umbe77/aoc-2019/utils"
)

func Execute() {
	var data string
	utils.ReadFile(fmt.Sprintf("./day8/input.txt"), func(line string) {
		data = line
	})
	fmt.Printf("Part 1: %d\n", Part1(data, 25, 6))
	fmt.Printf("Part 2: %d\n", Part2(data, 25, 6))
}

type Layer struct {
	Zeros int
	Ones  int
	Twos  int
}

func Part1(data string, width, height int) int {
	image := make([]Layer, 0)
	layerSize := width * height
	l := Layer{
		Zeros: 0,
		Ones:  0,
		Twos:  0,
	}
	for i, px := range data {
		switch string(px) {
		case "0":
			l.Zeros++
		case "1":
			l.Ones++
		case "2":
			l.Twos++
		}
		if (i+1)%layerSize == 0 {
			image = append(image, l)
			l = Layer{
				Zeros: 0,
				Ones:  0,
				Twos:  0,
			}
		}
	}
	sort.Slice(image, func(i, j int) bool {
		return image[i].Zeros < image[j].Zeros
	})
	// for _, v := range image {
	// 	fmt.Printf("%+v\n", v)
	// }
	os := image[0].Ones
	ts := image[0].Twos
	return os * ts
}

func Part2(data string, width, height int) int {
	image := make([][]int, 0)
	layerSize := width * height
	l := make([]int, 0)
	for i, px := range data {
		v, _ := strconv.Atoi(string(px))
		l = append(l, v)
		if (i+1)%layerSize == 0 {
			image = append(image, l)
			l = make([]int, 0)
		}
	}
	decodedImage := make([]int, width*height)
	for pxIndex := 0; pxIndex < width*height; pxIndex++ {
		for _, layer := range image {
			px := layer[pxIndex]
			if px < 2 {
				decodedImage[pxIndex] = px
				break
			}
		}
	}
	for i, v := range decodedImage {
		c := " "
		if v == 1 {
			c = "|"
		}
		fmt.Printf("%s", c)
		if (i+1)%width == 0 {
			fmt.Println()
		}
	}
	return 0
}
