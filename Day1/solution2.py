def main():
    list_1 = []
    list_2 = []

    first = True

    with open("input.txt") as input:
        for item in input:
            for number in item.split():
                number = int(number)
                if first:
                    list_1.append(number)
                    first = not first
                else:
                    list_2.append(number)
                    first = not first

    counts = {}

    for i in range(len(list_2)):
        if list_2[i] in counts:
            counts[list_2[i]] = counts[list_2[i]] + 1
        else:
            counts[list_2[i]] = 1

    total_dif = 0

    for i in range(len(list_1)):
        freq = counts[list_1[i]] if list_1[i] in counts else 0

        total_dif = total_dif + (list_1[i] * freq)

    print(total_dif)


if __name__ == "__main__":
    main()