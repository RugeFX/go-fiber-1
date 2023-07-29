package repositories

type BaseRepository[m interface{}] interface {
	GetAll() (m, error)
	Save(m) (m, error)
	Delete(id interface{}) (m, error)
	Update(m) (m, error)
}
