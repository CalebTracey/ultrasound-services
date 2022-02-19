package res

type UserAuthResponse struct {
	Success  bool
	Username string
	Message  Message
}

type Message struct {
	Count     int
	Status    string
	TimeTaken string
	ErrorLog  []ErrorLog
	HostName  string
}

type ErrorLog struct {
	RootCause  string
	Trace      string
	StatusCode string
}
