//package main
//import ("fmt"
//"sort")


func main() {
	fmt.Println("Hello World!")
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
	var T int
	fmt.Scan(&T)

	for ; T > 0; T-- {
		var PA, PB int
		var G1, G2 float32
		fmt.Scan(&PA, &PB, &G1, &G2)

		years := 0
		for {
			PA += int(float32(PA) * G1 / 100)
			PB += int(float32(PB) * G2 / 100)
			years++
			if PA > PB {
				break
			}
		}

		if years > 100 {
			fmt.Println("Mais de 1 seculo")
		} else {
			fmt.Printf("%d anos\n", years)
		}
	}
}