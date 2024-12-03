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

    list_1.sort()
    list_2.sort()
    
    total_dif = 0

    for i in range(len(list_1)):
        total_dif = total_dif + (abs(list_1[i] - list_2[i]))

    print(total_dif)


if __name__ == "__main__":
    main()