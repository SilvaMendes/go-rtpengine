# go-rtpengine

**go-rtpengine** is a Go library that provides bindings and utilities to interact with the [RTPengine](https://github.com/sipwise/rtpengine), a media proxy used in VoIP infrastructures such as Kamailio/Opensips. This project aims to simplify the integration of RTPengine's NG protocol into Go-based applications.

## Installation
``
go get github.com/SilvaMendes/go-rtpengine
``

## Features

- Native Go bindings for RTPengine NG protocol
- Support for key RTPengine commands: `offer`, `answer`, `delete`, `query`, `ping`, etc.
- Encoding and decoding of messages using Bencode
- Support for advanced RTPengine features:
  - DTLS and SDES configuration
  - SRTP crypto suites
  - ICE and OSRTP negotiation
  - Codec filtering and transcoding flags
  - RTCP multiplexing options
  - Media recording control

## Usage

Import the package and use the `Engine` struct to manage RTPengine connections and send commands.

``
import "github.com/SilvaMendes/go-rtpengine"
``

### Example: 
```
engine := rtpengine.NewEngine("127.0.0.1", 2223, "udp")
response, err := engine.Offer(request)
if err != nil {
    log.Fatal(err)
}
fmt.Println(response)
```

## License
``
This project is licensed under the Apache-2.0 License.
``
## 
``
Feel free to fork the repository and submit pull requests. Contributions are welcome!
``