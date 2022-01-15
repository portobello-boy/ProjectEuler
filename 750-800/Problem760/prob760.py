
import threading

BOUND = 10**18

THREADS = 8
THREADLOCK = threading.Lock()

total = 0

def g(m:int, n:int) -> int:
    return (m ^ n) + (m | n) + (m & n)

def G(lowN:int, highN:int) -> int:
    global total
    for n in range(lowN, highN):
        for k in range((n+1)//2):
            THREADLOCK.acquire()

            total += 2*g(k, n-k)

            THREADLOCK.release()
        if n % 2 == 0:
            THREADLOCK.acquire()

            total += g(int(n/2), int(n/2))

            THREADLOCK.release()


def main():
    threads = []
    threads.append(
        threading.Thread(target=G, args=(0, int(BOUND/THREADS)))
    )
    for i in range(1, THREADS-1):
        threads.append(
            threading.Thread(target=G, args=(int(i * BOUND/THREADS), int((i+1) * BOUND/THREADS)))
        )
    
    threads.append(
        threading.Thread(target=G, args=(int((THREADS-1) * BOUND/THREADS), BOUND+1))
    )

    for t in threads:
        t.start()

    for t in threads:
        t.join()

    print(total % 1000000007)

if __name__ == "__main__":
    main()
