package main

import "fmt"

// Primitive types
var (
	varFloat32       float32 = 0.1
	varFloat64       float64 = 0.1
	varString        string  = "Some String"
	varInteger       int     = 1
	varIntegerNeg    int     = -1
	varInteger32     int32   = 1
	varInteger64     int64   = 1
	varUnsignedInt   uint    = 1
	varUnsignedInt8  uint8   = 255
	varUnsignedInt32 uint32  = 1
	varUnsignedInt64 uint64  = 1
	varByte          byte    = 0x2
	varRune          rune    = 'a'
)

// Structs
type Player struct {
	name        string
	health      int
	attackPower float64
}

// Function
// func getHealth(player Player) int {
// 	return player.health
// }

// Function receivers are functions attached to structs
func (player Player) getHealth() int {
	return player.health
}

func main() {
	player_00 := Player{}
	player_01 := Player{
		name:        "Captain Jack",
		health:      100,
		attackPower: 33.3,
	}

	fmt.Printf("This is player 00: %v\n", player_00)
	// => This is player 00: { 0 0}
	fmt.Printf("This is player 01: %v\n", player_01)
	// => This is player 01: {Captain Jack 100 33.3}
	fmt.Printf("This is player 01: %+v\n", player_01)
	// => This is player 01: {name:Captain Jack health:100 attackPower:33.3}
	// fmt.Printf("Player health: %d\n", getHealth(player_01))
	// => Player health: 100
	fmt.Printf("Player health: %d\n", player_01.getHealth())
	// => Player health: 100

	// Maps
	emptyMap := map[string]int{}
	fmt.Printf("%v\n", emptyMap)
	// => map[]

	valueMap := map[string]int{
		"daz": 38,
	}
	valueMap["a"] = 1
	valueMap["b"] = 2
	fmt.Printf("%v\n", valueMap)
	// => map[a:1 b:2 daz:38]
}
