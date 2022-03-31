package api

import (
	"errors"
	"sort"
	"time"

	"github.com/google/uuid"
)

var ErrNotFound = errors.New("not found")

type Store interface {
	Create(description string) todo
	Check(id string, completed bool) (todo, error)
	Delete(id string)
	List() []todo
}

type store struct {
	todos map[string]*todo
}

type todo struct {
	// UUID, equal to the key in the todos map
	ID string `json:"id"`
	// The text of the Todo
	Description string `json:"description"`
	// Boolean of whether the Todo is completed
	Completed bool `json:"completed"`
	// Unix timestamp of creation
	CreatedAt int64 `json:"createdAt"`
}

func newStore() Store {
	s := &store{
		todos: make(map[string]*todo),
	}
	s.Seed()
	return s
}

func (s *store) Seed() {
	id1 := uuid.NewString()
	s.todos[id1] = &todo{
		ID:          id1,
		Description: "Pick up dry cleaning",
		Completed:   true,
		CreatedAt:   1,
	}

	id2 := uuid.NewString()
	s.todos[id2] = &todo{
		ID:          id2,
		Description: "Grab coffee",
		Completed:   false,
		CreatedAt:   2,
	}

	id3 := uuid.NewString()
	s.todos[id3] = &todo{
		ID:          id3,
		Description: "Solve world hunger",
		Completed:   false,
		CreatedAt:   3,
	}
}

func (s *store) Create(description string) todo {
	id := uuid.NewString()
	new := &todo{
		ID:          id,
		Description: description,
		CreatedAt:   time.Now().Unix(),
	}
	s.todos[id] = new
	return *new
}

func (s *store) Check(id string, completed bool) (todo, error) {
	var checked *todo = s.todos[id]

	if checked == nil {
		return *checked, ErrNotFound
	}

	checked.Completed = completed

	return *checked, nil
}

func (s *store) Delete(id string) {
	delete(s.todos, id)
}

func (s *store) List() []todo {
	todos := make([]todo, 0)

	for _, t := range s.todos {
		todos = append(todos, *t)
	}

	sort.Slice(todos, func(i, j int) bool {
		return todos[i].CreatedAt < todos[j].CreatedAt
	})

	return todos
}
