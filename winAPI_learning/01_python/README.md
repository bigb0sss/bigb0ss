# Python - WinAPI

| | Name | WinAPI | WinDLL | Note |
| :--- | :--- | :---: | :--- | :-- |
| 1 | messageBox.py | MessageBoxW | User32.dll | |
| 2 | openProcHandler.py | OpenProcess | Kernel32.dll | |
| 3 | procKiller.py | FindWindowA, GetWindowThreadProcessId, OpenProcess, TerminateProcess | User32.dll, Kernel32.dll | |
| 4 | createProc.py | CreateProcessW, PROCESS_INFORMATION, STARTUPINFOA | Kernel32.dll | |
| 5 | dnsCacheEntry.py | DnsGetCacheDataTable | Kernel32.dll, DNSAPI.dll | *Undocumented |

# Windows Token Privileges
Tokens are static; therefore, we cannot add/delete but we can enable/disable the current privileges (set by default).
* https://docs.microsoft.com/en-us/windows/win32/secauthz/privilege-constants

| | Privilege Value | Description | 
| :--- | :--- | :---: |
| 1 | SeDebugPrivilege | Required to debug and adjust the memory of a process owned by another account. User Right: Debug programs. |
