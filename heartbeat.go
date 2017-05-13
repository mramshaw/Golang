// Simple proof of concept for a heartbeat function.
//
// With a little bit of effort, could be modified to
// be a Stop-Loss alert (for instance if a stock quote
// falls below a specific value). Could also be modified
// for load-testing or benchmarking purposes.
//
// Possible enhancements might be to simulate multiple
// concurrent visitors or to specify a specific HTTP
// version (0.9, 1.0, 1.1 or 2) to verify compliance.
//
// The default URL is http://localhost
//
// The default polling time is every 5 minutes
//
// The default polling timeout is 10 seconds
// 
// The default variance (response size or time) is 5 percent
//
// When the first fetch completes, the returned length
// will be saved, and lengths that differ from this by
// more or less than the specified deviation on later
// fetches will generate alerts.
//
// In a similiar manner, significant deviations from
// the initial fetch time will also generate an alert.
//
// Fetches that do not complete within the timeout period
// will generate another type of alert.
//
// The first DNS lookup is ignored for variance purposes.
// This is because the very first lookup may well take a
// significant amount of time compared to subsequent
// lookups when this value will be served from cache.
//
// [With HTTP/2 the DNS lookup can be a significant part
//    of the overall round-trip time.]
//
// However any subsequent DNS lookups as a result of HTTP
// redirects will be considered for variance purposes.
//
// REQUIRES the net/http/httptrace package from Go 1.7
//
// ---------------------------------------------------
//
// TEST PLAN
//
// 1) Verify command line arguments are defaulted correctly
//
// 2) 'host unreachable'
//
//    Make sure web server is down
//
//     ./heartbeat
//     ./heartbeat https://localhost
//
//    Both tests should fail 'host unreachable'
//
// 3) 'host reachable' / x509 self-signed certificate error
//
//    Make sure web server is running (with HTTPS configured)
//
//     ./heartbeat
//     ./heartbeat https://localhost
//
//    First test should succeed (Ctrl-C to kill)
//    Second test should fail (x509 error due to self-signed certificate)
//
// 4) Length Variance
//
//     ./heartbeat http://localhost/test2.php 1 1 1 verbose
//
//     Wait until the variance range is displayed, 
//       then add or subtract characters from the
//       test2.php web page (or any other web page,
//       as appropriate) in order to trigger a
//       length variance.
//
//     Wait a minute or so, verify Length variance warning message
//
//     Wait another minute, verify Length variance warning message resets
//
// 5) Time Variance
//
//     Set sleep time in 'timer.php' to 5 (seconds)
//
//         [timer.php should take 5 seconds (plus overhead) to run]
//
//     ./heartbeat http://localhost/timer.php 1 6 1
//
//         [timeout is set to 6 seconds to allow timer.php to complete]
//
//     After first response, set sleep time in 'timer.php' to 4 (seconds)
//
//     Wait a minute or so, verify Time variance warning message
//
//     Wait another minute, verify Time variance warning message resets
//
// 6) Timeout
//
//     Set sleep time in 'timer.php' to 1 (seconds)
//
//         [timer.php should take 1 second (plus overhead) to run]
//
//     ./heartbeat http://localhost/timer.php 1 2 1
//
//         [timeout is set to 2 seconds to allow timer.php to complete]
//
//     After first response, set sleep time in 'timer.php' to 2 (seconds)
//
//     Wait a minute or so, verify Timeout warning message
//
//     Wait another minute, verify Timeout warning message repeats
//
// 7) Ignore first DNS lookup
//
//     ./heartbeat https://oracle.com 5 10 5 verbose
//
//         [polling time is set to 5 minutes to allow connections to expire]
//
//     Wait at least 10 minutes, verify that first DNS lookup is correctly ignored
//
//     [Update: DNS times seem to fluctuate wildly, 
//              perhaps best to use 'nslookup' to
//              determine FQDN and skip DNS completely.]
//
//     ./heartbeat https://137.254.120.50 5 10 5 verbose
//
//     [Update: Even using a FQDN seems to trigger DNS lookups.]
//
// ---------------------------------------------------
//
// @ Martin Ramshaw, April 2017 (mramshaw@alumni.concordia.ca)

package main

import (
    "fmt"
    "io"
    "io/ioutil"
    "net/http"
    "net/http/httptrace"
    "os"
    "runtime"
    "strconv"
    "time"
)

const (
    version         = "1.0"
    defaultURL      = "http://localhost"
    defaultPoll     =  5                   // every 5 minutes
    defaultTimeout  = 10                   // 10 seconds
    defaultVariance =  5                   // 5 percent (%)
)

var (
    wCount    uint64
    wLo       uint64
    wHi       uint64
    rTrip     int64
    rTime     int64
    rLo       int64
    rHi       int64
    verbose   bool
)

// ===============================================================

// transport is an http.RoundTripper that keeps track of the fetch
//   request and implements hooks to report HTTP tracing events.

type transport struct {
    current *http.Request
}

// Wraps http.DefaultTransport.RoundTrip to keep track of the current fetch.
func (trans *transport) RoundTrip(req *http.Request) (*http.Response, error) {

    trans.current = req
    return http.DefaultTransport.RoundTrip(req)
}

// Shows whether the connection has been used previously.
func (trans *transport) GotConn(info httptrace.GotConnInfo) {

    if verbose {
        fmt.Printf("Connection reused for '%v' ? %v - Was idle ? %v\n", trans.current.URL, info.Reused, info.WasIdle)
    }
}

// Indicates that a "100 Continue" message has been received - which is kind of a 'Normal' error message.
// Shouldn't show up often, probably indicates configuration issues (packet fragmentation or something).
// It should be handled transparently, but probably worth reporting in any case.
func (trans *transport) Got100Continue() {

    if verbose {
        fmt.Printf("'100 Continue' message received\n")
    }
}

// ===============================================================

func main() {

    fmt.Printf("\n== heartbeat %s (runtime: %s) == Ctrl-C to quit!\n", version, runtime.Version())
    fmt.Printf("\n")

    url      := defaultURL
    poll     := defaultPoll
    timeout  := defaultTimeout
    variance := defaultVariance

    if argsLength := len(os.Args); argsLength < 2 {
        fmt.Printf("   ... defaulting URL to %s\n",                     defaultURL)
        fmt.Printf("   ... defaulting Polling period to %v minutes\n",  defaultPoll)
        fmt.Printf("   ... defaulting Timeout period to %v seconds\n",  defaultTimeout)
        fmt.Printf("   ... defaulting Variance       to %v percent\n",  defaultVariance)
        fmt.Printf("\n")
    } else if argsLength < 3 {
        url  = os.Args[1]
        fmt.Printf("   ... defaulting Polling period to %v minutes\n",  defaultPoll)
        fmt.Printf("   ... defaulting Timeout period to %v seconds\n",  defaultTimeout)
        fmt.Printf("   ... defaulting Variance       to %v percent\n",  defaultVariance)
        fmt.Printf("\n")
    } else if argsLength < 4 {
        url = os.Args[1]
        poll = parseArg(os.Args[2], "polling period")
        fmt.Printf("   ... defaulting Timeout period to %v seconds\n",  defaultTimeout)
        fmt.Printf("   ... defaulting variance       to %v percent\n",  defaultVariance)
        fmt.Printf("\n")
    } else if argsLength < 5 {
        url = os.Args[1]
        poll = parseArg(os.Args[2], "polling period")
        timeout = parseArg(os.Args[3], "timeout period")
        fmt.Printf("   ... defaulting variance       to %v percent\n",  defaultVariance)
        fmt.Printf("\n")
    } else {
        if argsLength == 6 {
            if os.Args[5] == "verbose" {
                verbose = true
            } else {
                fmt.Printf("Invalid verbose mode: '%s'\n\n", os.Args[5])
                usage()
                os.Exit(2)
            }
        }
        url      = os.Args[1]
        poll     = parseArg(os.Args[2], "polling period")
        timeout  = parseArg(os.Args[3], "timeout period")
        variance = parseArg(os.Args[4], "variance")
    }

    fmt.Printf("Polling '%s' every %v minutes with a %v second timeout +/- %v percent variance\n", url, poll, timeout, variance)

    for {
        everLoop(url, poll, timeout, variance)			// Infinite loop, Ctrl-C to kill
    }
}

// Fetches the specified URL, redirecting as necessary, then sleeps.
// Variances will generate messages as will a response greater than
// the specified timeout period.
func everLoop(url string, p, to, v int) {

    defer time.Sleep(time.Duration(p) * time.Minute)

    t := &transport{}

    tStart := time.Now()
    if verbose {
        fmt.Printf("%s Starting HTTP Get now ...\n", tStart)
    }

    req, _ := http.NewRequest("GET", url, nil)

    var dnsTime,      connectTime          time.Time
    var totalDNStime, totalConnectionTime  time.Duration
    var firstDNStime                       time.Duration

    trace := &httptrace.ClientTrace {
        DNSStart:        func(sinfo httptrace.DNSStartInfo) {
            if verbose {
                dnsTime = time.Now()
                fmt.Printf("DNS lookup started for '%v'\n", sinfo.Host); // doesn't seem to reflect redirects
            }
        },
        DNSDone:         func(_ httptrace.DNSDoneInfo)  {
            dTime := time.Now().Sub(dnsTime)
            totalDNStime += dTime
            if firstDNStime == 0 {
                firstDNStime = dTime
            }
            if verbose {
                fmt.Printf("DNS lookup took: %d ms\n", int(time.Duration(dTime) / time.Millisecond))
            }
        },
        ConnectStart:    func(_, _ string) {
            connectTime = time.Now() // there can be many connections
        },
        ConnectDone:     func(net, addr string, err error) {
            if err != nil {
                fmt.Printf("Unable to connect to host '%v', net '%v':\n%v\n", addr, net, err)
                os.Exit(-1)
            } else {
                cTime := time.Now().Sub(connectTime)
                totalConnectionTime += cTime
                if verbose {
                    fmt.Printf("Connection:  %d ms\n", int(time.Duration(cTime) / time.Millisecond))
                }
            }
        },
        GotConn:         t.GotConn,
        Got100Continue:  t.Got100Continue,
    }
    req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))

    timeout := time.Duration(time.Duration(to) * time.Second)
    client  := &http.Client {
        Transport: t,
        Timeout:   timeout,
    }

    resp, err := client.Do(req)
    if err != nil {
        fmt.Printf("%s WARNING WARNING probable Timeout on request (use verbose option for more details)\n", time.Now())
        if verbose {
            fmt.Printf("Error on request:\n%v\n", err)
        }
        return
    }

    if verbose {
        fmt.Printf("Total DNS lookup time was: %v ms (First DNS lookup time was: %d ms)\n",
                   int(totalDNStime / time.Millisecond), int(firstDNStime / time.Millisecond))
        fmt.Printf("Total connection time was: %v ms\n", int(totalConnectionTime / time.Millisecond))
    }

    _, berr := verifyResponseBody(req, resp, v)
    if berr != nil {
        fmt.Printf("Error on response:\n%v\n", berr)
        return
    }

    elapsed := time.Since(tStart)
    varTime := elapsed - firstDNStime
    elapsed /= time.Millisecond  // reframe in milliseconds
    varTime /= time.Millisecond  // reframe in milliseconds

    tripTime := int64(elapsed)
    respTime := int64(varTime)
    respLo   := float64(varTime) * (1.0 - (float64(v) / 100.0))
    respHi   := float64(varTime) * (1.0 + (float64(v) / 100.0))
    if verbose {
        fmt.Printf("round trip took %v ms; %v ms ignoring first DNS, a %v%% variance is ~ %v - %v ms\n", tripTime, respTime, v, respLo, respHi)
        fmt.Printf("%s %s%s%d.%d %s\n", time.Now(), " - HTTP", "/", resp.ProtoMajor, resp.ProtoMinor, resp.Status)
    }
    if rTrip == 0 {
        rTrip = tripTime
        rTime = respTime
        rHi   = int64(respHi)
        rLo   = int64(respLo)
    } else {
        if respTime < rLo || respTime > rHi {
            fmt.Printf("%s WARNING WARNING previously %v ms, now %v ms\n", time.Now(), rTrip, tripTime)
            rTrip = tripTime
            rTime = respTime
            rLo   = int64(respLo)
            rHi   = int64(respHi)
        }
    }

    return
}

func verifyResponseBody(req *http.Request, resp *http.Response, v int) (written int64, err error) {

    defer resp.Body.Close()

    if isRedirected(resp) {
        if verbose {
            fmt.Printf("%s request was redirected with code %d\n", time.Now(), resp.StatusCode)
        }
        return		// if this is a redirect, don't care about body
    }

    w := ioutil.Discard

    byteCount, err := io.Copy(w, resp.Body)
    if err != nil {
        fmt.Printf("Failed to read response body; error:\n%s\n", err)
        return byteCount, err
    }

    bc := uint64(byteCount)
    lo := float64(bc) * (1.0 - (float64(v) / 100.0))
    hi := float64(bc) * (1.0 + (float64(v) / 100.0))
    if verbose {
        fmt.Printf("response body had %v bytes, a %v%% variance is ~ %v - %v\n", bc, v, lo, hi)
    }
    if wCount == 0 {
        wCount = bc
        wLo    = uint64(lo)
        wHi    = uint64(hi)
    } else {
        if bc < wLo || bc > wHi {
            fmt.Printf("%s WARNING WARNING previously %v bytes, now %v bytes\n", time.Now(), wCount, bc)
            wCount = bc
            wLo    = uint64(lo)
            wHi    = uint64(hi)
        }
    }

    return byteCount, nil
}

func isRedirected(resp *http.Response) bool {

    return resp.StatusCode > 299 && resp.StatusCode < 400
}

func parseArg(s, desc string) int {

    so, err := strconv.Atoi(s)
    if err != nil {
        fmt.Printf("Invalid %s: '%s'\n\n", desc, s)
        usage()
        os.Exit(2)
    }
    return so
}

func usage() {

    fmt.Printf("Usage is:\n")
    fmt.Printf("\n")
    fmt.Printf("    ./heartbeat URL poll timeout variance verbose\n")
    fmt.Printf("\n")
    fmt.Printf("      URL      [optional] website to heartbeat\n")
    fmt.Printf("                          defaults to  http://localhost\n")
    fmt.Printf("      poll     [optional] polling time in minutes\n")
    fmt.Printf("                          default value is 5\n")
    fmt.Printf("      timeout  [optional] timeout period in seconds\n")
    fmt.Printf("                          default value is 10\n")
    fmt.Printf("      variance [optional] time or response variance (percent)\n")
    fmt.Printf("                          default value is 5\n")
    fmt.Printf("      verbose  [optional] verbose mode\n")
    fmt.Printf("                          default value is Off\n")
    fmt.Printf("\n")
}
