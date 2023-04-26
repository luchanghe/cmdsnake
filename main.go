package main

import (
	"fmt"
	"snakeGame/Game"
)

/*
#include <stdio.h>
#include <windows.h>
void ConsoleCursor ();
//设置光标不可见
void sysInit(){
	HANDLE hOut;
	CONSOLE_CURSOR_INFO cur;
    cur.bVisible = 0;
	cur.dwSize = 1;
    hOut = GetStdHandle(STD_OUTPUT_HANDLE);
    SetConsoleCursorInfo(hOut, &cur);
}
*/
import "C"

func main() {
	C.sysInit()
	var weight, height, quicken int
	fmt.Println("请设定宽和高和加速(1-9)，三个值空格隔开")
	fmt.Scanf("%d %d %d", &weight, &height, &quicken)
	if weight <= 0 || height <= 0 || quicken <= 0 || quicken > 9 {
		panic("请输入整数,加速不能高于9")
	}
	plat := Game.NewPlat(weight, height)
	snake := Game.NewSnake()
	game := Game.NewGame(snake, plat, quicken)
	game.Listen()
}
