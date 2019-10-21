package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	"github.com/jlaffaye/ftp"
)

func main() {

	c, err := ftp.Dial("185.20.227.83:1212", ftp.DialWithTimeout(5*time.Second))
	if err != nil {
		log.Fatal(err)
	}

	err = c.Login("user", "password")
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("connected")
	}

	// для ввода команд
	myscanner := bufio.NewScanner(os.Stdin)
	myscanner.Scan()
	line := myscanner.Text()
	doparray := strings.Fields(line)
	i := -1
	for i < 1 {
		switch doparray[0] {
		case "mkdir":
			if err := c.MakeDir(doparray[1]); err != nil {
				fmt.Println(err)
			}
			break
		case "ls":
			if list, err := c.NameList("/"); err != nil {
				log.Fatal(err)
			} else {
				log.Printf("List: %v", list)
			}
			break
		case "rmdir":
			name := "/" + doparray[1]
			if err := c.RemoveDir(name); err != nil {
				log.Printf("Del dir: %s", err.Error())
			}
			break
		case "rm":
			name := "/" + doparray[1]
			err := c.Delete(name)
			if err != nil {
				log.Printf("Del file: %s", err.Error())
			}
			break
		case "download":
			name := "/" + doparray[1]
			res, err := c.Retr(name)
			if err != nil {
				fmt.Println(err.Error())
				break
			}
			buf, err := ioutil.ReadAll(res)
			if err != nil {
				fmt.Println(err)
				break
			}
			fmt.Println(string(buf))
			break
		case "upload":
			name := "/" + doparray[1]
			err := c.Stor(name, os.Stdin)
			if err != nil {
				log.Println(err)
			} else {
				fmt.Println("File uploaded")
			}
			break
		case "quit":
			if err := c.Quit(); err != nil {
				log.Printf("Quit: %s", err.Error())
				return
			}
			i = 1
			break
		}
		myscanner.Scan()
		line = myscanner.Text()
		doparray = strings.Fields(line)
	}
}
