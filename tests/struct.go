//go:generate gobatis struct.go
package tests

import (
	"net"
	"time"
)

type TestA1 struct {
	TableName struct{}         `db:"gobatis_testa"`
	ID        int64            `db:"id,pk,autoincr"`
	Field0    bool             `db:"field0,null"`
	Field1    int              `db:"field1,null"`
	Field2    uint             `db:"field2,null"`
	Field3    float32          `db:"field3,null"`
	Field4    float64          `db:"field4,null"`
	Field5    string           `db:"field5,null"`
	Field6    time.Time        `db:"field6,null"`
	Field7    net.IP           `db:"field7,null"`
	Field8    net.HardwareAddr `db:"field8,null"`
}

type TestA2 struct {
	TableName struct{}          `db:"gobatis_testa"`
	ID        int64             `db:"id,pk,autoincr"`
	Field0    *bool             `db:"field0,null"`
	Field1    *int              `db:"field1,null"`
	Field2    *uint             `db:"field2,null"`
	Field3    *float32          `db:"field3,null"`
	Field4    *float64          `db:"field4,null"`
	Field5    *string           `db:"field5,null"`
	Field6    *time.Time        `db:"field6,null"`
	Field7    *net.IP           `db:"field7,null"`
	Field8    *net.HardwareAddr `db:"field8,null"`
}

type TestA3 struct {
	TableName struct{}         `db:"gobatis_testa"`
	ID        int64            `db:"id,pk,autoincr"`
	Field0    bool             `db:"field0"`
	Field1    int              `db:"field1"`
	Field2    uint             `db:"field2"`
	Field3    float32          `db:"field3"`
	Field4    float64          `db:"field4"`
	Field5    string           `db:"field5"`
	Field6    time.Time        `db:"field6"`
	Field7    net.IP           `db:"field7"`
	Field8    net.HardwareAddr `db:"field8"`
}

type TestA4 struct {
	TableName struct{}          `db:"gobatis_testa"`
	ID        int64             `db:"id,pk,autoincr"`
	Field0    *bool             `db:"field0"`
	Field1    *int              `db:"field1"`
	Field2    *uint             `db:"field2"`
	Field3    *float32          `db:"field3"`
	Field4    *float64          `db:"field4"`
	Field5    *string           `db:"field5"`
	Field6    *time.Time        `db:"field6"`
	Field7    *net.IP           `db:"field7"`
	Field8    *net.HardwareAddr `db:"field8"`
}

type TestB1 struct {
	TableName struct{}         `db:"gobatis_testb"`
	ID        int64            `db:"id,pk,autoincr"`
	Field0    bool             `db:"field0,notnull"`
	Field1    int              `db:"field1,notnull"`
	Field2    uint             `db:"field2,notnull"`
	Field3    float32          `db:"field3,notnull"`
	Field4    float64          `db:"field4,notnull"`
	Field5    string           `db:"field5,notnull"`
	Field6    time.Time        `db:"field6,notnull"`
	Field7    net.IP           `db:"field7,notnull"`
	Field8    net.HardwareAddr `db:"field8,notnull"`
}

type TestB2 struct {
	TableName struct{}          `db:"gobatis_testb"`
	ID        int64             `db:"id,pk,autoincr"`
	Field0    *bool             `db:"field0,notnull"`
	Field1    *int              `db:"field1,notnull"`
	Field2    *uint             `db:"field2,notnull"`
	Field3    *float32          `db:"field3,notnull"`
	Field4    *float64          `db:"field4,notnull"`
	Field5    *string           `db:"field5,notnull"`
	Field6    *time.Time        `db:"field6,notnull"`
	Field7    *net.IP           `db:"field7,notnull"`
	Field8    *net.HardwareAddr `db:"field8,notnull"`
}

var (
	TestMAC net.HardwareAddr
	TestIP  net.IP
)

func init() {
	TestMAC, _ = net.ParseMAC("01:02:03:04:A5:A6")
	TestIP = net.ParseIP("192.168.1.1")
}

type ITest interface {
	InsertA1(v *TestA1) (int64, error)
	InsertA2(v *TestA2) (int64, error)
	InsertA3(v *TestA3) (int64, error)
	InsertA4(v *TestA4) (int64, error)
	InsertB1(v *TestB1) (int64, error)
	InsertB2(v *TestB2) (int64, error)
}