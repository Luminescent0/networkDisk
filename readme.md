<pre>
├─api
│      file.go
│      middle.go
│      router.go
│      user.go
│      
├─cmd
│      main.go
│      
├─dao
│      dao.go
│      file.go
│      user.go
│      
├─file
├─model
│      file.go
│      user.go
│      
├─service
│      file.go
│      user.go
│      
└─tool
        resp.go
</pre>
# 实现功能
1. 登陆注册
2. 上传文件（设置权限
3. 文件修改（路径 重命名 权限
4. 下载文件
5. 下载限速
6. 文件分享

# 接口说明
# login
## `POST` `/login`
 `application/x-www-form-urlencoded`

|请求参数 | 类型 | 备注 |
|--------|-----|------|
username|必选|用户名|
password|必选|密码|

|返回参数|说明|
|------|---|
|date|返回消息|
|token|用户token|

|date|说明|
|---|---|
|"info":服务器错误|服务器错误|
|“密码错误”|`password`与`username`不匹配|
|”用户不存在“|`username`不存在|
|用户token|登录成功|

# Register
## `POST` `/register`
`application/x-www-form-urlencoded`

|请求参数 | 类型 | 备注 |
|--------|-----|------|
username|必选|用户名|
password|必选|密码|

|返回参数|说明|
|---|---|
|data|返回消息|

|data|说明|
|---|---|
|"用户名格式有误(min=4,max-10)"|`username`格式有误|
|"密码格式有误(min=6,max=16)"|`password`格式有误|
|”服务器错误“|服务器错误|
|”用户名已经存在“|`username`已存在|
|“注册成功”|注册成功|
|“路径有误”|mkdirall出错|

# file `/:username`
## upload
### `POST` `/upload`
`form-data`
`application/x-www-form-urlencoded`

|请求参数 | 类型 | 备注 |
|--------|-----|------|
|upload|必选|需上传的文件|
|permit|可选|权限管理：默认为不公开|

|返回参数|说明|
|---|---|
|data|返回消息|

|data|说明|
|---|---|
|“保存失败”|无法存储|
|“插入数据失败”|数据库更新失败|
|“成功”|上传成功|

## delFile
### `DELETE` `/delete`
`application/x-www-form-urlencoded`

|请求参数 | 类型 | 备注 |
|--------|-----|------|
fileName|必选|被删除的文件名|
filePath|必选|文件路径|

|返回参数|说明|
|---|---|
|data|返回消息|

|data|说明|
|---|---|
|“请求失败”|无法删除|
|“数据删除失败”|数据库更新失败|
|“成功”|删除成功|

## changePath
## `PUT` `/changePath`
`application/x-www-form-urlencoded`

|请求参数 | 类型 | 备注 |
|--------|-----|------|
fileName|必选|文件名|
oldPath|必选|当前文件路径|
newPath|必选|修改后的文件路径|

|返回参数|说明|
|---|---|
|data|返回消息|

|data|说明|
|---|---|
|”新路径格式有误“|路径格式出错
|”请求失败“|无法移动文件|
|“数据更新失败”|数据库更新失败|
|“成功”|移动成功|

## changeName
## `PUT` `/changeName`
`application/x-www-form-urlencoded`

|请求参数 | 类型 | 备注 |
|--------|-----|------|
oldName|必选|原文件名|
localPath|必选|文件路径|
newName|必选|新文件名|

|返回参数|说明|
|---|---|
|data|返回消息|

|data|说明|
|---|---|
|”请求失败“|无法移动文件|
|“数据更新失败”|数据库更新失败|
|“成功”|修改成功|

## downloadByLink
`GET`

|返回参数|说明|
|---|---|
|data|返回消息|

|data|说明|
|---|---|
|”请求失败“|无法下载文件|

## share
### `POST` `/share`
`application/x-www-form-urlencoded`

|请求参数 | 类型 | 备注 |
|--------|-----|------|
fileName|必选|文件名|
path|必选|文件路径|

|返回参数|说明|
|---|---|
|data|返回消息|

|data|说明|
|---|---|
|“查询文件失败”|无法从数据库中查询到该文件的信息|
|“文件权限不允许公开下载”|文件权限为不公开|
|link|文件下载链接|

