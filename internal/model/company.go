package model

import "time"

// Company struct company
type Company struct {
	ID   string    `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty" db:"id"`
	Name string    `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty" db:"name"`
	Ask  uint32    `protobuf:"varint,2,opt,name=Ask,proto3" json:"Ask,omitempty"`
	Bid  uint32    `protobuf:"varint,3,opt,name=Bid,proto3" json:"Bid,omitempty"`
	Time time.Time `protobuf:"bytes,4,opt,name=Time,proto3" json:"Time,omitempty"`
}
