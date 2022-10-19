## 编译jxwaf docker镜像

```docker  build    -t  jxwaf . ``` 

## 获取jxwaf接入凭证

访问 [http://www.jxwaf.com](http://www.jxwaf.com) 并注册账号,在全局配置页面获取"api key"和"api password"

## jxwaf docker 实例
jxwaf管理中心实例：
```
wget https://raw.githubusercontent.com/jx-sec/jxwaf-docker-file/master/jxwaf-mini-server/docker-compose.yml
docker compose  up -d 
```
jxwaf最佳实践：
```
 docker run  --net=host -p80:80 --env JXWAF_API_KEY="<you api key>"   --env JXWAF_API_PASSWD="<you api password>" --env WAF_UPDATE_WEBSITE=http://update2.jxwaf.com/waf_update   --restart=always jxwaf 
```
