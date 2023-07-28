package znet

type Message struct {
	Id      uint32
	DataLen uint32
	Data    []byte
}

func (msg *Message) GetMsgId() uint32 {
	return msg.Id
}

// 获取消息长度
func (msg *Message) GetMsgLen() uint32 {
	return msg.DataLen
}

// 获取消息数据
func (msg *Message) GetData() []byte {
	return msg.Data
}

// SET
func (msg *Message) SetMsgId(id uint32) {
	msg.Id = id
}
func (msg *Message) SetMsgLen(len uint32) {
	msg.DataLen = len
}
func (msg *Message) SetData(data []byte) {
	msg.Data = data
}
