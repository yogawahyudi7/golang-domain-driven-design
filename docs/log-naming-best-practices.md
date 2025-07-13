# Log File Naming Best Practices

## 📋 **Standar Industri Penamaan File Log**

### 1. **Current Active Logs**
```
logs/
├── app.log          # Main application log (current)
├── error.log        # Error-only log (current)
├── access.log       # HTTP access log (current)
└── audit.log        # Security/audit log (current)
```

### 2. **Rotated Log Files (Lumberjack Format)**
```
logs/
├── app.log                 # Current active log
├── app.log.2025-07-13      # Yesterday's log
├── app.log.2025-07-12      # Day before yesterday
├── app.log.2025-07-11.gz   # Compressed older log
├── error.log               # Current error log
├── error.log.2025-07-13    # Yesterday's errors
└── access.log.2025-07-13   # Yesterday's access log
```

### 3. **Alternative Naming Conventions**

#### A. **Date-based Current Files**
```
logs/
├── app-2025-07-13.log      # Daily rotation
├── error-2025-07-13.log    # Daily error logs
└── access-2025-07-13.log   # Daily access logs
```

#### B. **Hourly Rotation**
```
logs/
├── app-2025-07-13-14.log   # Hour 14 (2 PM)
├── app-2025-07-13-15.log   # Hour 15 (3 PM)
└── app-2025-07-13-16.log   # Hour 16 (4 PM)
```

#### C. **Service-based Separation**
```
logs/
├── api/
│   ├── app.log
│   ├── error.log
│   └── access.log
├── worker/
│   ├── worker.log
│   └── error.log
└── scheduler/
    ├── scheduler.log
    └── error.log
```

## 🎯 **Recommended Best Practices**

### **1. File Naming Convention**
- **Use descriptive names**: `app.log`, `error.log`, `access.log`
- **Avoid timestamps in active files**: Let rotation handle dating
- **Use lowercase**: Consistent with Unix conventions
- **Use hyphens for separation**: `user-service.log`, `payment-api.log`

### **2. Rotation Strategy**
```go
// Recommended Lumberjack Configuration
&lumberjack.Logger{
    Filename:   "logs/app.log",
    MaxSize:    100,    // 100MB per file
    MaxBackups: 7,      // Keep 7 backup files
    MaxAge:     30,     // Keep logs for 30 days
    Compress:   true,   // Compress rotated files
    LocalTime:  true,   // Use local timezone
}
```

### **3. Log Types Separation**

#### **Application Logs** (`app.log`)
- General application events
- Business logic operations
- Startup/shutdown events
- Configuration changes

#### **Error Logs** (`error.log`)
- Application errors
- Exception stack traces
- Failed operations
- Critical system events

#### **Access Logs** (`access.log`)
- HTTP requests/responses
- API endpoint usage
- Authentication attempts
- Request timing and status

#### **Audit Logs** (`audit.log`)
- Security events
- User actions
- Data modifications
- Permission changes

## 📁 **Directory Structure Best Practices**

### **Option 1: Flat Structure (Small Applications)**
```
logs/
├── app.log
├── app.log.2025-07-13
├── error.log
├── error.log.2025-07-13
├── access.log
└── access.log.2025-07-13
```

### **Option 2: Categorized Structure (Large Applications)**
```
logs/
├── application/
│   ├── app.log
│   └── app.log.2025-07-13
├── errors/
│   ├── error.log
│   └── error.log.2025-07-13
├── access/
│   ├── access.log
│   └── access.log.2025-07-13
└── audit/
    ├── audit.log
    └── audit.log.2025-07-13
```

### **Option 3: Environment-based Structure**
```
logs/
├── development/
│   ├── app.log
│   └── error.log
├── staging/
│   ├── app.log
│   └── error.log
└── production/
    ├── app.log
    ├── error.log
    └── access.log
```

## ⚙️ **Implementation in Go**

### **Multiple Log Files Configuration**
```go
type Config struct {
    LogDir        string `json:"log_dir"`         // "logs"
    AppLogFile    string `json:"app_log_file"`    // "app.log"
    ErrorLogFile  string `json:"error_log_file"`  // "error.log"
    AccessLogFile string `json:"access_log_file"` // "access.log"
    DateFormat    string `json:"date_format"`     // "2006-01-02"
}
```

### **Log File Rotation Examples**
```go
// Daily rotation with date suffix
filename := fmt.Sprintf("app-%s.log", time.Now().Format("2006-01-02"))

// Hourly rotation
filename := fmt.Sprintf("app-%s.log", time.Now().Format("2006-01-02-15"))

// Service-based naming
filename := fmt.Sprintf("%s-service.log", serviceName)
```

## 🔍 **Log Analysis Considerations**

### **1. Searchability**
- Use consistent naming for easy grep/search
- Include service/component names
- Maintain chronological order

### **2. Monitoring Integration**
- Standard names for log aggregators (ELK, Splunk)
- Predictable file paths for monitoring tools
- Clear separation by log level/type

### **3. Troubleshooting**
- Quick identification by filename
- Easy correlation between related logs
- Clear timestamp in rotated files

## 🚀 **Production Recommendations**

### **1. High-Volume Applications**
```
logs/
├── api/
│   ├── access-2025-07-13.log     # Daily access logs
│   ├── error-2025-07-13.log      # Daily error logs
│   └── app-2025-07-13.log        # Daily app logs
└── archived/
    ├── api-access-2025-07-12.gz
    ├── api-error-2025-07-12.gz
    └── api-app-2025-07-12.gz
```

### **2. Microservices**
```
logs/
├── user-service/
│   ├── app.log
│   └── error.log
├── payment-service/
│   ├── app.log
│   └── error.log
└── notification-service/
    ├── app.log
    └── error.log
```

### **3. Container Environments**
```
# In containers, use stdout/stderr and let orchestrator handle files
logs/
├── containers/
│   ├── user-service-abc123.log
│   ├── payment-service-def456.log
│   └── api-gateway-ghi789.log
```

## 📊 **Monitoring and Alerting**

### **File Size Monitoring**
- Monitor active log file sizes
- Alert when files grow too large
- Ensure rotation is working properly

### **Retention Policy**
- Implement automatic cleanup
- Archive important logs
- Consider compliance requirements

### **Performance Impact**
- Monitor disk I/O from logging
- Use appropriate buffer sizes
- Consider async logging for high-volume

## 🛡️ **Security Considerations**

### **File Permissions**
```bash
# Recommended permissions
chmod 644 *.log        # Read for owner/group, read-only for others
chmod 755 logs/        # Directory permissions
```

### **Sensitive Data**
- Never log passwords or API keys
- Mask PII in access logs
- Use separate audit logs for security events

### **Log Integrity**
- Consider log signing for audit trails
- Implement tamper detection
- Regular backup of critical logs
