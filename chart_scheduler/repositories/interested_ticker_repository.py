from sqlalchemy import text
from chart_scheduler.constants.schema_names import SCHEMA_NAMES
from chart_scheduler.db.session import db_session

async def find_one(symbol : str):
  async with db_session() as session:
    query = await session.execute(text(f"""select symbol from {SCHEMA_NAMES.INTERESTED_TICKERS} where symbol = '{symbol}'limit 1"""))
    model = query.fetchone()
    if model is not None:
      return model
    else :
      return None