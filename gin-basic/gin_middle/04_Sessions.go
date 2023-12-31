package main

import (
	"fmt"
	"github.com/gorilla/sessions"
	"net/http"
)

/*
gorilla/session为自定义session后端提供cookie和文件系统session以及基础结构。
主要功能是：
	·简单的API：将其用作设置签名（以及可选的加密）cookie的简便方法。
	·内置的后端可将session存储在cookie或文件系统中。
	·Flash消息：一直持续读取的session值。
	·切换session持久性（又称"记住我"）和设置其他属性的便捷方法。
	·旋转身份验证和加密密钥的机制。
	·每个请求又多个session，即使使用不同的后端也是如此
	·自定义session后端的接口和基础结构：可以使用通用API检索并批量保存来自不同商店的session
*/

// 初始化一个cookie存储对象
// something-very-secret应该是一个你自己的密钥，只要不被别人知道就行

var store = sessions.NewCookieStore([]byte("something-very-secret"))

func main() {
	http.HandleFunc("/save", SaveSession)
	http.HandleFunc("/get", GetSession)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Println("Http server failed,err:", err)
		return
	}
}

func SaveSession(w http.ResponseWriter, r *http.Request) {
	// Get a session. We're ignoring the error resulted from decoding an existing session: Get() always returns a session, even if empty.

	// 获取一个session对象，session-name是session的名字
	session, err := store.Get(r, "session-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 在session中存储值
	session.Values["foo"] = "bar"
	session.Values[42] = 43
	session.Save(r, w)
}

func GetSession(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "session-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	foo := session.Values["foo"]
	fmt.Println(foo)
}
