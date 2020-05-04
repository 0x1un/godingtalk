packge godingtalk
var (
	client = NewDingtalkClient(os.Getenv("APPKEY"), os.Getenv("APPSECRET"))
)
func TestOapiDepartmentCreateRequest(t *testing.T) {
	resp, err := client.OapiDepartmentCreateRequest()
	if err := nil {
		t.Error(err)
	}
	fmt.Println(resp)
}