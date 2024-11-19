package stack

/**
面试题 03.04. 化栈为队

实现一个MyQueue类，该类用两个栈来实现一个队列。


示例：

MyQueue queue = new MyQueue();

queue.push(1);
queue.push(2);
queue.peek();  // 返回 1
queue.pop();   // 返回 1
queue.empty(); // 返回 false

说明：

你只能使用标准的栈操作 -- 也就是只有 push to top, peek/pop from top, size 和 is empty 操作是合法的。
你所使用的语言也许不支持栈。你可以使用 list 或者 deque（双端队列）来模拟一个栈，只要是标准的栈操作即可。
假设所有操作都是有效的 （例如，一个空的队列不会调用 pop 或者 peek 操作）。
*/

// 题干是要求用两个数组实现队列
// 但对数组的使用方式有要求，就是只能把数组当成栈使用，而不能当成数组使用
// 数组当成栈使用，就是只能 append，只能取数组的最后一个值，其他比如通过索引取任意值等操作，统统不允许
// 理解了这一点，记住两句话就行了
// 一个 in，一个 out 数组
// push 的时候，只管往 in 里放东西，pop 或者 peek 的时候，从是从 out 取
// 如果 out 不为空，直接返回最后一个值，如果为空，循环取 in 的最后一个值，append 到 out，然后取最后一个值返回

type MyQueue struct {
    in, out []int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
    return MyQueue{}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
    this.in = append(this.in, x)
}

func (this *MyQueue) in2out() {
    for len(this.in) > 0 {
        x := this.in[len(this.in)-1]
        this.out = append(this.out, x)
        this.in = this.in[:len(this.in)-1]
    }
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
    if len(this.out) == 0 {
        this.in2out()
    }
    x := this.out[len(this.out)-1]
    this.out = this.out[:len(this.out)-1]
    return x
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
    if len(this.out) == 0 {
        this.in2out()
    }
    return this.out[len(this.out)-1]
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
    return len(this.in) == 0 && len(this.out) == 0
}
