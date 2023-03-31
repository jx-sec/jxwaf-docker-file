## jxwaf docker 部署
### jxwaf mini server实例
#### 快速启动
```
wget https://raw.githubusercontent.com/jx-sec/jxwaf-docker-file/master/jxwaf-mini-server/docker-compose.yml
docker compose  up -d 
```
#### mini server + jxlog 快速启动
```
wget https://raw.githubusercontent.com/jx-sec/jxwaf-docker-file/master/docker-compose.yml
docker compose  up -d 
```
#### 参数说明
|  参数  | 说明 |
|  ----  | ----  |
| DBUSER | 数据库用户名，默认：root  |
|DBNAME  | 数据库库名，默认：jxwaf |
|DBPASSWORD| 数据库用户名密码，默认：jxwaf|
|DBHOST | 数据库HOST,默认：mariadb|
|DBPORT | 数据库端口：3306 | 

### jxwaf 节点实例
#### 获取jxwaf接入凭证
按照上一步“jxwaf管理中心实例”快速启动后，假设管理中心 IP 为 10.0.0.1,则打开网址 http://10.0.0.1:8088 进行注册,注册完后登录账号,在 系统管理 -> 基础配置 页面获取"API_KEY"和"API_PASSWORD"
#### 快速启动
```
docker run  -d  --rm    -p 80:80  -p 443:443 -e JXWAF_SERVER="http://10.0.8.1:8088" -e WAF_API_KEY=19a0953e-1b64-42c3-a4cb-a2c5d4b88285 -e WAF_API_PASSWORD=595e8383-ec7e-4f31-9b98-b55964028763  jxwaf/jxwaf:latest
```
#### 自定义端口启动
```
docker run  -d  --rm  -p 8443:8443 -e  -e HTTPs_PORT="8443" JXWAF_SERVER="http://10.0.8.1:8088" -e WAF_API_KEY=19a0953e-1b64-42c3-a4cb-a2c5d4b88285 -e WAF_API_PASSWORD=595e8383-ec7e-4f31-9b98-b55964028763  jxwaf/jxwaf:latest
```
#### 参数说明
|  参数  | 说明 |
|  ----  | ----  |
|JXWAF_SERVER|管理中心网址，示例：http://10.0.0.1:8088|
|WAF_API_KEY|管理中心API_KEY，示例：19a0953e-1b64-42c3-a4cb-a2c5d4b88285|
|WAF_API_PASSWORD|管理中心API_PASSWORD，示例：595e8383-ec7e-4f31-9b98-b55964028763|
|  HTTP_PORT | 自定义http端口，默认：80 |
|  HTTPs_PORT | 自定义https端口，默认：443 |
