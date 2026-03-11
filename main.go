package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	totalQuestions    = 5
	pointsPerQuestion = 20
)

func countdown() {
	fmt.Println("Гра почнеться через:")
	for i := 5; i > 0; i-- {
		fmt.Println(i)
		time.Sleep(1 * time.Second)
	}
	fmt.Println("GO!")
	fmt.Println()
}

func getValidIntInput() int {
	var num int

	for {
		_, err := fmt.Scanln(&num)
		if err != nil {
			fmt.Println("Введіть коректне число!")

			var discard string
			fmt.Scanln(&discard)

			continue
		}
		return num
	}
}

func main() {

	fmt.Println("Вітаємо у математичній грі!")
	fmt.Println("Вам потрібно розв'язати 5 прикладів.")
	fmt.Println("За кожну правильну відповідь ви отримуєте 20 балів.")
	fmt.Println("Максимум: 100 балів.")
	fmt.Println()

	countdown()

	score := 0
	startTime := time.Now()

	for i := 1; i <= totalQuestions; i++ {
		num1 := rand.Intn(90) + 10
		num2 := rand.Intn(90) + 10

		correctAnswer := num1 + num2

		fmt.Printf("Питання %d: %d + %d = ", i, num1, num2)
		userAnswer := getValidIntInput()

		if userAnswer == correctAnswer {
			fmt.Println("Правильно!")
			score += pointsPerQuestion
		} else {
			fmt.Printf("Неправильно! Правильна відповідь: %d\n", correctAnswer)
		}

		fmt.Println()
	}

	elapsedTime := time.Since(startTime)

	fmt.Println("===== Результат =====")
	fmt.Printf("Ваші бали: %d з 100\n", score)
	fmt.Printf("Час проходження: %.2f секунд\n", elapsedTime.Seconds())

	if score == 100 {
		fmt.Println("Відмінний результат!")
	} else if score >= 60 {
		fmt.Println("Гарна робота!")
	} else {
		fmt.Println("Спробуйте ще раз!")
	}

	time.Sleep(10 * time.Second)
}
