package main

import (
	"crypto/tls"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"golang.org/x/sync/errgroup"

	"github.com/se7enkings/quic-go"
	"github.com/se7enkings/quic-go/http3"
	"github.com/se7enkings/quic-go/internal/protocol"
	"github.com/se7enkings/quic-go/interop/http09"
	"github.com/se7enkings/quic-go/interop/utils"
)

var errUnsupported = errors.New("unsupported test case")

var tlsConf *tls.Config

func main() {
	logFile, err := os.Create("/logs/log.txt")
	if err != nil {
		fmt.Printf("Could not create log file: %s\n", err.Error())
		os.Exit(1)
	}
	defer logFile.Close()
	log.SetOutput(logFile)

	keyLog, err := utils.GetSSLKeyLog()
	if err != nil {
		fmt.Printf("Could not create key log: %s\n", err.Error())
		os.Exit(1)
	}
	if keyLog != nil {
		defer keyLog.Close()
	}

	tlsConf = &tls.Config{
		InsecureSkipVerify: true,
		KeyLogWriter:       keyLog,
	}
	testcase := os.Getenv("TESTCASE")
	if err := runTestcase(testcase); err != nil {
		if err == errUnsupported {
			fmt.Printf("unsupported test case: %s\n", testcase)
			os.Exit(127)
		}
		fmt.Printf("Downloading files failed: %s\n", err.Error())
		os.Exit(1)
	}
}

func runTestcase(testcase string) error {
	flag.Parse()
	urls := flag.Args()

	getLogWriter, err := utils.GetQLOGWriter()
	if err != nil {
		return err
	}
	quicConf := &quic.Config{GetLogWriter: getLogWriter}

	if testcase == "http3" {
		r := &http3.RoundTripper{
			TLSClientConfig: tlsConf,
			QuicConfig:      quicConf,
		}
		defer r.Close()
		return downloadFiles(r, urls, false)
	}

	r := &http09.RoundTripper{
		TLSClientConfig: tlsConf,
		QuicConfig:      quicConf,
	}
	defer r.Close()

	switch testcase {
	case "handshake", "transfer", "retry":
	case "multiconnect":
		return runMultiConnectTest(r, urls)
	case "versionnegotiation":
		return runVersionNegotiationTest(r, urls)
	case "resumption":
		return runResumptionTest(r, urls, false)
	case "zerortt":
		return runResumptionTest(r, urls, true)
	default:
		return errUnsupported
	}

	return downloadFiles(r, urls, false)
}

func runVersionNegotiationTest(r *http09.RoundTripper, urls []string) error {
	if len(urls) != 1 {
		return errors.New("expected at least 2 URLs")
	}
	protocol.SupportedVersions = []protocol.VersionNumber{0x1a2a3a4a}
	err := downloadFile(r, urls[0], false)
	if err == nil {
		return errors.New("expected version negotiation to fail")
	}
	if !strings.Contains(err.Error(), "No compatible QUIC version found") {
		return fmt.Errorf("expect version negotiation error, got: %s", err.Error())
	}
	return nil
}

func runMultiConnectTest(r *http09.RoundTripper, urls []string) error {
	for _, url := range urls {
		if err := downloadFile(r, url, false); err != nil {
			return err
		}
		if err := r.Close(); err != nil {
			return err
		}
	}
	return nil
}

func runResumptionTest(r *http09.RoundTripper, urls []string, use0RTT bool) error {
	if len(urls) < 2 {
		return errors.New("expected at least 2 URLs")
	}

	tlsConf.ClientSessionCache = tls.NewLRUClientSessionCache(1)

	// do the first transfer
	if err := downloadFiles(r, urls[:1], false); err != nil {
		return err
	}
	r.Close()

	// reestablish the connection, using the session ticket that the server (hopefully provided)
	defer r.Close()
	return downloadFiles(r, urls[1:], use0RTT)
}

func downloadFiles(cl http.RoundTripper, urls []string, use0RTT bool) error {
	var g errgroup.Group
	for _, u := range urls {
		url := u
		g.Go(func() error {
			return downloadFile(cl, url, use0RTT)
		})
	}
	return g.Wait()
}

func downloadFile(cl http.RoundTripper, url string, use0RTT bool) error {
	method := http.MethodGet
	if use0RTT {
		method = http09.MethodGet0RTT
	}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return err
	}
	rsp, err := cl.RoundTrip(req)
	if err != nil {
		return err
	}
	defer rsp.Body.Close()

	file, err := os.Create("/downloads" + req.URL.Path)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = io.Copy(file, rsp.Body)
	return err
}
