import string
def main():
    with open("input.txt") as f:
        contents = f.readlines()
    total = 0
    new_total = 0
    
    pts_map = dict(zip(string.ascii_letters,range(1,53)))

    for i in contents:
        compt1 , compt2 = split_word(i.strip())
        # print(f"{len(i)} \n {len(compt1)} : {compt1=},\n {len(compt2)} : {compt2=}") # debug
        repeated = get_repeated(compt1,compt2)
        total += pts_map[repeated]
    print(f"{total=}")
    
    group = []
    contents.append(" ")

    for j,i in enumerate(contents,start=1):
        # print(f"{i=} \n Iteration {j=}")
        if len(group) < 3:
            group.append(i)
            continue
        group.append(i)
        common_ltr = get_repeated_from3(group[:3])
        # print(f"{common_ltr=}")
        new_total += pts_map[common_ltr]
        # print(f"{group=}")
        group = group[3:]
        # print(f"modified{group=}")
        # print("---------------------------------------------")
    print(f"{new_total=}")
    

        
        
def split_word(word):
    quotient,remainder = divmod(len(word),2)
    first = word[:quotient + remainder]
    last = word[quotient + remainder:]
    return first , last
    
def get_repeated(word1,word2):
    for i in word1:
        if i in word2:
            return i

def get_repeated_from3(wordarr):
    # print(f"{wordarr=}")
    w1 , w2 , w3 = wordarr
    
    for i in w1:
        if i in w2 and i in w3:
            return i
        
if __name__ == "__main__":
    main()