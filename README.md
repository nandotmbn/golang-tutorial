# Tutorial Golang Gin Gonic

> Tutorial ini membutuhkan
>
> 1. Installed Go Languange
> 2. Code Editor (Visual Studio Code Preferred)
>

## Konten

1. Menginstall Golang
2. Write your first Golang code
3. Golang and beyond
4. RESTAPI with Gin-Gonic
5. Folder structure for Golang

### Menginstall Golang

Menginstall golang kadang agak tricky, kita coba untuk menyederhanakan.

1. Download Go language di link [https://go.dev/doc/install](https://go.dev/doc/install)
2. Klik Install
3. Lanjutkan seperti installasi pada umumnya.

### Menulis kode Golang pertamamu

Seperti biasa, kita akan coba menulis kode "Hello World"

1. Buka Code Editor
2. Open folder
3. Pada terminal, ketikkan command `go mod init <nama folder>`
   Akan terbuat file `go.mod`, untuk mengidentifikasi module yang baru saja kita buat.
4. Buat file baru dengan nama `index.go`
5. Tuliskan kode berikut

   ```go
    package main

    func main() {
      print("Hello World")
    }
   ```

6. Pada terminal, ketikkan command `go run index.go`

### Golang dan melampauinya

#### Modules

Konsep module pada golang tidak jauh dengan konsep module pada bahasa pemrograman lain.
Langsung saja ke implementasinya.
Buat struktur folder sebagai berikut

```bash
Folder
|-function
  |-add.go
  |-multiple
    |-index.go
  |-substract
    |-index.go
|-app.go
|-go.mod
```

* `function/add.go`

  ```go
  package function

  func Add(num1 int, num2 int) int {
    return num1 + num2
  }
  ```

* `function/substract/index.go`

  ```go
  package substract

  func Substract(num1 int, num2 int) int {
    return num1 - num2
  }
  ```

* `function/multiple/index.go`

  ```go
  package function

  func Multiplie(num1 int, num2 int) int {
    return num1 * num2
  }
  ```

* `index.go`
  
  ```go
  package main

  import (
    "tutorial/function"
    multiple "tutorial/function/multiple"
    "tutorial/function/substract"
  )

  func main() {
    println(function.Add(2, 3))
    println(substract.Substract(4, 1))
    println(multiple.Multiplier(3, 10))
  }

  ```

Ada banyak cara pemanggilan module yang kita buat sendiri.

#### Menginstall module package

Di golang, penginstallan module agak lain dengan bahasa pemrograman seperti JavaScript ataupun Python.

Yaitu dengan memanggil repository dari package module dengan cara:

1. Membuka terminal
2. Tuliskan `go get -u <repository>`
   contoh: `go get -u github.com/gin-gonic/gin`

Lalu anda dapat memanggil package module tersebut dengan keyword `import`
Contoh

```go
  import "github.com/gin-gonic/gin"
```

### RESTAPI dengan Gin-Gonic

Okay, seperti gelar yang akan kita dapat "Sarjana Terapan", kita akan menerapkan sedikit ilmu yang kita capai baru saja.

> Kenapa Gin?
>
> 1. Well documented
> 2. Broad community
> 3. Easy to use

Langsung saja

* `index.go`
  
  ```go
  package main

  import (
    // "tutorial/function"
    // multiple "tutorial/function/multiple"
    // "tutorial/function/substract"

    "github.com/gin-gonic/gin"
  )

  func main() {
    router := gin.Default()

    router.GET("/", func(c *gin.Context) {
      c.JSON(200, "Hello World")
    })

    router.Run("localhost:6000")
  }

  ```
