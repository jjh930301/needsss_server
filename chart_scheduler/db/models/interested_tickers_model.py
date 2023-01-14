from datetime import datetime
from sqlalchemy import Column , DECIMAL , String , ForeignKey , DateTime , BigInteger , Float , Date
from sqlalchemy.dialects.mysql import TINYINT 
from typing import List
from sqlalchemy.orm import relationship , backref
from chart_scheduler.db.session import Base , meta_obj
from chart_scheduler.constants.schema_names import SCHEMA_NAMES
# from chart_scheduler.db.models.kr_ticker_charts_model import KrTickerChartModel
from chart_scheduler.db.models.user_model import UserModel
# from chart_scheduler.db.models.kr_tickers_model import KrTickerModel

class InterestedTickersModel(Base) :
  __tablename__ = SCHEMA_NAMES.INTERESTED_TICKERS
  RELATIONSHIPS_TO_DICT = True
  __metaclass__=meta_obj

  symbol = Column(
    String(12), 
    ForeignKey("kr_ticker_charts.symbol"),
    index=True,
    primary_key=True
  )
  date : datetime = Column(
    Date , 
    nullable=False , 
    primary_key=True,
    index=True,
  )
  date_time : datetime = Column(
    DateTime , 
    nullable=False , 
    index=True,
  )
  # kospi : 1
  # kosdaq : 2
  # nasdaq : 3
  # dow : 4
  type : int = Column(TINYINT , nullable=False) 
  name : str = Column(String(100) , nullable=False , index=True)
  close : int = Column(DECIMAL(precision=11 , scale=3) , nullable=False)
  percent : float = Column(Float , nullable=False , default=0)
  volume : int = Column(BigInteger , nullable=False)
  sales_close : Column(DECIMAL(precision=11 , scale=3) , default=None)
  saled_at : datetime = Column(
    DateTime , 
    nullable=True , 
    default=None,
  )
  user_id = Column(String(36), ForeignKey('users.id'))
  user : UserModel = relationship(
    'UserModel',
    back_populates='interestes'
  )
  kr_ticker = relationship(
    'KrTickerModel',
    back_populates='interestes'
  )
  kr_charts : list = relationship(
    'KrTickerChartModel',
    back_populates='kr_ticker',
    uselist=True
  )
