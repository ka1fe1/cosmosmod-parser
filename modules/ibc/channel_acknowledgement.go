package ibc

import (
	"fmt"
	cdc "github.com/kaifei-bianjie/common-parser/codec"
	. "github.com/kaifei-bianjie/common-parser/modules"
	"github.com/kaifei-bianjie/common-parser/utils"
	. "github.com/kaifei-bianjie/cosmosmod-parser/modules"
)

type DocMsgAcknowledgement struct {
	PacketId        string `bson:"packet_id"`
	Packet          Packet `bson:"packet"`
	Acknowledgement string `bson:"acknowledgement"`
	ProofAcked      string `bson:"proof_acked"`
	ProofHeight     Height `bson:"proof_height"`
	Signer          string `bson:"signer"`
}

func (m *DocMsgAcknowledgement) GetType() string {
	return MsgTypeAcknowledgement
}

func (m *DocMsgAcknowledgement) BuildMsg(v interface{}) {

	msg := v.(*MsgAcknowledgement)
	m.Signer = msg.Signer
	m.ProofHeight = loadHeight(msg.ProofHeight)
	m.Acknowledgement = UnmarshalAcknowledgement(msg.Acknowledgement)
	m.ProofAcked = utils.MarshalJsonIgnoreErr(msg.ProofAcked)
	m.Packet = loadPacket(msg.Packet)
	m.PacketId = fmt.Sprintf("%v%v%v%v%v", msg.Packet.SourcePort, msg.Packet.SourceChannel,
		msg.Packet.DestinationPort, msg.Packet.DestinationChannel, msg.Packet.Sequence)

}

func (m *DocMsgAcknowledgement) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var addrs []string

	msg := v.(*MsgAcknowledgement)
	addrs = append(addrs, msg.Signer)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}

func UnmarshalAcknowledgement(bytesdata []byte) string {
	var result Acknowledgement
	cdc.GetMarshaler().MustUnmarshalJSON(bytesdata, &result)
	return result.String()
}
