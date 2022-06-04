package lib

func GetPosts() []byte {
	var myposts [2]DlogEntry

	myposts[0] = NewEntry("Hello world", "my first post", "hello-world")

	myposts[1] = NewEntry("Bye world", "my last post", "2")

	mylog := Dlog{
		Name:  "MANGO",
		Uname: "palash",
	}

	CreateNewLog(&mylog)
	getCon().Close()

	mylog.InsertNewEntry(&myposts[0])
	mylog.InsertNewEntry(&myposts[1])
	return []byte(mylog.FormatDlog())

}
