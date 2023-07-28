package znet

import "Zinxproject/zinx/ziface"

// 实现Router时，先嵌入这个BaseRouter基类，然后根据需求对这个基类进行重写

type BaseRouter struct {
}

// 抽一层BaseRouter出来的好处就是，有的Route可能用不到这三个所有的方法，那么抽出来一层，就可以继承BaseRouter来重写真正所需要的方法

// 在处理conn业务之前的钩子方法hook
func (br *BaseRouter) PreHandler(request ziface.IRequest) {

}

// 在处理conn业务的钩子方法hook
func (br *BaseRouter) Handler(request ziface.IRequest) {

}

// 在处理conn业务之后的钩子方法hook
func (br *BaseRouter) PostHandler(request ziface.IRequest) {

}
