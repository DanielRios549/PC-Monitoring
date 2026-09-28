package modules

type Agent struct {
    server string
}

func NewAgent() *Agent {
    instance := &Agent{
        server: "localhost:9003",
    }

    return instance
}

func (a *Agent) Start() {
    println("Starting Agent...")
}
