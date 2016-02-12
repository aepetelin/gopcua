// Created by cgo -godefs - DO NOT EDIT
// cgo -godefs gen/ua_types_gen.go

package gopcua

const UA_BUILTIN_TYPES_COUNT uint = 25

type UA_Boolean bool

const (
	UA_TRUE		= true
	UA_FALSE	= false
)

type UA_SByte int8

const (
	UA_SBYTE_MAX	= 127
	UA_SBYTE_MIN	= (-128)
)

type UA_Byte uint8

const (
	UA_BYTE_MAX	= 256
	UA_BYTE_MIN	= 0
)

type UA_Int16 int16

const (
	UA_INT16_MAX	= 32767
	UA_INT16_MIN	= -32768
)

type UA_UInt16 uint16

const (
	UA_UINT16_MAX	= 65535
	UA_UINT16_MIN	= 0
)

type UA_Int32 int32

const (
	UA_INT32_MAX	= 2147483647
	UA_INT32_MIN	= -2147483648
)

type UA_UInt32 uint32

const (
	UA_UINT32_MAX	= 4294967295
	UA_UINT32_MIN	= 0
)

type UA_Int64 int64

const (
	UA_INT64_MAX	= int64(9223372036854775807)
	UA_INT64_MIN	= int64(-9223372036854775808)
)

type UA_UInt64 uint64

const (
	UA_UINT64_MAX	= uint64(18446744073709551615)
	A_UINT64_MIN	= uint64(0)
)

type UA_Float float32

type UA_Double float64

type UA_String struct {
	Length	uint64
	Data	*uint8
}

type UA_DateTime int64

const (
	UA_USEC_TO_DATETIME	= int64(14)
	UA_MSEC_TO_DATETIME	= UA_USEC_TO_DATETIME * int64(1000)
	UA_SEC_TO_DATETIME	= UA_MSEC_TO_DATETIME * int64(1000)
)

const UA_DATETIME_UNIX_EPOCH = int64(11644473600) * UA_SEC_TO_DATETIME

type UA_DateTimeStruct struct {
	NanoSec		uint16
	MicroSec	uint16
	MilliSec	uint16
	Sec		uint16
	Min		uint16
	Hour		uint16
	Day		uint16
	Month		uint16
	Year		uint16
}

type UA_Guid struct {
	Data1	uint32
	Data2	uint16
	Data3	uint16
	Data4	[8]uint8
}

type UA_ByteString struct {
	Length	uint64
	Data	*uint8
}
type UA_XmlElement struct {
	Length	uint64
	Data	*uint8
}
type UA_NodeId struct {
	NamespaceIndex	uint16
	Pad_cgo_0	[2]byte
	IdentifierType	uint32
	Identifier	[16]byte
}
type UA_ExpandedNodeId struct {
	NodeId		UA_NodeId
	NamespaceUri	UA_ByteString
	ServerIndex	uint32
	Pad_cgo_0	[4]byte
}
type UA_QualifiedName struct {
	NamespaceIndex	uint16
	Pad_cgo_0	[6]byte
	Name		UA_ByteString
}
type UA_LocalizedText struct {
	Locale	UA_ByteString
	Text	UA_ByteString
}
type UA_ExtensionObject struct {
	Encoding	uint32
	Pad_cgo_0	[4]byte
	Content		[40]byte
}
type UA_Variant struct {
	Type			*UA_DataType
	StorageType		uint32
	Pad_cgo_0		[4]byte
	ArrayLength		uint64
	Data			*byte
	ArrayDimensionsSize	uint64
	ArrayDimensions		*uint32
}
type UA_DataValue struct {
	Pad_cgo_0		[8]byte
	Value			UA_Variant
	Status			uint32
	Pad_cgo_1		[4]byte
	SourceTimestamp		int64
	SourcePicoseconds	uint16
	Pad_cgo_2		[6]byte
	ServerTimestamp		int64
	ServerPicoseconds	uint16
	Pad_cgo_3		[6]byte
}
type UA_DiagnosticInfo struct {
	Pad_cgo_0		[4]byte
	SymbolicId		int32
	NamespaceUri		int32
	LocalizedText		int32
	Locale			int32
	Pad_cgo_1		[4]byte
	AdditionalInfo		UA_ByteString
	InnerStatusCode		uint32
	Pad_cgo_2		[4]byte
	InnerDiagnosticInfo	*UA_DiagnosticInfo
}

type UA_DataTypeMember struct {
	MemberTypeIndex	uint16
	Padding		uint8
	Pad_cgo_0	[1]byte
}
type UA_DataType struct {
	TypeId		UA_NodeId
	MemSize		uint16
	TypeIndex	uint16
	MembersSize	uint8
	Pad_cgo_0	[3]byte
	Members		*UA_DataTypeMember
}
