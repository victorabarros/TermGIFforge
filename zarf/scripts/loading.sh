#!/bin/bash

function loading_animation {
    local pid=$1
    local delay=0.75
    local spin='/-\|'
    local i=0

    while kill -0 $pid 2>/dev/null; do
        i=$(( (i+1) % 4 ))
        printf "\rProcessing... ${spin:i:1}"
        sleep $delay
    done
    printf "\rDone!       \n"
}

long_running_command() {
    sleep 10  # Simulating a long process here
}

long_running_command &
loading_animation $!


sleep 10 & pid=$!; spin='/-\|'; i=0; while kill -0 $pid 2>/dev/null; do i=$(( (i+1) % 4 )); printf "\rProcessing... ${spin:i:1}"; sleep 0.75; done; printf "\rDone!       \n" & pid=$!; spin='/-\|'; i=0; while kill -0 $pid 2>/dev/null; do i=$(( (i+1) % 4 )); printf "\rProcessing... ${spin:i:1}"; sleep 0.75; done; printf "\rDone!       \n"