#include<iostream>
#include"socket.h"
#include<random>
#include<cstdlib>

int main(int argc ,char **argv){
		if( argc <= 2 ) 
		{	std::cout<<"Usage "<<argv[0]<<' '<<"host"
		<<' '<<"port"<<" to connect db server"<<std::endl;
			exit(-argc);
		}
		std::random_device r;
		std::default_random_engine rd(r());
		std::uniform_int_distribution<int> id(1000000,9999999);
		std::string userid = std::to_string(id(rd));
		 
			std::string host(*(argv+1));
			Socket s(host,atoi(argv[2]));
			s.Write(userid);
			//接下来就可以发弹幕了
				while (true){
						std::string buf;
						buf = "卢本伟牛逼（破音）！";
						//while(std::cin >> buf){
							int ln = s.Write(buf);
							if (ln != 0) continue;
						//}
				}
}

