# 🩸 Hyperion Crypto Engine (Ultra Stream AES-256-CTR)

Ko'p tarmoqli (**Multi-threaded Golang Core**) va oqimli (**Stream Cipher Architecture**) shifrlash tizimi hamda uni boshqaruvchi **Advanced Bash Wrapper** (Launcher).

Ushbu loyiha istalgan hajmdagi (hatto 100GB+) ulkan fayl yoki papkalarni operativ xotirani (**RAM**) ortiqcha yuklamasdan, **32KB gacha bo'lgan kichik bufer quvurlari** (io.CopyBuffer) orqali **AES-256-CTR** algoritmida parallel ravishda soniyalar ichida shifrlaydi (`.locked`) va qayta tiklaydi.

---

## 🚀 Asosiy Imkoniyatlari

*   **RAM-Friendly Stream Cipher:** Fayllarni butunlay xotiraga yuklamaydi. 32KB li kichik bloklarda oqimli shifrlagani sababli past ko'rsatkichli qurilmalarda ham cheksiz hajmdagi fayllarni ko'tara oladi.
*   **AES-256-CTR & Dinamik IV:** Har bir shifrlanadigan fayl uchun mutlaqo takrorlanmas tasodifiy **IV (Initialization Vector)** blok yaratiladi va shifrlangan faylning boshiga xavfsiz tarzda joylanadi.
*   **Maksimal Yadrolar Integratsiyasi:** Go-lang goroutinalari `runtime.NumCPU() - 1` formulasiga ko'ra tizimdagi barcha yadrolarni parallel oqimlarga (Threads) safarbar etadi.
*   **Avtomatik Resurs Optimizatsiyasi:** Linux tizimidagi `ulimit` (Open Files limit) parametrlarini maksimal tezlik uchun avtomatik **65,535** ga ko'taradi.
*   **Xavfsiz Launcher (Bash Wrapper):** Fondagi osilib qolgan eski jarayonlarni tozalaydi, maqsadli papkaga (`test_files/`) ruxsat huquqlarini to'g'rilaydi va `Ctrl+C` bosilganda jarayonni xavfsiz to'xtatadi.

---

## 🖥️ Terminal Interfeysi (Interface Preview)

Dasturning terminaldagi ish faoliyati va chiroyli interaktiv menyusi quyidagicha ko'rinishga ega:

### 1. Launcher va Kompilyatsiya Jarayoni
Dastur `sudo` orqali ishga tushirilganda tizim xotirasini tozalaydi, Go kodini avtomatik Linux binary holatiga yig'adi va nazoratni kiber-yadroga topshiradi:

```text
[*] Tizim xotirasi tekshirilmoqda va eski osilgan oqimlar tozalanmoqda...
██╗  ██╗██╗   ██╗██████╗ ███████╗██████╗ ██╗ ██████╗ ███╗   ██╗
██║  ██║╚██╗ ██╔╝██╔══██╗██╔════╝██╔══██╗██║██╔═══██╗████╗  ██║
███████║ ╚████╔╝ ██████╔╝█████╗  ██████╔╝██║██║   ██║██╔██╗ ██║
██╔══██║  ╚██╔╝  ██╔═══╝ ██╔══╝  ██╔══██╗██║██║   ██║██║╚██╗██║
██║  ██║   ██║   ██║     ███████╗██║  ██║██║╚██████╔╝██║ ╚████║
╚═╝  ╚═╝   ╚═╝   ╚═╝     ╚══════╝╚═╝  ╚═╝╚═╝ ╚═════╝ ╚═╝  ╚═══╝
             [ ENGINE LAUNCHER ACTIVE ]

[*] Go kiber-yadrosi kompilyatsiya qilinmoqda...
[+] Kompilyatsiya muvaffaqiyatli tugadi. Dastur nazorati topshirilmoqda...
------------------------------------------------------------------------
```

### 2. Tizim Asosiy Menyusi (Core Dashboard)
Dastur ishga tushgach, foydalanuvchiga interaktiv boshqaruv panelini taqdim etadi:

```text
██╗  ██╗██╗   ██╗██████╗ ███████╗██████╗ ██╗ ██████╗ ███╗   ██╗
██║  ██║╚██╗ ██╔╝██╔══██╗██╔════╝██╔══██╗██║██╔═══██╗████╗  ██║
███████║ ╚████╔╝ ██████╔╝█████╗  ██████╔╝██║██║   ██║██╔██╗ ██║
██╔══██║  ╚██╔╝  ██╔═══╝ ██╔══╝  ██╔══██╗██║██║   ██║██║╚██╗██║
██║  ██║   ██║   ██║     ███████╗██║  ██║██║╚██████╔╝██║ ╚████║
╚═╝  ╚═╝   ╚═╝   ╚═╝     ╚══════╝╚═╝  ╚═╝╚═╝ ╚═════╝ ╚═╝  ╚═══╝
             [ V3.0 - ULTRA STREAM ENGINE ACTIVE ]

┌───[ CYBER ENGINE CORES ENGINE ]────────────────────┐
│  FAYLLARNI SHIFRLASH (LAUNCH ENCRYPTION)       │
│  FAYLLARNI QAYTA TIKLASH (LAUNCH DECRYPTION)   │
│  SESHINDAN TO'LIQ CHIQISH (EXIT ENGINE)        │
└────────────────────────────────────────────────────┘

┌── Amalni tanlang:
└── » 1

┌── Nishon papka nomini kiriting [Standart: test_files]:
└── » test_files
```

### 3. Parallel Shifrlash/Tiklash Operatsiyasi (Parallel Workers)
Rejim tanlangandan so'ng, tizim barcha CPU yadrolarini ishga solib, parallel ravishda fayllarni qayta ishlaydi va jonli loglarni chiqaradi:

```text
┌───[ OPERATSIYA PARAMETRLARI ]─────────────────────┐
│ Rejim     => ENCRYPT
│ Nishon    => /home/user/tools/crypto/test_files
│ Workers   => 7 Parallel Threads
└────────────────────────────────────────────────────┘

[SHIFRLANDI] ───> backup_database.sql
[SHIFRLANDI] ───> confidential_report.pdf
[SHIFRLANDI] ───> archive_2026.tar.gz

[+] OPERATSIYA YAKUNLANDI!
[*] Asosiy menyuga qaytish uchun [Enter] bosing...
```

---

## 🛠️ O'rnatish va Talablar

Skript to'g'ri ishlashi uchun kompyuteringizda **Go (Golang)** muhiti o'rnatilgan bo'lishi hamda sizda Linux tizimida **root (sudo)** huquqi bo'lishi lozim.

### 1. Golang muhitini o'rnatish:
```bash
sudo apt update
sudo apt install golang -y
```

### 2. Loyihani sozlash:
`main.go` va `run_crypto.sh` fayllari bitta papkada joylashganligiga ishonch hosil qiling va quyidagi buyruqlarni bering:
```bash
git clone https://github.com
cd LOYIHA-NOMI
chmod +x run_crypto.sh
```

---

## 🚀 Foydalanish

Tizimni ishga tushirish uchun shunchaki skriptni `sudo` bilan ishga tushiring:
```bash
sudo ./run_crypto.sh
```
Papkalar yoki fayllar ustida amallarni bajarish dasturning chiroyli interaktiv menyusi orqali boshqariladi.

---

## ⚠️ OGOHLANTIRISH (DISCLAIMER)

Ushbu vosita **faqat ta'lim, kiberxavfsizlikni o'rganish, ma'lumotlarni himoya qilish va kriptografik auditlar** uchun yaratilgan. Muallif ushbu dasturdan ransomware (tovlamachilik dasturlari) yoki boshqa har qanday noqonuniy zararli maqsadlarda foydalanish natijasida kelib chiqadigan oqibatlar uchun hech qanday javobgarlikni o'z zimmasiga olmaydi! Loyihadan faqat qonuniy va ruxsat berilgan doirada foydalaning.

---

## 🧑‍💻 Muallif va Aloqa

*   **Sizning Ismingiz** - [@SizningTelegramingiz](https://t.me/adhamjon0717)
*   GitHub: [@SizningGitHubingiz](https://github.com/Pydeveloperman)
