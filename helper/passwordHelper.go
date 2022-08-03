package helper

import "golang.org/x/crypto/bcrypt"

func HashPassword(pw string) (string, error) { // 유저 pw 해싱
	bytes, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPasswordHash(hashVal, userPw string) bool { //	유저의 pw와 db의 pw의 해싱 후 결과가 같은지 확인
	err := bcrypt.CompareHashAndPassword([]byte(hashVal), []byte(userPw))
	if err != nil {
		return false
	} else {
		return true
	}
}
