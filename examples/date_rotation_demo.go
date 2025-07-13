//go:build example
// +build example

package main //nolint:errcheck

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func main() {
	fmt.Println("=== SIMULASI LOG ROTATION DENGAN TANGGAL ===")

	// 1. Simulasi hari ini (2025-07-13)
	today := time.Date(2025, 7, 13, 14, 30, 0, 0, time.Local)
	fmt.Printf("🗓️ Hari ini: %s\n", today.Format("2006-01-02 15:04:05"))

	// 2. Simulasi besok (2025-07-14)
	tomorrow := today.AddDate(0, 0, 1)
	fmt.Printf("🗓️ Besok: %s\n\n", tomorrow.Format("2006-01-02 15:04:05"))

	// 3. Buat simulasi file log dengan tanggal
	logDir := "logs/demo"
	if err := os.MkdirAll(logDir, 0755); err != nil {
		fmt.Printf("❌ Error creating demo directory: %v\n", err)
		return
	}

	fmt.Println("📁 Simulasi struktur file log dengan rotasi harian:")

	// Simulasi file log untuk beberapa hari
	dates := []time.Time{
		today.AddDate(0, 0, -3), // 3 hari lalu
		today.AddDate(0, 0, -2), // 2 hari lalu
		today.AddDate(0, 0, -1), // kemarin
		today,                   // hari ini
		tomorrow,                // besok (simulasi)
	}

	logTypes := []string{"app", "error", "access"}

	for _, date := range dates {
		dateStr := date.Format("2006-01-02")

		for _, logType := range logTypes {
			// File aktif (hari ini dan besok)
			if date.Equal(today) || date.Equal(tomorrow) {
				activeFile := filepath.Join(logDir, fmt.Sprintf("%s.log", logType))
				createDemoLogFile(activeFile, logType, date, true)
			}

			// File rotated (dengan tanggal)
			if !date.Equal(today) && !date.Equal(tomorrow) {
				rotatedFile := filepath.Join(logDir, fmt.Sprintf("%s.log.%s", logType, dateStr))
				createDemoLogFile(rotatedFile, logType, date, false)
			}
		}
	}

	// 4. Tampilkan struktur file yang dibuat
	fmt.Println("\n📂 Struktur file log yang dibuat:")
	showDirectoryStructure(logDir)

	// 5. Simulasi log entry untuk "besok"
	fmt.Println("\n🔄 Simulasi: Ketika tanggal berganti ke besok...")
	fmt.Printf("   - File hari ini (%s) akan di-rename menjadi:\n", today.Format("2006-01-02"))
	fmt.Printf("     • app.log → app.log.%s\n", today.Format("2006-01-02"))
	fmt.Printf("     • error.log → error.log.%s\n", today.Format("2006-01-02"))
	fmt.Printf("     • access.log → access.log.%s\n", today.Format("2006-01-02"))
	fmt.Printf("   - File baru akan dibuat untuk tanggal %s\n", tomorrow.Format("2006-01-02"))

	// 6. Contoh isi log dengan timestamp yang berbeda
	fmt.Println("\n📄 Contoh isi log dengan timestamp:")
	showLogContent()

	fmt.Println("\n✅ Demo selesai! Silakan periksa folder 'logs/demo' untuk melihat file-file yang dibuat.")
}

func createDemoLogFile(filename, logType string, date time.Time, isActive bool) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Printf("❌ Error creating %s: %v\n", filename, err)
		return
	}
	defer file.Close()

	// Sample log entries dengan timestamp sesuai tanggal
	var sampleLogs []string

	switch logType {
	case "app":
		sampleLogs = []string{
			fmt.Sprintf(`{
  "level": "info",
  "message": "Application started",
  "timestamp": "%s",
  "version": "1.0.0"
}`, date.Format("2006-01-02T15:04:05+07:00")),
			fmt.Sprintf(`{
  "level": "info", 
  "message": "Database migration completed",
  "timestamp": "%s",
  "duration_ms": 1500
}`, date.Add(1*time.Minute).Format("2006-01-02T15:04:05+07:00")),
		}
	case "error":
		sampleLogs = []string{
			fmt.Sprintf(`{
  "level": "error",
  "message": "Database connection timeout",
  "error": "dial tcp: i/o timeout",
  "timestamp": "%s",
  "type": "database_error"
}`, date.Format("2006-01-02T15:04:05+07:00")),
		}
	case "access":
		sampleLogs = []string{
			fmt.Sprintf(`{
  "level": "info",
  "message": "HTTP request",
  "method": "GET",
  "path": "/api/users",
  "status_code": 200,
  "latency_ms": 45,
  "timestamp": "%s",
  "type": "http_access"
}`, date.Format("2006-01-02T15:04:05+07:00")),
		}
	}

	for _, log := range sampleLogs {
		file.WriteString(log + "\n\n")
	}

	if isActive {
		fmt.Printf("✅ Created active log: %s\n", filepath.Base(filename))
	} else {
		fmt.Printf("📄 Created rotated log: %s\n", filepath.Base(filename))
	}
}

func showDirectoryStructure(dir string) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		fmt.Printf("❌ Error reading directory: %v\n", err)
		return
	}

	fmt.Printf("logs/demo/\n")
	for i, entry := range entries {
		connector := "├──"
		if i == len(entries)-1 {
			connector = "└──"
		}

		info, _ := entry.Info()
		size := info.Size()
		modTime := info.ModTime().Format("2006-01-02 15:04")

		if entry.Name()[len(entry.Name())-4:] == ".log" && len(entry.Name()) == 10 {
			// Active log files
			fmt.Printf("%s %s (active) [%d bytes, %s]\n", connector, entry.Name(), size, modTime)
		} else {
			// Rotated log files
			fmt.Printf("%s %s (rotated) [%d bytes, %s]\n", connector, entry.Name(), size, modTime)
		}
	}
}

func showLogContent() {
	examples := []struct {
		filename string
		date     string
	}{
		{"app.log.2025-07-10", "2025-07-10 (3 hari lalu)"},
		{"app.log.2025-07-12", "2025-07-12 (kemarin)"},
		{"app.log", "2025-07-13 (hari ini)"},
	}

	for _, example := range examples {
		fmt.Printf("\n📋 %s - %s:\n", example.filename, example.date)
		fmt.Println("   {")
		fmt.Println("     \"timestamp\": \"" + example.date[:10] + "T14:30:00+07:00\",")
		fmt.Println("     \"level\": \"info\",")
		fmt.Println("     \"message\": \"Application event\"")
		fmt.Println("   }")
	}
}
