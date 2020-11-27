## 摸鱼小说应用
某PHP朋友说github找的摸鱼插件不能用了，让我帮他写一个简单能用的摸鱼看小说工具，于是给他整了一个比较简单的小工具


### 使用步骤
* 使用Go开发，无需安装依赖
* Windows下可直接使用提供的可执行文件运行

### 命令介绍
```
moyu E:\123456.txt 100  

// moyu 为可执行文件
// E:\123456.txt 为文本文件路径
// 100 为所要跳转行数(若不输入则载入上次阅读位置)
```
按任意不冲突按键为下一行，"~"，"."，"m" 键为退出并清屏保存阅读记录