# Logging Implementation with Logrus

This document describes the professional logging implementation using Logrus and Lumberjack for the Golang Domain-Driven Design application.

## Features

### 🚀 Production-Ready Features
- **Structured JSON Logging**: All logs are in JSON format for easy parsing and analysis
- **File Rotation**: Automatic log file rotation with configurable size, backup count, and retention
- **Multiple Output Targets**: Simultaneous output to console and file
- **Request ID Tracking**: Unique request IDs for tracing HTTP requests
- **Performance Monitoring**: Latency tracking and performance metrics
- **Specialized Log Types**: HTTP requests, database operations, business events, security events

### 📊 Log Types

1. **HTTP Request Logs**
   - Request ID, method, path, status code
   - Client IP, user agent, latency
   - Success/error status

2. **Database Operation Logs**
   - Operation type, table, duration
   - Query success/failure status

3. **Business Event Logs**
   - Domain-specific business logic events
   - User actions and state changes

4. **Security Event Logs**
   - Authentication attempts
   - Authorization failures
   - Security-related activities

5. **Performance Logs**
   - Operation timing and metrics
   - Resource usage monitoring

## Configuration

### Environment Variables

Add these to your `.env` file:

```env
# Logger Configuration
LOG_LEVEL=info
LOG_FORMAT=json
LOG_OUTPUT=both
LOG_PRETTY_PRINT=true
LOG_FILE_PATH=logs/app.log
LOG_MAX_SIZE=100
LOG_MAX_BACKUPS=3
LOG_MAX_AGE=28
LOG_COMPRESS=true
```

### Configuration Options

| Variable | Description | Default | Options |
|----------|-------------|---------|---------|
| `LOG_LEVEL` | Minimum log level | `info` | `trace`, `debug`, `info`, `warn`, `error`, `fatal`, `panic` |
| `LOG_FORMAT` | Log output format | `json` | `json`, `text` |
| `LOG_OUTPUT` | Output destination | `both` | `console`, `file`, `both` |
| `LOG_PRETTY_PRINT` | Enable pretty JSON formatting | `false` | `true`, `false` |
| `LOG_FILE_PATH` | Log file location | `logs/app.log` | Any valid file path |
| `LOG_MAX_SIZE` | Max file size (MB) | `100` | Integer value |
| `LOG_MAX_BACKUPS` | Number of backup files | `3` | Integer value |
| `LOG_MAX_AGE` | Max age of logs (days) | `28` | Integer value |
| `LOG_COMPRESS` | Compress old logs | `true` | `true`, `false` |

## Usage Examples

### Basic Logging

```go
// Initialize logger
logger := logger.New(config.Logger)

// Basic logging
logger.Info("Application started")
logger.Error("Something went wrong", "error", err)
logger.Debug("Debug information", "data", debugData)
```

### HTTP Request Logging

```go
// Automatic via middleware
func LoggerMiddleware(logger *logger.Logger) gin.HandlerFunc {
    return func(c *gin.Context) {
        // Generates unique request ID and logs automatically
        start := time.Now()
        requestID := uuid.New().String()
        c.Set("request_id", requestID)
        
        c.Next()
        
        logger.LogHTTPRequest(c.Request.Method, c.Request.URL.Path, 
            c.Writer.Status(), time.Since(start), c.ClientIP(), 
            c.Request.UserAgent(), requestID)
    }
}
```

### Database Operation Logging

```go
// Log database operations
logger.LogDBOperation("SELECT", "users", time.Since(start), nil)

// Log with error
if err != nil {
    logger.LogDBOperation("INSERT", "users", time.Since(start), err)
}
```

### Business Event Logging

```go
// Log business events
logger.LogBusinessEvent("user_registered", map[string]interface{}{
    "user_id": user.ID,
    "email": user.Email,
    "registration_method": "email",
})

logger.LogBusinessEvent("payment_processed", map[string]interface{}{
    "user_id": userID,
    "amount": amount,
    "currency": "USD",
    "payment_method": "credit_card",
})
```

### Security Event Logging

```go
// Log security events
logger.LogSecurityEvent("login_attempt", map[string]interface{}{
    "email": email,
    "ip_address": clientIP,
    "success": false,
    "reason": "invalid_password",
})

logger.LogSecurityEvent("jwt_token_expired", map[string]interface{}{
    "user_id": userID,
    "token_exp": expTime,
    "current_time": time.Now(),
})
```

### Performance Monitoring

```go
// Log performance metrics
start := time.Now()
result := performExpensiveOperation()
logger.LogPerformance("expensive_operation", time.Since(start), map[string]interface{}{
    "records_processed": len(result),
    "cache_hit": cacheHit,
})
```

### Pretty JSON Format Configuration

The logger supports pretty-printed JSON output for better readability during development:

```go
// Enable pretty print in development
logger := logger.New(&logger.Config{
    Format:      "json",
    PrettyPrint: true,  // Enable pretty formatting
    Output:      "both",
})

// Disable for production (better performance)
logger := logger.New(&logger.Config{
    Format:      "json", 
    PrettyPrint: false, // Compact single-line JSON
    Output:      "file",
})
```

#### Benefits of Pretty Print
- **Development**: Easier to read and debug log entries
- **Console Output**: Better formatting for terminal viewing  
- **Log Analysis**: Improved readability when examining individual entries

#### Performance Considerations
- **Development**: Use `LOG_PRETTY_PRINT=true` for better readability
- **Production**: Use `LOG_PRETTY_PRINT=false` for optimal performance and storage efficiency
- **File Size**: Pretty printing increases log file size due to indentation and newlines

## Log Format Examples

### Pretty JSON Format (Development)
When `LOG_PRETTY_PRINT=true` is set, logs are formatted with indentation for better readability:

#### HTTP Request Log (Pretty)
```json
{
  "caller": "golang-domain-driven-design/pkg/logger.(*Logger).LogHTTPRequest",
  "client_ip": "127.0.0.1",
  "file": "C:/path/to/logger.go:160",
  "latency_ms": 15,
  "level": "info",
  "message": "HTTP request completed successfully",
  "method": "POST",
  "path": "/api/v1/users",
  "request_id": "dc3c1ef9-c12c-4f2a-9e5e-26a38312d7e4",
  "status_code": 201,
  "timestamp": "2025-07-13T13:35:43+07:00",
  "type": "http_request",
  "user_agent": "curl/8.8.0"
}
```

#### Application Startup Log (Pretty)
```json
{
  "app_name": "golang-cloude",
  "caller": "main.main",
  "env": "development",
  "file": "C:/path/to/main.go:43",
  "level": "info",
  "message": "Starting application",
  "timestamp": "2025-07-13T11:10:36+07:00",
  "version": "1.0.0"
}
```

### Compact JSON Format (Production)
When `LOG_PRETTY_PRINT=false` is set, logs are compact single-line JSON for better performance:

#### HTTP Request Log (Compact)
```json
{"caller":"golang-domain-driven-design/pkg/logger.(*Logger).LogHTTPRequest","client_ip":"127.0.0.1","file":"C:/path/to/logger.go:160","latency_ms":15,"level":"info","message":"HTTP request completed successfully","method":"POST","path":"/api/v1/users","request_id":"dc3c1ef9-c12c-4f2a-9e5e-26a38312d7e4","status_code":201,"timestamp":"2025-07-13T13:35:43+07:00","type":"http_request","user_agent":"curl/8.8.0"}
```

#### Application Startup Log (Compact)
```json
{"app_name":"golang-cloude","caller":"main.main","env":"development","file":"C:/path/to/main.go:43","level":"info","message":"Starting application","timestamp":"2025-07-13T11:10:36+07:00","version":"1.0.0"}
```

### Database Operation Log
```json
{
  "caller": "golang-domain-driven-design/pkg/logger.(*Logger).LogDBOperation",
  "duration_ms": 25,
  "level": "info",
  "message": "Database operation completed successfully",
  "operation": "SELECT",
  "table": "users",
  "timestamp": "2025-07-13T13:40:22+07:00",
  "type": "database_operation"
}
```

## File Structure

```
logs/
├── app.log              # Current log file
├── app.log.2025-07-12   # Rotated log file
└── app.log.2025-07-11.gz # Compressed old log
```

## Integration Points

### 1. Main Application (`cmd/api/main.go`)
- Logger initialization with configuration
- Structured startup and shutdown logging

### 2. HTTP Middleware (`internal/interfaces/middleware/logger.go`)
- Request ID generation and tracking
- Automatic request/response logging
- Error and success status logging

### 3. Configuration (`internal/infrastructure/config/config.go`)
- Logger configuration structure
- Environment variable parsing
- Default value handling

### 4. Routes Setup (`internal/interfaces/routes/routes.go`)
- Logger middleware integration
- Route-specific logging configuration

## Best Practices

1. **Use Structured Logging**: Always include relevant context in log entries
2. **Request ID Tracking**: Use request IDs to trace requests across services
3. **Log Levels**: Use appropriate log levels (INFO for normal operations, ERROR for failures)
4. **Performance Logging**: Monitor critical operations with timing information
5. **Security Logging**: Log all authentication and authorization events
6. **Error Context**: Include error details and stack traces where appropriate

## Monitoring and Alerting

### Log Analysis
- Use tools like ELK Stack, Grafana, or similar for log analysis
- Set up alerts for error patterns and performance thresholds
- Monitor log file sizes and rotation

### Metrics to Monitor
- Error rates and patterns
- Response times and latency
- Authentication failures
- Database operation performance
- Application startup/shutdown events

## Production Recommendations

1. **Log Level**: Set to `info` or `warn` in production
2. **File Rotation**: Configure appropriate file sizes and retention
3. **Monitoring**: Set up log monitoring and alerting
4. **Storage**: Ensure adequate disk space for log files
5. **Security**: Restrict access to log files
6. **Backup**: Include logs in backup strategies

## Dependencies

- **Logrus**: `github.com/sirupsen/logrus` - Structured logging
- **Lumberjack**: `gopkg.in/natefinch/lumberjack.v2` - Log rotation
- **UUID**: `github.com/google/uuid` - Request ID generation
