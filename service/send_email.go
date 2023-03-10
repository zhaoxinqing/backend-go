package service

import (
	"io/ioutil"
	"log"

	"github.com/go-gomail/gomail"
)

type EmailInfo struct {
	ServerHost string // ServerHost 邮箱服务器地址，如腾讯企业邮箱为smtp.exmail.qq.com
	ServerPort int    // ServerPort 邮箱服务器端口，如腾讯企业邮箱为465

	FromEmail  string // FromEmail　发件人邮箱地址
	FromPasswd string //发件人邮箱密码（注意，这里是明文形式)

	Recipient []string //收件人邮箱
	CC        []string //抄送
}

var emailMessage *gomail.Message

/**
 * @Author: dcj
 * @Date: 2020-04-02 15:45:55
 * @Description: 发送邮件
 * @Param : subject[主题]、body[内容]、emailInfo[发邮箱需要的信息(参考EmailInfo)]
 * @Return:
 */
func SendEmail(body string) {

	emailInfo := &EmailInfo{
		ServerHost: "smtp.qq.com",
		ServerPort: 465,

		FromEmail:  "1375565592@qq.com",
		FromPasswd: "gyyibickelsvgdae",

		Recipient: []string{"ahead_zhaoxinqing@163.com"},
		CC:        []string{},
	}
	subject := "今日热点推送"
	// body = time.Now().Format("2006.01.02 15:04:05") + ": 知乎热榜、微博热榜、百度热点"

	emailMessage = gomail.NewMessage()
	//设置收件人
	emailMessage.SetHeader("To", emailInfo.Recipient...)
	//设置抄送列表
	if len(emailInfo.CC) != 0 {
		emailMessage.SetHeader("Cc", emailInfo.CC...)
	}
	// 第三个参数为发件人别名，如"dcj"，可以为空（此时则为邮箱名称）
	emailMessage.SetAddressHeader("From", emailInfo.FromEmail, "@LiaoWzqb")

	//主题
	emailMessage.SetHeader("Subject", subject)

	//正文
	// str := ReadFile("./doc/email.html")
	emailMessage.SetBody("text/html", body)

	d := gomail.NewDialer(emailInfo.ServerHost, emailInfo.ServerPort,
		emailInfo.FromEmail, emailInfo.FromPasswd)
	err := d.DialAndSend(emailMessage)
	if err != nil {
		log.Println("发送邮件失败： ", err)
	} else {
		log.Println("已成功发送邮件到指定邮箱")
	}
}

func ReadFile(filepath string) (str string) {
	bytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatal("Failed to read file: " + filepath)
	}
	return string(bytes)
}
