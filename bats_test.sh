#!/usr/bin/env bats

setup() {
    go build -o goCipherTest .
}

setup

@test "successful execution of ceasar with all required valid arguments" {
    eval "run ./goCipherTest ceasar -m hello"

    [[ "$status" -eq 0 ]]
    [[ "$output" == *"rovvy"* ]]
}


@test "successful execution of vigenere with all required valid arguments" {
    eval "run ./goCipherTest vigenere -m hello -k test"

    [[ "$status" -eq 0 ]]
    [[ "$output" == *"aideh"* ]]
}

@test "exits without passing subcommand" {
    eval "run ./goCipherTest"

    [[ "$status" -eq 1 ]]
}

@test "returns without passing arguments to ceasar subcommand" {
    eval "run ./goCipherTest ceasar"

    [[ "$status" -eq 0 ]]
    [[ "$output" == *"Please make sure to pass a message"* ]]
}

@test "returns without passing arguments to vigenere subcommand" {
    eval "run ./goCipherTest vigenere"

    [[ "$status" -eq 0 ]]
    [[ "$output" == *"Please make sure to pass a message"* ]]
}

teardown() {
    rm goCipherTest
}