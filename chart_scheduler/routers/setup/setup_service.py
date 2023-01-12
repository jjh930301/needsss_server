from fastapi import BackgroundTasks
from chart_scheduler.repositories.kr_ticker_charts_repository import before_insert_chart , fdr_chart_duplicate_insert , pykrx_chart_duplicate_insert
from chart_scheduler.repositories.kr_tickers_repository import pykrx_ticker_fundamental , update_index
from chart_scheduler.db.session import database
from chart_scheduler.enum.indexes import Indexes
import chart_scheduler.utils.finance_datareader as fdr
import chart_scheduler.utils.pykrx_util as pykrx
import time

async def before_chart(
  background_tasks : BackgroundTasks,
  day : str
) :
  try :
    symbols = await database.fetch_all(query='SELECT symbol FROM kr_tickers where market != 2')
    for symbol in symbols :
      background_tasks.add_task(
        before_insert_chart,
        symbol.symbol,
        day
      )
    return True
  except Exception as e: 
    print(e)

async def fdr_chart_insert(background_task : BackgroundTasks) :
  try :
    charts = await fdr.daily_values('KRX') 
    for chart in charts:
      background_task.add_task(
        fdr_chart_duplicate_insert,
        chart
      )
  except Exception as e :
    print(e)

async def pykrx_chart_insert(background_task : BackgroundTasks):
  try :
    charts = pykrx.daily_values("ALL")
    for chart in charts:
      background_task(
        pykrx_chart_duplicate_insert,
        chart
      )
    return True
  except Exception as e:
    print(e)
    return e

async def fundamental() :
  try:
    tickers = pykrx.get_market_fundamental()
    for ticker in tickers:
      await pykrx_ticker_fundamental(ticker)
  except Exception as e:
    print(e)
    return e

async def index():
  try :
    for index in Indexes :
      time.sleep(10)
      tickers = pykrx.get_index_portfolio_deposit_file(index=str(index.value))
      for ticker in tickers :
        await update_index(ticker , index.value)
  except Exception as e:
    print(e)
    return e