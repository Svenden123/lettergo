package main
import (
	"fmt"
	"math/rand"
	"net/smtp"
	"strconv"
	"time"
)

func main() {
	otp:=time.Now().Second()*rand.Intn(100)
	var n int
	var name,email string
	fmt.Printf("Enter Name ")
	fmt.Scanln(&name)
	fmt.Printf("Enter Email ")
	fmt.Scanln(&email)
	fmt.Println("Otp= ", otp)
	send(strconv.Itoa(otp), email)
	fmt.Println(&n)


}

func send(body string, to string) {
	from:="ivfert4499@gmail.com"
	pass:="volkodav123"
	msg:="Subject: Your personal OTP\n\n"+body
	err:=smtp.SendMail("smtp.gmail.com:587",
	smtp.PlainAuth("",from,pass,"smtp.gmail.com"),
	from, []string{to}, []byte(msg))

	if err!=nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Hurray Email Sent")
}



