from bigtree import Node, print_tree, find


def main():
    demo = "C:\\Users\\aryan\\OneDrive - vitbhopal.ac.in\\Documents\\sync\\Coding\\AdventOfCode\\2022\\Day7\\demo.txt"
    final = "input.txt"
    with open(final, "r") as f:
        commands = f.readlines()

    filetree = get_filetree(commands)
    print_tree(filetree)


def get_filetree(commands):
    root = Node("/")
    currentNode = root
    for command in commands:
        # print(ctr)
        # print("----------------")
        # print(command)
        command = command.strip()
        if command.startswith("$ cd"):
            _, cd, loc = command.split(" ")
            if loc == "..":
                # print(currentNode)
                currentNode = currentNode.parent
                # print(currentNode)
                continue
            else:
                if loc == currentNode.name:
                    continue

                exists = find(root, lambda node: node.name == loc)
                # siblings = [i for i in exists.siblings]
                # print(exists)
                if exists and currentNode is exists.parent:# and exists in siblings:
                    currentNode = exists
                    print(currentNode)
                    continue

        elif command.startswith("$ ls"):
            continue
        else:
            sizeOrdir, name = command.split(" ")
            if sizeOrdir.isalpha():  # dir condition
                exists = find(root, lambda node: node.name == name)
                if not exists:
                    newNode = Node(name)
                    currentNode >> newNode
            else:  # file condition
                exists = find(root, lambda node: node.name == name)
                if not exists:
                    newNode = Node(name+" --> "+sizeOrdir, size=sizeOrdir)
                    currentNode >> newNode
        if newNode.parent.name == "/":
            root = currentNode
        # print_tree(root)
        
    # print(currentNode)
    return root


if __name__ == "__main__":
    main()
