package main

import (
	"fmt"
	"strconv"
)

//基盤構造
//tokens　コマ
//currentTurn　現在のターン
type Board struct {
	tokens      []int
	CurrentTurn int
}

//ターン交換
//player1(o)のターンならplayer2(x)に交換する
//player2(x)のターンならplayer1(o)に交換する
func (b *Board) SwitchTurn() {
	if b.CurrentTurn == 1 {
		b.CurrentTurn = 2
	} else if b.CurrentTurn == 2 {
		b.CurrentTurn = 1
	}
}

// 勝敗判定
// player1(o)が勝つなら"o"を返す
// player2(x)が勝つなら"x"を返す
// 引き分けなら"d"を返す
func (b *Board) WinCheck() string {
	count_1p := 0
	count_2p := 0
	for x := 0; x < 8; x++ {
		for y := 0; y < 8; y++ {
			tokens := b.Get(x, y)
			// player1の駒
			if tokens == "o" {
				count_1p++
			} else if tokens == "x" {
				count_2p++
			}
		}
	}
	// 勝敗判定
	// TODO: 定数を利用
	if count_1p > count_2p {
		return "o"
	}
	if count_1p < count_2p {
		return "x"
	}
	if count_1p == count_2p {
		return "d"
	}
	return "." //予期以外　error
}

//基盤初期化
//中央のコマの初期化
//ターンの初期化
func (b *Board) BoardInitialization() { //碁盤を [........] 初期化
	b.tokens = make([]int, 64)
	for i := 0; i < 64; i++ {
		b.tokens[i] = 0
	}

	b.tokens[3+8*3] = 1 //中央コ初期化
	b.tokens[4+8*4] = 1
	b.tokens[3+8*4] = 2
	b.tokens[4+8*3] = 2

	b.CurrentTurn = 1
}

//基盤展示　b.tokens[]を8*8の基盤で返す
func (b *Board) BoardShow() string {
	var result string
	for y := 0; y < 8; y++ {
		for x := 0; x < 8; x++ {
			result += b.Get(x, y)
			result += " "
		}
		result += "\n"
	}
	return result
}

//コマが置く位置は基盤の中であるかどうかを判断
//基盤の外で置かれない
func (b *Board) notPutValid(x, y int) int { //error function
	if x > 7 || x < 0 || y > 7 || y < 0 { //碁盤の中でわない error
		//fmt.Printf("碁盤の中でわないので、もう一回入力してください\n")
		return 1
	} else if b.tokens[x+8*y] != 0 {
		//fmt.Printf("駒はもう存在していますので、もう一回入力してください\n") //重複　error
		return 2
	}
	return 0
}

// func (b *Board) Put(x, y int, u string) int { //int type 座標x, yを入力、int type tokens[]に保存
// 	if u == "o" { //player1
// 		b.tokens[x+8*y] = 1
// 		return 0 //順調
// 	} else if u == "x" { //player2
// 		b.tokens[x+8*y] = 2
// 		return 0 //順調
// 	}
// 	return 2 //予期以外　error
// }

//コマ投下結果を得る
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

//コマ逆転
//相手のコマが挟まれると、自分のコマにひっくり返される
//少なくとも相手のコマを1つ挟まなければコマを置くことはできない
func (b *Board) Reverse(x, y int) int {
	if b.notPutValid(x, y) != 0 { //putできないならreturn
		return 1 //error number
	}

	dx := [8]int{-1, -1, -1, 0, 0, 1, 1, 1} //左上，左，左下，上，下，右上，右，右下
	dy := [8]int{-1, 0, 1, -1, 1, -1, 0, 1}

	//var count int = 0

	flipcount := 0
	for i := 0; i < 8; i++ { //8方向, reverseできるのか、チェック
		nx, ny := x+dx[i], y+dy[i] //一回歩く
		count := 0

		for b.notPutValid(nx, ny) == 2 && b.tokens[nx+8*ny] != b.CurrentTurn { //この方向に駒があり、でもcurrentPlayerの駒でわないの場合
			nx += dx[i]
			ny += dy[i]
			count++ //どこまでreverse必要をチェック

			//fmt.Printf("%d\n", &count)
		}

		if count > 0 && b.notPutValid(nx, ny) == 2 { //現在の方向に、すべで駒をreverse
			for j := 0; j <= count; j++ { //上の”どこまで”
				nx -= dx[i]
				ny -= dy[i]
				b.tokens[nx+8*ny] = b.CurrentTurn //reverse
				flipcount++

				//fmt.Printf("%d\n", &flipcount)
			}
		}
	}

	if flipcount > 0 { //reverseできるなら、入力の座標に駒をおく
		b.tokens[x+8*y] = b.CurrentTurn
		b.SwitchTurn()
		return 0 //順調
	} else if flipcount == 0 { //reversできない、ターン続き
		return 2 //error number
	}

	return 0 //順調
}

// ボードの周りに座標がついた基盤展示　b.tokens[]を10*10の基盤で返す
func (b *Board) BoardShow_with_coordinate() string {
	var result string
	for y := 0; y < 10; y++ {
		// 上下の座標
		if y == 0 || y == 9 {
			for x := 0; x < 10; x++ {
				if x == 0 || x == 9 {
					result += "  "
				} else {
					result += strconv.Itoa(x)
					result += " "
				}
			}
			result += "\n"
			continue
		}
		for x := 0; x < 10; x++ {
			if x == 0 || x == 9 {
				result += strconv.Itoa(y)
				if x == 0 {
					result += " "
				}
			} else {
				result += b.Get(x-1, y-1)
				result += " "
			}
		}
		result += "\n"
	}
	return result
}

// turnのプレイヤーがボードに駒を置けるか判定
func (b *Board) IsPuttable(turn int) bool {
	for x := 0; x < 8; x++ {
		for y := 0; y < 8; y++ {
			// プレイヤー1が(x,y)に駒を配置できる
			if b.IsPlaceable_xy(x, y, turn) {
				return false
			}
		}
	}
	return true
}

// turnのプレイヤーが(x,y)に駒を置けるか判定
func (b *Board) IsPlaceable_xy(x, y, turn int) bool {
	// (x,y)に駒を配置できないならfalseを返す
	if b.notPutValid(x, y) != 0 {
		return false
	}

	// 方向オフセット: 左上，左，左下，上，下，右上，右，右下
	dx := [8]int{-1, -1, -1, 0, 0, 1, 1, 1}
	dy := [8]int{-1, 0, 1, -1, 1, -1, 0, 1}

	for i := 0; i < 8; i++ { //8方向, reverseできるのか、チェック
		nx, ny := x+dx[i], y+dy[i] // オフセットの方向に駒を1マス進める
		count := 0                 // リバースできる駒の数

		// (nx,ny)にturnプレイヤーでないに駒が存在する限り探索を続ける
		for b.notPutValid(nx, ny) == 2 && b.tokens[nx+8*ny] != turn {
			nx += dx[i]
			ny += dy[i]
			count++ //どこまでreverse必要をチェック
		}
		// 有効マスに駒を配置できる
		if count > 0 && b.notPutValid(nx, ny) == 2 {
			return true
		}
	}
	return false
}

// テスト用ボード作成
func (b *Board) make_board() {
	b.tokens = []int{
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 1, 0, 0, 0, 0, 0, 0,
		0, 0, 1, 0, 0, 0, 0, 0,
		0, 0, 0, 1, 0, 0, 0, 0,
		0, 0, 0, 0, 1, 0, 0, 0,
		0, 0, 0, 0, 0, 1, 0, 0,
		0, 0, 0, 0, 0, 0, 1, 0,
		0, 0, 0, 0, 0, 0, 0, 2,
	}
}

func (b *Board) CountTokens() (int, int) {
	count_1p := 0
	count_2p := 0
	for x := 0; x < 8; x++ {
		for y := 0; y < 8; y++ {
			tokens := b.Get(x, y)
			// player1の駒
			if tokens == "o" {
				count_1p++
			} else if tokens == "x" {
				count_2p++
			}
		}
	}
	return count_1p, count_2p
}

func (b *Board) ShowPlaceablePositions(turn int) {
	fmt.Println("駒を置ける位置：")
	for y := 0; y < 8; y++ {
		for x := 0; x < 8; x++ {
			if b.IsPlaceable_xy(x, y, turn) {
				fmt.Printf("(%d, %d) ", x+1, y+1)
			}
		}
	}
	fmt.Println()
}

func main() {
	var x, y, i int
	var ReverseError int
	board := Board{}
	board.BoardInitialization()                  //初期化
	fmt.Printf("please input like: x[space]y\n") //gameの説明。拡張したいならどうぞ
	// board.make_board()                           // テスト用

	board.CurrentTurn = 1
	for i = 0; i < 150; i++ {
		// 誰もコマを置けないならループを抜ける
		if board.IsPuttable(1) && board.IsPuttable(2) {
			break
		}

		fmt.Printf("%s", board.BoardShow_with_coordinate())

		// 現在のプレイヤーが駒を配置可能か判��する
		if board.IsPuttable(board.CurrentTurn) {
			fmt.Printf("Player%dは駒を配置できません\n", board.CurrentTurn)
			board.SwitchTurn()
		}

		if ReverseError == 2 {
			fmt.Printf("そこには置けません。もう一度入力してください。\n")
			board.ShowPlaceablePositions(board.CurrentTurn)
			ReverseError = 0
		}

		if board.CurrentTurn == 1 {
			count_1p, count_2p := board.CountTokens()
			fmt.Printf("Player1: Input (x,y) o:%d x:%d\n", count_1p, count_2p) //player1のターン。説明を拡張したいならどうぞ
		} else if board.CurrentTurn == 2 {
			count_1p, count_2p := board.CountTokens()
			fmt.Printf("Player2: Input (x,y) o:%d x:%d\n", count_1p, count_2p) //player2のターン。説明を拡張したいならどうぞ
		}
		// board.ShowPlaceablePositions(board.CurrentTurn)
		//fmt.Printf("%s", board.BoardShow())
		fmt.Scan(&x, &y)

		for board.notPutValid(x-1, y-1) != 0 { //errorの場合、もう一度入力
			println(x, y)
			if board.notPutValid(x-1, y-1) == 1 {
				fmt.Printf("碁盤の外です。もう一度入力してください。\n")
				// fmt.Printf(board.BoardShow())
			} else if board.notPutValid(x-1, y-1) == 2 {
				fmt.Printf("すでに駒が置かれています。もう一度入力してください。\n")
				// fmt.Printf(board.BoardShow())
			}
			board.ShowPlaceablePositions(board.CurrentTurn)
			fmt.Scan(&x, &y)
		}

		ReverseError = board.Reverse(x-1, y-1)
	}

	if board.WinCheck() == "o" { //player1 win
		fmt.Printf("player1 win\n")
	} else if board.WinCheck() == "x" { //player2 win
		fmt.Printf("player2 win\n")
	} else if board.WinCheck() == "d" { //draw
		fmt.Printf("draw\n")
	} else if i == 150 {
		fmt.Printf("rest round is 0, game over\n")
	}
	fmt.Printf("%s", board.BoardShow_with_coordinate())
}
