# 52주 최저고가 구하기
현재부터 n기간 동안의 종,시저,고가격 기준으로 최저, 최고가 찾기

## 테이블
project.tb_52_weeks
code_id, price_type, unit_type, unit,  highprice, lowprice 

### 예상 테이블 줄수
종목 * (52주, 1주, 2주, 3주, 1개월,2개월,3개월,4개월,5개월,6개월,7개월,8개월,9개월,10개월,11개월)찾는기간 수 * (저,고,종,시)기준가격    
2000*(12+3)*(4) = 120,000

### 프로세스
1. 테이블 초기화 하기.
   1. 전체 종목코드 조회
   2. 목록 조회 : 마지막날짜 보다 큰 해당종목의 가격+주단위 데이터 조회  
      1. 종목의 가격은 한줄에 저고시종 가가 있음으로 이것을 4개로 분리 시켜 가격종류별 목록(종시저고가 목록 ) 만들기
         1. 종목의 각각의 가격 목록을 일자 기준으로 정렬 시킨다. 일단 보류: sql로 소트 처리하니깐
         2. 종목의 일자가 최신 장열림 날짜 기준으로 몇 주나 몇 개월에 해당하는지 계산후 모델에 저장한다. <== sql로 day_cnt로 처리함.
   3. 목록 탐색하기
      1. 기간에 따라 탐색 범위가 달라짐
         1. 한개의 종목의 가격목록이 4개의 가격종류별 가격목록이 되고 
         2. 4개의 가격종류별 목록이 기간인 12개월과 123주 니깐 15개 4*15개의 최저최고가가 발생한다.
            1. 15개의 기간의 값의 차이는 크게 주단위와 월단위로 구분할수 있고
               1. 일수에 따라 구분할수 있다. 
               2. 하지만 특정일 기준으로 365일 전 가격목록에서부터 정확히 365/7= 52.1428571429 인것처럼 나오진 않는다.
                  1. 그러므로 7일을 일주일로 가정하고 30일을 한달로 가정한다는 기준이 필요하다.
                  2. 그럼 365/30 = 12.1666666667
               3. 목록을 기간별로 끊으려면 끊으려는 기준을 정해야함.
                  1. 7일은 1주일 14일은 2주 21일은 3주 30일은 한달 60일은 두달 이게 기준이
               
         3. 60개의 최고가 최저가의 정보를 디비에 저장한다. 
      2. 목록에서 가격값이 최고인 값과 최저인 값 찾기
   4. 저창 채널은 받은 후 저장(upsert 하지말고 insert로만 하기)
2. public.info 에 테이블 완료 됬다고 완료시간 업데이트 하기.


현재 가격이 52주와 얼마나 가까이에 있는 지를 파악한다.
그럼 테이블에는 얼마나 가까이에 있는지를 파악할수 있는 값이 존재 해야되지 않을까?
그리고 hiph_price와 low_price를 저장하지만 high_price_dt와 low_price_dt는 저장하지 않는다. 이것도 수정하자.

```
p1. sql로 실행

p1.1~4 golang
    dao select
    for
    chan
    chan
    dao insert
```
### config
meta.config에 추가할것 
기간 1~12개월 텍스트값, 1~3주 텍스트값   
config에서 기간 조회 dao 추가해야됨. .env로 하기    




