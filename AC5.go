package main
import (
    
    "fmt"

	"sort"
   
)


//Questao 1
func main(){
	fmt.Println("Hello World!")
}
//Questao 2


func main() {
    var A, B, C float64
    fmt.Println("A:")
    fmt.Scanln(&A)
    fmt.Println("B:")
    fmt.Scanln(&B)
    fmt.Println("C:")
    fmt.Scanln(&C)

	sides := []float64{A, B, C}
    sort.Slice(sides, func(i, j int) bool {
        return sides[i] > sides[j]
    })
    A, B, C = sides[0], sides[1], sides[2]

    // Conditions to determine the type of triangle
    if A >= B+C {
       fmt.Println("NAO FORMA TRIANGULO") 
    } else if A*A == B*B+C*C {
		fmt.Println ("TRIANGULO RETANGULO")
    } else if A*A > B*B+C*C {
		fmt.Println ("TRIANGULO OBTUSANGULO")
		fmt.Println( "TRIANGULO ISOSCELES")
	} else if A == B && B == C {
		fmt.Println ("TRIANGULO ACUTANGULO")
		fmt.Println ("TRIANGULO EQUILATERO")	
    } else if A*A < B*B+C*C {
        fmt.Println ("TRIANGULO ACUTANGULO")
		fmt.Println( "TRIANGULO ISOSCELES")
  
    } else if A == B || A == C || B == C {
		fmt.Println( "TRIANGULO ISOSCELES")
    }
    
}


//Questao 3
func solvePopulationProblem() {
	var T int
	fmt.Scanln(&T)

	for i := 0; i < T; i++ {
		var PA, PB, G1, G2 float64
		fmt.Scanf("%f %f %f %f", &PA, &PB, &G1, &G2)

		if G2 == 0 && PA >= PB {
			fmt.Println("Mais de 1 seculo")
			continue
		}

		years := 0
		for PA <= PB {
			PA = PA * (1 + G1 / 100)
			PB = PB * (1 + G2 / 100)
			years++
			if years > 100 {
				fmt.Println("Mais de 1 seculo")
				break
			}
		}
		if years <= 100 {
			fmt.Println(years, "anos")
		}
	}
}

func main() {
	solvePopulationProblem()
}
