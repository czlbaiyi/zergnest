protoc.exe --go_out=./go *.proto
::protoc.exe --cpp_out=./cpp *.proto
::protoc.exe --java_out=./java *.proto
protoc.exe --js_out=import_style=commonjs,binary:./js *.proto