from math import isnan
from sqlalchemy import text
from chart_scheduler.utils.finance_datareader import before_charts
from chart_scheduler.db.models.kr_ticker_charts_model import KrTickerChartModel
from chart_scheduler.db.session import database , db_session
from chart_scheduler.utils.date_util import get_day
from sqlalchemy.dialects.mysql import Insert
import sys

async def before_insert_chart(
  symbol : str,
  day : str
) :
  try :
    charts : list = await before_charts(
      symbol=symbol,
      day=day
    )
    if len(charts) != 0 and charts is not None :
      for chart in charts :
        if int(chart['open']) != 0 and \
          int(chart['high']) != 0 and \
          int(chart['low']) != 0 and \
          int(chart['close']) != 0 : 
          data = {
            'symbol' : symbol,
            'ticker_symbol': symbol,
            'date' : chart['date'], 
            'open' : chart['open'] if isinstance(chart['open']  , int) and isnan(isinstance(chart['open'] , float)) is False else 0 ,
            'high' : chart['high'] if isinstance(chart['high'] , int) and isnan(isinstance(chart['high'] , float)) is False else 0 ,
            'low' : chart['low'] if isinstance(chart['low'] , int) and isnan(isinstance(chart['low'] , float)) is False else 0 ,
            'close' : chart['close'] if isinstance(chart['close'] , int) and isnan(isinstance(chart['close'] , float)) is False else 0 ,
            'percent' : round(chart['change'] * 100 , 2) if isinstance(chart['change'] , float) and isnan(round(chart['change'] * 100 , 2)) is False else 0 ,
            'volume' : chart['volume'] if isinstance(chart['volume'] , int) and isnan(isinstance(chart['volume'] , float)) is False else 0,
            'symbol' : symbol,
          }
          insert_stmt = Insert(KrTickerChartModel).values(data)
          duplication = insert_stmt.on_duplicate_key_update(data)
          await database.execute(duplication)
          # print("Success {}".format(symbol) , file=sys.stdout)
  except Exception as e :
    e
  

async def fdr_chart_duplicate_insert(data : dict) :
  try :
    if int(data['Close']) != 0 and \
      int(data['Low']) != 0 and \
      int(data['High']) != 0 and \
      int(data['Open']) != 0 and \
      database.is_connected == True:
      obj = {
        'symbol' : data['Code'],
        'ticker_symbol': data['Code'],
        'date' : get_day(0),
        'open' : int(data['Open']),
        'high' : int(data['High']),
        'low' : int(data['Low']),
        'close' : int(data['Close']),
        'percent' : round(data['ChagesRatio'], 2),
        'volume' : int(data['Volume']),
      }
      insert_stmt = Insert(KrTickerChartModel).values(obj)
      duplication = insert_stmt.on_duplicate_key_update(obj)
      await database.execute(duplication)
  except Exception as e:
    e

async def pykrx_chart_duplicate_insert(data : dict) :
  # {
  #   '티커': '007980', 
  #   '시가': 1850, 
  #   '고가': 2100, 
  #   '저가': 1850, 
  #   '종가': 1975, 
  #   '거래량': 13079701, 
  #   '거래대금': 26352351970, 
  #   '등락률': 6.760000228881836
  # }
  try :
    if int(data['시가']) != 0 and \
      int(data['고가']) != 0 and \
      int(data['저가']) != 0 and \
      int(data['종가']) != 0 and\
      database.is_connected == True:
      
      obj = {
        'symbol' : data['티커'],
        'ticker_symbol': data['티커'],
        'date' : get_day(0),
        'open' : int(data['시가']),
        'high' : int(data['고가']),
        'low' : int(data['저가']),
        'close' : int(data['종가']),
        'volume' : int(data['거래량']),
        'percent' : round(data['등락률'], 2),
      }
      
      insert_stmt = Insert(KrTickerChartModel).values(obj)
      duplication = insert_stmt.on_duplicate_key_update(obj)
      await database.execute(duplication)
  except Exception as e:
    e