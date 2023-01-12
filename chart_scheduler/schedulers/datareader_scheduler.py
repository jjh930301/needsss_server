from fastapi import BackgroundTasks
from rocketry import Rocketry
from rocketry.conds import every, after_success , crontime , cron
from datetime import datetime , timedelta
from chart_scheduler.repositories.kr_tickers_repository import fdr_ticker_duplicate_insert
from chart_scheduler.repositories.kr_ticker_charts_repository import fdr_chart_duplicate_insert
import chart_scheduler.utils.finance_datareader as fdr
import asyncio

app = Rocketry(config={
  "task_execution": "async",
  'max_process_count': 5,
  'restarting' : 'replace'
})

# 종목 가져오기
@app.task(cron('* 1 * * *') & cron('* * * * 1-5'))
async def insert_kr_tickers():
  try:
    charts = fdr.daily_values('KRX') 
    for chart in charts:
      fdr_ticker_duplicate_insert(chart)
  except Exception as e :
    e

# 차트 업데이트9시 ~ 15시 
@app.task(cron('* 9-15 * * *') & cron('* * * * 1-5') & every('10 seconds'))
async def insert_kr_ticker_charts():
  try :
    charts = fdr.daily_values('KRX') 
    for chart in charts :
      await fdr_chart_duplicate_insert(chart)
  except Exception as e:
    print(e)

if __name__ == "__main__":
  asyncio.run(app.run())