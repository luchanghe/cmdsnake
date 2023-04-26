package Game

import "math/rand"

type Food struct {
	x int
	y int
}

func NewFood(x int, y int) *Food {
	return &Food{x, y}
}

func (f *Food) Range(g Game) {
	tempMap := make([]XY, 0)
	k := 0
	for i := 1; i < g.plat.weight; i++ {
		for j := 1; j < g.plat.weight; j++ {
			if g.hinder[i][j] != "蛇" && g.hinder[i][j] != "墙" {
				tempMap = append(tempMap, XY{i, j})
				k++
			}
		}
	}
	if len(tempMap) == 0 {
		f.x = 0
		f.y = 0
	} else {
		v := rand.Intn(k)
		f.x = tempMap[v].x
		f.y = tempMap[v].y
	}

}
