1. start etcd

etcd --listen-client-urls http://0.0.0.0:2379 ----advertise-client-urls http://0.0.0.0:2379

docker run -p 2379:2379 -p 2380:2380 --name etcd quay.io/coreos/etcd:latest /usr/local/bin/etcd

// 貌似2.x不通
////docker run -d -p 2379:2379 -p 2380:2380  -v /tmp:/data  --name etcd elcolio/etcd:latest


**使用这个镜像**

docker run -d -p 2379:2379 -p 2380:2380 --name etcd appcelerator/etcd

2. start provider

**使用下面的命令启动provider, 注意network设置和hostname设置**

docker run -d --network=host --hostname=172.17.0.1 --name provider-small registry.cn-hangzhou.aliyuncs.com/aliware2018/agent-demo provider-small

provider-small


docker run -d --network=host --hostname=172.17.0.1 --name provider-medium registry.cn-hangzhou.aliyuncs.com/aliware2018/agent-demo provider-medium

provider-medium


docker run -d --network=host --hostname=172.17.0.1 --name provider-large registry.cn-hangzhou.aliyuncs.com/aliware2018/agent-demo provider-large

provider-large


**如果不设置hostname,java 程序会报 主机找不到 的错误**

--hosname=172.17.0.1


3. start consumer

**使用下面的命令启动 consumer**

docker run -d --network=host --hostname=172.17.0.1 --name consumer registry.cn-hangzhou.aliyuncs.com/aliware2018/agent-demo consumer


4. 调试

docker run --network=host -ti --rm ellerbrock/alpine-bash-curl-ssl /bin/bash 

ip addr show docker0 | grep 'inet\b' | awk '{print $2}' | cut -d '/' -f 1


5. curl

**查看ip地址、执行curl命令等**

在 ellerbrock/alpine-bash-curl-ssl 容器中执行



- 往 provider-small agent 发消息
curl -d "interface=com.alibaba.dubbo.performance.demo.provider.IHelloService&method=hash&parameterTypesString=Ljava/lang/String;&parameter=helloword" -X POST http://172.17.0.1:30000


- consumer agent
curl -d "interface=com.alibaba.dubbo.performance.demo.provider.IHelloService&method=hash&parameterTypesString=Ljava/lang/String;&parameter=helloword" -X POST http://172.17.0.1:20000


- consumer
curl http://172.17.0.1:8087/invoke

