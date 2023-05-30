#!/bin/bash

PROGRAM=$(realpath "${PROGRAM:-$PROGRAM_DEFAULT}")

testFileToFile037() {
    printf 'Testing ebcdic 037 from file to utf8 file ... '

    rm -f /tmp/ebcdic037-out.dat

    $PROGRAM -i ebcdic037-in.dat -c 37 -o /tmp/ebcdic037-out.dat
    if [[ $? -ne 0 ]] ; then
        echo "ERROR: $PROGRAM returned error code $?"

        exit 1
    fi

    if cmp -s "/tmp/ebcdic037-out.dat" "./ebcdic037-out.dat"; then
        printf 'OK\n'
    else
        printf 'ERROR: result file does not match excepted result\n'

        exit 1
    fi
}

testFileToStdout037() {
    printf 'Testing ebcdic 037 from file to utf8 stdout ... '

    rm -f /tmp/ebcdic037-out.dat

    $PROGRAM -i ebcdic037-in.dat -c 037 > /tmp/ebcdic037-out.dat
    if [[ $? -ne 0 ]] ; then
        echo "ERROR: $PROGRAM returned error code $?"

        exit 1
    fi

    if cmp -s "/tmp/ebcdic037-out.dat" "./ebcdic037-out.dat"; then
        printf 'OK\n'
    else
        printf 'ERROR: result file does not match excepted result\n'

        exit 1
    fi
}

testStdinToFile037() {
    printf 'Testing ebcdic 037 from stdin to utf8 file ... '

    rm -f /tmp/ebcdic037-out.dat

    cat ./ebcdic037-in.dat | $PROGRAM -c EBCDIC037 -o /tmp/ebcdic037-out.dat
    if [[ $? -ne 0 ]] ; then
        echo "ERROR: $PROGRAM returned error code $?"

        exit 1
    fi

    if cmp -s "/tmp/ebcdic037-out.dat" "./ebcdic037-out.dat"; then
        printf 'OK\n'
    else
        printf 'ERROR: result file does not match excepted result\n'

        exit 1
    fi
}

main() {
    testFileToFile037

    testFileToStdout037

    testStdinToFile037
}

main
