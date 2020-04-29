package mail

import (
	"strconv"

	"github.com/eth_tools/util"
	"gopkg.in/gomail.v2"
)

// SendMail.
func SendMail(mailTo []string, subject string, body string) error {
	myConfig := new(util.Config)
	myConfig.InitConfig("common.conf")

	//定义邮箱服务器连接信息，如果是阿里邮箱 pass填密码，qq邮箱填授权码
	mailConn := map[string]string{
		"user": myConfig.Read("mail", "user"),
		"pass": myConfig.Read("mail", "pwd"),
		"host": myConfig.Read("mail", "host"),
		"port": myConfig.Read("mail", "port"),
	}

	port, _ := strconv.Atoi(mailConn["port"]) //转换端口类型为int

	m := gomail.NewMessage()
	m.SetHeader("From", "XD Game"+"<"+mailConn["user"]+">") //这种方式可以添加别名，即“XD Game”， 也可以直接用<code>m.SetHeader("From",mailConn["user"])</code> 读者可以自行实验下效果
	m.SetHeader("To", mailTo...)                            //发送给多个用户
	m.SetHeader("Subject", subject)                         //设置邮件主题
	m.SetBody("text/html", body)                            //设置邮件正文

	d := gomail.NewDialer(mailConn["host"], port, mailConn["user"], mailConn["pass"])

	err := d.DialAndSend(m)
	return err

}
