package main

/*Questao 1*/
/*

import (
    "fmt"
    "strconv"
)

func valorRealContrato(digitoFalho int, numeroContrato int) int {
    numeroContratoStr := strconv.Itoa(numeroContrato)
    numeroCorrigido := ""
    for _, d := range numeroContratoStr {
        if int(d-'0') != digitoFalho {
            numeroCorrigido += string(d)
        }
    }
    if numeroCorrigido == "" {
        return 0
    }
    numeroCorrigidoInt, _ := strconv.Atoi(numeroCorrigido)
    return numeroCorrigidoInt
}

func main() {
    var digitoFalho, numeroContrato int
    for {
        fmt.Scan(&digitoFalho, &numeroContrato)
        if digitoFalho == 0 && numeroContrato == 0 {
            break
        }
        fmt.Println(valorRealContrato(digitoFalho, numeroContrato))
    }
} */
/*Questao 2*/
/*

func contarLeds(numero string) int {
    ledsPorDigito := map[rune]int{
        '0': 6, '1': 2, '2': 5, '3': 5, '4': 4,
        '5': 5, '6': 6, '7': 3, '8': 7, '9': 6,
    }
    totalLeds := 0
    for _, digito := range numero {
        totalLeds += ledsPorDigito[digito]
    }
    return totalLeds
}

func main() {
    var n int
    fmt.Scan(&n)

    for i := 0; i < n; i++ {
        var valor string
        fmt.Scan(&valor)
        ledsNecessarios := contarLeds(valor)
        fmt.Printf("%d leds\n", ledsNecessarios)
    }
}*/

/*Questao3*/
/*

func main() {
	var n, v int
	cod := make([]byte, 50)

	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&cod)
		fmt.Scan(&v)

		for j := 0; j < 50; j++ {
			if cod[j] == 0 {
				break
			}
			if int(cod[j])-v < 'A' {
				fmt.Printf("%c", cod[j]-v+26)
			} else {
				fmt.Printf("%c", cod[j]-v)
			}
		}
		fmt.Println()
	}
}
*/

/*Questao4 */
/*package main

import (
	"fmt"
)

func main() {
	var name string

	for {
		_, err := fmt.Scan(&name)
		if err != nil {
			break // Termina o loop quando não há mais entrada
		}

		switch name {
		case "brasil", "portugal":
			fmt.Println("Feliz Natal!")
		case "inglaterra", "australia", "estados-unidos", "antardida", "canada":
			fmt.Println("Merry Christmas!")
		case "espanha", "argentina", "chile", "mexico":
			fmt.Println("Feliz Navidad!")
		case "marrocos", "siria":
			fmt.Println("Milad Mubarak!")
		case "italia", "libia":
			fmt.Println("Buon Natale!")
		case "alemanha":
			fmt.Println("Frohliche Weihnachten!")
		case "austria":
			fmt.Println("Frohe Weihnacht!")
		case "coreia":
			fmt.Println("Chuk Sung Tan!")
		case "grecia":
			fmt.Println("Kala Christougena!")
		case "suecia":
			fmt.Println("God Jul!")
		case "turquia":
			fmt.Println("Mutlu Noeller")
		case "belgica":
			fmt.Println("Zalig Kerstfeest!")
		case "japao":
			fmt.Println("Merii Kurisumasu!")
		case "irlanda":
			fmt.Println("Nollaig Shona Dhuit!")
		default:
			fmt.Println("--- NOT FOUND ---")
		}
	}
}
*/
