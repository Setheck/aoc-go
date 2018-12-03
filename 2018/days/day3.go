package days

import (
	"fmt"
	"strings"
)

// See https://adventofcode.com/2018/day/3 for description of problem

type Claim struct {
	id int
	*Square
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

func (s Square) Overlap(sqr Square) *Square {
	leftOverlap := s.Left() < sqr.Left() && s.Right() > sqr.Left()
	topOverlap := s.Top() < sqr.Top() && s.Bottom() > sqr.Top()
	rightOverlap := s.Right() > sqr.Right() && s.Left() < sqr.Right()
	botOverlap := s.Bottom() > sqr.Bottom() && s.Top() < sqr.Bottom()

	left, top, width, height := 0, 0, 0, 0
	if leftOverlap {
		left = sqr.Left()
	}
	if topOverlap {
		top = sqr.Top()
	}
	if rightOverlap {
		// right = s.Right()
	}
	if botOverlap {
		// bottom = sqr.Bottom()
	}
	return &Square{top: top, left: left, width: width, height: height}
}

func NewClaim(s string) *Claim {
	// #9 @ 109,286: 11x16
	reader := strings.NewReader(s)
	claim := &Claim{}
	if _, err := fmt.Fscanf(reader, "#%d @ %d,%d: %dx%d",
		&claim.id, &claim.left, &claim.top, &claim.width, &claim.height); err != nil {
		panic(err)
	}
	return claim
}

func OverlappingClaims(input []string) {
	allclaims := make([]Claim, 0, len(input))
	for _, s := range input {
		allclaims = append(allclaims, NewClaim(s))
	}
}
