{
    target = int($1)
    for (permutation = 0; permutation < 2 ^ (NF - 2); permutation++) {
        total = $2
        # -1 because awk is 1-based
        # -1 to ignore the 'target' value at the very beginning of the line
        # -1 to ignore the initial value of 'total'
        # dear god
        for (i = 3; i <= NF; i++) {
            mask = lshift(1, i - 3)
            op = and(permutation, mask)
            if (op) {
                total = total + $i
            } else {
                total = total * $i
            }
        }
        # print total
        if (total == target) {
            final_sum += total
            break
        }
    }
}

END {
    print final_sum
}
