BEGIN {
    FS = ","
}

/\|/ {
    current_successor_idx = index($0, "|")
    first = substr($0, 1, current_successor_idx - 1)
    second = substr($0, current_successor_idx + 1)
    successors_of[first][successor_lengths[first]++] = second
}

# okay yeah this is getting crazy
function sort_func(i1, v1, i2, v2) {
    if (v1 in successors_of) {
        for (succ in successors_of[v1]) {
            if (v2 == successors_of[v1][succ]) {
                return -1
            }
        }
    }
    if (v2 in successors_of) {
        for (succ in successors_of[v2]) {
            if (v1 == successors_of[v2][succ]) {
                return 1
            }
        }
    }
    return 0
}

/,/ {
    is_valid = 1
    for (page_count = 1; page_count <= NF && is_valid; page_count++) {
        current_page = $page_count
        current_page_idx = index($0, current_page)
        if (!(current_page in successors_of)) {
            continue
        }
        for (current_successor_idx in successors_of[current_page]) {
            current_successor = successors_of[current_page][current_successor_idx]
            current_successor_idx = index($0, current_successor)
            if (current_successor_idx == 0) {
                continue
            }
            if (current_page_idx > current_successor_idx) {
                is_valid = 0
                break
            }
        }
    }
    if (is_valid) {
        next
    }

    for (i = 1; i <= NF; i++) {
        row[i] = $i
    }
    asort(row, newrow, "sort_func")
    middle_number = newrow[int(NF / 2 + 1)]
    sum += middle_number
    # i have officially decided that awk is, in fact, not my beloved. this
    # "delete row" took like an hour to figure out. these auto-initialized
    # global variables are nuts
    delete row
}

END {
    print "sum:", sum
}
