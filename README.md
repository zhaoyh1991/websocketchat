# websocketchat
golang websocket
使用golang的websocket的一个包做的可以互相通信的小项目
#使用方法
启动项目之后
打开两个浏览器窗口f12之后在两个console中
#第一个console中
var ws1 = new WebSocket('ws://localhost:8081/ws?from=1');
ws1.onmessage=function(evt){alert(evt.data)}
#第二个console中
var ws2 = new WebSocket('ws://localhost:8081/ws?from=2');
ws2.onmessage=function(evt){alert(evt.data)}
#然后再第一个console中
ws1.send('{"from":"1","to":"2","content":"hello"}')
然后在第二个console中就会弹出消息了
也可以简单的用来做推送相关的功能
