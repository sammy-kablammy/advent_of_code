{
    left_column[NR] = $1
    frequency[$2] += 1
    linecount++
}

END {
    for (i in left_column) {
        num = left_column[i]
        sum += num * frequency[num]
    }
    print sum
}
