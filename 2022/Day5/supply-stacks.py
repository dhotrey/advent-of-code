def main():
    with open("input.txt", "r") as f:
        content = f.readlines()

    stack, instructions = separateStackfromInstructions(content)

    stackdct = getStack(stack)
    # print(stackdct)
    stackdct = remove_empty(stackdct)
    # print(stackdct)
    rearranged_stack = apply_instructions(stackdct, instructions)
    print(rearranged_stack)
    top_of_all_stacks = get_result(rearranged_stack)
    print(top_of_all_stacks)


def get_result(stack):
    result = ""
    for i in stack:
        result += stack[i][-1]
    return result


def remove_empty(stack):
    for i in stack:
        while " " in stack[i]:
            stack[i].remove(" ")
    return stack


def apply_instructions(stack, instructions):
    for instruction in instructions:
        _, qty, _, loc, _, dest = map(
            lambda x: int(x) if x.isnumeric() else x, instruction.strip().split(" ")
        )
        stacklen = len(stack[loc])
        items = stack[loc][stacklen - qty :]
        # items = items[::-1] # uncomment for part 1 solution
        stack[dest].extend(items)
        stack[loc] = stack[loc][:-qty]

    return stack


def separateStackfromInstructions(content: list):
    i = content.index("\n")
    return (content[:i], content[i + 1 :])


def getStack(stack: list) -> dict:
    stack = stack[::-1]
    stackdct = {i: [] for i in range(1, 10)}
    stacknum, i = 1, 0
    final_stack = []
    while i < len(stack):
        new_sub_stack = [*stack[i]]
        j = 0
        cleaned = []
        while j < len(new_sub_stack):
            if new_sub_stack[j].isalpha():
                cleaned.append(new_sub_stack[j])
                j += 1
            elif new_sub_stack[j].isnumeric():
                cleaned.append(new_sub_stack[j])
                j += 3
            elif new_sub_stack[j] == " " and new_sub_stack[j + 1] == " ":
                cleaned.append(" ")
                j += 4
                continue
            j += 1
        final_stack.append(cleaned)
        i += 1
    for i in final_stack:
        # print(i)
        for j in i:
            if (j.isalpha() or j.isspace()) and stacknum < 10:
                stackdct[stacknum].append(j)
                stacknum += 1
        stacknum = 1
    return stackdct


if __name__ == "__main__":
    main()
