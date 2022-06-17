
import threading

BOUND = 10**6

THREADS = 4
THREADLOCK = threading.Lock()

total = 14 # Includes n = 0, 1, 2

def g(m:int, n:int) -> int:
    return (m ^ n) + (m | n) + (m & n)

def G(lowN:int, highN:int) -> int:
    global total

    # print("{} - lowN: {}, highN: {}".format(threading.current_thread().name, lowN, highN))

    for n in range(lowN, highN):
        # if n > lowN and n % ((highN - lowN) // 10) == 0:
        #     print("{} - {:.2f}% - lowN: {}, n: {}, highN: {}".format(threading.current_thread().name, 100 * (n - lowN) / (highN - lowN), lowN, n, highN))

        # Known Values
        THREADLOCK.acquire()

        total += 4*n # g(0, n) + g(n, 0) = 4n

        THREADLOCK.release()
        
        for k in range(1, (n+1)//2):
            THREADLOCK.acquire()

            total += 2 * g(k, n-k)

            THREADLOCK.release()
        if n % 2 == 0:
            THREADLOCK.acquire()

            total += n # g(n/2, n/2) = 2*n/2 = n

            THREADLOCK.release()


def main():
    threads = []

    if THREADS == 1:
        threads.append(
            threading.Thread(target=G, args=(3, BOUND+1), name='t0')
        )

    elif THREADS > 1:
        threads.append(
            threading.Thread(target=G, args=(3, int(BOUND/THREADS)), name='t0')
        )

        for i in range(1, THREADS-1):
            threads.append(
                threading.Thread(target=G, args=(int(i * BOUND/THREADS), int((i+1) * BOUND/THREADS)), name='t{}'.format(i))
            )
        
        threads.append(
            threading.Thread(target=G, args=(int((THREADS-1) * BOUND/THREADS), BOUND+1), name='t{}'.format(THREADS))
        )

    for t in threads:
        t.start()

    for t in threads:
        t.join()

    # print(total % 1000000007)
    print(total)

if __name__ == "__main__":
    main()
