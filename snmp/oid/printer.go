package oid

// SNMP V1 needs the exact OIDs
// since it has no GetBulk() support
// Key 0 == Exact (V1)
// Key 1 == Bulk  (V2/3)
var PrinterOptions = map[string][]string{
    "hostname": {
        ".2.1.1.5.0",
        ".2.1.1.4.0",
    },
    "vendor_get": {
        ".2.1.1.2.0", // TODO: Verify if SNMPv1 version is correct
        ".2.1.1.2.0",
    },
    "printer_model": {
        ".2.1.25.3.2.1.3.1",
        ".2.1.25.3.2.1.3",
    },
    "toner_current": {
        ".2.1.43.11.1.1.9.1.1",
        ".2.1.43.11.1.1.9",
    },
    "toner_max": {
        ".2.1.43.11.1.1.8.1.1",
        ".2.1.43.11.1.1.8",
    },
}
