package main

import "fmt"

func main() {
	var s string
	s = "how are you?"
	s = `how are you?`

	fmt.Println(s)
	// result: how are you?

	s = "<htlm>\n\t<body>\"Hello\"</body>\n</html>"
	fmt.Println(s)
	// resilt:
	// <htlm>
	//     <body>"Hello"</body>
	// </html>

	s = `
<htlm>
	<body>"hello"</body>
</html>`
	fmt.Println(s)
	// result:
	// <htlm>
	//     <body>"hello"</body>
	// </html>

	fmt.Println("c:\\my\\dir\\file")
	// result: c:\my\dir\file
	fmt.Println(`c:\my\dir\file`)
	// result: c:\my\dir\file
}
