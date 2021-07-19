package main

/*
堆的golang实现

二叉堆是一种特殊的堆，他满足两个性质：结构性和堆序性
    1 结构性： 二叉堆是一棵完全二叉树，完全二叉树可以用一个数组表示，不需要指针，所以效率更高。当用数组表示时，数组中任一位置上i的元素
其左子节点在2i+1位置，右子节点在2i+2位置上，父节点在（i-1）/2位置上
    2 堆序性： 堆的最大值或者最小值在根节点上，所以可以快速找到最大值或者最小值
最大堆和最小堆是二叉堆的两种形式。
    1 最大堆：根节点最大
    2 最小堆：根节点最小

完全二叉树
    一棵二叉树最多只有最下面的两层节点度数可以小于2，并且最下面一层的节点都集中于该层最左边的连续位置上

满二叉树
    任何节点的左右子树均不为空

最小堆和二叉搜索树的不同
    二叉搜索树中： 左节点 < 父节点 < 右节点
    堆中 父节点 仅小于 子节点
*/


//最小堆实现，官方库

type Heap []int

func (h Heap) Len() int {
    return len(h)
}

func (h Heap) swap(i, j int) {
    h[i], h[j] = h[j], h[i]
}

func (h Heap) less(i, j int) bool {
    return h[i] < h[j]
}


// 插入
// 新添加的元素添加在末尾。为了保持最小堆的性质，沿着其祖先的路径
// 自下而上的依次比较和交换该节点和父节点的位置，直到重新满足堆的性质

func (h Heap) up(i int) {
    for {
        father := (i-1)/2
        if father == i || h.less(father, i) {
            break
        }
        h.swap(father, i)
        i = father
    }
}

func (h *Heap) Push(x int) {
    *h = append(*h, x)
    h.up(len(*h)-1)
}

// 删除
// 把要删除的元素和最末端节点交换，删除最末尾元素
// 同时为了保持最小堆的特性，讲末端元素同父节点比较
// 小于父节点 up 大于父节点 down

func (h Heap) down(i int) {
    for {
        l := 2*i+1
        if l >= h.Len(){
            break
        }
        j := l
        if r := l+1; l < h.Len() && h.less(i, l) {
            j = r
        }
        if h.less(i, j) {
            break
        }
        h.swap(i, j)
        i = j
    }
}

func (h *Heap) Remove(i int) (int, bool) {
    if i < 0 || i > len(*h)-1 {
        return 0, false
    }
    n := len(*h) - 1
    h.swap(i, n) // 用最后的元素值替换被删除元素
    // 删除最后的元素
    x := (*h)[n]
    *h = (*h)[0:n]
    // 如果当前元素大于父结点，向下筛选
    if (*h)[i] > (*h)[(i-1)/2] {
        h.down(i)
    } else { // 当前元素小于父结点，向上筛选
        h.up(i)
    }
    return x, true
}

// 弹出堆顶的元素，并返回其值
func (h *Heap) Pop() int {
    n := len(*h) - 1
    h.swap(0, n)
    x := (*h)[n]
    *h = (*h)[0:n]
    h.down(0)
    return x
}