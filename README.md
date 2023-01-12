관심종목 공유
- finance-datareader , pykrx 두개의 module을 사용해서 전체 종목에 대한 데이터를 interval로 업데이트 합니다.

**swagger 문서**
* chart_scheduler = http://localhost:3000/docs/
* gin 문서 = http://localhost:8090/docs/index.html

**config**
* touch .env
* ENV=development
* MYSQL_ROOT_PASSWORD=password
* MYSQL_DATABASE=DATABASE
* JWT_ACCESS_SECRET=access_secret
* JWT_REFRESH_SECRET=refresh_secret

* cd mysql
* mkdir data -> 데이터를 실제로 담아둘 dir 생성

**data insert**
* http://localhost:8090/docs/index.htm -> 유저 생성
* http://localhost:3000/docs 이동
* 유저 생성했을 떄 떨어지는 access_token으로 종목 -> 차트 insert
