package config

type PCs struct {
	ID        string      `json:"id"`
	IP        string      `json:"ip"`
	Name      string      `json:"name"`
}

type APs struct {
    ID        string      `json:"id"`
    IP        string      `json:"ip"`
	Snmp      Snmp        `json:"snmp"`
}
