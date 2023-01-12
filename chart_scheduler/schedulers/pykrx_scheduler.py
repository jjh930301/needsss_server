from fastapi import BackgroundTasks
from rocketry import Rocketry
from rocketry.conds import every, after_success , crontime , cron
from chart_scheduler.repositories.kr_ticker_charts_repository import pykrx_chart_duplicate_insert
from chart_scheduler.repositories.kr_tickers_repository import pykrx_ticker_fundamental
import chart_scheduler.utils.pykrx_util as pykrx
import asyncio

app = Rocketry(config={
  "task_execution": "async",
  'max_process_count': 5,
  'restarting' : 'replace'
})

# pykrx는 request가 많을 때 0으로 떨어지기 때문에 interval을 길게 둠
@app.task(cron('* 9-16 * * *') & cron('* * * * 1-5') & every('40 seconds'))
async def insert_kr_ticker_charts():
  try :
    print('start pykx')
    charts = pykrx.daily_values("ALL")
    for chart in charts:
      await pykrx_chart_duplicate_insert(chart)
    print('end pykrx')
  except Exception as e:
    print(e)

@app.task(cron('* 16 * * *') & cron('* * * * 1-5'))
async def ticker_fundamental():
  try:
    print('start fundamental')
    tickers = pykrx.get_market_fundamental()
    for ticker in tickers:
      await pykrx_ticker_fundamental(ticker)
  except Exception as e:
    print(e)
  
if __name__ == "__main__":
  asyncio.run(app.run())