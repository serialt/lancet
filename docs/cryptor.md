# Cryptor
Package cryptor contains some functions for data encryption and decryption. Support base64, md5, hmac, aes, des, rsa.

<div STYLE="page-break-after: always;"></div>

## Source:

- [https://github.com/duke-git/lancet/blob/main/cryptor/basic.go](https://github.com/duke-git/lancet/blob/main/cryptor/basic.go)
- [https://github.com/duke-git/lancet/blob/main/cryptor/encrypt.go](https://github.com/duke-git/lancet/blob/main/cryptor/encrypt.go)



<div STYLE="page-break-after: always;"></div>

## Usage:
```go
import (
    "github.com/serialt/lancet/cryptor"
)
```

<div STYLE="page-break-after: always;"></div>

## Index

- [AesEcbEncrypt](#AesEcbEncrypt)
- [AesEcbDecrypt](#AesEcbDecrypt)
- [AesCbcEncrypt](#AesCbcEncrypt)
- [AesCbcDecrypt](#AesCbcDecrypt)
- [AesCtrCrypt](#AesCtrCrypt)
- [AesCfbEncrypt](#AesCfbEncrypt)
- [AesCfbDecrypt](#AesCfbDecrypt)
- [AesOfbEncrypt](#AesOfbEncrypt)
- [AesOfbDecrypt](#AesOfbDecrypt)
- [Base64StdEncode](#Base64StdEncode)
- [Base64StdDecode](#Base64StdDecode)
- [DesEcbEncrypt](#DesEcbEncrypt)
- [DesEcbDecrypt](#DesEcbDecrypt)
- [DesCbcEncrypt](#DesCbcEncrypt)
- [DesCbcDecrypt](#DesCbcDecrypt)
- [DesCtrCrypt](#DesCtrCrypt)
- [DesCfbEncrypt](#DesCfbEncrypt)
- [DesCfbDecrypt](#DesCfbDecrypt)
- [DesOfbEncrypt](#DesOfbEncrypt)
- [DesOfbDecrypt](#DesOfbDecrypt)
- [HmacMd5](#HmacMd5)
- [HmacSha1](#HmacSha1)
- [HmacSha256](#HmacSha256)
- [HmacSha512](#HmacSha512)
- [Md5String](#Md5String)
- [Md5File](#Md5File)
- [Sha1](#Sha1)
- [Sha256](#Sha256)
- [Sha512](#Sha512)
- [GenerateRsaKey](#GenerateRsaKey)
- [RsaEncrypt](#RsaEncrypt)
- [RsaDecrypt](#RsaDecrypt)


<div STYLE="page-break-after: always;"></div>

## Documentation

### <span id="AesEcbEncrypt">AesEcbEncrypt</span>

<p>Encrypt data with key use AES ECB algorithm. Length of `key` param should be 16, 24 or 32.</p>

<b>Signature:</b>

```go
func AesEcbEncrypt(data, key []byte) []byte
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/serialt/lancet/cryptor"
)

func main() {
    data := "hello"
    key := "abcdefghijklmnop"

    encrypted := cryptor.AesEcbEncrypt([]byte(data), []byte(key))
    decrypted := cryptor.AesEcbDecrypt(encrypted, []byte(key))

    fmt.Println(string(decrypted))

    // Output:
    // hello
}
```

### <span id="AesEcbDecrypt">AesEcbDecrypt</span>

<p>Decrypt data with key use AES ECB algorithm. Length of `key` param should be 16, 24 or 32.</p>

<b>Signature:</b>

```go
func AesEcbDecrypt(encrypted, key []byte) []byte
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/serialt/lancet/cryptor"
)

func main() {
    data := "hello"
    key := "abcdefghijklmnop"

    encrypted := cryptor.AesEcbEncrypt([]byte(data), []byte(key))
    decrypted := cryptor.AesEcbDecrypt(encrypted, []byte(key))

    fmt.Println(string(decrypted))

    // Output:
    // hello
}
```

### <span id="AesCbcEncrypt">AesCbcEncrypt</span>

<p>Encrypt data with key use AES CBC algorithm. Length of `key` param should be 16, 24 or 32.</p>

<b>Signature:</b>

```go
func AesCbcEncrypt(data, key []byte) []byte
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/serialt/lancet/cryptor"
)

func main() {
    data := "hello"
    key := "abcdefghijklmnop"

    encrypted := cryptor.AesCbcEncrypt([]byte(data), []byte(key))
    decrypted := cryptor.AesCbcDecrypt(encrypted, []byte(key))

    fmt.Println(string(decrypted))

    // Output:
    // hello
}
```

### <span id="AesCbcDecrypt">AesCbcDecrypt</span>

<p>Decrypt data with key use AES CBC algorithm. Length of `key` param should be 16, 24 or 32.</p>

<b>Signature:</b>

```go
func AesCbcDecrypt(encrypted, key []byte) []byte
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/serialt/lancet/cryptor"
)

func main() {
    data := "hello"
    key := "abcdefghijklmnop"

    encrypted := cryptor.AesCbcEncrypt([]byte(data), []byte(key))
    decrypted := cryptor.AesCbcDecrypt(encrypted, []byte(key))

    fmt.Println(string(decrypted))

    // Output:
    // hello
}
```

### <span id="AesCtrCrypt">AesCtrCrypt</span>

<p>Encrypt or decrypt data with key use AES CTR algorithm. Length of `key` param should be 16, 24 or 32.</p>

<b>Signature:</b>

```go
func AesCtrCrypt(data, key []byte) []byte
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/serialt/lancet/cryptor"
)

func main() {
    data := "hello"
    key := "abcdefghijklmnop"

    encrypted := cryptor.AesCtrCrypt([]byte(data), []byte(key))
    decrypted := cryptor.AesCtrCrypt(encrypted, []byte(key))

    fmt.Println(string(decrypted))

    // Output:
    // hello
}
```

### <span id="AesCfbEncrypt">AesCfbEncrypt</span>

<p>Encrypt data with key use AES CFB algorithm. Length of `key` param should be 16, 24 or 32.</p>

<b>Signature:</b>

```go
func AesCfbEncrypt(data, key []byte) []byte
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/serialt/lancet/cryptor"
)

func main() {
    data := "hello"
    key := "abcdefghijklmnop"

    encrypted := cryptor.AesCfbEncrypt([]byte(data), []byte(key))
    decrypted := cryptor.AesCfbDecrypt(encrypted, []byte(key))

    fmt.Println(string(decrypted))

    // Output:
    // hello
}
```

### <span id="AesCfbDecrypt">AesCfbDecrypt</span>

<p>Decrypt data with key use AES CBC algorithm. Length of `key` param should be 16, 24 or 32.</p>

<b>Signature:</b>

```go
func AesCfbDecrypt(encrypted, key []byte) []byte
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/serialt/lancet/cryptor"
)

func main() {
    data := "hello"
    key := "abcdefghijklmnop"

    encrypted := cryptor.AesCfbEncrypt([]byte(data), []byte(key))
    decrypted := cryptor.AesCfbDecrypt(encrypted, []byte(key))

    fmt.Println(string(decrypted))

    // Output:
    // hello
}
```

### <span id="AesOfbEncrypt">AesOfbEncrypt</span>

<p>Enecrypt data with key use AES OFB algorithm. Length of `key` param should be 16, 24 or 32.</p>

<b>Signature:</b>

```go
func AesOfbEncrypt(data, key []byte) []byte
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/serialt/lancet/cryptor"
)

func main() {
    data := "hello"
    key := "abcdefghijklmnop"

    encrypted := cryptor.AesOfbEncrypt([]byte(data), []byte(key))
    decrypted := cryptor.AesCfbDecrypt(encrypted, []byte(key))

    fmt.Println(string(decrypted))

    // Output:
    // hello
}
```
### <span id="AesCfbDecrypt">AesOfbDecrypt</span>

<p>Decrypt data with key use AES OFB algorithm. Length of `key` param should be 16, 24 or 32.</p>

<b>Signature:</b>

```go
func AesOfbDecrypt(encrypted, key []byte) []byte
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/serialt/lancet/cryptor"
)

func main() {
    data := "hello"
    key := "abcdefghijklmnop"

    encrypted := cryptor.AesOfbEncrypt([]byte(data), []byte(key))
    decrypted := cryptor.AesCfbDecrypt(encrypted, []byte(key))

    fmt.Println(string(decrypted))

    // Output:
    // hello
}
```

### <span id="Base64StdEncode">Base64StdEncode</span>

<p>Encode string with base64 encoding.</p>

<b>Signature:</b>

```go
func Base64StdEncode(s string) string
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/serialt/lancet/cryptor"
)

func main() {
    base64Str := cryptor.Base64StdEncode("hello")
    fmt.Println(base64Str)

    // Output:
    // aGVsbG8=
}
```
### <span id="Base64StdDecode">Base64StdDecode</span>

<p>Decode a base64 encoded string.</p>

<b>Signature:</b>

```go
func Base64StdDecode(s string) string
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/serialt/lancet/cryptor"
)

func main() {
    str := cryptor.Base64StdDecode("aGVsbG8=")
    fmt.Println(str)

    // Output:
    // hello
}
```

### <span id="DesEcbEncrypt">DesEcbEncrypt</span>

<p>Encrypt data with key use DES ECB algorithm. Length of `key` param should be 8.</p>

<b>Signature:</b>

```go
func DesEcbEncrypt(data, key []byte) []byte
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/serialt/lancet/cryptor"
)

func main() {
    data := "hello"
    key := "abcdefgh"

    encrypted := cryptor.DesEcbEncrypt([]byte(data), []byte(key))

    decrypted := cryptor.DesEcbDecrypt(encrypted, []byte(key))

    fmt.Println(string(decrypted))

    // Output:
    // hello
}
```
### <span id="DesEcbDecrypt">DesEcbDecrypt</span>

<p>Decrypt data with key use DES ECB algorithm. Length of `key` param should be 8.</p>

<b>Signature:</b>

```go
func DesEcbDecrypt(encrypted, key []byte) []byte
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/serialt/lancet/cryptor"
)

func main() {
    data := "hello"
    key := "abcdefgh"

    encrypted := cryptor.DesEcbEncrypt([]byte(data), []byte(key))

    decrypted := cryptor.DesEcbDecrypt(encrypted, []byte(key))

    fmt.Println(string(decrypted))

    // Output:
    // hello
}
```

### <span id="DesCbcEncrypt">DesCbcEncrypt</span>

<p>Encrypt data with key use DES CBC algorithm. Length of `key` param should be 8.</p>

<b>Signature:</b>

```go
func DesCbcEncrypt(data, key []byte) []byte
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/serialt/lancet/cryptor"
)

func main() {
    data := "hello"
    key := "abcdefgh"

    encrypted := cryptor.DesCbcEncrypt([]byte(data), []byte(key))
    decrypted := cryptor.DesCbcDecrypt(encrypted, []byte(key))

    fmt.Println(string(decrypted))

    // Output:
    // hello
}
```

### <span id="DesCbcDecrypt">DesCbcDecrypt</span>

<p>Decrypt data with key use DES CBC algorithm. Length of `key` param should be 8.</p>

<b>Signature:</b>

```go
func DesCbcDecrypt(encrypted, key []byte) []byte
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/serialt/lancet/cryptor"
)

func main() {
    data := "hello"
    key := "abcdefgh"

    encrypted := cryptor.DesCbcEncrypt([]byte(data), []byte(key))
    decrypted := cryptor.DesCbcDecrypt(encrypted, []byte(key))

    fmt.Println(string(decrypted))

    // Output:
    // hello
}
```
### <span id="DesCtrCrypt">DesCtrCrypt</span>

<p>Encrypt or decrypt data with key use DES CTR algorithm. Length of `key` param should be 8.</p>

<b>Signature:</b>

```go
func DesCtrCrypt(data, key []byte) []byte
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/serialt/lancet/cryptor"
)

func main() {
    data := "hello"
    key := "abcdefgh"

    encrypted := cryptor.DesCtrCrypt([]byte(data), []byte(key))
    decrypted := cryptor.DesCtrCrypt(encrypted, []byte(key))

    fmt.Println(string(decrypted))

    // Output:
    // hello
}
```

### <span id="DesCfbEncrypt">DesCfbEncrypt</span>

<p>Encrypt data with key use DES CFB algorithm. Length of `key` param should be 8.</p>

<b>Signature:</b>

```go
func DesCfbEncrypt(data, key []byte) []byte
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/serialt/lancet/cryptor"
)

func main() {
    data := "hello"
    key := "abcdefgh"

    encrypted := cryptor.DesCfbEncrypt([]byte(data), []byte(key))
    decrypted := cryptor.DesCfbDecrypt(encrypted, []byte(key))

    fmt.Println(string(decrypted))

    // Output:
    // hello
}
```
### <span id="DesCfbDecrypt">DesCfbDecrypt</span>

<p>Decrypt data with key use DES CBC algorithm. Length of `key` param should be 8.</p>

<b>Signature:</b>

```go
func DesCfbDecrypt(encrypted, key []byte) []byte
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/serialt/lancet/cryptor"
)

func main() {
    data := "hello"
    key := "abcdefgh"

    encrypted := cryptor.DesCfbEncrypt([]byte(data), []byte(key))
    decrypted := cryptor.DesCfbDecrypt(encrypted, []byte(key))

    fmt.Println(string(decrypted))

    // Output:
    // hello
}
```
### <span id="DesOfbEncrypt">DesOfbEncrypt</span>

<p>Enecrypt data with key use DES OFB algorithm. Length of `key` param should be 8.</p>

<b>Signature:</b>

```go
func DesOfbEncrypt(data, key []byte) []byte
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/serialt/lancet/cryptor"
)

func main() {
    data := "hello"
    key := "abcdefgh"

    encrypted := cryptor.DesOfbEncrypt([]byte(data), []byte(key))
    decrypted := cryptor.DesOfbDecrypt(encrypted, []byte(key))

    fmt.Println(string(decrypted))

    // Output:
    // hello
}
```
### <span id="DesOfbDecrypt">DesOfbDecrypt</span>

<p>Decrypt data with key use DES OFB algorithm. Length of `key` param should be 8.</p>

<b>Signature:</b>

```go
func DesOfbDecrypt(encrypted, key []byte) []byte
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/serialt/lancet/cryptor"
)

func main() {
    data := "hello"
    key := "abcdefgh"

    encrypted := cryptor.DesOfbEncrypt([]byte(data), []byte(key))
    decrypted := cryptor.DesOfbDecrypt(encrypted, []byte(key))

    fmt.Println(string(decrypted))

    // Output:
    // hello
}
```

### <span id="HmacMd5">HmacMd5</span>

<p>Get the md5 hmac hash of string.</p>

<b>Signature:</b>

```go
func HmacMd5(data, key string) string
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/serialt/lancet/cryptor"
)

func main() {
      str := "hello"
    key := "12345"

    hms := cryptor.HmacMd5(str, key)
    fmt.Println(hms)

    // Output:
    // e834306eab892d872525d4918a7a639a
}
```
### <span id="HmacSha1">HmacSha1</span>

<p>Get the sha1 hmac hash of string.</p>

<b>Signature:</b>

```go
func HmacSha1(data, key string) string
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/serialt/lancet/cryptor"
)

func main() {
    str := "hello"
    key := "12345"

    hms := cryptor.HmacSha1(str, key)
    fmt.Println(hms)

    // Output:
    // 5c6a9db0cccb92e36ed0323fd09b7f936de9ace0
}
```
### <span id="HmacSha256">HmacSha256</span>

<p>Get the sha256 hmac hash of string</p>

<b>Signature:</b>

```go
func HmacSha256(data, key string) string
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/serialt/lancet/cryptor"
)

func main() {
    str := "hello"
    key := "12345"

    hms := cryptor.HmacSha256(str, key)
    fmt.Println(hms)

    // Output:
    // 315bb93c4e989862ba09cb62e05d73a5f376cb36f0d786edab0c320d059fde75
}
```

### <span id="HmacSha512">HmacSha512</span>

<p>Get the sha512 hmac hash of string.</p>

<b>Signature:</b>

```go
func HmacSha512(data, key string) string
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/serialt/lancet/cryptor"
)

func main() {
    str := "hello"
    key := "12345"

    hms := cryptor.HmacSha512(str, key)
    fmt.Println(hms)

    // Output:
    // dd8f1290a9dd23d354e2526d9a2e9ce8cffffdd37cb320800d1c6c13d2efc363288376a196c5458daf53f8e1aa6b45a6d856303d5c0a2064bff9785861d48cfc
}
```


### <span id="Md5String">Md5String</span>

<p>Get the md5 value of string.</p>

<b>Signature:</b>

```go
func Md5String(s string) string
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/serialt/lancet/cryptor"
)

func main() {
    str := "hello"

    md5Str := cryptor.Md5String(str)
    fmt.Println(md5Str)

    // Output:
    // 5d41402abc4b2a76b9719d911017c592
}
```

### <span id="Md5File">Md5File</span>

<p>Get the md5 value of file.</p>

<b>Signature:</b>

```go
func Md5File(filepath string) (string, error)
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/serialt/lancet/cryptor"
)

func main() {
    s := cryptor.Md5File("./main.go"))
    fmt.Println(s)
}
```

### <span id="Sha1">Sha1</span>

<p>Get the sha1 value of string.</p>

<b>Signature:</b>

```go
func Sha1(data string) string
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/serialt/lancet/cryptor"
)

func main() {
    str := "hello"

    result := cryptor.Sha1(str)
    fmt.Println(result)

    // Output:
    // aaf4c61ddcc5e8a2dabede0f3b482cd9aea9434d
}
```

### <span id="Sha256">Sha256</span>

<p>Get the sha256 value of string.</p>

<b>Signature:</b>

```go
func Sha256(data string) string
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/serialt/lancet/cryptor"
)

func main() {
    str := "hello"

    result := cryptor.Sha256(str)
    fmt.Println(result)

    // Output:
    // 2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824
}
```

### <span id="Sha512">Sha512</span>

<p>Get the sha512 value of string.</p>

<b>Signature:</b>

```go
func Sha512(data string) string
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/serialt/lancet/cryptor"
)

func main() {
    str := "hello"

    result := cryptor.Sha512(str)
    fmt.Println(result)

    // Output:
    // 9b71d224bd62f3785d96d46ad3ea3d73319bfbc2890caadae2dff72519673ca72323c3d99ba5c11d7c7acc6e14b8c5da0c4663475c2e5c3adef46f73bcdec043
}
```

### <span id="GenerateRsaKey">GenerateRsaKey</span>

<p>Create the rsa public and private key file in current directory.</p>

<b>Signature:</b>

```go
func GenerateRsaKey(keySize int, priKeyFile, pubKeyFile string) error
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/serialt/lancet/cryptor"
)

func main() {
    err := cryptor.GenerateRsaKey(4096, "rsa_private.pem", "rsa_public.pem")
    if err != nil {
        fmt.Println(err)
    }
}
```

### <span id="RsaEncrypt">RsaEncrypt</span>

<p>Encrypt data with public key file useing ras algorithm.</p>

<b>Signature:</b>

```go
func RsaEncrypt(data []byte, pubKeyFileName string) []byte
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/serialt/lancet/cryptor"
)

func main() {
    err := cryptor.GenerateRsaKey(4096, "rsa_private.pem", "rsa_public.pem")
    if err != nil {
        return
    }
      
    data := []byte("hello")
    encrypted := cryptor.RsaEncrypt(data, "rsa_public.pem")
    decrypted := cryptor.RsaDecrypt(encrypted, "rsa_private.pem")
      
    fmt.Println(string(decrypted))

    // Output:
    // hello
}
```


### <span id="RsaDecrypt">RsaDecrypt</span>

<p>Decrypt data with private key file useing ras algorithm.</p>

<b>Signature:</b>

```go
func RsaDecrypt(data []byte, privateKeyFileName string) []byte
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/serialt/lancet/cryptor"
)

func main() {
    err := cryptor.GenerateRsaKey(4096, "rsa_private.pem", "rsa_public.pem")
    if err != nil {
        return
    }
      
    data := []byte("hello")
    encrypted := cryptor.RsaEncrypt(data, "rsa_public.pem")
    decrypted := cryptor.RsaDecrypt(encrypted, "rsa_private.pem")
      
    fmt.Println(string(decrypted))

    // Output:
    // hello
}
```



