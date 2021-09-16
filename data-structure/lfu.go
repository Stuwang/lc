package data_structure


type LFUNode struct {
	Prev  *LFUNode
	Next  *LFUNode
	Key   int
	Val   int
	Count int
}

func (l *LFUNode) removeSelf() {
	l.Prev.Next = l.Next
	l.Next.Prev = l.Prev
}

func ContactLFUNode(l, r *LFUNode) {
	l.Next = r
	r.Prev = l
}

type LFUCache struct {
	Map  map[int]*LFUNode
	Head *LFUNode
	Cap  int
}

func Constructor(capacity int) LFUCache {
	res := LFUCache{
		Map:  map[int]*LFUNode{},
		Head: &LFUNode{},
		Cap:  capacity,
	}
	ContactLFUNode(res.Head, res.Head)
	return res
}

func (this *LFUCache) Get(key int) int {
	if n, found := this.Map[key]; found {
		n.Count++
		if n.Prev != this.Head {
			n.removeSelf()
			cur := n.Prev
			for cur != this.Head {
				if cur.Count > n.Count {
					break
				}
				cur = cur.Prev
			}
			ContactLFUNode(n, cur.Next)
			ContactLFUNode(cur, n)
		}
		return n.Val
	}
	return -1
}

func (this *LFUCache) Put(key int, value int) {
    if this.Cap == 0 {
		return
	}
	if n, found := this.Map[key]; found {
		n.Count++
		n.Val = value
		if n.Prev != this.Head {
			n.removeSelf()
			cur := n.Prev
			for cur != this.Head {
				if cur.Count > n.Count {
					break
				}
				cur = cur.Prev
			}
			ContactLFUNode(n, cur.Next)
			ContactLFUNode(cur, n)
		}
		return
	}
	if len(this.Map) >= this.Cap {
		// remove last
		last := this.Head.Prev
		delete(this.Map, last.Key)
		ContactLFUNode(last.Prev, this.Head)
	}
	newNode := &LFUNode{
		Prev:  nil,
		Next:  nil,
		Key:   key,
		Val:   value,
		Count: 1,
	}
	this.Map[key] = newNode
	cur := this.Head.Prev
	for cur != this.Head {
		if cur.Count > newNode.Count {
			break
		}
		cur = cur.Prev
	}
	ContactLFUNode(newNode, cur.Next)
	ContactLFUNode(cur, newNode)
}


/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
