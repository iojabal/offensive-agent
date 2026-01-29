 README.md
# Task-Based Remote Agent (Go)

This project is a **task-based remote agent** written in Go, designed for
**educational and research purposes**.  
It focuses on **clean architecture**, command dispatching, and execution flow,
rather than stealth or advanced evasion techniques.

The agent exposes a minimal interactive console that allows an operator to
execute system commands and internal agent commands in a controlled manner.

---

## ⚠️ Legal Disclaimer

**THIS TOOL IS FOR EDUCATIONAL AND AUTHORIZED TESTING ONLY.**

- Only use this tool on systems you own or have explicit written permission to test
- Unauthorized access to computer systems is illegal in most jurisdictions
- The authors assume no liability for misuse or damage caused by this tool
- This tool is intended for security professionals, researchers, and students in controlled environments

---

## Key Characteristics

- Task-based command execution (not a full interactive shell)
- Clear separation of responsibilities:
  - Dispatcher (session control)
  - Internal commands
  - System command executor
  - Persistence mechanisms
- OS-aware execution (Windows / Unix-like systems)
- Persistent working directory (`cd` handled internally)
- Windows persistence capabilities
- Simple and readable prompt

---

##  Architecture Overview

The agent is structured into clearly defined modules:

```
dispatcher/    → Session control and command routing
commands/      → Internal agent commands (info, help, persistence, etc.)
shell/         → System command executor (PowerShell / sh)
transport/     → Communication layer (TCP)
persistence/   → Windows persistence mechanisms (registry, startup)
  ├── windows/      → Windows-specific implementations
  │   ├── RegPersist.go  → Registry Run key manipulation
  │   └── windows.go     → Strategy pattern and dispatcher
  └── utils/        → Privilege detection utilities
```

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
- Includes persistence management

### Shell
- Executes system commands
- Maintains execution context (current working directory)
- Does not control session flow or UI

### Persistence Module
- Manages Windows persistence mechanisms
- Automatically detects privilege level (User/Admin)
- Supports multiple persistence strategies
- Provides enable/disable/status functionality

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

- `persistence <enable|disable|status> [strategy]`  
  Manages persistence mechanisms on the target system.
  
  **Subcommands:**
  - `enable <strategy>` - Activates specified persistence method
  - `disable` - Removes active persistence
  - `status` - Shows current persistence state
  
  **Available Strategies:**
  - `registry_run_key` - Registry Run key (HKCU/HKLM)
  - `startup_folder` - Windows Startup folder (not yet implemented)
  
  **Examples:**
  ```
  persistence enable registry_run_key
  persistence status
  persistence disable
  ```

- `exit`  
  Terminates the current session.

### System Commands

Any input that is **not** an internal agent command is treated as a system
command and executed using the appropriate shell for the operating system.

Examples:

```
- dir
- ls -la
- whoami
- ipconfig
```

---

##  Persistence Mechanisms

### Registry Run Key (registry_run_key)

**Description:**  
Creates a registry entry to execute the agent on user login.

**Privilege Detection:**
- **User privileges:** `HKEY_CURRENT_USER\Software\Microsoft\Windows\CurrentVersion\Run`
- **Admin privileges:** `HKEY_LOCAL_MACHINE\Software\Microsoft\Windows\CurrentVersion\Run`

**Registry Value:**
- **Name:** `SysBackdoor`
- **Data:** Full path to the agent executable

**Usage:**
```bash
persistence enable registry_run_key
```

**Output (User):**
```
HKCU\Software\Microsoft\Windows\CurrentVersion\Run registry key created for persistence.
```

**Output (Admin):**
```
HKLM\Software\Microsoft\Windows\CurrentVersion\Run registry key created for persistence.
```

**MITRE ATT&CK:** T1547.001 - Registry Run Keys / Startup Folder

### Startup Folder (startup_folder)

**Status:** Not yet implemented

**Planned Location:** `%APPDATA%\Microsoft\Windows\Start Menu\Programs\Startup`

---

##  Usage (Controlled Environment)

### Basic Setup

1. **Start a TCP listener:**
   ```bash
   nc -lvnp 443
   ```

2. **Run the agent on the target system**

3. **Interact with the agent through the task-based console**

### Persistence Workflow

1. **Check current persistence status:**
   ```
   persistence status
   ```

2. **Enable persistence:**
   ```
   persistence enable registry_run_key
   ```

3. **Verify persistence is active:**
   ```
   persistence status
   ```
   Output:
   ```
   Persistence Enabled: true
   Strategy: registry_run_key
   ```

4. **Test persistence:**
   - Reboot the system or log off/log in
   - Verify agent reconnects automatically

5. **Disable persistence when done:**
   ```
   persistence disable
   ```

---

##  Detection and Defense

### How to Detect

**Registry Monitoring:**
- Monitor `HKCU\Software\Microsoft\Windows\CurrentVersion\Run`
- Monitor `HKLM\Software\Microsoft\Windows\CurrentVersion\Run`
- Tools: Autoruns, Process Monitor, Sysmon (Event ID 13)

**Behavioral Indicators:**
- Unknown processes starting at login
- Suspicious registry value names (e.g., "SysBackdoor")
- Processes executing from unusual locations

### Defense Recommendations

1. **Enable registry audit policies:**
   ```powershell
   auditpol /set /subcategory:"Registry" /success:enable /failure:enable
   ```

2. **Use EDR/AV solutions** that monitor registry changes

3. **Regular autoruns audits:**
   ```powershell
   autorunsc.exe -a -nobanner -accepteula
   ```

4. **Implement Sysmon** for detailed logging:
   - Event ID 12: Registry object created/deleted
   - Event ID 13: Registry value set

---

##  Building and Development

### Prerequisites

- Go 1.16 or higher
- Windows operating system (for persistence features)

### Build

```bash
go build -o agent.exe
```

### Testing Persistence

**Test as regular user:**
```bash
.\agent.exe
> persistence enable registry_run_key
> persistence status
```

**Test as administrator:**
```powershell
# Run PowerShell as Administrator
.\agent.exe
> persistence enable registry_run_key
> persistence status
```

**Verify registry entries:**
```powershell
# Check HKCU
reg query "HKCU\Software\Microsoft\Windows\CurrentVersion\Run" /v SysBackdoor

# Check HKLM (requires admin)
reg query "HKLM\Software\Microsoft\Windows\CurrentVersion\Run" /v SysBackdoor
```

---

##  Educational Use Cases

This agent is designed for:

- **OSCP Preparation:** Understanding post-exploitation and persistence
- **Red Team Training:** Practicing agent deployment and persistence
- **Blue Team Training:** Learning to detect and remove persistence
- **CTF Competitions:** Authorized capture-the-flag events
- **Security Research:** Analyzing C2 architectures and persistence techniques

---

##  MITRE ATT&CK Mapping

| Technique | ID | Tactic |
|-----------|-----|--------|
| Registry Run Keys / Startup Folder | T1547.001 | Persistence, Privilege Escalation |
| Boot or Logon Autostart Execution | T1547 | Persistence, Privilege Escalation |
| Command and Control | TA0011 | Command and Control |

---

##  Known Limitations

- This is not a full interactive TTY or shell emulator
- No persistence is enabled by default (must be explicitly activated)
- The project prioritizes architectural clarity over feature completeness
- Startup folder persistence not yet implemented
- Registry operations don't return error status (intentional for stealth)
- The value name "SysBackdoor" is easily detectable
- Windows-only persistence (Unix-like systems not supported)

---

##  Roadmap

- [ ] Implement startup folder persistence
- [ ] Add scheduled task persistence
- [ ] Add WMI event subscription persistence
- [ ] Implement error handling for registry operations
- [ ] Add stealth improvements (randomized names)
- [ ] Cross-platform persistence (Linux, macOS)
- [ ] Persistence verification checks
- [ ] Advanced evasion techniques

---

##  References

- [MITRE ATT&CK - Persistence](https://attack.mitre.org/tactics/TA0003/)
- [Windows Autoruns](https://docs.microsoft.com/en-us/sysinternals/downloads/autoruns)
- [Sysmon Configuration](https://github.com/SwiftOnSecurity/sysmon-config)

---

## Disclaimer

This project is intended strictly for educational purposes and authorized
testing environments.
The author does not take responsibility for misuse of this software.

**Remember:** With great power comes great responsibility. Always obtain proper authorization before testing security controls.