package docarray

import (
	"google.golang.org/protobuf/encoding/protojson"
)

// Custom JSON marshalling for DocumentProto
func (x *DocumentProto) MarshalJSON() ([]byte, error) {
	return protojson.Marshal(x)
}

// Custom JSON unmarshalling for DocumentProto
func (x *DocumentProto) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, x)
}
