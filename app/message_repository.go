package app

type MessageRepository interface {
	Put(m Message) error
	FindById(id string) (*Message, error)
	FindAll() ([]Message, error)
}

type MessageRepositoryOnMemory struct {
	Data map[string]Message
}

func (r MessageRepositoryOnMemory) Put(m Message) error {
	r.Data[m.ID] = m
	return nil
}

func (r MessageRepositoryOnMemory) FindById(id string) (*Message, error) {
	m := r.Data[id]
	return &m, nil
}

func (r MessageRepositoryOnMemory) FindAll() ([]Message, error) {
	ms := []Message{}
	for _, m := range r.Data {
		ms = append(ms, m)
	}
	return ms, nil
}
