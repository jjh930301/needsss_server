from datetime import datetime, timedelta

def get_day(before) :
	return (
		(
				datetime.today() # Asia/Seoul
		) - timedelta(days=before) 
	).strftime('%Y-%m-%d') # YYYY-MM-dd

