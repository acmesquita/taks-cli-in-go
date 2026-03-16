package adapter

type Request struct {
	Command string
	Options map[string]string
}

func NewRequest(command string, options map[string]string) *Request {
	return &Request{Command: command, Options: options}
}

func (r *Request) GetCommand() string {
	return r.Command
}

func (r *Request) GetOptions() map[string]string {
	return r.Options
}
