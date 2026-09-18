package snmp

import (
	"errors"
	"fmt"
	// "pc-monitoring/models"
	"pc-monitoring/models/config"
	"pc-monitoring/snmp/oid"
	"time"

	g "github.com/gosnmp/gosnmp"
)

func V3(ip, context, user, pass, privpass string, isAP bool) ([]*config.Oid, error) {
	params := &g.GoSNMP{
		Target:        ip,
		Port:          161,
		Version:       g.Version3,
		SecurityModel: g.UserSecurityModel,
		MsgFlags:      g.AuthPriv,
		Timeout:       time.Duration(30) * time.Second,
		ContextName:  context,
		SecurityParameters: &g.UsmSecurityParameters{
			UserName: user,
			AuthenticationProtocol:   g.MD5,
			AuthenticationPassphrase: pass,
			PrivacyProtocol:          g.DES,
			PrivacyPassphrase:        privpass,
		},
	}

	err := params.Connect()

	if err != nil {
		// fmt.Printf("Connect() error: %v", err)
        return nil, errors.New("printer is offline")
	}

	defer func() {
		err := params.Conn.Close()

		if err != nil {
			fmt.Printf("Cannot Close Connection: %v", err)
		}
	}()

    oids := oid.PrinterOptions

    if isAP {
        oids = oid.APOptions
    }

	info := GetInfo(params, oids)

    return info, nil
}
