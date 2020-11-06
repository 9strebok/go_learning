package recursive_fact

func recursive_fact(n uint64) uint64 {
    if n <= 1 {
        return 1
    }
    return n * recursive_fact(n - 1)
}

