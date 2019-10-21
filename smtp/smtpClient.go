package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/smtp"
	"os"
	"strconv"
)

func Decrypt(text string, key string, size int) string {
	key_index := 0
	key_len := len(key)
	var decrypted string
	var old_code, new_code int

	for _, ch := range text {
		shift, err := strconv.Atoi(string(key[key_index]))
		if err != nil {
			log.Fatalf("Key must be number. %s", err)
		}
		old_code = int(ch)
		new_code = old_code - shift

		if ch >= 'a' && ch <= 'z' && new_code < int('a') ||
			ch >= 'A' && ch <= 'Z' && new_code < int('A') {
			new_code += size
		}

		decrypted += string(new_code)

		key_index++
		if key_index >= key_len {
			key_index = 0
		}
	}

	return decrypted

}

func main() {
	var email, message, title string
	password, _ := ioutil.ReadFile("/Users/go/src/smtp/password")
	log.Print(string(password))
	result := Decrypt(string(password), string(3), 26)
	log.Print(result)
	log.Print("write email to send")
	fmt.Scan(&email)
	auth := smtp.PlainAuth("", "kovajkin2011@yandex.ru", result, "smtp.gmail.com")
	log.Print("title:")
	fmt.Scan(&title)
	log.Print("message:")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	fmt.Println(text)
	log.Print(message)
	to := []string{email}
	msg := []byte("To:" + email + "\r\n" + "Subject: " + title + "\r\n" + "\r\n" + text + "\r\n")
	err := smtp.SendMail("smtp.gmail.com:587", auth, "kovajkin2011@yandex.ru", to, msg)
	if err != nil {
		log.Fatal(err)
	}
}
