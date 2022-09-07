package facade

import "fmt"
// 用一个对象操作其他对象，用户只关心这一个对象即可
// 为子系统的一组接口提供一个统一的入口。
// 外观模式定义了一个高层接口，这个接口使得子系统更加容易使用

// API 为facade 模块的外观接口，大部分代码使用此接口简化对facade类的访问。
type API interface {
	Test() string
}

func NewAPI() API {
	//返回的是实例指针 ？
	return &apiImpl{
		a: NewAModuleAPI(),
		b: NewBModuleAPI(),
	}
}



//facade implement   外观接口的实现
type apiImpl struct {
	a AModuleAPI    // 使用apiImpl的实例，可以操纵AmoduleAPI 和 BModuleAPI的实例
	b BModuleAPI
}

func (a *apiImpl) Test() string {
	aRet := a.a.TestA()  // 调用其他实例的方法
	bRet := a.b.TestB()
	return fmt.Sprintf("%s\n%s", aRet, bRet)
}

//NewAModuleAPI return new AModuleAPI
func NewAModuleAPI() AModuleAPI {
	return &aModuleImpl{}
}

//AModuleAPI ...
type AModuleAPI interface {
	TestA() string
}

type aModuleImpl struct{}

func (*aModuleImpl) TestA() string {
	return "A module running"
}

//NewBModuleAPI return new BModuleAPI
func NewBModuleAPI() BModuleAPI {
	return &bModuleImpl{}
}

//BModuleAPI ...
type BModuleAPI interface {
	TestB() string
}

type bModuleImpl struct{}

func (*bModuleImpl) TestB() string {
	return "B module running"
}
