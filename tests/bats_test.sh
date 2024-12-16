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

@test "successful execution of playfair with all required valid arguments" {
    eval "run ./goCipherTest playfair -m hello -k key"

    [[ "$status" -eq 0 ]]
    [[ "$output" == *"dbnvmi"* ]]
}

@test "returns without passing message to playfair subcommand" {
    eval "run ./goCipherTest playfair"

    [[ "$status" -eq 0 ]]
    [[ "$output" == *"Please make sure to pass a message"* ]]
}

@test "returns without passing key to playfair subcommand" {
    eval "run ./goCipherTest playfair -m hello"

    [[ "$status" -eq 0 ]]
    [[ "$output" == *"Please make sure to pass a key"* ]]
}

@test "returns without passing key to sbx subcommand" {
    eval "run ./goCipherTest sbx -m 61747461636b61746461776e"

    [[ "$status" -eq 0 ]]
    [[ "$output" == *"Please make sure to pass a key"* ]]
}

@test "returns without passing valid one-byte int key to sbx subcommand" {
    eval "run ./goCipherTest sbx -m 61747461636b61746461776e -k a"

    [[ "$status" -eq 1 ]]
    [[ "$output" == *"parse error"* ]]
}

@test "successful execution of sbx with all required valid arguments" {
    eval "run ./goCipherTest sbx -m 68656c6c6f -k 54"

    [[ "$status" -eq 0 ]]
    [[ "$output" == *"5e535a5a59"* ]]
}

teardown() {
    rm goCipherTest
}