#!/bin/bash

# 🩸 CRIMSON DARK CYBERPUNK PALITRASI
Reset="\033[0m"
Bold="\033[1m"
CrimsonRed="\033[38;5;196m"
BloodWine="\033[38;5;124m"
ElectricV="\033[38;5;129m"
BrightJade="\033[38;5;49m"
GhostGray="\033[38;5;244m"

# Root huquqini tekshirish
if [ "$EUID" -ne 0 ]; then
    echo -e "${CrimsonRed}[!] XATOLIK: Tizim darajasida shifrlash uchun skriptni 'sudo' bilan ishga tushiring!${Reset}"
    exit 1
fi

clear
echo -e "${GhostGray}[*] Tizim xotirasi tekshirilmoqda va eski osilgan oqimlar tozalanmoqda...${Reset}"
sudo killall crypto_core 2>/dev/null
sudo killall go 2>/dev/null

echo -e "${CrimsonRed}${Bold}"
echo "██╗  ██╗██╗   ██╗██████╗ ███████╗██████╗ ██╗ ██████╗ ███╗   ██╗"
echo "██║  ██║╚██╗ ██╔╝██╔══██╗██╔════╝██╔══██╗██║██╔═══██╗████╗  ██║"
echo "███████║ ╚████╔╝ ██████╔╝█████╗  ██████╔╝██║██║   ██║██╔██╗ ██║"
echo "██╔══██║  ╚██╔╝  ██╔═══╝ ██╔══╝  ██╔══██╗██║██║   ██║██║╚██╗██║"
echo "██║  ██║   ██║   ██║     ███████╗██║  ██║██║╚██████╔╝██║ ╚████║"
echo "╚═╝  ╚═╝   ╚═╝   ╚═╝     ╚══════╝╚═╝  ╚═╝╚═╝ ╚═════╝ ╚═╝  ╚═══╝"
echo -e "             ${ElectricV}[ ENGINE LAUNCHER ACTIVE ]${Reset}\n"

# Sinov papkasiga to'liq o'qish/yozish huquqini avtomat berish
if [ -d "test_files" ]; then
    sudo chmod -R 777 test_files/
fi

# Operatsiya vaqtida Ctrl+C bosilsa xotirani tozalash
cleanup() {
    echo -e "\n\n${CrimsonRed}[*] Launcher: Operatsiya majburiy to'xtatildi. Tizim tozalanmoqda...${Reset}"
    sudo killall crypto_core 2>/dev/null
    exit 0
}
trap cleanup SIGINT SIGTERM

# Go dasturini Linux tizim fayli ko'rinishida yig'ish va ishga tushirish
if [ -f "main.go" ]; then
    echo -e "${GhostGray}[*] Go kiber-yadrosi kompilyatsiya qilinmoqda...${Reset}"
    go build -o crypto_core main.go
    
    if [ -f "crypto_core" ]; then
        echo -e "${BrightJade}[+] Kompilyatsiya muvaffaqiyatli tugadi. Dastur nazorati topshirilmoqda...${Reset}"
        echo -e "${GhostGray}------------------------------------------------------------------------${Reset}"
        sleep 1
        
        # Dasturning o'zini root bilan ishga tushiramiz (Qolgani dastur ichida bo'ladi)
        sudo ./crypto_core
        
        # Dastur tugagach tozalash
        sudo rm -f crypto_core 2>/dev/null
    else
        echo -e "${CrimsonRed}[XATO] Go kodini yig'ib bo'lmadi! Sintaksis xatoliklarni tekshiring.${Reset}"
    fi
else
    echo -e "${CrimsonRed}[XATO] 'main.go' fayli joriy papkada topilmadi!${Reset}"
    exit 1
fi
