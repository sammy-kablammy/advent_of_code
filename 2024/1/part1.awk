# gawk gawk

{
    left_column[NR] = $1
    right_column[NR] = $2
    linecount++
}

END {
    asort(left_column)
    asort(right_column)
    for (i = 1; i <= linecount; i++) {
        distance = left_column[i] - right_column[i]
        if (distance < 0) {
            distance = 0 - distance
        }
        sum += distance
    }
    print sum
}
