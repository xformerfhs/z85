# z85

A package for Z85 encoding.

[![Go Report Card](https://goreportcard.com/badge/github.com/xformerfhs/z85)](https://goreportcard.com/report/github.com/xformerfhs/z85)
[![License](https://img.shields.io/github/license/xformerfhs/hashvalue)](https://github.com/xformerfhs/hashvalue/blob/main/LICENSE)

## Introduction

When using programs, binary data must be stored or transported in systems that can only handle printable characters.
A printable representation of the binary data must therefore be found.

There are several well-known encodings, like [hex](https://en.wikipedia.org/wiki/Hexadecimal) (aka Base16), [Base32](https://en.wikipedia.org/wiki/Base32), or [Base64](https://en.wikipedia.org/wiki/Base64).
While they are easy to use, they have several problems.

This package implements [Z85](https://rfc.zeromq.org/spec/32/) encoding as specified in the ZeroMQ universal messaging library.
This encoding processes 4 bytes at a time and produces 5 characters from them.
Therefore, it can only be used with binary data which has a byte length which is a multiple of 4.

> [!IMPORTANT]
> Data that does not have a length which is a multiple of 4 cannot be encoded.
> It is the duty of the calling application to pad such data and handle the padding and unpadding.

A benchmark shows that Z85 encoding needs about 90% more CPU than Base64 and 67% more CPU than Base32.
The reason for this is that Z85 uses division or scaled multiplication while the Base32 and Base64 algorithms need only bit shifting, bit masking and bit or-ing which are a lot faster than division.

## Functions

The library offers two public functions:

| Command          | Meaning                                                                                                 |
|------------------|---------------------------------------------------------------------------------------------------------|
| `Decode`         | Decodes a Z85 encoded string.                                                                           |
| `Encode`         | Encodes a byte slice in Z85.                                                                            |

## Errors

The functions may return the following named errors:

| Error               | Meaning                                                                    |
|---------------------|----------------------------------------------------------------------------|
| `ErrInvalidByte`    | An encoded string contains a byte that is not a valid Z85 encoding.        |
| `ErrInvalidLength`  | The supplied data has an invalid length.                                   |

There are two functions that can test a returned error:

| Function             | Meaning                                                   |
|----------------------|-----------------------------------------------------------|
| `IsErrInvalidByte`   | Reports whether the error is an `ErrInvalidByte` error.   |
| `IsErrInvalidLength` | Reports whether the error is an `ErrInvalidLength` error. |

## Examples

An example for encoding is this:

```go
...
   data := []byte{0x5b, 0x9b, 0xf1, 0x40}
   encoded, err := z85.Encode(data)
   if err != nil {
      return err
   }
   // encoded will contain the string "tBUdF".
...
```

An example for decoding is this:

```go
...
   encoded := `khpb9ZWRsz`
   data, err := z85.Decode(encoded)
   if err != nil {
      return err
   }
   // data will contain the bytes [0x3e, 0xdc, 0x70, 0xd2, 0xbf, 0xf1, 0x01, 0x1b].
...
```

## Contact

Frank Schwab ([Mail](mailto:github.sfdhi@slmails.com "Mail"))

## License

This source code is published under the [Apache License V2](https://www.apache.org/licenses/LICENSE-2.0.txt).
