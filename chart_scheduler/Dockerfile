FROM python:3.8-slim

ENV PYTHONDONTWRITEBYTECODE 1
ENV PYTHONUNBUFFERED 1 
ENV PYTHONPATH "${PYTHONPATH}:/code"

WORKDIR /code

COPY requirements.txt ./

RUN pip install --upgrade pip
RUN pip install -r requirements.txt

COPY . /code/chart_scheduler

CMD if [ "$ENV" = "production" ]; \
      then gunicorn chart_scheduler.main:app --preload --workers 100 --worker-class uvicorn.workers.UvicornWorker --bind 0.0.0.0:3000;  \
    else \
      uvicorn chart_scheduler.main:app --reload --host 0.0.0.0 --workers 2 --port 3000 --ws-ping-timeout 10000; fi
EXPOSE 3000
