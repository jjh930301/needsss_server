import databases
from sqlalchemy.ext.asyncio import create_async_engine, AsyncSession
from sqlalchemy import MetaData, create_engine
from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy.orm import sessionmaker , scoped_session 
from sqlalchemy.ext.asyncio import async_scoped_session
from asyncio import current_task
import os

# only mysql container
DATABASE_URL = "mysql+aiomysql://{}:{}@mysql:3306/{}?charset=utf8mb4"\
  .format(
    'root',
    os.environ['MYSQL_ROOT_PASSWORD'],
    os.environ['MYSQL_DATABASE']
  )
database = databases.Database(DATABASE_URL)

engine = create_async_engine(DATABASE_URL)

db_session = async_scoped_session(
  sessionmaker(
    autocommit=False,
    autoflush=False,
    bind=engine,
    class_=AsyncSession
  ),
  scopefunc=current_task,
)
Base = declarative_base()

meta_obj = MetaData()