# MetaICP Community Edition

[***zh-CN version***](README_zh.md)

*Get your very own virtual registration site quickly!*

## Use

### Linux

Find the latest version for your system on the Release page, download and run the following command:

```bash
apt install screen

screen -R icpservice
```

You should then see a New command line window with "New screen..." flashing on the screen. , run the following command:

```bash
cd path/to/release  # your release folder 

tar -zxvf metaicp.tar.gz  # unzip

chmod +x metaicp  # give power to run

vim ./pwd  # change password

vim ./notice  # change notice

vim ./port  # change run port

vim ./email  # change email info

vim ./domain  # change domain of the main site

./metaicp  # run service

```

Press `Ctrl+A` followed by `D` to exit the command line.

***Done.***


### Windows

Download the Windows Release, unzip it, and run 'metaicp.exe' from it.

***Done.***

### macOS X

Download the macOS X Release, unzip it, and run it in bash.

***Done.***

---

Your server will be running on the port in the `port` file (default 8080). Feel free to reverse proxy using 'nginx', 'Apache', etc.

## Discuss

This version is maintained by the community, so join and discuss!

[QQ Group](https://qm.qq.com/q/i34THscWk0)

![QQ QRcode](bin/qq.jpg)

## License

MIT license

**In the premise of not using this work for commercial purposes, be sure to retain the signature of this work!**
