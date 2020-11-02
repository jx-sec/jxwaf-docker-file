FROM centos:7

#USER root
RUN cd /tmp  \
    && yum install -y epel-release pcre-devel openssl-devel gcc cmake make  lua-devel  automak git \
    && git clone https://github.com/jx-sec/jxwaf.git  \
    && cd jxwaf \
    && tar zxvf openresty-1.15.8.3.tar.gz \
    && cd openresty-1.15.8.3 \
    && ./configure --prefix=/opt/jxwaf --with-http_v2_module --with-http_stub_status_module && gmake && gmake install \
    && mv /opt/jxwaf/nginx/conf/nginx.conf  /opt/jxwaf/nginx/conf/nginx.conf.bak \
    && cp ../conf/nginx.conf /opt/jxwaf/nginx/conf/ \
    && cp ../conf/full_chain.pem /opt/jxwaf/nginx/conf/ \
    && cp ../conf/private.key /opt/jxwaf/nginx/conf/ \
    && mkdir /opt/jxwaf/nginx/conf/jxwaf \
    && cp ../conf/jxwaf_config.json /opt/jxwaf/nginx/conf/jxwaf/ \
    && cp ../conf/GeoLite2-Country.mmdb /opt/jxwaf/nginx/conf/jxwaf/ \
    && cp -r ../lib/resty/jxwaf  /opt/jxwaf/lualib/resty/ \
    && /opt/jxwaf/nginx/sbin/nginx -t


CMD sed -i 's/"waf_api_key": ""/"waf_api_key": "'$JXWAF_API_KEY'"/'  /opt/jxwaf/nginx/conf/jxwaf/jxwaf_config.json \
    && sed -i 's/"waf_api_password": ""/"waf_api_password": "'$JXWAF_API_PASSWD'"/'    /opt/jxwaf/nginx/conf/jxwaf/jxwaf_config.json  \
    #&& sed -i 's#"waf_update_website": "http://update2.jxwaf.com/waf_update"#"waf_update_website": "'$WAF_UPDATE_WEBSITE'"#'    /opt/jxwaf/nginx/conf/jxwaf/jxwaf_config.json  \
    && sed -i 's#http://update2.jxwaf.com#$WAF_UPDATE_WEBSITE#g'    /opt/jxwaf/nginx/conf/jxwaf/jxwaf_config.json \
    && server_name=`ip addr | grep inet | awk '{ print $2; }' | sed 's/\/.*$//'|grep -v 127.0.0.1|head -1`\
    && server_mac=`hostname` \
    && sed -i "s/server_info_detail/$server_name|$server_mac/g" /opt/jxwaf/nginx/conf/jxwaf/jxwaf_config.json \
    && cat /opt/jxwaf/nginx/conf/jxwaf/jxwaf_config.json \
    && /opt/jxwaf/nginx/sbin/nginx \
    && tail -f /opt/jxwaf/nginx/logs/error.log