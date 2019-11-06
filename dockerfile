FROM scratch
#FROM alpine:latest

#VOLUME /data

#工作目录
WORKDIR /data

#游戏端口
EXPOSE 12000

CMD ["./bbs_server"]
