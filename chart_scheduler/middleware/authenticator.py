import os
from fastapi import Depends
from fastapi.responses import JSONResponse
import jwt
from chart_scheduler.db.session import db_session
from sqlalchemy import text

from fastapi.security import HTTPAuthorizationCredentials, HTTPBearer

security = HTTPBearer()

async def authenticator( 
    credentials: HTTPAuthorizationCredentials= Depends(security)
) :
  token = credentials.credentials   
  if token is None or token == "":
    return JSONResponse(
      content={
        "result" : None
      },
      status_code=403
    )
  try:
    verification = jwt.decode(
      token,
      os.environ["JWT_ACCESS_SECRET"],
      algorithms='HS256'
    )
  except jwt.DecodeError as e:
    return JSONResponse(
      content={
        "result" : None
      },
      status_code=403
    )
  try:
    async with db_session() as session :
      data = await session.execute(text(f"""select * from MARKET.users where mobile = '{verification['id']}' limit 1"""))
      user = data.fetchone()
      if user :
        return user
      else :
        return JSONResponse(
          content={
            "result" : None
          },
          status_code=403
        )
  except Exception as e:
      return JSONResponse(
        content={
          "result" : None
        },
        status_code=500
      )


