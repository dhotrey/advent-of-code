from filetree import Tree as dirtree


def main():
    with open("input.txt") as f:
        commands = f.readlines()

    filetree = get_filetree(commands)


def get_filetree(commands):
    tree = dirtree
    tree.set_root("/")

    for i in commands:
        i = i.strip()
        if "$" in i:  # i is a command
            if "cd" in i:
                _, loc = i.split(" ")

            else:
                ...
            ...
        else:
            ...


if __name__ == "__main__":
    main()
 