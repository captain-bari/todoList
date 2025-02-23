package service

type database interface {
	CloseConnection() error
}

type processor interface{} // Impliment Close for cleanup

type restServer interface {
	Run() error
	Stop() error
}
