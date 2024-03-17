

import "fmt"
//Questao 1
func Tower(height int, fromPole, toPole, withPole string) {
	if height >= 1 {
		Tower(height-1, fromPole, withPole, toPole)
		fmt.Printf("Move disk from %s to %s\n", fromPole, toPole)
		Tower(height-1, withPole, toPole, fromPole)
	}
}

func main() {
	numDisks := 3
	Tower(numDisks, "A", "C", "B")
}
//Questao 2
func findPair(arr []int, target int) (int, int) {
    left := 0
    right := len(arr) - 1

    for left < right {
        sum := arr[left] + arr[right]
        if sum == target {
            return arr[left], arr[right]
        } else if sum < target {
            left++
        } else {
            right--
        }
    }

    return -1, -1
}

func main() {
    arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
    target := 10
    num1, num2 := findPair(arr, target)
    if num1 == -1 && num2 == -1 {
        fmt.Println("Nenhum par encontrado.")
    } else {
        fmt.Printf("Par encontrado: %d e %d\n", num1, num2)
    }
}