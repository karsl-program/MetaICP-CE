package responses

import (
	"fmt"
	"io"
	"math/rand"
	"metaicp/jsons"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func RandVerifyCode() string {
	return fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
}

func GetNotice() string {
	config_file, _ := os.Open("notice")
	config_str, _ := io.ReadAll(config_file)
	config_file.Close()
	return string(config_str)
}

func GetDomain() string {
	config_file, _ := os.Open("domain")
	config_str, _ := io.ReadAll(config_file)
	config_file.Close()
	return string(config_str)
}

func GetPwd() string {
	config_file, _ := os.Open("pwd")
	config_str, _ := io.ReadAll(config_file)
	config_file.Close()
	return string(config_str)
}

func IndexNode(c *gin.Context) {
	isnotice := 1
	if GetNotice() == "" {
		isnotice = 0
	}
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title":    "首页",
		"mail":     GetMail(),
		"notice":   GetNotice(),
		"isnotice": isnotice,
	})
}

func AboutNode(c *gin.Context) {
	c.HTML(http.StatusOK, "about.html", gin.H{
		"title": "关于",
	})
}

func LoginNode(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{
		"title": "登录",
	})
}

func AdminNode(c *gin.Context) {
	if c.PostForm("pwd") != GetPwd() {
		c.String(http.StatusUnauthorized, "invalid password!")
		return
	}

	datas := jsons.ReadAllJSON().Datas

	c.HTML(http.StatusOK, "admin.html", gin.H{
		"title": "管理",
		"data":  datas,
		"pwd":   c.PostForm("pwd"),
	})
}

type AdminPost struct {
	Id  string `json:"id"`
	Pwd string `json:"pwd"`
}

func AdminAllowNode(c *gin.Context) {
	var posts AdminPost

	if err := c.ShouldBindJSON(&posts); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if posts.Pwd != GetPwd() {
		c.String(http.StatusUnauthorized, "invalid password")
		return
	}

	id := posts.Id

	if _, erra := strconv.Atoi(id); erra != nil {
		c.String(http.StatusBadRequest, "invalid id (type not int)")
		return
	}

	data := jsons.ReadAllJSON()
	datas := data.Datas

	i := 0

	flag := false

	for i = 0; i < len(datas); i++ {
		if strconv.Itoa(datas[i].Id) == id {
			flag = true
			if datas[i].Active {
				c.String(http.StatusOK, "invalid id (activated)")
				return
			} else {
				break
			}
		}
	}

	if !flag {
		c.String(http.StatusOK, "invalid id (not exist)")
		return
	} else {
		SendToMail(
			data.Datas[i].Email,
			"您的元ICP备"+strconv.Itoa(data.Datas[i].Id)+"号已通过",
			"<p>Dear "+data.Datas[i].SiteUser+",</p><p>您的元ICP备"+strconv.Itoa(data.Datas[i].Id)+"号已通过，祝贺您正式加入元ICP！</p><p>此致</p><p>敬礼</p><br><p>MetaICP Admin</p><p>"+time.Now().Format("2006-01-02 15:04:05")+"</p>",
		)
		data.Datas[i].Active = true
		jsons.SetAllJson(data)
		c.String(http.StatusOK, "allowed")
	}
}

func AdminBanNode(c *gin.Context) {
	var posts AdminPost

	if err := c.ShouldBindJSON(&posts); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if posts.Pwd != GetPwd() {
		c.String(http.StatusUnauthorized, "invalid password")
		return
	}

	id := posts.Id

	if _, erra := strconv.Atoi(id); erra != nil {
		c.String(http.StatusBadRequest, "invalid id (type not int)")
		return
	}

	data := jsons.ReadAllJSON()
	datas := data.Datas

	i := 0

	flag := false

	for i = 0; i < len(datas); i++ {
		if strconv.Itoa(datas[i].Id) == id {
			flag = true
			if datas[i].Active {
				c.String(http.StatusOK, "invalid id (activated)")
				return
			} else {
				break
			}
		}
	}

	if !flag {
		c.String(http.StatusOK, "invalid id (not exist)")
		return
	} else {
		SendToMail(
			data.Datas[i].Email,
			"您的元ICP备"+strconv.Itoa(data.Datas[i].Id)+"号已被驳回",
			"<p>Dear "+data.Datas[i].SiteUser+",</p><p>您的元ICP备"+strconv.Itoa(data.Datas[i].Id)+"号已被驳回，可能是因为您不满足以下的要求：</p><ul><li>不可访问（国内访问被屏蔽、多地DNS污染等）</li><li>站点违规，或含政治敏感内容</li><li>站点未按要求在页脚悬挂元备标识</li></ul><p>站点整改完毕后方可再次申请。</p><p>此致</p><p>敬礼</p><br><p>MetaICP Admin</p><p>"+time.Now().Format("2006-01-02 15:04:05")+"</p>",
		)
		data.Datas = append(data.Datas[:i], data.Datas[i+1:]...)
		jsons.SetAllJson(data)
		c.String(http.StatusOK, "banned")
	}
}

func SelectNode(c *gin.Context) {
	id := c.Param("id")
	datas := jsons.ReadAllJSON().Datas
	var i int
	flag := -1
	for i = 0; i < len(datas); i++ {
		if strconv.Itoa(datas[i].Id) == id {
			if datas[i].Active {
				flag = 1
			} else {
				flag = 0
			}
			break
		}
	}

	if flag == -1 {
		c.HTML(http.StatusOK, "select.html", gin.H{
			"title": "查询",
			"flag":  flag,
		})
		return
	}

	c.HTML(http.StatusOK, "select.html", gin.H{
		"title":      "查询",
		"flag":       flag,
		"sitename":   datas[i].SiteName,
		"siteuser":   datas[i].SiteUser,
		"domain":     datas[i].Site,
		"id":         id,
		"about":      datas[i].About,
		"maindomain": GetDomain(),
	})
}

func JoinNode(c *gin.Context) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	c.HTML(http.StatusOK, "join.html", gin.H{
		"title":     "加入",
		"randomnum": strconv.Itoa(time.Now().Year()) + fmt.Sprintf("%04d", r.Intn(10000)),
	})
}

func VerifyNode(c *gin.Context) {
	ids := c.PostForm("id")
	email := c.PostForm("email")
	name := c.PostForm("name")
	domain := c.PostForm("domain")
	siteuser := c.PostForm("siteuser")
	about := c.PostForm("about")
	messages := "请至邮箱 " + email + " 中查看验证码！"
	flag := 1
	if ids == "" || name == "" || domain == "" || siteuser == "" || about == "" || email == "" {
		messages = "参数不完整！请将所有信息填充完整！"
		flag = 0
	}

	id, erri := strconv.Atoi(ids)

	if erri != nil || (id/10000) != time.Now().Year() || id > 99999999 || id < 10000000 {
		messages = "备案号无效，请重新检查备案号！"
		flag = 0
	}

	emailreg := regexp.MustCompile(`(^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\.[a-zA-Z0-9-.]+$)`)
	emailmatch := emailreg.MatchString(email)

	if !emailmatch {
		messages = "邮箱无效，请重新检查！"
		flag = 0
	}

	alldatas := jsons.ReadAllJSON()
	verdatas := alldatas.Verify
	for i := 0; i < len(verdatas); i++ {
		if verdatas[i].Email == email {
			messages = "此邮箱正在等待验证，请更换邮箱！"
			flag = 0
			break
		}
	}

	datas := alldatas.Datas
	for i := 0; i < len(datas); i++ {
		if datas[i].Email == email && !datas[i].Active {
			messages = "在上一个备案通过前，此邮箱不得为第二网站的登记邮箱！"
			flag = 0
			break
		}
		if datas[i].Site == domain {
			messages = "此网站已有记录，请勿重复备案！"
			flag = 0
			break
		}
		if datas[i].Id == id {
			messages = "此号码已被他人注册，换一个号试试哦！"
			flag = 0
			break
		}
	}

	if flag == 1 {
		vercode := RandVerifyCode()
		SendToMail(
			email,
			"您的元ICP备验证码",
			"<p>Dear,</p><h2 style=\"text-align: center;\">验证码："+vercode+"</h2>",
		)
		datas := jsons.ReadAllJSON()
		datas.Verify = append(datas.Verify, jsons.VerifyJSON{
			Id:         id,
			Email:      email,
			SiteName:   name,
			SiteUser:   siteuser,
			Site:       domain,
			About:      about,
			VerifyCode: vercode,
		})
		jsons.SetAllJson(datas)
	}

	c.HTML(http.StatusOK, "verify.html", gin.H{
		"title":   "加入",
		"message": messages,
		"isok":    flag,
		"email":   email,
	})
}

func SubmitNode(c *gin.Context) {
	verify := c.PostForm("verify")
	email := c.PostForm("email")
	alldatas := jsons.ReadAllJSON()
	verdatas := alldatas.Verify
	p := 0
	flags := true
	for p = 0; p < len(verdatas); p++ {
		flags = false
		if verdatas[p].Email == email {
			break
		}
	}
	messages := ""
	if !flags {
		messages = "完成！详情请跳转到：" + GetDomain() + "/select/" + strconv.Itoa(verdatas[p].Id)
	}
	if verify == "" || email == "" {
		messages = "参数不完整，请重新检查！"
	}
	flag := ""
	var i int
	for i = 0; i < len(verdatas); i++ {
		if verdatas[i].Email == email {
			flag = verdatas[i].VerifyCode
			break
		}
	}
	if flag == "" {
		messages = "邮箱无效/没有记录！"
	} else {
		if flag != verify {
			messages = "验证码错误！"
		} else {
			datas := jsons.ReadAllJSON().Datas
			flag2 := false
			for k := 0; k < len(datas); k++ {
				if datas[k].Id == datas[i].Id {
					messages = "该号码已被他人先注册成功了，请换一个号码重新验证！"
					flag2 = true
					break
				}
			}
			if !flag2 {
				alldatas.Datas = append(alldatas.Datas, jsons.DatasJSON{
					Id:       verdatas[i].Id,
					Email:    verdatas[i].Email,
					SiteName: verdatas[i].SiteName,
					SiteUser: verdatas[i].SiteUser,
					Site:     verdatas[i].Site,
					About:    verdatas[i].About,
					Active:   false,
				})
			}
			alldatas.Verify = append(alldatas.Verify[:i], alldatas.Verify[i+1:]...)
			jsons.SetAllJson(alldatas)
		}
	}
	c.HTML(http.StatusOK, "verify.html", gin.H{
		"title":   "完成",
		"message": messages,
	})
}
