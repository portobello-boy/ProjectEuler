def numeral_to_number(s: str) -> int:
    n = 0
    for i in range(len(s)):
        match s[i]:
            case "M":
                n += 1000
            case "D":
                n += 500
            case "C":
                if i+1 < len(s):
                    match s[i+1]:
                        case "M":
                            n -= 100
                        case "D":
                            n -= 100
                        case _:
                            n += 100
                else:
                    n += 100
            case "L":
                n += 50
            case "X":
                if i+1 < len(s):
                    match s[i+1]:
                        case "C":
                            n -= 10
                        case "L":
                            n -= 10
                        case _:
                            n += 10
                else:
                    n += 10
            case "V":
                n += 5
            case "I":
                if i+1 < len(s):
                    match s[i+1]:
                        case "X":
                            n -= 1
                        case "V":
                            n -= 1
                        case _:
                            n += 1
                else:
                    n += 1
                        
                
        
    return n

def number_to_numeral(n: int) -> str:
    s = ""
    while n - 1000 >= 0:
        n -= 1000
        s += "M"
    if n - 900 >= 0:
        n -= 900
        s += "CM"
    while n - 500 >= 0:
        n -= 500
        s += "D"
    if n - 400 >= 0:
        n -= 400
        s += "CD"
    while n - 100 >= 0:
        n -= 100
        s += "C"
    if n - 90 >= 0:
        n -= 90
        s += "XC"
    while n - 50 >= 0:
        n -= 50
        s += "L"
    if n - 40 >= 0:
        n -= 40
        s += "XL"
    while n - 10 >= 0:
        n -= 10
        s += "X"
    if n - 9 >= 0:
        n -= 9
        s += "IX"
    while n - 5 >= 0:
        n -= 5
        s += "V"
    if n - 4 >= 0:
        n -= 4
        s += "IV"
    while n - 1 >= 0:
        n -= 1
        s += "I"
    return s

def main():
    sum = 0
    with open("./roman.txt") as readFile:
        for l in readFile:
            n = numeral_to_number(l.strip())
            s = number_to_numeral(n)
            
            print(f'{l.strip()} -- {s} -- {n}')
            sum += abs(len(l.strip()) - len(s))
    print(sum)
               
            
    # print(str(sum)[0:10])

if __name__ == "__main__":
    print(numeral_to_number("CMLXIX"))
    main()
    print(number_to_numeral(49))