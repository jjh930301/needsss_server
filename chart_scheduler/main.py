import asyncio
from fastapi import FastAPI
from chart_scheduler.db.session import database
from chart_scheduler.schedulers.datareader_scheduler import app as datareader
from chart_scheduler.schedulers.pykrx_scheduler import app as pykrx
from chart_scheduler.routers.health_router import router as health
from chart_scheduler.routers.setup.setup_router import router as setup
import socketio
import os

sio = socketio.Client()

app = FastAPI(docs_url='/docs')

@app.on_event("startup")
async def startup():
	sio.connect(os.environ['SOCKET_URI'])
	asyncio.create_task(datareader.serve())
	asyncio.create_task(pykrx.serve())
	await database.connect()
	

@app.on_event("shutdown")
async def shutdown():
	sio.disconnect()
	await database.disconnect()

app.include_router(health , prefix='/health')
app.include_router(setup , prefix='/setup')