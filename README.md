# go-api-protocols
## Table of Contents

 * [專案描述](#專案描述)
 * [執行專案](#執行專案)
 * [修改專案](#修改專案)

## 專案描述

1. REST 實作
2. GraphQL 實作
3. Json-RPC 實作
4. g-RPC 實作

## 執行專案

#### 執行應用程式

```bash
#到專案目錄下
$ cd path_to_dir/go-api-protocols

# 下載第三方套件
$ go mod download

# 生成swagger文檔
$ swag init 

# 編譯專案(輸出到當前目錄下,檔案名為main)
$ go build -o main . 

# 執行應用程式
$ ./main 
```
#### REST API文檔(swagger)
```bash
#網址打入(default host=>localhost:8080)
http://{host}/swagger/index.html
```
#### GraphQL API文檔
```bash
#網址打入(default host=>localhost:8081)
http://{host}/graphql
```

## 修改專案
### 修改rest
```bash
#refresh swagger doc
swag init
```
### 修改graphQL
```bash
#refresh schema
go run github.com/99designs/gqlgen generate 
```