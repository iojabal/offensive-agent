📄 README.md
# Task-Based Remote Agent (Go)

This project is a **task-based remote agent** written in Go, designed for
**educational and research purposes**.  
It focuses on **clean architecture**, command dispatching, and execution flow,
rather than stealth or advanced evasion techniques.

The agent exposes a minimal interactive console that allows an operator to
execute system commands and internal agent commands in a controlled manner.

---

## Key Characteristics

- Task-based command execution (not a full interactive shell)
- Clear separation of responsibilities:
  - Dispatcher (session control)
  - Internal commands
  - System command executor
- OS-aware execution (Windows / Unix-like systems)
- Persistent working directory (`cd` handled internally)
- Simple and readable prompt

---

##  Architecture Overview

The agent is structured into clearly defined modules:



dispatcher/ → Session control and command routing
commands/ → Internal agent commands (info, help, etc.)
shell/ → System command executor (PowerShell / sh)
transport/ → Communication layer (TCP)


### Dispatcher
- Controls the session lifecycle
- Prints the prompt
- Routes input to:
  - internal agent commands
  - system command execution
- Handles `exit`

### Commands
- Internal agent commands
- Do not execute system commands
- Do not manage connections

### Shell
- Executes system commands
- Maintains execution context (current working directory)
- Does not control session flow or UI

---

##  Available Commands

### Agent Commands

- `info`  
  Displays basic execution context information:
  - OS
  - Architecture
  - User
  - Current working directory
  - Process ID

- `help`  
  Lists available agent commands.

- `exit`  
  Terminates the current session.

### System Commands

Any input that is **not** an internal agent command is treated as a system
command and executed using the appropriate shell for the operating system.

Examples:


dir
ls -la
whoami
ipconfig


---

##  Usage (Controlled Environment)

1. Start a TCP listener:
   ```bash
   nc -lvnp 443


Run the agent on the target system.

Interact with the agent through the task-based console.

### Disclaimer

This project is intended strictly for educational purposes and authorized
testing environments.
The author does not take responsibility for misuse of this software.

### Notes

This is not a full interactive TTY or shell emulator.

No persistence is enabled by default.

The project prioritizes architectural clarity over feature completeness.