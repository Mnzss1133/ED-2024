package main
import (
    
    "fmt"

	"sort"
   
)


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