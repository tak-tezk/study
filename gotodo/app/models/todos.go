package models

import (
	"log"
	"time"
)

type Todo struct {
	ID        int
	Content   string
	UserID    int
	DueDate   *time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *User) CreateTodo(content string) (err error) {
	cmd := `insert into todos (
		content,
		user_id,
		due_date,
		created_at,
		updated_at) values (?, ?, ?, ?, ?)`

	_, err = Db.Exec(cmd, content, u.ID, nil, time.Now(), time.Now())
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func GetTodo(id int) (todo Todo, err error) {
	cmd := `select id, content, user_id, due_date, created_at, updated_at from todos
	where id = ?`
	todo = Todo{}

	err = Db.QueryRow(cmd, id).Scan(
		&todo.ID,
		&todo.Content,
		&todo.UserID,
		&todo.DueDate,
		&todo.CreatedAt,
		&todo.UpdatedAt)

	return todo, err
}

func GetTodos() (todos []Todo, err error) {
	cmd := `select id, content, user_id, due_date, created_at, updated_at  from todos`
	rows, err := Db.Query(cmd)
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		var todo Todo
		err = rows.Scan(&todo.ID,
			&todo.Content,
			&todo.UserID,
			&todo.DueDate,
			&todo.CreatedAt,
			&todo.UpdatedAt)
		if err != nil {
			log.Fatalln(err)
		}
		todos = append(todos, todo)
	}
	rows.Close()

	return todos, err
}

func (u *User) GetTodosByUser() (todos []Todo, err error) {
	cmd := `select id, content, user_id, due_date, created_at, updated_at from todos
	where user_id = ?`

	rows, err := Db.Query(cmd, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		var todo Todo
		err = rows.Scan(
			&todo.ID,
			&todo.Content,
			&todo.UserID,
			&todo.DueDate,
			&todo.CreatedAt,
			&todo.UpdatedAt)

		if err != nil {
			log.Fatalln(err)
		}
		todos = append(todos, todo)
	}
	rows.Close()

	return todos, err
}

func (t *Todo) UpdateTodo() error {
	cmd := `update todos set content = ?, user_id = ?, updated_at = ?
	where id = ?`
	_, err = Db.Exec(cmd, t.Content, t.UserID, time.Now(), t.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func (t *Todo) DeleteTodo() error {
	cmd := `delete from todos where id = ?`
	_, err = Db.Exec(cmd, t.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}
