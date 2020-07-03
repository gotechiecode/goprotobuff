# to check the different data format sizes run the below commands in terminal at the folder location
stat -f%z customer.xml
stat -f%z customer.json
stat -f%z customer.proto

#To Generate .go files from .proto files
# definition of below command is 'protoc' is to compile the .proto files, '--go_out' is a flag to generate files, 'go_out=.' will generate files in current directory, "*.proto" is to compile all .proto files in current directory
protoc --go_out=. *.proto 