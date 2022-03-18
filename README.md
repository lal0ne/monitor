# monitor

监控网站目录下的文件变更，可通过钉钉机器人发送告警信息。

## 配置文件说明

```yaml
# 记录目录下所有文件变更情况
logflag: false
log: monitor.log
# 钉钉配置文件
dingtalk:
  flag: true
  api: https://oapi.dingtalk.com/robot/send?access_token=XXXXXX
  # 自定义关键词
  key: XXX
# 监控配置
monitor:
  -
    # 监控目录
    path: /tmp/
    # 以下目录变更不做处理
    ignoreflag: false
    ignorepath:
      - /tmp/log/
    # 以下后缀名为黑名单，即该类型均告警[优先级没有ignorepath高]
    typeflag: true
    type:
      - ^php([0-9]*)
  -
    path: /root/
    ignoreflag: true
    ignorepath:
      - /root/log/
    typeflag: true
    type:
      - ^php([0-9]*)
```

## 运行

将可执行程序与配置文件放置同一文件夹，运行可执行程序即可。

## 下载

[下载地址](https://github.com/lal0ne/monitor/releases)

本人提供`linux`和`win`下的可执行程序。

## todo

- [x] 多个目录
- [x] 黑名单正则