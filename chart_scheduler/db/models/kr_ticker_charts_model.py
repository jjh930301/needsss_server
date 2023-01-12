from datetime import datetime
from sqlalchemy import BigInteger, Column, Date, DateTime, ForeignKey, Index, Integer, String , Float , DECIMAL
import sqlalchemy
from sqlalchemy.orm import relationship , backref
from chart_scheduler.db.session import Base , meta_obj
from chart_scheduler.constants.schema_names import SCHEMA_NAMES

class KrTickerChartModel(Base) :
  __tablename__ = SCHEMA_NAMES.KR_TICKER_CHARTS
  RELATIONSHIPS_TO_DICT = True
  __metaclass__=meta_obj
  __table_args__=(Index('partition_idx' , 'symbol' , 'date'),)

  symbol = Column(
    String(12), 
    index=True,
    primary_key=True
  )
  ticker_symbol = Column(
    String(12),
    nullable=False
  )
  date : datetime = Column(
    Date , 
    nullable=False , 
    primary_key=True,
    index=True,
  )
  open : int = Column(DECIMAL(precision=11 , scale=3) , nullable=False)
  high : int = Column(DECIMAL(precision=11 , scale=3) , nullable=False)
  low : int = Column(DECIMAL(precision=11 , scale=3) , nullable=False)
  close : int = Column(DECIMAL(precision=11 , scale=3) , nullable=False)
  volume : int = Column(BigInteger , nullable=False)
  percent : float = Column(Float , nullable=False , default=0)
  ticker = relationship(
    'KrTickerModel', # class name
    back_populates='charts',
  )