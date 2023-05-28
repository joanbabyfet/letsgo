package common

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/smtp"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/google/uuid"
	"github.com/mintance/go-uniqid"
	"github.com/spf13/viper"
)

// api成功响应
func Success(c *gin.Context, data interface{}, msg string) {
	if msg == "" {
		msg = "success"
	}
	if data == nil {
		data = make([]int, 0)
	}

	c.JSON(http.StatusOK, gin.H{
		"code":      0,
		"msg":       msg,
		"timestamp": time.Now().Unix(),
		"data":      data,
	})
	c.Abort() //目前的handler仍會執行到結束
}

// api失败响应
func Error(c *gin.Context, msg string, code int, data interface{}) {
	if msg == "" {
		msg = "error"
	}
	if code >= 0 {
		code = -1
	}
	if data == nil {
		data = make([]int, 0)
	}

	c.JSON(http.StatusOK, gin.H{
		"code":      code,
		"msg":       msg,
		"timestamp": time.Now().Unix(),
		"data":      data,
	})
	c.Abort() //目前的handler仍會執行到結束
}

// 生成唯一识别
func Random(c *gin.Context, method string, length int) string {
	if method == "" {
		method = "web"
	}
	if length == 0 {
		length = 32
	}
	str := ""

	switch method {
	case "web":
		// 即使同一个IP，同一款浏览器，要在微妙内生成一样的随机数，也是不可能的
		// 进程ID保证了并发，微妙保证了一个进程每次生成都会不同，IP跟AGENT保证了一个网段
		// md5(当前进程id + 在目前微秒时间生成唯一id + 当前ip + 当前浏览器)
		remote_addr := c.ClientIP()                                //当前ip
		user_agent := c.GetHeader("User-Agent")                    //当前浏览器
		microtime := strconv.FormatInt(time.Now().UnixMicro(), 10) //当前微秒时间生成唯一id
		pid := strconv.Itoa(os.Getpid())                           //当前进程id
		str = MD5(pid + GenUniqid(MD5(microtime)) + remote_addr + user_agent)
	}
	return str
}

// 获取md5哈希
func MD5(str string) string {
	data := []byte(str)
	return fmt.Sprintf("%x", md5.Sum(data)) //将[]byte转成16进制
}

// 类似php的uniqid()
func GenUniqid(str string) string {
	return uniqid.New(uniqid.Params{str, true})
}

// 生成UUID
func GenUUID() string {
	return uuid.New().String()
}

// 发送邮件
// TODO tls待优化
func SendMail(to string, subject string, body string, is_html bool) int {
	host := viper.GetString("mail.host")
	port := viper.GetString("mail.port")
	username := viper.GetString("mail.username")
	password := viper.GetString("mail.password")
	from_address := viper.GetString("mail.from_address")
	from_name := viper.GetString("mail.from_name")
	var content_type string
	status := 1

	//身份认证时不使用tls或ssl加密
	auth := smtp.PlainAuth("", username, password, host)
	if is_html {
		content_type = "Content-Type: text/html; charset=UTF-8"
	} else {
		content_type = "Content-Type: text/plain; charset=UTF-8"
	}
	msg := []byte("To: " + to + "\r\nFrom: " + from_name + "<" + from_address + ">" + "\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	arr_to := strings.Split(to, ";")
	err := smtp.SendMail(host+":"+port, auth, from_address, arr_to, msg)
	if err != nil {
		fmt.Println(err)
		status = -1
	}
	return status
}

// 发送telegram bot通知
func SendTelegramBot(text string, chat_id int64) int {
	status := 1
	token := viper.GetString("telegram_bot.token")

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		fmt.Println(err)
	}
	bot.Debug = false
	fmt.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	msg := tgbotapi.NewMessage(chat_id, text)
	if _, err := bot.Send(msg); err != nil {
		fmt.Println(err)
		status = -1
	}
	return status
}

// 发送line bot通知
func SendLineBot(text string) int {
	token := viper.GetString("line_bot.token")
	api_url := "https://notify-api.line.me/api/notify"
	status := 1

	data := url.Values{}
	data.Add("message", text)

	req, err := http.NewRequest("POST", api_url, strings.NewReader(data.Encode()))
	if err != nil {
		status = -1
		fmt.Println(err)
	}
	//设置请求头参数
	req.Header.Add("Authorization", "Bearer "+token)
	//内容类型不加会报错
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	//发送http请求
	client := &http.Client{}
	resp, err := client.Do(req)
	//关闭请求
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	if err != nil {
		status = -2
		fmt.Println("handle error")
	}
	return status
}
