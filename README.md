# diction-car

----------------

## grpc와 goroutine을 활용한 프로젝트입니다


1. go run main.go
2. bloomrpc 를 이용해서 proto 타입 형식을 만들어준다
3. 값을 확인 한다


* diction
  * 텍스트 파일에 있는 글자를 찾아서 보내준다
    * Kind : normal or goroutine
    * Word : 찾고 싶은 단어
    * Filename : 찾고 싶은 파일

![Alt text](/diction.png)

* Paralleler
  * 병렬 형식으로 값 찾아오기
  * 그냥 찾아오는 형식과 goroutine을 이용하여 찾아온다
    * 리퀘스트 값을 줄 필요는 없음

![Alt text](/paralleler.png)
