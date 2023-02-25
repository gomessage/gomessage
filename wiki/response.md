# Response约定

响应信息统一包含以下标准字段：
> - `code`：操作状态（0代表失败，1代表成功）
> - `msg`：描述信息（可以为空）
> - `result`：如果请求成功，此处会返回请求成功的结果集（可以为空）
> - `error`：如果请求失败，此处会返回请求失败的错误内容（可以为空）

Response成功示例：

```json
{
  "code": 1,
  "msg": "创建成功",
  "result": {
    "id": 5,
    "is_active": false,
    "name": "test",
    "description": "test"
  },
  "error": null
}
```

Response失败示例：

```json
{
  "code": 0,
  "msg": "命名空间已存在，不能重复创建",
  "result": null,
  "error": {
    "Code": 19,
    "ExtendedCode": 2067,
    "SystemErrno": 0
  }
}
```


