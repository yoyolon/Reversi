package main

import (
	"testing"
)

// 盤面からプレイヤー1("o")が(3,5)に駒を置くテスト
func TestPutToken_1P35_InitialBoard(t *testing.T) {
	// 盤面
	b := &Board{
		tokens: []int{
			0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 1, 2, 0, 0, 0,
			0, 0, 0, 2, 1, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
		},
	}
	// プレイヤー1が(3,5)に駒を置く
	b.CurrentTurn = 1      // プレイヤー1のターン
	b.Reverse(3, 5)        // (3,5)に駒を置く
	println(b.BoardShow()) // 盤面の表示(テストのテスト)
	// (3,5)が"o"になる
	if b.Get(3, 5) != "o" {
		t.Errorf("Failed to put \"o\" on (3,5)")
	}
	// (3,4)も"o"になる
	if b.Get(3, 4) != "o" {
		t.Errorf("Failed to reverse \"x\" on (3,4)")
	}
}

// 盤面からプレイヤー1("o")が(5,3)に駒を置くテスト
func TestPutToken_1P53_InitialBoard(t *testing.T) {
	// 盤面
	b := &Board{
		tokens: []int{
			0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 1, 2, 0, 0, 0,
			0, 0, 0, 2, 1, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
		},
	}
	// プレイヤー1が(5,3)に駒を置く
	b.CurrentTurn = 1      // プレイヤー1のターン
	b.Reverse(5, 3)        // (5,3)に駒を置く
	println(b.BoardShow()) // 盤面の表示(テストのテスト)
	// (5,3)が"o"になる
	if b.Get(5, 3) != "o" {
		t.Errorf("Failed to put \"o\" on (5,3)")
	}
	// (4,3)も"o"になる
	if b.Get(4, 3) != "o" {
		t.Errorf("Failed to reverse \"x\" on (4,3)")
	}
}

// テスト追加

// 盤面からプレイヤー1("o")が(4,2)に駒を置くテスト
func TestPutToken_1P42_InitialBoard(t *testing.T) {
	// 盤面
	b := &Board{
		tokens: []int{
			0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 1, 2, 0, 0, 0,
			0, 0, 0, 2, 1, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
		},
	}
	// プレイヤー1が(4,2)に駒を置く
	b.CurrentTurn = 1      // プレイヤー1のターン
	b.Reverse(4, 2)        // (4,2)に駒を置く
	println(b.BoardShow()) // 盤面の表示(テストのテスト)
	// (4,2)が"o"になる
	if b.Get(4, 2) != "o" {
		t.Errorf("Failed to put \"o\" on (4,2)")
	}
	// (4,3)も"o"になる
	if b.Get(4, 3) != "o" {
		t.Errorf("Failed to reverse \"x\" on (4,3)")
	}
}

// 盤面からプレイヤー1("o")が(2,4)に駒を置くテスト
func TestPutToken_1P24_InitialBoard(t *testing.T) {
	// 盤面
	b := &Board{
		tokens: []int{
			0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 1, 2, 0, 0, 0,
			0, 0, 0, 2, 1, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
		},
	}
	// プレイヤー1が(2,4)に駒を置く
	b.CurrentTurn = 1      // プレイヤー1のターン
	b.Reverse(2, 4)        // (2,4)に駒を置く
	println(b.BoardShow()) // 盤面の表示(テストのテスト)
	// (2,4)が"o"になる
	if b.Get(2, 4) != "o" {
		t.Errorf("Failed to put \"o\" on (2,4)")
	}
	// (3,4)も"o"になる
	if b.Get(3, 4) != "o" {
		t.Errorf("Failed to reverse \"x\" on (3,4)")
	}
}

// 盤面からプレイヤー2("x")が(2,3)に駒を置くテスト
func TestPutToken_2P23_InitialBoard(t *testing.T) {
	// 盤面
	b := &Board{
		tokens: []int{
			0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 1, 2, 0, 0, 0,
			0, 0, 0, 2, 1, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
		},
	}
	// プレイヤー2が(2,3)に駒を置く
	b.CurrentTurn = 2      // プレイヤー2のターン
	b.Reverse(2, 3)        // (2,3)に駒を置く
	println(b.BoardShow()) // 盤面の表示(テストのテスト)
	// (2,3)が"x"になる
	if b.Get(2, 3) != "x" {
		t.Errorf("Failed to put \"x\" on (2,3)")
	}
	// (3,3)も"x"になる
	if b.Get(3, 3) != "x" {
		t.Errorf("Failed to reverse \"o\" on (3,3)")
	}
}

// 盤面からプレイヤー2("x")が(3,2)に駒を置くテスト
func TestPutToken_2P32_InitialBoard(t *testing.T) {
	// 盤面
	b := &Board{
		tokens: []int{
			0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 1, 2, 0, 0, 0,
			0, 0, 0, 2, 1, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
		},
	}
	// プレイヤー2が(3,2)に駒を置く
	b.CurrentTurn = 2      // プレイヤー2のターン
	b.Reverse(3, 2)        // (3,2)に駒を置く
	println(b.BoardShow()) // 盤面の表示(テストのテスト)
	// (3,2)が"x"になる
	if b.Get(3, 2) != "x" {
		t.Errorf("Failed to put \"x\" on (3,2)")
	}
	// (3,3)も"x"になる
	if b.Get(3, 3) != "x" {
		t.Errorf("Failed to reverse \"o\" on (3,3)")
	}
}

// 盤面からプレイヤー2("x")が(4,5)に駒を置くテスト
func TestPutToken_2P45_InitialBoard(t *testing.T) {
	// 盤面
	b := &Board{
		tokens: []int{
			0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 1, 2, 0, 0, 0,
			0, 0, 0, 2, 1, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
		},
	}
	// プレイヤー2が(4,5)に駒を置く
	b.CurrentTurn = 2      // プレイヤー2のターン
	b.Reverse(4, 5)        // (4,5)に駒を置く
	println(b.BoardShow()) // 盤面の表示(テストのテスト)
	// (4,5)が"x"になる
	if b.Get(4, 5) != "x" {
		t.Errorf("Failed to put \"x\" on (4,5)")
	}
	// (4,3)も"x"になる
	if b.Get(4, 3) != "x" {
		t.Errorf("Failed to reverse \"o\" on (4,3)")
	}
}

// 盤面からプレイヤー2("x")が(5,4)に駒を置くテスト
func TestPutToken_2P54_InitialBoard(t *testing.T) {
	// 盤面
	b := &Board{
		tokens: []int{
			0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 1, 2, 0, 0, 0,
			0, 0, 0, 2, 1, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
		},
	}
	// プレイヤー2が(5,4)に駒を置く
	b.CurrentTurn = 2      // プレイヤー2のターン
	b.Reverse(5, 4)        // (5,4)に駒を置く
	println(b.BoardShow()) // 盤面の表示(テストのテスト)
	// (5,4)が"x"になる
	if b.Get(5, 4) != "x" {
		t.Errorf("Failed to put \"x\" on (5,4)")
	}
	// (4,3)も"x"になる
	if b.Get(4, 3) != "x" {
		t.Errorf("Failed to reverse \"o\" on (4,3)")
	}
}

// oxの左にxを置く
func TestPutToken_2P_Simple(t *testing.T) {
	// 盤面
	b := &Board{
		tokens: []int{
			0, 1, 2, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
		},
	}
	// プレイヤー2が(0,0)に駒を置く
	b.CurrentTurn = 2      // プレイヤー2のターン
	b.Reverse(0, 0)        // (0,0)に駒を置く
	println(b.BoardShow()) // 盤面の表示(テストのテスト)
	// (0,0)が"x"になる
	if b.Get(0, 0) != "x" {
		t.Errorf("Failed to put \"x\" on (0,0)")
	}
	// (1,0)も"x"になる
	if b.Get(1, 0) != "x" {
		t.Errorf("Failed to reverse \"o\" on (1,0)")
	}
}

// xoの左にoを置く
func TestPutToken_1P_Simple(t *testing.T) {
	// 盤面
	b := &Board{
		tokens: []int{
			0, 2, 1, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
		},
	}
	// プレイヤー1が(0,0)に駒を置く
	b.CurrentTurn = 1      // プレイヤー1のターン
	b.Reverse(0, 0)        // (0,0)に駒を置く
	println(b.BoardShow()) // 盤面の表示(テストのテスト)
	// (0,0)が"o"になる
	if b.Get(0, 0) != "o" {
		t.Errorf("Failed to put \"o\" on (0,0)")
	}
	// (1,0)も"o"になる
	if b.Get(1, 0) != "o" {
		t.Errorf("Failed to reverse \"x\" on (1,0)")
	}
}

// プレイヤー1("o")が斜めにひっくり返すテスト
func TestPutToken_1P_Diagonal(t *testing.T) {
	// 盤面
	b := &Board{
		tokens: []int{
			0, 0, 0, 0, 0, 0, 0, 0,
			0, 2, 0, 0, 0, 0, 0, 0,
			0, 0, 2, 0, 0, 0, 0, 0,
			0, 0, 0, 2, 0, 0, 0, 0,
			0, 0, 0, 0, 2, 0, 0, 0,
			0, 0, 0, 0, 0, 2, 0, 0,
			0, 0, 0, 0, 0, 0, 2, 0,
			0, 0, 0, 0, 0, 0, 0, 1,
		},
	}
	println("Before")
	println(b.BoardShow()) // 置く前の盤面
	// プレイヤー1が(0,0)に駒を置く
	b.CurrentTurn = 1 // プレイヤー1のターン
	b.Reverse(0, 0)   // (0,0)に駒を置く
	println("After")
	println(b.BoardShow()) // 盤面の表示(テストのテスト)
	// (0,0)が"o"になる
	if b.Get(0, 0) != "o" {
		t.Errorf("Failed to put \"o\" on (0,0)")
	}
	// (1,1)から(6,6)が"o"になる
	for i := 1; i <= 6; i++ {
		if b.Get(i, i) != "o" {
			t.Errorf("Failed to reverse \"x\" on (%d, %d)", i, i)
		}
	}
}

// プレイヤー2("X")が斜めにひっくり返すテスト
func TestPutToken_2P_Diagonal(t *testing.T) {
	// 盤面
	b := &Board{
		tokens: []int{
			0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 1, 0,
			0, 0, 0, 0, 0, 1, 0, 0,
			0, 0, 0, 0, 1, 0, 0, 0,
			0, 0, 0, 1, 0, 0, 0, 0,
			0, 0, 1, 0, 0, 0, 0, 0,
			0, 1, 0, 0, 0, 0, 0, 0,
			2, 0, 0, 0, 0, 0, 0, 0,
		},
	}
	println("Before")
	println(b.BoardShow()) // 置く前の盤面
	// プレイヤー2が(7,0)に駒を置く
	b.CurrentTurn = 2 // プレイヤー1のターン
	b.Reverse(7, 0)   // (2,0)に駒を置く
	println("After")
	println(b.BoardShow()) // 盤面の表示(テストのテスト)
	// (7,0)が"o"になる
	if b.Get(7, 0) != "x" {
		t.Errorf("Failed to put \"x\" on (0,0)")
	}
	// (6,1)から(1,6)が"x"になる
	for i := 1; i <= 6; i++ {
		if b.Get(7-i, i) != "x" {
			t.Errorf("Failed to reverse \"o\" on (%d, %d)", 7-i, i)
		}
	}
}

// 横一列をひっくり返すテスト
func TestPutToken_1P_Horizontal(t *testing.T) {
	// 盤面
	b := &Board{
		tokens: []int{
			0, 2, 2, 2, 2, 2, 2, 1,
			0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
		},
	}
	println("Before")
	println(b.BoardShow()) // 置く前の盤面
	// プレイヤー1が(0,0)に駒を置く
	b.CurrentTurn = 1 // プレイヤー1のターン
	b.Reverse(0, 0)   // (0,0)に駒を置く
	println("After")
	println(b.BoardShow()) // 盤面の表示(テストのテスト)
	// (0,0)が"o"になる
	if b.Get(0, 0) != "o" {
		t.Errorf("Failed to put \"o\" on (0,0)")
	}
	// (1,0)から(6,0)が"o"になる
	for i := 1; i <= 6; i++ {
		if b.Get(i, 0) != "o" {
			t.Errorf("Failed to reverse \"x\" on (%d, 0)", i)
		}
	}
}

// 縦一列をひっくり返すテスト
func TestPutToken_2P_Vertical(t *testing.T) {
	// 盤面
	b := &Board{
		tokens: []int{
			0, 0, 0, 0, 0, 0, 0, 2,
			0, 0, 0, 0, 0, 0, 0, 1,
			0, 0, 0, 0, 0, 0, 0, 1,
			0, 0, 0, 0, 0, 0, 0, 1,
			0, 0, 0, 0, 0, 0, 0, 1,
			0, 0, 0, 0, 0, 0, 0, 1,
			0, 0, 0, 0, 0, 0, 0, 1,
			0, 0, 0, 0, 0, 0, 0, 0,
		},
	}
	println("Before")
	println(b.BoardShow()) // 置く前の盤面
	// プレイヤー2が(7,7)に駒を置く
	b.CurrentTurn = 2 // プレイヤー2のターン
	b.Reverse(7, 7)   // (7,7)に駒を置く
	println("After")
	println(b.BoardShow()) // 盤面の表示(テストのテスト)
	// (7,7)が"x"になる
	if b.Get(7, 7) != "x" {
		t.Errorf("Failed to put \"x\" on (7,7)")
	}
	// (7,6)から(7,1)が"o"になる
	for i := 1; i <= 6; i++ {
		if b.Get(7, i) != "x" {
			t.Errorf("Failed to reverse \"o\" on (7, %d)", i)
		}
	}
}

// 縦一列と横一列をひっくり返すテスト
func TestPutToken_1P_VerticalHorizontal(t *testing.T) {
	// 盤面
	b := &Board{
		tokens: []int{
			0, 2, 2, 2, 2, 2, 2, 1,
			2, 0, 0, 0, 0, 0, 0, 0,
			2, 0, 0, 0, 0, 0, 0, 0,
			2, 0, 0, 0, 0, 0, 0, 0,
			2, 0, 0, 0, 0, 0, 0, 0,
			2, 0, 0, 0, 0, 0, 0, 0,
			2, 0, 0, 0, 0, 0, 0, 0,
			1, 0, 0, 0, 0, 0, 0, 0,
		},
	}
	println("Before")
	println(b.BoardShow()) // 置く前の盤面
	// プレイヤー1が(0,0)に駒を置く
	b.CurrentTurn = 1 // プレイヤー1のターン
	b.Reverse(0, 0)   // (0,0)に駒を置く
	println("After")
	println(b.BoardShow()) // 盤面の表示(テストのテスト)
	// (0,0)が"o"になる
	if b.Get(0, 0) != "o" {
		t.Errorf("Failed to put \"o\" on (0,0)")
	}
	// (1,0)から(6,0)が"o"になる
	for i := 1; i <= 6; i++ {
		if b.Get(i, 0) != "o" {
			t.Errorf("Failed to reverse \"x\" on (%d, 0)", i)
		}
	}
	// (0,1)から(0,6)が"o"になる
	for i := 1; i <= 6; i++ {
		if b.Get(0, i) != "o" {
			t.Errorf("Failed to reverse \"x\" on (0, %d)", i)
		}
	}
}

// プレイヤー2("x")が一度に複数の駒をひっくり返す
func TestPutToken_2P_Complex(t *testing.T) {
	// 盤面
	b := &Board{
		tokens: []int{
			0, 2, 0, 2, 0, 2, 0, 0,
			0, 0, 1, 1, 1, 0, 0, 0,
			0, 2, 1, 0, 1, 2, 0, 0,
			0, 0, 1, 1, 1, 0, 0, 0,
			0, 2, 0, 2, 0, 2, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
		},
	}
	println("Before")
	println(b.BoardShow()) // 置く前の盤面
	// プレイヤー1が(3,2)に駒を置く
	b.CurrentTurn = 2 // プレイヤー1のターン
	b.Reverse(3, 2)   // (3,2)に駒を置く
	println("After")
	println(b.BoardShow()) // 盤面の表示(テストのテスト)
	// (3,2)が"x"になる
	if b.Get(3, 2) != "x" {
		t.Errorf("Failed to put \"x\" on (3,2)")
	}
	// (3,2)の上一列
	if b.Get(2, 1) != "x" {
		t.Errorf("Failed to reverse \"o\" on (2,1)")
	}
	if b.Get(3, 1) != "x" {
		t.Errorf("Failed to reverse \"o\" on (3,1)")
	}
	if b.Get(4, 1) != "x" {
		t.Errorf("Failed to reverse \"o\" on (4,1)")
	}
	// (3,2)の同列
	if b.Get(2, 2) != "x" {
		t.Errorf("Failed to reverse \"o\" on (2,2)")
	}
	if b.Get(4, 2) != "x" {
		t.Errorf("Failed to reverse \"o\" on (4,2)")
	}
	// (3,2)の下一列
	if b.Get(2, 3) != "x" {
		t.Errorf("Failed to reverse \"o\" on (2,3)")
	}
	if b.Get(3, 3) != "x" {
		t.Errorf("Failed to reverse \"o\" on (3,3)")
	}
	if b.Get(4, 3) != "x" {
		t.Errorf("Failed to reverse \"o\" on (4,3)")
	}
}

// プレイヤー1("o")が駒を置けないテスト
func TestPutToken_P1_Unplaceable(t *testing.T) {
	// 盤面
	b := &Board{
		tokens: []int{
			0, 0, 1, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
		},
	}
	// プレイヤー1が(4,2)に駒を置く
	b.CurrentTurn = 1      // プレイヤー1のターン
	b.Reverse(0, 0)        // (0,0)に駒を置く
	println(b.BoardShow()) // 盤面の表示(テストのテスト)
	// (0,0)に置けない
	if b.Get(0, 0) != "." {
		t.Errorf("Wrong putting")
	}
}

// 引き分けになるテスト1
func TestCheckWin_Draw1(t *testing.T) {
	// 盤面
	b := &Board{
		tokens: []int{
			0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 1, 2, 0, 0, 0,
			0, 0, 0, 2, 1, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
		},
	}
	println(b.BoardShow()) // 盤面の表示(テストのテスト)
	// 引き分け
	if b.WinCheck() != "d" {
		t.Errorf("This game should be draw")
	}
}

// 引き分けになるテスト2
func TestCheckWin_Draw2(t *testing.T) {
	// 盤面
	b := &Board{
		tokens: []int{
			1, 2, 1, 2, 1, 2, 1, 2,
			1, 2, 1, 2, 1, 2, 1, 2,
			1, 2, 1, 2, 1, 2, 1, 2,
			1, 2, 1, 2, 1, 2, 1, 2,
			1, 2, 1, 2, 1, 2, 1, 2,
			1, 2, 1, 2, 1, 2, 1, 2,
			1, 2, 1, 2, 1, 2, 1, 2,
			1, 2, 1, 2, 1, 2, 1, 2,
		},
	}
	println(b.BoardShow()) // 盤面の表示(テストのテスト)
	// 引き分け
	if b.WinCheck() != "d" {
		t.Errorf("This game should be draw")
	}
}

// プレイヤー1("o")が勝利するテスト1
func TestCheckWin_P1_Win1(t *testing.T) {
	// 盤面(簡単)
	b := &Board{
		tokens: []int{
			0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 1, 0, 0, 0,
			0, 0, 0, 1, 1, 0, 0, 0,
			0, 0, 0, 2, 1, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
		},
	}
	println(b.BoardShow()) // 盤面の表示(テストのテスト)
	// プレイヤー1("o")が勝利
	if b.WinCheck() != "o" {
		t.Errorf("Player1(\"o\") should win")
	}
}

// プレイヤー1("o")が勝利するテスト2
func TestCheckWin_P1_Win2(t *testing.T) {
	// 盤面(僅差)
	b := &Board{
		tokens: []int{
			1, 1, 1, 1, 1, 1, 1, 1,
			0, 2, 2, 2, 2, 2, 2, 2,
			0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
		},
	}
	println(b.BoardShow()) // 盤面の表示(テストのテスト)
	// プレイヤー1("o")が勝利
	if b.WinCheck() != "o" {
		t.Errorf("Player1(\"o\") should win")
	}
}

// プレイヤー1("o")が勝利するテスト3
func TestCheckWin_P1_Win3(t *testing.T) {
	// 盤面(1だけ)
	b := &Board{
		tokens: []int{
			1, 0, 1, 0, 0, 1, 1, 1,
			0, 1, 0, 1, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 1,
			0, 0, 0, 1, 0, 0, 1, 0,
			0, 0, 1, 0, 0, 1, 0, 0,
			0, 1, 0, 0, 1, 0, 0, 0,
			0, 0, 0, 1, 0, 0, 0, 0,
			0, 0, 1, 0, 0, 0, 0, 0,
		},
	}
	println(b.BoardShow()) // 盤面の表示(テストのテスト)
	// プレイヤー1("o")が勝利
	if b.WinCheck() != "o" {
		t.Errorf("Player1(\"o\") should win")
	}
}

// プレイヤー2("x")が勝利するテスト1
func TestCheckWin_P2_Win1(t *testing.T) {
	// 盤面(簡単)
	b := &Board{
		tokens: []int{
			0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 1, 2, 0, 0, 0,
			0, 0, 0, 2, 2, 2, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
		},
	}
	println(b.BoardShow()) // 盤面の表示(テストのテスト)
	// プレイヤー2("x")が勝利
	if b.WinCheck() != "x" {
		t.Errorf("Player2(\"x\") should win")
	}
}

// プレイヤー2("x")が勝利するテスト2
func TestCheckWin_P2_Win2(t *testing.T) {
	// 盤面(僅差)
	b := &Board{
		tokens: []int{
			0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
			2, 2, 2, 2, 1, 1, 1, 1,
			1, 1, 1, 1, 1, 1, 1, 1,
			2, 2, 2, 2, 2, 2, 2, 2,
			0, 0, 0, 0, 0, 0, 0, 2,
			0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
		},
	}
	println(b.BoardShow()) // 盤面の表示(テストのテスト)
	// プレイヤー2("x")が勝利
	if b.WinCheck() != "x" {
		t.Errorf("Player2(\"x\") should win")
	}
}

// プレイヤー2("x")が勝利するテスト3
func TestCheckWin_P2_Win3(t *testing.T) {
	// 盤面(2だけ)
	b := &Board{
		tokens: []int{
			0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 2, 0, 0, 0, 0, 0,
			0, 2, 0, 0, 2, 0, 0, 0,
			0, 0, 2, 0, 2, 0, 0, 0,
			0, 0, 0, 2, 2, 0, 0, 0,
			0, 0, 0, 0, 2, 0, 0, 2,
			0, 0, 0, 0, 2, 0, 2, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
		},
	}
	println(b.BoardShow()) // 盤面の表示(テストのテスト)
	// プレイヤー2("x")が勝利
	if b.WinCheck() != "x" {
		t.Errorf("Player2(\"x\") should win")
	}
}
