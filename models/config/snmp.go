package config

// Printer or AP
type SnmpDevice struct {
	ID        string      `json:"id"`
	IP        string      `json:"ip"`
    SnmpConf  int8
	Snmp      Snmp        `json:"snmp" gorm:"foreignKey:SnmpConf;references:ID"`
}

type Snmp struct {
    ID        int8        `json:"id"`
	Version   int8        `json:"version"`
	Context   string      `json:"context"`
	User      string      `json:"user"`
	Pass      string      `json:"pass"`
	Privpass  string      `json:"privpass"`
}

type Oid struct {
    Name      string      `json:"name"`
    Oid       string      `json:"oid"`
    Value     string      `json:"value"`
}
