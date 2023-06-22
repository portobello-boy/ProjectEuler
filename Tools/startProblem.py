import os
import sys

if len(sys.argv) != 2:
    print('Invalid arguments - provide problem number')
    exit()
    
probNum = int(sys.argv[1])
block = probNum // 50
probPath = f'{50*block+1:03d}-{50*(block+1)}/Problem{probNum}'
# probPath = "{:03d}-{:03d}/Problem{:03d}".format((50*(probNum//50))+1, 50*((probNum//50)+1), probNum)

os.makedirs(probPath, exist_ok=True)
open(f'{probPath}/prob{probNum}.py', 'a').close()
