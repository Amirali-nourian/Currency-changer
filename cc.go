//تبدیل ارز به صورت آنلاین

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

// ساختار برای دریافت پاسخ API
type ExchangeRates struct {
	Rates map[string]float64 `json:"rates"`
}

const apiURL = "https://api.exchangerate-api.com/v4/latest/USD" // API نرخ ارز

func main() {
	// دریافت نرخ ارز از API
	response, err := http.Get(apiURL)
	if err != nil {
		log.Fatal("❌ خطا در دریافت داده‌ها:", err)
	}
	defer response.Body.Close()

	// پارس کردن داده‌های JSON
	var rates ExchangeRates
	if err := json.NewDecoder(response.Body).Decode(&rates); err != nil {
		log.Fatal("❌ خطا در پردازش داده‌ها:", err)
	}

	// دریافت ورودی از کاربر
	var amount float64
	var fromCurrency, toCurrency string

	fmt.Print("💰 مقدار پول را وارد کنید: ")
	fmt.Scanln(&amount)

	fmt.Print("🔄 ارز مبدا (مثلاً USD): ")
	fmt.Scanln(&fromCurrency)
	fromCurrency = strings.ToUpper(fromCurrency) // 🔥 اصلاح تبدیل به حروف بزرگ

	fmt.Print("🔄 ارز مقصد (مثلاً EUR): ")
	fmt.Scanln(&toCurrency)
	toCurrency = strings.ToUpper(toCurrency) // 🔥 اصلاح تبدیل به حروف بزرگ

	// بررسی اینکه ارزها در داده‌های API وجود دارند
	fromRate, fromExists := rates.Rates[fromCurrency]
	toRate, toExists := rates.Rates[toCurrency]

	if !fromExists || !toExists {
		fmt.Println("❌ ارز وارد شده نامعتبر است! لطفاً واحدهای ارزی معتبر وارد کنید.")
		os.Exit(1)
	}

	// محاسبه تبدیل ارز
	convertedAmount := (amount / fromRate) * toRate
	fmt.Printf("✅ %.2f %s معادل %.2f %s است.\n", amount, fromCurrency, convertedAmount, toCurrency)
}
