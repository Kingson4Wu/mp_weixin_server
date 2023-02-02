package ngrok_test

import (
	"fmt"
	"github.com/kingson4wu/mp_weixin_server/ngrok"
	"testing"
)

func TestParse(t *testing.T) {

	str := `{"tunnels":[{"name":"first","uri":"/api/tunnels/first","public_url":"https://cce7-120-230-98-139.ngrok.io","proto":"https","config":{"addr":"http://localhost:8989","inspect":true},"metrics":{"conns":{"count":25,"gauge":0,"rate1":2.119858701315161e-16,"rate5":0.000017386968598385996,"rate15":0.0006548635586677217,"p50":90286831083,"p90":90635994164.2,"p95":91156976713.5,"p99":91183139283},"http":{"count":25,"rate1":2.1198584327510557e-16,"rate5":0.000017323897078201087,"rate15":0.0006504444031429781,"p50":5143182,"p90":46321408.40000003,"p95":647652229.999999,"p99":893081699}}},{"name":"third","uri":"/api/tunnels/third","public_url":"https://cc24-120-230-98-139.ngrok.io","proto":"https","config":{"addr":"http://localhost:8988","inspect":true},"metrics":{"conns":{"count":13,"gauge":0,"rate1":2.964393875e-314,"rate5":2.878005173863151e-304,"rate15":4.934781738357966e-104,"p50":98514080543,"p90":121052288262,"p95":123155641214,"p99":123155641214},"http":{"count":49,"rate1":2.964393875e-314,"rate5":8.538876999288701e-304,"rate15":1.474831952094303e-103,"p50":7158002,"p90":19066033,"p95":21386196.499999996,"p99":25040753}}},{"name":"second","uri":"/api/tunnels/second","public_url":"https://f11c-120-230-98-139.ngrok.io","proto":"https","config":{"addr":"http://localhost:8787","inspect":true},"metrics":{"conns":{"count":0,"gauge":0,"rate1":0,"rate5":0,"rate15":0,"p50":0,"p90":0,"p95":0,"p99":0},"http":{"count":0,"rate1":0,"rate5":0,"rate15":0,"p50":0,"p90":0,"p95":0,"p99":0}}},{"name":"first (http)","uri":"/api/tunnels/first%20%28http%29","public_url":"http://cce7-120-230-98-139.ngrok.io","proto":"http","config":{"addr":"http://localhost:8989","inspect":true},"metrics":{"conns":{"count":0,"gauge":0,"rate1":0,"rate5":0,"rate15":0,"p50":0,"p90":0,"p95":0,"p99":0},"http":{"count":0,"rate1":0,"rate5":0,"rate15":0,"p50":0,"p90":0,"p95":0,"p99":0}}}],"uri":"/api/tunnels"}`
	fmt.Println(ngrok.Parse(str))

}
