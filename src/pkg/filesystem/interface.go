package filesystem

type FileSystem interface {
	Put()
	Get()
	Move()
	Delete()
}
