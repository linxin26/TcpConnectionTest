server/server.go 为服务端程序

client/client.go为客户端程序

build.sh用于编译client并生成Docker镜像，该镜像中包含了编译client所生成的客户端程序
docker-compose.yaml为docker compose的配置文件

##服务端启动
./server -bindAddr=0.0.0.0:28009

##客户端
./build.sh
docker-compose up -d --scale demo-client=2

