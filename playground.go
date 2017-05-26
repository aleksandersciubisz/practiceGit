package main

/*
Wzorcowy układ pliku to:

package main
import (
"fmt"
"inny pakiet"
)

func nazwa(wartości lub nic + typ) typ outputu {
	return /akcja/
}

func main () {

}
*/
import (
	"fmt"
	"math"
	"math/cmplx"
)

func add(x int, y int) int {
	return x + y
} //przykłady dodawania zmiennych

func golKrychowiaka() {
	goleKrychowiaka++
	goleKrychowiaka++

}

var (
	name     string
	age      int
	location string
)

/*
GOPATH powinien zawierać:
bin -binaries (compiled)
pkg -packages (compiled versions)
src -source code (organized by import path)
zaleca się tworzenie nowych programów w src
*/
var (
	name1, location1 string
	age1, numbers1   int
)

var name2 string
var age2 int
var location2 string

// + wartości początkowe

var (
	name3 string = "Sławek Peszko"
	age3  int    = 35
	// przy wartościach początkowych,
	// nie trzeba dodawać typów
	location3 = "Medelin"
)

// stałe dodajemy z przedrostkiem const
const golePeszkina = 0

var goleKrychowiaka = 1

func main() {
	fmt.Println(goleKrychowiaka)

	golKrychowiaka()
	fmt.Println(goleKrychowiaka)

	//wewnątrz fukcji, można używać znaku :=
	//czyli przypisania wartości do zmiennej
	name4, location4 := "Krychowiak", "Medelin"
	age4 := 33
	fmt.Printf("\n %s (lat %d) %s \n", name4,
		age4, location4)

	/*
		podstawowe znaczniki to:
		\n -nowa linia
		%s -string
		%d -integer
	*/

	// fmt.Println dodaje nową linię
	// fmt. Printf do fukcji i z formatami %
	fmt.Println("gole peszkina: \n",
		golePeszkina) //zauważ brak %d "pod gole peszki"

	fmt.Printf("W tym roku %s strzelił %d goli \n",
		name3, golePeszkina)

	fmt.Println(math.Pi)
	// Pi jest z dużej litery bo to wyeksportowana
	//część pakietu math. Duża litera = dostęp
	// z zewnątrz

	fmt.Println("\n dodawanie 11 i 2 \n result:",
		add(11, 2))

	//pointery są wyposażone w dwa znaczniki
	// & oznacza odwołanie (pointer of value)
	// * oznacza podległość (deference a pointer)
	//protip: go przekazuje wartości argumentów
	// jeżeli chcesz przekazać coś "po referencji"
	//to musisz użyć pointera

	/*
		podstawowe typy danych:
		-int
		-uint8, uint16, uint32, uint64 zestawy z bitami
		-int8, int16, int32, int64 wszystkie przypisane
		bity
		-float32, float64 set of IEEE-754-bit
		-byte, alias for uint8
		-rune, alias for int32
	*/
	//kilka przykładów podstawowych typów:
	var (
		peszkinToKozak bool       = true
		maxInt         uint64     = 1<<64 - 1
		complex1       complex128 = cmplx.Sqrt(-5 + 12i)
	)

	const f = "%T(%v)\n"
	fmt.Printf(f, peszkinToKozak, peszkinToKozak)
	fmt.Printf(f, maxInt, maxInt)
	fmt.Printf(f, complex1, complex1)
	const f2 = "%T\n(%v)\n"
	fmt.Printf(f2, peszkinToKozak, peszkinToKozak)
	fmt.Printf(f2, maxInt, maxInt)
	fmt.Printf(f2, complex1, complex1)

	//konwersja typów odbywa się wg schematu
	// T(v) :zmienia v na typ T
	//przykłady:

	var q1 int = 42
	var q2 float64 = float64(q1)
	var q3 uint = uint(q2)

	fmt.Printf("q1: %d\nq2: %v\nq3: %d\n",
		q1, q2, q3)
	fmt.Printf("typy q1, q2, q3 \n%T\n%T\n%T",
		q1, q2, q3)
}
