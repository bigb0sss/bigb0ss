# 14. Write a Python program to calculate number of days between two dates.
# Sample dates : (2014, 7, 2), (2014, 7, 11)
# Expected output : 9 days 

from datetime import date

firstDate = date(2014, 7, 2)
lastDate = date(2014, 7, 11)

diff = lastDate - firstDate
print("[+] First Date: {0}".format(firstDate))
print("[+] Last Date : {0}".format(lastDate))
print("[+] Difference: {0} Days".format(diff.days))
