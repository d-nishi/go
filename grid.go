package main

import (
  "fmt"
)

type Grid struct {
   height, width int
   grid []byte
}

func main() {

   w:=3
   h:=3

   g := NewGrid(w,h)
   g.Draw()
}

func NewGrid(x, y int) Grid {

   wth := 2*x + 2
   hgt := 2*y + 1

   g := make([]byte, wth*hgt)

   for i := 0; i < hgt; i += 2 {
	row0 := i * wth
	row1 := (i + 1) * wth
	for j := 0; j < wth-2; j += 2 {
		g[row0+j], g[row0+j+1] = '+', '-'
		if row1+j+1 <= wth*hgt {
			g[row1+j], g[row1+j+1] = '|', ' '
		}
	}
	g[row0+wth-2], g[row0+wth-1] = '+', '\n'
	if row1+wth < wth*hgt {
		g[row1+wth-2], g[row1+wth-1] = '|', '\n'
	}

   }

   return Grid {
	   height: y,
	   width:  x,
	   grid:   g,
   }

}

func (g Grid) Draw() {
	fmt.Print("\x0c", g, "\n")        // Print frame
}

func (g Grid) String() string{
	return string(g.grid)
}
