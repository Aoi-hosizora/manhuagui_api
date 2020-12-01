package service

import (
	"testing"
)

func TestXXX2(t *testing.T) {
	userService := &UserService{httpService: &HttpService{}}
	// log.Println(userService.Login("sasamori_sakuya", "chino1204"))
	// sasamori_sakuya%7c3326E7F19A7951B83EA773D6C6186838
	// log.Println(userService.CheckLogin("sasamori_sakuya%7c3326E7F19A7951B83EA773D6C6186838"))
	userService.GetShelfMangas("sasamori_sakuya%7c3326E7F19A7951B83EA773D6C6186838")
}
