package main

// Example: In composition, if the inner type has a common method than the outer type, when invoked, the inner is the receiver, not the outer; as opposed to OOP, where we'd have method override. 
// For more info of related topics, such as interface composition, see:
//  http://www.golangbootcamp.com/book/types
//  http://golang.org/doc/effective_go.html#embedding

import "fmt"

type User struct {
	Id             int
	Name, Location string
}

func (u *User) Greetings() string {
	return fmt.Sprintf("User: Hi %s from %s",
		u.Name, u.Location)
}

// If we didn't have this function, remember that Greetings would still be callable on Player, as functions defined on the inner type become part of the outer type's interface
// Also, remember that if the inner type implements n interfaces, the outer type automatically implement those. 
func (p *Player) Greetings() string {
	return fmt.Sprintf("Player: Hi %s from %s",
		p.Name, p.Location)}

// It's better to compose structs through pointers
type Player struct {
	*User
	GameId int
}

func main() {
	// Literal composition
	p := &Player{&User{ 1, "Optimus", "Mexico" } , 4423}
	fmt.Printf("player is %v", p.Greetings())
}
