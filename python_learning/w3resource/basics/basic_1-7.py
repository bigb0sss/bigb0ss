# 7. Write a Python program to accept a filename from the user and print the extension of that. 
# Sample filename : abc.java
# Output : java

value = input("[*] Enter Filename: ")
filename = value.split(".")[0]
ext = value.split(".")[-1]

print("[+] Filename:", filename)
print("[+] Extension:", ext)
