# go-gin
计划开发一个gin框架的go-web项目

### windows系统默认换行符
git config --global core.autocrlf true

### curl命令记录
curl http://localhost:8080/form -Body 'xxxxxxxxx' -Method 'POST'
curl 'http://localhost:8080/posts?id=9876&page=7'  -Body 'username=geektutu&password=1234' -Method 'POST'
curl 'http://localhost:8080/map?ids[jack]=001&ids[john]=002' -Body 'names[a]=Sam&names[b]=David' -Method 'POST'