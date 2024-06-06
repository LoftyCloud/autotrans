# **AutoTrans**项目开发记录

## **1.** **技术栈**

**1.1 gin**：提供支持前后端分离的**web框架**

gin web 框架文档：https://gin-gonic.com/zh-cn/docs/introduction/

**1.2 goini**：项目参数管理

一般程序启动的时候需要通过静态配置文件加载各种配置，go-zero 目前支持以下四种后缀名的配置文件：

  *.json

  *.toml

  *.yaml

  *.yml

本项目没有采用上述配置方式，而是采用了goini加载项目参数配置，其可以结构化地读取ini文件，一个ini文件中包括了**分区，键，值**以及注释。

goini说明文档：https://ini.unknwon.io/docs/intro/getting_started

**1.3 gorm**：数据库操作

gorm通过将go结构体映射到数据库表来简化数据库交互。使用过程中需要为数据库表声明对应的go结构体。

1）模型（model）：gorm将数据库表映射为go语言的结构体。每个结构体对应数据库中的一个表，每个结构体的字段对应表中的列。

2）迁移（migration）：gorm可以自动生成数据库表和字段，基于模型结构体。使用AutoMigrate方法，gorm会检查并自动更新数据库表结构。

3）CRUD 操作：gorm 提供了简洁的API用于创建（Create）、读取（Read）、更新（Update）和删除（Delete）数据库记录。

gorm说明文档：https://gorm.io/zh_CN/docs/

**1.4 jwt**（****JSON Web Tokens****）：用户身份验证

验证的作用：

1、传统做法：基于cookie和session的方式，但是不支持分布式环境，需要解决分布式session共享的问题（比如利用redis等）

当用户通过了安全认证后，则在服务端的session对象中保存该用户的认证信息，这样该用户对服务的访问被认为是安全的。这种模式的最大问题是没有分布式架构，不方便进行横向扩展，这种模式只适合于单体应用模式。如果需要进行服务集群则需要处理好共享session的问题。 如果一个庞大的系统需要按服务分解为多个独立的服务，使用分布式架构，则这种方式更难处理。使用jwt可以方便的处理上面提到的问题。

执行流程：

1）客户端：发送用户名密码

2）服务器：验证成功后，创建并保留一个跟客户端相关联的session，返回session_ID给客户端

3）客户端：将session_ID保存到**cookie**中

4）客户端：后续的请求都带上包含了session_ID的cookie

5）服务器：通过验证cookie中的session_ID确认用户身份

缺点：默认不支持非浏览器环境（没有cookie机制的环境）

2、基于token的方式：适用于前后端分离以及分布式的大环境。

执行流程：

1）客户端：发送用户名密码

2）服务器：验证成功后，生成一个跟客户端相关联的token，返回token给客户端

3）客户端：存储token到**本地**

4）客户端：后续的请求都带上token

5）服务器：通过验证token确认用户身份

JWT (Json Web Token) 是一种用于前后端身份认证的方法，一个 JWT 由 header、payload、signature 组成。

\-  header: 包含了 token 类型和算法类型

\-  payload: 包含了一些用户自定义或 jwt 预定义的一些数据，每一个数据叫一个 claim（**注意不要将敏感信息放入**）

\-  signature: 将 header 和 payload 经过 base64 编码，加上一个 secret 密钥，整体经过 header 中的算法加密后生成。

https://pkg.go.dev/github.com/golang-jwt/jwt/v5#section-readme

## **2** **开发过程**

**2.1** **项目依赖**

命令行运行以下命令：

go get -u github.com/gin-gonic/gin

go get gopkg.in/ini.v1

go get -u gorm.io/gorm

go get -u gorm.io/driver/sqlite

go get -u github.com/golang-jwt/jwt

**2.2** **项目结构**

.

├── api // 控制器接口

│  └── v1 // 另起文件夹便于版本控制

│    ├── login.go // 用户登录

│    └── user.go // 用户模块接口，更多模块的接口和功能有待实现

├── config // 项目配置

│  └── config.ini // 使用goini管理项目参数配置

├── middleware // 中间件

│  └── jwt.go // 使用jwt进行用户身份验证等

├── model // 管理数据库，建立数据表结构体模型（gorm）

│  ├── db.go

│  ├── EmptyBoxQuantity.go

│  ├── MaterialTransportRecord.go

│  ├── Point.go

│  └── User.go // 用户数据表模型，并实现对数据表进行增删改查的具体操作，被接口调用

├── routes // 路由接口（gin提供）

│  └── router.go

├── upload // 托管静态资源

├── utils // 公共工具，全局功能：提供了错误处理，项目配置读取功能

│  └── errmsg // 错误处理另起包（非必要）

├── web // 托管前端页面（暂无）

├── .gitignore

├── go.mod // go模组管理

├── go.sum

├── main.go // 运行main.go以启动项目

└── README.md

**2.3** **拟实现项目功能**

**2.3.1** **用户管理**

·   用户登录、登出功能，保障系统安全性；

·   用户增加、删除、修改功能，管理系统用户信息；

·   用户权限设置，限制不同用户的访问权限。

**2.3.2** **数据管理**

·   读取物料运输记录，包括满载样框的运输记录；

·   读取各物料装卸点空料框数量记录，实时更新系统中的空料框数量；

·   对敏感信息进行加密存储，保障数据安全性。

**2.3.3** **统计功能**

每次统计时生成在对应的表格内生成一条记录：

·   对小车行驶信息进行统计，包括电量、车速、行驶路程、满载/空载等信息；

·   统计任务信息，包括进行中任务以及代办任务；

·   统计生产信息，如总运量；

·   统计过程信息，如任务执行进度、小车位置信息等。

**2.4** **项目详细开发流程**

1、数据库设置（mysql）

1）新建一个mysql数据库；

2）新建一个用户管理该数据库（区别于AutoTrans的系统/项目用户）；项目开发过程中一般会使用navicat等软件管理数据库而不是直接用命令行，其提供的可视化界面可以方便操作。新建用户后【选择数据库】，选择【添加权限】，如图1。也可直接向数据库管理人员询问数据库名、用户名称密码、ip等信息。

![图1 新建数据库用户（区别于系统用户）](.\image\clip_image002.jpg)

创建数据库后不需要手动新建数据库表格，使用gorm进行**迁移**时会自动创建。

2、项目参数配置（goini）：使用goini管理参数配置，使得后期的修改查阅更加方便。

1）首先在config/config.ini文件中配置好正确的项目参数，包括服务器端口号，数据库设置等等。

2）其次在utils/setting.go中声明全局变量，随后在init()方法（utils包初始化时会执行init()初始化方法）中**使用****goini****读取****config.ini****中的参数**。

3）最后在项目的其他地方需要使用时直接进行调用即可。

3、路由入口（gin）

在routes/router.go文件中编写路由入口，在main.go中进行调用以初始化路由。路由入口的具体写法参见gin框架的：

1）路由参数：https://gin-gonic.com/zh-cn/docs/examples/param-in-path/

2）路由组：https://gin-gonic.com/zh-cn/docs/examples/grouping-routes/

3）http方法：https://gin-gonic.com/zh-cn/docs/examples/http-method/

4、读写数据库（gorm）

1）在model/db.go文件中设置数据库连接，之后需要在main.go中调用该方法，参见gorm连接数据库：https://gorm.io/zh_CN/docs/connecting_to_the_database.html

2）在model文件夹下为数据表分别建立对应的结构体模型，参见gorm模型：

https://gorm.io/zh_CN/docs/models.html

值得注意的是，gorm.Model是gorm预定义的结构体，包括了id、创建时间等等基本信息（如图2表中的前四项），因此在声明数据表结构体时可以免去相应字段的声明。

3）在model/db.go中配置自动迁移，需要分别传入每个表结构体对应的地址，参见gorm迁移：

https://gorm.io/zh_CN/docs/migration.html

4）在main.go中调用数据库初始化函数，此时若数据库与表设计不一致会进行自动迁移，打开数据库后运行项目即可创建对应的表。自动创建的用户管理功能与数据管理功能的数据表关系如图2。

![图2 gorm迁移创建的数据表](.\image\clip_image004.jpg)

5、接口与错误处理（gin，gorm）

1）错误处理：分别声明常量错误码（int类型）以及错误消息字典（map类型），根据错误码获取对应的错误消息，便于前端给用户返回提示信息。由于整个项目均需使用到错误处理，因此将其放在utils文件夹下。

2）**功能开发：**首先在routes/router.go中为需要实现的功能绑定路由接口。接着在model/文件夹中对应的模型处编写数据库操作的具体方法。最后在api/v1中编写接口逻辑，调用方法以完成功能实现。

其中，接口主要是负责接收数据、逻辑处理、调用方法、返回结果。功能的具体实现（这里指的是增删改查等针对数据库的操作）具体是在model/文件夹中的数据表模型中完成的。

总的业务流程：用户=>前端=>路由（routes/）=>接口（api/v1）=>方法（model/）=>数据库。

6、功能测试：配置好路由接口后，重新启动项目，此时可用的接口如图3所示。

![图3 配置的路由接口](.\image\clip_image006.jpg)

这里使用apipost软件（最常用的是postman软件和swagger包）进行接口测试，如接口测试文件所示。配置对应的方法和路径后（这里相当于模拟前端响应用户的操作向后端发送请求），【发送】响应，其中若包含了json数据需要注意数据的编写是否符合规范（在应用中由前端负责将数据打包成json格式）。

注：使用swagger自动化生成api文档的方法：

go install github.com/swaggo/swag/cmd/swag@latest

cd server

swag init

执行上面的命令后，server目录下会出现docs文件夹以及docs.go, swagger.json, swagger.yaml三个文件更新，启动go服务之后，在浏览器输入http://localhost:8888/swagger/index.html即可查看swagger文档。

7、登陆与验证（jwt中间件）

验证分为前端身份验证与后端身份验证。前端用户身份验证可以使用cookie进行，但是很多移动端如小程序等没有cookie，且存在其他的缺点。目前更加流行后端验证，其中使用较多的有jwt。

1）生成token；

2）验证token；

3）token中间件：

![image-20240605171706683](.\image\image-20240605171706683.png)
