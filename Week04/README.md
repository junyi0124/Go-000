# 学习笔记

## Week04作业
1. 按照自己的构想，写一个项目满足基本的目录结构的工程，代码需要包含对数据层、业务层、API注册、以及main函数对于服务的注册和启动，信号处理，使用Wire构建依赖。可以使用自己熟悉的框架。
1. 截止到周五晚上23:59

## 作业
- 在homework目录

## go工程化 - 如何分层放置代码
### 概述
- 需要及时补充DDD的知识
- 如果只是PoC或toy项目，可以都放在main.go中（方便，省事）
- 当多人协作时则需要一个共用的目录结构, 需要尽早定义统一的结构。建议开发一个kit-tool,用于方便快速生成项目模板,统一目录布局.

### 目录
1. cmd
    - 项目主干，其他内容分开目录放置
    - 分项目管理，应该: /cmd/myapp/main.go 而不是: /cmd/myapp.go
    - 除非有必要,不添加额外代码.

1. internal
    - 私有库代码
    - 针对每个项目都应该新建一个对应的目录, 而不是直接将.go文件放在本目录下.
    - 如果需要调用不暴露的公共函数, 可以在internal目录下添加pkg目录.

1. pkg
	- 可公开的库代码
    - 该目录可参考go标准库的组织方式.(按功能划分目录)

1. kit - 工具集
    - 每个公司都应该为不同的微服务建立一个统一的kit工具包项目(基础库/框架)和app目录.
    - 方便管理，建议公司级只有一个, 如果有更好内容，可以合并进来
    - 不允许有vendor、3rd-party库. 让应用选择第三方包，而不是kit选择第三方.
    - 必要包需要依赖: grpc、proto
    - 可考虑封装插件或者fork代码的方式引入依赖
    - 特点: 统一、标准库方式布局(命名)、高度抽象、支持插件

1. api
    - API协议定义目录
    - 先把api文件安排进这个目录

1. configs
    - 配置文件(推荐yaml)
    - 需要像对待代码一样，对待生产系统的配置信息

1. test
    - 较大的项目应该需要测试数据子目录, 如: /test/data 或 /test/testdata
    - 可以通过添加文件前缀`.`或`_`用于屏蔽go编译. 增加灵活性.


## Ref
- [github wire](https://github.com/google/wire/blob/master/_tutorial/README.md)
- [wire tutorial](https://github.com/google/wire/blob/master/_tutorial/README.md)
- [wire guide](https://github.com/google/wire/blob/master/docs/guide.md)
- [wire best-practices](hthttps://github.com/google/wire/blob/master/docs/best-practices.md)