# Fabric-SDK-Traceability

[![Build Status](https://travis-ci.org/chenwbyx/Fabric-SDK-Traceability.svg?branch=master)](https://travis-ci.org/chenwbyx/Fabric-SDK-Traceability)  ![](https://img.shields.io/badge/language-go-blue.svg)


> 一个基于 [Fabric-SDK-Go](https://github.com/hyperledger/fabric-sdk-go) 的商品溯源防伪查询系统。
> 在此之前请确保已安装 [Hyperledger Fabric](https://github.com/hyperledger/fabric)（本项目基于Hyperledger Fabric v1.1)

------

### 所需环境
    ```
    go 1.10
    docker 18.09.2
    docker-compose 1.12.0
    ```

### 安装及配置
* Step 1：
   * 创建并进入目录 $GOPATH/src/github.com/kongyixueyuan.com
   * 拉项目```git clone https://github.com/chenwbyx/Fabric-SDK-Traceability ```
   * 修改文件夹名字： ```mv Fabric-SDK-Traceability education```
   * 进入文件夹：```cd education```
* Step 2：
   * 敲命令：```make```
   * 看到如下输出表示运行成功
     ![img](https://github.com/chenwbyx/Fabric-SDK-Traceability/blob/master/img/update_findEduByCertNoAndName.png)
* Step 3：
   * 访问```http://localhost:8000```
     ![img](https://github.com/chenwbyx/Fabric-SDK-Traceability/blob/master/img/html_index.png)


------
作者 [@chenwb](https://github.com/chenwbyx/)  
2019 年 3 月 30 日