package main

/*
    二叉搜索树
    右子树所有节点均大于根节点， 左子树所有节点均小于根节点
    而且左右子树需要均为二叉搜索树
*/
type searchTreeNode struct {
    data int
    left *searchTreeNode
    right *searchTreeNode
}

func DeleteNode(root *searchTreeNode, value int)  *searchTreeNode {
    if root == nil {
        return nil
    }
    // 待删除节点在右子树
    if root.data < value {
        root.right = DeleteNode(root.right, value)
        return root
    }
    // 待删除节点在左子树
    if root.data > value {
        root.left = DeleteNode(root.left, value)
        return root
    }
    // 待删除节点为根节点
    // 找到右子树的最小值和根节点交换
    // 在删除掉交换后的那个节点
    if root.left == nil {
        return root.right
    }
    if root.right == nil {
        return root.left
    }
    minNode := root.right
    for minNode.left != nil {
        minNode = minNode.left
    }

    root.data = minNode.data
    root.right = DeleteNode(root.right, value)
    return root
}

func AddNode(root *searchTreeNode, value int) *searchTreeNode {
    if root == nil {
        return &searchTreeNode{
            data:  value,
        }
    }

    // 需要插入到右子树
    if root.data < value {
        root.right = AddNode(root.right, value)
        return root
    }
    // 需要插入到左子树
    if root.data > value {
        root.left = AddNode(root.left, value)
        return root
    }

    return root
}