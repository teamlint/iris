# 更新记录 <a href="HISTORY.md"> <img width="20px" src="https://iris-go.com/images/flag-unitedkingdom.svg?v=10" /></a> <a href="HISTORY_GR.md"> <img width="20px" src="https://iris-go.com/images/flag-greece.svg?v=10" /></a>

### 想得到免费即时的支持?

    https://github.com/teamlint/iris/issues
    https://chat.iris-go.com

### 获取历史版本?

    https://github.com/teamlint/iris/releases

### 我是否应该升级 Iris?

如果没有必要，不会强制升级。如果你已经准备好了，可以随时升级。

> Iris 使用 Golang 的 [vendor directory](https://docs.google.com/document/d/1Bz5-UB7g2uPBdOx-rw5t9MxJwkfpx90cqG9AFL0JAYo) 特性, 避免依赖包的更改带来影响。

**如何升级**: 打开命令行执行以下命令: `go get -u github.com/teamlint/iris` 或者等待自动更新。

# Mo, 15 Jenuary 2018 | v10.0.1

Translation is missing for this specific history entry, please navigate through [english version of HISTORY.md](HISTORY.md#mo-15-jenuary-2018--v1001) instead or check back later on.

# 2018 元旦 | v10.0.0 版本发布

我们必须感谢 [Mrs. Diana](https://www.instagram.com/merry.dii/) 帮我们绘制的漂亮 [logo](https://iris-go.com/images/icon.svg)!

如果有设计相关的需求，你可以[发邮件](mailto:Kovalenkodiana8@gmail.com)给他，或者通过 [instagram](https://www.instagram.com/merry.dii/) 给他发信息。

<p align="center">
<img width="145px" src="https://iris-go.com/images/icon.svg?v=a" />
</p>

在这个版本中，有许多内部优化改进，但只有两个重大变更和新增一个叫做 **hero** 的特性。

> 新版本有 75 + 的变更提交, 如果你需要升级 Iris 请仔细阅读本文档。 为什么版本 9 跳过了? 你猜...

## Hero 特性

新增包 [hero](hero) 可以绑定处理任何依赖 `handlers` 的对象或函数。Hero funcs 可以返回任何类型的值，并发送给客户端。

> 之前的绑定没有编辑器的支持, 新包 `hero` 为 Iris 带来真正的安全绑定。 Iris 会在服务器运行之前计算所有内容，所以它执行速度高，接近于原生性能。

下面你会看到我们为你准备的一些截图，以便于理解:

### 1. 路径参数 - 构建依赖

![](https://github.com/teamlint/explore/raw/master/iris/hero/hero-1-monokai.png)

### 2. 服务 - 静态依赖

![](https://github.com/teamlint/explore/raw/master/iris/hero/hero-2-monokai.png)

### 3. 请求之前 - 动态依赖

![](https://github.com/teamlint/explore/raw/master/iris/hero/hero-3-monokai.png)

`hero funcs` 非常容易理解，当你用过之后 **在也回不去了**.

示例：

- [基本用法](_examples/hero/basic/main.go)
- [使用概览](_examples/hero/overview)

## MVC

You have to understand the `hero` package in order to use the `mvc`, because `mvc` uses the `hero` internally for the controller's methods you use as routes, the same rules applied to those controller's methods of yours as well.

With this version you can register **any controller's methods as routes manually**, you can **get a route based on a method name and change its `Name` (useful for reverse routing inside templates)**, you can use any **dependencies** registered from `hero.Register` or `mvc.New(iris.Party).Register` per mvc application or per-controller, **you can still use `BeginRequest` and `EndRequest`**, you can catch **`BeforeActivation(b mvc.BeforeActivation)` to add dependencies per controller and `AfterActivation(a mvc.AfterActivation)` to make any post-validations**, **singleton controllers when no dynamic dependencies are used**, **Websocket controller, as simple as a `websocket.Connection` dependency** and more...

示例:

**如果你之前使用过 MVC ，请仔细阅读：MVC 包含一些破坏性的改进，但新的方式可以做更多，会让程序执行更快**

**请阅读我们为你准备的示例**

如果你现在需要升级，请对比新旧版本示例的不同，便于理解。

| NEW | OLD |
| -----------|-------------|
| [Hello world](_examples/mvc/hello-world/main.go) | [OLD Hello world](https://github.com/teamlint/iris/blob/v8/_examples/mvc/hello-world/main.go) |
| [Session Controller](_examples/mvc/session-controller/main.go) | [OLD Session Controller](https://github.com/teamlint/iris/blob/v8/_examples/mvc/session-controller/main.go) |
| [Overview - Plus Repository and Service layers](_examples/mvc/overview) | [OLD Overview - Plus Repository and Service layers](https://github.com/teamlint/iris/tree/v8/_examples/mvc/overview) |
| [Login showcase - Plus Repository and Service layers](_examples/mvc/login) | [OLD Login showcase - Plus Repository and Service layers](https://github.com/teamlint/iris/tree/v8/_examples/mvc/login) |
| [Singleton](_examples/mvc/singleton) |  **新增** |
| [Websocket Controller](_examples/mvc/websocket) |  **新增** |
| [Vue.js Todo MVC](_examples/tutorial/vuejs-todo-mvc) |  **新增** |

## context#PostMaxMemory

移除旧版本的常量 `context.DefaultMaxMemory` 替换为配置 `WithPostMaxMemory` 方法.

```go
// WithPostMaxMemory sets the maximum post data size
// that a client can send to the server, this differs
// from the overral request body size which can be modified
// by the `context#SetMaxRequestBodySize` or `iris#LimitRequestBodySize`.
//
// 默认值为 32MB 或者 32 << 20
func WithPostMaxMemory(limit int64) Configurator
```

如果你使用老版本的常量，你需要更改一行代码.

使用方式：

```go
import "github.com/teamlint/iris"

func main() {
    app := iris.New()
    // [...]

    app.Run(iris.Addr(":8080"), iris.WithPostMaxMemory(10 << 20))
}
```

## context#UploadFormFiles

新方法可以多文件上传, 应用于常见的上传操作, 它是一个非常有用的函数。

```go
// UploadFormFiles uploads any received file(s) from the client
// to the system physical location "destDirectory".
//
// The second optional argument "before" gives caller the chance to
// modify the *miltipart.FileHeader before saving to the disk,
// it can be used to change a file's name based on the current request,
// all FileHeader's options can be changed. You can ignore it if
// you don't need to use this capability before saving a file to the disk.
//
// Note that it doesn't check if request body streamed.
//
// Returns the copied length as int64 and
// a not nil error if at least one new file
// can't be created due to the operating system's permissions or
// http.ErrMissingFile if no file received.
//
// If you want to receive & accept files and manage them manually you can use the `context#FormFile`
// instead and create a copy function that suits your needs, the below is for generic usage.
//
// The default form's memory maximum size is 32MB, it can be changed by the
//  `iris#WithPostMaxMemory` configurator at main configuration passed on `app.Run`'s second argument.
//
// See `FormFile` to a more controlled to receive a file.
func (ctx *context) UploadFormFiles(
        destDirectory string,
        before ...func(string, string),
    ) (int64, error)
```

这里是相关示例 [here](_examples/http_request/upload-files/main.go).

## context#View

这里有个小更新，增加可选的第二个参数，用来绑定模版变量。提示：这种绑定方式，会忽略其他变量的绑定。
如果要忽略其他模版变量，之前是在 `ViewData` 上绑定一个空字符串，现在可以直接通过 View 方法添加。

```go
func(ctx iris.Context) {
    ctx.ViewData("", myItem{Name: "iris" })
    ctx.View("item.html")
}
```

等同于：

```go
func(ctx iris.Context) {
    ctx.View("item.html", myItem{Name: "iris" })
}
```

```html
html 模版中调用: {{.Name}}
```

## context#YAML

新增 `context#YAML` 函数, 解析结构体到 yaml。

```go
//使用 yaml 包的 Marshal 的方法解析，并发送到客户端。
func YAML(v interface{}) (int, error)
```

## Session#GetString

`sessions/session#GetString` 可以获取 session 的变量值（可以是 integer 类型），就像内存缓存、Context 上下文储存的值。
