"""
 The file, poker.txt, contains one-thousand random hands dealt to two players. 
 Each line of the file contains ten cards (separated by a single space): the first five are Player 1's cards and the last five are Player 2's cards. 
 You can assume that all hands are valid (no invalid characters or repeated cards), each player's hand is in no specific order, and in each hand there is a clear winner.
 How many hands does Player 1 win?
 
 This was solved using Cactus Kevs Hand Evaluator algorithm (see this video: https://www.youtube.com/watch?v=TM_sMACxSzY)
"""
import csv

# Constants
flushIdentifiers = set(["SF", "F"])
uniqueHandIdentifiers = set(["S", "HC"])

values = {"2": 1, "3": 2, "4": 4, "5": 8, "6": 16, "7": 32, "8": 64,
          "9": 128, "T": 256, "J": 512, "Q": 1024, "K": 2048, "A": 4096}
suits = {"S": 1, "H": 2, "D": 4, "C": 8}
ranks = {"2": 0, "3": 1, "4": 2, "5": 3, "6": 4, "7": 5,
         "8": 6, "9": 7, "T": 8, "J": 9, "Q": 10, "K": 11, "A": 12}
primeRank = {"2": 2, "3": 3, "4": 5, "5": 7, "6": 11, "7": 13,
             "8": 17, "9": 19, "T": 23, "J": 29, "Q": 31, "K": 37, "A": 41}

flushMask = 15 << 12
primeMask = 255

# Deck info
cardRankings = {}
flushesLookup = [0] * 7937  # Number of unique poker rankings
straightHighCardLookup = [0] * 7937  # Number of unique poker rankings
primeProductLookup = [0] * 4888
primeRankLookup = [0] * 4888

def buildCardRankings():
    for r in ranks.keys():
        for s in suits.keys():
            ranking = (((((values[r] << 4) | suits[s]) << 4)
                       | ranks[r]) << 8) | primeRank[r]
            cardRankings[r + s] = ranking

def buildFlushTable():
    with open("pokerEquivalenceClasses.csv", mode = "r") as file:
        csvFile = csv.reader(file)

        index = 1
        for lines in csvFile:
            if lines[2] in flushIdentifiers:
                rank = 0
                for c in lines[1].split(" "):
                    rank = rank | values[c]
                flushesLookup[rank] = index
            index += 1

def buildStraightHighCardLookup():
    with open("pokerEquivalenceClasses.csv", mode = "r") as file:
        csvFile = csv.reader(file)

        index = 1
        for lines in csvFile:
            if lines[2] in uniqueHandIdentifiers:
                rank = 0
                for c in lines[1].split(" "):
                    rank = rank | values[c]
                straightHighCardLookup[rank] = index
            index += 1

def buildPrimeProductLookup():
    with open("pokerEquivalenceClasses.csv", mode = "r") as file:
        csvFile = csv.reader(file)

        primeScorePairs = {}

        index = 0
        for lines in csvFile:
            if not (lines[2] in flushIdentifiers or lines[2] in uniqueHandIdentifiers):
                product = 1
                for c in lines[1].split(" "):
                    product = product * primeRank[c]
                primeProductLookup[index] = product
                primeScorePairs[product] = int(lines[0])
                index += 1
            
        primeProductLookup.sort()
        
        for i in range(0, len(primeProductLookup)):
            primeRankLookup[i] = primeScorePairs[primeProductLookup[i]]

def init():
    buildCardRankings()
    buildFlushTable()
    buildStraightHighCardLookup()
    buildPrimeProductLookup()

def isFlush(hand):
    cards = hand.split(" ")
    flushCheck = cardRankings[cards[0]] & flushMask

    for c in cards[1:]:
        flushCheck = flushCheck & (cardRankings[c] & flushMask)
    return bool(flushCheck)

def getHandRanking(hand):
    val = 0
    for c in hand.split(" "):
        val = val | cardRankings[c]
    return val >> 16

def getPrimeProduct(hand):
    product = 1
    for c in hand.split(" "):
        product *= (cardRankings[c] & primeMask)
    return product

def getRanking(hand):
    handRanking = getHandRanking(hand)
    if isFlush(hand):
        return flushesLookup[handRanking]
    elif straightHighCardLookup[handRanking] != 0:
        return straightHighCardLookup[handRanking]
    else:
        primeInd = primeProductLookup.index(getPrimeProduct(hand))
        return primeRankLookup[primeInd]

def main():
    init()

    p1Wins = 0
    p2Wins = 0

    with open("poker.txt", mode = "r") as file:
        for line in file:
            p1hand = line[0:14]
            p1ranking = getRanking(p1hand)
            p2hand = line[15:len(line)-1]
            p2ranking = getRanking(p2hand)
            
            if p1ranking < p2ranking:
                p1Wins += 1
            else:
                p2Wins += 1

    print("P1 Wins: {}, P2Wins: {}".format(p1Wins, p2Wins))

if __name__ == "__main__":
    main()
