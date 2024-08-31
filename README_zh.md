# MetaICP 社区版

[***英文版本***](README.md)

*快速构建自己的虚拟备案网站！*

## 使用

### Linux

在发布页面找到你的系统的最新版本，下载并运行以下命令:

```bash
apt install screen

screen -R icpservice
```

看到一个新的命令行窗口，屏幕上写着 `New screen…`，接着运行以下命令:

```bash
cd path/to/release  # 下载的 Release 所在的文件夹

tar -zxvf metaicp.tar.gz  # 解压

chmod +x metaicp  # 给予权限以运行

vim ./pwd  # 修改管理员密码

vim ./notice  # 修改公告

vim ./port  # 修改运行端口

vim ./email  # 修改邮箱信息

vim ./domain  # 修改主站域名

./metaicp  # 运行服务

```

按 `Ctrl+A` 并接着按下 `D` 退出命令行。

***完成***


### Windows

下载Windows版本的Release，解压，并双击运行其中的 `metaicp.exe`。

***完成***

### macOS X

下载macOS X发行版，解压，然后在命令行中运行主程序。

***完成***

---

您的服务将运行在 `port` 文件里的端口上（默认为 8080），可以自行使用 `nginx`、`Apache` 等服务进行反向代理。

## 讨论

此版本由社区维护，快加入我们并一起讨论吧！

[QQ群](https://qm.qq.com/q/i34THscWk0)

![QQ二维码](bin/qq.jpg)

## 许可证

MIT license

**在不将本作品用于商业目的的前提下，请务必保留本作品的署名！**
