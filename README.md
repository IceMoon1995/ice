# ice
像冰一般的~ 

ps：目前暂不开源,以后可能会传上去

弱鸡的扫描器探测探针，主要用于内网探测端口扫描相关信息，模拟mysql,vnc,rmi等协议。

害人之心不可有，防人之心不可无！

当内网扫描器进行扫描时，记录扫描日志，以供分析。

提供mysql反制（目前只支持windwos），rmi命令执行反制

## 使用方式

ice.exe 

默认监听tcp 21,22,23,53,80-90,1080,1099,1433,2375,3306,3389,5900,6379,7001端口，要防火墙放行

当远程扫描器，扫描以上监听端口时，会记录连接日志信息（默认iceLog.txt）以及发送报文信息， -l 2开启全部报文记录（默认只记录连接请求，dns记录以及反制日志）

开启53端口时，会默认启用dns转发，需配置本地dns到当前地址（如：127.0.0.1），会记录dns解析日志并转发到目标真实dns服务器（默认114.114.114.114:53），-ds指定目标dns服务器

需要配置本机dns服务器，对访问网站无任何影响
![Image text](https://github.com/IceMoon1995/ice/blob/master/img/dns%E8%A7%A3%E6%9E%90%E8%AE%B0%E5%BD%95.png)

当-p内包含3306，1099，5900端口时，默认开启协议模拟，可被指纹探测与识别，且默认开启3306（目前只写了windwos文件拖取逻辑），rmi蜜罐反制，目前只实现了jdk8的cc5,cc6链 -c指定执行系统命令！

nmap指纹识别结果
![Image text](https://github.com/IceMoon1995/ice/blob/master/img/nmap%E6%8E%A2%E6%B5%8B%E6%8C%87%E7%BA%B9.png)

mysql反制读取文件
![Image text](https://github.com/IceMoon1995/ice/blob/master/img/mysql%E5%8F%8D%E5%88%B6.jpg)

mysql蜜罐支持yml配置，主要可以修改file_list以配置需要读取文件信息
![Image text](https://github.com/IceMoon1995/ice/blob/master/img/mconfig%E9%85%8D%E7%BD%AE%E4%BF%A1%E6%81%AF.jpg)

## 参数说明
```
-p string
      指定监听的端口，如：21,22,23,53,80-90,1080,1099,1433,2375,3306,3389,5900,6379,7001
-o string
      生成的日志名称
-nlog bool
      不写日志
-c string
rmi等蜜罐反制成功后执行的命令，默认calc.exe，-c 指定其他系统命令
-l int
      日志记录等级，1、记录请求基本信息（源ip，目的ip）；2、记录请求详细信息（包含报文）
-ds string
      设置上级dns转发地址，默认114.114.114.114:53，格式 -ds 8.8.8.8:53
```
## TODO

[+] 协议流量以及恶意流量识别？感觉很难做，需要全流量保存分析？

## 新增功能

[+] dns转发，用于本机dns记录，默认监听本机0.0.0.0:53端口，远程dns服务器默认为114.114.114.114，-p内包含53时自带开启

[+] 新增mysql,vnc,rmi协议模拟，可被扫描器识别，默认开放在3306，5900，1099端口，当端口在-p内时，默认模拟，目前不支持修改

[+] 新增mysql反制蜜罐（目前仅针对windiws），新增rmi反制蜜罐（目前实现cc5,cc6链，需要运行时-c指定执行命令，默认calc.exe）
