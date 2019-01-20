package models

import (
	"github.com/astaxie/beego/orm"
	"crypto/md5"
	"encoding/hex"
	"time"
	"errors"
	"strings"
)

const TOKEN_EXPIRE_TIME = 86400

type AccessToken struct {
	Id            int
	AccountId     int
	Token         string
	CreateTime    int
	LastALiveTime int
}

func init() {
	orm.RegisterModel(new(AccessToken))
}

//生成token
func CreateToken(account_id int) (string, error) {
	o := orm.NewOrm()
	accessToken := new(AccessToken)
	accessToken.AccountId = account_id
	nowTimestamp := time.Now().Unix()
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(string(account_id) + string(nowTimestamp)))
	token := hex.EncodeToString(md5Ctx.Sum(nil))
	accessToken.Token = token
	_, err := o.Insert(accessToken)
	if err != nil {
		return "", errors.New("insert_error")
	}
	return token, err
}

//验证token
func VerifyToken(account_id int, token string) (bool, error) {
	o := orm.NewOrm()
	var accessToken AccessToken
	err := o.QueryTable("accesstoken").Filter("account_id", account_id).One(&accessToken)
	if err == orm.ErrNoRows {
		return false, errors.New("the_token_not_found")
	}
	if strings.Compare(token, accessToken.Token) != 0 {
		return false, errors.New("the_token_error")
	}
	if accessToken.CreateTime+TOKEN_EXPIRE_TIME < int(time.Now().Unix()) {
		return false, errors.New("the_token_expired")
	}
	return true, nil
}

//删除token
func DeleteToken(token string) {
	o := orm.NewOrm()
	o.Using("default")
	o.Delete(&AccessToken{Token: token})
}
