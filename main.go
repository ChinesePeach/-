package main

import (
	"bufio"
	"fmt"
	"github.com/jordan-wright/email"
	"github.com/xuri/excelize/v2"
	"io"
	"log"
	"net/smtp"
	"os"
	"strings"
	"time"
)

func readProHtml() (Project string, Html []byte) {
	file, err := os.Open("./ProjectHtml.txt")
	if err != nil {
		fmt.Println("ProjectHtml.txt文件打开失败 = ", err)
		os.Exit(1)
	}
	defer file.Close()
	//创建一个 *Reader ， 是带缓冲的
	reader := bufio.NewReader(file)
	tempNum := 1
	for {
		str, err := reader.ReadString('\n') //读到一个换行就结束
		if tempNum == 1 {
			Project, tempNum = str, tempNum+1
		} else {
			Html = append(Html, []byte(str)...)
		}
		if err == io.EOF { //io.EOF 表示文件的末尾
			break
		}
	}
	fmt.Println("ProjectHtml.txt文件读取结束...")
	return
}

func SendEmail(toEmail []string, subject string, Html []byte, from string, emailPassword string, emailAddress string, smtpPort string, smtpAddress string, attachFile int, attachFileName string) {
	// 简单设置l og 参数
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	em := email.NewEmail()
	//设置发送方的邮箱
	//em.From = "广西大学程序设计竞赛组委会 <gxucpc@163.com>"
	em.From = from
	// 设置接收方的邮箱
	em.To = toEmail
	//设置主题
	em.Subject = subject
	//设置文件发送的内容
	em.HTML = Html
	if attachFile == 1 {
		// 添加附件
		em.AttachFile("./" + attachFileName)
	}
	// 设置服务器相关的配置
	// emailAddress = gxucpc@163.com Password = CJDDNSQCBGDOPXTL

	err := em.Send(smtpAddress+":"+smtpPort, smtp.PlainAuth("", emailAddress, emailPassword, smtpAddress))
	if err != nil {
		log.Fatal(err)
		time.Sleep(5 * time.Second)
	}
	log.Println("send successfully ... ")
}

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("请输入运营商smtp服务器地址：")
	smtpAddress, err := reader.ReadString('\n')
	smtpAddress = strings.TrimSuffix(smtpAddress, "\n")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("请输入运营商smtp服务器端口：")
	smtpPort, err := reader.ReadString('\n')
	smtpPort = strings.TrimSuffix(smtpPort, "\n")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("请输入发件人邮箱：")
	emailAddress, err := reader.ReadString('\n')
	emailAddress = strings.TrimSuffix(emailAddress, "\n")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("请输入邮箱的授权码（不是邮箱登入密码）")
	emailPassword, err := reader.ReadString('\n')
	emailPassword = strings.TrimSuffix(emailPassword, "\n")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("请输入邮件From,例如:广西大学程序设计竞赛组委会 <gxucpc@163.com>")
	emailFrom, err := reader.ReadString('\n')
	emailFrom = strings.TrimSuffix(emailFrom, "\n")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	attachExit, attachFileName := 0, ""
	fmt.Println("是否存在附件（1: 存在  0: 不存在")
	fmt.Scanln(&attachExit)
	if attachExit == 1 {
		fmt.Println("请输入附件名称（建议使用英文名，汉字会乱码！| 名称写完整，包含文件类型：123.pdf）")
		attachFileName, err = reader.ReadString('\n')
		attachFileName = strings.TrimSuffix(attachFileName, "\n")
	}

	fmt.Println("请确认是否在同级目录下按照README创建对应的xlsx文件和txt文件")

	f, err := excelize.OpenFile("email.xlsx")
	if err != nil {
		println(err.Error())
		time.Sleep(5 * time.Second)
		return
	}

	Project, Html := readProHtml()

	rows, err := f.GetRows("Sheet1")
	if err != nil {
		fmt.Println(err)
		time.Sleep(5 * time.Second)
		return
	}

	for _, row := range rows {
		for _, colCell := range row {
			temp := []string{colCell}
			// 163限制15min不得超过200封邮件，设置延迟，平均一分钟10封
			// 如果对邮件数量不做要求时，可以直接把所有邮件放到temp里，不要team只放一个邮箱
			time.Sleep(5 * time.Second)
			SendEmail(temp, Project, Html, emailFrom, emailPassword, emailAddress, smtpPort, smtpAddress, attachExit, attachFileName)
		}
	}
	time.Sleep(time.Second)
	fmt.Println("邮件发送完成！")
	time.Sleep(5 * time.Second)
}
