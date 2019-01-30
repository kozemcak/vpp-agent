// Code generated by GoVPP binapi-generator. DO NOT EDIT.
//  source: vppapi/ioam_cache.api.json

/*
 Package ioam_cache is a generated from VPP binary API module 'ioam_cache'.

 It contains following objects:
	  2 messages
	  1 service

*/
package ioam_cache

import "git.fd.io/govpp.git/api"
import "github.com/lunixbochs/struc"
import "bytes"

// Reference imports to suppress errors if they are not otherwise used.
var _ = api.RegisterMessage
var _ = struc.Pack
var _ = bytes.NewBuffer

// VlAPIVersion represents version of the binary API module.
const VlAPIVersion = 0xb7452f41

// Services represents VPP binary API services:
//
//	"services": {
//	    "ioam_cache_ip6_enable_disable": {
//	        "reply": "ioam_cache_ip6_enable_disable_reply"
//	    }
//	},
//
type Services interface {
	IoamCacheIP6EnableDisable(*IoamCacheIP6EnableDisable) (*IoamCacheIP6EnableDisableReply, error)
}

/* Messages */

// IoamCacheIP6EnableDisable represents VPP binary API message 'ioam_cache_ip6_enable_disable':
//
//	"ioam_cache_ip6_enable_disable",
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
//	    "is_disable"
//	],
//	{
//	    "crc": "0x22324d89"
//	}
//
type IoamCacheIP6EnableDisable struct {
	IsDisable uint8
}

func (*IoamCacheIP6EnableDisable) GetMessageName() string {
	return "ioam_cache_ip6_enable_disable"
}
func (*IoamCacheIP6EnableDisable) GetCrcString() string {
	return "22324d89"
}
func (*IoamCacheIP6EnableDisable) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// IoamCacheIP6EnableDisableReply represents VPP binary API message 'ioam_cache_ip6_enable_disable_reply':
//
//	"ioam_cache_ip6_enable_disable_reply",
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
//	{
//	    "crc": "0xe8d4e804"
//	}
//
type IoamCacheIP6EnableDisableReply struct {
	Retval int32
}

func (*IoamCacheIP6EnableDisableReply) GetMessageName() string {
	return "ioam_cache_ip6_enable_disable_reply"
}
func (*IoamCacheIP6EnableDisableReply) GetCrcString() string {
	return "e8d4e804"
}
func (*IoamCacheIP6EnableDisableReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func init() {
	api.RegisterMessage((*IoamCacheIP6EnableDisable)(nil), "ioam_cache.IoamCacheIP6EnableDisable")
	api.RegisterMessage((*IoamCacheIP6EnableDisableReply)(nil), "ioam_cache.IoamCacheIP6EnableDisableReply")
}
