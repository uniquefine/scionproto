#include <stdlib.h>
#include <unistd.h>
#include <sys/types.h>
#include <netinet/in.h>
#include <arpa/inet.h>
#include <string.h>
#include <stdio.h>
#include <ncurses.h>
#include <signal.h>
#include <Python.h>
#include "SCIONSocket.h"
#include "SHA1.h"

#define BUFSIZE 1024

void die(int signum)
{
    endwin();
    exit(0);
}

int main()
{
    initscr();
    cbreak();
    noecho();

    signal(SIGINT, die);
    signal(SIGTERM, die);

    uint8_t hash[20];
    CSHA1 sha;
    sha.HashFile("randomfile");
    sha.Final();
    sha.GetHash(hash);

    SCIONAddr *addrs[1];
    SCIONAddr saddr;
    saddr.ISD_AD(1, 5);
    saddr.host.addrLen = 4;
    in_addr_t in = inet_addr("192.33.93.195");
    memcpy(saddr.host.addr, &in, 4);
    addrs[0] = &saddr;
    SCIONSocket s(SCION_PROTO_SDAMP, addrs, 1, 0, 8080);
    s.send((uint8_t *)hash, 20);
    printw("hash sent\n");
    refresh();
    FILE *fp = fopen("randomfile", "rb");
    char fbuf[BUFSIZE];
    int len = 0;
    while ((len = fread(fbuf, 1, BUFSIZE, fp)) > 0) {
        s.send((uint8_t *)fbuf, len);

        SCIONStats stats;
        memset(&stats, 0, sizeof(stats));
        s.getStats(&stats);
        clear();
        move(0, 0);
        printw("Current stats\n");
        for (int i = 0; i < MAX_TOTAL_PATHS; i++) {
            if (!(stats.exists[i]))
                continue;
            printw(">> Path %d <<\n", i);
            printw("Received: %d\n", stats.receivedPackets[i]);
            printw("Sent:     %d\n", stats.sentPackets[i]);
            printw("ACKed:    %d\n", stats.ackedPackets[i]);
            printw("RTT:      %d ms\n", stats.rtts[i] / 1000);
            printw("Loss:     %.2f\n", stats.lossRates[i]);
        }
        printw("\n");
        refresh();
    }
    printw("done sending file\n");
    refresh();
    fclose(fp);
    getch();
    endwin();
    exit(0);
}
