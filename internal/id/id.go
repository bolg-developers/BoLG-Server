package id

import "strconv"

type Manager struct {
	max     int
	current int
}

func NewManager(max int) *Manager {
	return &Manager{
		max:     max,
		current: -1, // 生成値が0から始まるように
	}
}

func (g *Manager) Increment() (int, bool) {
	if g.current == g.max {
		return 0, false
	}
	g.current++
	return g.current, true
}

func (g *Manager) Decrement() (int, bool) {
	if g.current == -1 {
		return 0, false
	}
	g.current--
	return g.current, true
}

func (g *Manager) CurrentId() int {
	return g.current
}

func ConvertString(id int32) string {
	return strconv.Itoa(int(id))
}
