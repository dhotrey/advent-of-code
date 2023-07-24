def main():
    with open("input.txt") as f:
        contents = f.readlines()

    total_overlaps = 0
    for i in contents:
        elf1 , elf2 = i.strip().split(",")
        
        if isOverlapping(elf1,elf2):
            total_overlaps += 1
            
    print(f"{total_overlaps=}")

def isOverlapping(elf1,elf2) -> bool:
    e1a , e1b = map(int,elf1.split("-"))
    e2a , e2b = map(int,elf2.split("-"))
    
    e1set = set(range(e1a,e1b+1))
    e2set = set(range(e2a,e2b+1))
    
    if e1set.issubset(e2set) or e2set.issubset(e1set):
        return True
    return False



if __name__ == "__main__":
    main()