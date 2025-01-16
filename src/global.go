package src

import "sync"

var messageCount = 0
var collecting = false
var photoDIR = "./downloads/"
var caption = ""
var wg sync.WaitGroup
