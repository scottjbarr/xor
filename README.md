# XOR Cypher

A Go implementation of the XOR Cypher.

I put this little bit of code together after once again hearing the XOR
crypher being spoken about. I decided to do a quick refresher. In a world
where PGP exists, you sometimes forget how good the simpler things can be.

Given the same key, passing plaintext to the function results in ciphertext,
and passing ciphertext to the function results in plaintext.


## Why?

The XOR Cipher is fast! There are implementations in any language you could
care to use.


## Usage

The example here shows encrypting and decrypting a string.

```
# Set the key. Use something secure, and of reasonable length.
$ export KEY=yoursecretthisisnotreallysecure

# encrypt and then decrypt
$ echo "Encrypting with xor is simple and fast." | xor-cypher | xor-cypher
Encrypting with xor is simple and fast.
```

The cipher can also be used as from code. See the [tests](xor_test.go) for
an example.


## Benchmarks

It is quite fast to encrypt and decrypt. The benchmarks will differ for you
on your machine, of course.

165ns/op is slightly better than 6M ops/sec.

```
$ go test -bench .
goos: linux
goarch: amd64
pkg: github.com/scottjbarr/xor
BenchmarkEncryptDecrypt-12          	10000000	       165 ns/op	      32 B/op	       1 allocs/op
```

## References

- https://en.wikipedia.org/wiki/XOR_cipher


## Copyright

The MIT License (MIT)

Copyright (c) 2019 Scott Barr

See the [LICENSE](LICENSE.md).
