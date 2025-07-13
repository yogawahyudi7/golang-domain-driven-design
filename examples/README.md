# Examples

This directory contains demonstration scripts for the logging functionality. Each example shows different aspects of the logging system implementation.

## Available Examples

### 1. Pretty Logging Demo (`pretty_logging_demo.go`)
Demonstrates the pretty JSON formatting feature using the custom PrettyJSONFormatter.

**Features shown:**
- Pretty formatted JSON logs
- Both info and error log types
- File output with indentation

**Run with:**
```bash
go run -tags=example examples/pretty_logging_demo.go
```

### 2. Log Naming Demo (`log_naming_demo.go`)
Shows the best practices for log file naming with timestamps and categorization.

**Features shown:**
- Date-based log file naming
- Multiple log file types (app, error, access)
- Automatic directory creation

**Run with:**
```bash
go run -tags=example examples/log_naming_demo.go
```

### 3. Date Rotation Demo (`date_rotation_demo.go`)
Demonstrates log file rotation based on dates and shows how logs are organized over time.

**Features shown:**
- Daily log rotation
- Historical log preservation
- Multiple days simulation

**Run with:**
```bash
go run -tags=example examples/date_rotation_demo.go
```

### 4. Log Rotation Demo (`log_rotation_demo.go`)
Advanced demonstration of Lumberjack log rotation with size-based rotation.

**Features shown:**
- Size-based log rotation
- Backup file management
- Compression options

**Run with:**
```bash
go run -tags=example examples/log_rotation_demo.go
```

### 5. Tomorrow Simulation (`tomorrow_simulation.go`)
Simulates logging across multiple days to demonstrate date-based file rotation.

**Features shown:**
- Cross-day logging simulation
- Date transition handling
- Multiple log file creation

**Run with:**
```bash
go run -tags=example examples/tomorrow_simulation.go
```

## Build Tags Explanation

All example files use build tags to prevent conflicts with the main application. The build tags used are:
- `//go:build example`
- `// +build example`

This approach allows:
1. **Normal builds** exclude examples: `go build` or `go run cmd/api/main.go`
2. **Example execution** includes specific examples: `go run -tags=example examples/filename.go`
3. **No main function conflicts** between examples and the main application

## Running All Examples

To run all examples in sequence:

```bash
# Pretty logging demonstration
go run -tags=example examples/pretty_logging_demo.go

# Log naming best practices
go run -tags=example examples/log_naming_demo.go

# Date-based rotation
go run -tags=example examples/date_rotation_demo.go

# Size-based rotation
go run -tags=example examples/log_rotation_demo.go

# Multi-day simulation
go run -tags=example examples/tomorrow_simulation.go
```

## Output

Each example creates log files in the `logs/` directory (which is excluded from git). The structure typically looks like:

```
logs/
├── app-2024-01-15.log
├── error-2024-01-15.log
├── access-2024-01-15.log
├── example-pretty-2024-01-15.log
└── rotation-demo-2024-01-15.log
```

## Best Practices Demonstrated

1. **Build Tags**: Isolating example code from production builds
2. **Log File Organization**: Date-based and type-based file naming
3. **Pretty Formatting**: Human-readable JSON logs for development
4. **Rotation Strategies**: Both size-based and date-based rotation
5. **Multiple Log Types**: Separating application, error, and access logs
