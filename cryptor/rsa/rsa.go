package rsa

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"
	"os"
	"runtime"
)

var RSA = &RSASecurity{}

type RSASecurity struct {
	pubStr string          //公钥字符串
	priStr string          //私钥字符串
	pubkey *rsa.PublicKey  //公钥
	prikey *rsa.PrivateKey //私钥
}

// // 设置公钥
// func (rsas *RSASecurity) SetPublicKey(pubStr string) (err error) {
// 	rsas.pubStr = pubStr
// 	rsas.pubkey, err = rsas.GetPublickey()
// 	return err
// }

// // 设置私钥
// func (rsas *RSASecurity) SetPrivateKey(priStr string) (err error) {
// 	rsas.priStr = priStr
// 	rsas.prikey, err = rsas.GetPrivatekey()
// 	return err
// }

// // *rsa.PublicKey
// func (rsas *RSASecurity) GetPrivatekey() (*rsa.PrivateKey, error) {
// 	return getPriKey([]byte(rsas.priStr))
// }

// // *rsa.PrivateKey
// func (rsas *RSASecurity) GetPublickey() (*rsa.PublicKey, error) {
// 	return getPubKey([]byte(rsas.pubStr))
// }

/*
	GenerateRSAKey
	bits: rsa 加密位数
*/

// GenerateRSAKey 创建RSA 常用位数 1024 2048 4096
func GenerateRSAKey(bits int) (priKey, pubKey []byte, err error) {
	var priWriter bytes.Buffer
	var pubWriter bytes.Buffer

	//GenerateKey函数使用随机数据生成器random生成一对具有指定字位数的RSA密钥
	//Reader是一个全局、共享的密码用强随机数生成器
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return
	}
	//保存私钥
	//通过x509标准将得到的ras私钥序列化为ASN.1 的 DER编码字符串
	X509PrivateKey := x509.MarshalPKCS1PrivateKey(privateKey)

	//使用pem格式对x509输出的内容进行编
	//构建一个pem.Block结构体对象
	privateBlock := pem.Block{Type: "RSA Private Key", Bytes: X509PrivateKey}
	//将数据保存到文件
	if err = pem.Encode(&priWriter, &privateBlock); err != nil {
		return
	}

	//保存公钥
	//获取公钥的数据
	publicKey := privateKey.PublicKey
	//X509对公钥编码
	X509PublicKey, err := x509.MarshalPKIXPublicKey(&publicKey)
	if err != nil {
		panic(err)
	}

	//创建一个pem.Block结构体对象
	publicBlock := pem.Block{Type: "RSA Public Key", Bytes: X509PublicKey}
	//保存到文件

	if err = pem.Encode(&pubWriter, &publicBlock); err != nil {
		return
	}
	priKey = priWriter.Bytes()
	pubKey = pubWriter.Bytes()

	return priKey, pubKey, nil
}

// GenerateRSAKeyInFile 创建公私钥并写入文件
func GenerateRSAKeyInFile(bits int, filePath string) (err error) {
	//GenerateKey函数使用随机数据生成器random生成一对具有指定字位数的RSA密钥
	//Reader是一个全局、共享的密码用强随机数生成器
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		panic(err)
	}
	//保存私钥
	//通过x509标准将得到的ras私钥序列化为ASN.1 的 DER编码字符串
	X509PrivateKey := x509.MarshalPKCS1PrivateKey(privateKey)
	//使用pem格式对x509输出的内容进行编码
	//创建文件保存私钥
	privateFile, err := os.Create(filePath + "-private.pem")
	if err != nil {
		panic(err)
	}
	defer privateFile.Close()
	//构建一个pem.Block结构体对象
	privateBlock := pem.Block{Type: "RSA Private Key", Bytes: X509PrivateKey}
	//将数据保存到文件
	if err = pem.Encode(privateFile, &privateBlock); err != nil {
		return
	}

	//保存公钥
	//获取公钥的数据
	publicKey := privateKey.PublicKey
	//X509对公钥编码
	X509PublicKey, err := x509.MarshalPKIXPublicKey(&publicKey)
	if err != nil {
		panic(err)
	}
	//pem格式编码
	//创建用于保存公钥的文件
	publicFile, err := os.Create(filePath + "-public.pem")
	if err != nil {
		panic(err)
	}
	defer publicFile.Close()
	//创建一个pem.Block结构体对象
	publicBlock := pem.Block{Type: "RSA Public Key", Bytes: X509PublicKey}
	//保存到文件
	if err = pem.Encode(publicFile, &publicBlock); err != nil {
		return err
	}
	return err
}

// GenerateRSAKeyWithPwd 创建带密码的RSA
func GenerateRSAKeyWithPwd(passwd string, bits int) (priKey, pubKey string, err error) {

	var priWriter bytes.Buffer
	var pubWriter bytes.Buffer

	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return
	}
	//通过x509标准将得到的ras私钥序列化为ASN.1 的 DER编码字符串
	x509PrivateKey := x509.MarshalPKCS1PrivateKey(privateKey)
	//使用pem格式对x509输出的内容进行编码

	//构建一个pem.Block结构体对象
	//privateBlock := pem.Block{Type: "RSA Private Key", Bytes: X509PrivateKey}
	privateBlock, err := x509.EncryptPEMBlock(rand.Reader, "RSA Private Key", x509PrivateKey, []byte(passwd), x509.PEMCipherAES256)
	if err != nil {
		return "", "", err
	}
	//将数据保存到文件
	err = pem.Encode(&priWriter, privateBlock)
	if err != nil {
		return
	}
	//X509对公钥编码
	X509PublicKey, err := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	if err != nil {
		return
	}

	//创建一个pem.Block结构体对象
	publicBlock := pem.Block{Type: "RSA Public Key", Bytes: X509PublicKey}

	err = pem.Encode(&pubWriter, &publicBlock)
	if err != nil {
		return
	}

	priKey = priWriter.String()
	pubKey = pubWriter.String()
	return priKey, pubKey, nil
}

// RSAEncrypt rsa加密
func RSAEncrypt(plainText, key []byte) (cryptText []byte, err error) {
	block, _ := pem.Decode(key)
	defer func() {
		if err := recover(); err != nil {
			switch err.(type) {
			case runtime.Error:
				log.Println("runtime err:", err, "Check that the key is correct")
			default:
				log.Println("error:", err)
			}
		}
	}()
	publicKeyInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	publicKey := publicKeyInterface.(*rsa.PublicKey)

	cipherText, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, plainText)
	if err != nil {
		return nil, err
	}
	return cipherText, nil
}

// RsaDecrypt rsa解密
func RsaDecrypt(cryptText, key []byte) (plainText []byte, err error) {
	block, _ := pem.Decode(key)

	defer func() {
		if err := recover(); err != nil {
			switch err.(type) {
			case runtime.Error:
				log.Println("runtime err:", err, "Check that the key is correct")
			default:
				log.Println("error:", err)
			}
		}
	}()
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return []byte{}, err
	}
	plainText, err = rsa.DecryptPKCS1v15(rand.Reader, privateKey, cryptText)
	if err != nil {
		return []byte{}, err
	}
	return plainText, nil
}

// RSAEncryptOAEP 公钥加密
func RSAEncryptOAEP(plainText []byte, pubCipherKey []byte) (cipherText []byte, err error) {

	block, _ := pem.Decode(pubCipherKey)

	//x509解码
	publicKeyInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		panic(err)
	}

	// 类型断言
	publickey := publicKeyInterface.(*rsa.PublicKey)

	cipherText, err = rsa.EncryptOAEP(sha256.New(), rand.Reader, publickey, plainText, []byte(nil))
	if err != nil {
		err = fmt.Errorf("rsa encryptOAEP: error from encryption: %s", err)
		return

	}
	return cipherText, nil

}

// RSADecryptOAEP 私钥解密
func RSADecryptOAEP(cipherText, privCipherKey []byte) (plainText []byte, err error) {
	//ciphertext, _ := hex.DecodeString(cipherText)

	lable := []byte(nil)

	// pem解码
	block, _ := pem.Decode(privCipherKey)

	// 向509解码
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		panic(err)
	}

	// 对密文进行解密
	plainText, err = rsa.DecryptOAEP(sha256.New(), rand.Reader, privateKey, cipherText, lable)
	if err != nil {
		err = fmt.Errorf("rsa decryptOAEP: error from decryption: %s", err)
		return

	}
	return plainText, nil
}

// RsaSignPKCS1v15 rsa签名
func RsaSignPKCS1v15(msg, Key []byte) (cryptText []byte, err error) {
	block, _ := pem.Decode(Key)
	defer func() {
		if err := recover(); err != nil {
			switch err.(type) {
			case runtime.Error:
				log.Println("runtime err:", err, "Check that the key is correct")
			default:
				log.Println("error:", err)
			}
		}
	}()
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	myHash := sha256.New()
	myHash.Write(msg)
	hashed := myHash.Sum(nil)
	sign, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed)
	if err != nil {
		return nil, err
	}
	return sign, nil
}

// RsaVerifySignPKCS1v15 rsa验证签名
func RsaVerifySignPKCS1v15(msg []byte, sign []byte, Key []byte) bool {
	block, _ := pem.Decode(Key)
	defer func() {
		if err := recover(); err != nil {
			switch err.(type) {
			case runtime.Error:
				log.Println("runtime err:", err, "Check that the key is correct")
			default:
				log.Println("error:", err)
			}
		}
	}()
	publicInterface, _ := x509.ParsePKIXPublicKey(block.Bytes)
	publicKey := publicInterface.(*rsa.PublicKey)
	myHash := sha256.New()
	myHash.Write(msg)
	hashed := myHash.Sum(nil)
	result := rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, hashed, sign)
	return result == nil
}
