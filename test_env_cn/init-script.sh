#!/bin/bash

while ! curl -s jxwafserver.demo:8000 > /dev/null; do
    sleep 1
done

echo '{"user_name":"test", "user_password":"123456", "waf_auth":"ee747988-612b-4790-b8ea-fb49c04fc1ea", "log_conf_remote":"true", "log_ip":"jxlog", "log_port":"8877", "log_response":"true", "log_all":"true", "report_conf":"true", "report_conf_ch_host":"clickhouse", "report_conf_ch_port":"9000", "report_conf_ch_user":"jxlog", "report_conf_ch_password":"jxlog", "report_conf_ch_database":"jxwaf"}' > /tmp/init_payload.json

curl -X POST "http://jxwafserver.demo:8000/demo_env_init" -H "Content-Type: application/json" -d @/tmp/init_payload.json

/opt/run