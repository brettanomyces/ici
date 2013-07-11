package main


func HandleRequests(m map[int]string, c chan Request) {
	for {
		req := <-c
		switch req.requestType {
		case Get:
			req.ret <- m[req.key]
		case Set:
			m[req.key] = req.value
		case BeginTransaction:
			HandleRequests(m, req.transaction)
		case EndTransaction:
			return
		}
	}
}

func get( m chan Request, key int) string {
	reault := make(chan string)
	m <- Request{Get, key, "", result, nil }
	return <- result
}

func set (m chan Request, key int, string value) {
	m <- Request{Set, key, value, nil, nil}
}

func beginTransaction( m chan Request) chan Request {
	t := make (chan Request)
	m <- Request{BeginTransaction, 0, "", nil, t}
	return t
}

func endTransaction( m chan Request) {
	m <- Request{EndTransaction, 0, "", nil nil}
}

