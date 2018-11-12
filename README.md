# **cdn upload admin**

> cdn upload admin api server

## 一、Build setup

### 1. 请确保已经安装好go和设置好$GOPATH、$GOROOT环境变量
### 2. 请确保已经安装好了mysql 
```
导入 ../sql/cdn-upload.sql 创建库和表

设置管理员账号密码， 密码是md5加密后的
use cdnupload;
INSERT INTO t_user (`is_admin`, `username`, `password`) VALUES (0, 'admin', '21232f297a57a5a743894a0e4a801fc3');

授权用户能访问 cdnupload库
GRANT all ON cdnupload.* to ${dbuser}@"${dbip}" Identified by "${dbpwd}" WITH GRANT OPTION;
flush privileges;
```
### 3. 编辑cfg.json更改root_path和db.addr的值为实际情况
```
{
    "debug": true,
    "http": {
        "listen": "0.0.0.0:9527",
        "root_path": "/data/resource/cdn_file/"
    },
    "db": {
        "addr": "${dbuser}:${dbpwd}@tcp(${dbip}:${dbport｝)/cdnupload?charset=utf8mb4",
        "idle": 10,
        "max": 100
    }
}
```
### 4. 安装依赖
```  go get```
### 5. 编译
```chmod u+x control && ./control build``` 
### 6. 启动
``` ./control start```
### 7. 查看日志
``` ./control tail```

## 二、设置 NGINX 反向代理

### 请确保已经安装好 NGINX，下面是相关的 nginx server配置
```
server {
    listen  80;
    server_name ${domain}; 
    access_log logs/cdnupload_access.log main;
    
    # 设置 header是在开发环境下前后端分离时支持跨域，在线上环境时可不用配置
    add_header 'Access-Control-Allow-Origin' '*';
    add_header 'Access-Control-Allow-Credentials' 'true';
    add_header 'Access-Control-Allow-Methods' 'GET, POST, OPTIONS';
    add_header 'Access-Control-Allow-Headers' '*';
    add_header 'Access-Control-Expose-Headers' '*';
    
    # 本地开发时vue在请求POST方法前会请求一次OPTIONS请求，go后端没有处理，直接在nginx这里处理返回200状态
    # vue webpack打包后的文件部署到线上服时可不用配置
    if ( $request_method = 'OPTIONS' ) {
        return 200;
    }
    # 前端静态资源文件目录，在vue-router中设置 model为history时Nginx需要以下配置try_files
    location / {
        root   /data/cdn-upload-admin/cdnadmin_webserver;
        index index.html;
        try_files $uri $uri/ /;
    }
    # api请求反向代理至后端， 定义最大上传文件大小为 1000M
    location /api {
        client_max_body_size 1000M;
        proxy_pass http://127.0.0.1:9527;
        proxy_redirect     off;
        proxy_set_header   Host             $host;
        proxy_set_header   X-Real-IP        $remote_addr;
        proxy_set_header   X-Forwarded-For  $proxy_add_x_forwarded_for;
    }
}
```
