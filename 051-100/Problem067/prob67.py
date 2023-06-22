
# Thanks, Cameron
def read(filename):
    with open(filename) as f:
        return [[int(num) for num in line.strip().split(' ')] for line in f]
    
def main():
    triangle = read("./triangle.txt")
    # print(triangle)
    
    for i in reversed(range(len(triangle)-1)):
        # print(triangle[i])
        for j in range(len(triangle[i])):
            triangle[i][j] += max(triangle[i+1][j], triangle[i+1][j+1])
            
    print(triangle[0][0])

if __name__ == "__main__":
    main()