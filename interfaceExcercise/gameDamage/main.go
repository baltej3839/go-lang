package main

import "fmt"


type Attacker interface {
	Attack(Targeter)
}

type Targeter interface {
	TakeDamage(int)
}

type Player struct {
	name  string
	speak bool
	health int 
}


type Zombie struct {
	noSleep      bool
	infinitywalk bool
	health int 
}



type Dragon struct {
	fly bool
	health int 
}


func (p *Player) TakeDamage(amount int) {
	fmt.Print("Damange talken")
}

func (p *Zombie) TakeDamage(amount int) {
	fmt.Print("Damange talken")
}

func (p *Dragon) TakeDamage(amount int) {
	fmt.Print("Damange talken")
}


func (p *Player) Attack(t Targeter) {
	t.TakeDamage(5)
}

func (d *Dragon) Attack(t Targeter) {
	t.TakeDamage(5)
}

func (z *Zombie) Attack(t Targeter) {
	t.TakeDamage(5)
}


func main() {
	var p=&Player{name:"Sam", speak:true, health:100}
	var d=&Dragon{fly:true, health:100}
	p.Attack(d)
}