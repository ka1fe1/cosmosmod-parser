package ibc

import (
	"fmt"
	. "github.com/kaifei-bianjie/common-parser/modules"
	"github.com/kaifei-bianjie/common-parser/utils"
	. "github.com/kaifei-bianjie/cosmosmod-parser/modules"
)

type DocMsgTimeout struct {
	PacketId         string `bson:"packet_id"`
	Packet           Packet `bson:"packet"`
	ProofUnreceived  string `bson:"proof_unreceived"`
	ProofHeight      Height `bson:"proof_height"`
	NextSequenceRecv int64  `bson:"next_sequence_recv"`
	Signer           string `bson:"signer"`
}

func (m *DocMsgTimeout) GetType() string {
	return MsgTypeTimeout
}

func (m *DocMsgTimeout) BuildMsg(v interface{}) {
	msg := v.(*MsgTimeout)
	m.Signer = msg.Signer
	m.NextSequenceRecv = int64(msg.NextSequenceRecv)
	m.ProofUnreceived = utils.MarshalJsonIgnoreErr(msg.ProofUnreceived)
	m.Packet = loadPacket(msg.Packet)
	m.ProofHeight = loadHeight(msg.ProofHeight)
	m.PacketId = fmt.Sprintf("%v%v%v%v%v", msg.Packet.SourcePort, msg.Packet.SourceChannel,
		msg.Packet.DestinationPort, msg.Packet.DestinationChannel, msg.Packet.Sequence)
}

func (m *DocMsgTimeout) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var addrs []string

	msg := v.(*MsgTimeout)
	addrs = append(addrs, msg.Signer)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
