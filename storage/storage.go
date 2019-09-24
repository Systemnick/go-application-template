package storage

type IStorage interface {
	FindRecord(key string) (string, error)
	SaveRecord(key string, value string) error
	DeleteRecord(key string) (string, error)
}

type Storage struct {
	Dsn string
}

// func (s Storage) FindRecord(key string) (string, error) {
//
// 	return "", fmt.Errorf("not implemented yet")
// }
//
// func (s Storage) SaveRecord(key string, value string) error {
//
// 	return fmt.Errorf("not implemented yet")
// }

type Call struct {
	CallSid string
	From string
	To string
	Channel  string
	Timeout  int
	TraceId string
	Modify string
	Preserve string
	scenario *Scenario
}

type Scenario []Action

type Action struct {
	Command string
	Params  map[string]string
	Count   int
}
