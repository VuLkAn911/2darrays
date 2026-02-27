package main

import "fmt"

func main() {
	//sub 1
	numbers := [3][5]int{}
	for i, row := range numbers {
		for j, _ := range row {
			fmt.Scan(&numbers[i][j])
		}
		fmt.Println(numbers)
		max_i, max_j := 0, 0
		max_value := numbers[0][0]

		for i, row := range numbers {
			for j, number := range row {
				if number > max_value {
					max_value = number
					max_i, max_j = i, j
				}
			}

		}
		fmt.Printf("max value: %d, position: [%d][%d]\n", max_value, max_i, max_j)
		// sub 2
		numbers2 := [11][11]string{}
		for i := range numbers2 {
			for j := range numbers2[i] {
				numbers2[i][j] = "."
			}
		}

		for i, row := range numbers2 {
			for j, _ := range row {
				if i == j {
					numbers2[i][j] = "*"
				}

				if i+j == len(numbers2)-1 {
					numbers2[i][j] = "*"
				}

				if i == len(numbers2)/2 || j == len(numbers2)/2 {
					numbers2[i][j] = "*"
				}

				fmt.Print(numbers2[i][j], "")
			}
			fmt.Println()
		}
		// sub 3
		numbers3 := [8][8]string{}
		for i := range numbers3 {
			for j := range numbers3[i] {

				if (i+j)%2 == 0 {
					numbers3[i][j] = "."
				} else {
					numbers3[i][j] = "*"
				}
				fmt.Print(numbers3[i][j], "")
			}
			fmt.Println()
		}

		// sub 4

		numbers4 := [4][4]int{}
		for i, row := range numbers4 {
			for j, _ := range row {
				fmt.Scan(&numbers4[i][j])
			}
		}
		for i, pair := range numbers4 {
			for j, _ := range pair {
				fmt.Print(numbers4[i][j], " ")
			}
			fmt.Println()
		}
		var num int
		var num1 int
		fmt.Scan(&num, &num1)
		buffer := numbers4[num-1]
		numbers4[num-1] = numbers4[num1-1]
		numbers4[num1-1] = buffer
		for i, pair := range numbers4 {
			for j, _ := range pair {
				fmt.Print(numbers4[i][j], " ")
			}
			fmt.Println()
		}

		// 		//sub 5

		numbers5 := [4][4]int{}
		for i, row := range numbers5 {
			for j, _ := range row {
				fmt.Scan(&numbers5[i][j])
			}
		}
		for i, pair := range numbers5 {
			for j, _ := range pair {
				fmt.Print(numbers5[i][j], " ")
			}
			fmt.Println()
		}
		var num2 int
		var num3 int
		fmt.Scan(&num2, &num3)

		for i, row := range numbers5 {
			buffer := row[num2-1]
			numbers5[i][num2-1] = row[num3-1]
			numbers5[i][num3-1] = buffer
		}
		for i, pair := range numbers5 {
			for j, _ := range pair {
				fmt.Print(numbers5[i][j], " ")
			}
			fmt.Println()
		}
	}
}
