package main

import (
	"fmt"
	"os"
	"math/rand"
	"time"
)

// main character structure
type Character struct{
	Name string
	Health int
	Choice int
	Weapon bool
	Potion int
}

// in the events of a bear attack
func bear() {
	fmt.Println("\nA bear storms in and kills you until you are dead.\n\nGAME OVER\n")
	os.Exit(0)
}

// function responsible for the attack exchange that accounts for the player having a weapon
func attack(attacker *Character, victim *Character){
	rand.Seed(time.Now().UnixNano())
	var damage int
	if attacker.Weapon == true {
		damage = rand.Intn(55) + 15
		victim.Health = victim.Health - damage
	} else {
		damage = rand.Intn(55)
		victim.Health = victim.Health - damage
	}
	if damage > 50{
		fmt.Println("\nCritical Strike!")
	}
	fmt.Printf("\n%v deals %v damage to %v!\n", attacker.Name, damage, victim.Name)
}

// the function responsible for the player healing by using potions
func (player *Character) heal(){
	if player.Potion == 1{
		player.Health += 50
		player.Potion = player.Potion - 1
		fmt.Printf("%v drank a potion healing him 50 hp!", player.Name)
	}
}

// the boss encounter
func (player *Character) bossRoom(){
	fmt.Printf("\n=-=-=-=-=-=-=-=-=-=-=\n\nAfter some time of walking in complete darknes, %v sees a light at the end of the tunnel.\n", player.Name)
	fmt.Println("\n=-=-=-=-=-=-=-=-=-=-=\n\nYou finally reach the end of the tunnel and enter a huge room.\nAt the end of it, you see a big scary looking man...\n\n")
	fmt.Println("Enter... SUPER EVIL DEVIL MAN!!!!\n\nSuper Evil Devil Man:\n'Whom'st've're are you, my child?'\n'Nevermind, WELCOME TO DIE!'\n")

	boss := &Character{Name: "Super Evil Devil Man", Health: 100, Choice: 0, Weapon: false, Potion: 0}
	fmt.Println("\nThe Battle begins!\n")
	for player.Health > 0{
		if boss.Health > 0 {
			fmt.Printf("\n\n%v's HP: %v\n%v's HP: %v\n\n",player.Name, player.Health, boss.Name, boss.Health)
			fmt.Println("OPTIONS:\n1) Attack\n2) Drink Potion\n3) Surrender\n\n")
			fmt.Scanln(&player.Choice)
			switch player.Choice{
			case 1:
				attack(player, boss)
				attack(boss, player)
			case 2:
				player.heal()
				attack(boss, player)

			case 3:
				fmt.Println("\nSuper Evil Devil Man:\n'Heh... nothing personnel kid...'\n\nSuper Evil Devil Man teleports behind you and slays you!\n\nGAME OVER")
				os.Exit(0)
		}
		}else{
			fmt.Println("\n=-=-=-=-=-=-=-=-=-=-=\n\nYou did it! You slayed the Super Evil Devil Man!\n\nYou finally reach the surface and find yourself on the streets of Brooklyn!\n")
			fmt.Println("Home sweet home...")
			fmt.Println("\n\nTHE END\n")
			os.Exit(0)
		}
	}
	if player.Health < 0 {
		fmt.Println("\nYou died and your corpse rots in the dungeon forever...\n\nGAME OVER")
		os.Exit(0)
	}
	}

// the tunnel where the player gets to choose between investigating a silhouette or not
func (player *Character) tunnel() {
	fmt.Println("\n=-=-=-=-=-=-=-=-=-=\n\nAs you proceed down the dark tunnel, you see a mysterious silhouette...\nDo you wish to reach for it?\n\nOPTIONS:\n1)Yes\n2)No\n")
	fmt.Scanln(&player.Choice)
	switch player.Choice {
	case 1:
		player.Weapon = true
		player.Potion = 1
		fmt.Printf("\n%v has found a sword and a healing potion!\n%v safely proceeds down the tunnel...\n", player.Name, player.Name)
		player.bossRoom()
	case 2:
		fmt.Printf("\nYou decided not to touch whatever was in the darkness.\n%v safely proceeds down the tunnel...\n", player.Name)
		player.bossRoom()
	default:
		bear()
	}
}

// the function responsible for the game beginning
func play() {
	var player Character
	player.Health = 100
	fmt.Println("\nWhat is your name Adventurer?\n")
	fmt.Scanln(&player.Name)
	fmt.Printf("\n=-=-=-=-=-=-=-=-=-=\n\n%v wakes up in a dark cavern...\n", player.Name)
	for player.Choice != 1 || player.Choice != 3{
		fmt.Printf("\nOPTIONS:\n1) Look around\n2) Scream\n3) Call %v's mommy\n\n", player.Name)
		fmt.Scanln(&player.Choice)
		switch player.Choice{
		case 2:
			fmt.Println("\nYou scream.\n\n  'AAAAAAAAAAAAAAAAAAAAAAA'\n\n...Nothing happens.\n")
		case 3:
			bear()
		case 1:
			fmt.Println("\nYou see a dark tunnel ahead...\nDo you wish to proceed?\n\nOPTIONS:\n1)Yes\n2)No, I wish to stay here\n")
			fmt.Scanln(&player.Choice)
				switch player.Choice{
				case 1:
					player.tunnel()
				default:
					bear()
				}
		}
	}
}

// Main method
func main() {
	var choice int
	// Title menu
	fmt.Println("\nWelcome mortal............\nTo the SECRET DARK UNDERGROUND AWESOME EVIL ADVENTURE\n\n")
	fmt.Println("Select an option and press Enter:")
	fmt.Println("\n1) New Game\n\nAnything else => Exit\n")
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		play()
	default:
		fmt.Println("\nExiting...\n")
	}
}

