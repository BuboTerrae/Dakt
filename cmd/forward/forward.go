package forward

import (
	"log"
	"net/http"

	"github.com/elazarl/goproxy"
)

func Forward(port string) {
	// log.Printf("Starting proxy server on port:%s\n", port)
	// http.HandleFunc("/", handler)
	// log.Printf("Proxy handler registered. Listening on :%s", port)
	// log.Fatal(http.ListenAndServe(":"+port, nil))
	
	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = true

	log.Printf("Starting proxy on port: %s", port)
	log.Fatal(http.ListenAndServe(":"+port, proxy))
}

// func handler(w http.ResponseWriter, r *http.Request) {
// 	log.Printf("RECEIVED REQUEST: %s %s %s (Host: %s)", r.Method, r.URL, r.Proto, r.Host)

// 	if r.Method == http.MethodConnect {
// 		log.Printf("→ CONNECT tunnel request to %s", r.Host)
// 		handleConnect(w, r)
// 		return
// 	}

// 	outReq, err := http.NewRequest(r.Method, r.URL.String(), r.Body)
// 	if err != nil {
// 		log.Printf("NewRequest error: %v", err)
// 		http.Error(w, err.Error(), http.StatusBadGateway)
// 		return
// 	}

// 	for k, vv := range r.Header {
// 		for _, v := range vv {
// 			outReq.Header.Add(k, v)
// 		}
// 	}

// 	removeHopHeaders(outReq.Header)

// 	resp, err := http.DefaultClient.Do(outReq)
// 	if err != nil {
// 		log.Printf("Forward err: %v", err)
// 		http.Error(w, err.Error(), http.StatusBadGateway)
// 		return
// 	}
// 	defer resp.Body.Close()

// 	copyHeader(w.Header(), resp.Header)
// 	w.WriteHeader(resp.StatusCode)
// 	io.Copy(w, resp.Body)
// }

// func handleConnect(w http.ResponseWriter, r *http.Request) {
// 	log.Printf("CONNECT to %s", r.Host)

// 	dest, err := net.DialTimeout("tcp", r.Host, 30*time.Second)
// 	if err != nil {
// 		log.Printf("Dial failed %s: %v", r.Host, err)
// 		http.Error(w, err.Error(), http.StatusBadGateway)
// 		return
// 	}
// 	defer dest.Close()

// 	// Critical part for browsers
// 	w.WriteHeader(http.StatusOK)
// 	if flusher, ok := w.(http.Flusher); ok {
// 		flusher.Flush()
// 	}

// 	// Hijack
// 	hijacker, ok := w.(http.Hijacker)
// 	if !ok {
// 		log.Println("Hijack not supported")
// 		http.Error(w, "Hijacking not supported", http.StatusInternalServerError)
// 		return
// 	}

// 	clientConn, _, err := hijacker.Hijack()
// 	if err != nil {
// 		log.Printf("Hijack error: %v", err)
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	defer clientConn.Close()

// 	log.Printf("✅ Tunnel open for %s", r.Host)

// 	go func() {
// 		_, err := io.Copy(dest, clientConn)
// 		if err != nil && err != io.EOF {
// 			log.Printf("cliet=>dest error: %v", err)
// 		}
// 	}()

// 	_, err = io.Copy(clientConn, dest)
// 	if err != nil && err != io.EOF {
// 		log.Printf("dest=>client error: %v", err)
// 	}

// 	log.Printf("Tunnel closed for %s", r.Host)
// }

// func copyHeader(dst, src http.Header) {
// 	for k, vv := range src {
// 		for _, v := range vv {
// 			dst.Add(k, v)
// 		}
// 	}
// }

// func removeHopHeaders(h http.Header) {
// 	for _, k := range []string{
// 		"Connection", "Proxy-Connection", "Keep-Alive",
// 		"Proxy-Authenticate", "Proxy-Authorization",
// 		"Te", "Trailers", "Transfer-Encoding", "Upgrade",
// 	} {
// 		h.Del(k)
// 	}
// }
