package jina

import (
	"encoding/json"

	docarray "github.com/deepankarm/client-go/docarray"
	structpb "google.golang.org/protobuf/types/known/structpb"
)

// Custom JSON marshalling for DataRequestProto
func (x *DataRequestProto) MarshalJSON() ([]byte, error) {
	// JSON request to Jina Gateway
	type tmpDataRequestProto struct {
		Header     *HeaderProto                 `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`         // header contains meta info defined by the user
		Parameters *structpb.Struct             `protobuf:"bytes,2,opt,name=parameters,proto3" json:"parameters,omitempty"` // extra kwargs that will be used in executor
		Routes     []*RouteProto                `protobuf:"bytes,3,rep,name=routes,proto3" json:"routes,omitempty"`         // status info on every routes
		Data       *docarray.DocumentArrayProto `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`             // container for docs and groundtruths
	}
	tmp := tmpDataRequestProto{
		Header:     x.Header,
		Parameters: x.Parameters,
		Routes:     x.Routes,
		Data:       x.Data.Documents.(*DataRequestProto_DataContentProto_Docs).Docs,
	}
	return json.Marshal(tmp)
}

// Custom JSON unmarshalling for DataRequestProto
func (x *DataRequestProto) UnmarshalJSON(data []byte) error {
	// JSON response from Jina Gateway
	type tmpDataRequestProto struct {
		Header     *HeaderProto              `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`         // header contains meta info defined by the user
		Parameters *structpb.Struct          `protobuf:"bytes,2,opt,name=parameters,proto3" json:"parameters,omitempty"` // extra kwargs that will be used in executor
		Routes     []*RouteProto             `protobuf:"bytes,3,rep,name=routes,proto3" json:"routes,omitempty"`         // status info on every routes
		Data       []*docarray.DocumentProto `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`             // container for docs and groundtruths
	}

	var tmp tmpDataRequestProto
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	x.Header = tmp.Header
	x.Parameters = tmp.Parameters
	x.Routes = tmp.Routes
	x.Data = &DataRequestProto_DataContentProto{
		Documents: &DataRequestProto_DataContentProto_Docs{
			Docs: &docarray.DocumentArrayProto{
				Docs: tmp.Data,
			},
		},
	}
	return nil
}
