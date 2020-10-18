# 10. Write a Python program that accepts an integer (n) and computes the value of n+nn+nnn.
# Sample value of n is 5
# Expected Result : 615

num = int(input("[*] Enter the number: "))
num1 = int("%s" % num)
num2 = int("%s%s" % (num,num))
num3 = int("%s%s%s" % (num,num,num))
print(num + num2 + num3)
