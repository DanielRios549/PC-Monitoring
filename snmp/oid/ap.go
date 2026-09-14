package oid

// SNMP V1 needs the exact OIDs
// since it has no GetBulk() support
// Key 0 == Exact (V1)
// Key 1 == Bulk  (V2/3)

// TODO: Separate V1 and V2/3 OIDs
var APOptions = map[string][]string{
    "hostname": {
        ".2.1.1.5.0",
        ".2.1.1.4.0",
    },
}
