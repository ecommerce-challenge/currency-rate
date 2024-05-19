/*
 * GSES BTC application
 *
 *  Тобі потрібно реалізувати сервіс з АРІ, який дозволить:  - дізнатись поточний курс долара (USD) у гривні (UAH) - підписати емейл на отримання інформації по зміні курсу.  Додаткові вимоги: 1. Сервіс має відповідати описаному ниже АРІ. <i>NB Закривати рішення аутентифікацією не потрібно</i>. 2. Eмейли з актуальний курсом мають відправлятись юзерам 1 раз на добу.  2. Всі дані, для роботи додатку повинні зберігатися в базі даних. Також потрібно реалізувати виконання міграції структури БД при піднятті сервісу. 3. В репозиторії повинен бути Dockerfile, та Docker Compose який дозволяє запустити систему в Docker. З матеріалом по Docker вам необхідно ознайомитись самостійно. 4. Також ти можеш додавати коментарі чи опис логіки виконання роботи в README.md документі. Правильна логіка може стати перевагою при оцінюванні, якщо ти не повністю виконаєш завдання.  Очікувані мови виконання завдання: PHP, Go, NodeJS.  Виконувати завдання іншими мовами можна, проте, це не буде перевагою. Виконане завдання необхідно завантажити на GitHub (публічний репозиторій) та сабмітнути виконане завдання в гугл-форму.  Ти можеш користуватися усією доступною інформацією, але виконуй завдання самостійно.  Успіхів!
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package exchangerate

import (
	"log"
	"net/http"
	"time"
)

func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		log.Printf(
			"%s %s %s %s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}
