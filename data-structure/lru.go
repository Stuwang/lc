package data_structure


type LinkNode struct {
	Prev *LinkNode
	Next *LinkNode
	Key  int
	Val  int
}

func (l *LinkNode) removeSelf() {
	l.Prev.Next = l.Next
	l.Next.Prev = l.Prev
}

func ContactNode(l, r *LinkNode) {
	l.Next = r
	r.Prev = l
}

type LRUCache struct {
	Map  map[int]*LinkNode
	Link *LinkNode
	Cap  int
}

func (this *LRUCache) removeLatest() *LinkNode {
	if this.Link.Next == this.Link {
		return nil
	}
	last := this.Link.Prev
	last.Prev.Next = last.Next
	last.Next.Prev = last.Prev
	return last
}

func Constructor(capacity int) LRUCache {
	res := LRUCache{
		Map:  make(map[int]*LinkNode),
		Link: &LinkNode{},
		Cap: capacity,
	}
	ContactNode(res.Link, res.Link)
	return res
}

func (this *LRUCache) Get(key int) int {
	if n, found := this.Map[key]; found {
		n.removeSelf()
		ContactNode(n, this.Link.Next)
		ContactNode(this.Link, n)
		return n.Val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if n, found := this.Map[key]; found {
		n.removeSelf()
		ContactNode(n, this.Link.Next)
		ContactNode(this.Link, n)
		n.Val = value
		return
	}
	if len(this.Map) >= this.Cap {
		last := this.removeLatest()
		if last != nil {
			delete(this.Map, last.Key)
		}
	}
	// new node
	newNode := &LinkNode{
		Prev: nil,
		Next: nil,
		Key:  key,
		Val:  value,
	}
	ContactNode(newNode, this.Link.Next)
	ContactNode(this.Link, newNode)
	this.Map[key] = newNode
}
