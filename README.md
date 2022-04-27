# Crypt Laravel
[![version](https://img.shields.io/badge/version-1.0.0-red "version")](https://gitlab.com/obt-product-manager/tour-manager)

Modul ini dibuat untuk memudahkan developer golang yang harus melakukan encrypt dan decrypt string yang fungsinya sama dengan fungsi [Encrypt & Decrypt Laravel](https://laravel.com/docs/9.x/encryption).

### Get Started
```
go get github.com/hantulautt/crypt_laravel
```
### Example
```
func main() {
    encrypt := crypt_laravel.EncryptString(privateKey, "test")
    decrypt := crypt_laravel.DecryptString(privateKey, encrypt)
    
    fmt.Println("ENCRYPT : " + encrypt)
    fmt.Println("DECRYPT : " + decrypt)
}

// OUTPUT
ENCRYPT : eyJpdiI6IlFuQk1ibVpuUkhOak1sZEVPRVl5Y1E9PSIsInZhbHVlIjoidmM1YVA5REhEenFXTFNIdnJyKzczdz09IiwibWFjIjoiZDUzZjM3MTIyZjc0YWExOTMyYjhmZmRiN2E5YjE1NDM1YTkwMzY4YTRkYzRjNTdhZmM1NDg3ODUwODBlZDk1NyIsInRhZyI6IiJ9
DECRYPT : test
```

### Contributor
- [A-K](https://github.com/hantulautt)
- [Mr. H]()