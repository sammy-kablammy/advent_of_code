# i'm now beginning to regret choosing awk

function is_safe(record, nf,  i) {
    is_increasing = 0
    if (record[2] - record[1] > 0) {
        is_increasing = 1
    }
    for (i = 1; i < nf; i++) {
        distance = record[i + 1] - record[i]
        # check distance requirement
        if (!(distance >= 1 && distance <= 3 || distance <= -1 && distance >= -3)) {
            return 0
        }
        # check increasing/decreasing requirement
        if (is_increasing && distance < 0 || !is_increasing && distance > 0) {
            return 0
        }
    }
    return 1
}

{
    # employ the Problem Dampener™
    for (i = 1; i <= NF; i++) {
        for (j = 1; j < i; j++) {
            record[j] = $j
        }
        for (j = i + 1; j <= NF; j++) {
            record[j - 1] = $j
        }
        safe = is_safe(record, NF - 1)
        if (safe) {
            safe_count += 1
            next
        }
    }
}

END {
    print safe_count
}
