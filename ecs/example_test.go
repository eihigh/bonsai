package ecs_test

import "github.com/eihigh/bonsai/ecs"

type foo struct {
	ecs.DeletedFlag
	x, y int
}

type mover interface {
	ecs.Entity
	move()
	getPos() (x, y int)
}

type drawer interface {
	ecs.Entity
	draw()
}

var movers []mover

var drawers []drawer

var screenWidth = 64

func newFoo() *foo {
	f := &foo{x: 0, y: 0}
	movers = append(movers, f)
	drawers = append(drawers, f)
	return f
}

func (f *foo) getPos() (x, y int) {
	return f.x, f.y
}

func (f *foo) move() {
	f.x += 8
	// Once f is outside the screen, it is deleted as a mover in the mover system.
	// Entities deleted in one system will also be deleted from other systems and
	// will eventually be garbage-collected.
}

func (f *foo) draw() {
	// draw something
}

func Example() {
	for i := 0; i < 99; i++ {
		newFoo()

		// mover system
		ecs.Slice(movers).SweepAll(func(_ int, m mover) bool {
			m.move()
			x, _ := m.getPos()
			if x > screenWidth {
				m.Delete()
			}
			return true
		})

		// When "range over func" is added to Go, it can be written as follows:
		// for _, m := range ecs.Slice(movers).SweepAll { m.move(); ... }

		// drawer system
		ecs.Slice(drawers).SweepAll(func(_ int, d drawer) bool {
			d.draw()
			return true
		})
	}
}
