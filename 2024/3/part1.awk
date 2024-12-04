BEGIN {
    # by splitting lines at "mul", all valid multiplication statements must now
    # appear at the very beginning of the line AND at most one per line, making
    # the later regex matching much simpler
    RS = "mul"
}

/^\([0-9]{1,3},[0-9]{1,3}\)/ {
    if (FNR == 1) {
        # avoid off-by-one caused by RS shenanigans, wherein "(xxx,xxx)"
        # (missing the "mul" part) would count as a valid mul(a,b) statement.
        # (this case doesn't even apply to my input but it bothers me, okay?)
        next
    }
    comma = index($0, ",")
    paren = index($0, ")")
    first = substr($0, 2, comma - 2)
    second = substr($0, comma + 1, paren - comma - 1)
    sum += first * second
}

END {
    print sum
}
