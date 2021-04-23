package sctpreceiver

import (
	"flag"
	"log"
	"nRIC/internal"
	"nRIC/internal/logger"
	"nRIC/pkg/nrice2t/managers/notificationmanager"
	"nRIC/pkg/nrice2t/sctp"
	"net"
	"strings"
)

type SctpReceiver struct {
	logger    *logger.Logger
	nManager  *notificationmanager.NotificationManager
	wconn	   *sctp.SCTPSndRcvInfoWrappedConn
	//messenger sctpCgo.SctpMessenger
}
func NewSctpReceiver(logger *logger.Logger, nManager *notificationmanager.NotificationManager) *SctpReceiver {
	return &SctpReceiver{
		logger:    logger,
		nManager:  nManager,
		wconn:	nil,
		//messenger: messenger,
	}
}

func (r *SctpReceiver) GetWconn() *sctp.SCTPSndRcvInfoWrappedConn {
	return r.wconn
}


func (r *SctpReceiver) Handle(conn net.Conn, bufsize int) error {
	for {
		mbuf := make([]byte, bufsize+128) // add overhead of SCTPSndRcvInfoWrappedConn
		n, err := conn.Read(mbuf)
		if err != nil {
			log.Printf("read failed: %v", err)
			return err
		}
		log.Printf("read: %d", n)
/*
		n, err = conn.Write(mbuf[:n])
		if err != nil {
			log.Printf("write failed: %v", err)
			return err
		}
		log.Printf("write: %d", n)*/

		mbuf2 := mbuf[32:n]
		//r.logger.Infof("mbuf 0-31 : %#v",mbuf[:32])
		//fmt.Printf("mbuf 0-31 : %x\n",mbuf[:32])
		//r.logger.Infof("#SctpReceiver.ListenAndHandle - Going to handle received message: %#v", mbuf2)

		_ = r.nManager.HandleMessage(mbuf2,conn)
	}
}

func (r *SctpReceiver) CreateConnection() (*sctp.SCTPListener) {
	var ip = flag.String("ip", internal.Nrice2tHost, "")
	var port = flag.Int("port", 8800, "")
	flag.Parse()

	ips := []net.IPAddr{}

	for _, i := range strings.Split(*ip, ",") {
		if a, err := net.ResolveIPAddr("ip", i); err == nil {
			log.Printf("\nResolved address '%s' to %s\n", i, a)
			ips = append(ips, *a)
		} else {
			log.Printf("Error resolving address '%s': %v\n", i, err)
		}
	}
	addr := &sctp.SCTPAddr{
		IPAddrs: ips,
		Port:    *port,
	}
	log.Printf("raw addr: %+v\n", addr.ToRawSockAddrBuf())

	ln, err := sctp.ListenSCTP("sctp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("Listen on  %s\n", ln.Addr())
	return ln
}

func (r *SctpReceiver) ListenAndHandle(ln *sctp.SCTPListener) (error) {
	var server = flag.Bool("server", true, "")
	var bufsize = flag.Int("bufsize", 4096, "")
	var sndbuf = flag.Int("sndbuf", 0, "")
	var rcvbuf = flag.Int("rcvbuf", 0, "")

	flag.Parse()
	log.Printf("ListenAndHandle: %s\n", ln.Addr())

	if *server {
		for {
			conn, err := ln.Accept()
			if err != nil {
				log.Fatalf("failed to accept: %v\n", err)
				return err
			}
			log.Printf("Accepted Connection from RemoteAddr: %s\n", conn.RemoteAddr())
			wconn := sctp.NewSCTPSndRcvInfoWrappedConn(conn.(*sctp.SCTPConn))

			//TBD: more wconn
			r.wconn = wconn

			if *sndbuf != 0 {
				err = wconn.SetWriteBuffer(*sndbuf)
				if err != nil {
					log.Fatalf("failed to set write buf: %v", err)
					return err
				}
			}
			if *rcvbuf != 0 {
				err = wconn.SetReadBuffer(*rcvbuf)
				if err != nil {
					log.Fatalf("failed to set read buf: %v", err)
					return err
				}
			}
			*sndbuf, err = wconn.GetWriteBuffer()
			if err != nil {
				log.Fatalf("failed to get write buf: %v", err)
				return err
			}
			*rcvbuf, err = wconn.GetReadBuffer()
			if err != nil {
				log.Fatalf("failed to get read buf: %v", err)
			}


			log.Printf("SndBufSize: %d, RcvBufSize: %d", *sndbuf, *rcvbuf)


			go r.Handle(wconn, *bufsize)
		}

	}
	return nil
}
