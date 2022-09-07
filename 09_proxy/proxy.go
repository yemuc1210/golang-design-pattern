package proxy
//代理模式用于延迟处理操作或者在进行实际操作前后进行其它处理。
// 用一个对象操作另一个对象，并提供相同的 API
// 在原操作基础上，提供新的定制化操作

type Subject interface {
	Do() string
}

type RealSubject struct{}

func (RealSubject) Do() string {
	return "real"
}

// 代理类
type Proxy struct {
	// 进行封装，实现代理
	real RealSubject
}

func (p Proxy) Do() string {
	var res string

	// 在调用真实对象之前的工作，检查缓存，判断权限，实例化真实对象等。。
	res += "pre:"

	// 调用真实对象
	res += p.real.Do()

	// 调用之后的操作，如缓存结果，对结果进行处理等。。
	res += ":after"

	return res
}
