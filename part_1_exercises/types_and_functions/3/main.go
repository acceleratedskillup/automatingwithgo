package main

import "fmt"

type Potion struct {
	potency  float64
	duration int
	quantity int
}

func (p *Potion) drink() {
	if p.quantity <= 0 {
		fmt.Println("No more potion left!")
		return
	}
	p.quantity--
	effect := p.potency * float64(p.duration)
	if effect > 50 {
		fmt.Println("You are completely invisible!")
	} else {
		fmt.Println("You are partially invisible. Be cautious!")
	}
}

func main() {
	invisibilityPotion := &Potion{
		potency:  0.8,
		duration: 60,
		quantity: 3,
	}
	invisibilityPotion.drink()
}
