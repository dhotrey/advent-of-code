def main():
    with open("input.txt") as f:
        content = f.readlines()
    
    points_mapping = {"X": 1, "Y": 2, "Z": 3}
    win_conditions = {("A", "B"), ("B", "C"), ("C", "A")}
    
    old_pts = 0
    actual_pts = 0

    for line in content:
        opponent, my_play = line.strip().split()

        old_pts += points_mapping[my_play]

        if (opponent, my_play) in win_conditions:
            old_pts += 6
        elif opponent == my_play:
            old_pts += 3

        actual_pts += get_actual_pts(opponent, my_play)

    print(f"oldpts={old_pts}")
    print(f"actualpts={actual_pts}")


def get_actual_pts(opponent, outcome):
    pts_mapping = {"X": 3, "Y": 3, "Z": 6}
    round_outcome_mapping = {"A": 1, "B": 2, "C": 3}

    pts = pts_mapping[outcome]

    if outcome == opponent:
        pts += 3

    pts += round_outcome_mapping[opponent]

    return pts


if __name__ == "__main__":
    main()
