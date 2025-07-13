//go:build example
// +build example

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func simulateTomorrowRotation() {
	fmt.Println("🔄 SIMULASI: Apa yang terjadi ketika tanggal berganti ke besok")
	fmt.Println(strings.Repeat("=", 60))

	// Current state - hari ini 2025-07-13
	today := "2025-07-13"
	tomorrow := "2025-07-14"

	fmt.Printf("\n📅 Skenario: Berganti dari %s ke %s\n", today, tomorrow)

	// 1. Sebelum rotasi (hari ini)
	fmt.Println("\n🕐 SEBELUM ROTASI (Malam hari 2025-07-13):")
	fmt.Println("logs/")
	fmt.Println("├── app.log          # Log aktif hari ini")
	fmt.Println("├── error.log        # Error log aktif")
	fmt.Println("├── access.log       # Access log aktif")
	fmt.Println("├── app.log.2025-07-12    # Kemarin")
	fmt.Println("├── error.log.2025-07-12  # Kemarin")
	fmt.Println("└── access.log.2025-07-12 # Kemarin")

	// 2. Proses rotasi (tengah malam)
	fmt.Println("\n🕐 SAAT ROTASI (00:00:00 tanggal 2025-07-14):")
	fmt.Println("Lumberjack akan melakukan:")
	fmt.Println("1. Rename: app.log → app.log.2025-07-13")
	fmt.Println("2. Rename: error.log → error.log.2025-07-13")
	fmt.Println("3. Rename: access.log → access.log.2025-07-13")
	fmt.Println("4. Create: app.log (baru untuk tanggal 2025-07-14)")
	fmt.Println("5. Create: error.log (baru)")
	fmt.Println("6. Create: access.log (baru)")

	// 3. Setelah rotasi (besok)
	fmt.Println("\n🕐 SETELAH ROTASI (Pagi hari 2025-07-14):")
	fmt.Println("logs/")
	fmt.Println("├── app.log               # Log aktif BESOK")
	fmt.Println("├── error.log             # Error log aktif BESOK")
	fmt.Println("├── access.log            # Access log aktif BESOK")
	fmt.Println("├── app.log.2025-07-13    # Kemarin (yang tadi)")
	fmt.Println("├── error.log.2025-07-13  # Kemarin")
	fmt.Println("├── access.log.2025-07-13 # Kemarin")
	fmt.Println("├── app.log.2025-07-12    # 2 hari lalu")
	fmt.Println("├── error.log.2025-07-12  # 2 hari lalu")
	fmt.Println("└── access.log.2025-07-12 # 2 hari lalu")

	// 4. Contoh timestamp dalam log
	fmt.Println("\n📄 TIMESTAMP DALAM LOG FILES:")
	fmt.Printf("\n• app.log.%s:\n", today)
	fmt.Printf("  {\n")
	fmt.Printf("    \"timestamp\": \"%sT23:59:59+07:00\",\n", today)
	fmt.Printf("    \"message\": \"Last log entry of the day\"\n")
	fmt.Printf("  }\n")

	fmt.Printf("\n• app.log (file baru untuk %s):\n", tomorrow)
	fmt.Printf("  {\n")
	fmt.Printf("    \"timestamp\": \"%sT00:00:01+07:00\",\n", tomorrow)
	fmt.Printf("    \"message\": \"First log entry of new day\"\n")
	fmt.Printf("  }\n")

	// 5. Demonstrasi dengan membuat file actual
	fmt.Println("\n🎯 DEMO PRAKTIS - Membuat file contoh:")
	createActualDemoFiles(today, tomorrow)
}

func createActualDemoFiles(today, tomorrow string) {
	demoDir := "logs/tomorrow-demo"
	os.MkdirAll(demoDir, 0755)

	// File untuk "hari ini" (akan di-rotate)
	todayFiles := []string{
		fmt.Sprintf("app.log.%s", today),
		fmt.Sprintf("error.log.%s", today),
		fmt.Sprintf("access.log.%s", today),
	}

	// File untuk "besok" (file baru)
	tomorrowFiles := []string{"app.log", "error.log", "access.log"}

	// Buat file "kemarin"
	for _, filename := range todayFiles {
		filepath := filepath.Join(demoDir, filename)
		file, _ := os.Create(filepath)
		content := fmt.Sprintf(`{
  "timestamp": "%sT23:59:59+07:00",
  "level": "info",
  "message": "Last entry of %s",
  "rotated": true
}`, today, today)
		file.WriteString(content)
		file.Close()
		fmt.Printf("✅ Created rotated file: %s\n", filename)
	}

	// Buat file "besok"
	for _, filename := range tomorrowFiles {
		filepath := filepath.Join(demoDir, filename)
		file, _ := os.Create(filepath)
		content := fmt.Sprintf(`{
  "timestamp": "%sT00:00:01+07:00", 
  "level": "info",
  "message": "First entry of %s",
  "active": true
}`, tomorrow, tomorrow)
		file.WriteString(content)
		file.Close()
		fmt.Printf("✅ Created active file: %s\n", filename)
	}

	fmt.Printf("\n📁 Files created in: %s\n", demoDir)

	// Show directory structure
	entries, _ := os.ReadDir(demoDir)
	fmt.Println("\nStructure:")
	fmt.Printf("%s/\n", demoDir)
	for i, entry := range entries {
		connector := "├──"
		if i == len(entries)-1 {
			connector = "└──"
		}

		if len(entry.Name()) > 10 && entry.Name()[len(entry.Name())-10:len(entry.Name())-4] == today {
			fmt.Printf("%s %s (rotated from yesterday)\n", connector, entry.Name())
		} else {
			fmt.Printf("%s %s (active for today)\n", connector, entry.Name())
		}
	}
}

func main() {
	simulateTomorrowRotation()

	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println("🎯 KESIMPULAN:")
	fmt.Println("• Log rotation terjadi otomatis saat berganti hari")
	fmt.Println("• File aktif (tanpa tanggal) selalu untuk hari ini")
	fmt.Println("• File dengan tanggal adalah file yang sudah di-rotate")
	fmt.Println("• Timestamp dalam log menunjukkan kapan event terjadi")
	fmt.Println("• Lumberjack menangani semua proses rotation secara otomatis")
	fmt.Println("\n✅ Periksa folder 'logs/tomorrow-demo' untuk melihat contoh!")
}
