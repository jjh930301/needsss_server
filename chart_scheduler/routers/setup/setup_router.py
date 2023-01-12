from fastapi import APIRouter, BackgroundTasks , Depends
from fastapi.responses import JSONResponse
import chart_scheduler.utils.finance_datareader as fdr
import chart_scheduler.routers.setup.setup_service as setup_service
from chart_scheduler.db.models.user_model import UserModel
from chart_scheduler.middleware.authenticator import authenticator
from chart_scheduler.repositories.kr_tickers_repository import fdr_ticker_duplicate_insert

router = APIRouter(tags=['setup'])

@router.get('/ticker')
async def ticker(
  background_tasks: BackgroundTasks,
  user : UserModel = Depends(authenticator)
) :
  try:
    charts = fdr.daily_values('KRX') 
    for chart in charts:
      background_tasks.add_task(
        fdr_ticker_duplicate_insert,
        chart
      )
    
    return JSONResponse(
      content={
        'result' : True
      },
      status_code=200
    )
  except Exception as e :
    e
  
@router.get('/fdr/before/{day}')
async def charts(
  background_tasks: BackgroundTasks,
  day : str,
  user : UserModel = Depends(authenticator)
):
  try:
    await setup_service.before_chart(
      background_tasks=background_tasks,
      day=day if day is None else '2022-01-01'
    )
    return JSONResponse(
      content={
        'result' : True
      },
      status_code=200
    )
  except Exception as e:
    e

@router.get('/pykrx/fundamental')
async def fundamental(
  background_task : BackgroundTasks,
  user : UserModel = Depends(authenticator)
) :
  try :
    await setup_service.fundamental(
      background_task=background_task
    )
    return JSONResponse(
      content={
        'result' : True
      },
      status_code=200
    )
  except Exception as e:
    return JSONResponse(
      content=None,
      status_code=500
    )

@router.get('/pykrx/index')
async def index(
  background_task :BackgroundTasks,
  user : UserModel = Depends(authenticator)
) :
  try :
      await setup_service.index(background_task=background_task)
      return JSONResponse(
        content={
          'result' : True
        },
        status_code=200
      )
  except Exception as e:
    return JSONResponse(
      content=None,
      status_code=500
    )