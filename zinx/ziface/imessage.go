package ziface

// 将请求的消息封装到一个Message中，定义抽象接口
type IMessage interface {
	// 获取消息id
	GetMsgId() uint32
	// 获取消息长度
	GetMsgLen() uint32
	// 获取消息数据
	GetData() []byte

	//SET
	SetMsgId(uint32)
	SetMsgLen(uint32)
	SetData([]byte)
}
