// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.7
// source: docarray.proto

package docarray

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

//*
// Represents a (quantized) dense n-dim array
type DenseNdArrayProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the actual array data, in bytes
	Buffer []byte `protobuf:"bytes,1,opt,name=buffer,proto3" json:"buffer,omitempty"`
	// the shape (dimensions) of the array
	Shape []uint32 `protobuf:"varint,2,rep,packed,name=shape,proto3" json:"shape,omitempty"`
	// the data type of the array
	Dtype string `protobuf:"bytes,3,opt,name=dtype,proto3" json:"dtype,omitempty"`
}

func (x *DenseNdArrayProto) Reset() {
	*x = DenseNdArrayProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_docarray_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DenseNdArrayProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DenseNdArrayProto) ProtoMessage() {}

func (x *DenseNdArrayProto) ProtoReflect() protoreflect.Message {
	mi := &file_docarray_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DenseNdArrayProto.ProtoReflect.Descriptor instead.
func (*DenseNdArrayProto) Descriptor() ([]byte, []int) {
	return file_docarray_proto_rawDescGZIP(), []int{0}
}

func (x *DenseNdArrayProto) GetBuffer() []byte {
	if x != nil {
		return x.Buffer
	}
	return nil
}

func (x *DenseNdArrayProto) GetShape() []uint32 {
	if x != nil {
		return x.Shape
	}
	return nil
}

func (x *DenseNdArrayProto) GetDtype() string {
	if x != nil {
		return x.Dtype
	}
	return ""
}

//*
// Represents a general n-dim array, can be either dense or sparse
type NdArrayProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Content:
	//	*NdArrayProto_Dense
	//	*NdArrayProto_Sparse
	Content isNdArrayProto_Content `protobuf_oneof:"content"`
	// the name of the ndarray class
	ClsName    string           `protobuf:"bytes,3,opt,name=cls_name,json=clsName,proto3" json:"cls_name,omitempty"`
	Parameters *structpb.Struct `protobuf:"bytes,4,opt,name=parameters,proto3" json:"parameters,omitempty"`
}

func (x *NdArrayProto) Reset() {
	*x = NdArrayProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_docarray_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NdArrayProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NdArrayProto) ProtoMessage() {}

func (x *NdArrayProto) ProtoReflect() protoreflect.Message {
	mi := &file_docarray_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NdArrayProto.ProtoReflect.Descriptor instead.
func (*NdArrayProto) Descriptor() ([]byte, []int) {
	return file_docarray_proto_rawDescGZIP(), []int{1}
}

func (m *NdArrayProto) GetContent() isNdArrayProto_Content {
	if m != nil {
		return m.Content
	}
	return nil
}

func (x *NdArrayProto) GetDense() *DenseNdArrayProto {
	if x, ok := x.GetContent().(*NdArrayProto_Dense); ok {
		return x.Dense
	}
	return nil
}

func (x *NdArrayProto) GetSparse() *SparseNdArrayProto {
	if x, ok := x.GetContent().(*NdArrayProto_Sparse); ok {
		return x.Sparse
	}
	return nil
}

func (x *NdArrayProto) GetClsName() string {
	if x != nil {
		return x.ClsName
	}
	return ""
}

func (x *NdArrayProto) GetParameters() *structpb.Struct {
	if x != nil {
		return x.Parameters
	}
	return nil
}

type isNdArrayProto_Content interface {
	isNdArrayProto_Content()
}

type NdArrayProto_Dense struct {
	Dense *DenseNdArrayProto `protobuf:"bytes,1,opt,name=dense,proto3,oneof"` // dense representation of the ndarray
}

type NdArrayProto_Sparse struct {
	Sparse *SparseNdArrayProto `protobuf:"bytes,2,opt,name=sparse,proto3,oneof"` // sparse representation of the ndarray
}

func (*NdArrayProto_Dense) isNdArrayProto_Content() {}

func (*NdArrayProto_Sparse) isNdArrayProto_Content() {}

//*
// Represents a sparse ndarray
type SparseNdArrayProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A 2-D int64 tensor of shape [N, ndims], which specifies the indices of the elements in the sparse tensor that contain nonzero values (elements are zero-indexed)
	Indices *DenseNdArrayProto `protobuf:"bytes,1,opt,name=indices,proto3" json:"indices,omitempty"`
	// A 1-D tensor of any type and shape [N], which supplies the values for each element in indices.
	Values *DenseNdArrayProto `protobuf:"bytes,2,opt,name=values,proto3" json:"values,omitempty"`
	// A 1-D int64 tensor of shape [ndims], which specifies the shape of the sparse tensor.
	Shape []uint32 `protobuf:"varint,3,rep,packed,name=shape,proto3" json:"shape,omitempty"`
}

func (x *SparseNdArrayProto) Reset() {
	*x = SparseNdArrayProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_docarray_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SparseNdArrayProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SparseNdArrayProto) ProtoMessage() {}

func (x *SparseNdArrayProto) ProtoReflect() protoreflect.Message {
	mi := &file_docarray_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SparseNdArrayProto.ProtoReflect.Descriptor instead.
func (*SparseNdArrayProto) Descriptor() ([]byte, []int) {
	return file_docarray_proto_rawDescGZIP(), []int{2}
}

func (x *SparseNdArrayProto) GetIndices() *DenseNdArrayProto {
	if x != nil {
		return x.Indices
	}
	return nil
}

func (x *SparseNdArrayProto) GetValues() *DenseNdArrayProto {
	if x != nil {
		return x.Values
	}
	return nil
}

func (x *SparseNdArrayProto) GetShape() []uint32 {
	if x != nil {
		return x.Shape
	}
	return nil
}

//*
// Represents the relevance model to `ref_id`
type NamedScoreProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value       float32 `protobuf:"fixed32,1,opt,name=value,proto3" json:"value,omitempty"`               // value
	OpName      string  `protobuf:"bytes,2,opt,name=op_name,json=opName,proto3" json:"op_name,omitempty"` // the name of the operator/score function
	Description string  `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`     // text description of the score
	RefId       string  `protobuf:"bytes,4,opt,name=ref_id,json=refId,proto3" json:"ref_id,omitempty"`    // the score is computed between doc `id` and `ref_id`
}

func (x *NamedScoreProto) Reset() {
	*x = NamedScoreProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_docarray_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NamedScoreProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NamedScoreProto) ProtoMessage() {}

func (x *NamedScoreProto) ProtoReflect() protoreflect.Message {
	mi := &file_docarray_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NamedScoreProto.ProtoReflect.Descriptor instead.
func (*NamedScoreProto) Descriptor() ([]byte, []int) {
	return file_docarray_proto_rawDescGZIP(), []int{3}
}

func (x *NamedScoreProto) GetValue() float32 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *NamedScoreProto) GetOpName() string {
	if x != nil {
		return x.OpName
	}
	return ""
}

func (x *NamedScoreProto) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *NamedScoreProto) GetRefId() string {
	if x != nil {
		return x.RefId
	}
	return ""
}

//*
// Represents a Document
type DocumentProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A hexdigest that represents a unique document ID
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Types that are assignable to Content:
	//	*DocumentProto_Blob
	//	*DocumentProto_Tensor
	//	*DocumentProto_Text
	Content isDocumentProto_Content `protobuf_oneof:"content"`
	// the depth of the recursive chunk structure
	Granularity uint32 `protobuf:"varint,5,opt,name=granularity,proto3" json:"granularity,omitempty"`
	// the width of the recursive match structure
	Adjacency uint32 `protobuf:"varint,6,opt,name=adjacency,proto3" json:"adjacency,omitempty"`
	// the parent id from the previous granularity
	ParentId string `protobuf:"bytes,7,opt,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty"`
	// The weight of this document
	Weight float32 `protobuf:"fixed32,8,opt,name=weight,proto3" json:"weight,omitempty"`
	// a uri of the document could be: a local file path, a remote url starts with http or https or data URI scheme
	Uri string `protobuf:"bytes,9,opt,name=uri,proto3" json:"uri,omitempty"`
	// modality, an identifier to the modality this document belongs to. In the scope of multi/cross modal search
	Modality string `protobuf:"bytes,10,opt,name=modality,proto3" json:"modality,omitempty"`
	// mime type of this document, for buffer content, this is required; for other contents, this can be guessed
	MimeType string `protobuf:"bytes,11,opt,name=mime_type,json=mimeType,proto3" json:"mime_type,omitempty"`
	// the offset of the doc
	Offset float32 `protobuf:"fixed32,12,opt,name=offset,proto3" json:"offset,omitempty"`
	// the position of the doc, could be start and end index of a string; could be x,y (top, left) coordinate of an image crop; could be timestamp of an audio clip
	Location []float32 `protobuf:"fixed32,13,rep,packed,name=location,proto3" json:"location,omitempty"`
	// list of the sub-documents of this document (recursive structure)
	Chunks []*DocumentProto `protobuf:"bytes,14,rep,name=chunks,proto3" json:"chunks,omitempty"`
	// the matched documents on the same level (recursive structure)
	Matches []*DocumentProto `protobuf:"bytes,15,rep,name=matches,proto3" json:"matches,omitempty"`
	// the embedding of this document
	Embedding *NdArrayProto `protobuf:"bytes,16,opt,name=embedding,proto3" json:"embedding,omitempty"`
	// a structured data value, consisting of field which map to dynamically typed values.
	Tags *structpb.Struct `protobuf:"bytes,17,opt,name=tags,proto3" json:"tags,omitempty"`
	// Scores performed on the document, each element corresponds to a metric
	Scores map[string]*NamedScoreProto `protobuf:"bytes,18,rep,name=scores,proto3" json:"scores,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Evaluations performed on the document, each element corresponds to a metric
	Evaluations map[string]*NamedScoreProto `protobuf:"bytes,19,rep,name=evaluations,proto3" json:"evaluations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// system-defined meta attributes represented in a structured data value.
	XMetadata *structpb.Struct `protobuf:"bytes,20,opt,name=_metadata,json=Metadata,proto3" json:"_metadata,omitempty"`
}

func (x *DocumentProto) Reset() {
	*x = DocumentProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_docarray_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DocumentProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DocumentProto) ProtoMessage() {}

func (x *DocumentProto) ProtoReflect() protoreflect.Message {
	mi := &file_docarray_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DocumentProto.ProtoReflect.Descriptor instead.
func (*DocumentProto) Descriptor() ([]byte, []int) {
	return file_docarray_proto_rawDescGZIP(), []int{4}
}

func (x *DocumentProto) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (m *DocumentProto) GetContent() isDocumentProto_Content {
	if m != nil {
		return m.Content
	}
	return nil
}

func (x *DocumentProto) GetBlob() []byte {
	if x, ok := x.GetContent().(*DocumentProto_Blob); ok {
		return x.Blob
	}
	return nil
}

func (x *DocumentProto) GetTensor() *NdArrayProto {
	if x, ok := x.GetContent().(*DocumentProto_Tensor); ok {
		return x.Tensor
	}
	return nil
}

func (x *DocumentProto) GetText() string {
	if x, ok := x.GetContent().(*DocumentProto_Text); ok {
		return x.Text
	}
	return ""
}

func (x *DocumentProto) GetGranularity() uint32 {
	if x != nil {
		return x.Granularity
	}
	return 0
}

func (x *DocumentProto) GetAdjacency() uint32 {
	if x != nil {
		return x.Adjacency
	}
	return 0
}

func (x *DocumentProto) GetParentId() string {
	if x != nil {
		return x.ParentId
	}
	return ""
}

func (x *DocumentProto) GetWeight() float32 {
	if x != nil {
		return x.Weight
	}
	return 0
}

func (x *DocumentProto) GetUri() string {
	if x != nil {
		return x.Uri
	}
	return ""
}

func (x *DocumentProto) GetModality() string {
	if x != nil {
		return x.Modality
	}
	return ""
}

func (x *DocumentProto) GetMimeType() string {
	if x != nil {
		return x.MimeType
	}
	return ""
}

func (x *DocumentProto) GetOffset() float32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *DocumentProto) GetLocation() []float32 {
	if x != nil {
		return x.Location
	}
	return nil
}

func (x *DocumentProto) GetChunks() []*DocumentProto {
	if x != nil {
		return x.Chunks
	}
	return nil
}

func (x *DocumentProto) GetMatches() []*DocumentProto {
	if x != nil {
		return x.Matches
	}
	return nil
}

func (x *DocumentProto) GetEmbedding() *NdArrayProto {
	if x != nil {
		return x.Embedding
	}
	return nil
}

func (x *DocumentProto) GetTags() *structpb.Struct {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *DocumentProto) GetScores() map[string]*NamedScoreProto {
	if x != nil {
		return x.Scores
	}
	return nil
}

func (x *DocumentProto) GetEvaluations() map[string]*NamedScoreProto {
	if x != nil {
		return x.Evaluations
	}
	return nil
}

func (x *DocumentProto) GetXMetadata() *structpb.Struct {
	if x != nil {
		return x.XMetadata
	}
	return nil
}

type isDocumentProto_Content interface {
	isDocumentProto_Content()
}

type DocumentProto_Blob struct {
	// the raw binary content of this document, which often represents the original document when comes into jina
	Blob []byte `protobuf:"bytes,2,opt,name=blob,proto3,oneof"`
}

type DocumentProto_Tensor struct {
	// the ndarray of the image/audio/video document
	Tensor *NdArrayProto `protobuf:"bytes,3,opt,name=tensor,proto3,oneof"`
}

type DocumentProto_Text struct {
	// a text document
	Text string `protobuf:"bytes,4,opt,name=text,proto3,oneof"`
}

func (*DocumentProto_Blob) isDocumentProto_Content() {}

func (*DocumentProto_Tensor) isDocumentProto_Content() {}

func (*DocumentProto_Text) isDocumentProto_Content() {}

type DocumentArrayProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Docs []*DocumentProto `protobuf:"bytes,1,rep,name=docs,proto3" json:"docs,omitempty"` // a list of Documents
}

func (x *DocumentArrayProto) Reset() {
	*x = DocumentArrayProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_docarray_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DocumentArrayProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DocumentArrayProto) ProtoMessage() {}

func (x *DocumentArrayProto) ProtoReflect() protoreflect.Message {
	mi := &file_docarray_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DocumentArrayProto.ProtoReflect.Descriptor instead.
func (*DocumentArrayProto) Descriptor() ([]byte, []int) {
	return file_docarray_proto_rawDescGZIP(), []int{5}
}

func (x *DocumentArrayProto) GetDocs() []*DocumentProto {
	if x != nil {
		return x.Docs
	}
	return nil
}

var File_docarray_proto protoreflect.FileDescriptor

var file_docarray_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x64, 0x6f, 0x63, 0x61, 0x72, 0x72, 0x61, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x64, 0x6f, 0x63, 0x61, 0x72, 0x72, 0x61, 0x79, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x57, 0x0a, 0x11, 0x44, 0x65, 0x6e, 0x73,
	0x65, 0x4e, 0x64, 0x41, 0x72, 0x72, 0x61, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x0a,
	0x06, 0x62, 0x75, 0x66, 0x66, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x62,
	0x75, 0x66, 0x66, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x68, 0x61, 0x70, 0x65, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0d, 0x52, 0x05, 0x73, 0x68, 0x61, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x64,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x64, 0x74, 0x79, 0x70,
	0x65, 0x22, 0xda, 0x01, 0x0a, 0x0c, 0x4e, 0x64, 0x41, 0x72, 0x72, 0x61, 0x79, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x33, 0x0a, 0x05, 0x64, 0x65, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1b, 0x2e, 0x64, 0x6f, 0x63, 0x61, 0x72, 0x72, 0x61, 0x79, 0x2e, 0x44, 0x65, 0x6e,
	0x73, 0x65, 0x4e, 0x64, 0x41, 0x72, 0x72, 0x61, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x48, 0x00,
	0x52, 0x05, 0x64, 0x65, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x06, 0x73, 0x70, 0x61, 0x72, 0x73,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x64, 0x6f, 0x63, 0x61, 0x72, 0x72,
	0x61, 0x79, 0x2e, 0x53, 0x70, 0x61, 0x72, 0x73, 0x65, 0x4e, 0x64, 0x41, 0x72, 0x72, 0x61, 0x79,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x48, 0x00, 0x52, 0x06, 0x73, 0x70, 0x61, 0x72, 0x73, 0x65, 0x12,
	0x19, 0x0a, 0x08, 0x63, 0x6c, 0x73, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x63, 0x6c, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x37, 0x0a, 0x0a, 0x70, 0x61,
	0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74,
	0x65, 0x72, 0x73, 0x42, 0x09, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x96,
	0x01, 0x0a, 0x12, 0x53, 0x70, 0x61, 0x72, 0x73, 0x65, 0x4e, 0x64, 0x41, 0x72, 0x72, 0x61, 0x79,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x35, 0x0a, 0x07, 0x69, 0x6e, 0x64, 0x69, 0x63, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x64, 0x6f, 0x63, 0x61, 0x72, 0x72, 0x61,
	0x79, 0x2e, 0x44, 0x65, 0x6e, 0x73, 0x65, 0x4e, 0x64, 0x41, 0x72, 0x72, 0x61, 0x79, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x52, 0x07, 0x69, 0x6e, 0x64, 0x69, 0x63, 0x65, 0x73, 0x12, 0x33, 0x0a, 0x06,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x64,
	0x6f, 0x63, 0x61, 0x72, 0x72, 0x61, 0x79, 0x2e, 0x44, 0x65, 0x6e, 0x73, 0x65, 0x4e, 0x64, 0x41,
	0x72, 0x72, 0x61, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x73, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x68, 0x61, 0x70, 0x65, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0d,
	0x52, 0x05, 0x73, 0x68, 0x61, 0x70, 0x65, 0x22, 0x79, 0x0a, 0x0f, 0x4e, 0x61, 0x6d, 0x65, 0x64,
	0x53, 0x63, 0x6f, 0x72, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x12, 0x17, 0x0a, 0x07, 0x6f, 0x70, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x6f, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x15, 0x0a, 0x06, 0x72,
	0x65, 0x66, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x72, 0x65, 0x66,
	0x49, 0x64, 0x22, 0xb3, 0x07, 0x0a, 0x0d, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x04, 0x62, 0x6c, 0x6f, 0x62, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0c, 0x48, 0x00, 0x52, 0x04, 0x62, 0x6c, 0x6f, 0x62, 0x12, 0x30, 0x0a, 0x06, 0x74, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x64, 0x6f, 0x63,
	0x61, 0x72, 0x72, 0x61, 0x79, 0x2e, 0x4e, 0x64, 0x41, 0x72, 0x72, 0x61, 0x79, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x48, 0x00, 0x52, 0x06, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x04,
	0x74, 0x65, 0x78, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x74, 0x65,
	0x78, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x67, 0x72, 0x61, 0x6e, 0x75, 0x6c, 0x61, 0x72, 0x69, 0x74,
	0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x67, 0x72, 0x61, 0x6e, 0x75, 0x6c, 0x61,
	0x72, 0x69, 0x74, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x64, 0x6a, 0x61, 0x63, 0x65, 0x6e, 0x63,
	0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x61, 0x64, 0x6a, 0x61, 0x63, 0x65, 0x6e,
	0x63, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x69, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x69, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x6f, 0x64,
	0x61, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x6f, 0x64,
	0x61, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x69, 0x6d, 0x65, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x69, 0x6d, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x02, 0x52, 0x08, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2f, 0x0a, 0x06, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x73,
	0x18, 0x0e, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x64, 0x6f, 0x63, 0x61, 0x72, 0x72, 0x61,
	0x79, 0x2e, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52,
	0x06, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x12, 0x31, 0x0a, 0x07, 0x6d, 0x61, 0x74, 0x63, 0x68,
	0x65, 0x73, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x64, 0x6f, 0x63, 0x61, 0x72,
	0x72, 0x61, 0x79, 0x2e, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x52, 0x07, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x12, 0x34, 0x0a, 0x09, 0x65, 0x6d,
	0x62, 0x65, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x64, 0x6f, 0x63, 0x61, 0x72, 0x72, 0x61, 0x79, 0x2e, 0x4e, 0x64, 0x41, 0x72, 0x72, 0x61, 0x79,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52, 0x09, 0x65, 0x6d, 0x62, 0x65, 0x64, 0x64, 0x69, 0x6e, 0x67,
	0x12, 0x2b, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x11, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x3b, 0x0a,
	0x06, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x18, 0x12, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e,
	0x64, 0x6f, 0x63, 0x61, 0x72, 0x72, 0x61, 0x79, 0x2e, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e,
	0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x06, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x12, 0x4a, 0x0a, 0x0b, 0x65, 0x76,
	0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x13, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x28, 0x2e, 0x64, 0x6f, 0x63, 0x61, 0x72, 0x72, 0x61, 0x79, 0x2e, 0x44, 0x6f, 0x63, 0x75, 0x6d,
	0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x65, 0x76, 0x61, 0x6c, 0x75,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x34, 0x0a, 0x09, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x52, 0x08, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x54, 0x0a, 0x0b,
	0x53, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2f, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x64,
	0x6f, 0x63, 0x61, 0x72, 0x72, 0x61, 0x79, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x64, 0x53, 0x63, 0x6f,
	0x72, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x1a, 0x59, 0x0a, 0x10, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2f, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x64, 0x6f, 0x63, 0x61, 0x72, 0x72,
	0x61, 0x79, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x64, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x09, 0x0a,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x41, 0x0a, 0x12, 0x44, 0x6f, 0x63, 0x75,
	0x6d, 0x65, 0x6e, 0x74, 0x41, 0x72, 0x72, 0x61, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2b,
	0x0a, 0x04, 0x64, 0x6f, 0x63, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x64,
	0x6f, 0x63, 0x61, 0x72, 0x72, 0x61, 0x79, 0x2e, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52, 0x04, 0x64, 0x6f, 0x63, 0x73, 0x42, 0x27, 0x5a, 0x25, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6a, 0x69, 0x6e, 0x61, 0x2d, 0x61,
	0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2d, 0x67, 0x6f, 0x2f, 0x64, 0x6f, 0x63, 0x61,
	0x72, 0x72, 0x61, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_docarray_proto_rawDescOnce sync.Once
	file_docarray_proto_rawDescData = file_docarray_proto_rawDesc
)

func file_docarray_proto_rawDescGZIP() []byte {
	file_docarray_proto_rawDescOnce.Do(func() {
		file_docarray_proto_rawDescData = protoimpl.X.CompressGZIP(file_docarray_proto_rawDescData)
	})
	return file_docarray_proto_rawDescData
}

var file_docarray_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_docarray_proto_goTypes = []interface{}{
	(*DenseNdArrayProto)(nil),  // 0: docarray.DenseNdArrayProto
	(*NdArrayProto)(nil),       // 1: docarray.NdArrayProto
	(*SparseNdArrayProto)(nil), // 2: docarray.SparseNdArrayProto
	(*NamedScoreProto)(nil),    // 3: docarray.NamedScoreProto
	(*DocumentProto)(nil),      // 4: docarray.DocumentProto
	(*DocumentArrayProto)(nil), // 5: docarray.DocumentArrayProto
	nil,                        // 6: docarray.DocumentProto.ScoresEntry
	nil,                        // 7: docarray.DocumentProto.EvaluationsEntry
	(*structpb.Struct)(nil),    // 8: google.protobuf.Struct
}
var file_docarray_proto_depIdxs = []int32{
	0,  // 0: docarray.NdArrayProto.dense:type_name -> docarray.DenseNdArrayProto
	2,  // 1: docarray.NdArrayProto.sparse:type_name -> docarray.SparseNdArrayProto
	8,  // 2: docarray.NdArrayProto.parameters:type_name -> google.protobuf.Struct
	0,  // 3: docarray.SparseNdArrayProto.indices:type_name -> docarray.DenseNdArrayProto
	0,  // 4: docarray.SparseNdArrayProto.values:type_name -> docarray.DenseNdArrayProto
	1,  // 5: docarray.DocumentProto.tensor:type_name -> docarray.NdArrayProto
	4,  // 6: docarray.DocumentProto.chunks:type_name -> docarray.DocumentProto
	4,  // 7: docarray.DocumentProto.matches:type_name -> docarray.DocumentProto
	1,  // 8: docarray.DocumentProto.embedding:type_name -> docarray.NdArrayProto
	8,  // 9: docarray.DocumentProto.tags:type_name -> google.protobuf.Struct
	6,  // 10: docarray.DocumentProto.scores:type_name -> docarray.DocumentProto.ScoresEntry
	7,  // 11: docarray.DocumentProto.evaluations:type_name -> docarray.DocumentProto.EvaluationsEntry
	8,  // 12: docarray.DocumentProto._metadata:type_name -> google.protobuf.Struct
	4,  // 13: docarray.DocumentArrayProto.docs:type_name -> docarray.DocumentProto
	3,  // 14: docarray.DocumentProto.ScoresEntry.value:type_name -> docarray.NamedScoreProto
	3,  // 15: docarray.DocumentProto.EvaluationsEntry.value:type_name -> docarray.NamedScoreProto
	16, // [16:16] is the sub-list for method output_type
	16, // [16:16] is the sub-list for method input_type
	16, // [16:16] is the sub-list for extension type_name
	16, // [16:16] is the sub-list for extension extendee
	0,  // [0:16] is the sub-list for field type_name
}

func init() { file_docarray_proto_init() }
func file_docarray_proto_init() {
	if File_docarray_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_docarray_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DenseNdArrayProto); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_docarray_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NdArrayProto); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_docarray_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SparseNdArrayProto); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_docarray_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NamedScoreProto); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_docarray_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DocumentProto); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_docarray_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DocumentArrayProto); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_docarray_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*NdArrayProto_Dense)(nil),
		(*NdArrayProto_Sparse)(nil),
	}
	file_docarray_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*DocumentProto_Blob)(nil),
		(*DocumentProto_Tensor)(nil),
		(*DocumentProto_Text)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_docarray_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_docarray_proto_goTypes,
		DependencyIndexes: file_docarray_proto_depIdxs,
		MessageInfos:      file_docarray_proto_msgTypes,
	}.Build()
	File_docarray_proto = out.File
	file_docarray_proto_rawDesc = nil
	file_docarray_proto_goTypes = nil
	file_docarray_proto_depIdxs = nil
}
