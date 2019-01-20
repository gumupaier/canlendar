## 仿google任务日历开发
采用beego框架，前端使用vue

#### 首次运行需执行如下命令 创建网络(docker开发暂时忽略，文件未修改)
```
docker network create app-calendar
``` 

#### 模块安装
请移步script目录下执行.sh文件


#### 项目运行（直接运行本条即可）
```
 bee run -gendoc=true -downdoc=true
```
如果不适用api文档直接运行
```cgo
bee run
```























#### 小计
* 关于json输出字段大写问题，可以在定义结构体时加小标签
* 传入参数可以通过自定义结构体传输，但有个小问题是swagger竟然没有相应的自动化填充生成，这应该是个缺陷
* post or put 方法如果传入参数在payload里，在beego中使用GetString或者GetInt是起不了作用的，可以通过序列化body再进行取值，参考上一条
* 关于数据输出，在gayhub上没有找到相关代码，所以每次输出响应格式我还得预先设定一个结构体或者多个，这方式感觉着实恶心，以后发现新方式再更新
* 槽点，建model时结构体可以塞入一个另一个结构体指针，查询返回的时候只给了一个id，其他字段不给，如在mapping的model里存有task指针,返回时虽然给了一个Task结构体，但里面除了id为正常值，其他均为对应类型零值,怀疑是我姿势不对？，但并没找到解决方式，暂且记下，那我就查询多次吧
* 启用调试模式  这样可以在console里看到sql语句
```cgo
orm.Debug=true
``` 
* orm一对一关系中，__方式查询不出所对应的字段值，不是报错，而是为空





