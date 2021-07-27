// 挿入, 検索, 削除が可能. insert.goに追記
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	Key int
	Parent *Node
	Left *Node
	Right *Node
}

type Tree struct {
	Root *Node
	NIL *Node
}

func NewTree() *Tree {
	tree := new(Tree)
	Nil := new(Node)
	(*tree).NIL = Nil
	(*tree).Root = Nil
	return tree
}

func (t *Tree) Insert(key int) {
	parent := (*t).NIL
	leaf := (*t).Root

	z := new(Node)
	(*z).Key = key
	(*z).Left = (*t).NIL
	(*z).Right = (*t).NIL

	for leaf != (*t).NIL {
		parent = leaf
		if (*z).Key < (*leaf).Key {
			leaf = (*leaf).Left
		} else {
			leaf = (*leaf).Right
		}
	}

	(*z).Parent = parent

	if parent == (*t).NIL {
		(*t).Root = z
	} else {
		if (*z).Key < (*parent).Key {
			(*parent).Left = z
		} else {
			(*parent).Right = z
		}
	}
}

// GetMinimum ではノードuから辿った時の木の最小値を持つノードを取得する
func (t *Tree) GetMinimum(u *Node) *Node {
	for (*u).Left != (*t).NIL {
		u = (*u).Left
	}
	return u
}

// Find では値kを持つノードを取得する
func (t *Tree) Find(k int) *Node {
	u := (*t).Root
	for u != (*t).NIL && k != (*u).Key {
		if k < (*u).Key {
			u = (*u).Left
		} else {
			u = (*u).Right
		}
	}
	return u
}

// GetSuccessor ではInOrderで巡回した時にノードuの次に来るノードを取得する
func (t *Tree) GetSuccessor(u *Node) *Node {
	// 左の子は巡回した時に前に来るので無視
	// もし右の子がある場合は右の子の一番小さい数字を持つノードが次に来る
	if (*u).Right != (*t).NIL {
		return (*t).GetMinimum((*u).Right)
	}
	// 右の子がない場合は親を辿って右の子がないか探していく
	y := (*u).Parent
	for y != (*t).NIL && u == (*y).Right {
		u = y
		y = (*y).Parent
	}
	return y
}

// Delete ではREADME.mdに記載した手順でノードの削除を行う。詳細説明はREADMEを参照。
func (t *Tree) Delete(u *Node) {
	switch {
	case (*u).Left == (*t).NIL && (*u).Right == (*t).NIL:  // zが子を持たない時
		switch {
		case (*u).Parent == (*t).NIL:  // zが根の時
			(*t).Root = (*t).NIL
		case u == (*(*u).Parent).Left:  // zが親の左の子の時
			(*(*u).Parent).Left = (*t).NIL
		default:  // zが親の右の子の時
			(*(*u).Parent).Right = (*t).NIL
		}
	case (*u).Left == (*t).NIL || (*u).Right == (*t).NIL:  // zが片方の子のみ持つ時
		x := map[bool]*Node{true: (*u).Left, false: (*u).Right}[(*u).Left != (*t).NIL]  // その片方の子をxとする
		switch {
		case (*u).Parent == (*t).NIL:  // zが根の時
			(*x).Parent = (*t).NIL
			(*t).Root = x
		case u == (*(*u).Parent).Left:  // zが親の左の子の時
			(*x).Parent = (*u).Parent
			(*(*u).Parent).Left = x
		default:  // zが親の右の子の時
			(*x).Parent = (*u).Parent
			(*(*u).Parent).Right = x
		}
	default:  // zが子を両方持つ場合
		y := (*t).GetSuccessor(u)
		(*t).Delete(y)
		(*u).Key = (*y).Key
	}
}

func (t *Tree) InOrder(u *Node) {
	if u == (*t).NIL {
		return
	}
	(*t).InOrder((*u).Left)
	fmt.Printf("%d ", (*u).Key)
	(*t).InOrder((*u).Right)
}

func (t *Tree) PreOrder(u *Node) {
	if u == (*t).NIL {
		return
	}
	fmt.Printf("%d ", (*u).Key)
	(*t).PreOrder((*u).Left)
	(*t).PreOrder((*u).Right)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	numCommands, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println("input error")
		return
	}

	tree := NewTree()

	for i := 0; i < numCommands; i++ {
		scanner.Scan()
		command := strings.Split(scanner.Text(), " ")

		if command[0] == "print" {
			fmt.Println("\nInOrder")
			tree.InOrder(tree.Root)
			fmt.Println("\nPreOrder")
			tree.PreOrder(tree.Root)
		} else if len(command) > 1 {  // 2つ目の入力値がある場合
			key, err := strconv.Atoi(command[1])
			if err != nil {
				fmt.Println("input error")
				return
			}
			if command[0] == "insert" {  // 挿入の場合：挿入
				tree.Insert(key)
			} else {  // 検索・削除の場合
				t := tree.Find(key)
				if t != tree.NIL {  // 該当するノードがある場合
					if command[0] == "find" {  // 検索の場合：検索結果はyes
						fmt.Println("yes")
					} else if command[0] == "delete" {  // 削除の場合：削除
						tree.Delete(t)
					} else {
						fmt.Println("input error")
						return
					}
				} else {
					fmt.Println("no")
				}
			}
		} else {
			fmt.Println("input error")
			return
		}
	}
}

