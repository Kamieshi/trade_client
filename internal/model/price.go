package model

import "time"

// Price struct company
type Price struct {
	CompanyID string    `protobuf:"bytes,1,opt,name=CompanyID,proto3" json:"CompanyID,omitempty" db:"id"`
	Name      string    `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty" db:"name"`
	Ask       uint32    `protobuf:"varint,2,opt,name=Ask,proto3" json:"Ask,omitempty"`
	Bid       uint32    `protobuf:"varint,3,opt,name=Bid,proto3" json:"Bid,omitempty"`
	Time      time.Time `protobuf:"bytes,4,opt,name=Time,proto3" json:"Time,omitempty"`
}
