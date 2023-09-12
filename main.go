package main

import "fmt"

type ZooKeeper struct {
	FirstName string
	LastName string
	Age int
	ZooCage ZooCage
}

type ZooCage struct {
	Number int
	Animal Animal
}

type Animal struct {
	Name string
	Age int
	Present bool
}

// check if Animal is present in a cage
func (Animal *Animal) IsPresent() bool {
	return Animal.Present
}

// Change presense of the animal in a cage from present to not present
func (Animal *Animal) Steal() {
	Animal.Present = false
}

// Change presense of the animal in a cage from not present to  present
func (Animal *Animal) Return() {
	Animal.Present = true
}

// Steal  or Return Animal
func (ZooKeeper *ZooKeeper) StealReturnAnimal() {
	if ZooKeeper.ZooCage.Animal.IsPresent() {
		ZooKeeper.ZooCage.Animal.Steal()
	} else {
		ZooKeeper.ZooCage.Animal.Return()
	}
	}

func main() {
// there are two zookeepers (Bob and Rob), four cages and four animals, each zookeepers is responible for two animals
	zk1 := ZooKeeper{
		FirstName: "Bob", LastName: "Bobenko", Age: 30, ZooCage: ZooCage{
			Number: 1, Animal: Animal{
				Name: "Wolf", Age: 5, Present: true,
			},
// was trying to add two animals under the same zookeepr but got next error from compiler
// duplicate field name in struct literal: Number
// duplicate field name in struct literal: Animal			
//			Number: 2, Animal: Animal{
//				Name: "Lion", Age: 4, Present: true,
//			},
		},
	}
	zk2 := ZooKeeper{
		FirstName: "Bob", LastName: "Bobenko", Age: 30, ZooCage: ZooCage{
			Number: 2, Animal: Animal{
				Name: "Lion", Age: 4, Present: true,
			},
		},
	}
	zk3 := ZooKeeper{
		FirstName: "Rob", LastName: "Robenko", Age: 40, ZooCage: ZooCage{
			Number: 3, Animal: Animal{
				Name: "Bear", Age: 3, Present: true,
			},
		},
	}
	zk4 := ZooKeeper{
		FirstName: "Rob", LastName: "Robenko", Age: 40, ZooCage: ZooCage{
			Number: 4, Animal: Animal{
				Name: "Fox", Age: 1, Present: true,
			},
		},
	}
// Going to check if all four variables are initialized 
	fmt.Println(zk1)
	fmt.Println(zk2)
	fmt.Println(zk3)
	fmt.Println(zk4)

	fmt.Printf("status of the animal zk1 : %+v\n", zk1)
	zk1.StealReturnAnimal()
	fmt.Printf("status of the amimal zk1 : %+v\n", zk1)
	fmt.Printf("Is animal in a cage? : %t\n", zk1.ZooCage.Animal.IsPresent())
	zk1.StealReturnAnimal()
	fmt.Printf("Is animal in a cage? : %t\n", zk1.ZooCage.Animal.IsPresent())
}
