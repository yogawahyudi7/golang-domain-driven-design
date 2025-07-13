# Panic Recovery Implementation

## Overview

Our Gin application now includes comprehensive panic recovery middleware that can detect, log, and gracefully handle panics that occur during HTTP request processing.

## Features

### 1. Custom Panic Recovery Middleware

Located in `internal/interfaces/middleware/panic_recovery.go`, our middleware provides:

- **Automatic Panic Detection**: Catches all panics that occur during request processing
- **Comprehensive Logging**: Logs panic details to multiple log files with different levels
- **Stack Trace Capture**: Full stack traces for debugging
- **Graceful Response**: Returns JSON error responses instead of server crashes
- **Security Monitoring**: Logs panics as security events for monitoring
- **Business Intelligence**: Tracks panics as business events for system health

### 2. Multi-Level Logging

When a panic occurs, the middleware logs to:

1. **Main Application Log**: Detailed panic information with context
2. **Error Log File**: Specific error file for panic tracking
3. **Security Log**: Security event for potential security issues
4. **Business Event Log**: System health monitoring

### 3. Response Handling

Instead of crashing the server, panics are handled gracefully:

```json
{
  "error": "Internal Server Error",
  "message": "An unexpected error occurred. Please try again later.",
  "request_id": "uuid-request-id"
}
```

## Implementation Details

### Middleware Registration

In `internal/interfaces/routes/routes.go`:

```go
// Add middleware in order
router.Use(middleware.LoggerMiddleware(appLogger))
router.Use(middleware.CORS())
router.Use(middleware.PanicRecoveryMiddleware(appLogger)) // Custom panic recovery
```

### Panic Information Logged

For each panic, the following information is captured:

- **Request Details**: Method, path, client IP, user agent
- **Panic Value**: The actual panic message/value
- **Stack Trace**: Full stack trace for debugging
- **Request ID**: For correlation across logs
- **Timestamp**: When the panic occurred

## Testing Panic Recovery

### Test Endpoints

We've added several test endpoints to verify panic recovery:

1. **`/test/panic`** - Basic string panic
2. **`/test/panic?type=nil`** - Nil pointer dereference
3. **`/test/panic?type=slice`** - Slice index out of bounds
4. **`/test/panic?type=map`** - Nil map access
5. **`/test/panic?type=custom&message=Your message`** - Custom panic

### Running Tests

1. **Start the server:**
   ```bash
   go run cmd/api/main.go
   ```

2. **Run the panic recovery demo:**
   ```bash
   go run -tags=example examples/panic_recovery_demo.go
   ```

3. **Manual testing with curl:**
   ```bash
   # Test normal endpoint
   curl http://localhost:8080/health
   
   # Test panic recovery
   curl http://localhost:8080/test/panic
   curl http://localhost:8080/test/panic?type=nil
   curl http://localhost:8080/test/panic?type=custom&message=Testing
   ```

## Log Output Examples

### Main Application Log

```json
{
  "level": "error",
  "message": "Panic recovered",
  "timestamp": "2025-07-13T14:30:00+07:00",
  "request_id": "abc-123-def",
  "method": "GET",
  "path": "/test/panic",
  "client_ip": "127.0.0.1",
  "user_agent": "curl/7.68.0",
  "panic_value": "This is a test panic!",
  "stack_trace": "goroutine 123 [running]:\\n..."
}
```

### Security Event Log

```json
{
  "level": "warning",
  "message": "Security event detected",
  "type": "security",
  "event": "panic_recovered",
  "user_id": "user_127.0.0.1",
  "client_ip": "127.0.0.1",
  "details": "panic_value: This is a test panic!",
  "timestamp": "2025-07-13T14:30:00+07:00"
}
```

### Business Event Log

```json
{
  "level": "info",
  "message": "Business event occurred",
  "type": "business_event",
  "event": "system_panic",
  "entity_id": "system",
  "user_id": "user_127.0.0.1",
  "request_id": "abc-123-def",
  "panic_message": "This is a test panic!",
  "endpoint": "/test/panic",
  "method": "GET",
  "recovery": "successful",
  "timestamp": "2025-07-13T14:30:00+07:00"
}
```

## Benefits

### 1. **Server Stability**
- Prevents server crashes from unhandled panics
- Maintains service availability during errors
- Graceful degradation instead of complete failure

### 2. **Debugging Support**
- Full stack traces for root cause analysis
- Request context for reproduction
- Correlation IDs for distributed tracing

### 3. **Monitoring & Alerting**
- Security events for potential attack detection
- Business events for health monitoring
- Structured logs for automated analysis

### 4. **User Experience**
- Consistent error responses
- No exposed internal error details
- Proper HTTP status codes

## Production Considerations

### 1. **Log Management**
- Monitor panic frequency
- Set up alerts for panic spikes
- Regular log rotation and archival

### 2. **Performance Impact**
- Minimal overhead during normal operation
- Stack trace generation only during panics
- Async logging where possible

### 3. **Security**
- Panic details not exposed to users
- Security events for monitoring
- Rate limiting on test endpoints (remove in production)

### 4. **Monitoring Integration**
- Integrate with APM tools
- Set up panic frequency alerts
- Dashboard for panic trends

## Comparison with Default Gin Recovery

| Feature | Default `gin.Recovery()` | Our Custom Middleware |
|---------|--------------------------|----------------------|
| Basic panic recovery | ✅ | ✅ |
| Detailed logging | ❌ | ✅ |
| Stack traces | ❌ | ✅ |
| Security events | ❌ | ✅ |
| Business events | ❌ | ✅ |
| Request correlation | ❌ | ✅ |
| Multi-file logging | ❌ | ✅ |
| Custom response format | ❌ | ✅ |

## Best Practices

1. **Remove test endpoints in production**
2. **Monitor panic frequency and patterns**
3. **Set up automated alerts for panic spikes**
4. **Regular review of panic logs for system improvement**
5. **Integrate with external monitoring systems**
6. **Consider adding rate limiting for panic-prone endpoints**
