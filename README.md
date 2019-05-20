## 编译jxwaf docker镜像

```docker  build    -t  jxwaf . ``` 

## 获取jxwaf接入凭证

访问 [http://www.jxwaf.com](http://www.jxwaf.com) 并注册账号,在全局配置页面获取"api key"和"api password"

## jxwaf实例
简单的jxwaf实例：
```
docker run -p80:80 --env JXWAF_API_KEY="<you api key>" --env JXWAF_API_PASSWD="<you api password>"  jxwaf
```
jxwaf最佳实践：
```
 docker run  --net=host -p80:80 --env JXWAF_API_KEY="<you api key>"   --env JXWAF_API_PASSWD="<you api password>"    --restart=always jxwaf 
```