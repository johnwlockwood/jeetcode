"""
    problems/populating-next-right-pointers-in-each-node-ii
# Definition for a Node.
class Node:
    def __init__(self, val: int = 0, left: 'Node' = None, right: 'Node' = None, next: 'Node' = None):
        self.val = val
        self.left = left
        self.right = right
        self.next = next
"""
import queue

class Solution:
    """
    BFS or Level order traversal of the binary tree
allows for a simple way to address the previous
node while iterating through each level and the
final item doesn't have a next node to use
as it's next.
    """
    def connect(self, root: 'Node') -> 'Node':
        if root is None:
            return None
        q = queue.Queue()
        q.put(root)
        
        while not q.empty():
            qs = q.qsize()
            prev = None
            for i in range(qs):
                item = q.get()
                if item.left:
                    q.put(item.left)
                if item.right:
                    q.put(item.right)
                if prev is not None:
                    prev.next = item
                prev = item
                
        return root


import collections
class Solution2:
    """
    This improves the previous solution by using
    the more efficient collections.deque and 
    avoiding keeping a reference to the prev node.
    """
    def connect(self, root: 'Node') -> 'Node':
        if root is None:
            return None
        q = collections.deque([root])
        
        while q:
            qs = len(q)
            for i in range(qs):
                node = q.popleft()
                if node.left:
                    q.append(node.left)
                if node.right:
                    q.append(node.right)
                if i < qs - 1:
                    node.next = q[0]
        return root
