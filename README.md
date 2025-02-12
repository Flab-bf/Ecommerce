# 说明 
#### 功能  
>
>- 基本功能实现
>- 其他：  
>1.嵌套评论  
>2.使用https  
>3.注意gorm语法防止SQL注入

#### 技术支持
>- go语言  
>- Hertz web框架  
>- 使用gorm进行数据库相关操作  
>- MySQL数据库  
>- 使用OpenSSL签发证书用于测试

#### 项目结构
- cmd/main.go：  
项目入口
- router/router.go：  
配置路由，定义接口的路由规则
- api目录：   
存放用户（如注册，登录等），商品（查找，加入购物车等），评论（查找，发表，回复等）相关接口  
- service目录：  
与dao层进行交互，处理user，merchandise，comment相关具体业务
- dao目录：   
实现连接数据库初始化，进行用户信息，商品信息，用户评论的数据处理，使用gorm与MySQL交互
- model目录：   
定义相关数据结构体，使用gorm与MySQL进行映射
- utils目录：   
成功/错误信息返回，https，jwt等功能函数的实现
- middleWare目录：  
实现jwt的验证功能
- ca_server目录：   
CA，后端服务器  证书，私钥等的存放

#### 注意事项
在修改数据库结构（如添加、修改表字段）后，可能需要使用数据库迁移工具（如 Gorm 的自动迁
移功能）或者手动执行 SQL 脚本来更新数据库表结构，以确保项目与数据库的一致性.