package prototype
// 通过复制一个对象来创建一个对象

//Cloneable 是原型对象需要实现的接口
type Cloneable interface {
	// 复制 方法
	Clone() Cloneable
}

// 定义结构体，实现接口
type PrototypeManager struct {
	prototypes map[string]Cloneable
}
// 构造函数
func NewPrototypeManager() *PrototypeManager {
	return &PrototypeManager{
		prototypes: make(map[string]Cloneable),
	}
}

func (p *PrototypeManager) Get(name string) Cloneable {
	return p.prototypes[name].Clone()
}

func (p *PrototypeManager) Set(name string, prototype Cloneable) {
	p.prototypes[name] = prototype
}
