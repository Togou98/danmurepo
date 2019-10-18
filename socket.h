#include<sys/socket.h>
#include<arpa/inet.h>
#include<sys/types.h>
#include<string>
#include<iostream>
#include<cstring>
#include<unistd.h>

class Socket{
		public:
		Socket();
		Socket(std::string host, int port);
		bool Connected() const {return connected;}
		int Write(std::string& s);
		private:
				int sockfd;
				struct sockaddr_in addr;
				bool connected;
};

Socket::Socket(){
		sockfd = 0;
		connected = false;
}
Socket::Socket(std::string host, int port){
		sockfd = socket(AF_INET , SOCK_STREAM, 0);
		memset(&addr,0,sizeof(addr));
		addr.sin_family = AF_INET;
		addr.sin_addr.s_addr = inet_addr("127.0.0.1");
		addr.sin_port = htons(port);
		if( (connect(sockfd, (const struct sockaddr*)&addr,sizeof(addr))) == 0){
				connected = true;
				std::cout<<"连接弹幕服务器成功"<<std::endl;
		}
}
int Socket::Write(std::string& s){
				if(this->Connected()){
					return write(sockfd,s.c_str(),s.size());
				}else{
						std::cerr<<"你发你ma呢"<<std::endl;
						exit(-100);
				}

}



