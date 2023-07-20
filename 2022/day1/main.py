def main():
    try:
        with open("elfs_food.txt", "r") as file:
            lines = file.read().split("\n")
    except IOError as e:
        print(f"Failed to read file: {e}")
        return

    calorie_counts = []
    count = 0
    for line in lines:
        if line == "":
            calorie_counts.append(count)
            count = 0
            continue

        try:
            cal = int(line)
        except ValueError as e:
            print(f"Failed to convert string to int: {e}")
            return
        count += cal
    calorie_counts.append(count)

    most_calories = max(calorie_counts)

    print("Most calories is:", most_calories)


if __name__ == "__main__":
    main()
