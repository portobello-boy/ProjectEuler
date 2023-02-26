import os
import sys

if len(sys.argv) != 2:
    print("Invalid arguments - provide problem number")
    exit()
    
probNum = int(sys.argv[1])
probPath = "{:03d}-{:03d}/Problem{:03d}".format((50*(probNum//50))+1, 50*((probNum//50)+1), probNum)

os.makedirs(probPath, exist_ok=True)
open("{}/prob{}.py".format(probPath, probNum), 'a').close()