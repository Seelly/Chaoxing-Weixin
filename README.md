# Chaoxing-Weixin
学习通未完成作业推送到微信
本项目爬取学习通未完成作业，并利用微信公众号自动推送能力实现通知。
# 第一步：注册微信公众号测试号
1、利用下面的链接注册一个微信公众测试号，微信公众平台测试号地址：
> https://mp.weixin.qq.com/debug/cgi-bin/sandbox?t=sandbox/login

2、打开并注册一个微信公众号测试号，扫码登录即可。

3、扫码登录成功后，就可以生成微信公众号测试号的appID和appsecret这两串信息

# 第二步：扫描测试号二维码
1、向下滑动，找到测试号二维码，使用微信扫描测试号二维码并关注。

2、用户列表会自动显示用户信息，并可以看到openid。
# 第三步：新增消息模板
1、点击新增测试模板，填入以下参考信息
```
今天还有 
{{course.DATA}} 
{{ work.DATA}} 
共 {{num.DATA}} 项作业未完成哦~
```
模板可以修改，但是代码中也要进行修改
2、复制模板消息的id
# 第四步：启动项目
1、将所需内容填入代码中
2、编译项目(开发环境为windows)
``` shell
要执行脚本的电脑为linux执行以下命令：
set GOOS=linux
set GOARCH=amd64
go build main.go
要在windows下运行：
go build main.go
```
3、将编译好的脚本文件添加到定时任务就完成配置了。


--------------
[学习通签到小程序-项目地址](https://github.com/Wenhorm/chaoxingxxt_sign/blob/main/README.md)
![](https://www.wenhorm.top/media/editor/mmexport1655269923965_20220615131619976981.jpg)
