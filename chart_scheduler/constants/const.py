def constant(func):
	def func_set(self, value):
		raise TypeError

	def func_get(self):
		return func()
	return property(func_get, func_set)