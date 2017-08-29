# coding_exercise

在一次聚会期间,哈利想点亮N个灯泡。他计算了所有灯泡与正门的距离。由于可用的电源接头有限，这些灯泡中只有一些会发光。哈里决定用最小长度的铜线将所有其余灯泡连接到发光的灯泡上,使得他们都能发光。可以确定的是，至少有一个灯泡连接到电源。写出一个的算法来帮助哈利求出让所有灯泡放光所需铜线的最小长度。
## 输入
函数/方法的输入包括三个整数——
```
numOfBulbs , 表示灯泡数(N)的整数。
glowState , 一个整数列表，表示各灯泡发光(1)或不发光(0)的初始状态。 
distanceOfBulb , 一个整数列表，表示灯泡到正门的距离。
```
## 输出
返回一个整数，表示让所有灯泡发光所需铜线的最小长度。

## 示例
### 输入:
```
numOfBulbs =3 
glowState=[1,0,0] 
distanceOfBulb=[1,5,6] 
```
### 输出:
5
### 解释:
将第二个灯泡连接到第一个灯泡所需的电线长度= 4 将第三个灯泡连接到第二个灯泡所需的电线长度= 1 电线总长度= 5(4+1)

## 测试用例:
### Testcase1:
Input:
```
7,
[1,0,1,1,0,1,1], 
[1,5,6,7,8,9,17] 
Expected Return Value: 2
```
### Testcase2:
Input:
```
8,
[0,1,0,0,1,1,0,0], 
[3,5,10,12,13,23,30,38] 
Expected Return Value: 20
```

## 答题框:
```
//import library packages needed by your program 
//some classes within a package may be restricted 
//define any class and method needed
//class begins , this class is required

public class Solution{
//method signature begins , this method is required
Public int minWireLength(int numOfBulbs , int [] gloeState , int [] distanceOfBulb){
//write your code here
}
//method signature ends
}
```
