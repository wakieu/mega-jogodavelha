package main

import (
	"errors"
	"fmt"
)

type Board struct {
	
	bigBoard [9]rune
	boards [9][9]rune
	currentQuadrant int
	turn rune

}

func NewBoard() Board {
	b := Board{}
	b.bigBoard = [9]rune{0,0,0,0,0,0,0,0,0}
	b.bigBoard[0] = 'X'
	b.bigBoard[1] = 'O'
	b.currentQuadrant = -1
	b.turn = 'X'

	for i := range b.boards {
		b.boards[i] = [9]rune{'_', '_', '_', '_', '_', '_', '_', '_', '_'}
	}

	return b
}

func convertCoords(x, y int) (int, int) {
	x1 := (((y-1) / 3) * 3) + (x-1) / 3
	y1 := ((y-1) % 3 * 3) + (x-1) % 3

	return x1, y1
}

func (b *Board) validate(x, y int) (error) {
	x, y = convertCoords(x, y)

	if b.boards[x][y] != '_' || b.bigBoard[x] != 0 {
		return errors.New("!! invalid coord")
	}
	if b.currentQuadrant != -1 && x != b.currentQuadrant {
		err := fmt.Sprintf("!! invalid quadrant (current: %d)", b.currentQuadrant + 1)
		return errors.New(err)
	}

	return nil
}

func (b *Board) play(x, y int) {
	x, y = convertCoords(x, y)

	b.boards[x][y] = b.turn

	b.checkQWin(x)

	if b.turn == 'X' { b.turn = 'O' } else { b.turn = 'X' }

	if b.bigBoard[y] == 0 {
		b.currentQuadrant = y
	} else {
		b.currentQuadrant = -1
	}
}

func (b *Board) checkQWin(q int) {
	if b.boards[q][0] != '_' && b.boards[q][0] == b.boards[q][1] && b.boards[q][1] == b.boards[q][2] ||
		b.boards[q][2] != '_' && b.boards[q][2] == b.boards[q][5] && b.boards[q][5] == b.boards[q][8] ||
		b.boards[q][8] != '_' && b.boards[q][8] == b.boards[q][7] && b.boards[q][7] == b.boards[q][6] ||
		b.boards[q][0] != '_' && b.boards[q][0] == b.boards[q][3] && b.boards[q][3] == b.boards[q][6] ||
		b.boards[q][4] != '_' && b.boards[q][0] == b.boards[q][4] && b.boards[q][4] == b.boards[q][8] ||
		b.boards[q][4] != '_' && b.boards[q][2] == b.boards[q][4] && b.boards[q][4] == b.boards[q][6] ||
		b.boards[q][4] != '_' && b.boards[q][3] == b.boards[q][4] && b.boards[q][4] == b.boards[q][5] ||
		b.boards[q][4] != '_' && b.boards[q][1] == b.boards[q][4] && b.boards[q][4] == b.boards[q][7] {
			b.bigBoard[q] = b.turn
	}
}

func (b *Board) checkWin() (bool, rune) {
	if b.bigBoard[0] != 0 && b.bigBoard[0] == b.bigBoard[1] && b.bigBoard[1] == b.bigBoard[2] ||
		b.bigBoard[2] != 0 && b.bigBoard[2] == b.bigBoard[5] && b.bigBoard[5] == b.bigBoard[8] ||
		b.bigBoard[8] != 0 && b.bigBoard[8] == b.bigBoard[7] && b.bigBoard[7] == b.bigBoard[6] ||
		b.bigBoard[0] != 0 && b.bigBoard[0] == b.bigBoard[3] && b.bigBoard[3] == b.bigBoard[6] ||
		b.bigBoard[4] != 0 && b.bigBoard[0] == b.bigBoard[4] && b.bigBoard[4] == b.bigBoard[8] ||
		b.bigBoard[4] != 0 && b.bigBoard[2] == b.bigBoard[4] && b.bigBoard[4] == b.bigBoard[6] ||
		b.bigBoard[4] != 0 && b.bigBoard[3] == b.bigBoard[4] && b.bigBoard[4] == b.bigBoard[5] ||
		b.bigBoard[4] != 0 && b.bigBoard[1] == b.bigBoard[4] && b.bigBoard[4] == b.bigBoard[7] {
			if b.turn == 'X' {
				return true,  'O'
			} else {
				return true,  'X'
			}
	}
	return false, 0
}

func getSection(q int) int {
	return 46 + ((q % 3) * 14) + ((q/3) * 352)
}

func (b *Board) overwriteQuadrants(s string) (string) {
	x1 := "             "
	x2 := "    \\\\ //    "
	x3 := "     \\//     "
	x4 := "     //\\     "
	x5 := "    // \\\\    "
	o1 := "             "
	o2 := "  //=====\\\\  "
	o3 := "  ||     ||  "
	o4 := "  ||     ||  "
	o5 := "  \\\\=====//  "

	for i, e := range b.bigBoard {
		if e == 0 { continue }
		p := getSection(i)
		if e == 'X' {
 			s = s[:p+44*0] + x1 + s[p+13+44*0:]
			s = s[:p+44*1] + x2 + s[p+13+44*1:]
			s = s[:p+44*2] + x3 + s[p+13+44*2:]
			s = s[:p+44*3] + x4 + s[p+13+44*3:]
			s = s[:p+44*4] + x5 + s[p+13+44*4:]
		} else {
			s = s[:p+44*0] + o1 + s[p+13+44*0:]
			s = s[:p+44*1] + o2 + s[p+13+44*1:]
			s = s[:p+44*2] + o3 + s[p+13+44*2:]
			s = s[:p+44*3] + o4 + s[p+13+44*3:]
			s = s[:p+44*4] + o5 + s[p+13+44*4:]
		}
	}

	return s
}

func (b *Board) render() {
	header    := "    A   B   C     D   E   F     G   H   I  "
	space     := "               |             |             " // len 43, with \n 44
	separator := "  -------------+-------------+-------------"
	fmt.Println()
	fmt.Println(header)

	// creates board string
	boardString := ""
	for i := 0; i < 9; i++ {
		boardString += fmt.Sprintln(space)

		boardString += fmt.Sprintf(
			"%d   %s   %s   %s  |  %s   %s   %s  |  %s   %s   %s  \n",
			i + 1,
			string(b.boards[0 + (i/3)*3][(i)%3*3 + 0]),
			string(b.boards[0 + (i/3)*3][(i)%3*3 + 1]),
			string(b.boards[0 + (i/3)*3][(i)%3*3 + 2]),
			string(b.boards[1 + (i/3)*3][(i)%3*3 + 0]),
			string(b.boards[1 + (i/3)*3][(i)%3*3 + 1]),
			string(b.boards[1 + (i/3)*3][(i)%3*3 + 2]),
			string(b.boards[2 + (i/3)*3][(i)%3*3 + 0]),
			string(b.boards[2 + (i/3)*3][(i)%3*3 + 1]),
			string(b.boards[2 + (i/3)*3][(i)%3*3 + 2]),
		)
		if i == 2 || i == 5 {
			boardString += fmt.Sprintln(space)
			boardString += fmt.Sprintln(separator)
		}
	}

	boardString = b.overwriteQuadrants(boardString)

	fmt.Print(boardString)
	fmt.Println(space)
	fmt.Println()
	fmt.Printf("Current quadrant: %v\n", b.currentQuadrant+1)
}
