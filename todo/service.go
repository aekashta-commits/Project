package todo

import (
	"errors"
	"time"
)

// управляем событиями в памяти
type TodoService struct {
	tasks  map[string]*Task
	events []Event
}

// конструктор для новых сервисов
func NewTodoService() *TodoService {
	return &TodoService{
		tasks:  make(map[string]*Task),
		events: make([]Event, 0),
	}
}

// метод для создания задачи
func (s *TodoService) AddTask(title string, description string) error {
	if _, exists := s.tasks[title]; exists {
		return errors.New("задача с таким заголовком уже существует")
	}

	newTask := &Task{
		Title:       title,
		Description: description,
		CreatedAt:   time.Now(),
		IsDone:      false,
		DoneAt:      nil, // до тех пор пока не выполнена
	}

	s.tasks[title] = newTask
	return nil
}

// метод для возврата всех текущих задач
func (s *TodoService) ListTasks() map[string]*Task {
	return s.tasks
}

// метод для отметки выполненных задач по заголовку
func (s *TodoService) DoneTask(title string) error {
	task, exists := s.tasks[title]
	if !exists {
		return errors.New("задача не найдена")
	}

	if task.IsDone {
		return errors.New("задача уже выполнена")
	}

	now := time.Now()
	task.IsDone = true
	task.DoneAt = &now

	return nil
}

// удачение задач по заголовку
func (s *TodoService) DeleteTask(title string) error {
	_, exists := s.tasks[title]
	if !exists {
		return errors.New("задача не найдена")
	}

	delete(s.tasks, title)
	return nil
}

// метод для записи вводенных задач и результата
func (s *TodoService) LogEvent(input string, err error) {
	errText := ""
	if err != nil {
		errText = err.Error() // сохранили если есть ошибка
	}

	event := Event{
		Input:     input,
		ErrorText: errText,
		CreatedAt: time.Now(),
	}

	s.events = append(s.events, event)
}

// метод для возврата списка всех произошедших событий
func (s *TodoService) ListEvents() []Event {
	return s.events
}
