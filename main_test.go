package main

import (
	"testing"
)

// 初期盤面からプレイヤー1("o")が(3,5)に駒を置くテスト
func TestPutToken_1P35_InitialBoard(t *testing.T) {
	// 初期盤面
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
		t.Errorf("Failed to reverse \"o\" on (3,4)")
	}
}

// 初期盤面からプレイヤー1("o")が(5,3)に駒を置くテスト
func TestPutToken_1P53_InitialBoard(t *testing.T) {
	// 初期盤面
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
	b.Reverse(5, 3)        // (5, 3)に駒を置く
	println(b.BoardShow()) // 盤面の表示(テストのテスト)
	// (5,3)が"o"になる
	if b.Get(5, 3) != "o" {
		t.Errorf("Failed to put \"o\" on (5,3)")
	}
	// (4,3)も"o"になる
	if b.Get(4, 3) != "o" {
		t.Errorf("Failed to reverse \"o\" on (4,3)")
	}
}
