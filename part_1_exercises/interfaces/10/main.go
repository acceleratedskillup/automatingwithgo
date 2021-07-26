package main

import (
	"bufio"
	"fmt"
	"strings"
)

type Potion struct {
	Ingredients []string
}

func (p *Potion) Scan(state fmt.ScanState, verb rune) error {
	token, _, err := state.ReadRune()
	if err != nil {
		return err
	}
	ingredient := string(token)
	p.Ingredients = append(p.Ingredients, ingredient)
	return nil
}

func main() {
	potion := &Potion{}
	scanner := bufio.NewScanner(strings.NewReader("Mercury DragonScale UnicornHorn"))
	scanner.Split(bufio.ScanRunes)

	for scanner.Scan() {
		fmt.Sscan(scanner.Text(), potion)
	}

	fmt.Println("Potion Ingredients:", potion.Ingredients)
}
