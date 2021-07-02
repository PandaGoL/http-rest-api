package apiserver

type APIServer struct {
	config *Config
}

//Новый объект (тело запроса)
func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
	}
}

//функция(метод)  Start, которя запускает БД и HTTP сервер
func (s *APIServer) Start() error {
	return nil
}
