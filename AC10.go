package main

//"bufio"

//"os"
//"strconv"
//"strings"

/*QUESTAO1
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	for i := 0; i < n; i++ {
		scanner.Scan()
		input := scanner.Text()
		numbers := strings.Split(input, " ")
		var lista []int

		for _, numStr := range numbers {
			num, _ := strconv.Atoi(numStr)
			lista = append(lista, num)
		}

		fmt.Println(mdc(lista[0], lista[1]))
	}
}

func mdc(n1 int, n2 int) int {
	for n2 != 0 {
		n1, n2 = n2, n1%n2
	}
	return n1
}
*/
/*Questao2
package main

import "fmt"

func abs(x int16) int16 {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	var x1, y1, x2, y2, resposta int16

	for {
		fmt.Scan(&x1, &y1, &x2, &y2)
		if x1 == 0 && y1 == 0 && x2 == 0 && y2 == 0 {
			break
		} else if x1 == x2 && y1 == y2 {
			fmt.Println("0")
		} else if x1 == x2 || y1 == y2 || abs(x1-x2) == abs(y1-y2) {
			fmt.Println("1")
		} else {
			fmt.Println("2")
		}
	}
}
*/
/* Questao3
package main

import "fmt"

func main() {
	var m, n, tmp, tmp2 int64
	var i int

	for {
		_, err := fmt.Scan(&m, &n)
		if err != nil {
			break
		}

		tmp = 1
		tmp2 = 1

		for i = int(m); i > 0; i-- {
			tmp *= m
			m--
		}

		for i = int(n); i > 0; i-- {
			tmp2 *= n
			n--
		}

		fmt.Println(tmp + tmp2)
	}
}
*/

/* Questao 4


func main() {
	var n, dias int
	var kg float64

	fmt.Scanln(&n)

	for i := 0; i < n; i++ {
		dias = 0
		fmt.Scanln(&kg)

		for kg > 1.0 {
			kg /= 2
			dias++
		}

		fmt.Printf("%d dias\n", dias)
	}
}

*/

/*Questao 5
package main

import (
	"fmt"
	"sort"
)

func main() {
	var N int
	fmt.Scanln(&N)

	numbers := make(map[int]int)

	for i := 0; i < N; i++ {
		var X int
		fmt.Scan(&X)
		numbers[X]++
	}

	// Obter os números distintos
	distinctNumbers := make([]int, 0, len(numbers))
	for num := range numbers {
		distinctNumbers = append(distinctNumbers, num)
	}

	// Ordenar os números distintos
	sort.Ints(distinctNumbers)

	// Imprimir a contagem de cada número distinto
	for _, num := range distinctNumbers {
		fmt.Printf("%d aparece %d vez(es)\n", num, numbers[num])
	}
}
*/
/*Questao 6
package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scanln(&n)

	for i := 0; i < n; i++ {
		var x float64
		fmt.Scanln(&x)
		prime := isPrime(x)
		printResult(prime)
	}
}

func printResult(prime bool) {
	if prime {
		fmt.Println("Prime")
	} else {
		fmt.Println("Not Prime")
	}
}

func isPrime(num float64) bool {
	if num < 2 {
		return false
	}
	if math.Mod(num, 2) == 0 {
		return num == 2
	}

	root := int(math.Sqrt(num))
	i := 3

	for i <= root {
		if math.Mod(num, float64(i)) == 0 {
			return false
		}
		i += 2
	}

	return true
}
*/
