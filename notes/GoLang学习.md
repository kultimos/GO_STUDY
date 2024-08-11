# GoLang学习

  ## 通过main函数初见语法注意点
  [](/main.go)
  
  ## 变量声明的方式
  - 包括全局变量如何声明、普通变量如何声明以及如何一次性声明多个普通变量
  [](/study/var/testVar.go)

  ## const全局常量的定义和iota的使用
  [](/study/const_iota/testConstIota.go)

  ## 函数多返回值的写法
  [](/study/return/testReturn.go)
  
  ## 几种导包方式
  - _ package,匿名导包,会在当前main方法调用时,调用导入包的init方法
  - name package,别名导包,可以通过别名.method()进行方法的调用
  - . package,该方式可以直接调用其method(),可读性不好,不建议
  - [](/study/import_init/init_package)

  ## GoLang中的指针
  [](/study/point/testPoint.go)
  
  ## defer关键字的使用
  - defer的执行是在当前function执行完毕后才会调用
  - 多个defer之间的执行顺序等同于入栈的顺序,先声明的后执行,后声明的先执行
  [](/study/defer/testDefer.go)

  ## slice可变长数组学习
  - 定长数组(不推荐) [](/study/slice/testNoSlice.go)
  - 可变长数组即切片 [](/study/slice/testSlice.go)
  - 几个核心的方法 [](/study/slice/sliceAppendSplit.go)
    - make([]int, 3, 5), 通过指定slice长度和容量的方式来创建一个数组
    - slice = append(slice, 162), 在一个数组末尾添加元素
    - s1 := slice[2:4], 截取slice索引2,3这一段数据给s1,但需要注意,此时s1和slice指向相同的内存地址,对二者任意一个的修改都会影响到对象
    - copy(s1, slice), 将slice的数据拷贝给s1,深拷贝,此后两个数组将毫无关联