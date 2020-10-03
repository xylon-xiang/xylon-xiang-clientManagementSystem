# Daily Report

由于夜间才停止打代码写日报，可能7天里,日报会在凌晨左右写前一天的。为避免时间差错，采用日本的28小时制

## 2020.10.2

2020.10.2 12 ： 05分记

今日内容:

* B/S架构与C/S架构
* webSocket入门，其相关概念和运行方式
* 项目架构设计（刚起步）
* 部分module的设计

## 2020.10.3

2020.10.3 12:23记

今日内容：

* 重构项目架构
* 确定教师机与学生机之间的通信使用Http和webSocket混合沟通. \
    ***webSocket Message format: "{verb}:{body}"***
    * 教师机：架Http Server， 通过Http接受信息，通过webSocket发送信息
    * 学生机：通过Http发送信息，通过webSocket接受信息
* 除去Module、DB operations, 包按照功能划分，而非层次划分
