# 数学接口

* 进制转换

## 进制转换

路径：`/api/math/convert/hexadecimal`

方法: `POST`

请求参数：

* method：方法名，可选值有（b：二进制、o：八进制、d：十进制、x：十六进制）：

  * b2o
  * o2b
  * b2d
  * d2b
  * b2x
  * x2b
  * o2d
  * d2o
  * o2x
  * x2o
  * d2x
  * x2d

* source：待转换数据

请求示例：

```shell
curl -X GET $API_BASE/api/math/convert/hexadecimal \
  -F source=0100 \
  -F method=b2x
```

返回值：

```json
{
  "success": true,
  "code": 20000,
  "message": "convert successfully",
  "payload": {
    "result": "4"
  }
}
```
