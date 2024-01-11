package gosayhello

// v1.0.0
// func SayHello() string {
// 	return "hello "
// }

// v1.0.1
// func SayHello() string {
// 	return "hello world"
// }

// v2.0.0
func SayHello(name string) string {
	return "hello " + name
}

//create git
//git init
//git add .
//git commit -m "first commit"
//git remote add origin https://github.com/jesslynch24/go-say-hello.git
//git push -u origin master
//git tag v1.0.0
//git push origin v1.0.0

//open app_say_hello
//create git
//go mod init github.com/jesslynch24/app-say-hello
//go get github.com/jesslynch24/go-say-hello

//major up
//misal untuk say hello ditambah param
//maka pada go.mod diganti nama modulenya misal ditambah/v2
