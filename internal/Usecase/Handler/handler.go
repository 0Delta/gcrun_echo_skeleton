package usecase_handler

type Handler interface {
	Run(interface{}) (string, error)
}
