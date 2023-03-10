package core

type Command struct {
	Id          int    `json:"id"`
	Name        string `json:"name" validate:"required"`
	Command     string `json:"command" validate:"required"`
	Description string `json:"description"`
}

type Store interface {
	Add(command Command) (*Command, error)
	Remove(id int) error
	Update(id int, command Command) error
	List() ([]Command, error)
	Get(id int) (*Command, error)
}
