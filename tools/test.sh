#----------------基于HTTP----------------#
# 查看状态
curl 127.0.0.1:12345/status
# 插入一个kv
curl -v  127.0.0.1:12345/cache/testkey -XPUT -d testvalue
# 查看key的val
curl 127.0.0.1:12345/cache/testkey
# 查看状态
curl 127.0.0.1:12345/status
# 删除key
curl 127.0.0.1:12345/cache/testkey -XDELETE
# 查看状态
curl 127.0.0.1:12345/status
#----------------基于TCP----------------#
../client/client.exe -c set -k testkey -v testvalue
../client/client.exe -c get -k testkey
curl 127.0.0.1:12345/status
../client/client.exe -c del -k testkey
curl 127.0.0.1:12345/status
