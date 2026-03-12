package main

import (
	"fmt"
	"math/rand"
	"mathcore/domain"
	"time"
)

const (
	totalQuestions    = 5
	pointsPerQuestion = 20
)

var id uint64 = 1

func menu() {

}

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

	var users []domain.User

	for {
		menu()

		choice := ""
		fmt.Scan(&choice)

		switch choice {
		case "1":
			u := play()
			users = append(users, u)
		case "2":
			for _, u := range users {
				fmt.Printf("Id: %v Name: %s Time: %v", u.Id, u.Name, u.Time)

			}
		case "3":
			return
		default:
		}
	}

}

func play() domain.User {

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

	fmt.Println("Введіть ім'я: ")
	name := ""

	fmt.Scan(&name)

	user := domain.User{
		Id:   id,
		Name: name,
		Time: elapsedTime,
	}
	id++

	return user
}
