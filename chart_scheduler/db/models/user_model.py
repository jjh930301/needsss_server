from sqlalchemy import Column , DECIMAL , String , ForeignKey , DateTime , BigInteger , Float , Date
from sqlalchemy.orm import relationship
from typing import List
from chart_scheduler.db.session import Base , meta_obj
from chart_scheduler.constants.schema_names import SCHEMA_NAMES
# from chart_scheduler.db.models.interested_tickers_model import InterestedTickersModel
import uuid

class UserModel(Base) :

  __tablename__ = SCHEMA_NAMES.USERS
  RELATIONSHIPS_TO_DICT = True
  __metaclass__=meta_obj

  id : str = Column(
    String(36), 
    primary_key=True, 
    default=uuid.uuid4
  )

  email : str = Column(String(100) , nullable=False)
  password : str = Column(String(255) , nullable=False)
  mobile : str = Column(String(12))
  interestes : list = relationship(
    'InterestedTickersModel',
    back_populates='user',
    uselist=True
  )
