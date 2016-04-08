FROM centos:centos7
MAINTAINER huangyang "503582241@qq.com"


#RUN echo "deb http://archive.ubuntu.com/ubuntu precise main universe"> /etc/apt/sources.list
#RUN apt-get update
#RUN apt-get install -y openssh-server
#RUN mkdir -p /var/run/sshd

#RUN echo "root:123456" | chpasswd 

# 容器需要开放SSH 22端口
#EXPOSE 22

# 容器需要开放Tomcat 8080端口
EXPOSE 4001

ENTRYPOINT ["./macblog.exe"]
