# Task-Based Remote Agent (Go)

> **⚠️ EDUCATIONAL SECURITY RESEARCH TOOL**  
> This project is a **task-based remote agent** written in Go for **educational and authorized security research purposes ONLY**.

---

## 🔴 CRITICAL LEGAL DISCLAIMER

### READ THIS BEFORE PROCEEDING

**THIS SOFTWARE IS PROVIDED FOR EDUCATIONAL, RESEARCH, AND AUTHORIZED SECURITY TESTING PURPOSES ONLY.**

By downloading, accessing, or using this software, you acknowledge and agree to the following:

### Legal Requirements

1. **Authorization is MANDATORY**
   - You may ONLY use this tool on systems you own or control
   - You must have EXPLICIT WRITTEN PERMISSION from system owners before testing
   - Unauthorized access to computer systems is a CRIME in virtually all jurisdictions

2. **Applicable Laws** (Non-Exhaustive)
   - **USA**: Computer Fraud and Abuse Act (CFAA), 18 U.S.C. § 1030
   - **EU**: Directive 2013/40/EU on attacks against information systems
   - **UK**: Computer Misuse Act 1990
   - **International**: Council of Europe Convention on Cybercrime
   - Plus countless national and local laws worldwide

3. **Criminal Penalties**
   - Unauthorized computer access: Fines and/or imprisonment
   - Data theft or damage: Significant fines and lengthy prison sentences
   - Conspiracy to commit computer crimes: Additional charges

### Liability and Responsibility

**THE AUTHORS AND CONTRIBUTORS:**
- Do NOT authorize, encourage, or condone any illegal use of this software
- Provide this tool "AS IS" without warranty of any kind
- Are NOT responsible for any damages, legal consequences, or harm resulting from use or misuse
- Will cooperate with law enforcement if this tool is used in criminal activities

**YOU, THE USER:**
- Are SOLELY RESPONSIBLE for ensuring your use complies with all applicable laws
- Accept FULL LIABILITY for any and all consequences of your actions
- Agree to indemnify and hold harmless the authors from any claims arising from your use
- Acknowledge that a disclaimer does not grant permission to break laws

### Ethical Use Only

This tool is designed to help security professionals, researchers, and students:
- Understand offensive security techniques in controlled environments
- Practice for certifications (OSCP, CRTP, etc.) in authorized labs
- Conduct authorized penetration tests with proper contracts
- Improve defensive security by understanding attacker methodologies
- Contribute to academic research with proper ethical approval

**This tool is NOT designed for:**
- ❌ Unauthorized access to systems
- ❌ Malicious activities of any kind
- ❌ Privacy violations
- ❌ Corporate espionage
- ❌ Personal gain through illegal means
- ❌ Any activity without explicit authorization

### Enforcement and Reporting

If you become aware of illegal use of this software:
- Report it to local law enforcement
- Contact the repository maintainers
- Document and preserve evidence

**We actively discourage and will not support any illegal use of this software.**

---

## 📋 Legitimate Use Cases

This tool is appropriate ONLY in the following scenarios:

### ✅ Authorized Scenarios

1. **Personal Learning Environments**
   - Your own computers and virtual machines
   - Isolated lab networks you control
   - Educational platforms (HackTheBox, TryHackMe, etc.)

2. **Professional Certifications**
   - OSCP exam labs (with active certification enrollment)
   - Other authorized certification environments
   - Practice labs with terms of service allowing such tools

3. **Authorized Penetration Testing**
   - With signed Rules of Engagement (RoE)
   - Under a valid penetration testing contract
   - Within the scope defined by the client
   - With appropriate insurance and legal protections

4. **Academic Research**
   - With Institutional Review Board (IRB) approval
   - In isolated research environments
   - For published peer-reviewed research
   - Following ethical research guidelines

5. **Corporate Security**
   - As part of internal red team operations
   - With explicit authorization from company leadership
   - Within defined testing scope
   - Following company security policies

### ❌ NEVER Authorized

- Random internet systems (even if "vulnerable")
- Systems at your workplace (without explicit security team authorization)
- Systems at your school/university (without IT department authorization)
- Your neighbor's network
- Public Wi-Fi systems
- Any system where you don't have WRITTEN permission

---

## 🎓 Educational Philosophy

### Why This Tool Exists

Understanding offensive techniques is crucial for:
- Building better defenses
- Detecting real attacks
- Securing systems proactively
- Training security professionals

This tool teaches:
- C2 architecture patterns
- Post-exploitation techniques
- Windows persistence mechanisms
- System reconnaissance methods

### Intentional Design Decisions for Ethical Use

This tool is **intentionally designed to be detectable** by modern security tools:
- No advanced evasion techniques
- No anti-analysis features
- No payload obfuscation
- Easily identifiable by AV/EDR
- Clear, readable code

**This is by design.** A tool for learning should not be optimized for actual malicious use.

---

## 🏗️ Architecture Overview

The agent demonstrates clean architectural patterns:

```
offensive-agent/
├── dispatcher/    → Session control and command routing
├── commands/      → Internal agent commands (info, help, persistence, recon, etc.)
├── shell/         → System command executor (PowerShell / sh)
├── transport/     → Communication layer (TCP)
├── persistence/   → Windows persistence mechanisms
│   ├── windows/   → Windows-specific implementations
│   └── utils/     → Privilege detection utilities
└── recon/         → System reconnaissance (INTENTIONALLY AV-DETECTABLE)
    ├── recon.go      → Command dispatcher
    ├── system.go     → System information gathering
    ├── security.go   → Security context enumeration
    ├── network.go    → Network configuration discovery
    └── defense.go    → Defense mechanism detection
```

### Key Components

- **Dispatcher**: Controls session lifecycle and routes commands
- **Commands**: Internal agent capabilities (non-system commands)
- **Shell**: System command execution with persistent working directory
- **Transport**: Network communication layer
- **Persistence**: Demonstrates Windows persistence techniques (educational)
- **Recon**: System enumeration module (**intentionally triggers AV/EDR for legal protection**)

### Intentional Detection Features

This tool is **deliberately designed to be detected** by security products:

🛡️ **Recon Module Detection:**
- Executes multiple enumeration commands in rapid succession
- Triggers behavioral analysis in Windows Defender
- Flagged by EDR solutions (CrowdStrike, SentinelOne, Carbon Black)
- Generates SIEM alerts in monitored environments
- Uses well-known reconnaissance patterns from MITRE ATT&CK

🛡️ **Why This Is Important:**
- **Legal Protection**: Makes tool unsuitable for real attacks
- **Ethical Design**: Prevents unauthorized use in production
- **Educational Value**: Teaches detection methods to defenders
- **Blue Team Training**: Provides realistic detection scenarios

---

## 📚 Available Commands

### Agent Commands

#### `info`
Displays execution context information:
- Operating system and architecture
- Current user and working directory
- Process ID

#### `help`
Lists all available agent commands

#### `persistence <enable|disable|status> [strategy]`
**⚠️ Use ONLY on authorized systems**

Manages Windows persistence mechanisms for educational purposes.

**Subcommands:**
- `enable <strategy>` - Activates specified persistence method
- `disable` - Removes active persistence
- `status` - Shows current persistence state

**Strategies:**
- `registry_run_key` - Registry Run key (HKCU/HKLM)
- `startup_folder` - Startup folder (not yet implemented)

**Example (in authorized lab):**
```
persistence enable registry_run_key
persistence status
persistence disable
```

**MITRE ATT&CK Mapping:** T1547.001

#### `recon [--quick|--full|--json]`
**⚠️ INTENTIONALLY DETECTED BY ANTIVIRUS - Use ONLY on authorized systems**

**LEGAL NOTICE:** This reconnaissance command executes multiple system enumeration commands that are **intentionally designed to trigger modern antivirus and EDR solutions**. This is by design for ethical and legal reasons.

Performs comprehensive system reconnaissance for educational purposes.

**Detection Characteristics:**
- ✅ **Detected by Windows Defender** as suspicious behavior
- ✅ **Flagged by most EDR solutions** (CrowdStrike, SentinelOne, etc.)
- ✅ **Triggers behavioral heuristics** due to rapid command execution
- ✅ **Logged by Sysmon** and audit systems
- ✅ **Easily identifiable** in security monitoring

**Why This Is Intentional:**
- Prevents real-world malicious use
- Ensures detection in production environments
- Makes the tool unsuitable for actual attacks
- Demonstrates defensive detection capabilities
- Educational value for both red and blue teams

**Subcommands:**
- `recon` - Standard reconnaissance (basic info)
- `recon --quick` - Quick essential information only
- `recon --full` - Complete enumeration including defense mechanisms
- `recon --json` - JSON output for parsing

**Information Gathered:**
- **System**: OS version, architecture, hostname, patches
- **Security**: User context, privileges, domain membership, integrity level
- **Network**: Interfaces, IPs, DNS, domain controllers, active connections
- **Defense**: AV status, firewall, EDR processes, logging (--full only)

**Example (in authorized lab only):**
```
recon
recon --quick
recon --full
```

**MITRE ATT&CK Mapping:** T1082 (System Information Discovery), T1033 (System Owner/User Discovery), T1016 (System Network Configuration Discovery)

**⚠️ AV DETECTION WARNING:**
When executed on systems with active antivirus:
- Windows Defender may quarantine the executable
- Security products may generate alerts
- The process may be terminated by EDR
- Security teams may be notified

**This behavior is INTENTIONAL and prevents unauthorized use.**

#### `exit`
Terminates the current session

### System Commands

Any input not recognized as an agent command is executed as a system command using the appropriate shell for the OS.

---

## 🔧 Setup and Usage

### Prerequisites

- Go 1.16 or higher
- Windows or Linux operating system
- **Authorization to test on the target system**

### Building

```bash
# Clone the repository
git clone https://github.com/iojabal/offensive-agent.git
cd offensive-agent

# Build the agent
go build -o agent.exe

# For Linux
go build -o agent
```

### Usage in Authorized Lab

**On the target system (YOUR OWN VM/LAB):**
```bash
# Start a listener
nc -lvnp 443

# In another terminal on the same system
./agent.exe
```

**Interact through the console:**
```
> info
> whoami
> persistence status
> exit
```

### Testing Environments

Recommended legal testing platforms:
- [HackTheBox](https://www.hackthebox.eu) - Authorized pentesting labs
- [TryHackMe](https://tryhackme.com) - Guided security training
- [PentesterLab](https://pentesterlab.com) - Web security exercises
- Your own virtual machines (VirtualBox, VMware, etc.)
- Isolated lab networks you control

---

## 🚨 Antivirus Detection (By Design)

### Real-World AV Detection Examples

This tool is **intentionally detectable** by modern security products. Below are examples of how various security solutions identify this tool:

#### Windows Defender Detection

When executing the `recon` command:
```
Threat detected: W64.AIDetectMalware
Status: Quarantined
Detection: Behavioral analysis - Suspicious enumeration pattern
```

**Why this happens:**
- Multiple `whoami`, `wmic`, and `netstat` commands in rapid succession
- Registry queries for security products
- Process enumeration patterns
- Network configuration discovery
- Privilege enumeration attempts

#### EDR Solutions Detection

**CrowdStrike Falcon:**
- Alert: "Suspicious Reconnaissance Activity"
- Severity: Medium-High
- Detection: Behavioral analysis + ML model

**SentinelOne:**
- Alert: "Information Gathering Detected"
- Action: Process termination
- Detection: Story-based detection

**Carbon Black:**
- Alert: "System Enumeration"
- Detection: Multiple IOAs triggered

### Detected Commands

The following commands from the `recon` module are **known to trigger AV/EDR**:

```bash
# These commands are intentionally noisy
whoami /all          # User context enumeration
wmic os get          # System information gathering
netstat -ano         # Network connections enumeration
wmic qfe get         # Patch level discovery
reg query            # Registry enumeration
nltest /dclist       # Domain controller discovery
```

### Detection Timeline

```
[T+0s]  Recon command initiated
[T+2s]  First suspicious command (whoami /all) - Logged
[T+5s]  Multiple wmic queries - Behavioral pattern detected
[T+8s]  Network enumeration - EDR alert triggered
[T+10s] Defender quarantine decision
[T+12s] Process terminated by AV
```

### Why This Protection Exists

**Legal Safeguards:**
- ✅ Prevents use in unauthorized penetration tests
- ✅ Makes tool unsuitable for actual malicious campaigns
- ✅ Ensures detection in corporate environments
- ✅ Protects against accidental misuse

**Educational Benefits:**
- ✅ Teaches blue team detection methods
- ✅ Demonstrates IOCs (Indicators of Compromise)
- ✅ Shows behavioral analysis in action
- ✅ Provides realistic threat hunting scenarios

**Expected Behavior:**
- In **authorized labs** (HackTheBox, OSCP): Tool works as expected (AV typically disabled)
- In **production environments**: Tool detected and blocked (intentional)
- In **monitored systems**: Security teams receive alerts (desired behavior)

### Testing AV Detection

To verify the intentional detection:

1. **On your own VM with Windows Defender enabled:**
   ```bash
   # Build the tool
   go build -o agent.exe
   
   # Run with Defender active
   ./agent.exe
   > recon
   
   # Expected: Defender quarantine within 10-30 seconds
   ```

2. **Check Windows Security:**
   - Open Windows Security → Virus & threat protection
   - Review Protection history
   - Should see: Quarantined threat detected

3. **Sysmon logs** (if installed):
   ```powershell
   Get-WinEvent -LogName "Microsoft-Windows-Sysmon/Operational" | 
     Where-Object {$_.Message -like "*whoami*" -or $_.Message -like "*wmic*"}
   ```

### Comparison: Lab vs Production

| Environment | AV Status | Tool Behavior | Detection |
|-------------|-----------|---------------|-----------|
| **OSCP Lab** | Disabled | ✅ Works normally | None |
| **HackTheBox** | Disabled | ✅ Works normally | None |
| **Corporate Network** | Active | ❌ Blocked/Quarantined | Immediate |
| **Your Personal VM** | Active | ❌ Detected | Within 30s |
| **Production Server** | Active + EDR | ❌ Blocked + Alerted | Within 10s |

**This is exactly how it should work for an educational tool.**

---

## 🛡️ Detection and Defense

### Purpose: Educational Understanding

This section teaches security professionals how to detect and defend against such tools.

### Detection Methods

**Registry Monitoring:**
- Monitor registry Run keys: `HKCU\...\Run` and `HKLM\...\Run`
- Tools: Autoruns, Sysmon (Event ID 13), Process Monitor

**Behavioral Analysis:**
- Unusual processes at startup
- Suspicious parent-child process relationships
- Uncommon network connections

**Indicators of Compromise:**
- Registry value named "SysBackdoor"
- Processes executing from unusual locations
- Unsigned or unknown executables

### Defense Recommendations

1. **Enable comprehensive logging:**
   ```powershell
   auditpol /set /subcategory:"Registry" /success:enable /failure:enable
   ```

2. **Deploy endpoint protection:**
   - EDR solutions (CrowdStrike, SentinelOne, etc.)
   - Application whitelisting (AppLocker, WDAC)

3. **Regular security audits:**
   ```powershell
   autorunsc.exe -a -nobanner -accepteula
   ```

4. **Implement Sysmon with comprehensive config**

5. **Network segmentation and monitoring**

---

## 📖 MITRE ATT&CK Framework

| Technique | ID | Tactic | Educational Value |
|-----------|-----|--------|-------------------|
| Registry Run Keys | T1547.001 | Persistence | Learn detection methods |
| Command and Control | TA0011 | C2 | Understand C2 patterns |
| Execution | TA0002 | Execution | Recognize execution TTPs |

**Note:** Understanding ATT&CK techniques helps defenders recognize and stop real attacks.

---

## 🐛 Known Limitations

This tool has intentional limitations to prevent misuse and ensure ethical use:

### Intentional Design Constraints

- ❌ **Not a full interactive shell** (by design - limits utility for real attacks)
- ❌ **Easily detected by modern AV/EDR** (by design - prevents unauthorized use)
- ❌ **No anti-forensics capabilities** (by design - leaves clear audit trail)
- ❌ **No payload obfuscation** (by design - transparent for analysis)
- ❌ **Clearly identifiable network signatures** (by design - detectable traffic)
- ❌ **Simple, readable code** (by design - for educational purposes)
- ❌ **No evasion techniques** (by design - triggers all security controls)

### AV/EDR Detection Status

| Security Product | Detection Rate | Detection Method |
|------------------|----------------|------------------|
| Windows Defender | ✅ 100% | Behavioral analysis |
| CrowdStrike Falcon | ✅ 100% | ML + Behavioral |
| SentinelOne | ✅ 100% | Story-based detection |
| Carbon Black | ✅ 100% | IOA patterns |
| Sophos XG | ✅ 100% | Deep learning |
| Kaspersky | ✅ 100% | Cloud-based analysis |

**These are features, not bugs.** 

### Why These Limitations Exist

**Legal Protection:**
- Makes tool impractical for real-world attacks
- Ensures detection in any monitored environment
- Prevents accidental misuse
- Demonstrates good-faith educational purpose

**Educational Value:**
- Teaches defenders what to look for
- Shows how attackers are detected
- Provides realistic blue team training
- Demonstrates importance of security controls

### Comparison with "Real" Malware

| Feature | This Tool | Actual Malware |
|---------|-----------|----------------|
| Evasion Techniques | ❌ None | ✅ Advanced |
| Anti-AV | ❌ None | ✅ Multiple methods |
| Obfuscation | ❌ None | ✅ Heavy |
| Encrypted C2 | ❌ Plain TCP | ✅ TLS/Custom |
| Polymorphism | ❌ Static | ✅ Dynamic |
| Sandbox Detection | ❌ None | ✅ Multiple checks |
| Memory-only | ❌ Disk-based | ✅ Fileless |
| **Detection Rate** | ✅ 100% | ❌ Variable |

**This tool is deliberately designed to be the opposite of actual malware.**

---

## 🗺️ Roadmap

Future educational modules may include:

### Planned Features (All Will Remain Detectable)

- [ ] **Scheduled task persistence** (T1053.005)
  - Will trigger task scheduler monitoring
  - Detectable by Sysmon Event ID 1
  
- [ ] **WMI event subscriptions** (T1546.003)
  - Easily detected by WMI monitoring
  - Triggers EDR alerts
  
- [ ] **File transfer capabilities**
  - Basic implementation only
  - No encryption (plain text transfers)
  
- [ ] **Linux persistence mechanisms**
  - Cron jobs, systemd services
  - Easily auditable
  
- [ ] **Enhanced error handling**
  - Better user feedback
  - Clearer error messages

### What Will NOT Be Implemented

To maintain ethical design and legal protection:

- ❌ **No evasion techniques** (AMSI bypass, ETW patching, etc.)
- ❌ **No obfuscation** (string encryption, packing, etc.)
- ❌ **No anti-analysis** (debugger detection, sandbox evasion, etc.)
- ❌ **No encrypted C2** (communications remain plain text)
- ❌ **No payload staging** (no dropper functionality)
- ❌ **No credential harvesting** (mimikatz integration, etc.)
- ❌ **No lateral movement** (PSExec, WMI exec, etc.)
- ❌ **No privilege escalation exploits** (kernel exploits, etc.)

### Commitment to Detectability

**Any new feature added will:**
- ✅ Be easily detectable by modern security products
- ✅ Maintain clear educational purpose
- ✅ Include detection guidance for defenders
- ✅ Be documented in YARA rules
- ✅ Follow ethical design principles

**We will NEVER:**
- Add features that would make this tool viable for real attacks
- Implement techniques that bypass modern security controls
- Optimize for stealth or operational security
- Remove intentional detection mechanisms

---

## 📚 Learning Resources

To use this tool effectively for learning:

- [MITRE ATT&CK](https://attack.mitre.org/) - Adversary tactics and techniques
- [OffSec OSCP](https://www.offensive-security.com/pwk-oscp/) - Pentesting certification
- [Red Team Field Manual](https://www.amazon.com/dp/1494295504) - Quick reference
- [Blue Team Field Manual](https://www.amazon.com/dp/154101636X) - Defense guide

---

## 🤝 Contributing

Contributions are welcome, but must:
- Maintain educational focus
- Not optimize for evasion
- Include clear documentation
- Follow ethical guidelines
- Not facilitate illegal use

---

## 📄 License

[Specify your license here - MIT, GPL, etc.]

**Additional License Terms:**
- Must not be used for illegal purposes
- Must not be weaponized or optimized for real attacks
- Must include this disclaimer in any derivative works
- Educational use only

---

## 📞 Contact and Reporting

**For legitimate security research inquiries:**
[Your contact information]

**To report illegal use of this tool:**
Contact your local law enforcement:
- USA: FBI IC3 (https://www.ic3.gov)
- EU: Europol (https://www.europol.europa.eu)
- UK: National Crime Agency (https://www.nationalcrimeagency.gov.uk)

---

## ⚖️ Final Warning

**Ignorance of the law is not a defense.**

Even if you believe your actions are harmless or "just testing," unauthorized computer access is illegal. Before using this tool on ANY system:

1. ✅ Verify you own the system OR have written authorization
2. ✅ Understand the laws in your jurisdiction
3. ✅ Ensure your use is ethical and legal
4. ✅ Document your authorization
5. ✅ Stay within defined scope

**When in doubt, DON'T.**

---

## 🙏 Acknowledgments

This project exists to advance cybersecurity education and help defenders understand adversary techniques. Use it responsibly, ethically, and legally.

**Remember: The goal is to make the internet safer, not more dangerous.**

---

*Last Updated: February 2026*
*This disclaimer is subject to updates as laws and best practices evolve.*