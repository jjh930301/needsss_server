from chart_scheduler.enum.market_type import MARKET_TYPE
from chart_scheduler.db.session import database
from chart_scheduler.db.models.kr_tickers_model import KrTickerModel
from sqlalchemy.dialects.mysql import Insert
from sqlalchemy import update

async def fdr_ticker_duplicate_insert(
    data : dict,
) :
  market = {
    "KOSDAQ" : MARKET_TYPE.KOSDAQ.value,
    "KOSPI" : MARKET_TYPE.KOSPI.value,
    "KONEX" : MARKET_TYPE.KONEX.value,
  }
  try :
    if type(data['Marcap']) is int: 
      if market[data['Market']] == MARKET_TYPE.KOSPI.value or market[data['Market']] == MARKET_TYPE.KOSDAQ.value :
        data = {
          'symbol':data['Code'],
          # 'ticker_symbol': data['Code'],
          'market':market[data['Market']],
          'name':data['Name'],
          'market_cap' : int(data['Marcap'])
        }
        insert_stmt = Insert(KrTickerModel).values(data)
        duplication = insert_stmt.on_duplicate_key_update(data)
        await database.execute(duplication) 
  except Exception as e:
    return None

async def pykrx_ticker_fundamental(
  data : dict
) :
  try:
    if int(data['티커']) != 0 and \
      int(data['BPS']) != 0 and \
      int(data['PER']) != 0 and \
      int(data['PBR']) != 0 and \
      int(data['EPS']) != 0 and \
      int(data['DIV']) != 0 and \
      int(data['DPS']) != 0 :
      
      ticker = update(KrTickerModel)\
        .where(KrTickerModel.symbol == data['티커'])\
        .values({
          'bps' : data['BPS'],
          'per' : data['PER'],
          'pbr' : data['PBR'],
          'eps' : data['EPS'],
          'div' : data['DIV'],
          'dps' : data['DPS'],
        })\
        .execution_options(synchronize_session='fetch')
      
      await database.execute(ticker)
  except Exception as e:
    return None

async def update_index(
  symbol : str,
  index : int
):
  try :
    ticker = update(KrTickerModel)\
      .where(KrTickerModel.symbol == symbol)\
      .values({
        'index' : index
      })\
      .execution_options(synchronize_session='fetch')
    await database.execute(ticker)
  except Exception as e:
    print(e)
    return None
