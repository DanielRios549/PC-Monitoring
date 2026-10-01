package config

type PC struct {
	ID        string      `json:"id"`
	IP        string      `json:"ip"`
	Name      string      `json:"name"`
}

type AP struct {
    ID        string      `json:"id"`
    IP        string      `json:"ip"`
	Snmp      Snmp        `json:"snmp"`
}
