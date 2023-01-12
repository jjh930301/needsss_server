from pykrx import stock
from chart_scheduler.utils.date_util import get_day
# https://github.com/sharebook-kr/pykrx
# 일자별 종목 등락 리스트
def daily_values(market : str) :
	try :
		# 티커 , 시가 , 고가 , 저가 , 종가 , 거래량 , 거래대금 , 등락률
		data : list = stock.get_market_ohlcv(get_day(0) , market=market)\
			.reset_index()\
			.to_dict('records')
		return data
	except Exception as e : 
		print(e)
		
# 일자별 시가총액 리스트
async def get_market_cap(
	start_date : str,
	end_date : str,
	ticker : str
) -> list:
	try :
		# 날짜 , 시가총액 , 거래대금 , 상장 주식수
		data : list = await stock.get_market_cap(start_date , end_date , ticker , 'd')\
			.reset_index()\
			.to_dict('records')
		return data
	except Exception as e:
		return None

# 1001 코스피
# 1028 코스피 200
# 1034 코스피 100
# 1035 코스피 50
# 1167 코스피 200 중소형주
# 1182 코스피 200 초대형제외 지수
# 1244 코스피200제외 코스피지수
# 1150 코스피 200 커뮤니케이션서비스
# 1151 코스피 200 건설
# 1152 코스피 200 중공업
# 1153 코스피 200 철강/소재
# 1154 코스피 200 에너지/화학
# 1155 코스피 200 정보기술
# 1156 코스피 200 금융
# 1157 코스피 200 생활소비재
# 1158 코스피 200 경기소비재
# 1159 코스피 200 산업재
# 1160 코스피 200 헬스케어
# 1005 음식료품
# 1006 섬유의복
# 1007 종이목재
# 1008 화학
# 1009 의약품
# 1010 비금속광물
# 1011 철강금속
# 1012 기계
# 1013 전기전자
# 1014 의료정밀
# 1015 운수장비
# 1016 유통업
# 1017 전기가스업
# 1018 건설업
# 1019 운수창고업
# 1020 통신업
# 1021 금융업
# 1022 은행
# 1024 증권
# 1025 보험
# 1026 서비스업
# 1027 제조업
# 1002 코스피 대형주
# 1003 코스피 중형주
# 1004 코스피 소형주
# 1224 코스피 200 비중상한 30%
# 1227 코스피 200 비중상한 25%
# 1232 코스피 200 비중상한 20%
async def get_index_chart(
	start_date : str,
	end_date : str,
	idx : str
) -> list:
	try :
		data : list = await stock\
			.get_index_ohlcv(start_date , end_date , idx , 'd')\
			.reset_index()\
			.to_dict('records')
		return data
	except Exception as e:
		return None

# 인덱스에 해당하는 종목을 가져 옵니다
def get_index_portfolio_deposit_file(
	index : str
) -> list :
	try:
		data : list = stock.get_index_portfolio_deposit_file(index)
		return data
	except Exception as e:
		return None
# 코스피 시장의 DIV/BPS/PER/EPS/PBR
def get_market_fundamental() :
	try :
		data : list = stock.get_market_fundamental(get_day(0) , market="ALL")\
			.reset_index()\
			.to_dict('records')
		return data
	except Exception as e :
		return None