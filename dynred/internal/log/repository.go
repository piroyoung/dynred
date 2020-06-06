package log

type Repository interface {
	Dump(log Log)
}
