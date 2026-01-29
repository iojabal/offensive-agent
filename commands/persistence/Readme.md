# Windows Persistence Module

A Go-based persistence module designed for educational purposes and authorized penetration testing engagements. This module implements common Windows persistence techniques.

## ⚠️ Legal Disclaimer

**THIS TOOL IS FOR EDUCATIONAL AND AUTHORIZED TESTING ONLY.**

- Only use this tool on systems you own or have explicit written permission to test
- Unauthorized access to computer systems is illegal in most jurisdictions
- The authors assume no liability for misuse or damage caused by this tool
- This tool is intended for security professionals, researchers, and students in controlled environments

##  Overview

This module provides a command-line interface to manage persistence mechanisms on Windows systems. It supports multiple persistence strategies and automatically adjusts techniques based on privilege level (user vs. administrator).

### Supported Persistence Techniques

| Strategy | Description | Privilege Required | Registry Location |
|----------|-------------|-------------------|-------------------|
| `registry_run_key` | Creates a registry Run key to execute on user login | User/Admin | `HKCU\Software\Microsoft\Windows\CurrentVersion\Run` (User)<br>`HKLM\Software\Microsoft\Windows\CurrentVersion\Run` (Admin) |
| `startup_folder` | Copies binary to Windows Startup folder | User | `%APPDATA%\Microsoft\Windows\Start Menu\Programs\Startup` |

##  Architecture

```
persistence/
├── persistence.go          # Main command handler
├── utils/
│   └── util.go            # Privilege escalation detection
└── windows/
    ├── windows.go         # Windows-specific persistence logic
    └── RegPersist.go      # Registry manipulation functions
```

##  Usage

### Basic Commands

#### Enable Persistence
```bash
persistence enable <strategy>
```

**Example:**
```bash
persistence enable registry_run_key
```

#### Disable Persistence
```bash
persistence disable
```

#### Check Status
```bash
persistence status
```

#### Get Help
```bash
persistence enable
```
This will display available persistence strategies.

### Command Workflow

1. **Check current status:**
   ```bash
   persistence status
   ```
   Output:
   ```
   Persistence Enabled: false
   Strategy: none
   ```

2. **Enable persistence with registry method:**
   ```bash
   persistence enable registry_run_key
   ```
   Output (User privileges):
   ```
   HKCU\Software\Microsoft\Windows\CurrentVersion\Run registry key created for persistence.
   ```
   Output (Admin privileges):
   ```
   HKLM\Software\Microsoft\Windows\CurrentVersion\Run registry key created for persistence.
   ```

3. **Verify persistence:**
   ```bash
   persistence status
   ```
   Output:
   ```
   Persistence Enabled: true
   Strategy: registry_run_key
   ```

4. **Disable persistence:**
   ```bash
   persistence disable
   ```

##  Technical Details

### Privilege Detection

The module automatically detects the current privilege level using the `IsElevated()` function:

- **Windows:** Checks if the current user is `SYSTEM`, `Administrator`, or `Administrador`
- **Unix-like:** Checks if effective UID is 0 (root)

Based on privilege level, the module selects the appropriate registry hive:
- **Elevated:** Uses `HKLM` (affects all users, requires admin)
- **Non-elevated:** Uses `HKCU` (affects current user only)

### Registry Run Key Method

**Location:**
- User: `HKEY_CURRENT_USER\Software\Microsoft\Windows\CurrentVersion\Run`
- Admin: `HKEY_LOCAL_MACHINE\Software\Microsoft\Windows\CurrentVersion\Run`

**Value Name:** `SysBackdoor`

**Value Data:** Full path to the executable

**Behavior:** The binary is executed automatically when:
- A user logs in (HKCU)
- Any user logs in (HKLM)

### Implementation Notes

1. **No Error Handling on Registry Ops:** The `exec.Command().Run()` calls do not check for errors. This is intentional for stealth but should be considered for production use.

2. **Executable Path Detection:** Uses `os.Executable()` to get the current binary path dynamically.

3. **Strategy Pattern:** The code uses a strategy pattern with a map to manage different persistence techniques.

##  Development

### Prerequisites

- Go 1.16 or higher
- Windows operating system (for testing)
- Administrator privileges (for HKLM testing)

### Building

```bash
go build -o persistence.exe
```

### Testing

**Test as regular user:**
```bash
.\persistence.exe
> persistence enable registry_run_key
> persistence status
```

**Test as administrator:**
```powershell
# Run PowerShell as Administrator
.\persistence.exe
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

##  Detection and Defense

### How to Detect This Technique

1. **Registry Monitoring:**
   - Monitor `HKCU\Software\Microsoft\Windows\CurrentVersion\Run`
   - Monitor `HKLM\Software\Microsoft\Windows\CurrentVersion\Run`
   - Tools: Autoruns, Process Monitor, Sysmon (Event ID 13)

2. **Startup Folder Monitoring:**
   - Monitor `%APPDATA%\Microsoft\Windows\Start Menu\Programs\Startup`
   - Look for unexpected executables

3. **Behavioral Indicators:**
   - Unusual processes starting at login
   - Unknown binaries with suspicious names (e.g., "SysBackdoor")
   - Processes executing from unusual locations

### Defense Recommendations

1. **Enable audit policies:**
   ```powershell
   auditpol /set /subcategory:"Registry" /success:enable /failure:enable
   ```

2. **Use EDR/AV solutions** that monitor registry changes

3. **Implement application whitelisting** (AppLocker/WDAC)

4. **Regular autoruns audits:**
   ```powershell
   autorunsc.exe -a -nobanner -accepteula
   ```

5. **Monitor with Sysmon:**
   - Event ID 12: Registry object created/deleted
   - Event ID 13: Registry value set

##  Educational Use Cases

This module is designed for:

- **OSCP Preparation:** Understanding Windows persistence mechanisms
- **Red Team Training:** Practicing post-exploitation techniques
- **Blue Team Training:** Learning to detect and remove persistence
- **CTF Competitions:** Authorized capture-the-flag events
- **Security Research:** Analyzing persistence techniques

##  MITRE ATT&CK Mapping

| Technique | ID | Tactic |
|-----------|-----|--------|
| Registry Run Keys / Startup Folder | T1547.001 | Persistence, Privilege Escalation |
| Boot or Logon Autostart Execution | T1547 | Persistence, Privilege Escalation |

**References:**
- [MITRE ATT&CK T1547.001](https://attack.mitre.org/techniques/T1547/001/)

##  Known Limitations

1. **Startup Folder Method:** Not yet implemented
2. **Error Handling:** Registry operations don't return error status
3. **Stealth:** The value name "SysBackdoor" is easily detectable
4. **Removal:** Only removes the specific strategy that was enabled
5. **Windows Only:** Unix-like system persistence not implemented

##  Future Enhancements

- [ ] Implement startup folder persistence
- [ ] Add scheduled task persistence
- [ ] Add WMI event subscription persistence
- [ ] Add service persistence
- [ ] Implement error handling for registry operations
- [ ] Add stealth improvements (randomized names, obfuscation)
- [ ] Cross-platform persistence (Linux, macOS)
- [ ] Persistence verification checks
- [ ] Logging and event tracking

##  References

- [Windows Persistence Techniques](https://attack.mitre.org/tactics/TA0003/)
- [Autoruns for Windows](https://docs.microsoft.com/en-us/sysinternals/downloads/autoruns)
- [Sysmon Configuration](https://github.com/SwiftOnSecurity/sysmon-config)

##  Contributing

This is an educational project. If you have suggestions or improvements:

1. Test thoroughly in isolated environments
2. Document any new techniques
3. Include detection/defense information
4. Follow responsible disclosure practices

##  License

This tool is provided as-is for educational purposes. Use responsibly and ethically.

---

**Remember:** With great power comes great responsibility. Always obtain proper authorization before testing security controls.