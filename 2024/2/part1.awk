{
    # assume NF > 1
    is_increasing = 0
    if ($2 - $1 > 0) {
        is_increasing = 1
    }
    for (i = 1; i < NF; i++) {
        distance = $(i + 1) - $i
        # check distance requirement
        if (!(distance >= 1 && distance <= 3 || distance <= -1 && distance >= -3)) {
            next
        }
        # check increasing/decreasing requirement
        if (is_increasing && distance < 0 || !is_increasing && distance > 0) {
            next
        }
    }
    safe_count += 1
}

END {
    print safe_count
}
