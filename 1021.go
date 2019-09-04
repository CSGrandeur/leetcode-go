// 其实主要的点就是跳过每个原语的第一个左括号，能省略几行的写法就是每次匹配完一个原语后下一个起点跳1位。
// 用 for in 这种就没法跳了，多加一个 flag 吧，很朴实哈。
func removeOuterParentheses(S string) string {
    var snum int = 0
    var res string = ""
    flag := false
    for _, ch := range S {
        if ch == '(' {
            snum ++
        } else {
            snum --
        }
        if snum > 1 {
            flag = true
        }
        if snum == 0 {
            flag = false
        }
        if flag == true {
            res += string(ch)
        }
    }
    return res
}