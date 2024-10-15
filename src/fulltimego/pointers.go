package main

import "fmt"

type Player struct {
	HP int
}

// If player is not a pointer, we are adjusting the copy of the player, not the actual player.
func TakeDamage(player *Player, amount int) {
	player.HP -= amount
	fmt.Println("Player is taking damage . New HP ->", player.HP)
}

func (p *Player) TakeDamage(amount int) {
	p.HP -= amount
	fmt.Println("player is taking damage. New HP ->", p.HP)
}

type Database struct {
	User string
}

type Server struct {
	db *Database
}

func (s *Server) GetUserFromDB() string {
	return s.db.User
}

func main() {
	player := Player{
		HP: 100,
	}
	TakeDamage(&player, 10)
	fmt.Println(player)

	s := &Server{
		//db: &Database{},
	}
	if s.db == nil {
		panic("db is nil")
	}
	s.GetUserFromDB()
}
