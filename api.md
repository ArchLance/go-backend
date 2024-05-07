# 用户规则管理

## 获取所有用户规则

>基本信息

**Path：/api/usr**

**Method：GET**

> 返回数据

| 名称   | 类型        | 备注                                                         |
| ------ | ----------- | ------------------------------------------------------------ |
| code   | int         |                                                              |
| status | string      | 返回状态                                                     |
| data   | interface{} | 用户规则数据，<br />{"id": int, "path": string, "usr": string} |



## 根据id获取用户规则

>基本信息

**Path：/api/usr/{id}**

**Method：GET**

> 返回数据

| 名称   | 类型        | 备注                                                         |
| ------ | ----------- | ------------------------------------------------------------ |
| code   | int         |                                                              |
| status | string      | 返回状态                                                     |
| data   | interface{} | 用户信息如下<br />{"id": int, "path": string, "usr": string} |

## 创建用户规则

>基本信息

**Path：/api/usr/create**

**Method：POST**

**接口描述：**

> 请求参数

**Body**

| 名称 | 类型   | 备注         |
| ---- | ------ | ------------ |
| path | string | 规则设置路径 |
| usr  | string | 用户         |

> 返回数据

| 名称   | 类型        | 备注 |
| ------ | ----------- | ---- |
| code   | int         |      |
| status | string      |      |
| data   | interface{} | nil  |

## 更新用户规则

>基本信息

**Path：/api/usr/{id}**

**Method：PATCH**

**接口描述：**

> 请求参数

**Body**

| 名称 | 类型   | 备注         |
| ---- | ------ | ------------ |
| path | string | 规则设置路径 |
| usr  | string | 用户         |

> 返回数据

| 名称   | 类型        | 备注 |
| ------ | ----------- | ---- |
| code   | int         |      |
| status | string      |      |
| data   | interface{} | nil  |

## 删除用户规则

>基本信息

**Path：/api/usr/{id}**

**Method：DELETE**

> 返回数据

| 名称   | 类型        | 备注     |
| ------ | ----------- | -------- |
| code   | int         |          |
| status | string      | 返回状态 |
| data   | interface{} | nil      |

# 进程规则管理

## 获取所有进程规则

>基本信息

**Path：/api/process**

**Method：GET**

> 返回数据

| 名称   | 类型        | 备注                                                         |
| ------ | ----------- | ------------------------------------------------------------ |
| code   | int         |                                                              |
| status | string      | 返回状态                                                     |
| data   | interface{} | 用户规则数据，<br />{"id": int, "path": string, "process": string} |



## 根据id获取进程规则

>基本信息

**Path：/api/process/{id}**

**Method：GET**

> 返回数据

| 名称   | 类型        | 备注                                                         |
| ------ | ----------- | ------------------------------------------------------------ |
| code   | int         |                                                              |
| status | string      | 返回状态                                                     |
| data   | interface{} | 用户信息如下<br />{"id": int, "path": string, "process": string} |

## 创建进程规则

>基本信息

**Path：/api/process/create**

**Method：POST**

**接口描述：**

> 请求参数

**Body**

| 名称    | 类型   | 备注         |
| ------- | ------ | ------------ |
| path    | string | 规则设置路径 |
| process | string | 进程         |

> 返回数据

| 名称   | 类型        | 备注 |
| ------ | ----------- | ---- |
| code   | int         |      |
| status | string      |      |
| data   | interface{} | nil  |

## 更新用户规则

>基本信息

**Path：/api/process/{id}**

**Method：PATCH**

**接口描述：**

> 请求参数

**Body**

| 名称    | 类型   | 备注         |
| ------- | ------ | ------------ |
| path    | string | 规则设置路径 |
| process | string | 用户         |

> 返回数据

| 名称   | 类型        | 备注 |
| ------ | ----------- | ---- |
| code   | int         |      |
| status | string      |      |
| data   | interface{} | nil  |

## 删除用户规则

>基本信息

**Path：/api/process/{id}**

**Method：DELETE**

> 返回数据

| 名称   | 类型        | 备注     |
| ------ | ----------- | -------- |
| code   | int         |          |
| status | string      | 返回状态 |
| data   | interface{} | nil      |

