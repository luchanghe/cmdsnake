package Game

import (
	"fmt"
	"math"
	"strings"
	"time"
)

/*
#include <windows.h>
#include <conio.h>
// 使用了WinAPI来移动控制台的光标
void gotoxy(int x,int y)
{
    COORD c;
    c.X=x,c.Y=y;
    SetConsoleCursorPosition(GetStdHandle(STD_OUTPUT_HANDLE),c);
}
// 从键盘获取一次按键，但不显示到控制台
int direct()
{
    return _getch();
}
*/
import "C" // go中可以嵌入C语言的函数

type Game struct {
	snake   *Snake
	plat    *Plat
	food    *Food
	hinder  [][]string
	status  int
	sorce   int
	quicken int
}

type XY struct {
	x int
	y int
}

func NewGame(snake *Snake, plat *Plat, quicken int) *Game {
	return &Game{snake: snake, plat: plat, food: NewFood(2, 2), hinder: [][]string{}, quicken: quicken}
}

func (g Game) Listen() {
	g.Flush()
	ticker := time.NewTicker(time.Second - time.Duration(g.quicken*100)*time.Millisecond)
	go g.Crawl()
	for {
		select {
		case <-ticker.C:
			if g.status == 1 {
				g.Flush()
				break
			}
			head := g.snake.body[0]
			x := head.x
			y := head.y
			if g.snake.direction == normal {
				break
			}
			//修正蛇皮走位
			if (g.snake.lastDirection == top && g.snake.direction == down) ||
				(g.snake.lastDirection == down && g.snake.direction == top) ||
				(g.snake.lastDirection == left && g.snake.direction == right) ||
				(g.snake.lastDirection == right && g.snake.direction == left) {
				g.snake.direction = g.snake.lastDirection
			}
			g.snake.lastDirection = g.snake.direction
			switch g.snake.direction {
			case top:
				x--
			case down:
				x++
			case left:
				y--
			case right:
				y++
			}

			//撞墙游戏结束
			if g.hinder[x][y] == "墙" {
				g.status = 1
			}
			//撞蛇(撞尾巴没事 因为要缩进)
			if g.hinder[x][y] == "蛇" {
				footer := g.snake.body[len(g.snake.body)-1]
				if x != footer.x && y != footer.y {
					g.status = 1
				}
			}
			if g.hinder[x][y] == "食" {
				g.snake.body = append([]XY{{x, y}}, g.snake.body[:]...)
				g.food.Range(g)
				if g.food.x == 0 && g.food.y == 0 {
					g.status = 2
				}
				g.sorce++
			} else {
				g.snake.body = append([]XY{{x, y}}, g.snake.body[0:len(g.snake.body)-1]...)
			}
			g.Flush()
		}
	}
}

func (g *Game) Flush() {
	p := g.plat
	s := g.snake
	g.hinder = p.Build()
	switch g.status {
	case 0:
		//游戏正常创建蛇体
		for _, v := range s.body {
			i := v.x
			j := v.y
			g.hinder[i][j] = "蛇"
		}

		g.hinder[g.food.x][g.food.y] = "食"
	case 1:
		//游戏结束绘画
		i := int(math.Floor(float64(p.weight/2 - 1)))
		j := int(math.Floor(float64(p.height/2) - 2))
		g.hinder[i][j] = "游"
		j++
		g.hinder[i][j] = "戏"
		j++
		g.hinder[i][j] = "结"
		j++
		g.hinder[i][j] = "束"
	case 2:
		//游戏结束绘画
		i := int(math.Floor(float64(p.weight/2 - 1)))
		j := int(math.Floor(float64(p.height/2) - 2))
		g.hinder[i][j] = "成"
		j++
		g.hinder[i][j] = "功"
		j++
		g.hinder[i][j] = "通"
		j++
		g.hinder[i][j] = "关"

	}

	C.gotoxy(0, 0)
	fmt.Printf("小键盘上下左右操作，空格重新开始\t\t\t\t\r\n当前分数：%d \r\n", g.sorce)
	fmt.Println(strings.Repeat("=", g.plat.weight*3))
	for _, value := range g.hinder {
		fmt.Println(value)
	}

}

func (g *Game) Crawl() {
	go func() {
		for {
			switch byte(C.direct()) {
			case 72:
				g.snake.direction = top
			case 75:
				g.snake.direction = left
			case 77:
				g.snake.direction = right
			case 80:
				g.snake.direction = down
			case 32:
				if g.status != 1 {
					break
				}
				g.status = 0
				g.sorce = 0
				g.snake = NewSnake()
				g.plat = NewPlat(g.plat.weight, g.plat.height)
				g.food.Range(*g)
				g.Flush()
			}
		}
	}()
}
