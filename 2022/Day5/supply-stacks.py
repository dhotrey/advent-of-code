def main():
    with open("input.txt") as f:
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
    print(stackdct)
    stacknum = 1
    for i in stack:
        print(i)
        for j in i:
            if (j.isalpha() or j == "   ") and stacknum < 10:
                stackdct[stacknum].append(j)
                stacknum += 1
                # print(stackdct)
        stacknum = 1
    return stackdct
            
    


if __name__ == "__main__":
    main()