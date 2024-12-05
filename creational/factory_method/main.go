// package main

// import "fmt"
// type : Transport interface {
// 	setName{name string}
// 	setPower{power int}
// 	getName{} string
// 	getPower{} int
// }

// type : Transport strust {
// 	name string
// 	power int
// }

// func (g* transport) setName(name string){
// 	g*name = name 
// }

// func (g* transport) getName(name string){
// 	return g*name
// }

// func (g* transport) setPower(power int){
// 	g*power = power 
// }

// func (g* transport) getPower(power int){
// 	return g*power 
// }

// type truct struct {
// 	transport
// }

// func newTruck() iTransport {
// 	return &truck{
// 		transport: transport(
// 			name: "Truck";
// 			power: 4,
// 		)
// 	}
// }


// type ships struct {
// 	transport
// }

// func newShips() iTransport {
// 	return &ships{
// 		transport: transport(
// 			name: "ships";
// 			power: 5,
// 		)
// 	}
// }

// func getTransport(transport:Type string) iTransport {
// 	if transportTypec == "Truck"{
// 		return newTruck{},nil
// 	}
// 	if transportTypec == "Ships"{
// 		return newShips{},nil
// 	}
// 	return nil,fmt.Errorf("Wrong transport type")

// }


// func main () {
// 	truck _ := getTransport("truck")
// 	ships_:= getTransport("ships")
// 	printDetails(truck)
// 	printDetails(ships)
// }

// func printDetails(g :Transport) {
// 	fmt.Printf("Name: %s", g.getName())
// 	fmt.Println()
// 	fmt.Printf("Power: %s", g.getPower())
// 	fmt.Println()
// }

//////////////////////////////////////////////////////////////////////////////////////////////////////////////


package main

import "fmt"

// Transport interface
type Transport interface {
	setName(name string)
	setPower(power int)
	getName() string
	getPower() int
}

// Transport struct
type transport struct {
	name  string
	power int
}

// Implementing Transport methods
func (g *transport) setName(name string) {
	g.name = name
}

func (g *transport) getName() string {
	return g.name
}

func (g *transport) setPower(power int) {
	g.power = power
}

func (g *transport) getPower() int {
	return g.power
}

// Truck struct
type truck struct {
	transport
}

func newTruck() Transport {
	return &truck{
		transport: transport{
			name:  "Truck",
			power: 4,
		},
	}
}

// Ships struct
type ships struct {
	transport
}

func newShips() Transport {
	return &ships{
		transport: transport{
			name:  "Ships",
			power: 5,
		},
	}
}

// Factory function
func getTransport(transportType string) (Transport, error) {
	if transportType == "Truck" {
		return newTruck(), nil
	}
	if transportType == "Ships" {
		return newShips(), nil
	}
	return nil, fmt.Errorf("Wrong transport type")
}

// Helper function
func printDetails(g Transport) {
	fmt.Printf("Name: %s\n", g.getName())
	fmt.Printf("Power: %d\n", g.getPower())
}

func main() {
	truck, err := getTransport("Truck")
	if err != nil {
		fmt.Println(err)
		return
	}
	ships, err := getTransport("Ships")
	if err != nil {
		fmt.Println(err)
		return
	}

	printDetails(truck)
	printDetails(ships)
}
