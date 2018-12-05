package days

import (
	"fmt"
	"log"
	"strings"
)

// See https://adventofcode.com/2018/day/3 for description of problem

const (
	fabricSize = 5000
)

// Claim relate a claim id to a square
type Claim struct {
	id int
	Square
}

// Id retrieve claim id
func (c Claim) Id() int {
	return c.id
}

// String return friendly string of claim
func (c Claim) String() string {
	return fmt.Sprintf("Id: %d, Square: %s", c.id, c.Square.String())
}

// Square it is a square
type Square struct {
	left   int
	top    int
	width  int
	height int
}

// NewSquare create a new instance of a square
func NewSquare(top, left, width, height int) Square {
	return Square{top: top, left: left, width: width, height: height}
}

// String return a friendly string representation of the square
func (s Square) String() string {
	return fmt.Sprintf("left: %d, top: %d, width: %d, height: %d, right(c): %d, bottom(c): %d",
		s.left, s.top, s.width, s.height, s.Right(), s.Bottom())
}

// Left return the left most edge
func (s Square) Left() int {
	return s.left
}

// Top return the top most edge
func (s Square) Top() int {
	return s.top
}

// Width return the square width
func (s Square) Width() int {
	return s.width
}

// Height return the square height
func (s Square) Height() int {
	return s.height
}

// Right calculate the right most edge of the square
func (s Square) Right() int {
	return s.left + s.width
}

// Bottom calculate the bottom most edge of the square
func (s Square) Bottom() int {
	return s.top + s.height
}

// Area return the area of the square
func (s Square) Area() int {
	return s.height * s.width
}

// Empty check if square is all zeros (empty)
func (s Square) Empty() bool {
	return s.top == 0 &&
		s.left == 0 &&
		s.width == 0 &&
		s.height == 0
}

// Equals check square equality
func (s Square) Equals(sqr Square) bool {
	return s.top == sqr.top &&
		s.left == sqr.left &&
		s.width == sqr.width &&
		s.height == sqr.height
}

// Overlap compare two squares, return a new square of overlap between the two compared.
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

// NewClaim generate a new claim from the given input line,
// example: #1 @ 2,3: 4x5 -> ClaimId: 1, Left: 2, Top: 3, Width: 4, Height: 5
func NewClaim(s string) Claim {
	reader := strings.NewReader(s)
	var id, left, top, width, height int
	if _, err := fmt.Fscanf(reader, "#%d @ %d,%d: %dx%d",
		&id, &left, &top, &width, &height); err != nil {
		log.Fatal(err)
	}
	return Claim{id: id, Square: NewSquare(top, left, width, height)}
}

// NoOverlap take a list of input lines and return the claim id of any claims that have no overlaps
func NoOverlap(input []string) []int {
	allclaims := make([]Claim, 0)
	for _, s := range input {
		allclaims = append(allclaims, NewClaim(s))
	}
	claimIds := make([]int, 0, len(allclaims))
	for i := 0; i < len(allclaims); i++ {
		overlap := false
		sq1 := allclaims[i].Square
		for j := 0; j < len(allclaims); j++ {
			if j == i {
				// skip comparing self because it will always overlap.
				continue
			}
			sq2 := allclaims[j].Square
			if sq1.Overlap(sq2).Area() > 0 {
				overlap = true
			}
		}
		if !overlap {
			claimIds = append(claimIds, allclaims[i].Id())
		}
	}
	return claimIds
}

// OverlappingClaims take a list of input lines, and return the total area overlapped by 2 or more claims
func OverlappingClaims(input []string) int {
	allclaims := make([]Claim, 0)
	for _, s := range input {
		allclaims = append(allclaims, NewClaim(s))
	}
	var grid [fabricSize][fabricSize]int
	for i := 0; i < len(allclaims); i++ {
		sq := allclaims[i].Square
		for l := sq.Left(); l < sq.Right(); l++ {
			for t := sq.Top(); t < sq.Bottom(); t++ {
				grid[l][t] += 1
			}
		}
	}
	total := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] > 1 {
				total += 1
			}
		}
	}

	return total
}
