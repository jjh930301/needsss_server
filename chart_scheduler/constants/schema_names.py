from chart_scheduler.constants.const import constant

class _SchemaNames(object):
	@constant
	def USER():
		return "users"

	@constant
	def KR_TICKERS():
		return "kr_tickers"

	@constant
	def KR_TICKER_CHARTS():
		return 'kr_ticker_charts'

	@constant
	def INTERESTED_TICKERS():
		return 'interested_tickers'
	
	@constant
	def USERS():
		return 'users'

SCHEMA_NAMES = _SchemaNames()