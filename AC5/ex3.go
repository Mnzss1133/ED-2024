package main 
import ("fmt")
func main() {
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