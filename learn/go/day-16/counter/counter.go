



package counter


type Counter struct{
	Value int

}

func (c *Counter) Inc(){
	c.Value++
}

func UnsafeAdd(c *Counter, n int){
	for i := 0; i < n; i++{
		go c.Inc()
	}
}