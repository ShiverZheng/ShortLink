package main

func main() {
	a := App{}
	a.Initiailize(getEnv())
	a.Run(":8000")
}
