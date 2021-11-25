# FFxivFate

基于HarmonicaBot的 最终幻想14（FFXIV）Fate通知转发机器人，支持频道消息

#### 使用方法

1. 安装 [go-cqhttp](https://github.com/Mrs4s/go-cqhttp) 并配置正向ws

   ```
     # 正向WS设置
     - ws:
         # 正向WS服务器监听地址
         host: 127.0.0.1
         # 正向WS服务器监听端口
         port: 6700
         middlewares:
           <<: *default # 引用默认中间件
   ```

2. 按格式修改conf文件夹下的配置文档
3. 运行
