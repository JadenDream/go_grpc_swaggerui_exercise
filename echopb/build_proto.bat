
set THEPATH=D:\Eclipse_Workspace\go_grpc_swaggerui_exercise\src\grpc-gateway-example\echopb
set GOOGLEAPIS=%GOPATH%\src\github.com\googleapis\googleapis
set DESCRIPTORPROTO=%PROTOBUFPATH%\google\protobuf\descriptor.proto

REM ½sÄ¶

protoc -I %THEPATH% -I %GOOGLEAPIS% -I %PROTOBUFPATH%\include --go_out=plugins=grpc:. %THEPATH%\service.proto

protoc -I %THEPATH% -I %GOOGLEAPIS% -I %PROTOBUFPATH%\include --grpc-gateway_out=logtostderr=true:. %THEPATH%\service.proto

protoc -I %THEPATH% -I %GOOGLEAPIS% -I %PROTOBUFPATH%\include --swagger_out=logtostderr=true:. %THEPATH%\service.proto

pause
