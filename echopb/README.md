# echopb proto windown環境編譯說明
環境: windown

## 編譯前置
1.先編譯 scripts 內的includetxt.go, 可產出scripts.exe

2.再將scripts.exe取到echopb資料夾內

## 編譯流程說明
1.執行 build_proto.bat 進行編譯 (產出pb.go , pb.gw.go , service.swagger.json)

2.執行scripts.exe 會產出(更新) swagger.pb.go

