package days

import (
	"fmt"
	"log"
	"strings"
)

// See https://adventofcode.com/2018/day/3 for description of problem

type Claim struct {
	id int
	Square
}

func (c Claim) Id() int {
	return c.id
}

type Square struct {
	left   int
	top    int
	width  int
	height int
}

func NewSquare(top, left, height, width int) Square {
	return Square{top: top, left: left, height: height, width: width}
}

func (s Square) Left() int {
	return s.left
}

func (s Square) Top() int {
	return s.top
}

func (s Square) Width() int {
	return s.width
}

func (s Square) Height() int {
	return s.height
}

func (s Square) Right() int {
	return s.left + s.width
}

func (s Square) Bottom() int {
	return s.top + s.height
}

func (s Square) LargestCorner() (int, int) {
	return s.left + s.width, s.top + s.height
}

func (s Square) Area() int {
	return s.height * s.width
}

func (s Square) Empty() bool {
	return s.top == 0 &&
		s.left == 0 &&
		s.width == 0 &&
		s.height == 0
}

func (s Square) Equals(sqr Square) bool {
	return s.top == sqr.top &&
		s.left == sqr.left &&
		s.width == sqr.width &&
		s.height == sqr.height
}

func (s Square) Overlap(sqr Square) Square {
	left, top, width, height := 0, 0, 0, 0

	left = Max(s.Left(), sqr.Left())
	right := Min(s.Right(), sqr.Right())
	top = Max(s.Top(), sqr.Top())
	bottom := Min(s.Bottom(), sqr.Bottom())
	if left < right && top < bottom {
		width = right - left
		height = bottom - top
	} else {
		top = 0
		left = 0
	}
	return Square{top: top, left: left, width: width, height: height}
}

func NewClaim(s string) Claim {
	// #9 @ 109,286: 11x16
	reader := strings.NewReader(s)
	var id, left, top, width, height int
	if _, err := fmt.Fscanf(reader, "#%d @ %d,%d: %dx%d",
		&id, &left, &top, &width, &height); err != nil {
		log.Fatal(err)
	}
	return Claim{id: id, Square: NewSquare(top, left, width, height)}
}

func OverlappingClaims(input []string) int {
	allclaims := make([]Claim, 0, len(input))
	for _, s := range input {
		allclaims = append(allclaims, NewClaim(s))
	}
	totalArea := 0
	// overlaps := make([]Square, 0, len(allclaims)/2)
	for i := 0; i < len(allclaims); i++ {
		for j := i + 1; j < len(allclaims); j++ {
			overlap := allclaims[i].Overlap(allclaims[j].Square)
			totalArea += allclaims[i].Area() + allclaims[j].Area() - overlap.Area()

			// overlaps = append(overlaps, overlap)
		}
	}
	// for i := 0; i < len(overlaps); i++ {
	// 	for j := i; j < len(overlaps); j++ {
	// 		totalArea -= overlaps[i].Overlap(overlaps[j]).Area()
	// 	}
	// }
	return totalArea
}
