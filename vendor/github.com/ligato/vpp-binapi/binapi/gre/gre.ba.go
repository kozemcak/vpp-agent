// Code generated by GoVPP binapi-generator. DO NOT EDIT.
//  source: vppapi/gre.api.json

/*
 Package gre is a generated from VPP binary API module 'gre'.

 It contains following objects:
	  4 messages
	  2 services

*/
package gre

import "git.fd.io/govpp.git/api"
import "github.com/lunixbochs/struc"
import "bytes"

// Reference imports to suppress errors if they are not otherwise used.
var _ = api.RegisterMessage
var _ = struc.Pack
var _ = bytes.NewBuffer

// VlAPIVersion represents version of the binary API module.
const VlAPIVersion = 0x83d2c14f

// Services represents VPP binary API services:
//
//	"services": {
//	    "gre_add_del_tunnel": {
//	        "reply": "gre_add_del_tunnel_reply"
//	    },
//	    "gre_tunnel_dump": {
//	        "reply": "gre_tunnel_details",
//	        "stream": true
//	    }
//	},
//
type Services interface {
	DumpGreTunnel(*GreTunnelDump) ([]*GreTunnelDetails, error)
	GreAddDelTunnel(*GreAddDelTunnel) (*GreAddDelTunnelReply, error)
}

/* Messages */

// GreAddDelTunnel represents VPP binary API message 'gre_add_del_tunnel':
//
//	"gre_add_del_tunnel",
//	[
//	    "u16",
//	    "_vl_msg_id"
//	],
//	[
//	    "u32",
//	    "client_index"
//	],
//	[
//	    "u32",
//	    "context"
//	],
//	[
//	    "u8",
//	    "is_add"
//	],
//	[
//	    "u8",
//	    "is_ipv6"
//	],
//	[
//	    "u8",
//	    "tunnel_type"
//	],
//	[
//	    "u32",
//	    "instance"
//	],
//	[
//	    "u8",
//	    "src_address",
//	    16
//	],
//	[
//	    "u8",
//	    "dst_address",
//	    16
//	],
//	[
//	    "u32",
//	    "outer_fib_id"
//	],
//	[
//	    "u16",
//	    "session_id"
//	],
//	{
//	    "crc": "0x9f03ede2"
//	}
//
type GreAddDelTunnel struct {
	IsAdd      uint8
	IsIPv6     uint8
	TunnelType uint8
	Instance   uint32
	SrcAddress []byte `struc:"[16]byte"`
	DstAddress []byte `struc:"[16]byte"`
	OuterFibID uint32
	SessionID  uint16
}

func (*GreAddDelTunnel) GetMessageName() string {
	return "gre_add_del_tunnel"
}
func (*GreAddDelTunnel) GetCrcString() string {
	return "9f03ede2"
}
func (*GreAddDelTunnel) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// GreAddDelTunnelReply represents VPP binary API message 'gre_add_del_tunnel_reply':
//
//	"gre_add_del_tunnel_reply",
//	[
//	    "u16",
//	    "_vl_msg_id"
//	],
//	[
//	    "u32",
//	    "context"
//	],
//	[
//	    "i32",
//	    "retval"
//	],
//	[
//	    "u32",
//	    "sw_if_index"
//	],
//	{
//	    "crc": "0xfda5941f"
//	}
//
type GreAddDelTunnelReply struct {
	Retval    int32
	SwIfIndex uint32
}

func (*GreAddDelTunnelReply) GetMessageName() string {
	return "gre_add_del_tunnel_reply"
}
func (*GreAddDelTunnelReply) GetCrcString() string {
	return "fda5941f"
}
func (*GreAddDelTunnelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// GreTunnelDump represents VPP binary API message 'gre_tunnel_dump':
//
//	"gre_tunnel_dump",
//	[
//	    "u16",
//	    "_vl_msg_id"
//	],
//	[
//	    "u32",
//	    "client_index"
//	],
//	[
//	    "u32",
//	    "context"
//	],
//	[
//	    "u32",
//	    "sw_if_index"
//	],
//	{
//	    "crc": "0x529cb13f"
//	}
//
type GreTunnelDump struct {
	SwIfIndex uint32
}

func (*GreTunnelDump) GetMessageName() string {
	return "gre_tunnel_dump"
}
func (*GreTunnelDump) GetCrcString() string {
	return "529cb13f"
}
func (*GreTunnelDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// GreTunnelDetails represents VPP binary API message 'gre_tunnel_details':
//
//	"gre_tunnel_details",
//	[
//	    "u16",
//	    "_vl_msg_id"
//	],
//	[
//	    "u32",
//	    "context"
//	],
//	[
//	    "u32",
//	    "sw_if_index"
//	],
//	[
//	    "u32",
//	    "instance"
//	],
//	[
//	    "u8",
//	    "is_ipv6"
//	],
//	[
//	    "u8",
//	    "tunnel_type"
//	],
//	[
//	    "u8",
//	    "src_address",
//	    16
//	],
//	[
//	    "u8",
//	    "dst_address",
//	    16
//	],
//	[
//	    "u32",
//	    "outer_fib_id"
//	],
//	[
//	    "u16",
//	    "session_id"
//	],
//	{
//	    "crc": "0x1a12b8c1"
//	}
//
type GreTunnelDetails struct {
	SwIfIndex  uint32
	Instance   uint32
	IsIPv6     uint8
	TunnelType uint8
	SrcAddress []byte `struc:"[16]byte"`
	DstAddress []byte `struc:"[16]byte"`
	OuterFibID uint32
	SessionID  uint16
}

func (*GreTunnelDetails) GetMessageName() string {
	return "gre_tunnel_details"
}
func (*GreTunnelDetails) GetCrcString() string {
	return "1a12b8c1"
}
func (*GreTunnelDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func init() {
	api.RegisterMessage((*GreAddDelTunnel)(nil), "gre.GreAddDelTunnel")
	api.RegisterMessage((*GreAddDelTunnelReply)(nil), "gre.GreAddDelTunnelReply")
	api.RegisterMessage((*GreTunnelDump)(nil), "gre.GreTunnelDump")
	api.RegisterMessage((*GreTunnelDetails)(nil), "gre.GreTunnelDetails")
}
