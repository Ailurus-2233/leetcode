import java.util.Deque;
import java.util.LinkedList;
import java.util.Queue;

class MaxQueue {
    Deque<Integer> deque;
    Queue<Integer> queue;

    public MaxQueue() {
        this.deque = new LinkedList<>();
        this.queue = new LinkedList<>();
    }
    
    public int max_value() {
        if (this.queue.isEmpty()) {
            return -1;
        }
        return this.deque.getFirst();
    }
    
    public void push_back(int value) {
        this.queue.add(value);
        while (!this.deque.isEmpty() && this.deque.getLast() < value) {
            this.deque.removeLast();
        }
        this.deque.addLast(value);
    }
    
    public int pop_front() {
        if (this.queue.isEmpty()) {
            return -1;
        }
        int ans = this.queue.remove();
        if (this.deque.getFirst() == ans) {
            this.deque.removeFirst();
        }
        return ans;
    }
}

/**
 * Your MaxQueue object will be instantiated and called as such:
 * MaxQueue obj = new MaxQueue();
 * int param_1 = obj.max_value();
 * obj.push_back(value);
 * int param_3 = obj.pop_front();
 */