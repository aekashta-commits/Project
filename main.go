package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"todo-list/todo"
)

func main() {

	service := todo.NewTodoService()

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Добро пожаловать в TODO-лист! Введите команду help для получения списка доступных команд.")

	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}

		input := scanner.Text()
		parts := strings.Fields(input)
		if len(parts) == 0 {
			continue
		}

		command := parts[0]

		switch command {
		case "help":
			fmt.Println("Доступные команды:")
			fmt.Println("  help — список команд")
			fmt.Println("  add {заголовок} {текст задачи} — добавить задачу (заголовок из ОДНОГО слова)")
			fmt.Println("  list — показать все задачи")
			fmt.Println("  done {заголовок} — выполнить задачу")
			fmt.Println("  del {заголовок} — удалить задачу")
			fmt.Println("  events — показать историю событий")
			fmt.Println("  exit — выход")

			// Передаем исходную строку и nil, т.к. ошибок нет
			service.LogEvent(input, nil)

		case "add":
			if len(parts) < 3 {
				err := fmt.Errorf("использование: add {заголовок} {текст задачи}")
				fmt.Println(err)
				service.LogEvent(input, err)
				continue
			}

			title := parts[1]
			// Собираем все оставшиеся слова обратно в одну строку через пробел
			description := strings.Join(parts[2:], " ")

			err := service.AddTask(title, description)
			if err != nil {
				fmt.Println("Ошибка добавления:", err)
			} else {
				fmt.Printf("Задача '%s' успешно добавлена!\n", title)
			}
			service.LogEvent(input, err)

		case "list":
			tasks := service.ListTasks()
			if len(tasks) == 0 {
				fmt.Println("Список задач пуст.")
			} else {
				fmt.Println("Текущие задачи:")
				for _, task := range tasks {
					status := "❌"
					if task.IsDone {
						status = "✅"
					}
					fmt.Printf("  %s %s - %s\n", status, task.Title, task.Description)
				}
			}
			service.LogEvent(input, nil)

		case "done":
			if len(parts) < 2 {
				err := fmt.Errorf("использование: done {заголовок}")
				fmt.Println(err)
				service.LogEvent(input, err)
				continue
			}

			title := parts[1]
			err := service.DoneTask(title)
			if err != nil {
				fmt.Println("Ошибка выполнения:", err)
			} else {
				fmt.Printf("Задача '%s' отмечена как выполненная!\n", title)
			}
			service.LogEvent(input, err)

		case "del":
			if len(parts) < 2 {
				err := fmt.Errorf("использование: del {заголовок}")
				fmt.Println(err)
				service.LogEvent(input, err)
				continue
			}

			title := parts[1]
			err := service.DeleteTask(title)
			if err != nil {
				fmt.Println("Ошибка удаления:", err)
			} else {
				fmt.Printf("Задача '%s' успешно удалена.\n", title)
			}
			service.LogEvent(input, err)

		case "events":
			events := service.ListEvents()
			if len(events) == 0 {
				fmt.Println("История событий пуста.")
			} else {
				fmt.Println("История системных событий:")
				for i, ev := range events {
					status := "Успешно"
					if ev.ErrorText != "" {
						status = "Ошибка: " + ev.ErrorText
					}
					fmt.Printf("%d. [%s] Команда: '%s' -> %s\n",
						i+1, ev.CreatedAt.Format("15:04:05"), ev.Input, status)
				}
			}
			service.LogEvent(input, nil)

		case "exit":
			fmt.Println("Выход из программы. До свидания!")
			return

		default:
			err := fmt.Errorf("неизвестная команда: %s", command)
			fmt.Println(err)
			service.LogEvent(input, err)
		}
	}
}
