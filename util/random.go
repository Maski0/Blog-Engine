package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

var paraGraph = [2]string{"Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
	"Sit eius assumenda ea consequuntur modi aut dignissimos optio sit officia earum. Vel praesentium distinctio ut voluptatem dolorem aut autem omnis ut aperiam molestiae et perferendis sunt ab galisum ducimus aut voluptas eaque? Qui enim numquam non quia ipsum aut iste dignissimos? Aut provident praesentium et numquam reiciendis est harum molestiae id placeat quos ut soluta officia qui nobis eligendi. Ea suscipit ipsa hic quaerat deleniti aut praesentium aliquid est aspernatur deleniti. Et exercitationem recusandae in delectus expedita ea delectus possimus non doloremque voluptate est neque provident."}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomString(n int) string {
	var sb strings.Builder

	k := len(alphabet)

	for i := 0; i < n; i++ {
		cs := alphabet[rand.Intn(k)]
		sb.WriteByte(cs)
	}
	return sb.String()
}

func RandomUsername() string {
	return RandomString(6)
}

func RandomEmail() string {
	return RandomString(6) + "gmail.com"
}

func RandomPassWordHash() string {
	return RandomString(10)
}

func RandomContent() string {
	return paraGraph[rand.Intn(2)]
}
