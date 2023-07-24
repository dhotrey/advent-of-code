import string
def main():
    with open("input.txt") as f:
        contents = f.readlines()
    total = 0
    
    pts_map = dict(zip(string.ascii_letters,range(1,53)))

    for i in contents:
        compt1 , compt2 = split_word(i.strip())
        # print(f"{len(i)} \n {len(compt1)} : {compt1=},\n {len(compt2)} : {compt2=}") # debug
        repeated = get_repeated(compt1,compt2)
        total += pts_map[repeated]
    print(total)
        
        
def split_word(word):
    quotient,remainder = divmod(len(word),2)
    first = word[:quotient + remainder]
    last = word[quotient + remainder:]
    return first , last
    
def get_repeated(word1,word2):
    for i in word1:
        if i in word2:
            return i
        
if __name__ == "__main__":
    main()