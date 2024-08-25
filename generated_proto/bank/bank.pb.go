// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.6.1
// source: proto/bank/bank.proto

package bank

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TransactionType int32

const (
	TransactionType_TRANSACTION_TYPE_UNSPECIFIED TransactionType = 0
	TransactionType_TRANSACTION_TYPE_IN          TransactionType = 1
	TransactionType_TRANSACTION_TYPE_OUT         TransactionType = 2
)

// Enum value maps for TransactionType.
var (
	TransactionType_name = map[int32]string{
		0: "TRANSACTION_TYPE_UNSPECIFIED",
		1: "TRANSACTION_TYPE_IN",
		2: "TRANSACTION_TYPE_OUT",
	}
	TransactionType_value = map[string]int32{
		"TRANSACTION_TYPE_UNSPECIFIED": 0,
		"TRANSACTION_TYPE_IN":          1,
		"TRANSACTION_TYPE_OUT":         2,
	}
)

func (x TransactionType) Enum() *TransactionType {
	p := new(TransactionType)
	*p = x
	return p
}

func (x TransactionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TransactionType) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_bank_bank_proto_enumTypes[0].Descriptor()
}

func (TransactionType) Type() protoreflect.EnumType {
	return &file_proto_bank_bank_proto_enumTypes[0]
}

func (x TransactionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TransactionType.Descriptor instead.
func (TransactionType) EnumDescriptor() ([]byte, []int) {
	return file_proto_bank_bank_proto_rawDescGZIP(), []int{0}
}

type TransferStatus int32

const (
	TransferStatus_TRANSFER_STATUS_UNSPECIFIED TransferStatus = 0
	TransferStatus_TRANSFER_STATUS_SUCCESS     TransferStatus = 1
	TransferStatus_TRANSFER_STATUS_FAILED      TransferStatus = 2
)

// Enum value maps for TransferStatus.
var (
	TransferStatus_name = map[int32]string{
		0: "TRANSFER_STATUS_UNSPECIFIED",
		1: "TRANSFER_STATUS_SUCCESS",
		2: "TRANSFER_STATUS_FAILED",
	}
	TransferStatus_value = map[string]int32{
		"TRANSFER_STATUS_UNSPECIFIED": 0,
		"TRANSFER_STATUS_SUCCESS":     1,
		"TRANSFER_STATUS_FAILED":      2,
	}
)

func (x TransferStatus) Enum() *TransferStatus {
	p := new(TransferStatus)
	*p = x
	return p
}

func (x TransferStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TransferStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_bank_bank_proto_enumTypes[1].Descriptor()
}

func (TransferStatus) Type() protoreflect.EnumType {
	return &file_proto_bank_bank_proto_enumTypes[1]
}

func (x TransferStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TransferStatus.Descriptor instead.
func (TransferStatus) EnumDescriptor() ([]byte, []int) {
	return file_proto_bank_bank_proto_rawDescGZIP(), []int{1}
}

type CurrentBalanceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountNumber string `protobuf:"bytes,1,opt,name=account_number,proto3" json:"account_number,omitempty"`
}

func (x *CurrentBalanceRequest) Reset() {
	*x = CurrentBalanceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_bank_bank_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CurrentBalanceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CurrentBalanceRequest) ProtoMessage() {}

func (x *CurrentBalanceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_bank_bank_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CurrentBalanceRequest.ProtoReflect.Descriptor instead.
func (*CurrentBalanceRequest) Descriptor() ([]byte, []int) {
	return file_proto_bank_bank_proto_rawDescGZIP(), []int{0}
}

func (x *CurrentBalanceRequest) GetAccountNumber() string {
	if x != nil {
		return x.AccountNumber
	}
	return ""
}

type CurrentBalanceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Amount      float64 `protobuf:"fixed64,1,opt,name=amount,proto3" json:"amount,omitempty"`
	CurrentDate string  `protobuf:"bytes,2,opt,name=current_date,proto3" json:"current_date,omitempty"`
}

func (x *CurrentBalanceResponse) Reset() {
	*x = CurrentBalanceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_bank_bank_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CurrentBalanceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CurrentBalanceResponse) ProtoMessage() {}

func (x *CurrentBalanceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_bank_bank_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CurrentBalanceResponse.ProtoReflect.Descriptor instead.
func (*CurrentBalanceResponse) Descriptor() ([]byte, []int) {
	return file_proto_bank_bank_proto_rawDescGZIP(), []int{1}
}

func (x *CurrentBalanceResponse) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *CurrentBalanceResponse) GetCurrentDate() string {
	if x != nil {
		return x.CurrentDate
	}
	return ""
}

type ExchangeRateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FromCurrency string `protobuf:"bytes,1,opt,name=from_currency,proto3" json:"from_currency,omitempty"`
	ToCurrency   string `protobuf:"bytes,2,opt,name=to_currency,proto3" json:"to_currency,omitempty"`
}

func (x *ExchangeRateRequest) Reset() {
	*x = ExchangeRateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_bank_bank_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExchangeRateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExchangeRateRequest) ProtoMessage() {}

func (x *ExchangeRateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_bank_bank_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExchangeRateRequest.ProtoReflect.Descriptor instead.
func (*ExchangeRateRequest) Descriptor() ([]byte, []int) {
	return file_proto_bank_bank_proto_rawDescGZIP(), []int{2}
}

func (x *ExchangeRateRequest) GetFromCurrency() string {
	if x != nil {
		return x.FromCurrency
	}
	return ""
}

func (x *ExchangeRateRequest) GetToCurrency() string {
	if x != nil {
		return x.ToCurrency
	}
	return ""
}

type ExchangeRateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FromCurrency string  `protobuf:"bytes,1,opt,name=from_currency,proto3" json:"from_currency,omitempty"`
	ToCurrency   string  `protobuf:"bytes,2,opt,name=to_currency,proto3" json:"to_currency,omitempty"`
	Rate         float64 `protobuf:"fixed64,3,opt,name=rate,proto3" json:"rate,omitempty"`
	Timestamp    string  `protobuf:"bytes,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *ExchangeRateResponse) Reset() {
	*x = ExchangeRateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_bank_bank_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExchangeRateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExchangeRateResponse) ProtoMessage() {}

func (x *ExchangeRateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_bank_bank_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExchangeRateResponse.ProtoReflect.Descriptor instead.
func (*ExchangeRateResponse) Descriptor() ([]byte, []int) {
	return file_proto_bank_bank_proto_rawDescGZIP(), []int{3}
}

func (x *ExchangeRateResponse) GetFromCurrency() string {
	if x != nil {
		return x.FromCurrency
	}
	return ""
}

func (x *ExchangeRateResponse) GetToCurrency() string {
	if x != nil {
		return x.ToCurrency
	}
	return ""
}

func (x *ExchangeRateResponse) GetRate() float64 {
	if x != nil {
		return x.Rate
	}
	return 0
}

func (x *ExchangeRateResponse) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

type Transaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountNumber string          `protobuf:"bytes,1,opt,name=account_number,proto3" json:"account_number,omitempty"`
	Type          TransactionType `protobuf:"varint,2,opt,name=type,proto3,enum=bank.TransactionType" json:"type,omitempty"`
	Amount        float64         `protobuf:"fixed64,3,opt,name=amount,proto3" json:"amount,omitempty"`
	Timestamp     string          `protobuf:"bytes,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Notes         string          `protobuf:"bytes,16,opt,name=notes,proto3" json:"notes,omitempty"`
}

func (x *Transaction) Reset() {
	*x = Transaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_bank_bank_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transaction) ProtoMessage() {}

func (x *Transaction) ProtoReflect() protoreflect.Message {
	mi := &file_proto_bank_bank_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transaction.ProtoReflect.Descriptor instead.
func (*Transaction) Descriptor() ([]byte, []int) {
	return file_proto_bank_bank_proto_rawDescGZIP(), []int{4}
}

func (x *Transaction) GetAccountNumber() string {
	if x != nil {
		return x.AccountNumber
	}
	return ""
}

func (x *Transaction) GetType() TransactionType {
	if x != nil {
		return x.Type
	}
	return TransactionType_TRANSACTION_TYPE_UNSPECIFIED
}

func (x *Transaction) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *Transaction) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

func (x *Transaction) GetNotes() string {
	if x != nil {
		return x.Notes
	}
	return ""
}

type TransactionSummary struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountNumber   string  `protobuf:"bytes,1,opt,name=account_number,proto3" json:"account_number,omitempty"`
	SumAmountIn     float64 `protobuf:"fixed64,2,opt,name=sum_amount_in,proto3" json:"sum_amount_in,omitempty"`
	SumAmountOut    float64 `protobuf:"fixed64,3,opt,name=sum_amount_out,proto3" json:"sum_amount_out,omitempty"`
	SumTotal        float64 `protobuf:"fixed64,4,opt,name=sum_total,proto3" json:"sum_total,omitempty"`
	TransactionDate string  `protobuf:"bytes,5,opt,name=transaction_date,proto3" json:"transaction_date,omitempty"`
}

func (x *TransactionSummary) Reset() {
	*x = TransactionSummary{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_bank_bank_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionSummary) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionSummary) ProtoMessage() {}

func (x *TransactionSummary) ProtoReflect() protoreflect.Message {
	mi := &file_proto_bank_bank_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionSummary.ProtoReflect.Descriptor instead.
func (*TransactionSummary) Descriptor() ([]byte, []int) {
	return file_proto_bank_bank_proto_rawDescGZIP(), []int{5}
}

func (x *TransactionSummary) GetAccountNumber() string {
	if x != nil {
		return x.AccountNumber
	}
	return ""
}

func (x *TransactionSummary) GetSumAmountIn() float64 {
	if x != nil {
		return x.SumAmountIn
	}
	return 0
}

func (x *TransactionSummary) GetSumAmountOut() float64 {
	if x != nil {
		return x.SumAmountOut
	}
	return 0
}

func (x *TransactionSummary) GetSumTotal() float64 {
	if x != nil {
		return x.SumTotal
	}
	return 0
}

func (x *TransactionSummary) GetTransactionDate() string {
	if x != nil {
		return x.TransactionDate
	}
	return ""
}

type TransferRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FromAccountNumber string  `protobuf:"bytes,1,opt,name=from_account_number,proto3" json:"from_account_number,omitempty"`
	ToAccountNumber   string  `protobuf:"bytes,2,opt,name=to_account_number,proto3" json:"to_account_number,omitempty"`
	Currency          string  `protobuf:"bytes,3,opt,name=currency,proto3" json:"currency,omitempty"`
	Amount            float64 `protobuf:"fixed64,4,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *TransferRequest) Reset() {
	*x = TransferRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_bank_bank_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransferRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransferRequest) ProtoMessage() {}

func (x *TransferRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_bank_bank_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransferRequest.ProtoReflect.Descriptor instead.
func (*TransferRequest) Descriptor() ([]byte, []int) {
	return file_proto_bank_bank_proto_rawDescGZIP(), []int{6}
}

func (x *TransferRequest) GetFromAccountNumber() string {
	if x != nil {
		return x.FromAccountNumber
	}
	return ""
}

func (x *TransferRequest) GetToAccountNumber() string {
	if x != nil {
		return x.ToAccountNumber
	}
	return ""
}

func (x *TransferRequest) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *TransferRequest) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type TransferResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FromAccountNumber string         `protobuf:"bytes,1,opt,name=from_account_number,proto3" json:"from_account_number,omitempty"`
	ToAccountNumber   string         `protobuf:"bytes,2,opt,name=to_account_number,proto3" json:"to_account_number,omitempty"`
	Currency          string         `protobuf:"bytes,3,opt,name=currency,proto3" json:"currency,omitempty"`
	Amount            float64        `protobuf:"fixed64,4,opt,name=amount,proto3" json:"amount,omitempty"`
	Status            TransferStatus `protobuf:"varint,5,opt,name=status,proto3,enum=bank.TransferStatus" json:"status,omitempty"`
	Timestamp         string         `protobuf:"bytes,6,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *TransferResponse) Reset() {
	*x = TransferResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_bank_bank_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransferResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransferResponse) ProtoMessage() {}

func (x *TransferResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_bank_bank_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransferResponse.ProtoReflect.Descriptor instead.
func (*TransferResponse) Descriptor() ([]byte, []int) {
	return file_proto_bank_bank_proto_rawDescGZIP(), []int{7}
}

func (x *TransferResponse) GetFromAccountNumber() string {
	if x != nil {
		return x.FromAccountNumber
	}
	return ""
}

func (x *TransferResponse) GetToAccountNumber() string {
	if x != nil {
		return x.ToAccountNumber
	}
	return ""
}

func (x *TransferResponse) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *TransferResponse) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *TransferResponse) GetStatus() TransferStatus {
	if x != nil {
		return x.Status
	}
	return TransferStatus_TRANSFER_STATUS_UNSPECIFIED
}

func (x *TransferResponse) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

var File_proto_bank_bank_proto protoreflect.FileDescriptor

var file_proto_bank_bank_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x61, 0x6e, 0x6b, 0x2f, 0x62, 0x61, 0x6e,
	0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x62, 0x61, 0x6e, 0x6b, 0x22, 0x3f, 0x0a,
	0x15, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x0e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x54,
	0x0a, 0x16, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x22, 0x0a, 0x0c, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f,
	0x64, 0x61, 0x74, 0x65, 0x22, 0x5d, 0x0a, 0x13, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x52, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x66,
	0x72, 0x6f, 0x6d, 0x5f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x79, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x6f, 0x5f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x74, 0x6f, 0x5f, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x63, 0x79, 0x22, 0x90, 0x01, 0x0a, 0x14, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x52, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x0d,
	0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x63, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x6f, 0x5f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x74, 0x6f, 0x5f, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x63, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x04, 0x72, 0x61, 0x74, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0xac, 0x01, 0x0a, 0x0b, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x26, 0x0a, 0x0e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x29,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x62,
	0x61, 0x6e, 0x6b, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12,
	0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x6e, 0x6f, 0x74, 0x65, 0x73, 0x22, 0xd4, 0x01, 0x0a, 0x12, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x26, 0x0a, 0x0e,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x12, 0x24, 0x0a, 0x0d, 0x73, 0x75, 0x6d, 0x5f, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x5f, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0d, 0x73, 0x75, 0x6d,
	0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x6e, 0x12, 0x26, 0x0a, 0x0e, 0x73, 0x75,
	0x6d, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6f, 0x75, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x0e, 0x73, 0x75, 0x6d, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6f,
	0x75, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x75, 0x6d, 0x5f, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x73, 0x75, 0x6d, 0x5f, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x12, 0x2a, 0x0a, 0x10, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x64, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x22, 0xa5, 0x01, 0x0a,
	0x0f, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x30, 0x0a, 0x13, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x66,
	0x72, 0x6f, 0x6d, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x12, 0x2c, 0x0a, 0x11, 0x74, 0x6f, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x74,
	0x6f, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x12, 0x1a, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x16, 0x0a, 0x06,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x22, 0xf2, 0x01, 0x0a, 0x10, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x13, 0x66, 0x72, 0x6f,
	0x6d, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x2c, 0x0a, 0x11, 0x74,
	0x6f, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x74, 0x6f, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2c, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e,
	0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2a, 0x66, 0x0a, 0x0f, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x20, 0x0a, 0x1c,
	0x54, 0x52, 0x41, 0x4e, 0x53, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x17,
	0x0a, 0x13, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x49, 0x4e, 0x10, 0x01, 0x12, 0x18, 0x0a, 0x14, 0x54, 0x52, 0x41, 0x4e, 0x53,
	0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4f, 0x55, 0x54, 0x10,
	0x02, 0x2a, 0x6a, 0x0a, 0x0e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x1f, 0x0a, 0x1b, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x46, 0x45, 0x52, 0x5f,
	0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x1b, 0x0a, 0x17, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x46, 0x45, 0x52,
	0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10,
	0x01, 0x12, 0x1a, 0x0a, 0x16, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x46, 0x45, 0x52, 0x5f, 0x53, 0x54,
	0x41, 0x54, 0x55, 0x53, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x02, 0x32, 0xc3, 0x02,
	0x0a, 0x0b, 0x42, 0x61, 0x6e, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x50, 0x0a,
	0x11, 0x47, 0x65, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e,
	0x63, 0x65, 0x12, 0x1b, 0x2e, 0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1c, 0x2e, 0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x42, 0x61,
	0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x4f, 0x0a, 0x12, 0x46, 0x65, 0x74, 0x63, 0x68, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x52, 0x61, 0x74, 0x65, 0x73, 0x12, 0x19, 0x2e, 0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x45, 0x78, 0x63,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1a, 0x2e, 0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x52, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01,
	0x12, 0x48, 0x0a, 0x15, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x69, 0x7a, 0x65, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x11, 0x2e, 0x62, 0x61, 0x6e, 0x6b,
	0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x18, 0x2e, 0x62,
	0x61, 0x6e, 0x6b, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x22, 0x00, 0x28, 0x01, 0x12, 0x47, 0x0a, 0x10, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x12, 0x15,
	0x2e, 0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x66, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28,
	0x01, 0x30, 0x01, 0x42, 0x1f, 0x5a, 0x1d, 0x67, 0x72, 0x70, 0x63, 0x62, 0x61, 0x6e, 0x6b, 0x2f,
	0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x62, 0x61, 0x6e, 0x6b, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_bank_bank_proto_rawDescOnce sync.Once
	file_proto_bank_bank_proto_rawDescData = file_proto_bank_bank_proto_rawDesc
)

func file_proto_bank_bank_proto_rawDescGZIP() []byte {
	file_proto_bank_bank_proto_rawDescOnce.Do(func() {
		file_proto_bank_bank_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_bank_bank_proto_rawDescData)
	})
	return file_proto_bank_bank_proto_rawDescData
}

var file_proto_bank_bank_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_proto_bank_bank_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_proto_bank_bank_proto_goTypes = []any{
	(TransactionType)(0),           // 0: bank.TransactionType
	(TransferStatus)(0),            // 1: bank.TransferStatus
	(*CurrentBalanceRequest)(nil),  // 2: bank.CurrentBalanceRequest
	(*CurrentBalanceResponse)(nil), // 3: bank.CurrentBalanceResponse
	(*ExchangeRateRequest)(nil),    // 4: bank.ExchangeRateRequest
	(*ExchangeRateResponse)(nil),   // 5: bank.ExchangeRateResponse
	(*Transaction)(nil),            // 6: bank.Transaction
	(*TransactionSummary)(nil),     // 7: bank.TransactionSummary
	(*TransferRequest)(nil),        // 8: bank.TransferRequest
	(*TransferResponse)(nil),       // 9: bank.TransferResponse
}
var file_proto_bank_bank_proto_depIdxs = []int32{
	0, // 0: bank.Transaction.type:type_name -> bank.TransactionType
	1, // 1: bank.TransferResponse.status:type_name -> bank.TransferStatus
	2, // 2: bank.BankService.GetCurrentBalance:input_type -> bank.CurrentBalanceRequest
	4, // 3: bank.BankService.FetchExchangeRates:input_type -> bank.ExchangeRateRequest
	6, // 4: bank.BankService.SummarizeTransactions:input_type -> bank.Transaction
	8, // 5: bank.BankService.TransferMultiple:input_type -> bank.TransferRequest
	3, // 6: bank.BankService.GetCurrentBalance:output_type -> bank.CurrentBalanceResponse
	5, // 7: bank.BankService.FetchExchangeRates:output_type -> bank.ExchangeRateResponse
	7, // 8: bank.BankService.SummarizeTransactions:output_type -> bank.TransactionSummary
	9, // 9: bank.BankService.TransferMultiple:output_type -> bank.TransferResponse
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_bank_bank_proto_init() }
func file_proto_bank_bank_proto_init() {
	if File_proto_bank_bank_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_bank_bank_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CurrentBalanceRequest); i {
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
		file_proto_bank_bank_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*CurrentBalanceResponse); i {
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
		file_proto_bank_bank_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*ExchangeRateRequest); i {
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
		file_proto_bank_bank_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*ExchangeRateResponse); i {
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
		file_proto_bank_bank_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*Transaction); i {
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
		file_proto_bank_bank_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*TransactionSummary); i {
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
		file_proto_bank_bank_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*TransferRequest); i {
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
		file_proto_bank_bank_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*TransferResponse); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_bank_bank_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_bank_bank_proto_goTypes,
		DependencyIndexes: file_proto_bank_bank_proto_depIdxs,
		EnumInfos:         file_proto_bank_bank_proto_enumTypes,
		MessageInfos:      file_proto_bank_bank_proto_msgTypes,
	}.Build()
	File_proto_bank_bank_proto = out.File
	file_proto_bank_bank_proto_rawDesc = nil
	file_proto_bank_bank_proto_goTypes = nil
	file_proto_bank_bank_proto_depIdxs = nil
}
