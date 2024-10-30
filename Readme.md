## Transaction Pruner

This project aims to identify and replay transactions that are dependent on a set of initial contracts. The process involves dynamically updating the list of known contracts as new transactions are encountered, ensuring that all relevant transactions are included in the replay list.

## Algorithm

Explained in interview process already.

## How to run project

`make run`

## How to build project

`make build`

Note: Make sure you have go installed on your machine and after build binary, please provider execution permission to the binary file.
`chmod +x ./build/pruner`
