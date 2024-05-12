package main

import ("fmt")
/*Questao1
func main() {
    var distancia int
    fmt.Scanln(&distancia)

    tempo := distancia * 2 
    fmt.Printf("%d minutos\n", tempo)
}*/

/*Questao 2 
package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scanln(&N)

	for i := 0; i < N; i++ {
		var M int
		fmt.Scanln(&M)

		produtos := make(map[string]float64)
		for j := 0; j < M; j++ {
			var produto string
			var preco float64
			fmt.Scan(&produto, &preco)
			produtos[produto] = preco
		}

		var P int
		fmt.Scanln(&P)

		total := 0.0
		for k := 0; k < P; k++ {
			var produto string
			var quantidade int
			fmt.Scan(&produto, &quantidade)
			total += float64(quantidade) * produtos[produto]
		}

		fmt.Printf("R$ %.2f\n", total)
	}
}


*/

/*Questao 3 
package main

import "fmt"

func main() {
    var C, P, F int
    fmt.Scanln(&C, &P, &F)

    if C*F <= P {
        fmt.Println("S")
    } else {
        fmt.Println("N")
    }
}

*/

/*Questao 4

package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scanln(&N)

	sequence := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scanln(&sequence[i])
	}

	maxMarked := 0
	currentMarked := 0
	for i := 0; i < N; i++ {
		if sequence[i] == 1 {
			currentMarked++
			maxMarked = max(maxMarked, currentMarked)
		} else {
			currentMarked = 0
		}
	}

	fmt.Println(maxMarked)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}



*/