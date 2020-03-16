**This is not my review I copied it from HTB forum. All credit is to the author (21y4d)

A couple of months after I earned my OSCP, I knew that my next step was going to be OSWE.

For the past 6 moths or so I’ve been busy preparing for the Offensive Security Web Expert (OSWE) certificate. I’ve had this certification on my plan, and once it was announced for the public in 2019, I started preparing to enroll in its course. I couple of months ago I registered for the OSWE course "Advnaced Web Attacks and Exploitation - AWAE". Last week, I had my 48-hour OSWE exam, and this morning I was informed that I successfully passed the exam and earned OSWE.

During my preparation period, and when I was preparing for the exam, I found that there’re no proper reviews for OSWE, since it is a relatively new certificate “2019”, and I found myself needing answers for several questions. So, I’ve decided to share with you a detailed OSWE review, for anyone who’s interested or planning to take OSWE.



Who it’s for "vs OSCP"

The very first thing that must be clearly understood, OSWE is not a successor to OSCP, nor is OSCE, but they are rather advanced courses in one of its fields.

OSCP is designed as an pentester certification, and hence it covers all of the main fields of pentesting, like Recon, Privilege Escalation, Network pivoting, Binary Exploitation, and Web Exploitation, at a professional level.
You can find my OSCP review here:

OSCP Exam review ''2019'' + Notes & Gift inside!

To be an “expert” in pentesting, there’s no one certificate that will cover expert level testing in all of the above fields, as the course would be 500 hours long, and the exam would 2 weeks long!

This is why for the “expert” level, there’s a separate certification in each field. For Advanced Binary/OS Exploitation there’re GXPN/OSCE/OSEE, for Advanced Web Exploitation there’s OSWE, and so on.

So, the question is, who is this certification for? Who should consider taking it?

In my humble opinion, as its name suggest, OSWE is for anyone looking to be an expert in web exploitation, where OSCP skills would just not be enough.

Such cases are when you are doing pentesting for critical web applications, like online banking platforms or globally used web apps, and you need to perform advanced pentesting at the source code level. OSWE might also be very appealing to people interested in bug bounty hunting, given that the majority of bug bounties are focused on web applications.

As for people who are purely doing blackbox pentesting, OSWE might not be very useful for them, since it is purely designed around whitebox testing. You might not find the benefits directly improving your work, as if you’re only doing blackbox testing, you may not be able to apply most of the things you learned at AWAE.

Still, I don’t want to discourage anyone thinking about OSWE from taking it. I think AWAE would still be a great learning experience, and you will definitely learn a lot of new stuff. Just make sure that you are spending your money and time on something that will benefit your work.

In general, AWAE is an advanced course, and OSWE is an "expert" level certificate. If you compare it to HTB boxes, it will definitely be around Hard/Insane difficulty, while PWK/OSCP would probably be around Easy/Medium difficulty.



Preparation & Practice Material

I always prefer to register only when I feel fully prepared to be certified, rather than going for the course and learning it's modules from scratch in its labs, since this might take much more lab time.

As mentioned earlier, when I wanted to start preparing for OSWE, I had no idea what topics I should learn, for the lack of review material. In the end, I set a plan for myself to follow before registering for AWAE, based on the course syllabus published by OffSec, which you can find here:
https://www.offensive-security.com/documentation/awae-syllabus.pdf

The following are the courses I took in web development to prepare for AWAE:

    NodeJS web app development
    C# .Net MVC web app development
    Java web app development

I did not go very deep in each, and simply learned what is necessary to understand how to read the code, undestand how url routing works in each language, and be able to understand the functionality of the web app.

You might also consider the following:

    Python exploit development
    PHP web app development
    SQL scripting
    JavaScript development
    JavaScript for Pentesters "Pentester Academy"
    WAP challenges "Pentester Academy"



My Gift to HTB members

My previous gift to HTB members, in my OSCP review, was the nmapAutomator tool, which I was glad to see well received, and got over 10k downloads from GitHub. You can find it here:
https://github.com/21y4d/nmapAutomator

While nmapAutomator took me a couple of weeks to finish, this one took me a couple of months of daily work, and a lot more lines of code, in many different languages!

Once I finished the above plan, I wanted to test my capabilities in whitebox testing before enrolling to the AWAE course. So, I looked through HTB for boxes with whitebox pentesting attack vectors, but didn’t find any!

At that time, I had owned almost all active and retired HTB boxes “except for perhaps 5”, and in almost none of them did I find whitebox attack paths. Even in the ones that did provide the source code, they were either very basic and can be done without the source code, or simply did not focus on whitebox pentesting.

So, I had the great idea of creating an HTB box that is completely focused on whitebox testing, and so I made sourceCode!

I wanted to make it a really enjoyable box with a great learning experience, and worked at every detail, which I hope will show in the final result.



sourceCode.htb

In summary, sourceCode requires you to be able to exploit scripts written in 6 different languages, none of which can be exploited without access to their source code, hence whitebox testing. The box also has a secondary theme, which I'll keep for when it's released.

As you go through it, sourceCode gets gradually more difficult. It does not focus on enumeration much, but on source code review instead. So, you always know what your next step is, and know what you must learn to exploit the current objective.

I even went ahead and included OSCE material, such that it is an excellent learning experience for both certificates, and can test your readiness for both.

I ranked it as Insane, which is appropriate once you complete the entire box. Still, I hope everyone will find something useful, even if they do not finish the entire box.

This was an excellent learning experience for me, as well, since I had to write from scratch scripts in 6 different languages, and make them in a way that they all interact together. I also had to make each of them vulnerable in a unique way, so I had to deeply understand each vulnerability and its whitebox testing technique, which turned out to be really useful for OSWE.

I submitted the box around 4 months ago, and I'm hoping it gets reviewed and approved soon. Shout out to @egotisticalSW for the great work he's been doing on reviewing box submissions.

I hope that sourceCode would be a great learning experience for all those considering taking OSWE or OSCE, and an excellent way to test your whitebox testing techniques before enrolling to OSWE.



AWAE Course/Lab

Before going through my Course/Exam review, I have to mention that I’m very restricted in what I’m allowed to say, so I will only mention in detail what is publicly announced by OffSec.

In general, I really enjoyed the Advanced Web Attacks and Exploitation (AWAE) course, and was learning new techniques at every module. When compared to OSCP, in my opinion, the AWAE course was miles better and much more enjoyable than PWK. This is not to say that PWK isn’t good, but rather because AWAE was so excellent and up-to-date.

I booked a couple of months lab time for $1600, and I think that was an overkill. The AWAE lab has 5 machines only for the course's 6 modules, and you should be easily able to complete all machines within just 2 weeks.

However, this will only be the case if you already finished the practice material I mentioned above, and had good whitebox testing skills. Otherwise, it might take you much more lab time to finish all machines and fully understand all modules, as they are quite advanced and complicated.

There are some areas in which AWAE can improve. First, AWAE does not provide enough vulnerability discovery techniques, and sometimes assumes the public knowledge of the vulnerability, while in real life, you must know techniques with which you can identify such vulnerabilities. Also, it would be better to teach a mixture of blackbox/whitebox discovery techniques, to make it easier to spot a vulnerability.

Another area in which AWAE can improve is huge code-base review techniques. If you have huge applications, without such techniques, none of the taught material would be helpful, so I think this is very critical.

Finally, AWAE would use a bit more extra-miles, as some vulnerabilities did not have extra practice material in the labs. It would also be good if such extra-miles covered the same vulnerability but in a different language, as what they did for deserialization, in which the course taught .Net deserializaiton, but there was an extra mile for a one in Java.

Once I finished all machines, I booked my exam “discussed later”, and created an exam preparation plan based on what I learned in the AWAE labs “discussed next”. Since I already had more than a month remaining in AWAE labs, I utilized the labs for my exam preparation, and completed all extra miles during that time, which I strongly recommend.



OSWE Exam Preparation

Once I finished my AWAE lab machines, and finished some extra miles, I wanted to use the time I had left by testing myself in each of the course modules.

My main plan was to find public exploits in each vulnerability type taught in AWAE, and then attempt to discover the vulnerability and write the exploit without reading the public exploit. However, this turned out to be difficult to arrange, since I had to go through so many vulnerabilities, and by the time I identified that the vulnerability suited the course, I already knew its attack vector, so it was no longer valid.

So, I created the following exam preparation plan for each course module. I included all AWAE extra-miles in the plan, and some others were based on HTB boxes, where I did not use the public exploit, but wrote it myself. You will need HTB VIP membership to access retired boxes, but a month in HTB “£10” is much cheaper than a month of AWAE "$400+".

Shout out to @ippsec for his great tool "http://ippsec.rocks/", which was quite handy in searching for topics for exam preparation.



Exam Preparation Plan

    .NET deserialization:
        DNN Extra-mile “AWAE labs”
        Never used in HTB

    Java deserialization:
        ManageEngine Extra-mile “AWAE labs”
        Arkham “HTB”

    PHP Type Juggling:
        ATutor Extra-mile “AWAE labs”
        Falafel “HTB”

    Advanced Authentication bypass:
        ATutor Extra-mile “AWAE labs”
        Didn’t find any in HTB

    CSRF to admin API execution:
        Atmail Extra-mile “AWAE labs”
        Didn’t find any in HTB

    Authenticated API to RCE:
        Atmail Extra-mile “AWAE labs”
        Zipper “HTB”
        HackBack “HTB”

    NodeJS Command injection:
        Bassmaster Safe-Eval Extra-mile “AWAE labs”
        Holiday “HTB”

    Boolean SQLi to RCE:
        Fighter “HTB”
        Using boolean SQLi instead of union, without receiving rev shell, with access to source code, with MSSQL debugging

    Time-Based SQLi data exfiltration:
        Writeup “HTB”
        Unattended “HTB”
        Help “HTB”
        Without SQLMap, without using/reading public exploit, using time-based SQLi instead of boolean

    Advanced File upload RCE:
        Ghoul “HTB“
        Help “HTB“
        Without using/reading public exploit



Python Exploit Development

During this time, among all of the above exercises, I also practiced python exploit development, which is a critical skill for the exam, without which you cannot pass. It doesn’t have to be python, but you have to be able to automate the exploitation process in a single exploit, and even chain multiple exploits together.

During my exam preparation, I wrote some fun exploits, such as:

    Fully interactive webshell over python, bypasses firewall, WAF, and AV
    Out of band SQL injection webshell, bypasses firewall, WAF, and AV
    Blind Time-Based/Boolean SQL injection script, blindly leaks data from DB

I wanted to share all of them with you here for practicing, however, since these are exploitation scripts, I fear they might be misused, so I don’t feel comfortable sharing them. I can, however, share the Blind SQLi script, since anyone can simply use SQLMap.

This script can show you what kind of exploits you should be able to develop, and should also be a good exercise to rewrite. You practice it on the SQLi boxes mentioned above. It’s around a 100 lines long, can exploit both time-based and boolean SQLi, and can be adapted to any SQL injectable url.

You can get it at:
https://github.com/21y4d/blindSQLi



Exam Scheduling

As mentioned earlier, once I completed all of the machines, I went ahead to book my exam. However, to my surprise, the earliest exam date was 3 months away!

This was an issue most people faced with OSWE exam booking. I think the reason is that “if I’m not mistaken” OffSec only takes one OSWE student at a time for their exam, since OSWE is a proctored exam as well. The OSWE exam is a 48-hour long session, and if only one student takes the exam every two days, at most 15 students will do it a month. If you have 50 or so registered students, you are looking at 3 months waiting time!

Still, if you keep checking the exam booking page every day, you will definitely find a closer booking. Every week or so somebody tends to cancel their exam, so you may be able to book their exam date.

My suggestion is to book the exam as early as possible, and then when you’re almost done with your exam preparation and ready to take the exam, you can keep checking the exam booking page for any cancellations. You will almost certainly find an appointment within 2-3 weeks of looking, and perhaps even within a week, which is what happened in my case.

You must note that if you do this, you will probably get an exam date only 3 days away, which is quite nerve racking, suddenly knowing that you’ll do the dreaded 48-hour exam in just a few of days. This is because people are not allowed to cancel their exam within 48-hours from the exam, and so people tend to cancel 3 or more days before the exam, if they feel not ready.

I booked the exam 2 months ahead “I think it was due mid March 2020”, and kept checking for any cancellations. Then someone cancelled a month away, so I booked that. Since you are only given 3 booking attempts, I only had one booking change left. When I saw a cancellation 3 days from then, I booked it, and had to take, because I had no more change attempts left. In any case, I was ready for the exam, so I didn’t hesitate, but it was still nerve racking at the time.



OSWE Exam

Then we come to the dreaded rigorous exam itself, which it absolutely was!

I started my exam at 9 AM, and submitted my report a 8 AM three days later! This means that I worked for 72 hours straight, and only slept around 15 hours during that time. My total work time, without the rest, was 35 hours out of 48, add to that 12 hours for writing the report. I was extremely tired after the exam, and needed 5 days to recover my strength. When compared to my OSCP time, I got the passing points in around 6 hours, and finished all machines in around 10.

Eventually, I was able to finish all objectives, got a score of 100/100 "assumed", and even found a couple of extra vulnerabilities to achieve some of the objectives. But that did not come easily or quickly!

The exam objectives difficulty was on par with the course’s extra miles, and perhaps slightly more difficult. That is not to say it was not difficult, because the course and it’s extra miles were very challenging, especially when you don't know their answers.

However, the most challenging part is vulnerability discovery, which when the code-base is several hundreds of thousands of lines long, can take a very long time if you have not practiced code review very well. Adding to that, the many possible vulnerabilities you must test, and the amount of fatigue you will feel after putting so many hours in, and you will find yourself struggling to keep your focus.

As mentioned earlier, I cannot say much about the exam itself, but I want to leave some notes for those preparing for the OSWE exam, without spoiling anything. Such notes would have really helped me during my preparation, but like I said, not many OSWE reviews are out there.

I know that many people will ask how the exam is designed, but unfortunately I’m not sure I can say that here. Basically, the exam would have a small number of machines, each has one debug and one real version. You have to do your whitebox testing on the debug server, until you find the vulnerability, and then write your exploit. Then, you must test your exploits on the real target, and eventually gather the flags. There are a number of objectives for each target, and you will start on the real target as an unauthenticated user.

My main advise would be to do the following before taking the exam:

    Have a plan! My plan was divided by language, and then sub-divided by authenticated and unauthenticated exploits. I had a prioritized list of potential attack vectors for each language, based on the AWAE course, such that I will attempt the attacks in that order if I got a machine in that language. I cannot share it here, to avoid any spoilers.

    Practice huge code base review in all languages taught in AWAE, such that you feel comfortable understanding the functionalities of the web app, and able to match the vulnerable function to the web UI or API.

In general, if you are well prepared and have a good plan at hand, each objective can be done in around 4-6 hours, which leaves plenty of time for resting and taking breaks.

However, if you do not have a good strategy, or do not have a good way of quickly reviewing huge code-bases, you will find yourself needing way more than the 48 hours provided. I think that is the point of the 48-hour time limit, to ensure that one not only knows how to identify and exploit the vulnerabilities, but also have good techniques on how to do that efficiently.



AWAE Course strengths and weaknesses

In summary, I think AWAE/OSWE was an excellent learning experience, which greatly improved my web application pentesting skills, and greatly helped me in my daily work in identifying hard to find vulnerabilities. If anyone think they fit in the “who it’s for” criteria I mentioned above, I strongly suggest taking AWAE, as you will definitely benefit from it.
Course Strengths:

    Covers advanced real-world vulnerabilities
    Professional teaching quality
    Excellent lab material
    Good for bug hunting
    Good OffSec support

Course Weaknesses:

    Needs more vulnerability discovery techniques
    Needs blackbox/whitebox mixture for vulnerability discovery
    Doesn’t go deep on source code review techniques
    Extra-miles did not cover all modules
    Needs more extra-miles in another language



Future Plans

I’ve been working on OSWE for quite some time, and have some ideas for my next step. Eventually, I’m thinking about going deep into OS/Binary exploitation, with: PACES, GXPN, OSCE, and OSEE. If anyone took OSCE and any of the others "GXPN, OSEE, PACES", I would love to hear your feedback on how to prioritize them.

But for my next step, I think I may have to step back a bit and improve my blue team skills, so I’m considering either Security+ or CISSP. I don't want to go very deep in the blue side, since it's not my main focus. If anyone has done any of them, I would also really appreciate the feedback.

For now, though, I'll work on ranking up to Guru or higher, here on HTB :)

Thanks a lot for taking the time to read my “detailed” OSWE review, you are awesome 

https://z-r0crypt.github.io/blog/2020/01/22/oswe/awae-preparation/

