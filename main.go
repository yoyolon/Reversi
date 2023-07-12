package main

import "fmt"

type Board struct {
	tokens []int
}

func (b *Board) BoardIinitializtion() { //碁盤を [........] 初期化
	b.tokens = make([]int, 64)
	for i := 0; i < 64; i++ {
		b.tokens[i] = 0
	}
}

func (b *Board) BoardShow() string {
	var result string
	for x := 0; x < 8; x++ {
		for y := 0; y < 8; y++ {
			result += b.Get(x, y)
			result += " "
		}
		result += "\n"
	}
	return result
}

func (b *Board) Put(x, y int, u string) int { //int type 座標x, yを入力、int type tokens[]に保存
	if x > 7 || x < 0 || y > 7 || y < 0 { //碁盤の中でわない error
		fmt.Printf("碁盤の中でわないので、もう一回入力してください\n")
		return 1
	}

	if b.tokens[x+8*y] != 0 {
		fmt.Printf("駒はもう存在していますので、もう一回入力してください\n") //重複　error
		return 1
	}

	if u == "o" { //player1
		b.tokens[x+8*y] = 1
		return 0 //順調
	} else if u == "x" { //player2
		b.tokens[x+8*y] = 2
		return 0 //順調
	}
	return 2 //予期以外　error
}

func (b *Board) Get(x, y int) string { //int type tokens[] --> 駒文字
	if b.tokens[x+8*y] == 1 { //player1
		return "o"
	} else if b.tokens[x+8*y] == 2 { //palyer2
		return "x"
	} else if b.tokens[x+8*y] == 0 { //駒は存在しない
		return "."
	}
	return "error" //予期以外
}

func main() {

}
