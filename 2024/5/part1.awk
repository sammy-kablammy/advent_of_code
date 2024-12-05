BEGIN {
    FS = ","
}

/\|/ {
    current_successor_idx = index($0, "|")
    first = substr($0, 1, current_successor_idx - 1)
    second = substr($0, current_successor_idx + 1)
    successors_of[first][successor_lengths[first]++] = second
}

/,/ {
    for (page_count = 1; page_count <= NF; page_count++) {
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
                next
            }
        }
    }

    # at this point, we know the current row is valid.
    middle_number = $(NF / 2 + 1)
    sum += middle_number
}

END {
    print "sum:", sum
}
