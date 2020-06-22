""" Node is defined as
class node:
    def __init__(self, data):
        self.data = data
        self.left = None
        self.right = None
"""

def checkBST(root):
    return f(root, -float('inf'), float('inf'))
    
def f(node, lower, upper):
    if node is None:
        return True
    return lower < node.data < upper \
            and f(node.left, lower, node.data) \
            and f(node.right, node.data, upper)
