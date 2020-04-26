# 常用方式

## 5种常用的算法 贪婪算法，动态规划算法，分治算法，回溯算法以及分支限界算法

### 动态规划法

通过把原问题分解为相对简单的子问题的方式求解复杂问题的方法。动态规划常常适用于有重叠子问题和最优子结构性质的问题

**基本思想:**

若要解一个给定问题，我们需要解其不同部分（即子问题），再合并子问题的解以得出原问题的解。 通常许多子问题非常相似，为此动态规划法试图仅仅解决每个子问题一次，从而减少计算量： 一旦某个给定子问题的解已经算出，则将其记忆化存储，以便下次需要同一个子问题解之时直接查表


**问题特征:**

  最优子结构：当问题的最优解包含了其子问题的最优解时，称该问题具有最优子结构性质。

  重叠子问题：在用递归算法自顶向下解问题时，每次产生的子问题并不总是新问题，有些子问题被反复计算多次。动态规划算法正是利用了这种子问题的重叠性质，对每一个子问题只解一次，而后将其解保存在一个表格中，在以后尽可能多地利用这些子问题的解。

**步骤**
  - 描述最优解的结构

  - 递归定义最优解的值

  - 按自底向上的方式计算最优解的值

  - 由计算出的结果构造一个最优解

动态规划只是在每一步的选择上，选择对当前点最有利的情况，但不代表当前的不利情况在后期的计算中不能变为其他点有利情况的一部分
  
 参考示例：
  - lc5_longest_palindrome.go 回文字符串
  - N: 689  三个无重叠子数组的最大和
  - N:64 最小路径和 


 ### 分治算法
 
 
  把一个复杂的问题分成若干个相同或相似的子问题，再把子问题分成更小的子问题… ， 知道最后子问题可以简单的直接求解，原问题的解即子问题的解的合并 
 
 
  使用场景： 
  
    1.当问题规模缩小到一定的程度就可以很容易解决
    2.该问题可以分解为若干个规模较小的相同问题
    3.该问题分解出的若干子问题的解可以合并为该问题的解
    4.每个子问题都是独立的，相互之间没有交集。
    
   示例： 快速排序、归并排序、大整数乘法、二分查找、递归
   
   
 #### 分治法和动态规划法

  分治与动态规划
  共同点：二者都要求原问题具有最优子结构性质,都是将原问题分而治之,分解成若干个规模较小(小到很容易解决的程序)的子问题.然后将子问题的解合并,形成原问题的解.

  不同点：分治法将分解后的子问题看成相互独立的，通过用递归来做。

       动态规划将分解后的子问题理解为相互间有联系,有重叠部分，需要记忆，通常用迭代来做

 
 
 ### 贪心算法

  在对问题求解时，总是做出在当前看来是最好的选择。 也就是说，不从整体最优上加以考虑，他所做出的仅是在某种意义上的局部最优解
  
  贪心算法没有固定的算法框架，算法设计的关键是贪心策略的选择。必须注意的是，贪心算法不是对所有问题都能得到整体最优解，
  选择的贪心策略必须具备 ``无后效性`` ，``即某个状态以后的过程不会影响以前的状态，只与当前状态有关``
  
  **贪心算法的一般步骤**
  
    1.建立数学模型来描述问题。
    
    2.把求解的问题分成若干个子问题。
    
    3.对每一子问题求解，得到子问题的局部最优解。
    
    4.把子问题的解局部最优解合成原来解问题的一个解
    
    ```
      从问题的某一初始解出发；
      
      while （能朝给定总目标前进一步）
      {
        利用可行的决策，求出可行解的一个解元素；
      }
      
      由所有解元素组合成问题的一个可行解；
    
    ```
    
    
 ### 回溯算法
 
 回溯法在问题的解空间树中，按深度优先策略， 从根结点出发搜索解空间树。搜索至解空间树任一点时，先判断该点是否包含问题的解，如果肯定不包含(剪枝过程)，则跳过对该结点为根的子树的搜索 ， 逐层向其祖先结点回溯，否则，进入该子树，继续按深度优先策略搜索
 
 为了避免生成那些不可能产生最优解的问题状态，要不断的利用限界函数，来剪枝那些实际上不可能产生最优解的活结点， 以减少问题的计算量。具有限界函数的深度优先生成法称为回溯法。(回溯法=穷举+剪枝)
 
 回溯法一般步骤
   1. 针对所给问题，定义问题的解空间
   2. 确定易搜索的解空间结构
   3. 以深度优先方式搜索解空间，并在搜索过程中用剪枝函数避免无效搜索。
   
   两个常用的剪枝函数：
     1. 约束函数：在扩展结点中剪去不满足约束的子树
     2. 限界函数：剪去得不到最优解的子树
     
     
 ### 分支限界算法
 
   分支限界算法是按照广度优先的方式对解空间树（状态空间树）进行搜索，从而求得最优解的算法。在搜索的过程中，采用限界函数（bound function）估算所有子节点的目标函数的可能取值，
   从而选择使目标函数取极值（极大值或者极小值）的节点作为扩展结点（如果限界值没有超过目前的最优解，则剪枝） 进行下一步搜索（重复 BFS -> 计算所有子节点限界 -> 选择最优子节点作为扩展结点的过程），从而不断调整搜索的方向，尽快找到问题的最优解 
 
   分支限界算法的一般步骤：
     1.将问题的解空间转化为图或者树的结构表示，维护一张活叶子表（可以是优先队列）。
     2.初始将根节点计算一个限界后加入活叶子表。
     3.当活叶子表不为空时，从活叶子表中取出一个限界最优的结点作为扩展结点，并将该节点去除出表。当活结点表为空时，算法结束。
     4.判断当前的扩展结点是否可以满足所有约束，并且得到一个可行解（该扩展结点是叶子节点）。
          如果是，判断优于当前最优解后，记录并更新最优解，随后将当前最优解与所有活叶子节点的限界做比较，对于限界差于最优解的活叶子结点，去除出活叶子表，并返回（3）。
          如果不是，则进入（5）。
     5.计算扩展结点的所有子节点是否满足约束条件，对于不满足约束条件的子节点，将以该节点为根的子树剪枝（丢弃）。
     6.根据限界函数，计算该节点满足约束的所有子节点的限界。对于限界差于当前最优解的子节点（ps：废了，没潜力），将以该子节点为根的子树丢弃；对于限界优于当前最优解的子节点（ps：还有潜力），将这些潜力节点作为活叶子结点添加到活叶子表，并返回（3）
     
   难点：
     1. 解空间的构造，即状态空间树的构造方法（节点的生成顺序）
     2. 剪枝函数的确定，即约束规则的确定
     3. 限界函数的确定，边界的评估方法
 
 #### 分支限界法与回溯法：
     求解目标：
         回溯法的求解目标是找出解空间树中满足约束条件的所有解，
         分支限界法的求解目标则是一次性对所有的可行子节点进行扩展，估算子节点目标函数的可能值，如果该子节点的目标函数值差于当前最优解，则丢弃；否则将其加入活叶子表，依次从表中选取使目标函数取极值的节点作为当前的扩展结点。重复这一过程，直到找到最优解
     
     搜索方式的不同：
         回溯采用深度优先搜索的方式去搜索解空间树。搜索过程中，对所有的子节点轮流进行深度优先搜索，一旦发现有不满足约束的子节点，则对该子节点为根的子树进行剪枝；否则就从该子节点深度优先搜索，直到搜索到一个满足约束条件的叶子节点，即求得一个可行解。
         分支限界采用广度优先搜索的方式去搜索解空间树。搜索过程中，先生成所有的子节点（分支），然后对所有分支计算一个函数值（限界），并根据这些函数值（计算出的上界或者下界），从中选择一个使目标函数最优（限界最优）的子节点作为扩展结点，使得搜索朝着最优解的方向快速推进，从而很快求得一个最优解