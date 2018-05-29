# 地域接口

* 获取地域详情
* 获取地域列表

## 获取地域详情

路径：`/api/area/:id`

方法: `POST`

路径参数：

* id：地域 ID

请求示例：

```shell
curl -X GET $API_BASE/api/area/1
```

返回值：

```json
{
  "success": true,
  "code": 20000,
  "message": "get area successfully",
  "payload": {
    "id": 1,
    "created_at": "0001-01-01T00:00:00Z",
    "updated_at": "0001-01-01T00:00:00Z",
    "deleted_at": null,
    "name": "北京",
    "parent_id": 0,
    "depth": 1
  }
}
```

## 获取地域列表

路径：`/api/areas?filter={filter}&query={query}`

方法: `POST`

路径参数：

* filter：地域查询条件，可选值有：

  * parent_id：父级地域 ID
  * depth：层级
  * name：地域名称（模糊查询）

* query：地域查询目标

请求示例：

```shell
curl -X GET $API_BASE/api/areas?filter=name&query=北京
```

返回值：

```json
{
  "success": true,
  "code": 20000,
  "message": "get areas successfully",
  "payload": {
    "page_no": 1,
    "page_size": -1,
    "current_size": 1,
    "total_size": 1,
    "data": [
      {
        "id": 1,
        "created_at": "0001-01-01T00:00:00Z",
        "updated_at": "0001-01-01T00:00:00Z",
        "deleted_at": null,
        "name": "北京",
        "parent_id": 0,
        "depth": 1
      }
    ]
  }
}
```
