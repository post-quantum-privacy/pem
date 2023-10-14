# PEM

PEM is a standards complient encoder for the PEM format. This libary enables the streaming of PEM encodes, something which is not supported within Go's standard libary. This enables the encoding and decoding of large PEM files without exhusting system memory.

## The Problem This Resolves
Go's standard `encoding/pem` package does not allow the streaming of files to be encoded or decoded, thus making large files exhust the device's RAM.
**This libarary allows data to be encoded and decoded in a stream** (io.Reader & io.Writer), without commiting large buffers to memory, or worse yet,
the entire file. 

## Import

```
go get github.com/post-quantum-privacy/pem
```