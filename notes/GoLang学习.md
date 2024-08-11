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