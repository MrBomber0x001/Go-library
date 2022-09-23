// sample program to show how to write a simple version of curl using
// the io.Reader and io.Writer interface support

package main


func init(){
	if len(os.Args) != 2 {
		fmt.Println("Usage: ./example2 <url>");
		os.Exist(1);
	}
}

func main(){
	// get the response [r] from web server
	r, err := http.Get(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	// copy from the body to Stdout
	io.Copy(os.stdout, r.Body)
	if err := r.Body.Close(); err != nil {
		fmt.Println(err)
	}

}