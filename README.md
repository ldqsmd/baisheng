# 
一般的 beego 项目的目录如下所示：
```
├── conf
│   └── app.conf
├── controllers
│   ├── admin
│   └── default.go
├── main.go
├── models
│   └── models.go
├── static
│   ├── css
│   ├── ico
│   ├── img
│   └── js
└── views
    ├── admin
    └── index.tpl
```
配置文件说明

```
appname     = baisheng          #项目名称
httpport    = 8080              #监听端口    
runmode     = dev

sessionon   = true              #开启session
sessiongcmaxlifetime   = 3600   #session生命周期 /秒


[mysql]
host        = "127.0.0.1"       #mysql地址
databse     = "bai_sheng"       #数据库    
port        = "3306"            #msyql端口
username    = "root"            #数据库账号
password    = "root"            #数据库密码

[email]
host        = "smtp.qq.com"     #smtp host
port        = "587"             #smtp 端口
from        = "MR.Zhang"        #发件人昵称
user        = ""                #发件人邮箱地址
pass        = ""                #发件人邮箱密码
turn        = "on"              #是否开启邮件提醒
 
```
