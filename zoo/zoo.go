package zoo

import "fmt"

type ZooKeeper struct {
	ZooKeeperName string
	ZooKeeperAge int
	ZooCage ZooCage
}

type ZooCage struct {
	CageNumber int
	IsOpen bool
	Animal Animal
}

type Animal struct {
	AnimalName string
	AnimalAge int
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

// Check Cage if its open or not
func (cage *ZooCage) IsCageOpen() bool {
	return cage.IsOpen
}

// Open the Cage
func (cage *ZooCage) OpenTheCage() {
	cage.IsOpen = true
}

// Close the Cage
func (cage *ZooCage) CloseTheCage() {
	cage.IsOpen = false
}

func Zoo() {
// there are two zookeepers (Bob and Rob), four cages and four animals, each zookeepers is responible for two animals
	zk1 := ZooKeeper{
		ZooKeeperName: "Bob",  ZooKeeperAge: 30, ZooCage: ZooCage{
			CageNumber: 1, IsOpen: false,  Animal: Animal{
				AnimalName: "Wolf", AnimalAge: 5, Present: true,
			},
		},
	}
	zk2 := ZooKeeper{
		ZooKeeperName: "Bob",  ZooKeeperAge: 30, ZooCage: ZooCage{
			CageNumber: 2, IsOpen: false, Animal: Animal{
				AnimalName: "Lion", AnimalAge: 4, Present: true,
			},
		},
	}
	zk3 := ZooKeeper{
		ZooKeeperName: "Rob", ZooKeeperAge: 40, ZooCage: ZooCage{
			CageNumber: 3,  IsOpen: false, Animal: Animal{
				AnimalName: "Bear", AnimalAge: 3, Present: true,
			},
		},
	}
	zk4 := ZooKeeper{
		ZooKeeperName: "Rob", ZooKeeperAge: 40, ZooCage: ZooCage{
			CageNumber: 4,  IsOpen: false, Animal: Animal{
				AnimalName: "Fox", AnimalAge: 1, Present: true,
			},
		},
	}
// Going to check if all four variables are initialized 
	fmt.Println(zk1)
	fmt.Println(zk2)
	fmt.Println(zk3)
	fmt.Println(zk4)

	fmt.Printf("status of the animal zk1 : %+v\n", zk1)
	fmt.Printf("Is Cage 1 open? : %t\n", zk1.ZooCage.IsCageOpen() )
	fmt.Println("Lets open the cage")
	zk1.ZooCage.OpenTheCage()
	fmt.Printf("Is Cage 1 open? : %t\n", zk1.ZooCage.IsCageOpen() )
	zk1.StealReturnAnimal()
	fmt.Printf("Is animal in a cage? : %t\n", zk1.ZooCage.Animal.IsPresent())
	fmt.Println("Lets return animal to a cage")
	zk1.StealReturnAnimal()
	fmt.Printf("Is animal in a cage? : %t\n", zk1.ZooCage.Animal.IsPresent())
	fmt.Printf("Is Cage 1 open? : %t\n", zk1.ZooCage.IsCageOpen() )
	fmt.Println("Lets close the cage")
	fmt.Printf("Is Cage 1 open? : %t\n", zk1.ZooCage.IsCageOpen() )
	fmt.Printf("status of the animal zk1 : %+v\n", zk1)
}