package stack

type Stack []byte

func (s *Stack) Push(v byte){
	*s = append((*s), v)
}

func (s *Stack) Pop() byte{
	l := len(*s)
	v := (*s)[l - 1]
	*s = (*s)[:l - 1]

	return v
}

func (s *Stack) Seek() byte{
	l := len(*s)
	v := (*s)[l - 1]

	return v
}

func (s *Stack) Len() int{
	return len(*s)
}

//num = "1432219", k = 3   1219
func removeKdigits(num string, k int) string {
	res := "0"
	stack := Stack{}

	//边界判断
	if len(num) == 0{
		return res
	}
	if k== 0 {
		return num
	}

	// k > 0 并且 值小于栈中的值就出栈，否则进栈，这样栈中留下的就是较小值（从前往后）
	for _,v :=range num{
		for k > 0 && stack.Len()!=0 && byte(v) < stack.Seek(){
			stack.Pop()
			k--
		}
		stack.Push(byte(v))
	}

	//注意  k 没用完，出现递增的情况  比如 12345 ，删除后面的即可
	for k>0 {
		stack.Pop()
		k--
	}

	//清除前缀为0
	i := 0
	for _,v := range stack{
		if v != '0' {
			break
		}else {
			i++
		}
	}
	res = string(stack[i:])

	//清空了返回0
	if len(res) == 0 {
		return "0"
	}
	return res
}




























/*class Solution {
//10 ms, 在所有 Java 提交中击败了83.60%的用户

//思想：贪心，让较小的数在高位，优先从最高为开始考虑
//数据结构：单调栈，压栈规则，如果比当前元素小，就向外弹；如果比当前元素大，就直接压
1342219   1219
public String removeKdigits(String num, int k) {
if(k >= num.length()){
return "0";
}

int n = num.length();
Stack<Integer> stack = new Stack<>();
int i = 0;
for( ; i < n; i++){

//如果k次数已经用完，说明，不可以再移掉数了，剩下的直接substring就可以了
//这样效率会提高很多
if(k == 0){
break;
}

int val = num.charAt(i) - '0';

//添加自己之前要把比自己大的全pop()出来
while(!stack.isEmpty() && val < stack.peek() && k > 0){
stack.pop();
k--;
}

//如果栈为空，并且自己是0的话不允许添加
if(!stack.isEmpty() || val != 0){
stack.push(val);
}


}

//这是因为在k还没有用完的情况下，遇到的剩下的序列全是递增的，则需要从后向前删除
while(k > 0){
stack.pop();
k--;
}

StringBuilder str = new StringBuilder();
while(!stack.isEmpty()){//添加的时候要从头部添加，否则就是逆序的
str.insert(0,String.valueOf(stack.pop()));
}
if(i < num.length()){//说明循环是从break出来的，那么就可以substring()
str.append(num.substring(i));
}
//这是因为 (10 , 1)这样的类似案例
return str.length() == 0 ? "0" : str.toString();

}
}*/

