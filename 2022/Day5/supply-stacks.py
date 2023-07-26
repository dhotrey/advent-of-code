def main():
    with open("/workspaces/AdventOfCode/2022/Day5/input.txt") as f:
        content = f.readlines()
        
    stack , instructions = separateStackfromInstructions(content)
    
    stackdct = getStack(stack)
    print(stackdct)
    

def separateStackfromInstructions(content:list):
    i = content.index("\n")
    return (content[:i],content[i+1:])

def getStack(stack:list)  -> dict:
    stack = stack[::-1]
    stackdct = {i: [] for i in range(1, 10)}
    stacknum, i = 1,0
    final_stack = []
    while i < len(stack):
        new_sub_stack = [*stack[i]]
        j = 0
        cleaned = []
        while j < len(new_sub_stack):
            if new_sub_stack[j].isalpha():
                cleaned.append(new_sub_stack[j])
                j +=1
            elif new_sub_stack[j].isnumeric():
                cleaned.append(new_sub_stack[j])
                j +=3
            elif new_sub_stack[j] == " " and new_sub_stack[j+1] == ' ' :
                cleaned.append(" ")
                j += 4
                continue
            j +=1
        final_stack.append(cleaned)
        i += 1
    for i in final_stack:
        print(i)
        for j in i:
            if (j.isalpha() or j.isspace()) and stacknum < 10:
                stackdct[stacknum].append(j)
                stacknum += 1
        stacknum = 1
    return stackdct
            



if __name__=="__main__":
    main()