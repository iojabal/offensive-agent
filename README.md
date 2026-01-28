# Task-Based Remote Agent (Research)

This project is a **simple task-based remote command execution agent** written in Go.
It is designed for **educational and research purposes**, focusing on architecture clarity
rather than stealth or advanced evasion techniques.

## Overview

The agent connects to a remote TCP listener and allows controlled command execution
while maintaining session state (current directory, user context).

It does **not** implement a full interactive shell or TTY emulation.

## Features

- TCP-based transport
- OS-aware command execution (Windows / Unix)
- Logical prompt with persistent working directory
- Built-in handling for `cd` and `exit`
- Stateless command execution (one process per command)

## What this project is NOT

- Not a full interactive shell
- Not designed to evade EDR or antivirus solutions
- Not intended for unauthorized use
- No data exfiltration or surveillance features

## Architecture

- `main.go`  
  Initializes the connection and starts the agent.

- `transport/TCP.go`  
  Implements a simple TCP transport abstraction.

- `shell/shell.go`  
  Handles task-based command execution and session state.

## Usage (Controlled Environment)

1. Start a TCP listener:
   ```bash
   nc -lvnp 443
