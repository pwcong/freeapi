# 编码接口

* 编码转换

## 编码转换

路径：`/api/code/convert/coding`

方法: `POST`

请求参数：

* method：方法名，可选值有：

  * unicode_encode
  * unicode_decode
  * url_encode
  * url_decode

* source：待转换数据

请求示例：

```shell
curl -X GET $API_BASE/api/code/convert/coding \
  -F source=你好 \
  -F method=unicode_encode
```

返回值：

```json
{
  "success": true,
  "code": 20000,
  "message": "convert successfully",
  "payload": {
    "result": "\\u4f60\\u597d"
  }
}
```
