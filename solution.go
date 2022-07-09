func maxProfit(prices []int) int {    

    var cp1, cp2, mp1, mp2 int = math.MaxInt, math.MaxInt, 0, 0

    for i:= 0; i < len(prices); i++ {
        if(prices[i]<cp1){
            cp1 = prices[i]
        }
        
        if(prices[i]-cp1>mp1){
            mp1 = prices[i]-cp1
        }
        
        if(prices[i]-mp1<cp2){
            cp2 = prices[i]-mp1
        }
        
        if(prices[i]-cp2>mp2){
            mp2 = prices[i]-cp2
        }
        
    }
    
    return mp2
}
