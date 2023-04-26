package Game

type Plat struct {
	weight int
	height int
}

func NewPlat(weight int, height int) *Plat {
	return &Plat{weight: weight, height: weight}
}

func (p *Plat) Build() [][]string {
	//渲染围墙和空白区
	hinder := make([][]string, p.weight)
	for i := 0; i < p.weight; i++ {
		for j := 0; j < p.height; j++ {
			if i == 0 || j == 0 || i == p.weight-1 || j == p.height-1 {
				hinder[i] = append(hinder[i], "墙")
			} else {
				hinder[i] = append(hinder[i], "  ")
			}
		}
	}
	return hinder
}
