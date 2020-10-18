# 12. Write a Python program to print the calendar of a given month and year.
# Note : Use 'calendar' module. 

import calendar
year = int(input("[*] Input the Year : "))
month = int(input("[*] Input the Month : "))
print("")
print(calendar.month(year, month))
