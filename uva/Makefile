CPP_FILES := $(wildcard *.cpp)
OBJ_FILES := $(CPP_FILES:.cpp=.o)

%.o: %.cpp
	g++ -lm -lcrypt -O2 -std=c++11 -pipe -o $@ $<
