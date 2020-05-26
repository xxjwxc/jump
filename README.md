# rmon
remote monitoring . 远程监控工具

# 支持
- 默认配置 (jump config),保存公共配置
- 登录(jump in),远程登录(支持tab键)
- 监控(jump run),远程监控(支持登录命令行执行，支持文件监控,tail -f)
- ptrace 远程进程调试


## 设置配置文件(jump config)

```
./jump config -d=~ -i=127.0.0.1 -p=123456 -u=ubuntu -P=22 -c="cd /var/log/,ls,ll"
```

### 更多：
```
./jump config -h
```

## 远程登录(jump in)

```
./jump in
```

## 远程登录(jump in)

```
./jump in
```

## 远程监控(jump run)

```
./jump run -f=nginx/access.log
```

### 更多：
```
./jump run -h
```