from fastapi import APIRouter
from fastapi.responses import JSONResponse
from chart_scheduler.db.session import database


router = APIRouter(tags=['health'])

@router.get('/check')
async def health_checker():
  return JSONResponse(
    {
      'result' : True,
      'database' : database.is_connected
    },
    status_code=200
  )
