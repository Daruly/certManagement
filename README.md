# certManagement

# Myserver — Production-ready HTTPS server with automatic TLS from Let's Encrypt

Проект: Реальный HTTPS сервер на Go с автоматическим получением и обновлением TLS сертификатов через Let's Encrypt (autocert).

## Быстрый старт

### 1. Клон

git clone https://github.com/Daruly/certManagement.git
cd certManagement
cd myserver


---

### 2. Установка зависимостей

- Установите Go >= 1.21 
- Установите зависимости (автоматически при запуске):

go mod tidy


---

### 3. Создание переменных окружения

**Требуется реальный домен, направленный на ваш сервер (A-запись в DNS)!**

- Укажите домены через запятую (пример: example.com,www.example.com)
- Email нужен для уведомлений от Let's Encrypt (например, истекает срок сертификата).


#### env PowerShell

$env:DOMAINS="example.com,www.example.com"
$env:EMAIL="your-email@example.com"


---

### 4. Запуск сервера

> Требуются права администратора для портов 80 и 443 (на Windows — запускать PowerShell от имени администратора).

go run main.go


---

### 5. Что происходит при запуске

- Сервер автоматически стартует:
  - HTTP на порту 80 (для Let's Encrypt ACME challenge)
  - HTTPS на порту 443 (основной сервер, выдаёт страницы и API)
- Сертификаты Let's Encrypt сохраняются в папку `certs/` и автоматически обновляются!
- Сервер корректно завершает работу по сигналу остановки (Ctrl+C).

---

### 6. Проверка результата

#### Через браузер

Откройте:

https://example.com/
https://example.com/health
https://example.com/api/info


Браузер должен показать зелёный замочек (действительный сертификат Let's Encrypt).

#### Через curl

curl https://example.com/
curl https://example.com/health
curl https://example.com/api/info