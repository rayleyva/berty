// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: errcode.proto

package errcode

import (
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = proto.Marshal
	_ = fmt.Errorf
	_ = math.Inf
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type ErrCode int32

const (
	Undefined                                  ErrCode = 0
	TODO                                       ErrCode = 666
	ErrNotImplemented                          ErrCode = 777
	ErrInternal                                ErrCode = 888
	ErrInvalidInput                            ErrCode = 100
	ErrInvalidRange                            ErrCode = 101
	ErrMissingInput                            ErrCode = 102
	ErrSerialization                           ErrCode = 103
	ErrDeserialization                         ErrCode = 104
	ErrStreamRead                              ErrCode = 105
	ErrStreamWrite                             ErrCode = 106
	ErrStreamTransform                         ErrCode = 110
	ErrStreamSendAndClose                      ErrCode = 111
	ErrStreamHeaderWrite                       ErrCode = 112
	ErrStreamHeaderRead                        ErrCode = 115
	ErrStreamSink                              ErrCode = 113
	ErrStreamCloseAndRecv                      ErrCode = 114
	ErrMissingMapKey                           ErrCode = 107
	ErrDBWrite                                 ErrCode = 108
	ErrDBRead                                  ErrCode = 109
	ErrCryptoRandomGeneration                  ErrCode = 200
	ErrCryptoKeyGeneration                     ErrCode = 201
	ErrCryptoNonceGeneration                   ErrCode = 202
	ErrCryptoSignature                         ErrCode = 203
	ErrCryptoSignatureVerification             ErrCode = 204
	ErrCryptoDecrypt                           ErrCode = 205
	ErrCryptoDecryptPayload                    ErrCode = 206
	ErrCryptoEncrypt                           ErrCode = 207
	ErrCryptoKeyConversion                     ErrCode = 208
	ErrCryptoCipherInit                        ErrCode = 209
	ErrCryptoKeyDerivation                     ErrCode = 210
	ErrMap                                     ErrCode = 300
	ErrForEach                                 ErrCode = 301
	ErrKeystoreGet                             ErrCode = 400
	ErrKeystorePut                             ErrCode = 401
	ErrOrbitDBInit                             ErrCode = 1000
	ErrOrbitDBOpen                             ErrCode = 1001
	ErrOrbitDBAppend                           ErrCode = 1002
	ErrOrbitDBDeserialization                  ErrCode = 1003
	ErrOrbitDBStoreCast                        ErrCode = 1004
	ErrIPFSAdd                                 ErrCode = 1050
	ErrIPFSGet                                 ErrCode = 1051
	ErrHandshakeOwnEphemeralKeyGenSend         ErrCode = 1100
	ErrHandshakePeerEphemeralKeyRecv           ErrCode = 1101
	ErrHandshakeRequesterAuthenticateBoxKeyGen ErrCode = 1102
	ErrHandshakeResponderAcceptBoxKeyGen       ErrCode = 1103
	ErrHandshakeRequesterHello                 ErrCode = 1104
	ErrHandshakeResponderHello                 ErrCode = 1105
	ErrHandshakeRequesterAuthenticate          ErrCode = 1106
	ErrHandshakeResponderAccept                ErrCode = 1107
	ErrHandshakeRequesterAcknowledge           ErrCode = 1108
	ErrContactRequestSameAccount               ErrCode = 1200
	ErrContactRequestContactAlreadyAdded       ErrCode = 1201
	ErrContactRequestContactBlocked            ErrCode = 1202
	ErrContactRequestContactUndefined          ErrCode = 1203
	ErrContactRequestIncomingAlreadyReceived   ErrCode = 1204
	ErrGroupMemberLogEventOpen                 ErrCode = 1300
	ErrGroupMemberLogEventSignature            ErrCode = 1301
	ErrGroupMemberUnknownGroupID               ErrCode = 1302
	ErrGroupSecretOtherDestMember              ErrCode = 1303
	ErrGroupSecretAlreadySentToMember          ErrCode = 1304
	ErrGroupInvalidType                        ErrCode = 1305
	ErrGroupMissing                            ErrCode = 1306
	ErrGroupActivate                           ErrCode = 1307
	ErrGroupDeactivate                         ErrCode = 1308
	ErrGroupInfo                               ErrCode = 1309
	// Event errors
	ErrEventListMetadata                   ErrCode = 1400
	ErrEventListMessage                    ErrCode = 1401
	ErrMessageKeyPersistencePut            ErrCode = 1500
	ErrMessageKeyPersistenceGet            ErrCode = 1501
	ErrBridgeInterrupted                   ErrCode = 1600
	ErrBridgeNotRunning                    ErrCode = 1601
	ErrMessengerInvalidDeepLink            ErrCode = 2000
	ErrMessengerDeepLinkRequiresPassphrase ErrCode = 2001
	ErrMessengerDeepLinkInvalidPassphrase  ErrCode = 2002
	ErrDBEntryAlreadyExists                ErrCode = 2100
	ErrDBAddConversation                   ErrCode = 2101
	ErrDBAddContactRequestOutgoingSent     ErrCode = 2102
	ErrDBAddContactRequestOutgoingEnqueud  ErrCode = 2103
	ErrDBAddContactRequestIncomingReceived ErrCode = 2104
	ErrDBAddContactRequestIncomingAccepted ErrCode = 2105
	ErrDBAddGroupMemberDeviceAdded         ErrCode = 2106
	ErrDBMultipleRecords                   ErrCode = 2107
	ErrReplayProcessGroupMetadata          ErrCode = 2200
	ErrReplayProcessGroupMessage           ErrCode = 2201
	// API internals errors
	ErrAttachmentPrepare                 ErrCode = 2300
	ErrAttachmentRetrieve                ErrCode = 2301
	ErrProtocolSend                      ErrCode = 2302
	ErrCLINoTermcaps                     ErrCode = 3001
	ErrServicesAuth                      ErrCode = 4000
	ErrServicesAuthNotInitialized        ErrCode = 4001
	ErrServicesAuthWrongState            ErrCode = 4002
	ErrServicesAuthInvalidResponse       ErrCode = 4003
	ErrServicesAuthServer                ErrCode = 4004
	ErrServicesAuthCodeChallenge         ErrCode = 4005
	ErrServicesAuthServiceInvalidToken   ErrCode = 4006
	ErrServicesAuthServiceNotSupported   ErrCode = 4007
	ErrServicesAuthUnknownToken          ErrCode = 4008
	ErrServicesAuthInvalidURL            ErrCode = 4009
	ErrServiceReplication                ErrCode = 4100
	ErrServiceReplicationServer          ErrCode = 4101
	ErrServiceReplicationMissingEndpoint ErrCode = 4102
	ErrBertyAccount                      ErrCode = 5000
	ErrBertyAccountNoIDSpecified         ErrCode = 5001
	ErrBertyAccountAlreadyOpened         ErrCode = 5002
	ErrBertyAccountInvalidIDFormat       ErrCode = 5003
	ErrBertyAccountLoggerDecorator       ErrCode = 5004
	ErrBertyAccountGRPCClient            ErrCode = 5005
	ErrBertyAccountOpenAccount           ErrCode = 5006
	ErrBertyAccountDataNotFound          ErrCode = 5007
	ErrBertyAccountMetadataUpdate        ErrCode = 5008
	ErrBertyAccountManagerOpen           ErrCode = 5009
	ErrBertyAccountManagerClose          ErrCode = 5010
	ErrBertyAccountInvalidCLIArgs        ErrCode = 5011
	ErrBertyAccountFSError               ErrCode = 5012
	ErrBertyAccountAlreadyExists         ErrCode = 5013
	ErrBertyAccountNoBackupSpecified     ErrCode = 5014
	ErrBertyAccountIDGenFailed           ErrCode = 5015
	ErrBertyAccountCreationFailed        ErrCode = 5016
)

var ErrCode_name = map[int32]string{
	0:    "Undefined",
	666:  "TODO",
	777:  "ErrNotImplemented",
	888:  "ErrInternal",
	100:  "ErrInvalidInput",
	101:  "ErrInvalidRange",
	102:  "ErrMissingInput",
	103:  "ErrSerialization",
	104:  "ErrDeserialization",
	105:  "ErrStreamRead",
	106:  "ErrStreamWrite",
	110:  "ErrStreamTransform",
	111:  "ErrStreamSendAndClose",
	112:  "ErrStreamHeaderWrite",
	115:  "ErrStreamHeaderRead",
	113:  "ErrStreamSink",
	114:  "ErrStreamCloseAndRecv",
	107:  "ErrMissingMapKey",
	108:  "ErrDBWrite",
	109:  "ErrDBRead",
	200:  "ErrCryptoRandomGeneration",
	201:  "ErrCryptoKeyGeneration",
	202:  "ErrCryptoNonceGeneration",
	203:  "ErrCryptoSignature",
	204:  "ErrCryptoSignatureVerification",
	205:  "ErrCryptoDecrypt",
	206:  "ErrCryptoDecryptPayload",
	207:  "ErrCryptoEncrypt",
	208:  "ErrCryptoKeyConversion",
	209:  "ErrCryptoCipherInit",
	210:  "ErrCryptoKeyDerivation",
	300:  "ErrMap",
	301:  "ErrForEach",
	400:  "ErrKeystoreGet",
	401:  "ErrKeystorePut",
	1000: "ErrOrbitDBInit",
	1001: "ErrOrbitDBOpen",
	1002: "ErrOrbitDBAppend",
	1003: "ErrOrbitDBDeserialization",
	1004: "ErrOrbitDBStoreCast",
	1050: "ErrIPFSAdd",
	1051: "ErrIPFSGet",
	1100: "ErrHandshakeOwnEphemeralKeyGenSend",
	1101: "ErrHandshakePeerEphemeralKeyRecv",
	1102: "ErrHandshakeRequesterAuthenticateBoxKeyGen",
	1103: "ErrHandshakeResponderAcceptBoxKeyGen",
	1104: "ErrHandshakeRequesterHello",
	1105: "ErrHandshakeResponderHello",
	1106: "ErrHandshakeRequesterAuthenticate",
	1107: "ErrHandshakeResponderAccept",
	1108: "ErrHandshakeRequesterAcknowledge",
	1200: "ErrContactRequestSameAccount",
	1201: "ErrContactRequestContactAlreadyAdded",
	1202: "ErrContactRequestContactBlocked",
	1203: "ErrContactRequestContactUndefined",
	1204: "ErrContactRequestIncomingAlreadyReceived",
	1300: "ErrGroupMemberLogEventOpen",
	1301: "ErrGroupMemberLogEventSignature",
	1302: "ErrGroupMemberUnknownGroupID",
	1303: "ErrGroupSecretOtherDestMember",
	1304: "ErrGroupSecretAlreadySentToMember",
	1305: "ErrGroupInvalidType",
	1306: "ErrGroupMissing",
	1307: "ErrGroupActivate",
	1308: "ErrGroupDeactivate",
	1309: "ErrGroupInfo",
	1400: "ErrEventListMetadata",
	1401: "ErrEventListMessage",
	1500: "ErrMessageKeyPersistencePut",
	1501: "ErrMessageKeyPersistenceGet",
	1600: "ErrBridgeInterrupted",
	1601: "ErrBridgeNotRunning",
	2000: "ErrMessengerInvalidDeepLink",
	2001: "ErrMessengerDeepLinkRequiresPassphrase",
	2002: "ErrMessengerDeepLinkInvalidPassphrase",
	2100: "ErrDBEntryAlreadyExists",
	2101: "ErrDBAddConversation",
	2102: "ErrDBAddContactRequestOutgoingSent",
	2103: "ErrDBAddContactRequestOutgoingEnqueud",
	2104: "ErrDBAddContactRequestIncomingReceived",
	2105: "ErrDBAddContactRequestIncomingAccepted",
	2106: "ErrDBAddGroupMemberDeviceAdded",
	2107: "ErrDBMultipleRecords",
	2200: "ErrReplayProcessGroupMetadata",
	2201: "ErrReplayProcessGroupMessage",
	2300: "ErrAttachmentPrepare",
	2301: "ErrAttachmentRetrieve",
	2302: "ErrProtocolSend",
	3001: "ErrCLINoTermcaps",
	4000: "ErrServicesAuth",
	4001: "ErrServicesAuthNotInitialized",
	4002: "ErrServicesAuthWrongState",
	4003: "ErrServicesAuthInvalidResponse",
	4004: "ErrServicesAuthServer",
	4005: "ErrServicesAuthCodeChallenge",
	4006: "ErrServicesAuthServiceInvalidToken",
	4007: "ErrServicesAuthServiceNotSupported",
	4008: "ErrServicesAuthUnknownToken",
	4009: "ErrServicesAuthInvalidURL",
	4100: "ErrServiceReplication",
	4101: "ErrServiceReplicationServer",
	4102: "ErrServiceReplicationMissingEndpoint",
	5000: "ErrBertyAccount",
	5001: "ErrBertyAccountNoIDSpecified",
	5002: "ErrBertyAccountAlreadyOpened",
	5003: "ErrBertyAccountInvalidIDFormat",
	5004: "ErrBertyAccountLoggerDecorator",
	5005: "ErrBertyAccountGRPCClient",
	5006: "ErrBertyAccountOpenAccount",
	5007: "ErrBertyAccountDataNotFound",
	5008: "ErrBertyAccountMetadataUpdate",
	5009: "ErrBertyAccountManagerOpen",
	5010: "ErrBertyAccountManagerClose",
	5011: "ErrBertyAccountInvalidCLIArgs",
	5012: "ErrBertyAccountFSError",
	5013: "ErrBertyAccountAlreadyExists",
	5014: "ErrBertyAccountNoBackupSpecified",
	5015: "ErrBertyAccountIDGenFailed",
	5016: "ErrBertyAccountCreationFailed",
}

var ErrCode_value = map[string]int32{
	"Undefined":                                  0,
	"TODO":                                       666,
	"ErrNotImplemented":                          777,
	"ErrInternal":                                888,
	"ErrInvalidInput":                            100,
	"ErrInvalidRange":                            101,
	"ErrMissingInput":                            102,
	"ErrSerialization":                           103,
	"ErrDeserialization":                         104,
	"ErrStreamRead":                              105,
	"ErrStreamWrite":                             106,
	"ErrStreamTransform":                         110,
	"ErrStreamSendAndClose":                      111,
	"ErrStreamHeaderWrite":                       112,
	"ErrStreamHeaderRead":                        115,
	"ErrStreamSink":                              113,
	"ErrStreamCloseAndRecv":                      114,
	"ErrMissingMapKey":                           107,
	"ErrDBWrite":                                 108,
	"ErrDBRead":                                  109,
	"ErrCryptoRandomGeneration":                  200,
	"ErrCryptoKeyGeneration":                     201,
	"ErrCryptoNonceGeneration":                   202,
	"ErrCryptoSignature":                         203,
	"ErrCryptoSignatureVerification":             204,
	"ErrCryptoDecrypt":                           205,
	"ErrCryptoDecryptPayload":                    206,
	"ErrCryptoEncrypt":                           207,
	"ErrCryptoKeyConversion":                     208,
	"ErrCryptoCipherInit":                        209,
	"ErrCryptoKeyDerivation":                     210,
	"ErrMap":                                     300,
	"ErrForEach":                                 301,
	"ErrKeystoreGet":                             400,
	"ErrKeystorePut":                             401,
	"ErrOrbitDBInit":                             1000,
	"ErrOrbitDBOpen":                             1001,
	"ErrOrbitDBAppend":                           1002,
	"ErrOrbitDBDeserialization":                  1003,
	"ErrOrbitDBStoreCast":                        1004,
	"ErrIPFSAdd":                                 1050,
	"ErrIPFSGet":                                 1051,
	"ErrHandshakeOwnEphemeralKeyGenSend":         1100,
	"ErrHandshakePeerEphemeralKeyRecv":           1101,
	"ErrHandshakeRequesterAuthenticateBoxKeyGen": 1102,
	"ErrHandshakeResponderAcceptBoxKeyGen":       1103,
	"ErrHandshakeRequesterHello":                 1104,
	"ErrHandshakeResponderHello":                 1105,
	"ErrHandshakeRequesterAuthenticate":          1106,
	"ErrHandshakeResponderAccept":                1107,
	"ErrHandshakeRequesterAcknowledge":           1108,
	"ErrContactRequestSameAccount":               1200,
	"ErrContactRequestContactAlreadyAdded":       1201,
	"ErrContactRequestContactBlocked":            1202,
	"ErrContactRequestContactUndefined":          1203,
	"ErrContactRequestIncomingAlreadyReceived":   1204,
	"ErrGroupMemberLogEventOpen":                 1300,
	"ErrGroupMemberLogEventSignature":            1301,
	"ErrGroupMemberUnknownGroupID":               1302,
	"ErrGroupSecretOtherDestMember":              1303,
	"ErrGroupSecretAlreadySentToMember":          1304,
	"ErrGroupInvalidType":                        1305,
	"ErrGroupMissing":                            1306,
	"ErrGroupActivate":                           1307,
	"ErrGroupDeactivate":                         1308,
	"ErrGroupInfo":                               1309,
	"ErrEventListMetadata":                       1400,
	"ErrEventListMessage":                        1401,
	"ErrMessageKeyPersistencePut":                1500,
	"ErrMessageKeyPersistenceGet":                1501,
	"ErrBridgeInterrupted":                       1600,
	"ErrBridgeNotRunning":                        1601,
	"ErrMessengerInvalidDeepLink":                2000,
	"ErrMessengerDeepLinkRequiresPassphrase":     2001,
	"ErrMessengerDeepLinkInvalidPassphrase":      2002,
	"ErrDBEntryAlreadyExists":                    2100,
	"ErrDBAddConversation":                       2101,
	"ErrDBAddContactRequestOutgoingSent":         2102,
	"ErrDBAddContactRequestOutgoingEnqueud":      2103,
	"ErrDBAddContactRequestIncomingReceived":     2104,
	"ErrDBAddContactRequestIncomingAccepted":     2105,
	"ErrDBAddGroupMemberDeviceAdded":             2106,
	"ErrDBMultipleRecords":                       2107,
	"ErrReplayProcessGroupMetadata":              2200,
	"ErrReplayProcessGroupMessage":               2201,
	"ErrAttachmentPrepare":                       2300,
	"ErrAttachmentRetrieve":                      2301,
	"ErrProtocolSend":                            2302,
	"ErrCLINoTermcaps":                           3001,
	"ErrServicesAuth":                            4000,
	"ErrServicesAuthNotInitialized":              4001,
	"ErrServicesAuthWrongState":                  4002,
	"ErrServicesAuthInvalidResponse":             4003,
	"ErrServicesAuthServer":                      4004,
	"ErrServicesAuthCodeChallenge":               4005,
	"ErrServicesAuthServiceInvalidToken":         4006,
	"ErrServicesAuthServiceNotSupported":         4007,
	"ErrServicesAuthUnknownToken":                4008,
	"ErrServicesAuthInvalidURL":                  4009,
	"ErrServiceReplication":                      4100,
	"ErrServiceReplicationServer":                4101,
	"ErrServiceReplicationMissingEndpoint":       4102,
	"ErrBertyAccount":                            5000,
	"ErrBertyAccountNoIDSpecified":               5001,
	"ErrBertyAccountAlreadyOpened":               5002,
	"ErrBertyAccountInvalidIDFormat":             5003,
	"ErrBertyAccountLoggerDecorator":             5004,
	"ErrBertyAccountGRPCClient":                  5005,
	"ErrBertyAccountOpenAccount":                 5006,
	"ErrBertyAccountDataNotFound":                5007,
	"ErrBertyAccountMetadataUpdate":              5008,
	"ErrBertyAccountManagerOpen":                 5009,
	"ErrBertyAccountManagerClose":                5010,
	"ErrBertyAccountInvalidCLIArgs":              5011,
	"ErrBertyAccountFSError":                     5012,
	"ErrBertyAccountAlreadyExists":               5013,
	"ErrBertyAccountNoBackupSpecified":           5014,
	"ErrBertyAccountIDGenFailed":                 5015,
	"ErrBertyAccountCreationFailed":              5016,
}

func (x ErrCode) String() string {
	return proto.EnumName(ErrCode_name, int32(x))
}

func (ErrCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4240057316120df7, []int{0}
}

type ErrDetails struct {
	Codes                []ErrCode `protobuf:"varint,1,rep,packed,name=codes,proto3,enum=berty.errcode.ErrCode" json:"codes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ErrDetails) Reset()         { *m = ErrDetails{} }
func (m *ErrDetails) String() string { return proto.CompactTextString(m) }
func (*ErrDetails) ProtoMessage()    {}
func (*ErrDetails) Descriptor() ([]byte, []int) {
	return fileDescriptor_4240057316120df7, []int{0}
}

func (m *ErrDetails) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ErrDetails.Unmarshal(m, b)
}

func (m *ErrDetails) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ErrDetails.Marshal(b, m, deterministic)
}

func (m *ErrDetails) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ErrDetails.Merge(m, src)
}

func (m *ErrDetails) XXX_Size() int {
	return xxx_messageInfo_ErrDetails.Size(m)
}

func (m *ErrDetails) XXX_DiscardUnknown() {
	xxx_messageInfo_ErrDetails.DiscardUnknown(m)
}

var xxx_messageInfo_ErrDetails proto.InternalMessageInfo

func (m *ErrDetails) GetCodes() []ErrCode {
	if m != nil {
		return m.Codes
	}
	return nil
}

func init() {
	proto.RegisterEnum("berty.errcode.ErrCode", ErrCode_name, ErrCode_value)
	proto.RegisterType((*ErrDetails)(nil), "berty.errcode.ErrDetails")
}

func init() { proto.RegisterFile("errcode.proto", fileDescriptor_4240057316120df7) }

var fileDescriptor_4240057316120df7 = []byte{
	// 1695 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x57, 0xc9, 0x73, 0x5c, 0x47,
	0x19, 0xcf, 0x24, 0xc4, 0x1a, 0xb5, 0x17, 0x7d, 0x6a, 0x3b, 0xde, 0xe2, 0x78, 0x14, 0x13, 0x67,
	0x82, 0x01, 0xab, 0x0a, 0x6e, 0xdc, 0x66, 0x93, 0x3c, 0x65, 0x2d, 0x53, 0x33, 0x32, 0xa9, 0xe2,
	0xd6, 0x7e, 0xfd, 0xe9, 0x4d, 0xa3, 0x37, 0xdd, 0xcf, 0xfd, 0x7a, 0x94, 0x0c, 0x67, 0xa0, 0xd8,
	0x49, 0xc8, 0xe6, 0x38, 0x50, 0xc5, 0xbe, 0x54, 0x41, 0x15, 0x4b, 0x58, 0x02, 0x17, 0xb8, 0xb1,
	0x64, 0xb1, 0x0d, 0x47, 0x38, 0x70, 0x63, 0xfb, 0x03, 0x4c, 0x15, 0x50, 0x54, 0xf7, 0xfb, 0x46,
	0x7a, 0x23, 0x0d, 0x22, 0x27, 0x8d, 0xbe, 0xef, 0xd7, 0xdf, 0xbe, 0x3d, 0x76, 0x14, 0xad, 0x8d,
	0x8c, 0xc4, 0xcb, 0xa9, 0x35, 0xce, 0xf0, 0xa3, 0xd7, 0xd1, 0xba, 0xd1, 0x65, 0x22, 0x9e, 0x3d,
	0x11, 0x9b, 0xd8, 0x04, 0xce, 0xa2, 0xff, 0x95, 0x83, 0x2e, 0x7c, 0x80, 0xb1, 0x96, 0xb5, 0x4d,
	0x74, 0x42, 0x25, 0x19, 0x7f, 0x0f, 0x7b, 0xd0, 0x63, 0xb3, 0xd3, 0xa5, 0x85, 0x07, 0x9e, 0x38,
	0xf6, 0xbe, 0x93, 0x97, 0x27, 0x44, 0x5c, 0x6e, 0x59, 0xdb, 0x30, 0x12, 0xbb, 0x39, 0xe8, 0xd2,
	0x73, 0xe7, 0xd8, 0x0c, 0x91, 0xf8, 0x51, 0x36, 0x7b, 0x4d, 0x4b, 0xdc, 0x54, 0x1a, 0x25, 0xdc,
	0xc7, 0x67, 0xd9, 0x3b, 0x36, 0xd6, 0x9b, 0xeb, 0x70, 0xeb, 0x41, 0x7e, 0x92, 0xcd, 0xb7, 0xac,
	0x5d, 0x33, 0xae, 0x3d, 0x48, 0x13, 0x1c, 0xa0, 0x76, 0x28, 0xe1, 0x93, 0x87, 0x38, 0xb0, 0xc3,
	0x2d, 0x6b, 0xdb, 0xda, 0xa1, 0xd5, 0x22, 0x81, 0x7b, 0x87, 0xf8, 0x71, 0x36, 0x17, 0x28, 0xdb,
	0x22, 0x51, 0xb2, 0xad, 0xd3, 0xa1, 0x03, 0x39, 0x49, 0xec, 0x0a, 0x1d, 0x23, 0x20, 0x11, 0x57,
	0x55, 0x96, 0x29, 0x1d, 0xe7, 0xc8, 0x4d, 0x7e, 0x82, 0x41, 0xcb, 0xda, 0x1e, 0x5a, 0x25, 0x12,
	0xf5, 0x11, 0xe1, 0x94, 0xd1, 0x10, 0xf3, 0x93, 0x8c, 0x07, 0x07, 0xb3, 0x09, 0x7a, 0x9f, 0xcf,
	0xb3, 0xa3, 0x1e, 0xed, 0x2c, 0x8a, 0x41, 0x17, 0x85, 0x04, 0xc5, 0x39, 0x3b, 0xb6, 0x43, 0x7a,
	0xd2, 0x2a, 0x87, 0xf0, 0x61, 0x7a, 0x9e, 0xd3, 0x36, 0xac, 0xd0, 0xd9, 0xa6, 0xb1, 0x03, 0xd0,
	0xfc, 0x0c, 0x7b, 0x68, 0x87, 0xde, 0x43, 0x2d, 0x6b, 0x5a, 0x36, 0x12, 0x93, 0x21, 0x18, 0x7e,
	0x9a, 0x9d, 0xd8, 0x61, 0x5d, 0x41, 0x21, 0xd1, 0xe6, 0xc2, 0x52, 0x7e, 0x8a, 0x1d, 0xdf, 0xc3,
	0x09, 0x9a, 0xb3, 0x09, 0x63, 0x7a, 0x4a, 0x6f, 0xc1, 0x8d, 0x09, 0x05, 0x41, 0x72, 0x4d, 0xcb,
	0x2e, 0x46, 0xdb, 0x60, 0xc9, 0x51, 0xf2, 0x7e, 0x55, 0xa4, 0x57, 0x71, 0x04, 0x5b, 0xfc, 0x58,
	0x9e, 0xc9, 0x7a, 0xae, 0x2c, 0xf1, 0x19, 0x09, 0xff, 0x07, 0x15, 0x03, 0x7e, 0x9e, 0x9d, 0xf1,
	0xb9, 0xb2, 0xa3, 0xd4, 0x99, 0xae, 0xd0, 0xd2, 0x0c, 0x96, 0x51, 0xa3, 0xcd, 0xc3, 0xf1, 0xeb,
	0x12, 0x7f, 0x98, 0x9d, 0xdc, 0xe1, 0x5f, 0xc5, 0x51, 0x81, 0xf9, 0x9b, 0x12, 0x7f, 0x84, 0x9d,
	0xde, 0x61, 0xae, 0x19, 0x1d, 0x61, 0x81, 0xfd, 0xdb, 0x12, 0x3f, 0x15, 0x82, 0x94, 0xb3, 0x7b,
	0x2a, 0xd6, 0xc2, 0x0d, 0x2d, 0xc2, 0xef, 0x4a, 0xfc, 0x9d, 0xec, 0xfc, 0x7e, 0xc6, 0x07, 0xd1,
	0xaa, 0x4d, 0x15, 0xe5, 0xaf, 0x5f, 0x2f, 0xf1, 0x87, 0x82, 0x3b, 0x39, 0xa8, 0x89, 0x91, 0xff,
	0x0b, 0x6f, 0x94, 0xf8, 0x39, 0x76, 0x6a, 0x2f, 0xb9, 0x23, 0x46, 0x89, 0x11, 0x12, 0xde, 0x9c,
	0x7c, 0xd4, 0xd2, 0xf9, 0xa3, 0xb7, 0xf6, 0x79, 0xd1, 0x30, 0x7a, 0x1b, 0x6d, 0xe6, 0x15, 0xdd,
	0x2e, 0xf1, 0xd3, 0x21, 0xfc, 0x39, 0xb3, 0xa1, 0xd2, 0x3e, 0xda, 0xb6, 0x56, 0x0e, 0xee, 0xec,
	0x7b, 0xd6, 0x44, 0xab, 0xb6, 0x73, 0xfb, 0xee, 0x96, 0xf8, 0x61, 0x76, 0xc8, 0x87, 0x5b, 0xa4,
	0xf0, 0xdd, 0xfb, 0xf9, 0x5c, 0x88, 0xf2, 0x92, 0xb1, 0x2d, 0x11, 0xf5, 0xe1, 0x7b, 0xf7, 0xf3,
	0xe3, 0xa1, 0x68, 0xae, 0xe2, 0x28, 0x73, 0xc6, 0xe2, 0x32, 0x3a, 0x78, 0xe6, 0x81, 0x3d, 0xc4,
	0xce, 0xd0, 0xc1, 0xb3, 0x63, 0xe2, 0xba, 0xbd, 0xae, 0x5c, 0xb3, 0x1e, 0x34, 0xff, 0x65, 0x66,
	0x92, 0xb8, 0x9e, 0xa2, 0x86, 0xbf, 0xce, 0x90, 0x73, 0x44, 0xac, 0xa5, 0x29, 0x6a, 0x09, 0x7f,
	0x9b, 0xa1, 0x14, 0x12, 0x79, 0x6f, 0x45, 0xff, 0x7d, 0x86, 0xfc, 0x23, 0x7e, 0xcf, 0x6b, 0x6e,
	0x88, 0xcc, 0xc1, 0x3f, 0x66, 0xc8, 0xea, 0x76, 0x67, 0xa9, 0x57, 0x93, 0x12, 0x6e, 0x95, 0x0b,
	0x04, 0x6f, 0xf1, 0x2b, 0x65, 0x5e, 0x65, 0x17, 0x5a, 0xd6, 0x5e, 0x11, 0x5a, 0x66, 0x7d, 0xb1,
	0x85, 0xeb, 0x4f, 0xe9, 0x56, 0xda, 0xc7, 0x01, 0x5a, 0x91, 0xe4, 0xd5, 0xe0, 0x8b, 0x1c, 0x5e,
	0x2f, 0xf3, 0x8b, 0x6c, 0xa1, 0x08, 0xec, 0x20, 0xda, 0x22, 0x32, 0x94, 0xe8, 0x1b, 0x65, 0xbe,
	0xc8, 0x2e, 0x15, 0x61, 0x5d, 0xbc, 0x31, 0xc4, 0xcc, 0xa1, 0xad, 0x0d, 0x5d, 0x1f, 0xb5, 0xf3,
	0xe9, 0xc7, 0xba, 0x79, 0x3a, 0x97, 0x0d, 0x6f, 0x96, 0xf9, 0xbb, 0xd8, 0x63, 0x93, 0x0f, 0xb2,
	0xd4, 0x68, 0x89, 0xb6, 0x16, 0x45, 0x98, 0xba, 0x5d, 0xe8, 0x5b, 0x65, 0x5e, 0x61, 0x67, 0xa7,
	0xca, 0xbe, 0x82, 0x49, 0x62, 0xe0, 0xf6, 0x14, 0x00, 0xc9, 0xca, 0x01, 0x77, 0xca, 0xfc, 0x71,
	0xf6, 0xe8, 0xff, 0xb5, 0x0e, 0xee, 0x96, 0xf9, 0x02, 0x7b, 0xf8, 0x00, 0xa3, 0xe0, 0xf7, 0xfb,
	0xc2, 0xb1, 0x2b, 0x29, 0xda, 0xd2, 0xe6, 0xa9, 0x04, 0x65, 0x8c, 0xf0, 0x87, 0x32, 0x7f, 0x94,
	0x9d, 0x0b, 0x93, 0x52, 0x3b, 0x11, 0x39, 0x02, 0xf5, 0xc4, 0x00, 0x6b, 0x51, 0x64, 0x86, 0xda,
	0xc1, 0xf7, 0x67, 0x29, 0x00, 0x93, 0x10, 0xfa, 0xaf, 0x96, 0x58, 0x14, 0x72, 0x54, 0x93, 0x12,
	0x25, 0xfc, 0x60, 0x96, 0x3f, 0xc6, 0x2a, 0xff, 0x0b, 0x5a, 0x4f, 0x4c, 0xb4, 0x85, 0x12, 0x7e,
	0x38, 0x4b, 0x4e, 0x4e, 0x45, 0xed, 0x8e, 0xea, 0x1f, 0xcd, 0xf2, 0xf7, 0xb2, 0x27, 0xf6, 0xe1,
	0xda, 0x3a, 0x32, 0x03, 0xa5, 0x63, 0xd2, 0xdc, 0xc5, 0x08, 0xd5, 0x36, 0x4a, 0x78, 0x75, 0x96,
	0x82, 0xbb, 0x6c, 0xcd, 0x30, 0x5d, 0xc5, 0xc1, 0x75, 0xb4, 0x2b, 0x26, 0x6e, 0x6d, 0xa3, 0x76,
	0xa1, 0x7a, 0x9f, 0x67, 0x64, 0xdd, 0x14, 0xc0, 0xee, 0x68, 0x78, 0x81, 0x51, 0x44, 0x0a, 0xa8,
	0x6b, 0xda, 0x47, 0x4c, 0x07, 0x4a, 0xbb, 0x09, 0x2f, 0x32, 0x7e, 0x81, 0x3d, 0x32, 0x86, 0xf4,
	0x30, 0xb2, 0xe8, 0xd6, 0x5d, 0x1f, 0xfd, 0x28, 0x77, 0xf9, 0x0b, 0x78, 0x89, 0x91, 0x93, 0x05,
	0x0c, 0x59, 0xdc, 0x43, 0xed, 0x36, 0x0c, 0xe1, 0x6e, 0x32, 0xea, 0x8d, 0x5c, 0x78, 0xbe, 0x4b,
	0x36, 0x46, 0x29, 0xc2, 0xcb, 0x8c, 0x9f, 0x08, 0xbb, 0x24, 0x37, 0x24, 0x1f, 0xa9, 0x70, 0x8b,
	0x51, 0x0b, 0x06, 0x6a, 0x2d, 0x72, 0x7e, 0x1a, 0x20, 0xbc, 0xc2, 0x68, 0xd2, 0x05, 0x72, 0x13,
	0xc5, 0x98, 0xf1, 0x45, 0xc6, 0xe7, 0xd9, 0x91, 0x5d, 0xf9, 0x9b, 0x06, 0xbe, 0xc4, 0xf8, 0x99,
	0xb0, 0x07, 0x82, 0xe7, 0x2b, 0xca, 0xdb, 0xec, 0x84, 0x14, 0x4e, 0xc0, 0xbd, 0xb1, 0x35, 0x05,
	0x56, 0x96, 0x89, 0x18, 0xe1, 0x9f, 0x8c, 0x2a, 0x8e, 0x08, 0x57, 0x71, 0xd4, 0xf1, 0xe3, 0x2b,
	0x73, 0xa8, 0xa3, 0x30, 0x46, 0xfe, 0x78, 0xf8, 0x20, 0x84, 0xef, 0xe5, 0x3f, 0x1d, 0x26, 0xc5,
	0x75, 0xab, 0x64, 0x8c, 0x61, 0xbf, 0xda, 0x61, 0xea, 0x97, 0xee, 0x2f, 0x8f, 0x90, 0xe2, 0x9c,
	0xb5, 0x66, 0x5c, 0x77, 0xa8, 0xb5, 0x77, 0xf8, 0x57, 0x47, 0x0a, 0x62, 0x51, 0xc7, 0x38, 0x5e,
	0xb8, 0x4d, 0xc4, 0x74, 0xc5, 0x2f, 0xa4, 0xdb, 0x73, 0xfc, 0xdd, 0xec, 0xf1, 0x22, 0x62, 0xcc,
	0xf2, 0x15, 0xa3, 0x2c, 0x66, 0x1d, 0x91, 0x65, 0x69, 0xdf, 0x8a, 0x0c, 0xe1, 0xce, 0x1c, 0xbf,
	0xc4, 0x2e, 0x4e, 0x03, 0x93, 0xd8, 0x02, 0xf6, 0xee, 0x1c, 0x4d, 0xfa, 0x66, 0xbd, 0xa5, 0x9d,
	0x1d, 0x51, 0xfe, 0x5a, 0x4f, 0xab, 0xcc, 0x65, 0xf0, 0x2a, 0x90, 0x37, 0xcd, 0x7a, 0x4d, 0x4a,
	0x1a, 0xe7, 0xf9, 0xc0, 0xfb, 0x31, 0xd0, 0xd0, 0x1a, 0xb3, 0x0a, 0xe5, 0xbb, 0x3e, 0x74, 0xb1,
	0x51, 0x3a, 0xf6, 0x55, 0x00, 0x3f, 0x01, 0xb2, 0xe6, 0x00, 0x60, 0x4b, 0xdf, 0x18, 0xe2, 0x50,
	0xc2, 0x4f, 0x81, 0xdc, 0x9c, 0x82, 0x1d, 0xf7, 0xc4, 0x4e, 0x33, 0xfc, 0xec, 0x6d, 0x80, 0xf3,
	0x49, 0x81, 0x12, 0x5e, 0x03, 0xda, 0x86, 0x01, 0x5c, 0xa8, 0xfb, 0x26, 0x6e, 0xab, 0x08, 0xf3,
	0xde, 0xfe, 0xf9, 0xae, 0xbb, 0xab, 0xc3, 0xc4, 0xa9, 0x34, 0xc1, 0x2e, 0x46, 0xc6, 0xca, 0x0c,
	0x7e, 0x01, 0xd4, 0x0f, 0x5d, 0x4c, 0x13, 0x31, 0xea, 0x58, 0x13, 0x61, 0x96, 0x91, 0x1c, 0xaa,
	0xac, 0x9b, 0xf3, 0xd4, 0x56, 0xd3, 0x30, 0x79, 0x89, 0xbd, 0x3c, 0x4f, 0x1a, 0x6a, 0xce, 0x89,
	0xa8, 0xef, 0xcf, 0xb1, 0x8e, 0xc5, 0x54, 0x58, 0x84, 0x7f, 0xcd, 0xf3, 0xb3, 0xe1, 0xe8, 0xd8,
	0x65, 0x75, 0xd1, 0x59, 0x85, 0xdb, 0x08, 0xff, 0x9e, 0xa7, 0x3e, 0xe9, 0xf8, 0xab, 0x31, 0x32,
	0x49, 0x58, 0x07, 0xff, 0x99, 0x1f, 0xef, 0xe1, 0x95, 0xf6, 0x9a, 0xd9, 0x40, 0x3b, 0x88, 0x44,
	0x9a, 0xc1, 0x6b, 0xa7, 0x08, 0xdc, 0x43, 0xeb, 0x7d, 0xcb, 0xfc, 0x58, 0x85, 0x2f, 0x57, 0xc8,
	0x81, 0x22, 0xd5, 0x9f, 0x85, 0x5a, 0xb9, 0xb0, 0xc7, 0x50, 0xc2, 0x57, 0x2a, 0xb4, 0xe4, 0x8a,
	0x98, 0x27, 0xad, 0xd1, 0x71, 0xcf, 0xf9, 0x46, 0xfb, 0x6a, 0x85, 0x82, 0x58, 0xe4, 0x8f, 0x6f,
	0xc3, 0x30, 0x9c, 0x33, 0x84, 0xaf, 0x55, 0xc8, 0x8f, 0x22, 0xc8, 0xff, 0x46, 0x0b, 0x5f, 0xaf,
	0x50, 0x84, 0x8a, 0x3c, 0x7f, 0xc0, 0x36, 0xfa, 0x22, 0x49, 0x7c, 0xa9, 0xc2, 0x37, 0x2a, 0x54,
	0x57, 0x7b, 0x9f, 0xab, 0x08, 0xc7, 0xa3, 0xc3, 0x6c, 0xa1, 0x86, 0x6f, 0x1e, 0x00, 0x5c, 0x33,
	0xae, 0x37, 0x4c, 0x53, 0x63, 0x7d, 0xea, 0xbf, 0x55, 0xa1, 0xee, 0x2a, 0x02, 0x69, 0xdc, 0xe5,
	0xa2, 0xbe, 0x3d, 0xcd, 0x6f, 0x52, 0x76, 0xad, 0xbb, 0x02, 0xdf, 0xd9, 0xe3, 0x92, 0xcf, 0xef,
	0xf8, 0x82, 0xfa, 0xe8, 0xc2, 0xa4, 0xf4, 0x02, 0x8f, 0x9c, 0xfe, 0xd8, 0x02, 0x2d, 0x97, 0xfd,
	0x08, 0x9a, 0x78, 0x2d, 0x2d, 0x53, 0xa3, 0xb4, 0x83, 0x8f, 0x2f, 0x50, 0xea, 0xea, 0xfe, 0xf0,
	0x1f, 0x6f, 0xa7, 0x4f, 0x54, 0x29, 0x6a, 0x45, 0xea, 0x9a, 0x69, 0x37, 0x7b, 0x29, 0x46, 0x6a,
	0x53, 0xf9, 0x83, 0x7e, 0x1a, 0x84, 0x7a, 0xd9, 0x2f, 0x06, 0x94, 0xf0, 0xa9, 0x2a, 0x25, 0xaf,
	0x08, 0x19, 0x5f, 0xfb, 0xcd, 0x25, 0x63, 0x07, 0xc2, 0xc1, 0xa7, 0xa7, 0x81, 0x56, 0x4c, 0x1c,
	0x66, 0x48, 0x64, 0xac, 0x70, 0xc6, 0xc2, 0x67, 0xaa, 0x14, 0xae, 0x22, 0x68, 0xb9, 0xdb, 0x69,
	0x34, 0x12, 0xe5, 0x3b, 0xfe, 0xb3, 0x55, 0xda, 0x52, 0x45, 0xbe, 0xb7, 0x62, 0xec, 0xd0, 0xe7,
	0xaa, 0x14, 0xb3, 0x22, 0xa0, 0x29, 0x9c, 0x58, 0x33, 0x6e, 0xc9, 0x0c, 0xb5, 0x84, 0xcf, 0x57,
	0xa9, 0x5a, 0x8b, 0x88, 0x71, 0xa3, 0x5d, 0x4b, 0xa5, 0xaf, 0xc6, 0x67, 0xa6, 0xa9, 0x59, 0x15,
	0x5a, 0xc4, 0x68, 0xc3, 0x32, 0x7c, 0x76, 0x9a, 0x1a, 0x02, 0xe4, 0x5f, 0x0b, 0x5f, 0x98, 0xa6,
	0x86, 0x62, 0xd2, 0x58, 0x69, 0xd7, 0x6c, 0x9c, 0xc1, 0x73, 0x55, 0xba, 0x4f, 0x8b, 0x98, 0xa5,
	0x5e, 0xcb, 0x5a, 0x63, 0xe1, 0xf9, 0x03, 0xe2, 0x4e, 0x33, 0xf4, 0x85, 0x2a, 0x5d, 0x29, 0x93,
	0xd9, 0xab, 0x8b, 0x68, 0x6b, 0x98, 0xee, 0x66, 0xf0, 0xc5, 0x69, 0xde, 0xb4, 0x9b, 0xcb, 0xa8,
	0x97, 0x84, 0x4a, 0x50, 0xc2, 0x4b, 0xd3, 0x6c, 0x6d, 0x58, 0x0c, 0x85, 0x44, 0x98, 0x9b, 0xd5,
	0xfa, 0xc5, 0xdb, 0x7f, 0x3e, 0x7f, 0xdf, 0x87, 0x2a, 0xf9, 0x97, 0xa3, 0xc3, 0xa8, 0xbf, 0x18,
	0x7e, 0x2e, 0xc6, 0x66, 0x31, 0xdd, 0x8a, 0x17, 0xe9, 0x5b, 0xf2, 0xfa, 0xa1, 0xf0, 0xfd, 0xf9,
	0xfe, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0xac, 0x80, 0xf6, 0x21, 0xb5, 0x0e, 0x00, 0x00,
}
