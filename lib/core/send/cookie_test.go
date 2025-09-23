package send

import (
	"testing"

	"main/lib/core/mock"
)

func TestCookie(t *testing.T) {
	client := mock.NewClient()
	Cookie(client, "cookie", "monster")
	writer := client.Writer.(*mock.ResponseWriter)
	if writer.MockHeader.Get("Set-Cookie") != "cookie=monster; Path=/; HttpOnly" {
		t.Fatal("cookie should be monster")
	}
}
