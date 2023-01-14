from datetime import datetime
from typing import List
from sqlalchemy import Column, DECIMAL, String
from sqlalchemy.orm import relationship , backref
from sqlalchemy.dialects.mysql import TINYINT , DATE , SMALLINT
from chart_scheduler.db.models.kr_ticker_charts_model import KrTickerChartModel
from chart_scheduler.db.session import Base , meta_obj
from chart_scheduler.constants.schema_names import SCHEMA_NAMES

class KrTickerModel(Base) :
  __tablename__ = SCHEMA_NAMES.KR_TICKERS
  RELATIONSHIPS_TO_DICT = True
  __metaclass__=meta_obj

  symbol : str = Column(String(12) , nullable=False , index=True , primary_key=True)
  market : int = Column(TINYINT , nullable=True)
  index : int = Column(SMALLINT , nullable=True)
  name : str = Column(String(100) , nullable=False , index=True)
  bps : int = Column(DECIMAL(precision=11 , scale=2) , nullable=True , default=0)
  per : int = Column(DECIMAL(precision=11 , scale=2) , nullable=True , default=0)
  pbr : int = Column(DECIMAL(precision=11 , scale=2) , nullable=True , default=0)
  eps : int = Column(DECIMAL(precision=11 , scale=2) , nullable=True , default=0)
  div : int = Column(DECIMAL(precision=11 , scale=2) , nullable=True , default=0)
  dps : int = Column(DECIMAL(precision=11 , scale=2) , nullable=True , default=0)
  sector : str = Column(String(150) , nullable=True)
  industry : str = Column(String(150) , nullable=True)
  listing_date : datetime = Column(DATE , nullable=True)
  settle_month : str = Column(String(10) , nullable=True)
  representative : str = Column(String(100) , nullable=True)
  homepage : str = Column(String(200) , nullable=True)
  market_cap : int = Column(DECIMAL(precision=17 , scale=0) , nullable=True , default=None) 
  region : str = Column(String(10) , nullable=True)
  charts : List[KrTickerChartModel] = relationship(
    'KrTickerChartModel',
    back_populates='ticker',
    uselist=True
  )