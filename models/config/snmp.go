package config

type Printers struct {
	ID        string      `json:"id"`
	IP        string      `json:"ip"`
	Snmp      Snmp        `json:"snmp"`
}

 type Snmp struct {
	Version   int8        `json:"version"`
	Context   string      `json:"context"`
	User      string      `json:"user"`
	Pass      string      `json:"pass"`
	Privpass  string      `json:"privpass"`
}
