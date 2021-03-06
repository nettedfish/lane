lane
====

Lane package provides queue, priority queue, stack and deque data structures
implementations. Its was designed with simplicity, performance, and concurrent
usage in mind.

## Installation

```
	go get github.com/oleiade/lane
```

## Usage


#### Priority Queue

Pqueue is a *heap priority queue* data structure implementation. It can be whether max or min ordered, is synchronized and is safe for concurrent operations. It performs insertion and max/min removal in O(log N) time.

##### Example

```go
	// Let's create a new max ordered priority queue
	var priorityQueue *PQueue = NewPQueue(MINPQ)

	// And push some prioritized content into it
	priorityQueue.Push("easy as", 3)
	priorityQueue.Push("123", 2)
	priorityQueue.Push("do re mi", 4)
	priorityQueue.Push("abc", 1)

	// Now let's take a look at the min element in
	// the priority queue
	headValue, headPriority := priorityQueue.Head()
	fmt.Println(headValue)    // "abc"
	fmt.Println(headPriority) // 1

	// Okay the song order seems to be preserved, let's
	// roll
	var jacksonFive []string = make([]string, priorityQueue.Size())

	for i := 0; i < len(jacksonFive); i++ {
		value, _ := priorityQueue.Pop()

		jacksonFive[i] = value.(string)
	}

	fmt.Println(strings.Join(jacksonFive, " "))
```

#### Deque

Deque is a *head-tail linked list data* structure implementation. It is based on a doubly linked list container, so that every operations time complexity is O(1). Every operations over an instiated Deque are synchronized and safe for concurrent usage.

##### Example

```go
	// Let's create a new deque data structure
	var deque *Deque = NewDeque()

	// And push some content into it using the Append
	// and Prepend methods
	deque.Append("easy as")
	deque.Prepend("123")
	deque.Append("do re mi")
	deque.Prepend("abc")

	// Now let's take a look at what are the first and
	// last element stored in the Deque
	firstValue := deque.First()
	lastValue := deque.Last()
	fmt.Println(firstValue) // "abc"
	fmt.Println(lastValue)  // 1

	// Okay now let's play with the Pop and Shift
	// methods to bring the song words together
	var jacksonFive []string = make([]string, deque.Size())

	for i := 0; i < len(jacksonFive); i++ {
		value := deque.Shift()
		jacksonFive[i] = value.(string)
	}

	// abc 123 easy as do re mi
	fmt.Println(strings.Join(jacksonFive, " "))
```

#### Queue

Queue is a **FIFO** ( *First in first out* ) data structure implementation. It is based on a deque container and focuses its API on core functionalities: Enqueue, Dequeue, Head, Size, Empty. Every operations time complexity is O(1). As it is implemented using a Deque container, every operations over an instiated Queue are synchronized and safe for concurrent usage.

##### Example

```go
	// Create a new queue and pretend we're handling starbucks
	// clients
	var queue *Queue = NewQueue()

	// Let's add the incoming clients to the queue
	queue.Enqueue("grumpyClient")
	queue.Enqueue("happyClient")
	queue.Enqueue("ecstaticClient")

	fmt.Println(queue.Head) // grumpyClient

	// Let's handle the clients asynchronously
	for client := queue.Dequeue(); client != nil; {
		go fmt.Println(client)
	}
```

#### Stack

Stack is a **LIFO** ( *Last in first out* ) data structure implementation. It is based on a deque container and focuses its API on core functionalities: Push, Pop, Head, Size, Empty. Every operations time complexity is O(1). As it is implemented using a Deque container, every operations over an instiated Stack are synchronized and safe for concurrent usage.

##### Example

```go
	// Create a new stack and put some plates over it
	var stack *Stack = NewStack()

	// Let's put some plates on the stack
	stack.Push("redPlate")
	stack.Push("bluePlate")
	stack.Push("greenPlate")

	fmt.Println(stack.Head) // greenPlate

	// What's on top of the stack?
	value := stack.Pop()
	fmt.Println(value.(string)) // greenPlate

	stack.Push("yellowPlate")
	value = stack.Pop()
	fmt.Println(value.(string)) // yellowPlate

	// What's on top of the stack?
	value = stack.Pop()
	fmt.Println(value.(string)) // bluePlate

	// What's on top of the stack?
	value = stack.Pop()
	fmt.Println(value.(string)) // redPlate
```


## Documentation

For a more detailled overview of lane, please refer to [Documentation](http://godoc.org/github.com/oleiade/lane)


[![Bitdeli Badge](https://d2weczhvl823v0.cloudfront.net/oleiade/lane/trend.png)](https://bitdeli.com/free "Bitdeli Badge")

