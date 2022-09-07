package singleton

import "sync"
// 单例模式，一个类只能有一个实例

// Singleton 是单例模式接口，导出的
// 通过该接口可以避免 GetInstance 返回一个包私有类型的指针
type Singleton interface {
	foo()
}

// singleton 是单例模式类，包私有的
type singleton struct{}
// 实现foo方法，实现Singleton接口
func (s singleton) foo() {}
//使用懒惰模式的单例模式，使用双重检查加锁保证线程安全
var (
	instance *singleton
	once     sync.Once
)

//GetInstance 用于获取单例模式对象
func GetInstance() Singleton {
	once.Do(func() {  // Once，只会执行一次
		instance = &singleton{}
	})

	return instance
}
