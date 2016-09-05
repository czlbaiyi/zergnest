protoc --go_out=./go *.proto
protoc --js_out=import_style=commonjs,binary:./js *.proto