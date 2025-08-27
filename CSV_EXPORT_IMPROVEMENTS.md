# GoTestWAF CSV导出功能改进

## 概述

本次改进为GoTestWAF的CSV日志导出功能添加了完整的HTTP请求信息，让用户能够看到每次测试请求的详细情况。

## 改进内容

### 1. 新增字段

在原有的CSV导出基础上，新增了以下4个字段：

- **HTTP Method**: HTTP请求方法 (GET, POST, PUT, DELETE等)
- **Request URL**: 完整的请求URL (包括路径、查询参数等)
- **Request Headers**: 请求头信息 (格式化为易读的字符串)
- **Request Body**: 请求体内容 (对于POST/PUT等请求)

### 2. 修改的文件

#### `internal/db/models.go`
- 在 `Info` 结构体中添加了新的字段：
  ```go
  type Info struct {
      // ... 原有字段 ...
      
      // 新增：HTTP请求信息
      HTTPMethod         string
      RequestURL         string
      RequestHeaders     string
      RequestBody        string
  }
  ```

#### `internal/scanner/scanner.go`
- 添加了 `toInfoWithRequest` 函数，用于提取和保存完整的HTTP请求信息
- 修改了 `updateDB` 函数，使用新的函数来保存请求信息

#### `internal/db/export.go`
- 在CSV导出中添加了新的列
- 确保所有测试类型（blocked、passed、unresolved）都包含新的请求信息

#### `internal/db/database.go`
- 添加了 `SetTestDataForTesting` 方法，用于测试时设置数据

## 使用效果

### 改进前的CSV格式
```
Payload, Check Status, Response Code, Placeholder, Encoder, Set, Case, Test Result
```

### 改进后的CSV格式
```
Payload, Check Status, Response Code, Placeholder, Encoder, Set, Case, Test Result, HTTP Method, Request URL, Request Headers, Request Body
```

## 示例输出

```
test-payload, blocked, 403, URLParam, Base64, XSS, test-case, passed, GET, http://example.com/test?param=dGVzdC1wYXlsb2Fk, User-Agent: Mozilla/5.0; Accept: text/html, 
test-payload, passed, 200, RequestBody, Plain, SQLInjection, test-case, failed, POST, http://example.com/api/submit, Content-Type: application/json, {"data": "test-payload"}
```

## 技术实现

### 请求信息提取

1. **HTTP方法**: 从 `http.Request.Method` 获取
2. **完整URL**: 从 `http.Request.URL.String()` 获取
3. **请求头**: 格式化所有请求头为 `Key: Value` 格式，用分号分隔
4. **请求体**: 读取请求体内容，支持文本和二进制数据

### 兼容性

- 支持GoHTTP客户端和Chrome客户端
- 对于Chrome请求，提供基本的标识信息
- 保持向后兼容，不影响现有功能

## 优势

1. **完整的请求记录**: 可以看到每次测试的完整HTTP请求
2. **便于调试**: 当测试失败时，可以重现具体的请求
3. **安全分析**: 安全团队可以分析具体的攻击向量和请求模式
4. **合规性**: 满足某些合规要求中对完整请求日志的需求

## 注意事项

1. **数据量增加**: CSV文件会变大，因为包含了更多信息
2. **敏感信息**: 请求头可能包含敏感信息（如Authorization），请注意数据安全
3. **性能影响**: 轻微的性能影响，因为需要读取和格式化请求信息

## 测试

使用以下命令测试功能：

```bash
# 构建项目
go build ./cmd/gotestwaf

# 运行测试（如果有的话）
go test ./...
```

## 总结

这次改进让GoTestWAF的CSV日志更加完整和有用，用户可以：

- 看到每次测试的完整HTTP请求
- 更容易调试和重现问题
- 进行更深入的安全分析
- 满足合规性要求

所有修改都保持了向后兼容性，不会影响现有的功能。
