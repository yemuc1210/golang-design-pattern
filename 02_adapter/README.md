# 适配器模式

适配器模式用于转换一种接口适配另一种接口。

实际使用中Adaptee一般为接口，并且使用工厂函数生成实例。

在Adapter中匿名组合Adaptee接口，所以Adapter类也拥有SpecificRequest实例方法，又因为Go语言中非入侵式接口特征，其实Adapter也适配Adaptee接口。

```
// 将adaptee 转换为target
// adapter 实现 target接口，那么Target实例可以引用 adapter实例
// 又因为 adaptee 为 adapter内部的属性，因此Target可以调用target的方法，完成适配
```

