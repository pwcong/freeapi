# 附件接口

- 上传附件

## 上传附件

路径：`/api/attachment/upload`

方法: `POST`

请求参数:

- file: 文件

请求示例：

```shell
curl -X POST $API_BASE/api/attachment/upload \
  -H 'content-type: multipart/form-data; boundary=xxx' \
  -F 'file=@xxx'
```

返回值：

```json
{
  "success": true,
  "message": "upload successfully",
  "code": 20000,
  "payload": {
    "url": "xxx"
  }
}
```

## 上传多附件

路径：`/api/attachment/uploads`

方法: `POST`

请求参数:

- files: 文件

请求示例：

```shell
curl -X POST $API_BASE/api/attachment/upload \
  -H 'content-type: multipart/form-data; boundary=xxx' \
  -F 'files=@xxx' \
  -F 'files=@xxx'
```

返回值：

```json
{
  "success": true,
  "message": "upload successfully",
  "code": 20000,
  "payload": {
    "urls": ["xxx", "xxx"]
  }
}
```
