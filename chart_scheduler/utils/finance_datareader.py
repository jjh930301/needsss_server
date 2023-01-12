import FinanceDataReader as fdr
from chart_scheduler.utils.date_util import get_day

# 심볼	설명
# KRX   한국 시장
# KS11	KOSPI 지수
# KQ11	KOSDAQ 지수
# KS50	KOSPI 50 지수
# KS100 KOSPI 100
# DJI	다우존스 지수
# IXIC	나스닥 지수
# US500 S & P 500 지수
# JP225	닛케이 225 선물
# STOXX50E	Euro Stoxx 50
# CSI300	CSI 300 (중국)
# HSI	항셍 (홍콩)
# FTSE	영국 FTSE
# DAX	독일 DAX 30
# CAC	프랑스 CAC 40

def daily_values(market_type : str) -> list:
	try:
		index = fdr.StockListing(market_type).reset_index()
		# insert KRX tickers
		tickers = index.to_dict('records')

		return tickers
	except Exception as e:
		print(e)
		return None

async def before_charts(symbol : str , day : str):
	try:
		charts : list = fdr.DataReader(symbol=symbol , start=day).reset_index()\
			.rename(columns={
				"Close": "close", 
				"Date": "date",
				"Open": "open", 
				"High": "high", 
				"Low": "low", 
				"Volume": "volume",
				"Change" : "change"
			}).to_dict('records')
		return charts
	except Exception as e :
		print(e)
		return None