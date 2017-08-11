# echopb proto windown環境編譯說明
環境: windown

## 編譯前置
1.先編譯 scripts 內的includetxt.go, 可產出scripts.exe

2.再將scripts.exe取到echopb資料夾內

## 編譯流程說明
1.執行 build_proto.bat 進行編譯 (產出pb.go , pb.gw.go , service.swagger.json)

2.執行scripts.exe 會產出(更新) swagger.pb.go


## 額外說明
編譯產生的 service.swagger.json 為service.proto的api文件

scripts.exe則是要將該api文件轉成 swagger.pb.go

而swagger.pb.go的作用則為在編譯服務時產出以下連結的對應json文件

https://localhost:10000/swagger.json



