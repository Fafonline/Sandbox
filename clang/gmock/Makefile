all: test

build: 
	clang++  -Isrc/  -lgtest -lgmock test/test.cpp -o ./test/test  -std=c++14 
test: build
	./test/test
