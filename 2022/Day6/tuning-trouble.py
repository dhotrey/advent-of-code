import sys
sys.setrecursionlimit(5000)
def main():
    with open("input.txt") as f:
        contents = f.readline()
        # print(len(contents))

    char_pos = 1
    marker_pos = get_marker_pos(contents,char_pos)
    print(marker_pos)

def get_marker_pos(contents,char_pos):
    marker = contents[0:14] # make marker[0:4] for part1 sol
    if isUnique(marker):
        return char_pos + 13 # change 13 to 3 for part1 sol
    else:
        return get_marker_pos(contents[1:],char_pos+1)
        
    
def isUnique(marker):
    return True if len(marker) == len(set(marker)) else False
if __name__ == "__main__":
    main()